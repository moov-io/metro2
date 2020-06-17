package segments

import (
	"gopkg.in/check.v1"
)

func (s *SegmentTest) TestK1Segment(c *check.C) {
	segment := NewK1Segment()
	_, err := segment.Parse(s.sampleK1Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.sampleK1Segment)
	c.Assert(segment.Description(), check.Equals, K1SegmentDescription)
}

func (s *SegmentTest) TestK1SegmentWithInvalidData(c *check.C) {
	segment := NewK1Segment()
	_, err := segment.Parse("ERROR" + s.sampleK1Segment)
	c.Assert(err, check.Not(check.IsNil))
}

func (s *SegmentTest) TestPackedK1SegmentWithInvalidCreditorClassification(c *check.C) {
	segment := &K1Segment{}
	_, err := segment.Parse(s.sampleK1Segment)
	c.Assert(err, check.IsNil)
	segment.CreditorClassification = 22
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of creditor classification")
}

func (s *SegmentTest) TestK2Segment(c *check.C) {
	segment := NewK2Segment()
	_, err := segment.Parse(s.sampleK2Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.sampleK2Segment)
	c.Assert(segment.Description(), check.Equals, K2SegmentDescription)
}

func (s *SegmentTest) TestK2SegmentWithInvalidData(c *check.C) {
	segment := NewK2Segment()
	_, err := segment.Parse("ERROR" + s.sampleK2Segment)
	c.Assert(err, check.Not(check.IsNil))
}

func (s *SegmentTest) TestK2SegmentWithInvalidPurchasedIndicator(c *check.C) {
	segment := &K2Segment{}
	_, err := segment.Parse(s.sampleK2Segment)
	c.Assert(err, check.IsNil)
	segment.PurchasedIndicator = 3
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of purchased indicator")
}

func (s *SegmentTest) TestK2SegmentWithInvalidPurchasedName(c *check.C) {
	segment := &K2Segment{}
	_, err := segment.Parse(s.sampleK2Segment)
	c.Assert(err, check.IsNil)
	segment.PurchasedName = "err"
	segment.PurchasedIndicator = PurchasedIndicatorRemove
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of purchased name")
}

func (s *SegmentTest) TestK3Segment(c *check.C) {
	segment := NewK3Segment()
	_, err := segment.Parse(s.sampleK3Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.sampleK3Segment)
	c.Assert(segment.Description(), check.Equals, K3SegmentDescription)
}

func (s *SegmentTest) TestK3SegmentWithInvalidData(c *check.C) {
	segment := NewK3Segment()
	_, err := segment.Parse("ERROR" + s.sampleK3Segment)
	c.Assert(err, check.Not(check.IsNil))
}

func (s *SegmentTest) TestK3SegmentWithInvalidAgencyIdentifier(c *check.C) {
	segment := &K3Segment{}
	_, err := segment.Parse(s.sampleK3Segment)
	c.Assert(err, check.IsNil)
	segment.AgencyIdentifier = 5
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of agency identifier")
}

func (s *SegmentTest) TestK3SegmentWithInvalidAccountNumber(c *check.C) {
	segment := &K3Segment{}
	_, err := segment.Parse(s.sampleK3Segment)
	c.Assert(err, check.IsNil)
	segment.AccountNumber = "error"
	segment.AgencyIdentifier = AgencyIdentifierNotApplicable
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of account number")
}

func (s *SegmentTest) TestK4Segment(c *check.C) {
	segment := NewK4Segment()
	_, err := segment.Parse(s.sampleK4Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.sampleK4Segment)
	c.Assert(segment.Description(), check.Equals, K4SegmentDescription)
}

func (s *SegmentTest) TestK4SegmentWithInvalidData(c *check.C) {
	segment := NewK4Segment()
	_, err := segment.Parse("ERROR" + s.sampleK4Segment)
	c.Assert(err, check.Not(check.IsNil))
}

func (s *SegmentTest) TestK4SegmentWithInvalidSpecializedPaymentIndicator(c *check.C) {
	segment := &K4Segment{}
	_, err := segment.Parse(s.sampleK4Segment)
	c.Assert(err, check.IsNil)
	segment.SpecializedPaymentIndicator = 3
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of specialized payment indicator")
}
