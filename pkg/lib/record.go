// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

// General record interface
type Record interface {
	Name() string
	Parse([]byte) (int, error)
	String() string
	Bytes() []byte
	Validate() error
	Length() int
	BlockSize() int
	AddApplicableSegment(Segment) error
	GetSegments(string) []Segment
}

const (
	// UnpackedSegmentLength indicates length of unpacked segment
	UnpackedRecordLength = 426
	// PackedSegmentLength indicates length of packed segment
	PackedRecordLength = 366

	// HeaderRecordName indicates name of header record
	HeaderRecordName = "header"
	// BaseSegmentName indicates name of base segment
	BaseSegmentName = "base"
	// TrailerRecordName indicates name of trailer record
	TrailerRecordName = "trailer"
	// PackedHeaderRecordName indicates name of packed header record
	PackedHeaderRecordName = "headerPacked"
	// PackedBaseSegmentName indicates length of name base record
	PackedBaseSegmentName = "basePacked"
	// PackedTrailerRecordName indicates length of name trailer record
	PackedTrailerRecordName = "trailerPacked"
	// TrailerIdentifier indicates record identifier of trailer record
	TrailerIdentifier = "TRAILER"
	// HeaderIdentifier indicates record identifier of header record
	HeaderIdentifier = "HEADER"
)

// NewHeaderRecord returns a new header record
func NewHeaderRecord() Record {
	return &HeaderRecord{
		RecordIdentifier: HeaderIdentifier,
	}
}

// NewPackedHeaderRecord returns a new packed header record
func NewPackedHeaderRecord() Record {
	return &PackedHeaderRecord{
		RecordIdentifier: HeaderIdentifier,
	}
}

// NewTrailerRecord returns a new trailer record
func NewTrailerRecord() Record {
	return &TrailerRecord{
		RecordIdentifier: TrailerIdentifier,
	}
}

// NewPackedTrailerRecord returns a new packed trailer record
func NewPackedTrailerRecord() Record {
	return &PackedTrailerRecord{
		RecordIdentifier: TrailerIdentifier,
	}
}

// NewBaseSegment returns a new base segment
func NewBaseSegment() Record {
	return &BaseSegment{}
}

// NewPackedBaseSegment returns a new packed base segment
func NewPackedBaseSegment() Record {
	return &PackedBaseSegment{}
}
