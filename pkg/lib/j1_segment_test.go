// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"bytes"
	"gopkg.in/check.v1"
)

func (t *SegmentTest) TestJ1Segment(c *check.C) {
	segment := NewJ1Segment()
	_, err := segment.Parse(t.sampleJ1Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.sampleJ1Segment))
	c.Assert(segment.Name(), check.Equals, J1SegmentName)
	c.Assert(segment.Length(), check.Equals, J1SegmentLength)
}

func (t *SegmentTest) TestJ1SegmentWithInvalidData(c *check.C) {
	segment := NewJ1Segment()
	_, err := segment.Parse(append([]byte("ERROR"), t.sampleJ1Segment...))
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestJ1SegmentWithInvalidGenerationCode(c *check.C) {
	segment := J1Segment{}
	_, err := segment.Parse(t.sampleJ1Segment)
	c.Assert(err, check.IsNil)
	segment.GenerationCode = "0"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "generation code in j1 segment has an invalid value")
}

func (t *SegmentTest) TestJ1SegmentWithInvalidTelephoneNumber(c *check.C) {
	segment := &J1Segment{}
	_, err := segment.Parse(t.sampleJ1Segment)
	c.Assert(err, check.IsNil)
	segment.TelephoneNumber = 0
	err = segment.Validate()
	c.Assert(err, check.IsNil)
}

func (t *SegmentTest) TestJ1SegmentWithInvalidData2(c *check.C) {
	_, err := NewJ1Segment().Parse(t.sampleJ1Segment[:16])
	c.Assert(err, check.Not(check.IsNil))
}
