package segments

import (
	"gopkg.in/check.v1"
)

func (s *SegmentTest) TestHeaderRecord(c *check.C) {
	segment := NewHeaderRecord()
	_, err := segment.Parse(s.sampleHeaderRecord)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.sampleHeaderRecord)
	c.Assert(segment.Description(), check.Equals, HeaderRecordDescription)
}

func (s *SegmentTest) TestHeaderRecordWithInvalidData(c *check.C) {
	segment := NewHeaderRecord()
	_, err := segment.Parse("ERROR" + s.sampleHeaderRecord)
	c.Assert(err, check.Not(check.IsNil))
}

func (s *SegmentTest) TestPackedHeaderRecord(c *check.C) {
	segment := NewPackedHeaderRecord()
	_, err := segment.Parse(s.samplePackedHeaderRecord)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.samplePackedHeaderRecord)
	c.Assert(segment.Description(), check.Equals, PackedHeaderRecordDescription)
}

func (s *SegmentTest) TestPackedHeaderRecordWithInvalidData(c *check.C) {
	segment := NewPackedHeaderRecord()
	_, err := segment.Parse("ERROR" + s.samplePackedHeaderRecord)
	c.Assert(err, check.Not(check.IsNil))
}
