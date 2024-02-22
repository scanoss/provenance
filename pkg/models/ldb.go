package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/google/uuid"
)

type ProvenanceItem struct {
	Country string
	Source  string
}
type ProvenanceResult struct {
	Provenance []ProvenanceItem
	Purl       string
	Version    string
}

type UrlItem struct {
	UrlHash  string
	PurlName string
	Version  string
}

type PivotItem struct {
	UrlHash  string
	FileHash string
}

type PurlItem struct {
	PurlHash string
	Date1    string
	Date2    string
	Date3    string
	Int1     string
	Int2     string
	Int3     string
	Country  string
}

var LDBProvenanceTableName string
var LDBPivotTableName string
var LDBBinPath string
var LDBEncBinPath string

// Checks if the LBD exists and returns the list of available tables
func PingLDB(ldbname string) ([]string, error) {
	var ret []string
	entry, err := os.ReadDir("/var/lib/ldb/" + ldbname)
	if err != nil {
		return []string{}, errors.New("Problems opening LDB " + ldbname)
	}
	for e := range entry {
		if entry[e].IsDir() {
			ret = append(ret, entry[e].Name())
		}
	}

	return ret, nil
}

// Single item worker for Provenance. From a MD5 of a file enqueues a list of ProvenanceItem

func QueryBulkPivotLDB(keys []string) (map[string][]string, error) {
	ret := make(map[string][]string)
	name := fmt.Sprintf("/tmp/%s-pivot.txt", uuid.New().String())
	f, err := os.Create(name)
	if err != nil {
		return map[string][]string{}, err
	}
	for job := range keys {
		if keys[job] != "" {
			line := fmt.Sprintf("select from %s key %s csv hex 32\n", LDBPivotTableName, keys[job])
			f.WriteString(line)
		}
	}
	f.Close()
	_, err = os.Stat(LDBBinPath)
	if os.IsNotExist(err) {

		return map[string][]string{}, errors.New("LDB console not found")
	}

	ldbCmd := exec.Command(LDBBinPath, "-f", name)

	buffer, errLDB := ldbCmd.Output()
	fmt.Println(errLDB)

	//split results line by line
	//each row contains 3 values: <UrlMD5>,<FileMD5>,unknown
	lines := strings.Split(string(buffer), "\n")

	for i := range lines {
		fields := strings.Split(lines[i], ",")
		if len(fields) == 3 {
			ret[fields[0]] = append(ret[fields[0]], fields[1])
		}
	}
	os.Remove(name)
	return ret, nil
}

func QueryBulkProvenanceLDB(items map[string][]string) map[string][]ProvenanceItem {
	countries := make(map[string][]ProvenanceItem)
	name := fmt.Sprintf("/tmp/%s-provenance.txt", uuid.New().String())
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return map[string][]ProvenanceItem{}
	}
	added := make(map[string]bool)
	for job := range items {
		fileHashes := items[job]
		for r := range fileHashes {
			if _, exist := added[fileHashes[r]]; !exist {
				line := fmt.Sprintf("select from %s key %s csv hex 16\n", LDBProvenanceTableName, fileHashes[r])
				f.WriteString(line)
				added[fileHashes[r]] = true
			}
		}
	}
	f.Close()

	ldbCmd := exec.Command(LDBBinPath, "-f", name)
	buffer, _ := ldbCmd.Output()
	lines := strings.Split(string(buffer), "\n")
	for i := range lines {
		fields := strings.Split(lines[i], ",")
		if len(fields) == 3 {
			source := ""
			if fields[1] == "0" {
				source = "URL or e-mail found on code"
			} else if fields[1] == "1" {
				source = "Country declared on code"
			} else if fields[1] == "2" {
				source = "Contributor profile declared"
			}

			country := ProvenanceItem{Source: source, Country: fields[2]}
			countries[fields[0]] = append(countries[fields[0]], country)
		}
		/*if len(fields) == 2 {
			country := ProvenanceItem{Source: "URL or email in file", Country: fields[1]}
			countries[fields[0]] = append(countries[fields[0]], country)
		}*/
	}
	//os.Remove(name)
	return countries
}

func QueryBulkPurlLDB(items []InternalQuery) map[string][]PurlItem {
	purls := make(map[string][]PurlItem)
	name := fmt.Sprintf("/tmp/%s-purl.txt", uuid.New().String())
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return map[string][]PurlItem{}
	}
	added := make(map[string]string)
	for job := range items {
		fileHash := fmt.Sprintf("%016x", md5.Sum([]byte(items[job].CompletePurl)))
		if _, exist := added[fileHash]; !exist {
			line := fmt.Sprintf("select from oss/purl key %s csv hex 16\n", fileHash)
			f.WriteString(line)
			added[fileHash] = items[job].CompletePurl
		}

	}
	f.Close()

	ldbCmd := exec.Command(LDBEncBinPath, "-f", name)
	buffer, _ := ldbCmd.Output()
	lines := strings.Split(string(buffer), "\n")
	for i := range lines {
		fields := strings.Split(lines[i], ",")
		/*	if len(fields) == 3 {
			country := ProvenanceItem{Source: fields[1], Country: fields[2]}
			countries[fields[0]] = append(countries[fields[0]], country)
		}*/
		if len(fields) == 8 {
			purlItem := PurlItem{PurlHash: fields[0], Date1: fields[1], Date2: fields[2], Date3: fields[3], Int1: fields[4], Int2: fields[5], Int3: fields[6], Country: fields[7]}
			purls[added[purlItem.PurlHash]] = append(purls[added[purlItem.PurlHash]], purlItem)
		}
	}
	//os.Remove(name)
	return purls
}

func ContainsTable(arr []string, value string) bool {
	for r := range arr {
		if arr[r] == value {
			return true
		}
	}
	return false

}
