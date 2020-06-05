// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import "errors"

type Segment interface {
	Description() string
	Parse(record string) error
	String() string
	Validate() error
}

const (
	HeaderRecordCharacterLength       = 426
	HeaderRecordDescription           = "Header Record (Character)"
	BaseSegmentLength                 = 426
	BaseSegmentDescription            = "Base Segment (Character)"
	J2SegmentLength                   = 200
	J2SegmentDescription              = "J2 Segment"
	K1SegmentLength                   = 34
	K1SegmentDescription              = "K1 Segment"
	K2SegmentLength                   = 34
	K2SegmentDescription              = "K2 Segment"
	K3SegmentLength                   = 40
	K3SegmentDescription              = "K3 Segment"
	K4SegmentLength                   = 30
	K4SegmentDescription              = "K4 Segment"
	L1SegmentLength                   = 54
	L1SegmentDescription              = "L1 Segment"
	N1SegmentLength                   = 146
	N1SegmentDescription              = "N1 Segment"
	TrailerRecordLength               = 426
	TrailerRecordDescription          = "Trailer Record"
	PackedHeaderRecordCharacterLength = 366
	PackedHeaderRecordDescription     = "Header Record (Packed)"
	PackedBaseSegmentLength           = 366
	PackedBaseSegmentDescription      = "Base Segment (Packed)"
	PackedTrailerRecordLength         = 366
	PackedTrailerRecordDescription    = "Trailer Record (Packed)"
)

var (
	ErrAlphanumeric         = errors.New("is not alphanumeric")
	ErrUpperAlpha           = errors.New("is not uppercase A-Z or 0-9")
	ErrNumeric              = errors.New("is not number")
	ErrPhoneNumber          = errors.New("is not phone number")
	ErrSegmentInvalidLength = errors.New("invalid segment length")
	ErrSegmentInvalidType   = errors.New("not support segment type")
	ErrSegmentParseType     = errors.New("don't parse type")
	ErrTimestamp            = errors.New("invalid timestamp")
	ErrDate                 = errors.New("invalid date")
	ErrRequired             = errors.New("required field is empty")
	ErrSegmentParse         = errors.New("don't parse struct")
)

func NewBaseSegment() Segment {
	return &BaseSegment{}
}
