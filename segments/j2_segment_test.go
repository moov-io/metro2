package segments

import (
	"gopkg.in/check.v1"

	"github.com/moov-io/metro2/utils"
)

func (s *SegmentTest) TestJ2Segment(c *check.C) {
	segment := NewJ2Segment()
	err := segment.Parse(s.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.sampleJ2Segment)
	c.Assert(segment.Description(), check.Equals, J2SegmentDescription)
}

func (s *SegmentTest) TestJ2SegmentWithInvalidData(c *check.C) {
	segment := NewJ2Segment()
	err := segment.Parse(s.sampleJ2Segment + "ERROR")
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err, check.DeepEquals, utils.ErrSegmentLength)
}

func (s *SegmentTest) TestJ2SegmentWithInvalidGenerationCode(c *check.C) {
	segment := J2Segment{}
	err := segment.Parse(s.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	segment.GenerationCode = "0"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of generation code")
}

func (s *SegmentTest) TestJ2SegmentWithInvalidTelephoneNumber(c *check.C) {
	segment := &J2Segment{}
	err := segment.Parse(s.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	segment.TelephoneNumber = 0
	err = segment.Validate()
	c.Assert(err, check.IsNil)
}

func (s *SegmentTest) TestJ2SegmentWithInvalidAddressIndicator(c *check.C) {
	segment := J2Segment{}
	err := segment.Parse(s.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	segment.AddressIndicator = "0"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of address indicator")
}

func (s *SegmentTest) TestJ2SegmentWithInvalidResidenceCode(c *check.C) {
	segment := J2Segment{}
	err := segment.Parse(s.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	segment.ResidenceCode = "0"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of residence code")
}
