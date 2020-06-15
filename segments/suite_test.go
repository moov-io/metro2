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
	sampleBaseSegment         string
	samplePackedBaseSegment   string
	sampleHeaderRecord        string
	samplePackedHeaderRecord  string
	sampleTrailerRecord       string
	samplePackedTrailerRecord string
	sampleJ1Segment           string
	sampleJ2Segment           string
	sampleK1Segment           string
	sampleK2Segment           string
	sampleK3Segment           string
	sampleK4Segment           string
	sampleL1Segment           string
	sampleN1Segment           string
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

	f, err = os.Open(filepath.Join("..", "testdata", "packed_base_segment.dat"))
	c.Assert(err, check.IsNil)
	s.samplePackedBaseSegment = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "header_record.dat"))
	c.Assert(err, check.IsNil)
	s.sampleHeaderRecord = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "packed_header_record.dat"))
	c.Assert(err, check.IsNil)
	s.samplePackedHeaderRecord = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "trailer_record.dat"))
	c.Assert(err, check.IsNil)
	s.sampleTrailerRecord = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "packed_trailer_record.dat"))
	c.Assert(err, check.IsNil)
	s.samplePackedTrailerRecord = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "j1_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleJ1Segment = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "j2_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleJ2Segment = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "k1_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleK1Segment = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "k2_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleK2Segment = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "k3_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleK3Segment = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "k4_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleK4Segment = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "l1_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleL1Segment = readStringFromFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "n1_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleN1Segment = readStringFromFile(f)
}

func (s *SegmentTest) TearDownSuite(c *check.C) {}

func (s *SegmentTest) SetUpTest(c *check.C) {}

func (s *SegmentTest) TearDownTest(c *check.C) {}
