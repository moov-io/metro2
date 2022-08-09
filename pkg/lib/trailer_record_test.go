// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"bytes"
	"testing"

	"gopkg.in/check.v1"
)

func TestTrailerRecordErr(t *testing.T) {
	record := &TrailerRecord{}
	if _, err := record.Parse([]byte("12345")); err == nil {
		t.Error("expected error")
	}
	packedRecord := &PackedTrailerRecord{}
	if _, err := packedRecord.Parse([]byte("12345")); err == nil {
		t.Error("expected error")
	}
}

func (t *SegmentTest) TestTrailerRecord(c *check.C) {
	segment := NewTrailerRecord()
	_, err := segment.Parse(t.sampleTrailerRecord)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.sampleTrailerRecord))
	c.Assert(segment.Name(), check.Equals, TrailerRecordName)
	c.Assert(segment.Length(), check.Equals, UnpackedRecordLength)
	c.Assert(segment.BlockSize(), check.Equals, UnpackedRecordLength)
	c.Assert(segment.GetSegments(K1SegmentName), check.IsNil)
	_sub := NewJ1Segment()
	c.Assert(segment.AddApplicableSegment(_sub), check.NotNil)
}

func (t *SegmentTest) TestTrailerRecordWithInvalidData(c *check.C) {
	segment := NewTrailerRecord()
	_, err := segment.Parse(append([]byte("ERROR"), t.sampleTrailerRecord...))
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestPackedTrailerRecord(c *check.C) {
	segment := NewPackedTrailerRecord()
	_, err := segment.Parse(t.samplePackedTrailerRecord)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.samplePackedTrailerRecord))
	c.Assert(segment.Name(), check.Equals, PackedTrailerRecordName)
	c.Assert(segment.Length(), check.Equals, PackedRecordLength)
	c.Assert(segment.BlockSize(), check.Equals, PackedRecordLength)
	c.Assert(segment.GetSegments(K1SegmentName), check.IsNil)
	_sub := NewJ1Segment()
	c.Assert(segment.AddApplicableSegment(_sub), check.NotNil)
}

func (t *SegmentTest) TestPackedTrailerRecordWithInvalidData(c *check.C) {
	segment := NewPackedTrailerRecord()
	_, err := segment.Parse(append([]byte("ERROR"), t.samplePackedTrailerRecord...))
	c.Assert(err, check.Not(check.IsNil))
}
