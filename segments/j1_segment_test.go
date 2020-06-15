package segments

import (
	"gopkg.in/check.v1"

	"github.com/moov-io/metro2/utils"
)

func (s *SegmentTest) TestJ1Segment(c *check.C) {
	segment := NewJ1Segment()
	err := segment.Parse(s.sampleJ1Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.sampleJ1Segment)
	c.Assert(segment.Description(), check.Equals, J1SegmentDescription)
}

func (s *SegmentTest) TestJ1SegmentWithInvalidData(c *check.C) {
	segment := NewJ1Segment()
	err := segment.Parse(s.sampleJ1Segment + "ERROR")
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err, check.DeepEquals, utils.ErrSegmentLength)
}

func (s *SegmentTest) TestJ1SegmentWithInvalidGenerationCode(c *check.C) {
	segment := J1Segment{}
	err := segment.Parse(s.sampleJ1Segment)
	c.Assert(err, check.IsNil)
	segment.GenerationCode = "0"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "is an invalid value of generation code")
}

func (s *SegmentTest) TestJ1SegmentWithInvalidTelephoneNumber(c *check.C) {
	segment := &J1Segment{}
	err := segment.Parse(s.sampleJ1Segment)
	c.Assert(err, check.IsNil)
	segment.TelephoneNumber = 0
	err = segment.Validate()
	c.Assert(err, check.IsNil)
}
