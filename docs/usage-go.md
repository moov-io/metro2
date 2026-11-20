---
layout: page
title: Go library
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Go library

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.26 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/metro2/releases/latest) as well. We highly recommend you use a tagged release for production.

Full API documentation is on [pkg.go.dev](https://pkg.go.dev/github.com/moov-io/metro2):

- [`pkg/file`](https://pkg.go.dev/github.com/moov-io/metro2/pkg/file) — `File`, `NewFile`, `NewFileFromReader`, and `CreateFile`
- [`pkg/lib`](https://pkg.go.dev/github.com/moov-io/metro2/pkg/lib) — header, base, trailer, and optional segments (J1, J2, K1–K4, L1, N1)
- [`pkg/utils`](https://pkg.go.dev/github.com/moov-io/metro2/pkg/utils) — `CharacterFileFormat`, `PackedFileFormat`, and shared constants

```
# Clone the repository
git clone git@github.com:moov-io/metro2.git

# Add a tagged package to your Go module (do not use go get -u; it can bump
# the Go version in go.mod and leave missing go.sum entries)
go get github.com/moov-io/metro2/pkg/file@v0.10.0

go doc github.com/moov-io/metro2/pkg/file
```

## Package overview

`file.File` is the main entry point:

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

Pass `utils.CharacterFileFormat` (`"character"`) or `utils.PackedFileFormat` (`"packed"`) to `NewFile`. `NewFileFromReader` and `CreateFile` detect JSON versus raw Metro 2, and packed versus character format, from the contents.

Records and segments implement:

```
type Record
    Name() string
    Parse([]byte, bool) (int, error)
    String() string
    Bytes() []byte
    Validate() error
    Length() int
    BlockSize() int
    AddApplicableSegment(Segment) error
    GetSegments(string) []Segment

type Segment
    Name() string
    Parse([]byte, bool) (int, error)
    String() string
    Bytes() []byte
    Validate() error
    Length() int
```

Use `lib.HeaderRecord`, `lib.BaseSegment`, and `lib.TrailerRecord` for character files, and `lib.PackedHeaderRecord`, `lib.PackedBaseSegment`, and `lib.PackedTrailerRecord` for packed files. Optional appendable segments (`lib.J1Segment`, `lib.J2Segment`, `lib.K1Segment`, `lib.K2Segment`, `lib.K3Segment`, `lib.K4Segment`, `lib.L1Segment`, `lib.N1Segment`) are always character-encoded.

## Read, validate, and write a file

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/moov-io/metro2/pkg/file"
)

func main() {
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
}
```

This example is also the `ExampleNewFileFromReader` test in `pkg/file`.

## Build a file from scratch

A Metro 2 file has one header, one or more base (data) records, and a trailer. Base records can include optional J1/J2 (associated consumers) and K/L/N segments. After the data records are in place, call `GeneratorTrailer` (or `metro2 convert -g`) to fill in the trailer counts, then `Validate`.

Character format is a 426-byte unpacked record (fixed or variable blocked). Packed format is a 366-byte variable-blocked record with binary descriptor words. See [File structure](/file-structure/) and each bureau's Metro 2 specification for which format to send.

The following program builds a character-format file, attaches a J1 segment in JSON and a K1 segment in code, generates the trailer, and validates the result. It is tested as `Example_createFile` in `pkg/file`.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/moov-io/metro2/pkg/file"
	"github.com/moov-io/metro2/pkg/lib"
	"github.com/moov-io/metro2/pkg/utils"
)

func main() {
	f, err := file.NewFile(utils.CharacterFileFormat)
	if err != nil {
		log.Fatal(err)
	}

	header := &lib.HeaderRecord{}
	if err := json.Unmarshal([]byte(`{
		"recordDescriptorWord": 426,
		"recordIdentifier": "HEADER",
		"transUnionProgramIdentifier": "5555555555",
		"activityDate": "2002-08-20T00:00:00Z",
		"dateCreated": "1999-05-10T00:00:00Z",
		"programDate": "1999-05-10T00:00:00Z",
		"programRevisionDate": "1999-05-10T00:00:00Z",
		"reporterName": "YOUR BUSINESS NAME HERE",
		"reporterAddress": "LINE ONE OF YOUR ADDRESS LINE TWO OF YOUR ADDRESS LINE THERE OF YOUR ADDRESS",
		"reporterTelephoneNumber": 1234567890
	}`), header); err != nil {
		log.Fatal(err)
	}
	if err := f.SetRecord(header); err != nil {
		log.Fatal(err)
	}

	// Base-segment JSON is wrapped as {"base": {...}, "j1": [...], "k1": {...}}.
	base := lib.NewBaseSegment().(*lib.BaseSegment)
	if err := json.Unmarshal([]byte(`{
		"base": {
			"recordDescriptorWord": 426,
			"processingIndicator": 1,
			"timeStamp": "2020-01-01T01:11:22Z",
			"identificationNumber": "6666666",
			"consumerAccountNumber": "553723456",
			"portfolioType": "M",
			"accountType": "48",
			"dateOpened": "2019-11-22T00:00:00Z",
			"highestCredit": 2438,
			"termsDuration": "LOC",
			"actualPaymentAmount": 100,
			"accountStatus": "11",
			"paymentHistoryProfile": "BBBBBBBBBBBBBBBBBBBBBBBB",
			"currentBalance": 2233,
			"dateAccountInformation": "2002-09-20T00:00:00Z",
			"surname": "SMITH-JONES",
			"firstName": "JUNIOR",
			"socialSecurityNumber": 159328759,
			"dateBirth": "1972-03-18T00:00:00Z",
			"telephoneNumber": 5175555555,
			"ecoaCode": "1",
			"countryCode": "US",
			"firstLineAddress": "123 MAIN STREET",
			"city": "LANSING",
			"state": "MI",
			"zipCode": "72654"
		},
		"j1": [
			{
				"segmentIdentifier": "J1",
				"surname": "BEAUCHAMP",
				"firstName": "KEVIN",
				"socialSecurityNumber": 445112877,
				"dateBirth": "2020-01-02T00:00:00Z",
				"ecoaCode": "2"
			}
		]
	}`), base); err != nil {
		log.Fatal(err)
	}

	k1 := lib.NewK1Segment().(*lib.K1Segment)
	k1.OriginalCreditorName = "ORIGINAL CREDITOR"
	k1.CreditorClassification = 1
	if err := base.AddApplicableSegment(k1); err != nil {
		log.Fatal(err)
	}
	// Include optional segments in the record descriptor word.
	base.RecordDescriptorWord = lib.UnpackedRecordLength + lib.J1SegmentLength + lib.K1SegmentLength

	if err := f.AddDataRecord(base); err != nil {
		log.Fatal(err)
	}

	trailer, err := f.GeneratorTrailer()
	if err != nil {
		log.Fatal(err)
	}
	if err := f.SetRecord(trailer); err != nil {
		log.Fatal(err)
	}
	if err := f.Validate(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(f.String(true))
}
```

`SetType` can recast a packed file to character format (or the reverse). It converts the record struct types but does not rewrite descriptor words, so a packed file that is recast to character still starts with `03700366HEADER` when the original block and record descriptor words were 370 and 366.
