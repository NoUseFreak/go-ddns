
appname := go-ddns
version ?= $(shell git describe --tags --abbrev=0)
build_date := $(shell date +%Y-%m-%d\ %H:%M)


default: clean all package

clean:
	rm -rf build/

.PHONY: test
test:
	go run cmd/go-ddns/main.go local-fixtures/aws-route53.yml
	
all:
	go get github.com/mitchellh/gox
	mkdir -p build
	gox \
		-ldflags="-s -X 'main.version=${version}' -X 'main.buildDate=${build_date}'" \
		-output="build/{{.OS}}_{{.Arch}}/${appname}" \
		./cmd/${appname}

package:
	$(shell rm -rf build/archive)
	$(shell rm -rf build/archive)
	$(eval UNIX_FILES := $(shell ls build | grep -v ${appname} | grep -v windows))
	$(eval WINDOWS_FILES := $(shell ls build | grep -v ${appname} | grep windows))
	@mkdir -p build/archive
	@for f in $(UNIX_FILES); do \
		echo Packaging $$f && \
		(cd $(shell pwd)/build/$$f && tar -czf ../archive/$$f.tar.gz ${appname}*); \
	done
	@for f in $(WINDOWS_FILES); do \
		echo Packaging $$f && \
		(cd $(shell pwd)/build/$$f && zip ../archive/$$f.zip ${appname}*); \
	done
	ls -lah build/archive/


