// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/moov-io/metro2/pkg/utils"
)

var _ Record = (*BaseSegment)(nil)
var _ Segment = (*BaseSegment)(nil)
var _ Record = (*PackedBaseSegment)(nil)
var _ Segment = (*PackedBaseSegment)(nil)

// BaseSegment holds the base segment
type BaseSegment struct {
	// Contains a value equal to the length of the block of data and must be reported when using the packed format or
	// when reporting variable length records.  This value includes the four bytes reserved for this field.
	// Report the standard IBM variable record length conventions.
	//
	// This field is not required when reporting fixed length, fixed block records.
	BlockDescriptorWord int `json:"blockDescriptorWord,omitempty"`

	// Contains a value equal to the length of the physical record.  This value includes the four bytes reserved
	// for this field.
	// The length of each segment should be included in the RDW.
	//
	// For example:
	//  Base Segment =   426
	//  J2 Segment   =   200
	//  K1 Segment   =    34
	//  RDW          =  0660
	//
	// For fixed block, the RDW will remain the same for each record.
	// For variable block, the RDW will change depending on the size of each record.
	RecordDescriptorWord int `json:"recordDescriptorWord" validate:"required"`

	// Report a constant of 1.
	ProcessingIndicator int `json:"processingIndicator"`

	// Contains date and time of actual account information update.
	// Format for packed date is 0MMDDYYYYHHMMSSs — where s is the sign.
	// Format is MMDDYYYYHHMMSS for character date.
	TimeStamp utils.Time `json:"timeStamp"`

	// Used to uniquely identify a data furnisher.
	// Report your internal code to identify each branch, office, and/or credit central where information is verified.
	// For accounts reported by servicers, the Identification Number should refer to the current holder of the note.
	//
	// This number must be unique, at least 5 digits long, and should not include embedded blanks or special characters.
	// Entire field should never be zero, blank or 9 filled.
	//
	// This field must be consistent on a month-to-month basis to avoid duplication of information.
	// Notify consumer reporting agencies before adding, deleting, or changing the identifiers in this field.
	IdentificationNumber string `json:"identificationNumber" validate:"required"`

	// Report the internal cycle code for this account.
	// Field is required if reporting by cycles; otherwise blank fill.
	CycleIdentifier string `json:"cycleIdentifier,omitempty"`

	// Report the individual's complete and unique account number as extracted from your file.
	// Do not include embedded blanks or special characters.
	//
	// Do not report the Social Security Number, in whole or in part, within the Consumer Account Number.
	//
	// Account number scrambling and encryption methods for security purposes are permitted.
	// Contact the consumer reporting agencies for information regarding the Metro 2® scrambling techniques.
	ConsumerAccountNumber string `json:"consumerAccountNumber" validate:"required"`

	// Contains the one-character abbreviation for type of portfolio.  Values available:
	//
	//  C = Line of Credit
	//  I = Installment
	//  M = Mortgage
	//  O = Open
	//  R = Revolving
	//
	// Refer to the Glossary of Terms for definitions of each Portfolio Type.
	PortfolioType string `json:"portfolioType"`

	// Report the specific code that identifies the account classification.
	// Exhibit 1 provides a numeric listing of type codes that specify industry usage, and Exhibit 2 provides an alphabetic listing of type codes within their corresponding Portfolio Types.
	AccountType string `json:"accountType" validate:"required"`

	// Report the date the account was originally opened.
	// Retain the original Date Opened regardless of future activity, such as transfer, refinance, lost or stolen card, etc.
	//
	// Valid Dates Opened must be reported – field cannot be zero or blank filled, nor contain a date in the future.
	//
	// For companies who report returned checks, such as collection agencies, report the date of the check.
	//
	// Format for character date is MMDDYYYY.  Format for packed date is 0MMDDYYYYs — where s is the sign.
	// If the day is not available, use 01.
	DateOpened utils.Time `json:"dateOpened"`

	// Report the following values in whole dollars only:
	//
	//  Line of Credit = Assigned credit limit*
	//  Installment = Zero fill
	//  Mortgage = Zero fill
	//  Open = Assigned credit limit*, if applicable; otherwise, zero fill
	//  Revolving = Assigned credit limit*
	//
	// * For closed accounts, continue to report the last assigned   credit limit.
	CreditLimit int `json:"creditLimit,omitempty"`

	// Report the following values in whole dollars only:
	//
	//  Line of Credit = Highest amount of credit utilized by the consumer
	//  Installment = Original amount of the loan excluding interest payments
	//  Mortgage = Original amount of the loan excluding interest payments
	//  Open = Highest amount of credit utilized by the consumer, if applicable
	//  Revolving = Highest amount of credit utilized by the consumer
	//
	// For companies who report returned checks, such as collection agencies, report the original amount of the check, excluding fees and interest.
	HighestCredit int `json:"highestCredit" validate:"required"`

	// Contains the duration of credit extended.
	//
	//  Line of Credit = Constant of LOC
	//  Installment = Number of months
	//  Mortgage = Number of years
	//  Open = Constant of 001, One payment as scheduled
	//  Revolving = Constant of REV
	//
	// Exhibit 3 provides the calculations necessary to convert Terms Duration to monthly.
	TermsDuration string `json:"termsDuration" validate:"required"`

	// Report the frequency for payments due.  Values available:
	//
	//  D  = Deferred (Refer to Note)
	//  P  = Single Payment Loan
	//  W  = Weekly
	//  B  = Biweekly
	//  E  = Semimonthly
	//  M   = Monthly
	//  L   = Bimonthly
	//  Q  = Quarterly
	//  T  = Tri-annually
	//  S  = Semiannually
	//  Y  = Annually
	//
	// Exhibit 3 provides definitions of the Terms Frequency Codes.
	// Note: When reporting Deferred loans, report the Deferred  Payment Start Date in the K4 Segment.
	TermsFrequency string `json:"termsFrequency,omitempty"`

	// Report the dollar amount of the scheduled monthly payment due for this reporting period, whether principal, interest only or a combination of the two.
	// When a balloon payment is also due during the reporting period, the balloon payment amount should be included to represent the entire monthly payment amount due.
	//
	// Report in whole dollars only.
	// When the account is paid in full, the Scheduled Monthly Payment Amount should be zero filled.
	//
	//  Line of Credit = Minimum amount due based on balance, not including any amounts past due
	//  Installment = Regular monthly payment
	//  Mortgage = Regular monthly payment, including the principal, interest, and escrow due this month
	//  Open = Zero fill
	//  Revolving = Minimum amount due based on balance, not including any amounts past due
	//
	// Exhibit 3 provides the calculations necessary to convert payment amounts to monthly.
	ScheduledMonthlyPaymentAmount int `json:"scheduledMonthlyPaymentAmount,omitempty"`

	// Report the dollar amount of the monthly payment actually received for this reporting period in whole dollars only.
	// If multiple payments are made during the reporting period, the total amount should be reported.
	ActualPaymentAmount int `json:"actualPaymentAmount,omitempty"`

	// Contains the status code that properly identifies the current condition of the account as of the Date of Account Information
	// Exhibit 4 provides a description of these codes.
	// The Payment Rating (Field 17B) must also be reported when the Account Status Code is 05, 13, 65, 88, 89, 94, or 95.
	// Special Comments (Field 19) may be used in conjunction with the Account Status to further define the account.
	// For examples of how Account Statuses, Payment Ratings and Special Comments interact.
	AccountStatus string `json:"accountStatus" validate:"required"`

	// When the Account Status (Field 17A) contains 05, 13, 65, 88, 89, 94 or 95, this field must also be reported.  The Payment Rating must be blank filled for all other Account Status Codes.
	// The Payment Rating contains a code that properly identifies whether the account was current, past due, in collections or  charged off prior to the status and within the current month’s reporting period.
	//
	// Values available:
	//  0 = Current account (0–29 days past the due date)
	//  1 = 30-59 days past the due date
	//  2 = 60-89 days past the due date
	//  3 = 90-119 days past the due date
	//  4 = 120-149 days past the due date
	//  5 = 150-179 days past the due date
	//  6 = 180 or more days past the due date
	//  G = Collection
	//  L = Charge-off
	//
	// For example, if the account was paid on March 22, 2019, but the consumer was 30 days past the due date on March 10, 2019 prior to paying the account, report Account Status Code = 13 and Payment Rating = 1.
	PaymentRating string `json:"paymentRating,omitempty"`

	// Contains up to 24 months of consecutive payment activity for the previous 24 reporting periods prior to the Date of Account Information (Field 24) being reported.
	// Report one month’s payment history in each byte from the left to right in most recent to least recent order.
	// The first byte should represent the Account Status Code reported in the previous reporting period.  Refer to Exhibit 5 for examples of reporting payment history, which includes examples for month-end reporters, as well as examples for reporters who submit data on other days of the month (e.g., 1st, 15th, etc.).  Values available:
	//
	//  0 = 0 payments past due (current account)
	//  1 = 30 - 59 days past due date
	//  2 =  60 - 89 days past due date
	//  3 =  90 - 119 days past due date
	//  4 =  120 - 149 days past due date
	//  5 =  150 - 179 days past due date
	//  6 =  180 or more days past due date
	//  B =  No payment history available prior to this time – either because the account was not open or because the payment history cannot be furnished.  A “B” may not be embedded within other values.
	//  D =  No payment history available this month.   “D” may be embedded in the payment pattern.
	//  E =  Zero balance and current account           (Applies to Credit Cards and Lines of Credit)
	//  G =  Collection
	//  H =  Foreclosure Completed
	//  J =  Voluntary Surrender
	//  K = Repossession
	//  L =  Charge-off
	//
	// No other values are acceptable in this field.
	// If a full 24 months of history are not available for reporting, the ending positions of this field should be B-filled.
	// The Payment History Profile is intended to be used to report monthly history, regardless of the Terms Frequency.
	// Reporting of the Payment History Profile provides a method for automated correction of erroneously reported history.
	//
	// For important information:
	// • Paid accounts
	// • First-time reporters
	PaymentHistoryProfile string `json:"paymentHistoryProfile" validate:"required"`

	// Used in conjunction with Account Status and Payment Rating (to further define the account (e.g., closed accounts or adjustments pending).
	// The Special Comment Code must be reported each month as long as the condition applies.
	// If more than one Special Comment applies to an account, it is the data furnisher’s decision to report the comment that is deemed most important from a business perspective for the current reporting period.
	//
	// If no Special Comment is applicable, blank fill.
	//
	// Exhibit 6 provides a list of available comments by category within Portfolio Type and Exhibit 7 provides a list of codes in alphabetical sequence.
	// Both exhibits include definitions and usage guidelines.
	//
	// For examples of how Account Statuses, Payment Ratings and Special Comments interact.
	SpecialComment string `json:"specialComment,omitempty"`

	// Allows the reporting of a condition that is required for legal compliance.
	// This condition may refer to accounts closed at consumer’s request, accounts in dispute under the Fair Credit Reporting Act (FCRA), the Fair Debt Collection Practices Act (FDCPA) or the Fair Credit Billing Act (FCBA).
	//
	// The code should be reported one time and will remain on file until another Compliance Condition Code or the XR (Removal code) is reported.  As an option, the code may be reported each month as long as the condition applies.
	// Regardless of the method of reporting, the code will be deleted only when another Compliance Condition Code or the XR (Removal code) is reported.
	//
	// Exhibit 8 provides a list of Compliance Condition Codes and examples that demonstrate how to report these codes.
	// For questions about the use of Compliance Condition Codes or how long to report them, data furnishers should refer to their internal policies and procedures.
	ComplianceConditionCode string `json:"complianceConditionCode,omitempty"`

	// Report the outstanding current balance on the account as of the Date of Account Information.
	//
	// The Current Balance should contain the principal balance including Balloon Payment Amounts (when applicable), as well as applicable interest, late charges, fees, insurance payments and escrow that are due during the current reporting period.
	// The Current Balance may exceed the Highest Credit, Original Loan Amount or Credit Limit.
	//
	// The Current Balance should not include future interest, escrow, fees or insurance payments.
	//
	// This amount, which should be reported in whole dollars only, may increase or decline from month to month.
	// Credit balances (negative balances) should be reported as zero.
	CurrentBalance int `json:"currentBalance" validate:"omitempty"`

	// Report the total amount of payments that are 30 days or more past due in whole dollars only.
	// This field should include late charges and fees, if applicable.
	// Do not include current amount due in this field.
	//
	// Note: If the Account Status is current (Status Code 11), this field should be zero.
	AmountPastDue int `json:"amountPastDue,omitempty"`

	// For Status Codes 64 and 97 (all portfolio types), report the original amount charged to loss, regardless of the declining balance.
	// Report whole dollars only.
	//
	// If payments are received from the consumer, report the outstanding balance in the Current Balance and Amount Past Due fields.
	OriginalChargeOffAmount int `json:"originalChargeOffAmount,omitempty"`

	// All account information in the Base Segment, such as Account Status and Current Balance, must be reported as of the date in this field.
	//
	// For Account Status Codes 11, 71, 78, 80, 82-84, 88, 89, 93-97, DA and DF, report a date within the current month’s reporting period, as noted below:
	//
	// • Cycle Reporters – Report the date of the current month’s billing cycle.  This method is preferred to facilitate accurate and timely reporting of account information.
	// • Monthly Reporters – Report the date within the current month’s reporting period that represents the most recent update to the account, such as mid-month (03/15/2019) or end of month (03/31/2019).  The Date of Account Information may represent the consumer’s billing date as long as the date is within the current month’s reporting period.
	// 						 A historic or future date should not be reported.
	//
	// For Account Status Codes 13 and 61–65, report the date paid, unless the account was closed due to inactivity; then report the date within the current reporting period when the account was closed to further charges.
	// For accounts reported with bankruptcy Consumer Information Indicators, refer to Frequently Asked Questions 27 and 28 for guidelines on reporting the Date of Account Information.
	//
	// Format for character date is MMDDYYYY.  Format for packed date is 0MMDDYYYYs – where s is the sign.
	// Notes: This date must not reflect a future date.
	// For guidelines on reporting paid, closed or inactive accounts, refer to FAQs 39, 40 and 41.
	DateAccountInformation utils.Time `json:"dateAccountInformation" validate:"required"`

	// This date is used to ensure compliance with the Fair Credit Reporting Act.
	// The date in the Date of First Delinquency field must be determined each reporting period based on the following hierarchy:
	//
	// 1. For Account Status Codes 61-65, 71, 78, 80, 82-84, 88-89 and 93-97, report the date of the first 30-day delinquency that led to the status being reported.  This date should be 30 days after the Due Date.  If a delinquent account becomes current, the Date of First Delinquency should be zero filled.  Then if the account goes delinquent again, the Date of First Delinquency starts over with the new first delinquency date.
	// 2. For Account Status Codes 05 and 13, if the Payment Rating is 1, 2, 3, 4, 5, 6, G or L, report the date of the first 30-day delinquency that led to the Payment Rating being reported.  This date should be 30 days after the Due Date.
	// 3. For Consumer Information Indicators A-H and Z (Bankruptcies), 1A (Personal Receivership) and V-Y (Reaffirmation of Debt Rescinded with Bankruptcy Chapters), if the account is current (Account Status Code 11 or Account Status Code 05 or 13 with Payment Rating 0), report the date of the bankruptcy/personal receivership petition or notification.  Even though the account is not delinquent, this date is required for purging purposes.
	//    Notes: In hierarchy rule #3, Account Status 13 is included for scenarios when merchandise is redeemed.  Refer to FAQ 31 for further guidance.
	//           Consumer Information Indicators W, X & Y are obsolete as of September 2010 and may no longer be reported.
	//
	// If none of the conditions listed in the above hierarchy apply, the Date of First Delinquency should be zero filled.
	// The Date of First Delinquency is used by the consumer reporting agencies for purging purposes.  Format for character date is MMDDYYYY.  Format for packed date is 0MMDDYYYYs — where s is the sign.
	// Notes:
	// • Refer to Exhibit 9 for detailed reporting instructions, examples and excerpts from the Fair Credit Reporting Act.
	// • First-time reporters should refer to Frequently Asked Question 22 for important information.
	DateFirstDelinquency utils.Time `json:"dateFirstDelinquency,omitempty"`

	// For all portfolio types, contains the date the account was closed to further purchases, paid in full, transferred or sold.  For Line of Credit, Open or Revolving accounts, there may be a balance due.
	//
	// Format for character date is MMDDYYYY.
	// Format for packed date is 0MMDDYYYYs — where s is the sign.
	// If not applicable, zero fill.
	DateClosed utils.Time `json:"dateClosed,omitempty"`

	// Report the date the most recent payment was received, whether full or partial payment is made.
	//
	// Format for character date is MMDDYYYY.
	// Format for packed date is 0MMDDYYYYs — where s is the sign.
	// If the day is not available, use 01.
	DateLastPayment utils.Time `json:"dateLastPayment,omitempty"`

	// Contains one of the following values that designates the interest type:
	//
	//  F = Fixed
	//  V = Variable/Adjustable
	//
	// If indicator not available or unknown, blank fill.
	// Note: Report indicator ‘V’ for loans where the interest rate will be variable at some point, even if the interest rate starts as fixed.
	InterestTypeIndicator string `json:"interestTypeIndicator"`

	// Report the last name of the primary consumer.
	// Titles and prefixes should not be reported.
	//
	// If the surname contains multiple names, such as Paternal Name-Maternal Name, hyphenate the surnames.
	// For example, “SMITH-JONES” or “MARTINEZ-REYES” requires the hyphen.
	//
	// If the surname contains separate words, the hyphen is not required.
	// For example, report “VAN DYKE” or “DE LA CRUZ” with a space between each word.
	//
	// Other than the hyphen, do not report special characters in any of the Consumer Name fields.
	//
	// The Generation Code should be reported in Generation Code.
	//
	// Notes: Do not report minors.
	//        The name fields should not contain messages, such as “Parent of”, “Baby”, “Daughter”, “Child”, etc.
	//        Do not report trustee or estate accounts.
	//        In cases where the debt is included in a revocable trust and  the consumer retains contractual responsibility, report the consumer’s Full Name, Address, Social Security Number, and Date of Birth within the Base Segment fields.
	//        Do not report the name of the trust.
	Surname string `json:"surname" validate:"required"`

	// Report the full first name of the primary consumer.
	// Names should not be abbreviated.
	// Examples: Report first name “JUNIOR” (not “JR”); report “ROBERT” (not “ROBT”).
	//
	// If reporting multiple first names, hyphenate the first names.
	//
	// Note: If a consumer uses only initials for first and middle names (e.g., A.J.), the first name initial should be reported in the First Name field (e.g., A)            and the middle initial should be reported in the            Middle Name field (e.g., J).
	FirstName string `json:"firstName" validate:"required"`

	// Report the middle name or middle initial of the primary consumer, if available.
	//
	// If reporting multiple middle names, hyphenate the middle names.
	MiddleName string `json:"middleName,omitempty"`

	// Used to distinguish Jr., Sr., II, III, etc.  If not applicable, blank fill.
	// Values available:
	//  J = Junior
	//  3 = III
	//  6 = VI
	//  9 = IX
	//  S = Senior
	//  4 = IV
	//  7 = VII
	//  2 = II
	//  5 = V
	//  8 = VIII
	GenerationCode string `json:"generationCode,omitempty"`

	// Report the Social Security Number (SSN) of the primary consumer.
	// Report only valid U.S.-issued SSNs.
	//
	// Reporting of this information is required as the Social Security Number greatly enhances accuracy in matching to the correct consumer.
	// If the consumer does not have a SSN or one is not available for reporting, zero- or 9-fill all positions.
	//
	// Notes:
	//   If the Social Security Number is not reported, the Date of Birth is required to be reported.
	//   Do not report Individual Tax Identification Numbers  (ITINs) in this field.  ITINs do not prove identity outside the tax system and should not be offered or accepted as identification for non-tax purposes, per the Social Security Administration.
	//   Do not report Credit Profile Numbers (CPNs) in this field.  The CPN should not be used for credit reporting purposes and does not replace the Social Security Number.
	SocialSecurityNumber int `json:"socialSecurityNumber" validate:"required"`

	// Report the full Date of Birth of the primary consumer, including the month, day and year.
	// Reporting of this information is required as the Date of Birth greatly enhances accuracy in matching to the correct consumer.
	// Format for character date is MMDDYYYY.  Format for packed date is 0MMDDYYYYs — where s is the sign.
	//
	// Notes:  If the Date of Birth is not reported, the Social Security Number is required to be reported.
	//         When reporting Authorized Users (ECOA Code 3), the full Date of Birth (MMDDYYYY) must be reported for all newly-added Authorized Users on all pre-existing and newly-opened accounts, even if the Social Security Number is reported.
	//         Do not report accounts of consumers who are too young to enter into a binding contract.
	DateBirth utils.Time `json:"dateBirth" validate:"required"`

	// Contains the telephone number of the primary consumer (Area Code + 7 digits).
	TelephoneNumber int64 `json:"telephoneNumber"`

	// Defines the relationship of the primary consumer to the account and designates the account as joint, individual, etc., in compliance with the Equal Credit Opportunity Act.
	//
	// Exhibit 10 provides a list of ECOA Codes, their definitions and usage.
	//
	// For important information:
	// • Guidelines on reporting consumers who are personally liable for business accounts.
	// • Usage guidelines on ECOA Codes T (Terminated) and Z (Delete Consumer).
	//
	// Note: Codes 0 (Undesignated), 4 (Joint) and 6 (On-Behalf-Of) are obsolete as of September 2003 and may no longer be reported.
	ECOACode string `json:"ecoaCode" validate:"required"`

	// Contains a value that indicates a special condition of the account that applies to the primary consumer.
	//
	// This special condition may be that a bankruptcy was filed, discharged, dismissed or withdrawn; a debt was reaffirmed; or the consumer cannot be located or is now located.
	//
	// The indicator should be reported one time and will remain on file until another Consumer Information Indicator or a Removal value is reported.
	// As an option, the indicator may be reported each month as long as the condition applies.
	//
	// Regardless of the method of reporting, the indicator will be deleted only when another Consumer Information Indicator or a Removal value (Q, S, U) is reported.
	//
	// Exhibit 11 provides a list of Consumer Information Indicators and examples that demonstrate how to report these codes.
	//
	// For reporting guidelines, refer to Frequently Asked Questions 23 through 32 (bankruptcy) and 61 (personal receivership).
	ConsumerInformationIndicator string `json:"consumerInformationIndicator,omitempty"`

	// Contains the standard two-character country abbreviation.
	//
	// Exhibit 12 provides a list of the Country Codes.
	CountryCode string `json:"countryCode"`

	// Contains billing/mailing address for the primary consumer.
	// If the consumer has a U.S. address and a foreign address, report the U.S. address.
	// If the consumer has never used the U.S. address as a billing/mailing address (e.g., a property address), report the foreign address.
	//
	// If the billing/mailing address does not belong specifically to the consumer, such as a financial counseling site or bill paying service, report the consumer’s home address.
	//
	// The First Line of Address usually includes street number, direction, street name, and type of thoroughfare.
	//
	// If the billing/mailing address is a PO Box or Rural Route, include Box or Route followed by the number (e.g., PO Box 100).
	// Do not report both a street address and a PO Box.
	//
	// If the billing/mailing address is a private mailbox (PMB), the street address should be reported in the First Line of Address (e.g., 5678 Main Street).
	// The PMB number should be reported in the Second Line of Address (e.g., PMB 1234).
	// As an alternative, the entire address can be reported in the First Line of Address; for example, 5678 Main Street PMB 1234.
	//
	// Eliminate internal messages such as: “Do not mail”, “Attorney”, “Charge-off”, “Chapter 13”, “Fraud”, “Trustee”, “Estate of”, “Care of”, “M/R” (Mail Returned), etc.
	//
	// Exhibit 13 provides general rules for address reporting.
	// Do not enter data furnisher's address in this field.
	FirstLineAddress string `json:"firstLineAddress" validate:"required"`

	// Contains second line of address, if needed, such as apartment or unit number, or private mailbox number (PMB).
	//
	// Eliminate internal messages such as: “Do not mail”, “Attorney”, “Charge-off”, “Chapter 13”, “Fraud”, “Trustee”, or “Estate of”, “Care of”, “M/R” (Mail Returned), etc.
	SecondLineAddress string `json:"secondLineAddress,omitempty"`

	// Contains city name for address of primary consumer.
	// Truncate rightmost positions if city name is greater than 20 characters or use standard 13-character U.S. Postal Service city abbreviations.
	City string `json:"city" validate:"required"`

	// Contains the standard U.S. Postal Service state abbreviation for the address of the primary consumer.
	//
	// Exhibit 14 provides a list of State Codes.
	State string `json:"state" validate:"required"`

	// Report the Zip Code of the primary consumer’s address.
	// Use entire field if reporting 9-digit zip codes.
	// Otherwise, leftjustify and blank fill.
	ZipCode string `json:"zipCode" validate:"required"`

	// Contains one of the following values for the address reported in fields 40-44:
	//
	//  C  =  Confirmed/Verified address Note: Value ‘C’ enables reporting a confirmed or verified address after receiving an address discrepancy notification from a consumer reporting agency.  Report ‘C’ one time after the address is confirmed.
	//  Y = Known to be address of primary consumer
	//  N = Not confirmed address
	//  M = Military address
	//  S  =  Secondary Address
	//  B = Business address — not consumer's residence
	//  U = Non-deliverable address/Returned mail
	//  D  = Data reporter’s default address
	//  P  =  Bill Payer Service  — not consumer’s residence
	//
	// If indicator not available or unknown, blank fill.
	AddressIndicator string `json:"addressIndicator"`

	// Contains the one-character residence code of the address reported in fields 40-44.
	//
	//  Values available:
	//    O = Owns
	//    R = Rents
	//
	// If not available or unknown, blank fill.
	ResidenceCode string `json:"residenceCode"`

	j1Segments []Segment
	j2Segments []Segment
	k1Segment  Segment
	k2Segment  Segment
	k3Segment  Segment
	k4Segment  Segment
	l1Segment  Segment
	n1Segment  Segment

	converter
	validator
}

// PackedBaseSegment holds the packed base segment
type PackedBaseSegment BaseSegment

type baseJson BaseSegment

type dataRecordJson struct {
	Base       baseJson  `json:"base,omitempty"`
	J1Segments []Segment `json:"j1,omitempty"`
	J2Segments []Segment `json:"j2,omitempty"`
	K1Segment  Segment   `json:"k1,omitempty"`
	K2Segment  Segment   `json:"k2,omitempty"`
	K3Segment  Segment   `json:"k3,omitempty"`
	K4Segment  Segment   `json:"k4,omitempty"`
	L1Segment  Segment   `json:"l1,omitempty"`
	N1Segment  Segment   `json:"n1,omitempty"`
}

// Name returns name of base segment
func (r *BaseSegment) Name() string {
	return BaseSegmentName
}

// Parse takes the input record string and parses the base segment values
func (r *BaseSegment) Parse(record []byte) (int, error) {
	if len(record) < UnpackedRecordLength {
		return 0, utils.NewErrSegmentLength("base segment")
	}

	fields := reflect.ValueOf(r).Elem()
	length, err := r.parseRecordValues(fields, baseSegmentCharacterFormat, record, &r.validator, "base segment")
	if err != nil {
		return length, err
	}

	offset := UnpackedRecordLength
	if r.BlockDescriptorWord > 0 {
		offset = UnpackedRecordLength + 4
	}
	r.j1Segments = []Segment{}
	r.j2Segments = []Segment{}

	if len(record) < offset {
		return 0, utils.NewErrSegmentLength("base segment")
	}
	read, err := readApplicableSegments(record[offset:], r)
	if err != nil {
		return 0, err
	}

	if r.BlockDescriptorWord > 0 && read+offset > r.BlockDescriptorWord {
		return 0, utils.NewErrFailedParsing()
	}
	if r.BlockDescriptorWord == 0 && read+offset > r.RecordDescriptorWord {
		return 0, utils.NewErrFailedParsing()
	}

	if r.BlockDescriptorWord > 0 {
		return r.BlockDescriptorWord, nil
	}
	return r.RecordDescriptorWord, nil
}

// String writes the base segment struct to a 426 character string.
func (r *BaseSegment) String() string {
	var buf strings.Builder
	specifications := r.toSpecifications(baseSegmentCharacterFormat)
	fields := reflect.ValueOf(r).Elem()
	blockSize := r.RecordDescriptorWord
	if r.BlockDescriptorWord > 0 {
		blockSize += 4
	}
	buf.Grow(blockSize)
	for _, spec := range specifications {
		value := r.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	applicableSegment := ""
	for _, sub := range r.j1Segments {
		applicableSegment += sub.String()
	}
	for _, sub := range r.j2Segments {
		applicableSegment += sub.String()
	}
	if r.k1Segment != nil {
		applicableSegment += r.k1Segment.String()
	}
	if r.k2Segment != nil {
		applicableSegment += r.k2Segment.String()
	}
	if r.k3Segment != nil {
		applicableSegment += r.k3Segment.String()
	}
	if r.k4Segment != nil {
		applicableSegment += r.k4Segment.String()
	}
	if r.l1Segment != nil {
		applicableSegment += r.l1Segment.String()
	}
	if r.n1Segment != nil {
		applicableSegment += r.n1Segment.String()
	}
	buf.WriteString(applicableSegment)

	if blockSize > buf.Len() {
		buf.WriteString(strings.Repeat(blankString, blockSize-buf.Len()))
	}

	return buf.String()
}

// Bytes return raw byte array
func (r *BaseSegment) Bytes() []byte {
	return []byte(r.String())
}

// Validate performs some checks on the record and returns an error if not Validated
func (r *BaseSegment) Validate() error {
	return r.validateRecord(r, baseSegmentCharacterFormat, "base segment")
}

// BlockSize returns size of block
func (r *BaseSegment) BlockSize() int {
	return r.BlockDescriptorWord
}

// Length returns size of segment
func (r *BaseSegment) Length() int {
	return r.RecordDescriptorWord
}

// GetSegments returns list of applicable segments by segment name
func (r *BaseSegment) GetSegments(name string) []Segment {
	var ret []Segment
	switch name {
	case J1SegmentName:
		return r.j1Segments
	case J2SegmentName:
		return r.j2Segments
	case K1SegmentName:
		if r.k1Segment == nil {
			return nil
		}
		ret = append(ret, r.k1Segment)
	case K2SegmentName:
		if r.k2Segment == nil {
			return nil
		}
		ret = append(ret, r.k2Segment)
	case K3SegmentName:
		if r.k3Segment == nil {
			return nil
		}
		ret = append(ret, r.k3Segment)
	case K4SegmentName:
		if r.k4Segment == nil {
			return nil
		}
		ret = append(ret, r.k4Segment)
	case L1SegmentName:
		if r.l1Segment == nil {
			return nil
		}
		ret = append(ret, r.l1Segment)
	case N1SegmentName:
		if r.n1Segment == nil {
			return nil
		}
		ret = append(ret, r.n1Segment)
	default:
		return nil
	}
	return ret
}

// AddApplicableSegment will add new applicable segment into record
func (r *BaseSegment) AddApplicableSegment(s Segment) error {
	err := s.Validate()
	if err != nil {
		return err
	}
	switch s.Name() {
	case J1SegmentName:
		r.j1Segments = append(r.j1Segments, s)
	case J2SegmentName:
		r.j2Segments = append(r.j2Segments, s)
	case K1SegmentName:
		r.k1Segment = s
	case K2SegmentName:
		r.k2Segment = s
	case K3SegmentName:
		r.k3Segment = s
	case K4SegmentName:
		r.k4Segment = s
	case L1SegmentName:
		r.l1Segment = s
	case N1SegmentName:
		r.n1Segment = s
	}

	return nil
}

// MarshalJSON returns JSON blob
func (r *BaseSegment) MarshalJSON() ([]byte, error) {
	dummy := dataRecordJson{}
	base := baseJson{}

	fromFields := reflect.ValueOf(r).Elem()
	toFields := reflect.ValueOf(&base).Elem()
	for i := 0; i < fromFields.NumField(); i++ {
		fieldName := fromFields.Type().Field(i).Name
		fromField := fromFields.FieldByName(fieldName)
		toField := toFields.FieldByName(fieldName)
		if fromField.IsValid() && toField.CanSet() {
			toField.Set(fromField)
		}
	}

	dummy.Base = base
	dummy.J1Segments = r.j1Segments
	dummy.J2Segments = r.j2Segments
	dummy.K1Segment = r.k1Segment
	dummy.K2Segment = r.k2Segment
	dummy.K3Segment = r.k3Segment
	dummy.K4Segment = r.k4Segment
	dummy.L1Segment = r.l1Segment
	dummy.N1Segment = r.n1Segment

	return json.Marshal(dummy)
}

// UnmarshalJSON parses a JSON blob
func (r *BaseSegment) UnmarshalJSON(data []byte) error {

	dummy := make(map[string]interface{})
	err := json.Unmarshal(data, &dummy)
	if err != nil {
		return fmt.Errorf("invalid json format (%s)", err.Error())
	}

	r.j1Segments = []Segment{}
	r.j2Segments = []Segment{}

	for key, record := range dummy {
		buf, err := json.Marshal(record)
		if err != nil {
			return fmt.Errorf("invalid %s segment (%s)", key, err.Error())
		}

		switch key {
		case BaseSegmentName:
			base := baseJson{}
			err := json.Unmarshal(buf, &base)
			if err != nil {
				return fmt.Errorf("unabled to parse %s segment (%s)", key, err.Error())
			}

			fromFields := reflect.ValueOf(&base).Elem()
			toFields := reflect.ValueOf(r).Elem()
			for i := 0; i < fromFields.NumField(); i++ {
				fieldName := fromFields.Type().Field(i).Name
				fromField := fromFields.FieldByName(fieldName)
				toField := toFields.FieldByName(fieldName)
				if fromField.IsValid() && toField.CanSet() {
					toField.Set(fromField)
				}
			}
		case J1SegmentName, J2SegmentName:
			var list []interface{}
			err := json.Unmarshal(buf, &list)
			if err != nil {
				return fmt.Errorf("unabled to parse %s segment (%s)", key, err.Error())
			}

			for _, subSegment := range list {
				subBuf, err := json.Marshal(subSegment)
				if err != nil {
					return err
				}
				err = unmarshalApplicableSegments(key, subBuf, r)
				if err != nil {
					return err
				}
			}
		default:
			err = unmarshalApplicableSegments(key, buf, r)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *BaseSegment) ValidateIdentificationNumber() error {
	if validFilledString(r.IdentificationNumber) {
		return utils.NewErrInvalidValueOfField("identification number", "base segment")
	}
	return nil
}

func (r *BaseSegment) ValidatePortfolioType() error {
	switch r.PortfolioType {
	case PortfolioTypeCredit, PortfolioTypeInstallment, PortfolioTypeMortgage, PortfolioTypeOpen, PortfolioTypeRevolving:
		return nil
	}
	return utils.NewErrInvalidValueOfField("portfolio type", "base segment")
}

func (r *BaseSegment) ValidateTermsDuration() error {
	switch r.TermsDuration {
	case TermsDurationCredit, TermsDurationOpen, TermsDurationRevolving:
		return nil
	}
	_, err := strconv.Atoi(r.TermsDuration)
	if err != nil {
		return utils.NewErrInvalidValueOfField("terms duration", "base segment")
	}
	return nil
}

func (r *BaseSegment) ValidateTermsFrequency() error {
	switch r.TermsFrequency {
	case TermsFrequencyDeferred, TermsFrequencyPayment, TermsFrequencyWeekly, TermsFrequencyBiweekly,
		TermsFrequencySemimonthly, TermsFrequencyMonthly, TermsFrequencyBimonthly, TermsFrequencyQuarterly,
		TermsFrequencyTriAnnually, TermsFrequencySemiannually, TermsFrequencyAnnually, "":
		return nil
	}
	return utils.NewErrInvalidValueOfField("terms frequency", "base segment")
}

func (r *BaseSegment) ValidatePaymentRating() error {
	switch r.AccountStatus {
	case AccountStatus05, AccountStatus13, AccountStatus65, AccountStatus88, AccountStatus89, AccountStatus94, AccountStatus95:
		switch r.PaymentRating {
		case PaymentRatingCurrent, PaymentRatingPast30, PaymentRatingPast60, PaymentRatingPast90,
			PaymentRatingPast120, PaymentRatingPast150, PaymentRatingPast180, PaymentRatingCollection, PaymentRatingChargeOff:
			return nil
		}
		return utils.NewErrInvalidValueOfField("payment rating", "base segment")
	}

	if r.PaymentRating == "" {
		return nil
	}
	return utils.NewErrInvalidValueOfField("payment rating", "base segment")
}

func (r *BaseSegment) ValidatePaymentHistoryProfile() error {
	if len(r.PaymentHistoryProfile) != 24 {
		return utils.NewErrInvalidValueOfField("payment history profile", "base segment")
	}
	for i := 0; i < len(r.PaymentHistoryProfile); i++ {
		switch r.PaymentHistoryProfile[i] {
		case PaymentHistoryPast0, PaymentHistoryPast30, PaymentHistoryPast60, PaymentHistoryPast90,
			PaymentHistoryPast120, PaymentHistoryPast150, PaymentHistoryPast180, PaymentHistoryNoPayment,
			PaymentHistoryNoPaymentMonth, PaymentHistoryZero, PaymentHistoryCollection,
			PaymentHistoryForeclosureCompleted, PaymentHistoryVoluntarySurrender, PaymentHistoryRepossession,
			PaymentHistoryChargeOff:
			continue
		}
		return utils.NewErrInvalidValueOfField("payment history profile", "base segment")
	}
	return nil
}

func (r *BaseSegment) ValidateInterestTypeIndicator() error {
	switch r.InterestTypeIndicator {
	case InterestIndicatorFixed, InterestIndicatorVariable, "":
		return nil
	}
	return utils.NewErrInvalidValueOfField("interest type indicator", "base segment")
}

func (r *BaseSegment) ValidateTelephoneNumber() error {
	if err := r.isPhoneNumber(r.TelephoneNumber, "base segment"); err != nil {
		return err
	}
	return nil
}

// Name returns name of packed base segment
func (r *PackedBaseSegment) Name() string {
	return PackedBaseSegmentName
}

// Parse takes the input record string and parses the packed base segment values
func (r *PackedBaseSegment) Parse(record []byte) (int, error) {
	if len(record) < PackedRecordLength {
		return 0, utils.NewErrSegmentLength("packed base segment")
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
		spec, ok := baseSegmentPackedFormat[fieldName]
		if !ok || !field.IsValid() {
			return 0, utils.NewErrInvalidValueOfField(fieldName, "packed base segment")
		}

		if len(record) < spec.Start+spec.Length+offset {
			return 0, utils.NewErrSegmentLength("packed base segment")
		}
		data := string(record[spec.Start+offset : spec.Start+spec.Length+offset])
		if err := r.isValidType(spec, data, fieldName, "packed base segment"); err != nil {
			return 0, err
		}

		value, err := r.parseValue(spec, data, fieldName, "packed base segment")
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

	offset = PackedRecordLength + 4
	r.j1Segments = []Segment{}
	r.j2Segments = []Segment{}

	if len(record) < offset {
		return 0, utils.NewErrSegmentLength("packed base segment")
	}
	read, err := readApplicableSegments(record[offset:], r)
	if err != nil {
		return 0, err
	}

	if read+offset > r.BlockDescriptorWord {
		return 0, utils.NewErrFailedParsing()
	}

	return r.BlockDescriptorWord, nil
}

// String writes the packed base segment struct to a 426 character string.
func (r *PackedBaseSegment) String() string {
	var buf strings.Builder
	specifications := r.toSpecifications(baseSegmentPackedFormat)
	fields := reflect.ValueOf(r).Elem()
	blockSize := r.RecordDescriptorWord
	if r.BlockDescriptorWord > 0 {
		blockSize += 4
	}
	buf.Grow(blockSize)
	for _, spec := range specifications {
		value := r.toString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	applicableSegment := ""
	for _, sub := range r.j1Segments {
		applicableSegment += sub.String()
	}
	for _, sub := range r.j2Segments {
		applicableSegment += sub.String()
	}
	if r.k1Segment != nil {
		applicableSegment += r.k1Segment.String()
	}
	if r.k2Segment != nil {
		applicableSegment += r.k2Segment.String()
	}
	if r.k3Segment != nil {
		applicableSegment += r.k3Segment.String()
	}
	if r.k4Segment != nil {
		applicableSegment += r.k4Segment.String()
	}
	if r.l1Segment != nil {
		applicableSegment += r.l1Segment.String()
	}
	if r.n1Segment != nil {
		applicableSegment += r.n1Segment.String()
	}
	buf.WriteString(applicableSegment)

	if blockSize > buf.Len() {
		buf.WriteString(strings.Repeat(blankString, blockSize-buf.Len()))
	}

	return buf.String()
}

// Bytes return raw byte array
func (r *PackedBaseSegment) Bytes() []byte {
	return []byte(r.String())
}

// Validate performs some checks on the record and returns an error if not Validated
func (r *PackedBaseSegment) Validate() error {
	fields := reflect.ValueOf(r).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if spec, ok := baseSegmentPackedFormat[fieldName]; ok {
			if spec.Required == required {
				fieldValue := fields.FieldByName(fieldName)
				if fieldValue.IsZero() {
					return utils.NewErrFieldRequired(fieldName, "packed base segment")
				}
			}
		}

		funcName := r.validateFuncName(fieldName)
		method := reflect.ValueOf(r).MethodByName(funcName)
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

// BlockSize returns size of block
func (r *PackedBaseSegment) BlockSize() int {
	return r.BlockDescriptorWord
}

// Length returns size of segment
func (r *PackedBaseSegment) Length() int {
	return r.RecordDescriptorWord
}

// GetSegments returns list of applicable segments by segment name
func (r *PackedBaseSegment) GetSegments(name string) []Segment {
	var ret []Segment
	switch name {
	case J1SegmentName:
		return r.j1Segments
	case J2SegmentName:
		return r.j2Segments
	case K1SegmentName:
		if r.k1Segment == nil {
			return nil
		}
		ret = append(ret, r.k1Segment)
	case K2SegmentName:
		if r.k2Segment == nil {
			return nil
		}
		ret = append(ret, r.k2Segment)
	case K3SegmentName:
		if r.k3Segment == nil {
			return nil
		}
		ret = append(ret, r.k3Segment)
	case K4SegmentName:
		if r.k4Segment == nil {
			return nil
		}
		ret = append(ret, r.k4Segment)
	case L1SegmentName:
		if r.l1Segment == nil {
			return nil
		}
		ret = append(ret, r.l1Segment)
	case N1SegmentName:
		if r.n1Segment == nil {
			return nil
		}
		ret = append(ret, r.n1Segment)
	default:
		return nil
	}
	return ret
}

// AddApplicableSegment will add new applicable segment into record
func (r *PackedBaseSegment) AddApplicableSegment(s Segment) error {
	err := s.Validate()
	if err != nil {
		return err
	}
	switch s.Name() {
	case J1SegmentName:
		r.j1Segments = append(r.j1Segments, s)
	case J2SegmentName:
		r.j2Segments = append(r.j2Segments, s)
	case K1SegmentName:
		r.k1Segment = s
	case K2SegmentName:
		r.k2Segment = s
	case K3SegmentName:
		r.k3Segment = s
	case K4SegmentName:
		r.k4Segment = s
	case L1SegmentName:
		r.l1Segment = s
	case N1SegmentName:
		r.n1Segment = s
	}

	return nil
}

// MarshalJSON returns JSON blob
func (r *PackedBaseSegment) MarshalJSON() ([]byte, error) {
	dummy := dataRecordJson{}
	base := baseJson{}

	fromFields := reflect.ValueOf(r).Elem()
	toFields := reflect.ValueOf(&base).Elem()
	for i := 0; i < fromFields.NumField(); i++ {
		fieldName := fromFields.Type().Field(i).Name
		fromField := fromFields.FieldByName(fieldName)
		toField := toFields.FieldByName(fieldName)
		if fromField.IsValid() && toField.CanSet() {
			toField.Set(fromField)
		}
	}

	dummy.Base = base
	dummy.J1Segments = r.j1Segments
	dummy.J2Segments = r.j2Segments
	dummy.K1Segment = r.k1Segment
	dummy.K2Segment = r.k2Segment
	dummy.K3Segment = r.k3Segment
	dummy.K4Segment = r.k4Segment
	dummy.L1Segment = r.l1Segment
	dummy.N1Segment = r.n1Segment

	return json.Marshal(dummy)
}

// UnmarshalJSON parses a JSON blob
func (r *PackedBaseSegment) UnmarshalJSON(data []byte) error {

	dummy := make(map[string]interface{})
	err := json.Unmarshal(data, &dummy)
	if err != nil {
		return fmt.Errorf("invalid json format (%s)", err.Error())
	}

	r.j1Segments = []Segment{}
	r.j2Segments = []Segment{}

	for key, record := range dummy {
		buf, err := json.Marshal(record)
		if err != nil {
			return fmt.Errorf("invalid %s segment (%s)", key, err.Error())
		}

		switch key {
		case BaseSegmentName:
			base := baseJson{}
			err := json.Unmarshal(buf, &base)
			if err != nil {
				return fmt.Errorf("unabled to parse %s segment (%s)", key, err.Error())
			}
			fromFields := reflect.ValueOf(&base).Elem()
			toFields := reflect.ValueOf(r).Elem()
			for i := 0; i < fromFields.NumField(); i++ {
				fieldName := fromFields.Type().Field(i).Name
				fromField := fromFields.FieldByName(fieldName)
				toField := toFields.FieldByName(fieldName)
				if fromField.IsValid() && toField.CanSet() {
					toField.Set(fromField)
				}
			}
		case J1SegmentName, J2SegmentName:
			var list []interface{}
			err := json.Unmarshal(buf, &list)
			if err != nil {
				return fmt.Errorf("unabled to parse %s segment (%s)", key, err.Error())
			}
			for _, subSegment := range list {
				subBuf, err := json.Marshal(subSegment)
				if err != nil {
					return err
				}
				err = unmarshalApplicableSegments(key, subBuf, r)
				if err != nil {
					return err
				}
			}
		default:
			err = unmarshalApplicableSegments(key, buf, r)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *PackedBaseSegment) ValidateIdentificationNumber() error {
	if validFilledString(r.IdentificationNumber) {
		return utils.NewErrInvalidValueOfField("identification number", "packed base segment")
	}
	return nil
}

func (r *PackedBaseSegment) ValidatePortfolioType() error {
	switch r.PortfolioType {
	case PortfolioTypeCredit, PortfolioTypeInstallment, PortfolioTypeMortgage, PortfolioTypeOpen, PortfolioTypeRevolving:
		return nil
	}
	return utils.NewErrInvalidValueOfField("portfolio type", "packed base segment")
}

func (r *PackedBaseSegment) ValidateTermsDuration() error {
	switch r.TermsDuration {
	case TermsDurationCredit, TermsDurationOpen, TermsDurationRevolving:
		return nil
	}
	_, err := strconv.Atoi(r.TermsDuration)
	if err != nil {
		return utils.NewErrInvalidValueOfField("terms duration", "packed base segment")
	}
	return nil
}

func (r *PackedBaseSegment) ValidateTermsFrequency() error {
	switch r.TermsFrequency {
	case TermsFrequencyDeferred, TermsFrequencyPayment, TermsFrequencyWeekly, TermsFrequencyBiweekly,
		TermsFrequencySemimonthly, TermsFrequencyMonthly, TermsFrequencyBimonthly, TermsFrequencyQuarterly,
		TermsFrequencyTriAnnually, TermsFrequencySemiannually, TermsFrequencyAnnually, "":
		return nil
	}
	return utils.NewErrInvalidValueOfField("terms frequency", "packed base segment")
}

func (r *PackedBaseSegment) ValidatePaymentRating() error {
	switch r.AccountStatus {
	case AccountStatus05, AccountStatus13, AccountStatus65, AccountStatus88, AccountStatus89, AccountStatus94, AccountStatus95:
		switch r.PaymentRating {
		case PaymentRatingCurrent, PaymentRatingPast30, PaymentRatingPast60, PaymentRatingPast90,
			PaymentRatingPast120, PaymentRatingPast150, PaymentRatingPast180, PaymentRatingCollection, PaymentRatingChargeOff:
			return nil
		}
		return utils.NewErrInvalidValueOfField("payment rating", "packed base segment")
	}

	if r.PaymentRating == "" {
		return nil
	}
	return utils.NewErrInvalidValueOfField("payment rating", "packed base segment")
}

func (r *PackedBaseSegment) ValidatePaymentHistoryProfile() error {
	if len(r.PaymentHistoryProfile) != 24 {
		return utils.NewErrInvalidValueOfField("payment history profile", "packed base segment")
	}
	for i := 0; i < len(r.PaymentHistoryProfile); i++ {
		switch r.PaymentHistoryProfile[i] {
		case PaymentHistoryPast0, PaymentHistoryPast30, PaymentHistoryPast60, PaymentHistoryPast90,
			PaymentHistoryPast120, PaymentHistoryPast150, PaymentHistoryPast180, PaymentHistoryNoPayment,
			PaymentHistoryNoPaymentMonth, PaymentHistoryZero, PaymentHistoryCollection,
			PaymentHistoryForeclosureCompleted, PaymentHistoryVoluntarySurrender, PaymentHistoryRepossession,
			PaymentHistoryChargeOff:
			continue
		}
		return utils.NewErrInvalidValueOfField("payment history profile", "packed base segment")
	}
	return nil
}

func (r *PackedBaseSegment) ValidateInterestTypeIndicator() error {
	switch r.InterestTypeIndicator {
	case InterestIndicatorFixed, InterestIndicatorVariable, "":
		return nil
	}
	return utils.NewErrInvalidValueOfField("interest type indicator", "packed base segment")
}

func (r *PackedBaseSegment) ValidateTelephoneNumber() error {
	if err := r.isPhoneNumber(r.TelephoneNumber, "packed base segment"); err != nil {
		return err
	}
	return nil
}

func readApplicableSegments(record []byte, f Record) (int, error) {
	var segment Segment
	offset := 0

	for offset+2 < len(record) {
		switch string(record[offset : offset+2]) {
		case J1SegmentIdentifier:
			segment = NewJ1Segment()
		case J2SegmentIdentifier:
			segment = NewJ2Segment()
		case K1SegmentIdentifier:
			segment = NewK1Segment()
		case K2SegmentIdentifier:
			segment = NewK2Segment()
		case K3SegmentIdentifier:
			segment = NewK3Segment()
		case K4SegmentIdentifier:
			segment = NewK4Segment()
		case L1SegmentIdentifier:
			segment = NewL1Segment()
		case N1SegmentIdentifier:
			segment = NewN1Segment()
		default:
			return offset, nil
		}
		read, err := segment.Parse(record[offset:])
		if err != nil {
			return 0, err
		}
		err = f.AddApplicableSegment(segment)
		if err != nil {
			return 0, err
		}
		offset += read
	}

	return offset, nil
}

func unmarshalApplicableSegments(description string, data []byte, r Record) error {
	var segment Segment

	switch description {
	case J1SegmentName:
		segment = NewJ1Segment()
	case J2SegmentName:
		segment = NewJ2Segment()
	case K1SegmentName:
		segment = NewK1Segment()
	case K2SegmentName:
		segment = NewK2Segment()
	case K3SegmentName:
		segment = NewK3Segment()
	case K4SegmentName:
		segment = NewK4Segment()
	case L1SegmentName:
		segment = NewL1Segment()
	case N1SegmentName:
		segment = NewN1Segment()
	default:
		return nil
	}

	err := json.Unmarshal(data, segment)
	if err != nil {
		return fmt.Errorf("unabled to parse %s segment (%s)", description, err.Error())
	}
	return r.AddApplicableSegment(segment)
}
