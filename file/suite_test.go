// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"encoding/json"
	"gopkg.in/check.v1"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/metro2/segments"
	"github.com/moov-io/metro2/utils"
)

func Test(t *testing.T) { check.TestingT(t) }

type FileTest struct {
	unpackedFixedLengthFile     string
	unpackedVariableBlockedFile string
	unpackedVariableBlockedJson string
}

var _ = check.Suite(&FileTest{})

func (t *FileTest) SetUpSuite(c *check.C) {
	f, err := os.Open(filepath.Join("..", "testdata", "unpacked_fixed_file.dat"))
	c.Assert(err, check.IsNil)
	t.unpackedFixedLengthFile = utils.ReadStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "unpacked_variable_file.dat"))
	c.Assert(err, check.IsNil)
	t.unpackedVariableBlockedFile = utils.ReadStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "unpacked_variable_file.json"))
	c.Assert(err, check.IsNil)
	t.unpackedVariableBlockedJson = utils.ReadStringFromFile(f)
}

func (t *FileTest) TearDownSuite(c *check.C) {}

func (t *FileTest) SetUpTest(c *check.C) {}

func (t *FileTest) TearDownTest(c *check.C) {}

func (t *FileTest) TestFileUnmarshal(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal([]byte(t.unpackedVariableBlockedJson), f)
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestFileMarshal(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	f.AddApplicableSegment(segments.NewN1Segment())
	f.AddApplicableSegment(segments.NewL1Segment())
	jsonStr, err := json.Marshal(f)
	c.Assert(err, check.IsNil)
	var out bytes.Buffer
	err = json.Indent(&out, jsonStr, "", "\t")
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestUnpackedFixedLengthFileParse(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.unpackedFixedLengthFile)
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestUnpackedVariableBlockedFileParse(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.unpackedVariableBlockedFile)
	c.Assert(err, check.IsNil)
}
