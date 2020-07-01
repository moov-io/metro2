# metro2

[![GoDoc](https://godoc.org/github.com/moov-io/metro2?status.svg)](https://godoc.org/github.com/moov-io/metro2)
[![Build Status](https://travis-ci.com/moov-io/metro2.svg?branch=master)](https://travis-ci.com/moov-io/metro2)
[![Coverage Status](https://codecov.io/gh/moov-io/metro2/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/metro2)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/metro2)](https://goreportcard.com/report/github.com/moov-io/metro2)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/metro2/master/LICENSE)

Metro2 is an open-source consumer credit history report for credit report file creation and validation.

## Getting Started

Metro2 is primarily a Go library which can read and write credit reporting specifications. We write unit tests and fuzz the code to help ensure our code is production ready for everyone. Metro2 uses [Go Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies and suggests Go 1.14 or greater.

To clone our code and verify our tests on your system run:

```
$ git clone git@github.com:moov-io/metro2.git
$ cd metro2

$ go test ./...
ok   	github.com/moov-io/metro2	0.710s	coverage: 98.1% of statements
```

## Commands

Each interaction that the library supports is exposed in a command-line option:

 Command | Info
 ------- | -------
`convert` | The convert command allows users to convert from a metro file to another format file. Result will create a metro file.
`print` | The print command allows users to print a metro file with special file format (json, metro).
`validator` | The validator command allows users to validate a metro file.
`web` | The web command will launch a web server with endpoints to management metro files.

### Web server endpoints

Method | Endpoint | Content-Type | Info
 ------- | ------- | ------- | -------
 `POST` | `/convert` | multipart/form-data | convert metro file. will download new file.
 `GET` | `/health` | text/plain | check web server.
 `POST` | `/print` | multipart/form-data | print metro file.
 `POST` | `/validator` | multipart/form-data | validate metro file.

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
