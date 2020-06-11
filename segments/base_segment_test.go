package segments

import (
	"github.com/moov-io/ach"
	"gopkg.in/check.v1"
)

func (s *SegmentTest) TestBaseSegment(c *check.C) {
	segment := NewBaseSegment()
	err := segment.Parse(s.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.sampleBaseSegment)
	c.Assert(segment.Description(), check.Equals, BaseSegmentDescription)
}

func (s *SegmentTest) TestBaseSegmentWithInvalidData(c *check.C) {
	segment := NewBaseSegment()
	err := segment.Parse(s.sampleBaseSegment + "ERROR")
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err, check.DeepEquals, ErrSegmentLength)
}

func (s *SegmentTest) TestBaseSegmentWithIdentificationNumber(c *check.C) {
	segment := &BaseSegment{}
	err := segment.Parse(s.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	segment.IdentificationNumber = ""
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err, check.DeepEquals, ach.ErrFieldRequired)
}

func (s *SegmentTest) TestBaseSegmentWithInvalidPortfolioType(c *check.C) {
	segment := &BaseSegment{}
	err := segment.Parse(s.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	segment.PortfolioType = "A"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of portfolio type")
}

func (s *SegmentTest) TestBaseSegmentWithInvalidTermsDuration(c *check.C) {
	segment := &BaseSegment{}
	err := segment.Parse(s.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	segment.TermsDuration = "AAA"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of terms duration")
}

func (s *SegmentTest) TestBaseSegmentWithInvalidPaymentHistoryProfile(c *check.C) {
	segment := &BaseSegment{}
	err := segment.Parse(s.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	segment.PaymentHistoryProfile = "Z"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of payment history profile")
}

func (s *SegmentTest) TestBaseSegmentWithInvalidInterestTypeIndicator(c *check.C) {
	segment := &BaseSegment{}
	err := segment.Parse(s.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	segment.InterestTypeIndicator = "Z"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of interest type indicator")
}

func (s *SegmentTest) TestBaseSegmentWithInvalidTelephoneNumber(c *check.C) {
	segment := &BaseSegment{}
	err := segment.Parse(s.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	segment.TelephoneNumber = 0
	err = segment.Validate()
	c.Assert(err, check.IsNil)
}

func (s *SegmentTest) TestPackedBaseSegment(c *check.C) {
	segment := NewPackedBaseSegment()
	err := segment.Parse(s.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.samplePackedBaseSegment)
	c.Assert(segment.Description(), check.Equals, PackedBaseSegmentDescription)
}

func (s *SegmentTest) TestPackedBaseSegmentWithInvalidData(c *check.C) {
	segment := NewPackedBaseSegment()
	err := segment.Parse(s.samplePackedBaseSegment + "ERROR")
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err, check.DeepEquals, ErrSegmentLength)
}

func (s *SegmentTest) TestPackedBaseSegmentWithIdentificationNumber(c *check.C) {
	segment := &PackedBaseSegment{}
	err := segment.Parse(s.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	segment.IdentificationNumber = ""
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err, check.DeepEquals, ach.ErrFieldRequired)
}

func (s *SegmentTest) TestPackedBaseSegmentWithInvalidPortfolioType(c *check.C) {
	segment := &PackedBaseSegment{}
	err := segment.Parse(s.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	segment.PortfolioType = "A"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of portfolio type")
}

func (s *SegmentTest) TestPackedBaseSegmentWithInvalidTermsDuration(c *check.C) {
	segment := &PackedBaseSegment{}
	err := segment.Parse(s.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	segment.TermsDuration = "AAA"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of terms duration")
}

func (s *SegmentTest) TestPackedBaseSegmentWithInvalidPaymentHistoryProfile(c *check.C) {
	segment := &PackedBaseSegment{}
	err := segment.Parse(s.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	segment.PaymentHistoryProfile = "Z"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of payment history profile")
}

func (s *SegmentTest) TestPackedBaseSegmentWithInvalidInterestTypeIndicator(c *check.C) {
	segment := &PackedBaseSegment{}
	err := segment.Parse(s.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	segment.InterestTypeIndicator = "Z"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of interest type indicator")
}

func (s *SegmentTest) TestPackedBaseSegmentWithInvalidTelephoneNumber(c *check.C) {
	segment := &PackedBaseSegment{}
	err := segment.Parse(s.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	segment.TelephoneNumber = 0
	err = segment.Validate()
	c.Assert(err, check.IsNil)
}
