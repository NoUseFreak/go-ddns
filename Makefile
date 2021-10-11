
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

publish:
	@if [ "${version}" == "" ]; then echo "Missing version"; exit 1; fi
	# git tag ${version}
	docker buildx use $(appname)-builder \
		|| (docker buildx create --name $(appname)-builder && docker buildx use $(appname)-builder)
	
	docker buildx inspect --bootstrap
	docker buildx build -t nousefreak/$(appname):$(version) \
	 	--platform linux/amd64,linux/arm64,linux/arm --push .
	docker buildx rm $(appname)-builder
	echo "Published version $(version)"
