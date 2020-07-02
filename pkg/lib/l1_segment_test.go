// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"gopkg.in/check.v1"
)

func (t *SegmentTest) TestL1Segment(c *check.C) {
	segment := NewL1Segment()
	_, err := segment.Parse(t.sampleL1Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, t.sampleL1Segment)
	c.Assert(segment.Name(), check.Equals, L1SegmentName)
	c.Assert(segment.Length(), check.Equals, L1SegmentLength)
}

func (t *SegmentTest) TestL1SegmentWithInvalidData(c *check.C) {
	segment := NewL1Segment()
	_, err := segment.Parse("ERROR" + t.sampleL1Segment)
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestL1SegmentWithInvalidNewConsumerAccountNumber(c *check.C) {
	segment := L1Segment{}
	_, err := segment.Parse(t.sampleL1Segment)
	c.Assert(err, check.IsNil)
	segment.NewConsumerAccountNumber = "error"
	segment.ChangeIndicator = ChangeIndicatorIdentificationNumber
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of new consumer account number")
}

func (t *SegmentTest) TestL1SegmentWithInvalidNewIdentificationNumber(c *check.C) {
	segment := L1Segment{}
	_, err := segment.Parse(t.sampleL1Segment)
	c.Assert(err, check.IsNil)
	segment.NewIdentificationNumber = "error"
	segment.ChangeIndicator = ChangeIndicatorAccountNumber
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of new identification number")
}

func (t *SegmentTest) TestL1SegmentWithInvalidChangeIndicator(c *check.C) {
	segment := L1Segment{}
	_, err := segment.Parse(t.sampleL1Segment)
	c.Assert(err, check.IsNil)
	segment.ChangeIndicator = 5
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of change indicator")
}

func (t *SegmentTest) TestL1SegmentWithInvalidData2(c *check.C) {
	_, err := NewL1Segment().Parse(t.sampleL1Segment[:16])
	c.Assert(err, check.Not(check.IsNil))
}
