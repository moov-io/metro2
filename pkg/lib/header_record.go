// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"reflect"
	"strings"
	"unicode"

	"github.com/moov-io/metro2/pkg/utils"
)

var _ Record = (*HeaderRecord)(nil)
var _ Segment = (*HeaderRecord)(nil)
var _ Record = (*PackedHeaderRecord)(nil)
var _ Segment = (*PackedHeaderRecord)(nil)

// HeaderRecord holds the header record
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
	ActivityDate utils.Time `json:"activityDate" validate:"required"`

	// Contains the date the media was generated.
	// A future date should not be reported.
	// Format is MMDDYYYY.
	DateCreated utils.Time `json:"dateCreated" validate:"required"`

	// Contains the date your reporting format was developed.
	// Format is MMDDYYYY.
	// If the day is not available, use 01.
	ProgramDate utils.Time `json:"programDate"`

	// Contains the last date your reporting format was revised.
	// Format is MMDDYYYY.
	// If the day is not available, use 01.
	ProgramRevisionDate utils.Time `json:"programRevisionDate"`

	// Contains the name of the processing company sending the data; i.e., data furnisher or processor.
	// If multiple Header Records are provided, the Reporter Name on the second and subsequent Headers may be repeated or blank filled.
	ReporterName string `json:"reporterName" validate:"required"`

	// Contains the complete mailing address of the processing company; i.e., street address, city, state and zip code.
	ReporterAddress string `json:"reporterAddress" validate:"required"`

	// Contains the telephone number (Area Code + number) of the company sending the data; i.e., data furnisher or processor.
	ReporterTelephoneNumber int64 `json:"reporterTelephoneNumber"`

	// Contains the name of the software vendor that provided the Metro 2® Format software.
	SoftwareVendorName string `json:"softwareVendorName,omitempty"`

	// Contains the version number of the Metro 2® Format software.
	SoftwareVersionNumber string `json:"softwareVersionNumber,omitempty"`

	// Contains a unique identification number assigned by this consumer reporting agency.
	PRBCProgramIdentifier string `json:"prbcProgramIdentifier,omitempty"`

	converter
	validator
}

// PackedHeaderRecord holds the packed header record
type PackedHeaderRecord HeaderRecord

// Name returns name of header record
func (r *HeaderRecord) Name() string {
	return HeaderRecordName
}

// Parse takes the input record string and parses the header record values
func (r *HeaderRecord) Parse(record []byte) (int, error) {
	if len(record) < UnpackedRecordLength {
		return 0, utils.NewErrSegmentLength("header record")
	}

	fields := reflect.ValueOf(r).Elem()
	length, err := r.parseRecordValues(fields, headerRecordCharacterFormat, record, &r.validator, "header record")
	if err != nil {
		return length, err
	}

	if r.BlockDescriptorWord > 0 {
		return r.BlockDescriptorWord, nil
	}
	return r.RecordDescriptorWord, nil
}

// String writes the header record struct to a 426 character string.
func (r *HeaderRecord) String() string {
	var buf strings.Builder
	specifications := r.toSpecifications(headerRecordCharacterFormat)
	fields := reflect.ValueOf(r).Elem()
	blockSize := r.BlockDescriptorWord
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
func (r *HeaderRecord) Bytes() []byte {
	return []byte(r.String())
}

// Validate performs some checks on the record and returns an error if not Validated
func (r *HeaderRecord) Validate() error {
	return r.validateRecord(r, headerRecordCharacterFormat, "header record")
}

// BlockSize returns size of block
func (r *HeaderRecord) BlockSize() int {
	return r.BlockDescriptorWord
}

// Length returns size of segment
func (r *HeaderRecord) Length() int {
	return r.RecordDescriptorWord
}

// GetSegments returns list of applicable segments by segment name
func (r *HeaderRecord) GetSegments(string) []Segment {
	return nil
}

// AddApplicableSegment will add new applicable segment into record
func (r *HeaderRecord) AddApplicableSegment(s Segment) error {
	return utils.NewErrApplicableSegment("header record", s.Name())
}

// Name returns name of packed header record
func (r *PackedHeaderRecord) Name() string {
	return PackedHeaderRecordName
}

// Parse takes the input record string and parses the packed header record values
func (r *PackedHeaderRecord) Parse(record []byte) (int, error) {
	if len(record) < PackedRecordLength {
		return 0, utils.NewErrSegmentLength("packed header record")
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
		spec, ok := headerRecordPackedFormat[fieldName]
		if !ok || !field.IsValid() {
			return 0, utils.NewErrInvalidValueOfField(fieldName, "packed header record")
		}

		if len(record) < spec.Start+spec.Length+offset {
			return 0, utils.NewErrSegmentLength("packed header record")
		}
		data := string(record[spec.Start+offset : spec.Start+spec.Length+offset])
		if err := r.isValidType(spec, data, fieldName, "packed header record"); err != nil {
			return 0, err
		}

		value, err := r.parseValue(spec, data, fieldName, "packed header record")
		if err != nil {
			return 0, err
		}
		// set value
		if value.IsValid() && field.CanSet() {
			switch value.Interface().(type) {
			case int, int64:
				if fieldName == "BlockDescriptorWord" {
					if !utils.IsVariableLength(record) {
						return 0, utils.NewErrBlockDescriptorWord()
					}
					offset += 4
				}
				field.SetInt(value.Interface().(int64))
			case string:
				field.SetString(value.Interface().(string))
			case utils.Time:
				field.Set(value)
			}
		}
	}

	if r.BlockDescriptorWord > 0 {
		return r.BlockDescriptorWord, nil
	}
	return r.RecordDescriptorWord, nil
}

// String writes the packed header record struct to a 426 character string.
func (r *PackedHeaderRecord) String() string {
	var buf strings.Builder
	specifications := r.toSpecifications(headerRecordPackedFormat)
	fields := reflect.ValueOf(r).Elem()
	blockSize := r.BlockDescriptorWord
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
func (r *PackedHeaderRecord) Bytes() []byte {
	return []byte(r.String())
}

// Validate performs some checks on the record and returns an error if not Validated
func (r *PackedHeaderRecord) Validate() error {
	fields := reflect.ValueOf(r).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if spec, ok := headerRecordPackedFormat[fieldName]; ok {
			if spec.Required == required {
				fieldValue := fields.FieldByName(fieldName)
				if fieldValue.IsZero() {
					return utils.NewErrFieldRequired(fieldName, "packed header record")
				}
			}
		}
	}

	return nil
}

// BlockSize returns size of block
func (r *PackedHeaderRecord) BlockSize() int {
	return r.BlockDescriptorWord
}

// Length returns size of segment
func (r *PackedHeaderRecord) Length() int {
	return r.RecordDescriptorWord
}

// GetSegments returns list of applicable segments by segment name
func (r *PackedHeaderRecord) GetSegments(string) []Segment {
	return nil
}

// AddApplicableSegment will add new applicable segment into record
func (r *PackedHeaderRecord) AddApplicableSegment(s Segment) error {
	return utils.NewErrApplicableSegment("packed header record", s.Name())
}
