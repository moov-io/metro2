// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"reflect"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/moov-io/metro2/utils"
)

// TrailerRecord holds the trailer record
type TrailerRecord struct {
	// Contains a value equal to the length of the physical record. This value includes the four bytes reserved for this field.
	// If fixed-length records are being reported, the Trailer Record should be the same length as all the data records.
	// The Trailer Record should be padded with blanks to fill the needed number of positions.
	RecordDescriptorWord int `json:"recordDescriptorWord" validate:"required"`

	// Contains a constant of HEADER, which is used to identify this record.
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

// PackedTrailerRecord holds the packed trailer record
type PackedTrailerRecord TrailerRecord

// Description returns description of trailer record
func (s *TrailerRecord) Description() string {
	return TrailerRecordDescription
}

// Parse takes the input record string and parses the trailer record values
func (s *TrailerRecord) Parse(record string) error {
	if utf8.RuneCountInString(record) != TrailerRecordLength {
		return utils.ErrSegmentLength
	}

	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		// skip local variable
		if !unicode.IsUpper([]rune(fieldName)[0]) {
			continue
		}

		field := fields.FieldByName(fieldName)
		spec, ok := trailerRecordCharacterFormat[fieldName]
		if !ok || !field.IsValid() {
			return utils.ErrValidField
		}

		data := record[spec.Start : spec.Start+spec.Length]
		if err := s.isValidType(spec, data); err != nil {
			return err
		}

		value, err := s.parseValue(spec, data)
		if err != nil {
			return err
		}

		// set value
		if value.IsValid() && field.CanSet() {
			switch value.Interface().(type) {
			case int, int64:
				field.SetInt(value.Interface().(int64))
			case string:
				field.SetString(value.Interface().(string))
			case time.Time:
				field.Set(value)
			}
		}
	}

	return nil
}

// String writes the trailer record struct to a 426 character string.
func (s *TrailerRecord) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(trailerRecordCharacterFormat)
	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return ""
	}

	buf.Grow(TrailerRecordLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *TrailerRecord) Validate() error {
	fields := reflect.ValueOf(s).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if !fields.IsValid() {
			return utils.ErrValidField
		}

		if spec, ok := trailerRecordCharacterFormat[fieldName]; ok {
			if spec.Required == required {
				fieldValue := fields.FieldByName(fieldName)
				if fieldValue.IsZero() {
					return utils.ErrFieldRequired
				}
			}
		}

		funcName := s.validateFuncName(fieldName)
		method := reflect.ValueOf(s).MethodByName(funcName)
		if method.IsValid() {
			response := method.Call(nil)
			if len(response) == 0 {
				continue
			}

			err := method.Call(nil)[0]
			if !err.IsNil() {
				return err.Interface().(error)
			}
		}
	}

	return nil
}

// Description returns description of packed trailer record
func (s *PackedTrailerRecord) Description() string {
	return PackedTrailerRecordDescription
}

// Parse takes the input record string and parses the packed trailer record values
func (s *PackedTrailerRecord) Parse(record string) error {
	if utf8.RuneCountInString(record) != PackedSegmentLength {
		return utils.ErrSegmentLength
	}

	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		// skip local variable
		if !unicode.IsUpper([]rune(fieldName)[0]) {
			continue
		}

		field := fields.FieldByName(fieldName)
		spec, ok := trailerRecordPackedFormat[fieldName]
		if !ok || !field.IsValid() {
			return utils.ErrValidField
		}

		data := record[spec.Start : spec.Start+spec.Length]
		if err := s.isValidType(spec, data); err != nil {
			return err
		}

		value, err := s.parseValue(spec, data)
		if err != nil {
			return err
		}

		// set value
		if value.IsValid() && field.CanSet() {
			switch value.Interface().(type) {
			case int, int64:
				field.SetInt(value.Interface().(int64))
			case string:
				field.SetString(value.Interface().(string))
			case time.Time:
				field.Set(value)
			}
		}
	}

	return nil
}

// String writes the trailer record struct to a 426 character string.
func (s *PackedTrailerRecord) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(trailerRecordPackedFormat)
	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return ""
	}

	buf.Grow(PackedSegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *PackedTrailerRecord) Validate() error {
	fields := reflect.ValueOf(s).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if !fields.IsValid() {
			return utils.ErrValidField
		}

		if spec, ok := trailerRecordPackedFormat[fieldName]; ok {
			if spec.Required == required {
				fieldValue := fields.FieldByName(fieldName)
				if fieldValue.IsZero() {
					return utils.ErrFieldRequired
				}
			}
		}

		funcName := s.validateFuncName(fieldName)
		method := reflect.ValueOf(s).MethodByName(funcName)
		if method.IsValid() {
			response := method.Call(nil)
			if len(response) == 0 {
				continue
			}

			err := method.Call(nil)[0]
			if !err.IsNil() {
				return err.Interface().(error)
			}
		}
	}

	return nil
}
