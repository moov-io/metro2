// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/check.v1"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/metro2/pkg/lib"
	"github.com/moov-io/metro2/pkg/utils"
)

func TestFile__Crashers(t *testing.T) {
	paths := readCrasherInputFilePaths(t)
	for i := range paths {

		f, err := os.Open(paths[i])
		if err != nil {
			t.Fatal(err)
		}

		if testing.Verbose() {
			t.Logf("parsing %s", paths[i])
		}

		if _, err := NewFileFromReader(f); err == nil {
			t.Errorf("expected error with %s", paths[i])
		} else {
			t.Logf("error with %s\n  %#v", paths[i], err)
		}

		if testing.Verbose() {
			t.Logf("read %s without crashing", paths[i])
		}

		f.Close()
	}
}

func readCrasherInputFilePaths(t *testing.T) []string {
	t.Helper()

	basePath := filepath.Join("..", "..", "test", "testdata", "crashers")
	fds, err := os.ReadDir(basePath)
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
	f, err := NewFile(utils.CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal(t.unpackedVariableBlockedJson, f)
	c.Assert(err, check.IsNil)

	raw, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", "unpacked_variable_file.dat"))
	c.Assert(err, check.IsNil)

	rawStr := strings.ReplaceAll(string(raw), "\r\n", "\n")
	c.Assert(strings.Compare(f.String(true), rawStr), check.Equals, 0)

	buf, err := json.Marshal(f)
	c.Assert(err, check.IsNil)
	var out bytes.Buffer
	err = json.Indent(&out, buf, "", "  ")
	c.Assert(err, check.IsNil)
	jsonStr := out.String()
	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	c.Assert(jsonStr, check.Equals, string(t.unpackedVariableBlockedJson))
}

func (t *FileTest) TestParseWithUnpackedVariableBlockedFileParse(c *check.C) {
	f, err := NewFile(utils.CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.unpackedVariableBlockedRaw)
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestJsonWithUnpackedFixedLength(c *check.C) {
	f, err := NewFile(utils.CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal(t.unpackedFixedLengthJson, f)
	c.Assert(err, check.IsNil)
	buf, err := json.Marshal(f)
	c.Assert(err, check.IsNil)
	var out bytes.Buffer
	err = json.Indent(&out, buf, "", "  ")
	c.Assert(err, check.IsNil)
	jsonStr := out.String()
	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	c.Assert(jsonStr, check.Equals, string(t.unpackedFixedLengthJson))
}

func (t *FileTest) TestParseWithUnpackedFixedLength(c *check.C) {
	f, err := NewFile(utils.CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.unpackedFixedLengthRaw)
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
	f, err := NewFile(utils.CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.unpackedFixedLengthRaw)
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
	f, err := NewFile(utils.PackedFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal(t.packedJson, f)
	c.Assert(err, check.IsNil)
	buf, err := json.Marshal(f)
	c.Assert(err, check.IsNil)
	var out bytes.Buffer
	err = json.Indent(&out, buf, "", "  ")
	c.Assert(err, check.IsNil)
	jsonStr := out.String()
	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	c.Assert(jsonStr, check.Equals, string(t.packedJson))
}

func (t *FileTest) TestParseWithPackedFileParse(c *check.C) {
	f, err := NewFile(utils.PackedFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.packedRaw)
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestFileSetBlock(c *check.C) {
	f, err := NewFile(utils.CharacterFileFormat)
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
	orgHeader, err := f.GetRecord(utils.HeaderRecordName)
	c.Assert(err, check.IsNil)
	origin := orgHeader.BlockSize()
	err = f.SetRecord(&newSegment)
	c.Assert(err, check.IsNil)
	newHeader, err := f.GetRecord(utils.HeaderRecordName)
	c.Assert(err, check.IsNil)
	c.Assert(origin, check.Not(check.Equals), newHeader.BlockSize())
}

func (t *FileTest) TestFileDataRecord(c *check.C) {
	f, err := NewFile(utils.CharacterFileFormat)
	c.Assert(err, check.IsNil)
	segment := lib.NewBaseSegment()
	err = json.Unmarshal(t.baseSegmentJson, segment)
	c.Assert(err, check.IsNil)
	err = f.AddDataRecord(segment)
	c.Assert(err, check.IsNil)
	list := f.GetDataRecords()
	c.Assert(len(list), check.Equals, 1)
}

func (t *FileTest) TestGeneratorTrailer(c *check.C) {
	f, err := NewFile(utils.CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal(t.unpackedFixedLengthJson, f)
	c.Assert(err, check.IsNil)
	trailer, err := f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	err = f.SetRecord(trailer)
	c.Assert(err, check.IsNil)
	err = f.Validate()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestGeneratorPackedTrailer(c *check.C) {
	f, err := NewFile(utils.PackedFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal(t.packedJson, f)
	c.Assert(err, check.IsNil)
	trailer, err := f.GeneratorTrailer()
	c.Assert(err, check.IsNil)
	err = f.SetRecord(trailer)
	c.Assert(err, check.IsNil)
	err = f.Validate()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestFileValidate(c *check.C) {
	f, err := NewFile(utils.PackedFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal(t.packedJson, f)
	c.Assert(err, check.IsNil)
	err = f.Validate()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestGetRecord(c *check.C) {
	f, err := NewFile(utils.PackedFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal(t.packedJson, f)
	c.Assert(err, check.IsNil)
	_, err = f.GetRecord(lib.TrailerRecordName)
	c.Assert(err, check.IsNil)
	_, err = f.GetRecord(lib.BaseSegmentName)
	c.Assert(err, check.NotNil)
}

func (t *FileTest) TestCreateFile(c *check.C) {
	_, err := CreateFile(t.packedJson)
	c.Assert(err, check.IsNil)

	f, err := CreateFile(t.packedRaw)
	c.Assert(err, check.IsNil)

	raw, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", "packed_file.dat"))
	c.Assert(err, check.IsNil)

	c.Assert(strings.Compare(f.String(false), string(raw)), check.Equals, 0)
}

func (t *FileTest) TestNewFileFromReader(c *check.C) {
	_, err := NewFileFromReader(t.packedJsonReader)
	c.Assert(err, check.IsNil)

	f, err := NewFileFromReader(t.packedRawReader)
	c.Assert(err, check.IsNil)

	raw, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", "packed_file.dat"))
	c.Assert(err, check.IsNil)

	c.Assert(strings.Compare(f.String(false), string(raw)), check.Equals, 0)
}

func (t *FileTest) TestCreateFileFailed(c *check.C) {

	r1 := bytes.NewReader(t.packedRaw[8:])
	c.Assert(r1, check.NotNil)

	_, err := NewFileFromReader(r1)
	c.Assert(err, check.NotNil)

	data := `{
  "header": {
    "recordDescriptorWord": 480,
    "recordIdentifier": "error",
  }
}`
	r2 := bytes.NewReader([]byte(data))

	_, err = NewFileFromReader(r2)
	c.Assert(err, check.NotNil)
}

func (t *FileTest) TestWithUnknownFileType(c *check.C) {
	_, err := NewFile("unknown")
	c.Assert(err, check.NotNil)
}

func TestFile__Reader(t *testing.T) {

	t.Run("Read with unpacked fixed file", func(t *testing.T) {

		fd, err := os.Open(filepath.Join("..", "..", "test", "testdata", "unpacked_fixed_file.dat"))
		if err != nil {
			t.Fatalf("Can not open local file: %s: \n", err)
		}
		defer fd.Close()

		f, err := NewReader(fd).Read()
		require.NoError(t, err)

		// ensure we have a validated file structure
		err = f.Validate()
		require.NoError(t, err)
	})

	t.Run("Read with unpacked variable file", func(t *testing.T) {

		fd, err := os.Open(filepath.Join("..", "..", "test", "testdata", "unpacked_variable_file.dat"))
		if err != nil {
			t.Fatalf("Can not open local file: %s: \n", err)
		}
		defer fd.Close()

		f, err := NewReader(fd).Read()
		require.NoError(t, err)

		// ensure we have a validated file structure
		err = f.Validate()
		require.NoError(t, err)
	})

	t.Run("Read with packed file", func(t *testing.T) {

		fd, err := os.Open(filepath.Join("..", "..", "test", "testdata", "packed_file.dat"))
		if err != nil {
			t.Fatalf("Can not open local file: %s: \n", err)
		}
		defer fd.Close()

		f, err := NewReader(fd).Read()
		require.NoError(t, err)

		// ensure we have a validated file structure
		err = f.Validate()
		require.NoError(t, err)
	})
}
