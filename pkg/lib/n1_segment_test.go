package lib

import (
	"bytes"
	"gopkg.in/check.v1"
)

func (t *SegmentTest) TestN1Segment(c *check.C) {
	segment := NewN1Segment()
	_, err := segment.Parse(t.sampleN1Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.sampleN1Segment))
	c.Assert(segment.Name(), check.Equals, N1SegmentName)
	c.Assert(segment.Length(), check.Equals, N1SegmentLength)
}

func (t *SegmentTest) TestN1SegmentWithInvalidData(c *check.C) {
	segment := NewN1Segment()
	_, err := segment.Parse(t.sampleN1Segment[2:])
	c.Assert(err, check.Not(check.IsNil))

	_, err = segment.Parse(t.sampleN1Segment[:16])
	c.Assert(err, check.Not(check.IsNil))
}
