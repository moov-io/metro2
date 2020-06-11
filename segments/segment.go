// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"errors"
)

// General segment interface
type Segment interface {
	Description() string
	Parse(record string) error
	String() string
	Validate() error
}

const (
	// HeaderRecordCharacterLength indicates length of header record
	HeaderRecordCharacterLength = 426
	// BaseSegmentLength indicates length of base segment
	BaseSegmentLength = 426
	// J2SegmentLength indicates length of J2 segment
	J2SegmentLength = 200
	// K1SegmentLength indicates length of K1 segment
	K1SegmentLength = 34
	// K2SegmentLength indicates length of K2 segment
	K2SegmentLength = 34
	// K3SegmentLength indicates length of K3 segment
	K3SegmentLength = 40
	// K4SegmentLength indicates length of K4 segment
	K4SegmentLength = 30
	// L1SegmentLength indicates length of L1 segment
	L1SegmentLength = 54
	// N1SegmentLength indicates length of N1 segment
	N1SegmentLength = 146
	// TrailerRecordLength indicates length of trailer record
	TrailerRecordLength = 426
	// PackedSegmentLength indicates length of packed segment
	PackedSegmentLength = 366

	// HeaderRecordCharacterLength indicates description of header record
	HeaderRecordDescription = "Header Record (Character)"
	// BaseSegmentLength indicates description of base segment
	BaseSegmentDescription = "Base Segment (Character)"
	// J2SegmentLength indicates description of J2 segment
	J2SegmentDescription = "J2 Segment"
	// K1SegmentLength indicates description of K1 segment
	K1SegmentDescription = "K1 Segment"
	// K2SegmentLength indicates description of K2 segment
	K2SegmentDescription = "K2 Segment"
	// K3SegmentLength indicates description of K3 segment
	K3SegmentDescription = "K3 Segment"
	// K4SegmentLength indicates description of K4 segment
	K4SegmentDescription = "K4 Segment"
	// L1SegmentLength indicates description of L1 segment
	L1SegmentDescription = "L1 Segment"
	// N1SegmentLength indicates description of N1 segment
	N1SegmentDescription = "N1 Segment"
	// TrailerRecordLength indicates description of trailer record
	TrailerRecordDescription = "Trailer Record"
	// PackedHeaderRecordLength indicates description of packed header record
	PackedHeaderRecordDescription = "Header Record (Packed)"
	// PackedBaseSegmentLength indicates length of description base segment
	PackedBaseSegmentDescription = "Base Segment (Packed)"
	// PackedTrailerRecordLength indicates length of description trailer record
	PackedTrailerRecordDescription = "Trailer Record (Packed)"
)

var (
	//ErrUpperAlpha is given when a field is not numeric characters
	ErrNumeric = errors.New("is not numeric characters")
	//ErrUpperAlpha is given when a field is an invalid phone number
	ErrPhoneNumber = errors.New("is an invalid phone number")
	//ErrValidYear is given when there's an invalid date
	ErrValidDate = errors.New("is an invalid Date")
	//ErrValidYear is given when a segment has an invalid length
	ErrSegmentLength = errors.New("has an invalid length")
	//ErrValidField is given when there's an invalid field
	ErrValidField = errors.New("is an invalid field")
)

// NewBaseSegment returns a new base segment
func NewBaseSegment() Segment {
	return &BaseSegment{}
}

// NewPackedBaseSegment returns a new packed base segment
func NewPackedBaseSegment() Segment {
	return &PackedBaseSegment{}
}

// NewHeaderRecord returns a new header record
func NewHeaderRecord() Segment {
	return &HeaderRecord{}
}

// NewPackedHeaderRecord returns a new packed header record
func NewPackedHeaderRecord() Segment {
	return &PackedHeaderRecord{}
}

// NewTrailerRecord returns a new trailer record
func NewTrailerRecord() Segment {
	return &TrailerRecord{}
}

// NewPackedTrailerRecord returns a new packed trailer record
func NewPackedTrailerRecord() Segment {
	return &PackedTrailerRecord{}
}
