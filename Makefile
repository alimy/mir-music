.PHONY: default build serve generate api fmt clean distclean

TAGS = release portal
ASSETS_DATA_FILES := $(shell find assets | sed 's/  /\\ /g')

LDFLAGS += -X "github.com/alimy/mir-music/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "github.com/alimy/mir-music/version.GitHash=$(shell git rev-parse HEAD)"

default: build

build: fmt generate
	go build -ldflags '$(LDFLAGS)' -tags '$(TAGS)' -o mir-music

serve: fmt generate
	go run -ldflags '$(LDFLAGS)' -tags '$(TAGS)' github.com/alimy/mir-music serve --debug --addr :8013

fmt:
	go fmt ./...

generate: $(ASSETS_DATA_FILES)
	-rm -f pkg/assets/assets_gen.go
	go generate pkg/assets/assets.go
	gofmt -s -w pkg/assets

api:
	docker run -it --rm -p 8080:80 -v $(PWD)/api/openapi.yaml:/usr/share/nginx/html/openapi.yaml -e SPEC_URL=openapi.yaml redocly/redoc

clean:
	go clean -r ./...
	-rm -f mir-music

distclean: clean
	-rm -f pkg/assets/assets_gen.go