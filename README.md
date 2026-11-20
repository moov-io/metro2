[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://moov-io.github.io/metro2/">Project Documentation</a>
  ·
  <a href="https://moov-io.github.io/metro2/api/#overview">API Endpoints</a>
  ·
  <a href="https://moov.io/blog/craft/metro2-api-guide/">API Guide</a>
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
[![Slack Channel](https://img.shields.io/badge/Slack-Moov-e01563?logo=slack&logoColor=white)](https://slack.moov.io/)
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
{"status":"valid file"}
```

Convert a file from JSON to raw Metro format:
```
curl -X POST --form "file=@./test/testdata/packed_file.json" --form "format=metro" http://localhost:8080/convert
```
The response is a packed Metro 2 file (`Content-Type: application/octet-stream`) that starts with binary block and record descriptor words (`01 72 00 00 01 6e 00 00`, then `HEADER`), not ASCII text. To get character-encoded records, add `--form "type=character"`.


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

Enable the [Artifact Registry](https://cloud.google.com/artifact-registry/docs/docker/store-docker-container-images) and Cloud Run APIs for your project and associate a [billing account](https://cloud.google.com/billing/docs/how-to/manage-billing-account) if needed. Then, open the Cloud Shell terminal and run the following commands, substituting your unique project ID and a region close to you (for example `us-central1`):

```
gcloud services enable artifactregistry.googleapis.com run.googleapis.com

gcloud artifacts repositories create moov   --repository-format=docker   --location=us-central1   --description="Docker images"

gcloud auth configure-docker us-central1-docker.pkg.dev

docker pull moov/metro2
docker tag moov/metro2 us-central1-docker.pkg.dev/<PROJECT-ID>/moov/metro2
docker push us-central1-docker.pkg.dev/<PROJECT-ID>/moov/metro2
```

Deploy the container to Cloud Run:
```
gcloud run deploy metro2 --image us-central1-docker.pkg.dev/<PROJECT-ID>/moov/metro2 --port 8080 --region us-central1
```

You can also deploy the public image directly with `gcloud run deploy metro2 --image moov/metro2:latest --port 8080 --region us-central1`. Upon a successful deploy you will be given a URL where the API has been deployed:

```
https://YOUR-METRO2-APP-URL.a.run.app
```

Now you can complete a health check:
```
curl https://YOUR-METRO2-APP-URL.a.run.app/health
```
You should get this response:
```
{"status":"alive"}
```

### Data persistence
By design, Metro2  **does not persist** (save) any data about the files or entry details created. The only storage occurs in memory of the process and upon restart Metro2 will have no files or data saved. Also, no in-memory encryption of the data is performed.

### Command line

Metro2 has a command line interface to manage Metro 2 files and launch a web service.

Download a release binary from the [latest GitHub release](https://github.com/moov-io/metro2/releases/latest). Assets are named `metro2-linux-amd64`, `metro2-darwin-amd64`, and `metro2.exe` (Windows):

```
curl -sSL -o metro2 "https://github.com/moov-io/metro2/releases/download/v0.10.0/metro2-linux-amd64"
chmod +x metro2
```

Or install with Go 1.26 or newer:

```
go install github.com/moov-io/metro2/cmd/metro2@latest
```

From a source checkout, `make build` writes `./bin/metro2`.

```
metro2 --help

Moov Metro 2 file reader, writer, and validator

Usage:
  metro2 [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  convert     Convert metro file format
  help        Help about any command
  print       Print metro file
  validator   Validate metro file
  version     Print the version of metro2
  web         Launches web server

Flags:
  -h, --help           help for metro2
      --input string   input file (default is $PWD/metro.json)

Use "metro2 [command] --help" for more information about a command.
```

The CLI exits `0` on success and `1` on error (missing input, validation failure, or a convert path whose parent directory does not exist). `convert` supports `--format json|metro`, `--type character|packed`, `-n`/`--newline`, and `-g`/`--generate`. See the [command line documentation](https://moov-io.github.io/metro2/usage-command-line/) for packed versus character format, sample output, and exit codes.

Each interaction that the library supports is exposed in a command-line option:

 Command | Info
 ------- | -------
`convert` | The convert command allows users to convert a metro file to a specified file format (json, metro). The result will create a new file.
`print` | The print command allows users to print a metro file in a specified file format (json, metro).
`validator` | The validator command allows users to validate a metro file.
`version` | Print the library version.
`web` | The web command will launch a web server with endpoints to manage metro files.


### Go library

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.26 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/metro2/releases/latest) as well. We highly recommend you use a tagged release for production.

### Package Installation
```
# Clone the repository
git clone git@github.com:moov-io/metro2.git

# Add a tagged package to your Go module (avoid go get -u; it can bump
# the Go version in go.mod and leave missing go.sum entries)
go get github.com/moov-io/metro2/pkg/file@v0.10.0

go doc github.com/moov-io/metro2/pkg/file
```

### Package Overview

The Go API lives in [`pkg/file`](https://pkg.go.dev/github.com/moov-io/metro2/pkg/file), [`pkg/lib`](https://pkg.go.dev/github.com/moov-io/metro2/pkg/lib), and [`pkg/utils`](https://pkg.go.dev/github.com/moov-io/metro2/pkg/utils). `file.File` is the main entry point:

```
func NewFile(format string) (File, error)
func NewFileFromReader(r io.Reader) (File, error)
func CreateFile(buf []byte) (File, error)

type File
	GetType() string
	SetType(string) error
	SetRecord(lib.Record) error
	AddDataRecord(lib.Record) error
	GetRecord(string) (lib.Record, error)
	GetDataRecords() []lib.Record
	GeneratorTrailer() (lib.Record, error)
	Parse(record []byte, isVariableLength bool) error
	Bytes() []byte
	String(newline bool) string
	ConcurrentString(newline bool, goroutines int) string
	Validate() error
```

Use `utils.CharacterFileFormat` or `utils.PackedFileFormat` with `NewFile`. Records (`lib.HeaderRecord`, `lib.BaseSegment`, `lib.TrailerRecord`, and the packed variants) implement `Parse([]byte, bool)`, `String()`, `Bytes()`, and `Validate()`. Optional J1, J2, K1–K4, L1, and N1 segments attach to a base record with `AddApplicableSegment`. See the [Go library docs](https://moov-io.github.io/metro2/usage-go/) for details and a from-scratch example.

### Package Examples

Read a Metro 2 file, validate it, and write it back out:

```go
r, err := os.Open("test/testdata/unpacked_fixed_file.dat")
if err != nil {
	log.Fatal(err)
}
defer r.Close()

f, err := file.NewFileFromReader(r)
if err != nil {
	log.Fatal(err)
}
if err := f.Validate(); err != nil {
	log.Fatal(err)
}

fmt.Print(f.String(true))
```

To build a file from scratch, create a `file.File`, `SetRecord` a header, `AddDataRecord` one or more base segments (with optional J1/K segments), then `GeneratorTrailer` / `SetRecord` and `Validate`. A complete, tested example is in the [Go library docs](https://moov-io.github.io/metro2/usage-go/) and as `ExampleNewFileFromReader` / `Example_createFile` in `pkg/file`.

## Learn about Metro 2
- [Intro to Metro 2](https://www.cdiaonline.org/resources/furnishers-of-data-overview/metro2-information/)
- [Metro 2 File Structure](https://www.collect.org/cv11/Help/metro2format.html)

## Getting help

 channel | info
 ------- | -------
Twitter [@moov](https://twitter.com/moov)	| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/metro2/issues) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) to get started!

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.26 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/metro2/releases/latest) as well. We highly recommend you use a tagged release for production.

### Releasing

To make a release of metro2 update `version.go` with the next version number and push a tag (i.e. `git push origin v1.0.0`) so CI can create the GitHub Release. Release notes live on [GitHub Releases](https://github.com/moov-io/metro2/releases); `CHANGELOG.md` is frozen for historical entries.

### Testing

We maintain a comprehensive suite of unit tests and recommend table-driven testing when a particular function warrants several very similar test cases. To run all test files in the current directory, use `go test`. Current overall coverage can be found on [Codecov](https://app.codecov.io/gh/moov-io/metro2/).

### Fuzzing

We currently run fuzzing over Metro 2 in the form of a [Github Action](https://github.com/moov-io/metro2/actions/workflows/fuzz.yml). Please report crash examples to [`oss@moov.io`](mailto:oss@moov.io). Thanks!

## Related projects

As part of Moov's initiative to offer open source fintech infrastructure, we have a large collection of active projects you may find useful:

- [Moov Watchman](https://github.com/moov-io/watchman) offers search functions over numerous trade sanction lists from the United States and European Union.

- [Moov Fed](https://github.com/moov-io/fed) implements utility services for searching the United States Federal Reserve System such as ABA routing numbers, financial institution name lookup, and FedACH and Fedwire routing information.

- [Moov Wire](https://github.com/moov-io/wire) implements an interface to write files for the Fedwire Funds Service, a real-time gross settlement funds transfer system operated by the United States Federal Reserve Banks.

- [Moov ACH](https://github.com/moov-io/ach) provides ACH file generation and parsing, supporting all Standard Entry Codes for the primary method of money movement throughout the United States.

- [Moov Image Cash Letter](https://github.com/moov-io/imagecashletter) implements Image Cash Letter (ICL) files used for Check21, X.9 or check truncation files for exchange and remote deposit in the U.S.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
