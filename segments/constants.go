package segments

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
