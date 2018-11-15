.PHONY: build test run release
.SILENT: test run docker-run

## vars
DEPLOY_OS = $(shell echo `uname` | tr '[:upper:]' '[:lower:]')

programName=std2tch
arch=amd64

PLATFORMS := windows linux darwin
os = $(word 1, $@)

version ?= latest

${PLATFORMS}:
	GOOS=${os} GOARCH=${arch} go build -o ${programName}.${os}.${version}

test:
	go test -v -cover ./...

run:
	./${programName}.${DEPLOY_OS}.${VERSION}

release: windows linux darwin
	
docker-build:
	docker-compose up builder

docker-run:
	if [ ! -f std2tch.linux.latest ]; then \
		$(MAKE) docker-build ; \
	fi

	docker-compose up app