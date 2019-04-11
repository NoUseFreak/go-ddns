
appname := go-ddns
version := $(shell git describe --tags --abbrev=0)
build_date := $(shell date +%Y-%m-%d\ %H:%M)

.PHONY: test
test:
	go run cmd/go-ddns/main.go local-fixtures/aws-route53.yml
	
all:
	go get github.com/mitchellh/gox
	mkdir -p build
	gox \
		-ldflags="-s -X 'main.version=${version}' -X 'main.buildDate=${build_date}'" \
		-output="build/{{.OS}}_{{.Arch}}/${appname}"
