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
