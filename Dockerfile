FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go generate ./pkg/cmd/server.go
RUN go build -o ./scanoss-geoprovenance ./cmd/server

FROM debian:buster-slim

WORKDIR /app
 
COPY --from=build /app/scanoss-geoprovenance /app/scanoss-geoprovenance

EXPOSE 50056

ENTRYPOINT ["./scanoss-geoprovenance"]
#CMD ["--help"]
