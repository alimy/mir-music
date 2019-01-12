#!/bin/sh
set -x
set -e

# Create user for Gin
addgroup -S mirmusic
adduser -G mirmusic -H -D -g 'mirmusic User' mirmusic -h /data/mirmusic -s /bin/bash && usermod -p '*' mirmusic && passwd -u mirmusic
echo "export MIRMUSIC_CUSTOM=${MIRMUSIC_CUSTOM}" >> /etc/profile

# Final cleaning
rm /app/mirmusic/docker/finalize.sh
rm /app/mirmusic/docker/nsswitch.conf
