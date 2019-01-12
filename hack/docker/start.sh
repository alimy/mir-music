#!/bin/sh

create_volume_subfolder() {
    # Create VOLUME subfolder
    for f in /data/mirmusic/data /data/mirmusic/conf /data/mirmusic/log; do
        if ! test -d $f; then
            mkdir -p $f
        fi
    done
}

setids() {
    PUID=${PUID:-1000}
    PGID=${PGID:-1000}
    groupmod -o -g "$PGID" mirmusic
    usermod -o -u "$PUID" mirmusic
}

setids
create_volume_subfolder

# Exec CMD or S6 by default if nothing present
if [ $# -gt 0 ];then
    exec "$@"
else
    exec /bin/s6-svscan /app/mirmusic/docker/s6/
fi
