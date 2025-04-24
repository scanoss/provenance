#!/bin/bash

##########################################
#
# This script is designed to run by Systemd SCANOSS Geo Provenance API service.
# It rotates scanoss log file and starts SCANOSS Geo Provenance API.
# Install it in /usr/local/bin
#
################################################################
DEFAULT_ENV="prod"
ENVIRONMENT="${1:-$DEFAULT_ENV}"
LOGFILE=/var/log/scanoss/geoprovenance/scanoss-geoprovenance-api-$ENVIRONMENT.log
CONF_FILE=/usr/local/etc/scanoss/geoprovenance/app-config-${ENVIRONMENT}.json

# Rotate log
if [ -f "$LOGFILE" ] ; then
  echo "rotating logfile..."
  TIMESTAMP=$(date '+%Y%m%d-%H%M%S')
  BACKUP_FILE=$LOGFILE.$TIMESTAMP
  cp "$LOGFILE" "$BACKUP_FILE"
  gzip -f "$BACKUP_FILE"
fi
echo > "$LOGFILE"

echo > $LOGFILE
# Start scanoss-geoprovenance-api
echo "starting SCANOSS Geo Provenance API"
exec /usr/local/bin/scanoss-geoprovenance-api --json-config "$CONF_FILE" > "$LOGFILE" 2>&1