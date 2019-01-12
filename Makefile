.PHONY: default build serve bindata api fmt clean distclean

TAGS = release portal
PORTAL_DATA_FILES := $(shell find portal | sed 's/  /\\ /g')
ASSETS_DATA_FILES := $(shell find assets | sed 's/  /\\ /g')
GENERATED := pkg/assets/bindata.go pkg/portal/bindata.go

LDFLAGS += -X "github.com/alimy/mir-music/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "github.com/alimy/mir-music/version.GitHash=$(shell git rev-parse HEAD)"

default: build

build: fmt bindata
	go build -ldflags '$(LDFLAGS)' -tags '$(TAGS)' -o mir-music

serve: fmt bindata
	go run -ldflags '$(LDFLAGS)' -tags '$(TAGS)' github.com/alimy/mir-music serve --debug --addr :8013

fmt:
	go fmt ./...

bindata: $(GENERATED)

pkg/assets/bindata.go: $(ASSETS_DATA_FILES)
	rm -rf $@
	go-bindata -nomemcopy -pkg=assets \
	     -prefix=assets \
         -debug=$(if $(findstring debug,$(TAGS)),true,false) \
         -o=$@ assets/...
	gofmt -s -w pkg/assets

pkg/portal/bindata.go: $(PORTAL_DATA_FILES)
	rm -rf $@
	go-bindata -nomemcopy -pkg=portal -tags=portal \
	     -prefix=portal \
         -debug=$(if $(findstring debug,$(TAGS)),true,false) \
         -o=$@ portal/...
	gofmt -s -w pkg/portal

api:
	docker run -it --rm -p 8080:80 -v $(PWD)/api/openapi.yaml:/usr/share/nginx/html/openapi.yaml -e SPEC_URL=openapi.yaml redocly/redoc

clean:
	go clean -r ./...
	rm -f mir-music

distclean: clean
	rm -rf pkg/assets/bindata.go pkg/portal/bindata.go