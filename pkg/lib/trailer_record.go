// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"github.com/moov-io/metro2/pkg/utils"
	"reflect"
	"strings"
	"unicode"
)

var _ Record = (*TrailerRecord)(nil)
var _ Segment = (*TrailerRecord)(nil)
var _ Record = (*PackedTrailerRecord)(nil)
var _ Segment = (*PackedTrailerRecord)(nil)

// TrailerRecord holds the trailer record
type TrailerRecord struct {

	// Contains a value equal to the length of the physical record. This value includes the four bytes reserved for this field.
	// If fixed-length records are being reported, the Trailer Record should be the same length as all the data records.
	// The Trailer Record should be padded with blanks to fill the needed number of positions.
	RecordDescriptorWord int `json:"recordDescriptorWord" validate:"required"`

	// Contains a constant of TRAILER, which is used to identify this record.
	RecordIdentifier string `json:"recordIdentifier" validate:"required"`

	// Contains the total number of Base Segments being reported.
	TotalBaseRecords int `json:"totalBaseRecords" validate:"required"`

	// Contains the total number of Base Segments with Status Code DF.
	TotalStatusCodeDF int `json:"totalStatusCodeDF"`

	// Contains the total number of J1 Segments being reported. Do not count blank- or 9-filled segments.
	TotalConsumerSegmentsJ1 int `json:"totalConsumerSegmentsJ1,omitempty"`

	// Contains the total number of J2 Segments being reported. Do not count blank- or 9-filled segments.
	TotalConsumerSegmentsJ2 int `json:"totalConsumerSegmentsJ2,omitempty"`

	// Contains the number of blocks on the file, if applicable.
	BlockCount int `json:"blockCount"`

	// Contains the total number of Base Segments with Status Code DA.
	TotalStatusCodeDA int `json:"totalStatusCodeDA"`

	// Contains the total number of Base Segments with Status Code 05.
	TotalStatusCode05 int `json:"totalStatusCode05"`

	// Contains the total number of Base Segments with Status Code 11.
	TotalStatusCode11 int `json:"totalStatusCode11"`

	// Contains the total number of Base Segments with Status Code 13.
	TotalStatusCode13 int `json:"totalStatusCode13"`

	// Contains the total number of Base Segments with Status Code 61.
	TotalStatusCode61 int `json:"totalStatusCode61"`

	// Contains the total number of Base Segments with Status Code 62.
	TotalStatusCode62 int `json:"totalStatusCode62"`

	// Contains the total number of Base Segments with Status Code 63.
	TotalStatusCode63 int `json:"totalStatusCode63"`

	// Contains the total number of Base Segments with Status Code 64.
	TotalStatusCode64 int `json:"totalStatusCode64"`

	// Contains the total number of Base Segments with Status Code 65.
	TotalStatusCode65 int `json:"totalStatusCode65"`

	// Contains the total number of Base Segments with Status Code 71.
	TotalStatusCode71 int `json:"totalStatusCode71"`

	// Contains the total number of Base Segments with Status Code 78.
	TotalStatusCode78 int `json:"totalStatusCode78"`

	// Contains the total number of Base Segments with Status Code 80.
	TotalStatusCode80 int `json:"totalStatusCode80"`

	// Contains the total number of Base Segments with Status Code 82.
	TotalStatusCode82 int `json:"totalStatusCode82"`

	// Contains the total number of Base Segments with Status Code 83.
	TotalStatusCode83 int `json:"totalStatusCode83"`

	// Contains the total number of Base Segments with Status Code 84.
	TotalStatusCode84 int `json:"totalStatusCode84"`

	// Contains the total number of Base Segments with Status Code 88.
	TotalStatusCode88 int `json:"totalStatusCode88"`

	// Contains the total number of Base Segments with Status Code 89.
	TotalStatusCode89 int `json:"totalStatusCode89"`

	// Contains the total number of Base Segments with Status Code 93.
	TotalStatusCode93 int `json:"totalStatusCode93"`

	// Contains the total number of Base Segments with Status Code 94.
	TotalStatusCode94 int `json:"totalStatusCode94"`

	// Contains the total number of Base Segments with Status Code 95.
	TotalStatusCode95 int `json:"totalStatusCode95"`

	// Contains the total number of Base Segments with Status Code 96.
	TotalStatusCode96 int `json:"totalStatusCode96"`

	// Contains the total number of Base Segments with Status Code 97.
	TotalStatusCode97 int `json:"totalStatusCode97"`

	// Contains the total number of records with ECOA Code Z being reported in the Base Segment, in the J1 Segment and in the J2 Segment.
	TotalECOACodeZ int `json:"totalECOACodeZ"`

	// Contains the total number of records with employment being reported in the N1 Segment.
	TotalEmploymentSegments int `json:"totalEmploymentSegments"`

	// Contains the total number of records with Original Creditors being reported in the K1 Segment.
	TotalOriginalCreditorSegments int `json:"totalOriginalCreditorSegments"`

	// Contains the total number of records with Purchased From/Sold To being reported in the K2 Segment.
	TotalPurchasedToSegments int `json:"totalPurchasedToSegments"`

	// Contains the total number of records with Mortgage Information being reported in the K3 Segment.
	TotalMortgageInformationSegments int `json:"totalMortgageInformationSegments"`

	// Contains the total number of records with Specialized Payment Information being reported in the K4 Segment.
	TotalPaymentInformationSegments int `json:"totalPaymentInformationSegments"`

	// Contains the total number of Consumer Account Number and/or Identification Number changes being reported in the L1 Segment.
	TotalChangeSegments int `json:"totalChangeSegments"`

	// Contains the total number of valid Social Security Numbers reported in the Base Segment, in the J1 Segment and in  the J2 Segment.
	// Do not count zero- or 9-filled SSNs.
	TotalSocialNumbersAllSegments int `json:"totalSocialNumbersAllSegments"`

	// Contains the total number of valid Social Security Numbers reported in the Base Segment. Do not count zero- or 9-filled SSNs.
	TotalSocialNumbersBaseSegments int `json:"totalSocialNumbersBaseSegments"`

	// Contains the total number of valid Social Security Numbers reported in the J1 Segment. Do not count zero- or 9-filled SSNs.
	TotalSocialNumbersJ1Segments int `json:"totalSocialNumbersJ1Segments"`

	// Contains the total number of valid Social Security Numbers reported in the J2 Segment. Do not count zero- or 9-filled SSNs.
	TotalSocialNumbersJ2Segments int `json:"totalSocialNumbersJ2Segments"`

	// Contains the total number of valid Dates of Birth reported in the Base Segment, in the J1 Segment and in the J2 Segment.
	// Do not count zero-filled Dates of Birth.
	TotalDatesBirthAllSegments int `json:"totalDatesBirthAllSegments"`

	// Contains the total number of valid Dates of Birth reported in the Base Segment. Do not count zero-filled Dates of Birth.
	TotalDatesBirthBaseSegments int `json:"totalDatesBirthBaseSegments"`

	// Contains the total number of valid Dates of Birth reported in the J1 Segment. Do not count zero-filled Dates of Birth.
	TotalDatesBirthJ1Segments int `json:"totalDatesBirthJ1Segments"`

	// Contains the total number of valid Dates of Birth reported in the J2 Segment. Do not count zero-filled Dates of Birth.
	TotalDatesBirthJ2Segments int `json:"totalDatesBirthJ2Segments"`

	// Contains the total number of valid Telephone Numbers reported in the Base Segment, in the J1 Segment and in the J2 Segment.
	// Do not count zero-filled Telephone Numbers.
	TotalTelephoneNumbersAllSegments int `json:"totalTelephoneNumbersAllSegments"`

	converter
	validator
}

type TrailerInformation TrailerRecord

// Name returns name of trailer record
func (r *TrailerRecord) Name() string {
	return TrailerRecordName
}

// Parse takes the input record string and parses the trailer record values
func (r *TrailerRecord) Parse(record []byte) (int, error) {
	if len(record) < UnpackedRecordLength {
		return 0, utils.NewErrSegmentLength("trailer record")
	}

	fields := reflect.ValueOf(r).Elem()
	length, err := r.parseRecordValues(fields, trailerRecordCharacterFormat, record, &r.validator, "trailer record")
	if err != nil {
		return length, err
	}

	return r.RecordDescriptorWord, nil
}

// String writes the trailer record struct to a 426 character string.
func (r *TrailerRecord) String() string {
	var buf strings.Builder
	specifications := r.toSpecifications(trailerRecordCharacterFormat)
	fields := reflect.ValueOf(r).Elem()
	blockSize := r.RecordDescriptorWord
	if blockSize == 0 {
		blockSize = r.RecordDescriptorWord
	}
	buf.Grow(blockSize)
	for _, spec := range specifications {
		value := r.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}
	if blockSize > buf.Len() {
		buf.WriteString(strings.Repeat(blankString, blockSize-buf.Len()))
	}

	return buf.String()
}

// Bytes return raw byte array
func (r *TrailerRecord) Bytes() []byte {
	return []byte(r.String())
}

// Validate performs some checks on the record and returns an error if not Validated
func (r *TrailerRecord) Validate() error {
	return r.validateRecord(r, trailerRecordCharacterFormat, "trailer record")
}

// BlockSize returns size of block
func (r *TrailerRecord) BlockSize() int {
	return r.RecordDescriptorWord
}

// Length returns size of record
func (r *TrailerRecord) Length() int {
	return r.RecordDescriptorWord
}

// GetSegments returns list of applicable segments by segment name
func (r *TrailerRecord) GetSegments(string) []Segment {
	return nil
}

// AddApplicableSegment will add new applicable segment into record
func (r *TrailerRecord) AddApplicableSegment(s Segment) error {
	return utils.NewErrApplicableSegment("trailer record", s.Name())
}

// PackedTrailerRecord holds the packed trailer record
type PackedTrailerRecord TrailerRecord

// Name returns name of packed trailer record
func (r *PackedTrailerRecord) Name() string {
	return PackedTrailerRecordName
}

// Parse takes the input record string and parses the packed trailer record values
func (r *PackedTrailerRecord) Parse(record []byte) (int, error) {
	if len(record) < PackedRecordLength {
		return 0, utils.NewErrSegmentLength("packed trailer record")
	}

	fields := reflect.ValueOf(r).Elem()
	offset := 0
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		// skip local variable
		if !unicode.IsUpper([]rune(fieldName)[0]) {
			continue
		}
		field := fields.FieldByName(fieldName)
		spec, ok := trailerRecordPackedFormat[fieldName]
		if !ok || !field.IsValid() {
			return 0, utils.NewErrInvalidValueOfField(fieldName, "packed trailer record")
		}
		data := string(record[spec.Start+offset : spec.Start+spec.Length+offset])
		if err := r.isValidType(spec, data, fieldName, "packed trailer record"); err != nil {
			return 0, err
		}
		value, err := r.parseValue(spec, data, fieldName, "packed trailer record")
		if err != nil {
			return 0, err
		}
		// set value
		if value.IsValid() && field.CanSet() {
			switch value.Interface().(type) {
			case int, int64:
				field.SetInt(value.Interface().(int64))
			case string:
				field.SetString(value.Interface().(string))
			}
		}
	}

	return r.RecordDescriptorWord, nil
}

// String writes the trailer record struct to a 426 character string.
func (r *PackedTrailerRecord) String() string {
	var buf strings.Builder
	specifications := r.toSpecifications(trailerRecordPackedFormat)
	fields := reflect.ValueOf(r).Elem()
	blockSize := r.RecordDescriptorWord
	if blockSize == 0 {
		blockSize = r.RecordDescriptorWord
	}
	buf.Grow(blockSize)
	for _, spec := range specifications {
		value := r.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}
	if blockSize > buf.Len() {
		buf.WriteString(strings.Repeat(blankString, blockSize-buf.Len()))
	}

	return buf.String()
}

// Bytes return raw byte array
func (r *PackedTrailerRecord) Bytes() []byte {
	return []byte(r.String())
}

// Validate performs some checks on the record and returns an error if not Validated
func (r *PackedTrailerRecord) Validate() error {
	fields := reflect.ValueOf(r).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if spec, ok := trailerRecordPackedFormat[fieldName]; ok {
			if spec.Required == required {
				fieldValue := fields.FieldByName(fieldName)
				if fieldValue.IsZero() {
					return utils.NewErrFieldRequired(fieldName, "packed trailer record")
				}
			}
		}
	}

	return nil
}

// BlockSize returns size of block
func (r *PackedTrailerRecord) BlockSize() int {
	return r.RecordDescriptorWord
}

// Length returns size of record
func (r *PackedTrailerRecord) Length() int {
	return r.RecordDescriptorWord
}

// GetSegments returns list of applicable segments by segment name
func (r *PackedTrailerRecord) GetSegments(string) []Segment {
	return nil
}

// AddApplicableSegment will add new applicable segment into record
func (r *PackedTrailerRecord) AddApplicableSegment(s Segment) error {
	return utils.NewErrApplicableSegment("packed header record", s.Name())
}
