// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"bytes"
	"gopkg.in/check.v1"
)

func (t *SegmentTest) TestL1Segment(c *check.C) {
	segment := NewL1Segment()
	_, err := segment.Parse(t.sampleL1Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.sampleL1Segment))
	c.Assert(segment.Name(), check.Equals, L1SegmentName)
	c.Assert(segment.Length(), check.Equals, L1SegmentLength)
}

func (t *SegmentTest) TestL1SegmentWithInvalidData(c *check.C) {
	segment := NewL1Segment()
	_, err := segment.Parse(append([]byte("ERROR"), t.sampleL1Segment...))
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
	c.Assert(err.Error(), check.DeepEquals, "new consumer account number in l1 segment has an invalid value")
}

func (t *SegmentTest) TestL1SegmentWithInvalidNewIdentificationNumber(c *check.C) {
	segment := L1Segment{}
	_, err := segment.Parse(t.sampleL1Segment)
	c.Assert(err, check.IsNil)
	segment.NewIdentificationNumber = "error"
	segment.ChangeIndicator = ChangeIndicatorAccountNumber
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "new identification number in l1 segment has an invalid value")
}

func (t *SegmentTest) TestL1SegmentWithInvalidChangeIndicator(c *check.C) {
	segment := L1Segment{}
	_, err := segment.Parse(t.sampleL1Segment)
	c.Assert(err, check.IsNil)
	segment.ChangeIndicator = 5
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "change indicator in l1 segment has an invalid value")
}

func (t *SegmentTest) TestL1SegmentWithInvalidData2(c *check.C) {
	_, err := NewL1Segment().Parse(t.sampleL1Segment[:16])
	c.Assert(err, check.Not(check.IsNil))
}
