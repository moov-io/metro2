// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/metro2/pkg/utils"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type SegmentTest struct {
	sampleBaseSegment         []byte
	samplePackedBaseSegment   []byte
	sampleHeaderRecord        []byte
	samplePackedHeaderRecord  []byte
	sampleTrailerRecord       []byte
	samplePackedTrailerRecord []byte
	sampleJ1Segment           []byte
	sampleJ2Segment           []byte
	sampleK1Segment           []byte
	sampleK2Segment           []byte
	sampleK3Segment           []byte
	sampleK4Segment           []byte
	sampleL1Segment           []byte
	sampleN1Segment           []byte
}

var _ = check.Suite(&SegmentTest{})

func (t *SegmentTest) SetUpSuite(c *check.C) {
	f, err := os.Open(filepath.Join("..", "..", "test", "testdata", "base_segment.dat"))
	c.Assert(err, check.IsNil)
	t.sampleBaseSegment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "packed_base_segment.dat"))
	c.Assert(err, check.IsNil)
	t.samplePackedBaseSegment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "header_record.dat"))
	c.Assert(err, check.IsNil)
	t.sampleHeaderRecord = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "packed_header_record.dat"))
	c.Assert(err, check.IsNil)
	t.samplePackedHeaderRecord = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "trailer_record.dat"))
	c.Assert(err, check.IsNil)
	t.sampleTrailerRecord = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "packed_trailer_record.dat"))
	c.Assert(err, check.IsNil)
	t.samplePackedTrailerRecord = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "j1_segment.dat"))
	c.Assert(err, check.IsNil)
	t.sampleJ1Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "j2_segment.dat"))
	c.Assert(err, check.IsNil)
	t.sampleJ2Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "k1_segment.dat"))
	c.Assert(err, check.IsNil)
	t.sampleK1Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "k2_segment.dat"))
	c.Assert(err, check.IsNil)
	t.sampleK2Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "k3_segment.dat"))
	c.Assert(err, check.IsNil)
	t.sampleK3Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "k4_segment.dat"))
	c.Assert(err, check.IsNil)
	t.sampleK4Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "l1_segment.dat"))
	c.Assert(err, check.IsNil)
	t.sampleL1Segment = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "n1_segment.dat"))
	c.Assert(err, check.IsNil)
	t.sampleN1Segment = utils.ReadFile(f)
}

func (t *SegmentTest) TearDownSuite(c *check.C) {}

func (t *SegmentTest) SetUpTest(c *check.C) {}

func (t *SegmentTest) TearDownTest(c *check.C) {}
