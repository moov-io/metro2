// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"github.com/moov-io/metro2/utils"
	"gopkg.in/check.v1"
	"os"
	"path/filepath"
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

func (s *SegmentTest) SetUpSuite(c *check.C) {
	f, err := os.Open(filepath.Join("..", "testdata", "base_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleBaseSegment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "packed_base_segment.dat"))
	c.Assert(err, check.IsNil)
	s.samplePackedBaseSegment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "header_record.dat"))
	c.Assert(err, check.IsNil)
	s.sampleHeaderRecord = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "packed_header_record.dat"))
	c.Assert(err, check.IsNil)
	s.samplePackedHeaderRecord = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "trailer_record.dat"))
	c.Assert(err, check.IsNil)
	s.sampleTrailerRecord = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "packed_trailer_record.dat"))
	c.Assert(err, check.IsNil)
	s.samplePackedTrailerRecord = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "j1_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleJ1Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "j2_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleJ2Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "k1_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleK1Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "k2_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleK2Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "k3_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleK3Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "k4_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleK4Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "l1_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleL1Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "testdata", "n1_segment.dat"))
	c.Assert(err, check.IsNil)
	s.sampleN1Segment = utils.ReadFile(f)
}

func (s *SegmentTest) TearDownSuite(c *check.C) {}

func (s *SegmentTest) SetUpTest(c *check.C) {}

func (s *SegmentTest) TearDownTest(c *check.C) {}
