// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/metro2/pkg/lib"
	"gopkg.in/check.v1"
)

func TestFile__Crashers(t *testing.T) {
	paths := readCrasherInputFilepaths(t)
	for i := range paths {
		bs, err := ioutil.ReadFile(paths[i])
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("parsing %s", paths[i])

		if _, err := CreateFile(bs); err == nil {
			t.Error("expected error")
		}

		if testing.Verbose() {
			t.Logf("read %s without crashing", paths[i])
		}
	}
}

func readCrasherInputFilepaths(t *testing.T) []string {
	t.Helper()

	basePath := filepath.Join("..", "..", "test", "testdata", "crashers")
	fds, err := ioutil.ReadDir(basePath)
	if err != nil {
		t.Fatal(err)
	}

	var out []string
	for i := range fds {
		if strings.HasSuffix(fds[i].Name(), ".output") {
			continue
		}
		out = append(out, filepath.Join(basePath, fds[i].Name()))
	}
	return out
}

func (t *FileTest) TestJsonWithUnpackedVariableBlocked(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal([]byte(t.unpackedVariableBlockedJson), f)
	c.Assert(err, check.IsNil)
	c.Assert(f.String(), check.Equals, t.unpackedVariableBlockedFile)
	buf, err := json.Marshal(f)
	c.Assert(err, check.IsNil)
	var out bytes.Buffer
	err = json.Indent(&out, buf, "", "  ")
	c.Assert(err, check.IsNil)
	jsonStr := out.String()
	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	c.Assert(jsonStr, check.Equals, t.unpackedVariableBlockedJson)
}

func (t *FileTest) TestParseWithUnpackedVariableBlockedFileParse(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.unpackedVariableBlockedFile)
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestJsonWithUnpackedFixedLength(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal([]byte(t.unpackedFixedLengthJson), f)
	c.Assert(err, check.IsNil)
	buf, err := json.Marshal(f)
	c.Assert(err, check.IsNil)
	var out bytes.Buffer
	err = json.Indent(&out, buf, "", "  ")
	c.Assert(err, check.IsNil)
	jsonStr := out.String()
	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	c.Assert(jsonStr, check.Equals, t.unpackedFixedLengthJson)
}

func (t *FileTest) TestParseWithUnpackedFixedLength(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.unpackedFixedLengthFile)
	c.Assert(err, check.IsNil)
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	trailer := f.GetDataRecords()[0]
	a := trailer.(*lib.BaseSegment)
	a.AccountStatus = lib.AccountStatusDF
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatusDA
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus05
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus11
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus13
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus61
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus63
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus64
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus65
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestParseWithUnpackedFixedLength2(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.unpackedFixedLengthFile)
	c.Assert(err, check.IsNil)
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	trailer := f.GetDataRecords()[0]
	a := trailer.(*lib.BaseSegment)
	a.AccountStatus = lib.AccountStatus71
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus78
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus80
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus82
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus83
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus84
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus88
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus89
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus93
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus94
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus95
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus96
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	a.AccountStatus = lib.AccountStatus97
	_, err = f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestJsonWithPackedBlocked(c *check.C) {
	f, err := NewFile(PackedFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal([]byte(t.packedJson), f)
	c.Assert(err, check.IsNil)
	buf, err := json.Marshal(f)
	c.Assert(err, check.IsNil)
	var out bytes.Buffer
	err = json.Indent(&out, buf, "", "  ")
	c.Assert(err, check.IsNil)
	jsonStr := out.String()
	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	c.Assert(jsonStr, check.Equals, t.packedJson)
}

func (t *FileTest) TestParseWithPackedFileParse(c *check.C) {
	f, err := NewFile(PackedFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.packedFile)
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestFileSetBlock(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	jsonStr := `{
		"blockDescriptorWord": 430,
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
	  }`
	newSegment := lib.HeaderRecord{}
	err = json.Unmarshal([]byte(jsonStr), &newSegment)
	c.Assert(err, check.IsNil)
	orgHeader, err := f.GetRecord(HeaderRecordName)
	c.Assert(err, check.IsNil)
	origin := orgHeader.BlockSize()
	err = f.SetRecord(&newSegment)
	c.Assert(err, check.IsNil)
	newHeader, err := f.GetRecord(HeaderRecordName)
	c.Assert(err, check.IsNil)
	c.Assert(origin, check.Not(check.Equals), newHeader.BlockSize())
}

func (t *FileTest) TestFileDataRecord(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	segment := lib.NewBaseSegment()
	err = json.Unmarshal([]byte(t.baseSegmentJson), segment)
	c.Assert(err, check.IsNil)
	err = f.AddDataRecord(segment)
	c.Assert(err, check.IsNil)
	list := f.GetDataRecords()
	c.Assert(len(list), check.Equals, 1)
}

func (t *FileTest) TestGeneratorTrailer(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal([]byte(t.unpackedFixedLengthJson), f)
	c.Assert(err, check.IsNil)
	trailer, err := f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	err = f.SetRecord(trailer)
	c.Assert(err, check.IsNil)
	err = f.Validate()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestGeneratorPackedTrailer(c *check.C) {
	f, err := NewFile(PackedFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal([]byte(t.packedJson), f)
	c.Assert(err, check.IsNil)
	trailer, err := f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	err = f.SetRecord(trailer)
	c.Assert(err, check.IsNil)
	err = f.Validate()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestFileValidate(c *check.C) {
	f, err := NewFile(PackedFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal([]byte(t.packedJson), f)
	c.Assert(err, check.IsNil)
	err = f.Validate()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestGetRecord(c *check.C) {
	f, err := NewFile(PackedFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal([]byte(t.packedJson), f)
	c.Assert(err, check.IsNil)
	_, err = f.GetRecord(lib.TrailerRecordName)
	c.Assert(err, check.IsNil)
	_, err = f.GetRecord(lib.BaseSegmentName)
	c.Assert(err, check.NotNil)
}

func (t *FileTest) TestCreateFile(c *check.C) {
	_, err := CreateFile([]byte(t.packedJson))
	c.Assert(err, check.IsNil)
	f, err := CreateFile([]byte(t.packedFile))
	c.Assert(err, check.IsNil)
	c.Assert(f.String(), check.Equals, t.packedFile)
}

func (t *FileTest) TestCreateFileFailed(c *check.C) {
	_, err := CreateFile([]byte(t.packedFile[8:]))
	c.Assert(err, check.NotNil)
	data := `{
  "header": {
    "recordDescriptorWord": 480,
    "recordIdentifier": "error",
  }
}`
	_, err = CreateFile([]byte(data))
	c.Assert(err, check.NotNil)
}
