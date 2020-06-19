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
	"strings"
	"testing"

	"github.com/moov-io/metro2/utils"
)

func Test(t *testing.T) { check.TestingT(t) }

type FileTest struct {
	unpackedFixedLengthFile     string
	unpackedFixedLengthJson     string
	unpackedVariableBlockedFile string
	unpackedVariableBlockedJson string
}

var _ = check.Suite(&FileTest{})

func (t *FileTest) SetUpSuite(c *check.C) {
	f, err := os.Open(filepath.Join("..", "testdata", "unpacked_fixed_file.dat"))
	c.Assert(err, check.IsNil)
	t.unpackedFixedLengthFile = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "unpacked_variable_file.dat"))
	c.Assert(err, check.IsNil)
	t.unpackedVariableBlockedFile = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "unpacked_fixed_file.json"))
	c.Assert(err, check.IsNil)
	t.unpackedFixedLengthJson = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "unpacked_variable_file.json"))
	c.Assert(err, check.IsNil)
	t.unpackedVariableBlockedJson = utils.ReadFile(f)
}

func (t *FileTest) TearDownSuite(c *check.C) {}

func (t *FileTest) SetUpTest(c *check.C) {}

func (t *FileTest) TearDownTest(c *check.C) {}

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

func (t *FileTest) TestJsonWithUnpackedFixedLength(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal([]byte(t.unpackedFixedLengthJson), f)
	c.Assert(err, check.IsNil)
	c.Assert(f.String(), check.Equals, t.unpackedFixedLengthFile)
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
}

func (t *FileTest) TestParseWithUnpackedVariableBlockedFileParse(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.unpackedVariableBlockedFile)
	c.Assert(err, check.IsNil)
}
