package segments

import "gopkg.in/check.v1"

func (s *SegmentTest) TestHeaderRecord(c *check.C) {
	segment := NewHeaderRecord()
	err := segment.Parse(s.sampleHeaderRecord)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.sampleHeaderRecord)
	c.Assert(segment.Description(), check.Equals, HeaderRecordDescription)
}

func (s *SegmentTest) TestHeaderRecordWithInvalidData(c *check.C) {
	segment := NewHeaderRecord()
	err := segment.Parse(s.sampleHeaderRecord + "ERROR")
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err, check.DeepEquals, ErrSegmentInvalidLength)
}

func (s *SegmentTest) TestHeaderRecordWithInvalidActivityDate(c *check.C) {
	segment := &HeaderRecord{}
	err := segment.Parse(s.sampleHeaderRecord)
	c.Assert(err, check.IsNil)
	segment.ActivityDate = 0
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err, check.DeepEquals, ErrRequired)
}

func (s *SegmentTest) TestPackedHeaderRecord(c *check.C) {
	segment := NewPackedHeaderRecord()
	err := segment.Parse(s.samplePackedHeaderRecord)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.samplePackedHeaderRecord)
	c.Assert(segment.Description(), check.Equals, PackedHeaderRecordDescription)
}

func (s *SegmentTest) TestPackedHeaderRecordWithInvalidData(c *check.C) {
	segment := NewPackedHeaderRecord()
	err := segment.Parse(s.samplePackedHeaderRecord + "ERROR")
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err, check.DeepEquals, ErrSegmentInvalidLength)
}

func (s *SegmentTest) TestPackedHeaderRecordWithInvalidActivityDate(c *check.C) {
	segment := &PackedHeaderRecord{}
	err := segment.Parse(s.samplePackedHeaderRecord)
	c.Assert(err, check.IsNil)
	segment.ActivityDate = 0
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err, check.DeepEquals, ErrRequired)
}
