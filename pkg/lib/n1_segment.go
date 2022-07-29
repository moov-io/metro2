package lib

import (
	"github.com/moov-io/metro2/pkg/utils"
	"reflect"
	"strings"
)

var _ Segment = (*N1Segment)(nil)

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
func (s *N1Segment) Parse(record []byte) (int, error) {
	if len(record) < N1SegmentLength {
		return 0, utils.NewErrSegmentLength("n1 segment")
	}

	fields := reflect.ValueOf(s).Elem()
	length, err := s.parseRecordValues(fields, n1SegmentFormat, record, &s.validator, "n1 segment")
	if err != nil {
		return length, err
	}

	return N1SegmentLength, nil
}

// String writes the n1 segment struct to a 146 character string.
func (s *N1Segment) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(n1SegmentFormat)
	fields := reflect.ValueOf(s).Elem()
	buf.Grow(N1SegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Bytes return raw byte array
func (s *N1Segment) Bytes() []byte {
	return []byte(s.String())
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *N1Segment) Validate() error {
	return s.validateRecord(s, n1SegmentFormat, "n1 segment")
}

// Length returns size of segment
func (s *N1Segment) Length() int {
	return N1SegmentLength
}
