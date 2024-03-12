PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
VERSION := $(shell grep -Eo '(v[0-9]+[\.][0-9]+[\.][0-9]+(-[a-zA-Z0-9]*)?)' version.go)

.PHONY: build docker release

build:
	go fmt ./...
	@mkdir -p ./bin/
	CGO_ENABLED=0 go build -o ./bin/metro2 github.com/bloomcredit/moov-metro2/cmd/metro2

.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	go test ./... -count 1 -race
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	COVER_THRESHOLD=85.0 time ./lint-project.sh
endif

check-openapi:
	docker run \
	-v ${PWD}/api/openapi.yaml:/projects/openapi.yaml \
	wework/speccy lint --verbose /projects/openapi.yaml

.PHONY: client
client:
ifeq ($(OS),Windows_NT)
	@echo "Please generate client on macOS or Linux, currently unsupported on windows."
else
# Versions from https://github.com/OpenAPITools/openapi-generator/releases
	@chmod +x ./openapi-generator
	@rm -rf ./client
	OPENAPI_GENERATOR_VERSION=6.6.0 ./openapi-generator generate --git-user-id=moov-io --git-repo-id=metro2 --package-name client -i ./api/openapi.yaml -g go -o ./pkg/client
	rm -f ./pkg/client/go.mod ./pkg/client/go.sum ./pkg/client/.travis.yml ./pkg/client/git_push.sh
	go fmt ./...
	go build github.com/bloomcredit/moov-metro2/pkg/client
	go test ./pkg/client/...
endif

dist: clean build
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/metro2.exe github.com/bloomcredit/moov-metro2/cmd/metro2
else
	CGO_ENABLED=0 GOOS=$(PLATFORM) go build -o bin/metro2-$(PLATFORM)-amd64 github.com/bloomcredit/moov-metro2/cmd/metro2
endif

docker: clean
# Docker image
	docker build --pull -t moov/metro2:$(VERSION) -f Dockerfile .
	docker tag moov/metro2:$(VERSION) moov/metro2:latest
# OpenShift Docker image
	docker build --pull -t quay.io/moov/metro2:$(VERSION) -f Dockerfile-openshift --build-arg VERSION=$(VERSION) .
	docker tag quay.io/moov/metro2:$(VERSION) quay.io/moov/metro2:latest

.PHONY: clean
clean:
ifeq ($(OS),Windows_NT)
	@echo "Skipping cleanup on Windows, currently unsupported."
else
	@rm -rf cover.out coverage.txt misspell* staticcheck*
	@rm -rf ./bin/ openapi-generator-cli-*.jar metro2.db ./storage/ lint-project.sh
endif

.PHONY: cover-test cover-web
cover-test:
	go test -coverprofile=cover.out ./...
cover-web:
	go tool cover -html=cover.out

release: docker AUTHORS
	go vet ./...
	go test -coverprofile=cover-$(VERSION).out ./...
	git tag -f $(VERSION)

# From https://github.com/genuinetools/img
.PHONY: AUTHORS
AUTHORS:
	@$(file >$@,# This file lists all individuals having contributed content to the repository.)
	@$(file >>$@,# For how it is generated, see `make AUTHORS`.)
	@echo "$(shell git log --format='\n%aN <%aE>' | LC_ALL=C.UTF-8 sort -uf)" >> $@

release-push:
	docker push moov/metro2:$(VERSION)
	docker push moov/metro2:latest

quay-push:
	docker push quay.io/moov/metro2:$(VERSION)
	docker push quay.io/moov/metro2:latest
