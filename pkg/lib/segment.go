// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

// General segment interface
type Segment interface {
	Name() string
	Parse([]byte) (int, error)
	String() string
	Bytes() []byte
	Validate() error
	Length() int
}

const (
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

	// J1SegmentLength indicates name of J1 segment
	J1SegmentName = "j1"
	// J2SegmentLength indicates name of J2 segment
	J2SegmentName = "j2"
	// K1SegmentLength indicates name of K1 segment
	K1SegmentName = "k1"
	// K2SegmentLength indicates name of K2 segment
	K2SegmentName = "k2"
	// K3SegmentLength indicates name of K3 segment
	K3SegmentName = "k3"
	// K4SegmentLength indicates name of K4 segment
	K4SegmentName = "k4"
	// L1SegmentLength indicates name of L1 segment
	L1SegmentName = "l1"
	// N1SegmentLength indicates name of N1 segment
	N1SegmentName = "n1"

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

// NewJ1Segment returns a new j1 segment
func NewJ1Segment() Segment {
	return &J1Segment{
		SegmentIdentifier: J1SegmentIdentifier,
	}
}

// NewJ2Segment returns a new j1 segment
func NewJ2Segment() Segment {
	return &J2Segment{
		SegmentIdentifier: J2SegmentIdentifier,
	}
}

// NewK1Segment returns a new k1 segment
func NewK1Segment() Segment {
	return &K1Segment{
		SegmentIdentifier: K1SegmentIdentifier,
	}
}

// NewK2Segment returns a new k2 segment
func NewK2Segment() Segment {
	return &K2Segment{
		SegmentIdentifier: K2SegmentIdentifier,
	}
}

// NewK3Segment returns a new k3 segment
func NewK3Segment() Segment {
	return &K3Segment{
		SegmentIdentifier: K3SegmentIdentifier,
	}
}

// NewK4Segment returns a new k4 segment
func NewK4Segment() Segment {
	return &K4Segment{
		SegmentIdentifier: K4SegmentIdentifier,
	}
}

// NewL1Segment returns a new l1 segment
func NewL1Segment() Segment {
	return &L1Segment{
		SegmentIdentifier: L1SegmentIdentifier,
	}
}

// NewN1Segment returns a new n1 segment
func NewN1Segment() Segment {
	return &N1Segment{
		SegmentIdentifier: N1SegmentIdentifier,
	}
}
