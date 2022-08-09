// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"bytes"
	"testing"

	"gopkg.in/check.v1"
)

func TestHeaderRecordErr(t *testing.T) {
	record := &HeaderRecord{}
	if _, err := record.Parse([]byte("12345")); err == nil {
		t.Error("expected error")
	}
}

func (t *SegmentTest) TestHeaderRecord(c *check.C) {
	segment := NewHeaderRecord()
	_, err := segment.Parse(t.sampleHeaderRecord)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.sampleHeaderRecord))
	c.Assert(segment.Name(), check.Equals, HeaderRecordName)
	c.Assert(segment.Length(), check.Equals, UnpackedRecordLength)
	c.Assert(segment.BlockSize(), check.Equals, UnpackedRecordLength+4)
	c.Assert(segment.GetSegments(K1SegmentName), check.IsNil)
	_sub := NewJ1Segment()
	c.Assert(segment.AddApplicableSegment(_sub), check.NotNil)
}

func (t *SegmentTest) TestHeaderRecordWithInvalidData(c *check.C) {
	segment := NewHeaderRecord()
	_, err := segment.Parse(append([]byte("ERROR"), t.sampleHeaderRecord...))
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestPackedHeaderRecord(c *check.C) {
	segment := NewPackedHeaderRecord()
	_, err := segment.Parse(t.samplePackedHeaderRecord)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.samplePackedHeaderRecord))
	c.Assert(segment.Name(), check.Equals, PackedHeaderRecordName)
	c.Assert(segment.Length(), check.Equals, PackedRecordLength)
	c.Assert(segment.BlockSize(), check.Equals, PackedRecordLength+4)
	c.Assert(segment.GetSegments(K1SegmentName), check.IsNil)
	_sub := NewJ1Segment()
	c.Assert(segment.AddApplicableSegment(_sub), check.NotNil)
}

func (t *SegmentTest) TestPackedHeaderRecordWithInvalidData(c *check.C) {
	segment := NewPackedHeaderRecord()
	_, err := segment.Parse(append([]byte("ERROR"), t.samplePackedHeaderRecord...))
	c.Assert(err, check.Not(check.IsNil))
}
