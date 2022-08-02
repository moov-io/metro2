// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"reflect"
	"strings"

	"github.com/moov-io/metro2/pkg/utils"
)

var _ Segment = (*J2Segment)(nil)

// J2Segment holds the j2 segment
type J2Segment struct {
	// Contains a constant of J2.
	SegmentIdentifier string `json:"segmentIdentifier" validate:"required"`

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
	//  Security Number, and Date of Birth within the J2
	//  Segment fields. Do not report the name of the
	//  trust.
	Surname string `json:"surname" validate:"required"`

	// Report the full first name of the associated consumer. Names should not be abbreviated.
	// Examples: Report first name “JUNIOR” (not “JR”); report “ROBERT” (not “ROBT”).
	// If reporting multiple first names, hyphenate the first names.
	// Note: If a consumer uses only initials for first and
	//  middle names (e.g., A.J.), the first name initial
	//  should be reported in the First Name field
	//  (e.g., A) and the middle initial should be reported
	//  in the Middle Name field (e.g., J).
	FirstName string `json:"firstName" validate:"required"`

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
	SocialSecurityNumber int `json:"socialSecurityNumber" validate:"required"`

	// Report the full Date of Birth of the associated consumer, including the month, day and year.
	// Reporting of this information is required as the Date of Birth greatly enhances accuracy in matching to the correct consumer.
	// Format is MMDDYYYY.
	// Notes:
	//  If the Date of Birth is not reported, the Social Security Number is required to be reported.
	//  When reporting Authorized Users (ECOA Code 3), the full Date of Birth (MMDDYYYY) must be reported for all newly-added
	//  Authorized Users on all pre-existing and newly-opened accounts, even if the Social Security Number is reported.
	//  Do not report accounts of consumers who are too young to enter into a binding contract.
	DateBirth utils.Time `json:"dateBirth" validate:"required"`

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
	ECOACode string `json:"ecoaCode" validate:"required"`

	// Contains a value that indicates a special condition of the account that applies to the associated consumer.
	// This special condition may be that a bankruptcy was filed, discharged, dismissed or withdrawn; a debt was reaffirmed; or the consumer cannot be located or is now located.
	// The indicator should be reported one time and will remain on file until another Consumer Information Indicator or a Removal value is reported.
	// As an option, the indicator may be reported each month as long as the condition applies.
	// Regardless of the method of reporting, the indicator will be deleted only when another Consumer Information Indicator or a Removal value (Q, S, U) is reported.
	// Exhibit 11 provides a list of Consumer Information Indicators and examples that demonstrate how to report these codes.
	ConsumerInformationIndicator string `json:"consumerInformationIndicator,omitempty"`

	// Contains the standard two-character country abbreviation.
	CountryCode string `json:"countryCode"`

	// Contains billing/mailing address for the associated consumer.
	// If the consumer has a U.S. address and a foreign address, report the U.S. address. If the consumer has never used the U.S.
	// address as a billing/mailing address (e.g., a property address), report the foreign address.
	// If the billing/mailing address does not belong specifically to the consumer, such as a financial counseling site or bill paying service, report the consumer’s home address.
	// The First Line of Address usually includes street number, direction, street name, and type of thoroughfare.
	// If the billing/mailing address is a PO Box or Rural Route, include Box or Route followed by the number (e.g., PO Box 100).
	// Do not report both a street address and a PO Box.
	// If the billing/mailing address is a private mailbox (PMB), the street address should be reported in the First Line of Address (e.g., 5678 Main Street).
	// The PMB number should be reported in the Second Line of Address (e.g., PMB 1234).
	// As an alternative, the entire address can be reported in the First Line of Address; for example, 5678 Main Street PMB 1234.
	// Eliminate internal messages such as: “Do not mail”, “Attorney”, “Charge-off”, “Chapter 13”, “Fraud”, “Trustee”, “Estate of”, “Care of”, “M/R” (Mail Returned), etc.
	// Do not enter data furnisher's address in this field.
	FirstLineAddress string `json:"firstLineAddress" validate:"required"`

	// Contains second line of address, if needed, such as apartment or unit number, or private mailbox number (PMB).
	// Eliminate internal messages such as: “Do not mail”, “Attorney”, “Charge-off”, “Chapter 13”, “Fraud”, “Trustee”, “Estate of”, “Care of”, “M/R” (Mail Returned), etc.
	SecondLineAddress string `json:"secondLineAddress,omitempty"`

	// Contains city name for address of associated consumer.
	// Truncate rightmost positions if city name is greater than 20 characters or use standard 13-character U.S. Postal Service city abbreviations.
	City string `json:"city"  validate:"required"`

	// Contains the standard U.S. Postal Service state abbreviation for the address of the associated consumer.
	State string `json:"state"  validate:"required"`

	// Report the Zip Code of the associated consumer’s address.
	// Use entire field if reporting 9-digit zip codes. Otherwise, leftjustify and blank fill.
	ZipCode string `json:"zipCode"  validate:"required"`

	// Contains one of the following values for the address
	// C = Confirmed/Verified address
	// Note:
	//  Value ‘C’ enables reporting a confirmed or verified address after receiving an address discrepancy notification from a consumer reporting agency.
	//  Report ‘C’ one time after the address is confirmed.
	//
	// Y = Known to be address of associated consumer
	// N = Not confirmed address
	// M = Military address
	// S = Secondary address
	// B = Business address — not consumer's residence
	// U = Non-deliverable address/Returned mail
	// D = Data reporter’s default address
	// P = Bill Payer Service — not consumer’s residence
	//
	//If indicator not available or unknown, blank fill.
	AddressIndicator string `json:"addressIndicator"`

	// Contains the one-character residence code of the address reported in fields 13-17. Values available:
	//  O = Owns
	//  R = Rents
	// If not available or unknown, blank fill.
	ResidenceCode string `json:"residenceCode"`

	converter
	validator
}

// Name returns name of j2 segment
func (s *J2Segment) Name() string {
	return J2SegmentName
}

// Parse takes the input record string and parses the j2 segment values
func (s *J2Segment) Parse(record []byte) (int, error) {
	if len(record) < J2SegmentLength {
		return 0, utils.NewErrSegmentLength("j2 segment")
	}

	fields := reflect.ValueOf(s).Elem()
	length, err := s.parseRecordValues(fields, j2SegmentFormat, record, &s.validator, "j2 segment")
	if err != nil {
		return length, err
	}

	return J2SegmentLength, nil
}

// String writes the j2 segment struct to a 200 character string.
func (s *J2Segment) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(j2SegmentFormat)
	fields := reflect.ValueOf(s).Elem()
	buf.Grow(J2SegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Bytes return raw byte array
func (s *J2Segment) Bytes() []byte {
	return []byte(s.String())
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *J2Segment) Validate() error {
	return s.validateRecord(s, j2SegmentFormat, "j2 segment")
}

// Length returns size of segment
func (s *J2Segment) Length() int {
	return J2SegmentLength
}

// validation of generation code
func (s *J2Segment) ValidateGenerationCode() error {
	switch s.GenerationCode {
	case GenerationCodeJunior, GenerationCodeSenior, GenerationCode2, GenerationCode3, GenerationCode4,
		GenerationCode5, GenerationCode6, GenerationCode7, GenerationCode8, GenerationCode9:
		return nil
	}
	return utils.NewErrInvalidValueOfField("generation code", "j1 segment")
}

// validation of telephone number
func (s *J2Segment) ValidateTelephoneNumber() error {
	if err := s.isPhoneNumber(s.TelephoneNumber, "j2 segment"); err != nil {
		return err
	}
	return nil
}

// validation of address indicator
func (s *J2Segment) ValidateAddressIndicator() error {
	switch s.AddressIndicator {
	case AddressIndicatorConfirmed, AddressIndicatorKnown, AddressIndicatorNotConfirmed, AddressIndicatorMilitary,
		AddressIndicatorSecondary, AddressIndicatorBusiness, AddressIndicatorNonDeliverable,
		AddressIndicatorData, AddressIndicatorBill, blankString:
		return nil
	}
	return utils.NewErrInvalidValueOfField("address indicator", "j2 segment")
}

// validation of residence code
func (s *J2Segment) ValidateResidenceCode() error {
	switch s.ResidenceCode {
	case ResidenceCodeOwns, ResidenceCodeRents, blankString:
		return nil
	}
	return utils.NewErrInvalidValueOfField("residence code", "j2 segment")
}
