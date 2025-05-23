# SCANOSS Provenance Service Deployment Support
This folder contains convenience utilities for deploying, configuring and running the SCANOSS provenance API service.

## Setup
The [scripts](.) folder contains an [env_setup.sh](env-setup.sh) script which attempts to do the following:
* Set up the default folders
* Set permissions
* Set up service registration ([scanoss-geoprovenance-api.service](scanoss-geoprovenance-api.service))
* Copy in binaries (if `scanoss-geoprovenance-api` exist in the folder)
* Copy in preferred configuration (if `app-config-prod.json` exists in the folder)

A sample development configuration file can be found in [config/app-config-dev.json](../config/app-config-dev.json).

Logs are written by default to `/var/log/scanoss/geoprovenance/scanoss-geoprovenance-api-prod.log`.

Configuration is written by default to: `/usr/local/etc/scanoss/geoprovenance`.

## Installation
Running the [env_setup.sh](env-setup.sh) on the target server, takes care of installation. Simply run:
```shell
./env_setup.sh
```

This will copy the configuration files to `/usr/local/etc/scanoss/geoprovenance` the binaries to `/usr/local/bin` and the service registration to `/etc/systemd/system`. 

It will also redirect logging to `/var/log/scanoss/geoprovenance`.

### Multi-service Registration
If there is a need to deploy more than one API service on the same server, this can be achieved by using a different `ENVIRONMENT` name.

Create a copy of the [scanoss-geoprovenance-api.service](scanoss-geoprovenance-api.service) using the following command:
```shell
cp scanoss-geoprovenance-api.service scanoss-geoprovenance-api-<env>.service
```

Where `env` is the name of this edition of the service (i.e. dev).

The `app-config-prod.json` file will also need to be copied:
```shell
cp app-config-prod.json app-config-<env>.json
```
**Note:** Please remember to use a different port number.

Finally, run the environment setup script using:
```shell
./env_setup.sh <env>
```

This will search for these specific service & config files and place them in the correct location.

Details for starting/stopping the service will be displayed in the console at the end of installation.