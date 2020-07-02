// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/check.v1"

	"github.com/moov-io/metro2/pkg/utils"
)

func Test(t *testing.T) { check.TestingT(t) }

type FileTest struct {
	unpackedFixedLengthFile     string
	unpackedFixedLengthJson     string
	unpackedVariableBlockedFile string
	unpackedVariableBlockedJson string
	packedFile                  string
	packedJson                  string
	baseSegmentJson             string
}

var _ = check.Suite(&FileTest{})

func (t *FileTest) SetUpSuite(c *check.C) {
	f, err := os.Open(filepath.Join("..", "..", "testdata", "unpacked_fixed_file.dat"))
	c.Assert(err, check.IsNil)
	t.unpackedFixedLengthFile = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "testdata", "unpacked_fixed_file.json"))
	c.Assert(err, check.IsNil)
	t.unpackedFixedLengthJson = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "testdata", "unpacked_variable_file.dat"))
	c.Assert(err, check.IsNil)
	t.unpackedVariableBlockedFile = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "testdata", "unpacked_variable_file.json"))
	c.Assert(err, check.IsNil)
	t.unpackedVariableBlockedJson = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "testdata", "packed_file.dat"))
	c.Assert(err, check.IsNil)
	t.packedFile = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "testdata", "packed_file.json"))
	c.Assert(err, check.IsNil)
	t.packedJson = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "testdata", "base_segment.json"))
	c.Assert(err, check.IsNil)
	t.baseSegmentJson = utils.ReadFile(f)
}

func (t *FileTest) TearDownSuite(c *check.C) {}

func (t *FileTest) SetUpTest(c *check.C) {}

func (t *FileTest) TearDownTest(c *check.C) {}
