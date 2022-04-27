// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

var (
	// ErrInvalidMetroFile is given when there's an invalid file
	ErrInvalidMetroFile = errors.New("the file isn't metro file")
	// ErrNonHeaderRecord is given when there isn't header record
	ErrNonHeaderRecord = errors.New("the file must have header record")
)

// NewErrApplicableSegment returns a error that can't support applicable segment
func NewErrApplicableSegment(recordName, segName string) error {
	return fmt.Errorf("%s has not applicable segment (%s)", recordName, segName)
}

// NewErrInvalidSegment returns a error that has invalid segment
func NewErrInvalidSegment(name string) error {
	return fmt.Errorf("has an invalid segment (%s)", name)
}

// NewErrFailedParsing returns a error that is failed parsing raw data
func NewErrFailedParsing() error {
	return fmt.Errorf("is failed parsing raw data")
}

// NewErrBlockDescriptorWord returns a error that
func NewErrBlockDescriptorWord() error {
	return fmt.Errorf("should be block descriptor word")
}

// NewErrInvalidRecord returns a error if is invalid record
func NewErrInvalidRecord(name string) error {
	return fmt.Errorf("is an invalid record (%s)", name)
}

// NewErrInvalidValueOfField returns a error that has invalid value
func NewErrInvalidValueOfField(fieldName, recordName string) error {
	return fmt.Errorf("%s in %s has an invalid value", toSnakeCase(fieldName), recordName)
}

// NewErrNonAlphanumeric is given when a field has non-alphanumeric characters
func NewErrNonAlphanumeric(fieldName, recordName string) error {
	return fmt.Errorf("%s in %s has not alphanumeric characters", toSnakeCase(fieldName), recordName)
}

// NewErrUpperAlpha is given when a field is not in uppercase
func NewErrUpperAlpha(fieldName, recordName string) error {
	return fmt.Errorf("%s in %s has not uppercase A-Z or 0-9", toSnakeCase(fieldName), recordName)
}

// NewErrNumeric is given when a field is not numeric characters
func NewErrNumeric(fieldName, recordName string) error {
	return fmt.Errorf("%s in %s has not numeric characters", toSnakeCase(fieldName), recordName)
}

// NewErrValidDate is given when there's an invalid date
func NewErrValidDate(fieldName, recordName string) error {
	return fmt.Errorf("%s in %s has an invalid date", toSnakeCase(fieldName), recordName)
}

// NewErrFieldRequired is given when a field is required
func NewErrFieldRequired(fieldName, recordName string) error {
	return fmt.Errorf("%s in %s is a required field", toSnakeCase(fieldName), recordName)
}

// NewErrPhoneNumber is given when a field is an invalid phone number
func NewErrPhoneNumber(recordName string) error {
	return fmt.Errorf("telephone number in %s has an invalid phone number", recordName)
}

// NewErrSegmentLength is given when a segment has an invalid length
func NewErrSegmentLength(recordName string) error {
	return fmt.Errorf("%s has an invalid length", recordName)
}
