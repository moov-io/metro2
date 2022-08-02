// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"reflect"
	"strings"

	"github.com/moov-io/metro2/pkg/utils"
)

var _ Segment = (*K1Segment)(nil)
var _ Segment = (*K2Segment)(nil)
var _ Segment = (*K3Segment)(nil)
var _ Segment = (*K4Segment)(nil)

// K1Segment holds the k1 segment
type K1Segment struct {
	// Contains a constant of K1.
	SegmentIdentifier string `json:"segmentIdentifier"  validate:"required"`

	// This field is required and the content is dependent on the type of reporter.
	// Collection Agencies: Report the name of the company/creditor, including any partnering affinity name1, that originally opened the account for the consumer,
	//  even if the account had been turned over to multiple collection agencies.
	// Debt Buyers: Report the name of the company/creditor, including any partnering affinity name1, that originally opened the account for the consumer,
	//  even if the account had been sold multiple times to different debt buyers.
	// Refer to the K2 Segment for “purchased from” information.
	// Companies Reporting Returned Checks: Report the name of the payee; i.e., name of company to which the check was written.
	// Refer to Frequently Asked Question 15 for additional guidelines on reporting returned checks.
	// Student Loan Guarantors/U.S. Department of Education: Report the name of the original student loan lender.
	// U.S. Treasury: Report the name of the government agency that is the original creditor.
	// One of the following three options should be used when reporting a creditor’s name that would reveal sensitive information about the consumer.
	// 1. Report the name of the institution, but do not include reference to the type of service.
	//    For example, use the hospital name without identifying that it was the psychiatric unit that provided care. If a hospital’s name reveals sensitive information, abbreviate the name.
	// 2. Use the corporate name if it is different from the commercial name of a mental institution or drug rehabilitation center.
	// 3. Do not report the account if either of the above two options would not sufficiently protect the consumer’s privacy
	OriginalCreditorName string `json:"originalCreditorName"  validate:"required"`

	// Contains a code which must be reported to indicate the general type of business for the Original Creditor Name.
	// Values available:
	//  01 = Retail
	//  02 = Medical/Health Care Required when reporting medical debts and returned checks from providers of medical services, products or devices
	//  03 = Oil Company
	//  04 = Government
	//  05 = Personal Services
	//  06 = Insurance
	//  07 = Educational
	//  08 = Banking
	//  09 = Rental/Leasing
	//  10 = Utilities
	//  11 = Cable/Cellular
	//  12 = Financial (other non-banking financial institutions)
	//  13 = Credit Union
	//  14 = Automotive
	//  15 = Check Guarantee
	CreditorClassification int `json:"creditorClassification"  validate:"required"`

	converter
	validator
}

// Name returns name of K1 segment
func (s *K1Segment) Name() string {
	return K1SegmentName
}

// Parse takes the input record string and parses the k1 segment values
func (s *K1Segment) Parse(record []byte) (int, error) {
	if len(record) < K1SegmentLength {
		return 0, utils.NewErrSegmentLength("k1 segment")
	}

	fields := reflect.ValueOf(s).Elem()
	length, err := s.parseRecordValues(fields, k1SegmentFormat, record, &s.validator, "k1 segment")
	if err != nil {
		return length, err
	}

	return K1SegmentLength, nil
}

// String writes the k1 segment struct to a 34 character string.
func (s *K1Segment) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(k1SegmentFormat)
	fields := reflect.ValueOf(s).Elem()
	buf.Grow(K1SegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Bytes return raw byte array
func (s *K1Segment) Bytes() []byte {
	return []byte(s.String())
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *K1Segment) Validate() error {
	return s.validateRecord(s, k1SegmentFormat, "k1 segment")
}

// Length returns size of segment
func (s *K1Segment) Length() int {
	return K1SegmentLength
}

// validation of creditor classification
func (s *K1Segment) ValidateCreditorClassification() error {
	switch s.CreditorClassification {
	case CreditorClassificationRetail, CreditorClassificationMedical, CreditorClassificationOil,
		CreditorClassificationGovernment, CreditorClassificationPersonal, CreditorClassificationInsurance,
		CreditorClassificationEducational, CreditorClassificationBanking, CreditorClassificationRental,
		CreditorClassificationUtilities, CreditorClassificationCable, CreditorClassificationFinancial,
		CreditorClassificationCredit, CreditorClassificationAutomotive, CreditorClassificationGuarantee:
		return nil
	}
	return utils.NewErrInvalidValueOfField("creditor classification", "k1 segment")
}

// K2Segment holds the k2 segment
type K2Segment struct {
	// Contains a constant of K2.
	SegmentIdentifier string `json:"segmentIdentifier"  validate:"required"`

	// Contains a code representing the type of information being reported. Values available:
	// 1 = Purchased From Name
	// 2 = Sold To Name
	// 9 = Remove Previously Reported K2 Segment Information
	PurchasedIndicator int `json:"purchasedIndicator"  validate:"required"`

	// Contains the name of the company from which the account was purchased or to which the account was sold.
	// If field 2 = 9, this field should be blank filled.
	PurchasedName string `json:"purchasedName"  validate:"required"`

	converter
	validator
}

// Name returns name of K2 segment
func (s *K2Segment) Name() string {
	return K2SegmentName
}

// Parse takes the input record string and parses the k2 segment values
func (s *K2Segment) Parse(record []byte) (int, error) {
	if len(record) < K2SegmentLength {
		return 0, utils.NewErrSegmentLength("k2 segment")
	}

	fields := reflect.ValueOf(s).Elem()
	length, err := s.parseRecordValues(fields, k2SegmentFormat, record, &s.validator, "k2 segment")
	if err != nil {
		return length, err
	}

	return K2SegmentLength, nil
}

// String writes the k2 segment struct to a 34 character string.
func (s *K2Segment) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(k2SegmentFormat)
	fields := reflect.ValueOf(s).Elem()
	buf.Grow(K2SegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Bytes return raw byte array
func (s *K2Segment) Bytes() []byte {
	return []byte(s.String())
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *K2Segment) Validate() error {
	return s.validateRecord(s, k2SegmentFormat, "k2 segment")
}

// Length returns size of segment
func (s *K2Segment) Length() int {
	return K2SegmentLength
}

// validation of purchased indicator
func (s *K2Segment) ValidatePurchasedIndicator() error {
	switch s.PurchasedIndicator {
	case PurchasedIndicatorFromName, PurchasedIndicatorToName, PurchasedIndicatorRemove:
		return nil
	}
	return utils.NewErrInvalidValueOfField("purchased indicator", "k2 segment")
}

// validation of purchased name
func (s *K2Segment) ValidatePurchasedName() error {
	if s.PurchasedIndicator == PurchasedIndicatorRemove {
		if !validFilledString(s.PurchasedName) {
			return utils.NewErrInvalidValueOfField("purchased name", "k2 segment")
		}
	}
	return nil
}

// K3Segment holds the k3 segment
type K3Segment struct {
	// Contains a constant of K3.
	SegmentIdentifier string `json:"segmentIdentifier"  validate:"required"`

	// Contains a code indicating which secondary marketing agency has interest in this loan. Values available:
	// 00 = Agency Identifier not applicable (Used when reporting MIN only)
	// 01 = Fannie Mae
	// 02 = Freddie Mac
	AgencyIdentifier int `json:"agencyIdentifier,omitempty"`

	// Contains the account number as assigned by the secondary marketing agency. Do not include embedded blanks or special characters.
	// If field 2 = 00, this field should be blank filled.
	AccountNumber string `json:"accountNumber,omitempty"`

	// Contains the Mortgage Identification Number assigned to a mortgage loan. Do not include embedded blanks or special characters.
	// The MIN indicates that the loan is registered with the Mortgage Electronic Registration Systems, Inc.
	// (MERS), the electronic registry for tracking the ownership of mortgage rights.
	// For more information, see http://www.mersinc.org.
	MortgageIdentificationNumber string `json:"mortgageIdentificationNumber"`

	converter
	validator
}

// Name returns name of K3 segment
func (s *K3Segment) Name() string {
	return K3SegmentName
}

// Parse takes the input record string and parses the k3 segment values
func (s *K3Segment) Parse(record []byte) (int, error) {
	if len(record) < K3SegmentLength {
		return 0, utils.NewErrSegmentLength("k3 segment")
	}

	fields := reflect.ValueOf(s).Elem()
	length, err := s.parseRecordValues(fields, k3SegmentFormat, record, &s.validator, "k3 segment")
	if err != nil {
		return length, err
	}

	return K3SegmentLength, nil
}

// Bytes return raw byte array
func (s *K3Segment) Bytes() []byte {
	return []byte(s.String())
}

// String writes the k3 segment struct to a 40 character string.
func (s *K3Segment) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(k3SegmentFormat)
	fields := reflect.ValueOf(s).Elem()
	buf.Grow(K3SegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *K3Segment) Validate() error {
	return s.validateRecord(s, k3SegmentFormat, "k3 segment")
}

// Length returns size of segment
func (s *K3Segment) Length() int {
	return K3SegmentLength
}

// validation of agency identifier
func (s *K3Segment) ValidateAgencyIdentifier() error {
	switch s.AgencyIdentifier {
	case AgencyIdentifierNotApplicable, AgencyIdentifierFannieMae, AgencyIdentifierFreddieMac:
		return nil
	}
	return utils.NewErrInvalidValueOfField("agency identifier", "k3 segment")
}

// validation of account number
func (s *K3Segment) ValidateAccountNumber() error {
	if s.AgencyIdentifier == AgencyIdentifierNotApplicable {
		if !validFilledString(s.AccountNumber) {
			return utils.NewErrInvalidValueOfField("account number", "k3 segment")
		}
	}
	return nil
}

// K4Segment holds the k4 segment
type K4Segment struct {
	// Contains a constant of K4.
	SegmentIdentifier string `json:"segmentIdentifier"  validate:"required"`

	// Contains a code describing the specialized payment arrangements.
	// Values available:
	// 01 = Balloon Payment
	// 02 = Deferred Payment
	SpecializedPaymentIndicator int `json:"specializedPaymentIndicator"  validate:"required"`

	// Report the date the first payment is due for deferred loans.
	// Format is MMDDYYYY. If the day is not available, use 01.
	DeferredPaymentStartDate utils.Time `json:"deferredPaymentStartDate,omitempty"`

	// Report the date the balloon payment is due, if applicable.
	// Format is MMDDYYYY. If the day is not available, use 01.
	BalloonPaymentDueDate utils.Time `json:"balloonPaymentDueDate,omitempty"`

	// Report the amount of the balloon payment in whole dollars only.
	BalloonPaymentAmount int `json:"balloonPaymentAmount"`

	converter
	validator
}

// Name returns name of K4 segment
func (s *K4Segment) Name() string {
	return K4SegmentName
}

// Parse takes the input record string and parses the k4 segment values
func (s *K4Segment) Parse(record []byte) (int, error) {
	if len(record) < K4SegmentLength {
		return 0, utils.NewErrSegmentLength("k4 segment")
	}

	fields := reflect.ValueOf(s).Elem()
	length, err := s.parseRecordValues(fields, k4SegmentFormat, record, &s.validator, "k4 segment")
	if err != nil {
		return length, err
	}

	return K4SegmentLength, nil
}

// Bytes return raw byte array
func (s *K4Segment) Bytes() []byte {
	return []byte(s.String())
}

// String writes the k4 segment struct to a 30 character string.
func (s *K4Segment) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(k4SegmentFormat)
	fields := reflect.ValueOf(s).Elem()
	buf.Grow(K4SegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *K4Segment) Validate() error {
	return s.validateRecord(s, k4SegmentFormat, "k4 segment")
}

// Length returns size of segment
func (s *K4Segment) Length() int {
	return K4SegmentLength
}

// validation of specialized payment indicator
func (s *K4Segment) ValidateSpecializedPaymentIndicator() error {
	switch s.SpecializedPaymentIndicator {
	case SpecializedBalloonPayment, SpecializedDeferredPayment:
		return nil
	}
	return utils.NewErrInvalidValueOfField("specialized payment indicator", "k4 segment")
}
