// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

type HeaderRecord struct {
	// Contains a value equal to the length of the block of data and must be reported when using the packed format or
	// when reporting variable length records.  This value includes the four bytes reserved for this field.
	// Report the standard IBM variable record length conventions.
	//
	// This field is not required when reporting fixed length, fixed block records.
	BlockDescriptorWord int `json:"blockDescriptorWord,omitempty"`

	// Contains a value equal to the length of the physical record.
	// This value includes the four bytes reserved for this field.
	// The recording technique is Numeric for the 426 format and Binary for the 366 format.
	//
	//  • Numeric: The entire four bytes are used. Example: F0F4F2F6.
	//  • Binary: The hexadecimal value should be in the first two bytes of the field and the last two bytes should contain binary zeros. Example: 016E0000.
	//
	// If fixed-length records are being reported, the Header Record should be the same length as all the data records.
	// The Header Record should be padded with blanks to fill the needed number of positions.
	RecordDescriptorWord int `json:"recordDescriptorWord" validate:"required"`

	// Contains a constant of HEADER, which is used to identify this record.
	RecordIdentifier string `json:"recordIdentifier" validate:"required"`

	// Contains the cycle identifier for the information being reported, if reporting by cycles.
	// If data contains more than	one cycle, report the first cycle identifier found on the data.
	CycleIdentifier string `json:"cycleIdentifier,omitempty"`

	// Contains a unique identification number assigned by this consumer reporting agency.
	InnovisProgramIdentifier string `json:"innovisProgramIdentifier,omitempty"`

	// Contains a unique identification number assigned by this consumer reporting agency.
	EquifaxProgramIdentifier string `json:"equifaxProgramIdentifier,omitempty"`

	// Contains a unique identification number assigned by this consumer reporting agency.
	ExperianProgramIdentifier string `json:"experianProgramIdentifier,omitempty"`

	// Contains a unique identification number assigned by this consumer reporting agency.
	TransUnionProgramIdentifier string `json:"transUnionProgramIdentifier,omitempty"`

	// Signifies date of most recent update to accounts.
	// If accounts are updated on different dates, use most recent.
	// A future date should not be reported.
	// Format is MMDDYYYY.
	ActivityDate int `json:"activityDate" validate:"required"`

	// Contains the date the media was generated.
	// A future date should not be reported.
	// Format is MMDDYYYY.
	DateCreated int `json:"dateCreated" validate:"required"`

	// Contains the date your reporting format was developed.
	// Format is MMDDYYYY.
	// If the day is not available, use 01.
	ProgramDate int `json:"programDate"`

	// Contains the last date your reporting format was revised.
	// Format is MMDDYYYY.
	// If the day is not available, use 01.
	ProgramRevisionDate int `json:"programRevisionDate"`

	// Contains the name of the processing company sending the data; i.e., data furnisher or processor.
	// If multiple Header Records are provided, the Reporter Name on the second and subsequent Headers may be repeated or blank filled.
	ReporterName string `json:"reporterName" validate:"required"`

	// Contains the complete mailing address of the processing company; i.e., street address, city, state and zip code.
	ReporterAddress string `json:"reporterAddress" validate:"required"`

	// Contains the telephone number (Area Code + number) of the company sending the data; i.e., data furnisher or processor.
	ReporterTelephoneNumber int `json:"reporterTelephoneNumber"`

	// Contains the name of the software vendor that provided the Metro 2® Format software.
	SoftwareVendorName string `json:"softwareVendorName,omitempty"`

	// Contains the version number of the Metro 2® Format software.
	SoftwareVersionNumber string `json:"softwareVersionNumber,omitempty"`

	// Contains a unique identification number assigned by this consumer reporting agency.
	PRBCProgramIdentifier string `json:"prbcProgramIdentifier,omitempty"`

	converter
	validator
}

type PackedHeaderRecord HeaderRecord

func (s *HeaderRecord) Description() string {
	return HeaderRecordDescription
}

func (s *HeaderRecord) Parse(record string) error {
	if utf8.RuneCountInString(record) != BaseSegmentLength {
		return ErrSegmentInvalidLength
	}

	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return ErrSegmentParse
	}

	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		// skip local variable
		if !unicode.IsUpper([]rune(fieldName)[0]) {
			continue
		}

		field := fields.FieldByName(fieldName)
		spec, ok := headerRecordCharacterFormat[fieldName]
		if !ok || !field.IsValid() {
			return ErrSegmentInvalidType
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
			}
		}
	}

	return nil
}

func (s *HeaderRecord) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(headerRecordCharacterFormat)
	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return ""
	}

	buf.Grow(BaseSegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

func (s *HeaderRecord) Validate() error {
	fields := reflect.ValueOf(s).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if !fields.IsValid() {
			return ErrSegmentParse
		}

		if spec, ok := headerRecordCharacterFormat[fieldName]; ok {
			if spec.Required == required {
				fieldValue := fields.FieldByName(fieldName)
				if fieldValue.IsZero() {
					return ErrRequired
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

func (s *PackedHeaderRecord) Description() string {
	return PackedHeaderRecordDescription
}

func (s *PackedHeaderRecord) Parse(record string) error {
	if utf8.RuneCountInString(record) != PackedHeaderRecordCharacterLength {
		return ErrSegmentInvalidLength
	}

	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return ErrSegmentParse
	}

	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		// skip local variable
		if !unicode.IsUpper([]rune(fieldName)[0]) {
			continue
		}

		field := fields.FieldByName(fieldName)
		spec, ok := headerRecordPackedFormat[fieldName]
		if !ok || !field.IsValid() {
			return ErrSegmentInvalidType
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
			}
		}
	}

	return nil
}

func (s *PackedHeaderRecord) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(headerRecordPackedFormat)
	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return ""
	}

	buf.Grow(BaseSegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

func (s *PackedHeaderRecord) Validate() error {
	fields := reflect.ValueOf(s).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if !fields.IsValid() {
			return ErrSegmentParse
		}

		if spec, ok := headerRecordPackedFormat[fieldName]; ok {
			if spec.Required == required {
				fieldValue := fields.FieldByName(fieldName)
				if fieldValue.IsZero() {
					return ErrRequired
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
