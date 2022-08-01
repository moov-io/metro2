// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/check.v1"

	"github.com/moov-io/metro2/pkg/utils"
)

func Test(t *testing.T) { check.TestingT(t) }

type FileTest struct {
	unpackedFixedLengthRaw      []byte
	unpackedFixedLengthJson     []byte
	unpackedVariableBlockedRaw  []byte
	unpackedVariableBlockedJson []byte
	packedRaw                   []byte
	packedJson                  []byte
	baseSegmentJson             []byte
	packedJsonReader            io.Reader
	packedRawReader             io.Reader
}

var _ = check.Suite(&FileTest{})

func (t *FileTest) SetUpSuite(c *check.C) {
	f, err := os.Open(filepath.Join("..", "..", "test", "testdata", "unpacked_fixed_file.dat"))
	c.Assert(err, check.IsNil)
	t.unpackedFixedLengthRaw = utils.ReadFile(f)
	f.Close()

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "unpacked_fixed_file.json"))
	c.Assert(err, check.IsNil)
	t.unpackedFixedLengthJson = utils.ReadFile(f)
	f.Close()

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "unpacked_variable_file.dat"))
	c.Assert(err, check.IsNil)
	t.unpackedVariableBlockedRaw = utils.ReadFile(f)
	f.Close()

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "unpacked_variable_file.json"))
	c.Assert(err, check.IsNil)
	t.unpackedVariableBlockedJson = utils.ReadFile(f)
	f.Close()

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "packed_file.dat"))
	c.Assert(err, check.IsNil)
	t.packedRaw = utils.ReadFile(f)
	f.Close()

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "packed_file.json"))
	c.Assert(err, check.IsNil)
	t.packedJson = utils.ReadFile(f)
	f.Close()

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "base_segment.json"))
	c.Assert(err, check.IsNil)
	t.baseSegmentJson = utils.ReadFile(f)
	f.Close()

	t.packedJsonReader, err = os.Open(filepath.Join("..", "..", "test", "testdata", "packed_file.json"))
	c.Assert(err, check.IsNil)

	t.packedRawReader, err = os.Open(filepath.Join("..", "..", "test", "testdata", "packed_file.dat"))
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TearDownSuite(c *check.C) {
	err := t.packedJsonReader.(io.ReadCloser).Close()
	c.Assert(err, check.IsNil)

	err = t.packedRawReader.(io.ReadCloser).Close()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) SetUpTest(c *check.C) {}

func (t *FileTest) TearDownTest(c *check.C) {}
