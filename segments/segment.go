// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

// General segment interface
type Segment interface {
	Description() string
	Parse(string) (int, error)
	String() string
	Validate() error
	BlockSize() int
	Length() int
}

const (
	// UnpackedSegmentLength indicates length of unpacked segment
	UnpackedSegmentLength = 426
	// PackedSegmentLength indicates length of packed segment
	PackedSegmentLength = 366
	// J1SegmentLength indicates length of J1 segment
	J1SegmentLength = 100
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

	// HeaderRecordCharacterLength indicates description of header record
	HeaderRecordDescription = "headerCharacter"
	// BaseSegmentLength indicates description of base segment
	BaseSegmentDescription = "baseCharacter"
	// J1SegmentLength indicates description of J1 segment
	J1SegmentDescription = "j1"
	// J2SegmentLength indicates description of J2 segment
	J2SegmentDescription = "j2"
	// K1SegmentLength indicates description of K1 segment
	K1SegmentDescription = "k1"
	// K2SegmentLength indicates description of K2 segment
	K2SegmentDescription = "k2"
	// K3SegmentLength indicates description of K3 segment
	K3SegmentDescription = "k3"
	// K4SegmentLength indicates description of K4 segment
	K4SegmentDescription = "k4"
	// L1SegmentLength indicates description of L1 segment
	L1SegmentDescription = "l1"
	// N1SegmentLength indicates description of N1 segment
	N1SegmentDescription = "n1"
	// TrailerRecordLength indicates description of trailer record
	TrailerRecordDescription = "trailer"
	// PackedHeaderRecordLength indicates description of packed header record
	PackedHeaderRecordDescription = "headerPacked"
	// PackedBaseSegmentLength indicates length of description base segment
	PackedBaseSegmentDescription = "basePacked"
	// PackedTrailerRecordLength indicates length of description trailer record
	PackedTrailerRecordDescription = "trailerPacked"

	// J1SegmentIdentifier indicates segment identifier of J1 segment
	J1SegmentIdentifier = "J1"
	// J2SegmentIdentifier indicates segment identifier of J2 segment
	J2SegmentIdentifier = "J2"
	// K1SegmentIdentifier indicates segment identifier of K1 segment
	K1SegmentIdentifier = "K1"
	// K2SegmentIdentifier indicates segment identifier of K2 segment
	K2SegmentIdentifier = "K2"
	// K3SegmentIdentifier indicates segment identifier of K3 segment
	K3SegmentIdentifier = "K3"
	// K4SegmentIdentifier indicates segment identifier of K4 segment
	K4SegmentIdentifier = "K4"
	// L1SegmentIdentifier indicates segment identifier of L1 segment
	L1SegmentIdentifier = "L1"
	// N1SegmentIdentifier indicates segment identifier of N1 segment
	N1SegmentIdentifier = "N1"
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

// NewJ1Segment returns a new j1 segment
func NewJ1Segment() Segment {
	return &J1Segment{}
}

// NewJ2Segment returns a new j1 segment
func NewJ2Segment() Segment {
	return &J2Segment{}
}

// NewK1Segment returns a new k1 segment
func NewK1Segment() Segment {
	return &K1Segment{}
}

// NewK2Segment returns a new k2 segment
func NewK2Segment() Segment {
	return &K2Segment{}
}

// NewK3Segment returns a new k3 segment
func NewK3Segment() Segment {
	return &K3Segment{}
}

// NewK4Segment returns a new k4 segment
func NewK4Segment() Segment {
	return &K4Segment{}
}

// NewL1Segment returns a new l1 segment
func NewL1Segment() Segment {
	return &L1Segment{}
}

// NewN1Segment returns a new n1 segment
func NewN1Segment() Segment {
	return &N1Segment{}
}
