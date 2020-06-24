// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"reflect"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/moov-io/metro2/utils"
)

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
func (s *K1Segment) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < K1SegmentLength {
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
		spec, ok := k1SegmentFormat[fieldName]
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
			case int, int64:
				field.SetInt(value.Interface().(int64))
			case string:
				field.SetString(value.Interface().(string))
			case time.Time:
				field.Set(value)
			}
		}
	}

	return K1SegmentLength, nil
}

// String writes the k1 segment struct to a 34 character string.
func (s *K1Segment) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(k1SegmentFormat)
	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return ""
	}

	buf.Grow(K1SegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *K1Segment) Validate() error {
	fields := reflect.ValueOf(s).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if !fields.IsValid() {
			return utils.ErrValidField
		}

		if spec, ok := k1SegmentFormat[fieldName]; ok {
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
	return utils.NewErrValidValue("creditor classification")
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
func (s *K2Segment) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < K2SegmentLength {
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
		spec, ok := k2SegmentFormat[fieldName]
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
			case int, int64:
				field.SetInt(value.Interface().(int64))
			case string:
				field.SetString(value.Interface().(string))
			case time.Time:
				field.Set(value)
			}
		}
	}

	return K2SegmentLength, nil
}

// String writes the k2 segment struct to a 34 character string.
func (s *K2Segment) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(k2SegmentFormat)
	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return ""
	}

	buf.Grow(K2SegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *K2Segment) Validate() error {
	fields := reflect.ValueOf(s).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if !fields.IsValid() {
			return utils.ErrValidField
		}

		if spec, ok := k2SegmentFormat[fieldName]; ok {
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
	return utils.NewErrValidValue("purchased indicator")
}

// validation of purchased name
func (s *K2Segment) ValidatePurchasedName() error {
	if s.PurchasedIndicator == PurchasedIndicatorRemove {
		if !validFilledString(s.PurchasedName) {
			return utils.NewErrValidValue("purchased name")
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
func (s *K3Segment) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < K3SegmentLength {
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
		spec, ok := k3SegmentFormat[fieldName]
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
			case int, int64:
				field.SetInt(value.Interface().(int64))
			case string:
				field.SetString(value.Interface().(string))
			case time.Time:
				field.Set(value)
			}
		}
	}

	return K3SegmentLength, nil
}

// String writes the k3 segment struct to a 40 character string.
func (s *K3Segment) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(k3SegmentFormat)
	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return ""
	}

	buf.Grow(K3SegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *K3Segment) Validate() error {
	fields := reflect.ValueOf(s).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if !fields.IsValid() {
			return utils.ErrValidField
		}

		if spec, ok := k3SegmentFormat[fieldName]; ok {
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
	return utils.NewErrValidValue("agency identifier")
}

// validation of account number
func (s *K3Segment) ValidateAccountNumber() error {
	if s.AgencyIdentifier == AgencyIdentifierNotApplicable {
		if !validFilledString(s.AccountNumber) {
			return utils.NewErrValidValue("account number")
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
	DeferredPaymentStartDate time.Time `json:"deferredPaymentStartDate,omitempty"`

	// Report the date the balloon payment is due, if applicable.
	// Format is MMDDYYYY. If the day is not available, use 01.
	BalloonPaymentDueDate time.Time `json:"balloonPaymentDueDate,omitempty"`

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
func (s *K4Segment) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < K4SegmentLength {
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
		spec, ok := k4SegmentFormat[fieldName]
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
			case int, int64:
				field.SetInt(value.Interface().(int64))
			case string:
				field.SetString(value.Interface().(string))
			case time.Time:
				field.Set(value)
			}
		}
	}

	return K4SegmentLength, nil
}

// String writes the k4 segment struct to a 30 character string.
func (s *K4Segment) String() string {
	var buf strings.Builder
	specifications := s.toSpecifications(k4SegmentFormat)
	fields := reflect.ValueOf(s).Elem()
	if !fields.IsValid() {
		return ""
	}

	buf.Grow(K4SegmentLength)
	for _, spec := range specifications {
		value := s.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.String()
}

// Validate performs some checks on the record and returns an error if not Validated
func (s *K4Segment) Validate() error {
	fields := reflect.ValueOf(s).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if !fields.IsValid() {
			return utils.ErrValidField
		}

		if spec, ok := k4SegmentFormat[fieldName]; ok {
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
	return utils.NewErrValidValue("specialized payment indicator")
}
