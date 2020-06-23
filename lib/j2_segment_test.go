// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"gopkg.in/check.v1"
)

func (t *SegmentTest) TestJ2Segment(c *check.C) {
	segment := NewJ2Segment()
	_, err := segment.Parse(t.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, t.sampleJ2Segment)
	c.Assert(segment.Name(), check.Equals, J2SegmentName)
}

func (t *SegmentTest) TestJ2SegmentWithInvalidData(c *check.C) {
	segment := NewJ2Segment()
	_, err := segment.Parse("ERROR" + t.sampleJ2Segment)
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestJ2SegmentWithInvalidGenerationCode(c *check.C) {
	segment := J2Segment{}
	_, err := segment.Parse(t.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	segment.GenerationCode = "0"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of generation code")
}

func (t *SegmentTest) TestJ2SegmentWithInvalidTelephoneNumber(c *check.C) {
	segment := &J2Segment{}
	_, err := segment.Parse(t.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	segment.TelephoneNumber = 0
	err = segment.Validate()
	c.Assert(err, check.IsNil)
}

func (t *SegmentTest) TestJ2SegmentWithInvalidAddressIndicator(c *check.C) {
	segment := J2Segment{}
	_, err := segment.Parse(t.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	segment.AddressIndicator = "0"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of address indicator")
}

func (t *SegmentTest) TestJ2SegmentWithInvalidResidenceCode(c *check.C) {
	segment := J2Segment{}
	_, err := segment.Parse(t.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	segment.ResidenceCode = "0"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of residence code")
}
