# metro2

[![GoDoc](https://godoc.org/github.com/moov-io/metro2?status.svg)](https://godoc.org/github.com/moov-io/metro2)
[![Build Status](https://github.com/moov-io/metro2/workflows/Go/badge.svg)](https://github.com/moov-io/metro2/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/metro2/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/metro2)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/metro2)](https://goreportcard.com/report/github.com/moov-io/metro2)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/metro2/master/LICENSE)

Metro2 implements a reader, writer, and validator for consumer credit history reports in an HTTP server and Go library. The HTTP server is available in a [Docker image](#docker) and the Go package `github.com/moov-io/metro2` is available.

Docs: [API Endpoints](https://moov-io.github.io/metro2/api/)

## Getting Started

### Docker

We publish a [public Docker image `moov/metro2`](https://hub.docker.com/r/moov/metro2/tags) on Docker Hub with emetro2 tagged release of Wire. No configuration is required to serve on `:8080`. We also have docker images for [OpenShift](https://quay.io/repository/moov/metro2?tab=tags).

Start the Docker image:
```
docker run -p 8080:8080 moov/metro2:latest
```

Upload a file and validate it
```
curl -XPOST --form "file=@./test/testdata/packed_file.json" http://localhost:8080/validator
```
```
valid file
```

Convert a file between formats
```
curl -XPOST --form "file=@./test/testdata/packed_file.json" http://localhost:8080/convert -v
```
```
{"header":{"recordDescriptorWord":480,"recordIdentifier":"HEADER","transUnionProgramIdentifier":"5555555555","activityDate":"2002-08-20T00:00:00Z", ...
```

### Go Library

There is a Go library which can read and write credit reporting specifications. We write unit tests and fuzz the code to help ensure our code is production ready for everyone. Metro2 uses [Go Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies and suggests Go 1.14 or greater.

To clone our code and verify our tests on your system run:

```
$ git clone git@github.com:moov-io/metro2.git
$ cd metro2

$ go test ./...
ok   	github.com/moov-io/metro2	0.710s	coverage: 98.1% of statements
```

## Commands

Metro2 has command line interface to manage metro2 files and to lunch web service.

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
`convert` | The convert command allows users to convert from a metro file to another format file. Result will create a metro file.
`print` | The print command allows users to print a metro file with special file format (json, metro).
`validator` | The validator command allows users to validate a metro file.
`web` | The web command will launch a web server with endpoints to manage metro files.

### file convert

```
metro2 convert --help

Usage:
   convert [output] [flags]

Flags:
      --format string   format of metro file(required) (default "json")
  -g, --generate        generate trailer record
  -h, --help            help for convert

Global Flags:
      --input string   input file (default is $PWD/metro.json)
```

The output parameter is the full path name to convert new metro2 file.
The format parameter is supported 2 types, "json" and  "metro".
The generate parameter will replace new generated trailer record in the file.
The input parameter is source metro2 file, supported raw type file and json type file.

example:
```
metro2 convert output/output.json --input testdata/packed_file.json --format json
```

### file print

```
metro2 print --help

Usage:
   print [flags]

Flags:
      --format string   print format (default "json")
  -h, --help            help for print

Global Flags:
      --input string   input file (default is $PWD/metro.json)
```

The format parameter is supported 2 types, "json" and  "metro".
The input parameter is source metro2 file, supported raw type file and json type file.

example:
```
metro2 print --input testdata/packed_file.dat --format json
{
  "header": {
    "blockDescriptorWord": 370,
    "recordDescriptorWord": 366,
    "recordIdentifier": "HEADER",
    "transUnionProgramIdentifier": "5555555555",
    "activityDate": "2002-08-20T00:00:00Z",
    "dateCreated": "1999-05-10T00:00:00Z",
    "programDate": "1999-05-10T00:00:00Z",
    "programRevisionDate": "1999-05-10T00:00:00Z",
    "reporterName": "YOUR BUSINESS NAME HERE",
    "reporterAddress": "LINE ONE OF YOUR ADDRESS LINE TWO OF YOUR ADDRESS LINE THERE OF YOUR ADDRESS",
    "reporterTelephoneNumber": 1234567890
  },
  ...
}
```

### file validate

```
metro2 validator --help

Usage:
   validator [flags]

Flags:
  -h, --help   help for validator

Global Flags:
      --input string   input file (default is $PWD/metro.json)
```

The input parameter is source metro2 file, supported raw type file and json type file.

example:
```
metro2 validator --input testdata/packed_file.dat
Error: is an invalid value of TotalConsumerSegmentsJ1

```

### web server

```
metro2 web --help

Usage:
   web [flags]

Flags:
  -h, --help          help for web
      --port string   port of the web server (default "8080")
  -t, --test          test server

Global Flags:
      --input string   input file (default is $PWD/metro.json)
```

The port parameter is port number of web service.

example:
```
metro2 web
```

Web server have some endpoints to manage metro2 file

Method | Endpoint | Content-Type | Info
 ------- | ------- | ------- | -------
 `POST` | `/convert` | multipart/form-data | convert metro file. will download new file.
 `GET` | `/health` | text/plain | check web server.
 `POST` | `/print` | multipart/form-data | print metro file.
 `POST` | `/validator` | multipart/form-data | validate metro file.

web page example to use metro2 web server:

```
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Single file upload</title>
</head>
<body>
<h1>Upload single file with fields</h1>

<form action="http://localhost:8080/convert" method="post" enctype="multipart/form-data">
    Format: <input type="text" name="format"><br>
    Files: <input type="file" name="file"><br><br>
    <input type="submit" value="Submit">
</form>
</body>
</html>
```

## Docker

You can run the [moov/metro2 Docker image](https://hub.docker.com/r/moov/metro2) which defaults to starting the HTTP server.

```
docker run -p 8080:8080 moov/metro2:latest
```

## Getting Help

 channel | info
 ------- | -------
  Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/metro2/issues) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#metro2`) to have an interactive discussion about the development of the project.

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) to get started! [Checkout our issues](https://github.com/moov-io/metro2/issues) for something to help out with.

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go 1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/metro2/releases/latest) as well. We highly recommend you use a tagged release for production.

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
