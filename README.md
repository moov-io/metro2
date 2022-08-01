[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://moov-io.github.io/metro2/">Project Documentation</a>
  ·
  <a href="https://moov-io.github.io/metro2/api/#overview">API Endpoints</a>
  ·
  <a href="https://moov.io/blog/education/metro2-api-guide/">API Guide</a>
  ·
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://godoc.org/github.com/moov-io/metro2?status.svg)](https://godoc.org/github.com/moov-io/metro2)
[![Build Status](https://github.com/moov-io/metro2/workflows/Go/badge.svg)](https://github.com/moov-io/metro2/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/metro2/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/metro2)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/metro2)](https://goreportcard.com/report/github.com/moov-io/metro2)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/metro2?label=project%20size)](https://github.com/moov-io/metro2)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/metro2/master/LICENSE)
[![Slack Channel](https://slack.moov.io/badge.svg?bg=e01563&fgColor=fffff)](https://slack.moov.io/)
[![Docker Pulls](https://img.shields.io/docker/pulls/moov/metro2)](https://hub.docker.com/r/moov/metro2)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/metro2)](https://github.com/moov-io/metro2)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

# moov-io/metro2

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease-of-use.

Metro2 implements a reader, writer, and validator for consumer credit history reports in an HTTP server and Go library. The HTTP server is available in a [Docker image](#docker) and the Go package `github.com/moov-io/metro2` is available.

## Table of contents

- [Project status](#project-status)
- [Usage](#usage)
  - As an API
    - [Docker](#docker)
    - [Google Cloud](#google-cloud-run)
    - [Data persistence](#data-persistence)
  - [As a command line tool](#command-line)
  - As a Package
      - [Go Library](#go-library)
      - [Installation](#package-installation)  
      - [Overview](#package-overview)
      - [Examples](#package-examples)
- [Learn about Metro 2](#learn-about-metro-2)
- [Getting help](#getting-help)
- [Supported and tested platforms](#supported-and-tested-platforms)
- [Contributing](#contributing)
- [Related projects](#related-projects)

## Project status

Moov Metro2 is actively used in multiple production environments. Please star the project if you are interested in its progress. If you have layers above Metro2 to simplify tasks, perform business operations, or found bugs we would appreciate an issue or pull request. Thanks!

## Usage

The Metro2 project implements an HTTP server and [Go library](https://pkg.go.dev/github.com/moov-io/metro2) for creating and modifying files in Metro 2 format, which is used for consumer credit history reporting by U.S. credit bureaus.

### Docker

We publish a [public Docker image `moov/metro2`](https://hub.docker.com/r/moov/metro2/) on Docker Hub with each tagged release of Metro2. No configuration is required to serve on `:8080`. We also have Docker images for [OpenShift](https://quay.io/repository/moov/metro2?tab=tags) published as `quay.io/moov/metro2`.

Pull & start the Docker image:
```
docker pull moov/metro2:latest
docker run -p 8080:8080 moov/metro2:latest
```

Upload a file and validate it:
```
curl -X POST --form "file=@./test/testdata/packed_file.json" http://localhost:8080/validator
```
```
valid file
```

Convert a file from JSON to raw Metro format:
```
curl -X POST --form "file=@./test/testdata/packed_file.json" --form "format=metro" http://localhost:8080/convert
```
```
0480HEADER                           555555555508202002051019990510199905101999YOUR BUSINESS NAME HERE                 LINE ONE OF YOUR ADDRESS LINE TWO OF YOUR ADDRESS LINE THERE OF YOUR ADDRESS                    1234567890
...
```

Convert a file from raw Metro format to JSON:
```
curl -X POST --form "file=@./test/testdata/packed_file.dat" --form "format=json" http://localhost:8080/convert
```
```
{"header": {"blockDescriptorWord": 370,"recordDescriptorWord": 366,"recordIdentifier": "HEADER","transUnionProgramIdentifier": "5555555555","activityDate": "2002-08-20T00:00:00Z", ...
```

### Google Cloud Run

To get started in a hosted environment you can deploy this project to the Google Cloud Platform.

From your [Google Cloud dashboard](https://console.cloud.google.com/home/dashboard) create a new project and call it:
```
moov-metro2-demo
```

Enable the [Container Registry](https://cloud.google.com/container-registry) API for your project and associate a [billing account](https://cloud.google.com/billing/docs/how-to/manage-billing-account) if needed. Then, open the Cloud Shell terminal and run the following Docker commands, substituting your unique project ID:

```
docker pull moov/metro2
docker tag moov/metro2 gcr.io/<PROJECT-ID>/metro2
docker push gcr.io/<PROJECT-ID>/metro2
```

Deploy the container to Cloud Run:
```
gcloud run deploy --image gcr.io/<PROJECT-ID>/metro2 --port 8080
```

Select your target platform to `1`, service name to `metro2`, and region to the one closest to you (enable Google API service if a prompt appears). Upon a successful build you will be given a URL where the API has been deployed:

```
https://YOUR-METRO2-APP-URL.a.run.app
```

Now you can complete a health check:
```
curl https://YOUR-METRO2-APP-URL.a.run.app/health
```
You should get this response:
```
{"health":true}
```

### Data persistence
By design, Metro2  **does not persist** (save) any data about the files or entry details created. The only storage occurs in memory of the process and upon restart Metro2 will have no files or data saved. Also, no in-memory encryption of the data is performed.

### Command line

Metro2 has a command line interface to manage Metro 2 files and launch a web service.

```
metro2 --help

Usage:
   [command]

Available Commands:
  convert     Convert metro file format
  help        Help about any command
  print       Print metro file
  validator   Validate metro file
  web         Launches web server

Flags:
  -h, --help           help for this command
      --input string   input file (default is $PWD/metro.json)

Use " [command] --help" for more information about a command.
```

Each interaction that the library supports is exposed in a command-line option:

 Command | Info
 ------- | -------
`convert` | The convert command allows users to convert a metro file to a specified file format (json, metro). The result will create a new file.
`print` | The print command allows users to print a metro file in a specified file format (json, metro).
`validator` | The validator command allows users to validate a metro file.
`web` | The web command will launch a web server with endpoints to manage metro files.


### Go library

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go v1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/metro2/releases/latest) as well. We highly recommend you use a tagged release for production.

### Package Installation
```
$ git@github.com:moov-io/metro2.git

# Pull down into the Go Module cache
$ go get -u github.com/moov-io/metro2

$ go doc github.com/moov-io/metro2
```

### Package Overview

Package file defines main interfaces to a metro file.

```
func NewFile(format string) (File, error)

func NewFileFromReader(r io.Reader) (File, error)

type File
	GetType() string
	SetType(string) error
	SetRecord(lib.Record) error
	AddDataRecord(lib.Record) error
	GetRecord(string) (lib.Record, error)
	GetDataRecords() []lib.Record
	GeneratorTrailer() (lib.Record, error)
	Parse(record string) error
	String(newline bool) string
	Validate() error
```

File has some records for header, body, trailer

Records have distinguished by file type

```
type Record
	Name() string
	Parse(string) (int, error)
	String() string
	Validate() error
	Length() int
	BlockSize() int
	AddApplicableSegment(Segment) error
	GetSegments(string) []Segment

type BaseSegment        // for character file format

type PackedBaseSegment  // for packed file format

type HeaderRecord       // for character file format

type PackedHeaderRecord // for packed file format

type TrailerRecord      // for character file format

type PackedTrailerRecord // for packed file format
```

Record has some segments to describe special form

Base Record can use correctly the features of appendable segments  

```
type Record
	Name() string
	Parse(string) (int, error)
	String() string
	Validate() error
	Length() int
	BlockSize() int
	AddApplicableSegment(Segment) error
	GetSegments(string) []Segment

type J1Segment

type J2Segment

type K1Segment

type K2Segment

type K3Segment

type K4Segment

type L1Segment

type N1Segment
```

### Package Examples

File object has type for metro packed file and metro character file, should specify file type when creating new file
```
f, err := file.NewFile(utils.PackedFileFormat) 
err = json.Unmarshal([]byte(buf), f)
```

File object will create from metro file buffer directly, CreateFile function will create new file and will set file type based on metro file body
```
r, _ = os.Open(filepath.Join("...))
f, err := file.NewFileFromReader(r) 
```

File object will manage records using member functions

Please use exact record object that has same type (header, base, trailer)

```
f, err := file.NewFile(utils.CharacterFileFormat)
newHeader := lib.HeaderRecord{}
err = f.SetRecord(&newSegment)
exsitedHeader, err := f.GetRecord(utils.HeaderRecordName)
newBase := lib.HeaderRecord{}
...
```

Similarly, Record can manage the appendable segments 

## Learn about Metro 2
- [Intro to Metro 2](https://www.cdiaonline.org/resources/furnishers-of-data-overview/metro2-information/)
- [Metro 2 File Structure](https://www.collect.org/cv11/Help/metro2format.html)

## Getting help

 channel | info
 ------- | -------
Twitter [@moov](https://twitter.com/moov)	| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) to get started!

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go v1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/metro2/releases/latest) as well. We highly recommend you use a tagged release for production.

### Releasing

To make a release of metro2 simply open a pull request with `CHANGELOG.md` and `version.go` updated with the next version number and details. You'll also need to push the tag (i.e. `git push origin v1.0.0`) to origin in order for CI to make the release.

### Testing

We maintain a comprehensive suite of unit tests and recommend table-driven testing when a particular function warrants several very similar test cases. To run all test files in the current directory, use `go test`. Current overall coverage can be found on [Codecov](https://app.codecov.io/gh/moov-io/metro2/).

### Fuzzing

We currently run fuzzing over ImageCashLetter in the form of a [`moov/metro2`](https://hub.docker.com/r/moov/metro2fuzz) Docker image. You can [read more](./test/fuzz-reader/README.md) or run the image and report crasher examples to [`security@moov.io`](mailto:security@moov.io). Thanks!

## Related projects
As part of Moov's initiative to offer open source fintech infrastructure, we have a large collection of active projects you may find useful:

- [Moov Watchman](https://github.com/moov-io/watchman) offers search functions over numerous trade sanction lists from the United States and European Union.

- [Moov Fed](https://github.com/moov-io/fed) implements utility services for searching the United States Federal Reserve System such as ABA routing numbers, financial institution name lookup, and FedACH and Fedwire routing information.

- [Moov Wire](https://github.com/moov-io/wire) implements an interface to write files for the Fedwire Funds Service, a real-time gross settlement funds transfer system operated by the United States Federal Reserve Banks.

- [Moov ACH](https://github.com/moov-io/ach) provides ACH file generation and parsing, supporting all Standard Entry Codes for the primary method of money movement throughout the United States.

- [Moov Image Cash Letter](https://github.com/moov-io/imagecashletter) implements Image Cash Letter (ICL) files used for Check21, X.9 or check truncation files for exchange and remote deposit in the U.S.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
