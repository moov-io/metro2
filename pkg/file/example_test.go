// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/moov-io/metro2/pkg/file"
	"github.com/moov-io/metro2/pkg/lib"
	"github.com/moov-io/metro2/pkg/utils"
)

func ExampleNewFileFromReader() {
	r, err := os.Open(filepath.Join("..", "..", "test", "testdata", "unpacked_fixed_file.dat"))
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

	fmt.Println(f.String(false)[:10])
	// Output: 0470HEADER
}

func Example_createFile() {
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

	fmt.Println(f.String(true)[:10])
	// Output: 0426HEADER
}
