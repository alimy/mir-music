FROM golang:alpine AS binaryBuilder
# Install build deps
RUN apk --no-cache --no-progress add --virtual build-deps build-base git
WORKDIR /go/src/github.com/alimy/mir-music
COPY . .
RUN export GO111MODULE=on && make build

FROM alpine:latest
# Install system utils & Mir-Music runtime dependencies
ADD https://github.com/tianon/gosu/releases/download/1.11/gosu-amd64 /usr/sbin/gosu
RUN chmod +x /usr/sbin/gosu \
  && echo http://dl-2.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories \
  && apk --no-cache --no-progress add \
    bash \
    shadow \
    s6

ENV GINMUSIC_CUSTOM /data/mirmusic

# Configure LibC Name Service
COPY hack/docker/nsswitch.conf /etc/nsswitch.conf

WORKDIR /app/mirmusic
COPY hack/docker ./docker
COPY --from=binaryBuilder /go/src/github.com/alimy/mir-music/mir-music .

RUN ./docker/finalize.sh

# Configure Docker Container
VOLUME ["/data"]
EXPOSE 8013
ENTRYPOINT ["/app/mirmusic/docker/start.sh"]
CMD ["/bin/s6-svscan", "/app/mirmusic/docker/s6/"]