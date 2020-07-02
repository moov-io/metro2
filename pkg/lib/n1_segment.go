package lib

import (
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/moov-io/metro2/pkg/utils"
)

// N1Segment holds the n1 segment
type N1Segment struct {
	// Contains a constant of N1.
	SegmentIdentifier string `json:"segmentIdentifier"  validate:"required"`

	// Report the name of the employer for the consumer reported in the Base Segment.
	EmployerName string `json:"employerName"  validate:"required"`

	// Contains the mailing address for the employer in Field 2 and usually includes street number, direction, street name and type of thoroughfare.
	FirstLineEmployerAddress string `json:"firstLineEmployerAddress"`

	// Contains second line of employer’s address, if needed.
	SecondLineEmployerAddress string `json:"secondLineEmployerAddress"`

	// Contains city name for employer’s address.
	// Truncate rightmost positions if city name is greater than 20 characters or use standard 13-character U.S.
	// Postal Service city abbreviations.
	EmployerCity string `json:"employerCity"`

	// Contains the standard U.S. Postal Service state abbreviation for the address of the employer.
	EmployerState string `json:"employerState"`

	// Report the zip code of the employer’s address.
	// Use entire field if reporting 9-digit zip codes. Otherwise, left-justify and blank fill.
	ZipCode string `json:"zipCode"`

	// Report title or position for consumer reported in the Base Segment (the employee).
	Occupation string `json:"occupation"`

	converter
	validator
}

// Name returns name of N1 segment
func (s *N1Segment) Name() string {
	return N1SegmentName
}

// Parse takes the input record string and parses the n1 segment values
func (s *N1Segment) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < N1SegmentLength {
		return 0, utils.ErrSegmentLength
	}

	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return 0, utils.ErrValidField
	}

	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		// skip local variable
		if !unicode.IsUpper([]rune(fieldName)[0]) {
			continue
		}

		field := fields.FieldByName(fieldName)
		spec, ok := n1SegmentFormat[fieldName]
		if !ok || !field.IsValid() {
			return 0, utils.ErrValidField
		}

		data := record[spec.Start : spec.Start+spec.Length]
		if err := s.isValidType(spec, data); err != nil {
			return 0, err
		}

		value, err := s.parseValue(spec, data)
		if err != nil {
			return 0, err
		}

		// set value
		if value.IsValid() && field.CanSet() {
			switch value.Interface().(type) {
			case string:
				field.SetString(value.Interface().(string))
			}
		}
	}

	return N1SegmentLength, nil
}

// String writes the n1 segment struct to a 146 character string.
func (s *N1Segment) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(n1SegmentFormat)
	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return ""
	}

	buf.Grow(N1SegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *N1Segment) Validate() error {
	fields := reflect.ValueOf(s).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if !fields.IsValid() {
			return utils.ErrValidField
		}

		if spec, ok := n1SegmentFormat[fieldName]; ok {
			if spec.Required == required {
				fieldValue := fields.FieldByName(fieldName)
				if fieldValue.IsZero() {
					return utils.ErrFieldRequired
				}
			}
		}
	}

	return nil
}

// Length returns size of segment
func (s *N1Segment) Length() int {
	return N1SegmentLength
}
