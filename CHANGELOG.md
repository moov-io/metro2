## v0.5.0 (Released 2022-09-23)

IMPROVEMENTS

- added NewFileFromReader and Reader, replace string with byte array
- added reader that read by record from io, reduce the number of allocations needed using byte array instead of string
- cmd/metro2: fixup server to add timeouts, resolve gosec error
- fix: update ssn validation, reject more invalid SSNs

BUILD

- fix(deps): update golang.org/x/oauth2 digest to f213421
- fix(deps): update module github.com/moov-io/base to v0.35.0

## v0.4.1 (Released 2022-02-28)

BUG FIXES

- file: fix `TrailerRecord.TotalConsumerSegmentsJ1` calculation

## v0.4.0 (Released 2022-04-27)

IMPROVEMENTS

- file: support creating files with one record per line

BUG FIXES

- file: fix character encoding issue when converting json to metro2
- file: fix file size string format when converting metro2 back to json

## v0.3.3 (Released 2021-07-16)

IMPROVEMENTS

- Update schema to support leases

BUILD

- fix(deps): update golang.org/x/oauth2 commit hash to a41e5a7
- fix(deps): update module github.com/spf13/cobra to v1.2.1
- fix: Dockerfile-fuzz to reduce vulnerabilities

## v0.3.2 (Released 2021-06-22)

BUG FIXES

- client: expand telephone numbers to int64
- header: expand reporter telephone number to int64

IMPROVEMENTS

- client: rengerate, fix bugs and issues
- fix some errors, replace test files with valid files

BUILD

- build: upgrade github.com/gogo/protobuf to v1.3.2
- fix(deps): update golang.org/x/oauth2 commit hash to d040287

## v0.3.1 (Released 2021-02-26)

BUILD

- build: fix Docker build issue in v0.3.0 while deploying

## v0.3.0 (Released 2021-02-26)

IMPROVEMENTS

- docs: lots of improvements across the Go types, markdown, etc
- pkg: improved error messages
- site: Create GitHub Pages documentation site with Jekyll (#46)

BUILD

- chore(deps): update golang docker tag to v1.16

## v0.2.7 (Released 2020-09-13)

IMPROVEMENTS

- docs: include section for the docker image
- file: print filepaths with error

BUILD

- chore(deps): update golang docker tag to v1.15
- chore(deps): update module gorilla/mux to v1.8.0

## v0.2.6 (Released 2020-07-24)

BUG FIXES

- file: fix crash from negative offsets

BUILD

- chore(deps): update golang.org/x/oauth2 commit hash to bf48bf1

## v0.2.5 (Released 2020-07-10)

BUG FIXES

- cmd/metro2: bind to all interfaces
- server: check 'true' better

BUILD

- build: add maintainer label

## v0.2.4 (Released 2020-07-09)

BUG FIXES

- lib: fix common out-of-bounds error checks

## v0.2.3 (Released 2020-07-08)

IMPROVEMENTS

- file: fix panic from out of range lookup
- fuzz: improve detection to collect files into the corpus

## v0.2.2 (Released 2020-07-07)

IMPROVEMENTS

- Small tweaks to fuzzing container

## v0.2.1 (Released 2020-07-07)

ADDITIONS

- Created fuzzing container (`docker run moov/metro2fuzz:v0.2.1`)
- Added OpenAPI specification
- Added generated Go client (`github.com/moov-io/metro2/pkg/client`)

## v0.2.0 (Released 2020-07-03)

ADDITIONS

- cmd/server: add web server
- build: Publish Openshift images on each release

BUILD

- build: enable codeql scanning and use Github Actions for CI

## v0.1.1 (Released 2020-06-29)

This is the initial release of Metro2 for credit history parsing and creation. Included is a Go library and command-line utility.
