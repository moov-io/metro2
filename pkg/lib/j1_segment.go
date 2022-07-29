// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"reflect"
	"strings"

	"github.com/moov-io/metro2/pkg/utils"
)

var _ Segment = (*J1Segment)(nil)

// J1Segment holds the j1 segment
type J1Segment struct {
	// Contains a constant of J1.
	SegmentIdentifier string `json:"segmentIdentifier"  validate:"required"`

	// Report the last name of the associated consumer. Titles and prefixes should not be reported.
	// If the surname contains multiple names, such as Paternal Name-Maternal Name, hyphenate the surnames.
	// For example, “SMITH-JONES” or “MARTINEZ-REYES” requires the hyphen.
	// If the surname contains separate words, the hyphen is not required.
	// For example, report “VAN DYKE” or “DE LA CRUZ” with a space between each word.
	// Other than the hyphen, do not report special characters in any of the Consumer Name fields.
	// The Generation Code should be reported in Field 6.
	// Notes: Do not report minors. The name fields should not
	//  contain messages, such as “Parent of”, “Baby”,
	//  “Daughter”, “Child”, etc.
	//  Do not report trustee or estate accounts. In cases
	//  where the debt is included in a revocable trust and
	//  the consumer retains contractual responsibility,
	//  report the consumer’s Full Name, Address, Social
	//  Security Number, and Date of Birth within the J1
	//  Segment fields. Do not report the name of the
	//  trust.
	Surname string `json:"surname"  validate:"required"`

	// Report the full first name of the associated consumer. Names should not be abbreviated.
	// Examples: Report first name “JUNIOR” (not “JR”); report “ROBERT” (not “ROBT”).
	// If reporting multiple first names, hyphenate the first names.
	// Note: If a consumer uses only initials for first and
	//  middle names (e.g., A.J.), the first name initial
	//  should be reported in the First Name field
	//  (e.g., A) and the middle initial should be reported
	//  in the Middle Name field (e.g., J).
	FirstName string `json:"firstName"  validate:"required"`

	// Report the middle name or middle initial of the associated consumer, if available.
	// If reporting multiple middle names, hyphenate the middle names.
	MiddleName string `json:"middleName,omitempty"`

	// Used to distinguish Junior, Senior, II, III, IV, etc.
	// If not applicable, blank fill.
	// Values available:
	//  J = Junior 3 = III 6 = VI 9 = IX
	//  S = Senior 4 = IV 7 = VII
	//  2 = II 5 = V 8 = VIII
	GenerationCode string `json:"generationCode,omitempty"`

	// Report the Social Security Number (SSN) of the associated consumer. Report only valid U.S.-issued SSNs.
	// Reporting of this information is required as the Social Security Number greatly enhances accuracy in matching to the correct consumer.
	// If the consumer does not have a SSN or one is not available for reporting, zero- or 9-fill all positions.
	// Notes:
	//  If the Social Security Number is not reported, the Date of Birth is required to be reported.
	//  Do not report Individual Tax Identification Numbers (ITINs) in this field.
	//  ITINs do not prove identity outside the tax system and should not be offered or accepted as identification for non-tax purposes,
	//   per the Social Security Administration.
	//  Do not report Credit Profile Numbers (CPNs) in this field.
	//  The CPN should not be used for credit reporting purposes and does not replace the Social Security Number.
	SocialSecurityNumber int `json:"socialSecurityNumber"  validate:"required"`

	// Report the full Date of Birth of the associated consumer, including the month, day and year.
	// Reporting of this information is required as the Date of Birth greatly enhances accuracy in matching to the correct consumer.
	// Format is MMDDYYYY.
	// Notes:
	//  If the Date of Birth is not reported, the Social Security Number is required to be reported.
	//  When reporting Authorized Users (ECOA Code 3), the full Date of Birth (MMDDYYYY) must be reported for all newly-added
	//  Authorized Users on all pre-existing and newly-opened accounts, even if the Social Security Number is reported.
	//  Do not report accounts of consumers who are too young to enter into a binding contract.
	DateBirth utils.Time `json:"dateBirth"  validate:"required"`

	// Contains the telephone number of the associated consumer (Area Code + 7 digits).
	TelephoneNumber int64 `json:"telephoneNumber"`

	// Defines the relationship of the associated consumer to the account and designates the account as joint, individual, etc.,
	// in compliance with the Equal Credit Opportunity Act.
	// Exhibit 10 provides a list of ECOA Codes, their definitions and usage.
	// For important information:
	// • Guidelines on reporting consumers who are personally liable for business accounts
	// • Usage guidelines on ECOA Codes T (Terminated) and Z (Delete Consumer)
	// Note:
	//   Codes 0 (Undesignated), 4 (Joint) and 6 (On-Behalf-Of) are obsolete as of September 2003 and may no longer be reported.
	ECOACode string `json:"ecoaCode"  validate:"required"`

	// Contains a value that indicates a special condition of the account that applies to the associated consumer.
	// This special condition may be that a bankruptcy was filed, discharged, dismissed or withdrawn; a debt was reaffirmed; or the consumer cannot be located or is now located.
	// The indicator should be reported one time and will remain on file until another Consumer Information Indicator or a Removal value is reported.
	// As an option, the indicator may be reported each month as long as the condition applies.
	// Regardless of the method of reporting, the indicator will be deleted only when another Consumer Information Indicator or a Removal value (Q, S, U) is reported.
	// Exhibit 11 provides a list of Consumer Information Indicators and examples that demonstrate how to report these codes.
	ConsumerInformationIndicator string `json:"consumerInformationIndicator,omitempty"`

	converter
	validator
}

// Name returns name of j1 segment
func (s *J1Segment) Name() string {
	return J1SegmentName
}

// Parse takes the input record string and parses the j1 segment values
func (s *J1Segment) Parse(record []byte) (int, error) {
	if len(record) < J1SegmentLength {
		return 0, utils.NewErrSegmentLength("j1 segment")
	}

	fields := reflect.ValueOf(s).Elem()
	length, err := s.parseRecordValues(fields, j1SegmentFormat, record, &s.validator, "j1 segment")
	if err != nil {
		return length, err
	}

	return J1SegmentLength, nil
}

// String writes the j1 segment struct to a 100 character string.
func (s *J1Segment) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(j1SegmentFormat)
	fields := reflect.ValueOf(s).Elem()
	buf.Grow(J1SegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Bytes return raw byte array
func (s *J1Segment) Bytes() []byte {
	return []byte(s.String())
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *J1Segment) Validate() error {
	return s.validateRecord(s, j1SegmentFormat, "j1 segment")
}

// Length returns size of segment
func (s *J1Segment) Length() int {
	return J1SegmentLength
}

// validation of generation code
func (s *J1Segment) ValidateGenerationCode() error {
	switch s.GenerationCode {
	case GenerationCodeJunior, GenerationCodeSenior, GenerationCode2, GenerationCode3, GenerationCode4,
		GenerationCode5, GenerationCode6, GenerationCode7, GenerationCode8, GenerationCode9:
		return nil
	}
	return utils.NewErrInvalidValueOfField("generation code", "j1 segment")
}

// validation of telephone number
func (s *J1Segment) ValidateTelephoneNumber() error {
	if err := s.isPhoneNumber(s.TelephoneNumber, "j1 segment"); err != nil {
		return err
	}
	return nil
}
