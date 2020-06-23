// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"gopkg.in/check.v1"
)

func (t *SegmentTest) TestJ1Segment(c *check.C) {
	segment := NewJ1Segment()
	_, err := segment.Parse(t.sampleJ1Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, t.sampleJ1Segment)
	c.Assert(segment.Name(), check.Equals, J1SegmentName)
}

func (t *SegmentTest) TestJ1SegmentWithInvalidData(c *check.C) {
	segment := NewJ1Segment()
	_, err := segment.Parse("ERROR" + t.sampleJ1Segment)
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestJ1SegmentWithInvalidGenerationCode(c *check.C) {
	segment := J1Segment{}
	_, err := segment.Parse(t.sampleJ1Segment)
	c.Assert(err, check.IsNil)
	segment.GenerationCode = "0"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of generation code")
}

func (t *SegmentTest) TestJ1SegmentWithInvalidTelephoneNumber(c *check.C) {
	segment := &J1Segment{}
	_, err := segment.Parse(t.sampleJ1Segment)
	c.Assert(err, check.IsNil)
	segment.TelephoneNumber = 0
	err = segment.Validate()
	c.Assert(err, check.IsNil)
}
