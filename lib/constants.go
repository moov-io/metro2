// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import "regexp"

// External Constants
//

const (
	// type of portfolio, Line of Credit
	PortfolioTypeCredit = "C"
	// type of portfolio, Installment
	PortfolioTypeInstallment = "I"
	// type of portfolio, Mortgage
	PortfolioTypeMortgage = "M"
	// type of portfolio, Open
	PortfolioTypeOpen = "O"
	// type of portfolio, Revolving
	PortfolioTypeRevolving = "R"
	// duration of credit extended, Line of Credit
	TermsDurationCredit = "LOC"
	// duration of credit extended, Open
	TermsDurationOpen = "001"
	// duration of credit extended, Revolving
	TermsDurationRevolving = "REV"
	// frequency for payments due, Deferred (Refer to Note)
	TermsFrequencyDeferred = "D"
	// frequency for payments due, Single Payment Loan
	TermsFrequencyPayment = "P"
	// frequency for payments due, Weekly
	TermsFrequencyWeekly = "W"
	// frequency for payments due, Biweekly
	TermsFrequencyBiweekly = "B"
	// frequency for payments due, Semimonthly
	TermsFrequencySemimonthly = "E"
	// frequency for payments due, Monthly
	TermsFrequencyMonthly = "M"
	// frequency for payments due, Bimonthly
	TermsFrequencyBimonthly = "L"
	// frequency for payments due, Quarterly
	TermsFrequencyQuarterly = "Q"
	// frequency for payments due, Tri-annually
	TermsFrequencyTriAnnually = "T"
	// frequency for payments due, Semiannually
	TermsFrequencySemiannually = "S"
	// frequency for payments due, Annually
	TermsFrequencyAnnually = "Y"
	// code that properly identifies whether the account was current, past due, in collections or charged off
	// Current account (0–29 days past the due date)
	PaymentRatingCurrent = "0"
	// code that properly identifies whether the account was current, past due, in collections or charged off
	// 30-59 days past the due date
	PaymentRatingPast30 = "1"
	// code that properly identifies whether the account was current, past due, in collections or charged off
	// 60-89 days past the due date
	PaymentRatingPast60 = "2"
	// code that properly identifies whether the account was current, past due, in collections or charged off
	// 90-119 days past the due date
	PaymentRatingPast90 = "3"
	// code that properly identifies whether the account was current, past due, in collections or charged off
	// 120-149 days past the due date
	PaymentRatingPast120 = "4"
	// code that properly identifies whether the account was current, past due, in collections or charged off
	// 150-179 days past the due date
	PaymentRatingPast150 = "5"
	// code that properly identifies whether the account was current, past due, in collections or charged off
	// 180 or more days past the due date
	PaymentRatingPast180 = "6"
	// code that properly identifies whether the account was current, past due, in collections or charged off
	// Collection
	PaymentRatingCollection = "G"
	// code that properly identifies whether the account was current, past due, in collections or charged off
	// Charge-off
	PaymentRatingChargeOff = "L"
	// consecutive payment activity, 0 payments past due (current account)
	PaymentHistoryPast0 = '0'
	// consecutive payment activity, 30 - 59 days past due date
	PaymentHistoryPast30 = '1'
	// consecutive payment activity, 60 - 89 days past due date
	PaymentHistoryPast60 = '2'
	// consecutive payment activity, 90 - 119 days past due date
	PaymentHistoryPast90 = '3'
	// consecutive payment activity, 120 - 149 days past due date
	PaymentHistoryPast120 = '4'
	// consecutive payment activity, 150 - 179 days past due date
	PaymentHistoryPast150 = '5'
	// consecutive payment activity, 180 or more days past due date
	PaymentHistoryPast180 = '6'
	// consecutive payment activity, No payment history available prior to this time
	PaymentHistoryNoPayment = 'B'
	// consecutive payment activity, No payment history available this month.
	PaymentHistoryNoPaymentMonth = 'D'
	// consecutive payment activity, Zero balance and current account
	PaymentHistoryZero = 'E'
	// consecutive payment activity, Collection
	PaymentHistoryCollection = 'G'
	// consecutive payment activity, Foreclosure Completed
	PaymentHistoryForeclosureCompleted = 'H'
	// consecutive payment activity, Voluntary Surrender
	PaymentHistoryVoluntarySurrender = 'J'
	// consecutive payment activity, Repossession
	PaymentHistoryRepossession = 'K'
	// consecutive payment activity, Charge-off
	PaymentHistoryChargeOff = 'L'
	//  status code that properly identifies the current condition of the account, "DF"
	AccountStatusDF = "DF"
	//  status code that properly identifies the current condition of the account, "DA"
	AccountStatusDA = "DA"
	//  status code that properly identifies the current condition of the account, "11"
	AccountStatus11 = "11"
	//  status code that properly identifies the current condition of the account, "61"
	AccountStatus61 = "61"
	//  status code that properly identifies the current condition of the account, "62"
	AccountStatus62 = "62"
	//  status code that properly identifies the current condition of the account, "63"
	AccountStatus63 = "63"
	//  status code that properly identifies the current condition of the account, "64"
	AccountStatus64 = "64"
	//  status code that properly identifies the current condition of the account, "71"
	AccountStatus71 = "71"
	//  status code that properly identifies the current condition of the account, "78"
	AccountStatus78 = "78"
	//  status code that properly identifies the current condition of the account, "80"
	AccountStatus80 = "80"
	//  status code that properly identifies the current condition of the account, "82"
	AccountStatus82 = "82"
	//  status code that properly identifies the current condition of the account, "83"
	AccountStatus83 = "83"
	//  status code that properly identifies the current condition of the account, "84"
	AccountStatus84 = "84"
	//  status code that properly identifies the current condition of the account, "93"
	AccountStatus93 = "93"
	//  status code that properly identifies the current condition of the account, "96"
	AccountStatus96 = "96"
	//  status code that properly identifies the current condition of the account, "97"
	AccountStatus97 = "97"
	//  status code that properly identifies the current condition of the account, "05"
	AccountStatus05 = "05"
	//  status code that properly identifies the current condition of the account, "13"
	AccountStatus13 = "13"
	//  status code that properly identifies the current condition of the account, "65"
	AccountStatus65 = "65"
	//  status code that properly identifies the current condition of the account, "88"
	AccountStatus88 = "88"
	//  status code that properly identifies the current condition of the account, "89"
	AccountStatus89 = "89"
	//  status code that properly identifies the current condition of the account, "94"
	AccountStatus94 = "94"
	//  status code that properly identifies the current condition of the account, "95"
	AccountStatus95 = "95"
	// designates the interest type, Fixed
	InterestIndicatorFixed = "F"
	// designates the interest type, Variable/Adjustable
	InterestIndicatorVariable = "V"
	// Consumer Account Number Change ONLY
	ChangeIndicatorAccountNumber = 1
	// Identification Number Change ONLY
	ChangeIndicatorIdentificationNumber = 2
	// Consumer Account Number AND Identification Number Change
	ChangeIndicatorBothNumber = 3
	// Generation Code Junior
	GenerationCodeJunior = "J"
	// Generation Code Senior
	GenerationCodeSenior = "S"
	// Generation Code 2
	GenerationCode2 = "2"
	// Generation Code 3
	GenerationCode3 = "3"
	// Generation Code 4
	GenerationCode4 = "4"
	// Generation Code 5
	GenerationCode5 = "5"
	// Generation Code 6
	GenerationCode6 = "6"
	// Generation Code 7
	GenerationCode7 = "7"
	// Generation Code 8
	GenerationCode8 = "8"
	// Generation Code 9
	GenerationCode9 = "9"
	// Confirmed/Verified address
	AddressIndicatorConfirmed = "C"
	// Known to be address of associated consumer
	AddressIndicatorKnown = "Y"
	// Not confirmed address
	AddressIndicatorNotConfirmed = "N"
	// Military address
	AddressIndicatorMilitary = "M"
	// Secondary address
	AddressIndicatorSecondary = "S"
	// Business address — not consumer's residence
	AddressIndicatorBusiness = "B"
	// Non-deliverable address/Returned mail
	AddressIndicatorNonDeliverable = "U"
	// Data reporter’s default address
	AddressIndicatorData = "D"
	// Bill Payer Service — not consumer’s residence
	AddressIndicatorBill = "P"
	// Residence Code Owns
	ResidenceCodeOwns = "O"
	// Residence Code Rents
	ResidenceCodeRents = "R"
	// Creditor Classification
	CreditorClassificationRetail = 1
	// Creditor Classification Medical/Health Care
	CreditorClassificationMedical = 2
	// Creditor Classification Oil Company
	CreditorClassificationOil = 3
	// Creditor Classification Government
	CreditorClassificationGovernment = 4
	// Creditor Classification Personal Services
	CreditorClassificationPersonal = 5
	// Creditor Classification Insurance
	CreditorClassificationInsurance = 6
	// Creditor Classification Educational
	CreditorClassificationEducational = 7
	// Creditor Classification Banking
	CreditorClassificationBanking = 8
	// Creditor Classification Rental/Leasing
	CreditorClassificationRental = 9
	// Creditor Classification Utilities
	CreditorClassificationUtilities = 10
	// Creditor Classification Cable/Cellular
	CreditorClassificationCable = 11
	// Creditor Classification Financial
	CreditorClassificationFinancial = 12
	// Creditor Classification Credit Union
	CreditorClassificationCredit = 13
	// Creditor Classification Automotive
	CreditorClassificationAutomotive = 14
	// Creditor Classification Check Guarantee
	CreditorClassificationGuarantee = 15
	// Purchased From Name
	PurchasedIndicatorFromName = 1
	// Sold To Name
	PurchasedIndicatorToName = 2
	// Remove Previously Reported K2 Segment Information
	PurchasedIndicatorRemove = 9
	// Agency Identifier not applicable
	AgencyIdentifierNotApplicable = 0
	// Agency Identifier Fannie Mae
	AgencyIdentifierFannieMae = 1
	// Agency Identifier Freddie Mac
	AgencyIdentifierFreddieMac = 2
	// Specialized Payment Indicator Balloon Payment
	SpecializedBalloonPayment = 1
	// Specialized Payment Indicator Deferred Payment
	SpecializedDeferredPayment = 2
	// ECOA Code Z
	ECOACodeZ = "Z"
)

// Internal Constants
//

var (
	upperAlphanumericRegex = regexp.MustCompile(`[^ A-Z0-9!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	alphanumericRegex      = regexp.MustCompile(`[^ \w!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	phoneRegex             = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	numericRegex           = regexp.MustCompile(`[0-9a-fA-F]`)
)

const (
	zeroString          = "0"
	blankString         = " "
	nineString          = "9"
	timestampFormat     = "01022006150405"
	dateFormat          = "01022006"
	nullable            = ""
	required            = "Y"
	applicable          = "A"
	timestampSizeStr    = "14"
	dateSizeStr         = "8"
	packedTimestampSize = 8
	packedDateSize      = 5
	int64size           = 8
)

// field types
const (
	alphanumeric = 1 << iota
	alpha
	numeric
	timestamp
	date
	descriptor
	packedTimestamp
	packedDate
	packedNumber
)

// field type options
const (
	zeroFill = 1 << 14
	omitted  = 1 << 15
)
