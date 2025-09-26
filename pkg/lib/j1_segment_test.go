// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"bytes"
	"encoding/json"

	"gopkg.in/check.v1"

	"github.com/moov-io/metro2/pkg/utils"
)

func (t *SegmentTest) TestJ1Segment(c *check.C) {
	segment := NewJ1Segment()
	_, err := segment.Parse(t.sampleJ1Segment, false)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.sampleJ1Segment))
	c.Assert(segment.Name(), check.Equals, J1SegmentName)
	c.Assert(segment.Length(), check.Equals, J1SegmentLength)
}

func (t *SegmentTest) TestJ1SegmentWithInvalidData(c *check.C) {
	segment := NewJ1Segment()
	_, err := segment.Parse(append([]byte("ERROR"), t.sampleJ1Segment...), false)
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestJ1SegmentWithEmptyGenerationCode(c *check.C) {
	jsonStr := `{
      "segmentIdentifier": "J1",
      "surname": "BEAUCHAMP",
      "firstName": "KEVIN",
      "socialSecurityNumber": 445112877,
      "dateBirth": "2020-01-02T00:00:00Z",
      "telephoneNumber": 4335552333,
      "ecoaCode": "2",
      "consumerInformationIndicator": "R"
    }`
	segment := NewJ1Segment()
	err := json.Unmarshal([]byte(jsonStr), &segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.Name(), check.Equals, J1SegmentName)
	c.Assert(segment.Length(), check.Equals, J1SegmentLength)
}

func (t *SegmentTest) TestJ1SegmentWithInvalidGenerationCode(c *check.C) {
	segment := J1Segment{}
	_, err := segment.Parse(t.sampleJ1Segment, false)
	c.Assert(err, check.IsNil)
	segment.GenerationCode = "0"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "generation code in j1 segment has an invalid value")
}

func (t *SegmentTest) TestJ1SegmentWithInvalidTelephoneNumber(c *check.C) {
	segment := &J1Segment{}
	_, err := segment.Parse(t.sampleJ1Segment, false)
	c.Assert(err, check.IsNil)
	segment.TelephoneNumber = 0
	err = segment.Validate()
	c.Assert(err, check.IsNil)
}

func (t *SegmentTest) TestJ1SegmentWithInvalidData2(c *check.C) {
	_, err := NewJ1Segment().Parse(t.sampleJ1Segment[:16], false)
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestJ1SegmentWithSocialSecurityNumber(c *check.C) {
	segment := &J1Segment{}
	_, err := segment.Parse(t.sampleJ1Segment, false)
	c.Assert(err, check.IsNil)

	segment.SocialSecurityNumber = 0
	err = segment.Validate()
	c.Assert(err, check.Equals, nil)

	segment.DateBirth = utils.Time{}
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestJ1SegmentWithDateBirth(c *check.C) {
	segment := &J1Segment{}
	_, err := segment.Parse(t.sampleJ1Segment, false)
	c.Assert(err, check.IsNil)

	segment.DateBirth = utils.Time{}
	err = segment.Validate()
	c.Assert(err, check.Equals, nil)

	segment.SocialSecurityNumber = 0
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
}
