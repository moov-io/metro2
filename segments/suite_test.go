// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"bufio"
	"gopkg.in/check.v1"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type SegmentTest struct {
	sampleBaseSegment        string
	samplePackedBaseSegment  string
	sampleHeaderRecord       string
	samplePackedHeaderRecord string
}

var _ = check.Suite(&SegmentTest{})

func readStringFromFile(f *os.File) string {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "")
}

func (s *SegmentTest) SetUpSuite(c *check.C) {
	f, err := os.Open(filepath.Join("..", "testdata", "base_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleBaseSegment = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "header_record.dat"))
	c.Assert(err, check.IsNil)
	s.sampleHeaderRecord = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "packed_header_record.dat"))
	c.Assert(err, check.IsNil)
	s.samplePackedHeaderRecord = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "packed_base_segment.dat"))
	c.Assert(err, check.IsNil)
	s.samplePackedBaseSegment = readStringFromFile(f)
}

func (s *SegmentTest) TearDownSuite(c *check.C) {}

func (s *SegmentTest) SetUpTest(c *check.C) {}

func (s *SegmentTest) TearDownTest(c *check.C) {}
