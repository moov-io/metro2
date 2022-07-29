// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"bytes"
	"encoding/json"
	"testing"

	"gopkg.in/check.v1"
)

func TestBaseSegmentErr(t *testing.T) {
	record := &BaseSegment{}
	if _, err := record.Parse([]byte("12345")); err == nil {
		t.Error("expected error")
	}
}

func (t *SegmentTest) TestBaseSegment(c *check.C) {
	segment := NewBaseSegment()
	_, err := segment.Parse(t.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.sampleBaseSegment))
	c.Assert(segment.Name(), check.Equals, BaseSegmentName)
	c.Assert(segment.Length(), check.Equals, 1264)
	c.Assert(segment.BlockSize(), check.Equals, 1268)

	list := segment.GetSegments(J1SegmentName)
	c.Assert(len(list), check.Equals, 1)
	list = segment.GetSegments(J2SegmentName)
	c.Assert(len(list), check.Equals, 2)
	list = segment.GetSegments(K1SegmentName)
	c.Assert(len(list), check.Equals, 1)
	list = segment.GetSegments(K2SegmentName)
	c.Assert(len(list), check.Equals, 1)
	list = segment.GetSegments(K3SegmentName)
	c.Assert(len(list), check.Equals, 1)
	list = segment.GetSegments(K4SegmentName)
	c.Assert(len(list), check.Equals, 1)
	list = segment.GetSegments(L1SegmentName)
	c.Assert(len(list), check.Equals, 1)
	list = segment.GetSegments(N1SegmentName)
	c.Assert(len(list), check.Equals, 1)
	list = segment.GetSegments("unknown")
	c.Assert(len(list), check.Equals, 0)
}

func (t *SegmentTest) TestBaseSegmentWithInvalidData(c *check.C) {
	segment := NewBaseSegment()
	_, err := segment.Parse(append([]byte("ERROR"), t.sampleBaseSegment...))
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestBaseSegmentWithIdentificationNumber(c *check.C) {
	segment := &BaseSegment{}
	_, err := segment.Parse(t.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	segment.IdentificationNumber = ""
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestBaseSegmentWithInvalidPortfolioType(c *check.C) {
	segment := &BaseSegment{}
	_, err := segment.Parse(t.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	segment.PortfolioType = "A"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "portfolio type in base segment has an invalid value")
}

func (t *SegmentTest) TestBaseSegmentWithInvalidTermsDuration(c *check.C) {
	segment := &BaseSegment{}
	_, err := segment.Parse(t.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	segment.TermsDuration = "AAA"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "terms duration in base segment has an invalid value")
}

func (t *SegmentTest) TestBaseSegmentWithInvalidPaymentHistoryProfile(c *check.C) {
	segment := &BaseSegment{}
	_, err := segment.Parse(t.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	segment.PaymentHistoryProfile = "Z"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "payment history profile in base segment has an invalid value")
}

func (t *SegmentTest) TestBaseSegmentWithInvalidInterestTypeIndicator(c *check.C) {
	segment := &BaseSegment{}
	_, err := segment.Parse(t.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	segment.InterestTypeIndicator = "Z"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "interest type indicator in base segment has an invalid value")
}

func (t *SegmentTest) TestBaseSegmentWithInvalidTelephoneNumber(c *check.C) {
	segment := &BaseSegment{}
	_, err := segment.Parse(t.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	segment.TelephoneNumber = 0
	err = segment.Validate()
	c.Assert(err, check.IsNil)
}

func (t *SegmentTest) TestPackedBaseSegment(c *check.C) {
	segment := NewPackedBaseSegment()
	_, err := segment.Parse(t.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.samplePackedBaseSegment))
	c.Assert(segment.Name(), check.Equals, PackedBaseSegmentName)
	c.Assert(segment.Length(), check.Equals, 1106)
	c.Assert(segment.BlockSize(), check.Equals, 1110)

	list := segment.GetSegments(J1SegmentName)
	c.Assert(len(list), check.Equals, 1)
	list = segment.GetSegments(J2SegmentName)
	c.Assert(len(list), check.Equals, 2)
	list = segment.GetSegments(K1SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = segment.GetSegments(K2SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = segment.GetSegments(K3SegmentName)
	c.Assert(len(list), check.Equals, 1)
	list = segment.GetSegments(K4SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = segment.GetSegments(L1SegmentName)
	c.Assert(len(list), check.Equals, 1)
	list = segment.GetSegments(N1SegmentName)
	c.Assert(len(list), check.Equals, 1)
	list = segment.GetSegments("unknown")
	c.Assert(len(list), check.Equals, 0)
}

func (t *SegmentTest) TestPackedBaseSegmentWithInvalidData(c *check.C) {
	segment := NewPackedBaseSegment()
	_, err := segment.Parse(append([]byte("ERROR"), t.samplePackedBaseSegment...))
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestPackedBaseSegmentWithIdentificationNumber(c *check.C) {
	segment := &PackedBaseSegment{}
	_, err := segment.Parse(t.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	segment.IdentificationNumber = ""
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestPackedBaseSegmentWithInvalidPortfolioType(c *check.C) {
	segment := &PackedBaseSegment{}
	_, err := segment.Parse(t.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	segment.PortfolioType = "A"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "portfolio type in packed base segment has an invalid value")
}

func (t *SegmentTest) TestPackedBaseSegmentWithInvalidTermsDuration(c *check.C) {
	segment := &PackedBaseSegment{}
	_, err := segment.Parse(t.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	segment.TermsDuration = "AAA"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "terms duration in packed base segment has an invalid value")
}

func (t *SegmentTest) TestPackedBaseSegmentWithInvalidPaymentHistoryProfile(c *check.C) {
	segment := &PackedBaseSegment{}
	_, err := segment.Parse(t.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	segment.PaymentHistoryProfile = "Z"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "payment history profile in packed base segment has an invalid value")
}

func (t *SegmentTest) TestPackedBaseSegmentWithInvalidInterestTypeIndicator(c *check.C) {
	segment := &PackedBaseSegment{}
	_, err := segment.Parse(t.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	segment.InterestTypeIndicator = "Z"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "interest type indicator in packed base segment has an invalid value")
}

func (t *SegmentTest) TestPackedBaseSegmentWithInvalidTelephoneNumber(c *check.C) {
	segment := &PackedBaseSegment{}
	_, err := segment.Parse(t.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	segment.TelephoneNumber = 0
	err = segment.Validate()
	c.Assert(err, check.IsNil)
}

func (t *SegmentTest) TestBaseRecordApplicableSegment(c *check.C) {
	f := NewBaseSegment()
	jsonStr := `{
		  "segmentIdentifier": "J1",
		  "surname": "BEAUCHAMP",
		  "firstName": "KEVIN",
		  "generationCode": "S",
		  "socialSecurityNumber": 445112877,
		  "dateBirth": "2020-01-02T00:00:00Z",
		  "telephoneNumber": 4335552333,
		  "ecoaCode": "2",
		  "consumerInformationIndicator": "R"
		}`
	newSegment := NewJ1Segment()
	err := json.Unmarshal([]byte(jsonStr), &newSegment)
	c.Assert(err, check.IsNil)
	err = f.AddApplicableSegment(newSegment)
	c.Assert(err, check.IsNil)
	err = f.AddApplicableSegment(NewK1Segment())
	c.Assert(err, check.NotNil)
	err = f.AddApplicableSegment(NewK2Segment())
	c.Assert(err, check.NotNil)
	err = f.AddApplicableSegment(NewK4Segment())
	c.Assert(err, check.NotNil)
	list := f.GetSegments(J1SegmentName)
	c.Assert(len(list), check.Equals, 1)
	list = f.GetSegments(J2SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = f.GetSegments(K1SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = f.GetSegments(K2SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = f.GetSegments(K3SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = f.GetSegments(K4SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = f.GetSegments(L1SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = f.GetSegments(N1SegmentName)
	c.Assert(len(list), check.Equals, 0)
}

func (t *SegmentTest) TestBaseRecordApplicableSingleSegment(c *check.C) {
	f := NewBaseSegment()
	jsonStr := `{
		"segmentIdentifier": "K3",
		"mortgageIdentificationNumber": "Mortgage Number"
	  }`
	newSegment := NewK3Segment()
	err := json.Unmarshal([]byte(jsonStr), &newSegment)
	c.Assert(err, check.IsNil)
	err = f.AddApplicableSegment(newSegment)
	c.Assert(err, check.IsNil)
	list := f.GetSegments(K3SegmentName)
	c.Assert(len(list), check.Equals, 1)
}

func (t *SegmentTest) TestBaseSegmentJson(c *check.C) {
	segment := NewBaseSegment()
	_, err := segment.Parse(t.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	buf, err := json.Marshal(segment)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal(buf, segment)
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.sampleBaseSegment))
	c.Assert(segment.Name(), check.Equals, BaseSegmentName)
}

func (t *SegmentTest) TestPackedBaseRecordApplicableSegment(c *check.C) {
	f := NewPackedBaseSegment()
	jsonStr := `{
		  "segmentIdentifier": "J1",
		  "surname": "BEAUCHAMP",
		  "firstName": "KEVIN",
		  "generationCode": "S",
		  "socialSecurityNumber": 445112877,
		  "dateBirth": "2020-01-02T00:00:00Z",
		  "telephoneNumber": 4335552333,
		  "ecoaCode": "2",
		  "consumerInformationIndicator": "R"
		}`
	newSegment := NewJ1Segment()
	err := json.Unmarshal([]byte(jsonStr), &newSegment)
	c.Assert(err, check.IsNil)
	err = f.AddApplicableSegment(newSegment)
	c.Assert(err, check.IsNil)
	list := f.GetSegments(J1SegmentName)
	c.Assert(len(list), check.Equals, 1)
	list = f.GetSegments(J2SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = f.GetSegments(K1SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = f.GetSegments(K2SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = f.GetSegments(K3SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = f.GetSegments(K4SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = f.GetSegments(L1SegmentName)
	c.Assert(len(list), check.Equals, 0)
	list = f.GetSegments(N1SegmentName)
	c.Assert(len(list), check.Equals, 0)
}

func (t *SegmentTest) TestPackedBaseRecordApplicableSingleSegment(c *check.C) {
	f := NewPackedBaseSegment()
	jsonStr := `{
		"segmentIdentifier": "K3",
		"mortgageIdentificationNumber": "Mortgage Number"
	  }`
	newSegment := NewK3Segment()
	err := json.Unmarshal([]byte(jsonStr), &newSegment)
	c.Assert(err, check.IsNil)
	err = f.AddApplicableSegment(newSegment)
	c.Assert(err, check.IsNil)
	list := f.GetSegments(K3SegmentName)
	c.Assert(len(list), check.Equals, 1)
}

func (t *SegmentTest) TestPackedBaseSegmentJson(c *check.C) {
	segment := NewPackedBaseSegment()
	_, err := segment.Parse(t.samplePackedBaseSegment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	buf, err := json.Marshal(segment)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal(buf, segment)
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.samplePackedBaseSegment))
	c.Assert(segment.Name(), check.Equals, PackedBaseSegmentName)
}
