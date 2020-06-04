// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import "errors"

type Segment interface {
	Parse(record string) error
	String() string
	Validate() error
}

const (
	BaseSegmentLength = 426
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
)
