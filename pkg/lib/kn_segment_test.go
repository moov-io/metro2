// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"bytes"
	"gopkg.in/check.v1"
)

func (t *SegmentTest) TestK1Segment(c *check.C) {
	segment := NewK1Segment()
	_, err := segment.Parse(t.sampleK1Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.sampleK1Segment))
	c.Assert(segment.Name(), check.Equals, K1SegmentName)
	c.Assert(segment.Length(), check.Equals, K1SegmentLength)
}

func (t *SegmentTest) TestK1SegmentWithInvalidData(c *check.C) {
	segment := NewK1Segment()
	_, err := segment.Parse(append([]byte("ERROR"), t.sampleK1Segment...))
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestPackedK1SegmentWithInvalidCreditorClassification(c *check.C) {
	segment := &K1Segment{}
	_, err := segment.Parse(t.sampleK1Segment)
	c.Assert(err, check.IsNil)
	segment.CreditorClassification = 22
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "creditor classification in k1 segment has an invalid value")
}

func (t *SegmentTest) TestK2Segment(c *check.C) {
	segment := NewK2Segment()
	_, err := segment.Parse(t.sampleK2Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.sampleK2Segment))
	c.Assert(segment.Name(), check.Equals, K2SegmentName)
	c.Assert(segment.Length(), check.Equals, K2SegmentLength)
}

func (t *SegmentTest) TestK2SegmentWithInvalidData(c *check.C) {
	segment := NewK2Segment()
	_, err := segment.Parse(append([]byte("ERROR"), t.sampleK2Segment...))
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestK2SegmentWithInvalidPurchasedIndicator(c *check.C) {
	segment := &K2Segment{}
	_, err := segment.Parse(t.sampleK2Segment)
	c.Assert(err, check.IsNil)
	segment.PurchasedIndicator = 3
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "purchased indicator in k2 segment has an invalid value")
}

func (t *SegmentTest) TestK2SegmentWithInvalidPurchasedName(c *check.C) {
	segment := &K2Segment{}
	_, err := segment.Parse(t.sampleK2Segment)
	c.Assert(err, check.IsNil)
	segment.PurchasedName = "err"
	segment.PurchasedIndicator = PurchasedIndicatorRemove
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "purchased name in k2 segment has an invalid value")
}

func (t *SegmentTest) TestK3Segment(c *check.C) {
	segment := NewK3Segment()
	_, err := segment.Parse(t.sampleK3Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.sampleK3Segment))
	c.Assert(segment.Name(), check.Equals, K3SegmentName)
	c.Assert(segment.Length(), check.Equals, K3SegmentLength)
}

func (t *SegmentTest) TestK3SegmentWithInvalidData(c *check.C) {
	segment := NewK3Segment()
	_, err := segment.Parse(append([]byte("ERROR"), t.sampleK3Segment...))
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestK3SegmentWithInvalidAgencyIdentifier(c *check.C) {
	segment := &K3Segment{}
	_, err := segment.Parse(t.sampleK3Segment)
	c.Assert(err, check.IsNil)
	segment.AgencyIdentifier = 5
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "agency identifier in k3 segment has an invalid value")
}

func (t *SegmentTest) TestK3SegmentWithInvalidAccountNumber(c *check.C) {
	segment := &K3Segment{}
	_, err := segment.Parse(t.sampleK3Segment)
	c.Assert(err, check.IsNil)
	segment.AccountNumber = "error"
	segment.AgencyIdentifier = AgencyIdentifierNotApplicable
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "account number in k3 segment has an invalid value")
}

func (t *SegmentTest) TestK4Segment(c *check.C) {
	segment := NewK4Segment()
	_, err := segment.Parse(t.sampleK4Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.sampleK4Segment))
	c.Assert(segment.Name(), check.Equals, K4SegmentName)
	c.Assert(segment.Length(), check.Equals, K4SegmentLength)
}

func (t *SegmentTest) TestK4SegmentWithInvalidData(c *check.C) {
	segment := NewK4Segment()
	_, err := segment.Parse(append([]byte("ERROR"), t.sampleK4Segment...))
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestK4SegmentWithInvalidSpecializedPaymentIndicator(c *check.C) {
	segment := &K4Segment{}
	_, err := segment.Parse(t.sampleK4Segment)
	c.Assert(err, check.IsNil)
	segment.SpecializedPaymentIndicator = 3
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "specialized payment indicator in k4 segment has an invalid value")
}

func (t *SegmentTest) TestNSegmentWithInvalidData(c *check.C) {
	_, err := NewK1Segment().Parse(t.sampleK1Segment[:16])
	c.Assert(err, check.Not(check.IsNil))
	_, err = NewK2Segment().Parse(t.sampleK2Segment[:16])
	c.Assert(err, check.Not(check.IsNil))
	_, err = NewK3Segment().Parse(t.sampleK3Segment[:16])
	c.Assert(err, check.Not(check.IsNil))
	_, err = NewK4Segment().Parse(t.sampleK4Segment[:16])
	c.Assert(err, check.Not(check.IsNil))
}
