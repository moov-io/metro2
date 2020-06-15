package segments

import (
	"gopkg.in/check.v1"

	"github.com/moov-io/metro2/utils"
)

func (s *SegmentTest) TestL1Segment(c *check.C) {
	segment := NewL1Segment()
	err := segment.Parse(s.sampleL1Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.sampleL1Segment)
	c.Assert(segment.Description(), check.Equals, L1SegmentDescription)
}

func (s *SegmentTest) TestL1SegmentWithInvalidData(c *check.C) {
	segment := NewL1Segment()
	err := segment.Parse(s.sampleL1Segment + "ERROR")
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err, check.DeepEquals, utils.ErrSegmentLength)
}

func (s *SegmentTest) TestL1SegmentWithInvalidNewConsumerAccountNumber(c *check.C) {
	segment := L1Segment{}
	err := segment.Parse(s.sampleL1Segment)
	c.Assert(err, check.IsNil)
	segment.NewConsumerAccountNumber = "error"
	segment.ChangeIndicator = ChangeIndicatorIdentificationNumber
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of new consumer account number")
}

func (s *SegmentTest) TestL1SegmentWithInvalidNewIdentificationNumber(c *check.C) {
	segment := L1Segment{}
	err := segment.Parse(s.sampleL1Segment)
	c.Assert(err, check.IsNil)
	segment.NewIdentificationNumber = "error"
	segment.ChangeIndicator = ChangeIndicatorAccountNumber
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of new identification number")
}

func (s *SegmentTest) TestL1SegmentWithInvalidChangeIndicator(c *check.C) {
	segment := L1Segment{}
	err := segment.Parse(s.sampleL1Segment)
	c.Assert(err, check.IsNil)
	segment.ChangeIndicator = 5
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of change indicator")
}
