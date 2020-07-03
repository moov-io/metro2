PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
VERSION := $(shell grep -Eo '(v[0-9]+[\.][0-9]+[\.][0-9]+(-[a-zA-Z0-9]*)?)' version.go)

.PHONY: build docker release

build:
	go fmt ./...
	@mkdir -p ./bin/
	CGO_ENABLED=1 go build -o ./bin/metro2 github.com/moov-io/metro2/cmd/metro2

.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	@echo "Skipping checks on Windows, currently unsupported."
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	time ./lint-project.sh
endif

dist: clean build
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/metro2.exe github.com/moov-io/metro2/cmd/metro2
else
	CGO_ENABLED=0 GOOS=$(PLATFORM) go build -o bin/metro2-$(PLATFORM)-amd64 github.com/moov-io/metro2/cmd/metro2
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

release-push:
	docker push moov/metro2:$(VERSION)
	docker push moov/metro2:latest

quay-push:
	docker push quay.io/moov/metro2:$(VERSION)
	docker push quay.io/moov/metro2:latest
