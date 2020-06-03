// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

type Field struct {
	Start    int
	Length   int
	Type     int
	Required string
}

type Specification struct {
	Key   int
	Name  string
	Field Field
}

const (
	NonRequired = "N"
	Required    = "Y"
	Applicable  = "A"
)

const (
	Alphanumeric = 1 << iota
	Alpha
	Numeric
	Binary
	Packed
	ZeroFill
	Omitted
)

var (
	BaseSegmentCharacterFormat = map[string]Field{
		"BlockDescriptorWord":           {0, 4, Numeric | Omitted, Applicable},
		"RecordDescriptorWord":          {0, 4, Numeric, Required},
		"ProcessingIndicator":           {4, 1, Numeric, NonRequired},
		"TimeStamp":                     {5, 14, Numeric, NonRequired},
		"Reserved1":                     {19, 1, Alphanumeric | ZeroFill, NonRequired},
		"IdentificationNumber":          {20, 20, Alphanumeric, Required},
		"CycleIdentifier":               {40, 2, Alphanumeric, Applicable},
		"ConsumerAccountNumber":         {42, 30, Alphanumeric, Required},
		"PortfolioType":                 {72, 1, Alphanumeric, Required},
		"AccountType":                   {73, 2, Alphanumeric, Required},
		"DateOpened":                    {75, 8, Numeric, Required},
		"CreditLimit":                   {83, 9, Numeric | ZeroFill, Applicable},
		"HighestCredit":                 {92, 9, Numeric, Required},
		"TermsDuration":                 {101, 3, Alphanumeric, Required},
		"TermsFrequency":                {104, 1, Alphanumeric, Applicable},
		"ScheduledMonthlyPaymentAmount": {105, 9, Numeric, Applicable},
		"ActualPaymentAmount":           {114, 9, Numeric, Applicable},
		"AccountStatus":                 {123, 2, Alphanumeric, Required},
		"PaymentRating":                 {125, 1, Alphanumeric, Applicable},
		"PaymentHistoryProfile":         {126, 24, Alphanumeric, Required},
		"SpecialComment":                {150, 2, Alphanumeric, Applicable},
		"ComplianceConditionCode":       {152, 2, Alphanumeric, Applicable},
		"CurrentBalance":                {154, 9, Numeric, Required},
		"AmountPastDue":                 {163, 9, Numeric, Applicable},
		"OriginalChargeOffAmount":       {172, 9, Numeric, Applicable},
		"DateAccountInformation":        {181, 8, Numeric, Required},
		"DateFirstDelinquency":          {189, 8, Numeric, Applicable},
		"DateClosed":                    {197, 8, Numeric | ZeroFill, Applicable},
		"DateLastPayment":               {205, 8, Numeric, Applicable},
		"InterestTypeIndicator":         {213, 1, Alphanumeric, NonRequired},
		"Reserved2":                     {214, 17, Alphanumeric, NonRequired},
		"Surname":                       {231, 25, Alphanumeric, Required},
		"FirstName":                     {256, 20, Alphanumeric, Required},
		"MiddleName":                    {276, 20, Alphanumeric, Applicable},
		"GenerationCode":                {296, 1, Alphanumeric, Applicable},
		"SocialSecurityNumber":          {297, 9, Numeric, Required},
		"DateBirth":                     {306, 8, Numeric, Required},
		"TelephoneNumber":               {314, 10, Numeric, NonRequired},
		"ECOACode":                      {324, 1, Alphanumeric, Required},
		"ConsumerInformationIndicator":  {325, 2, Alphanumeric, Applicable},
		"CountryCode":                   {327, 2, Alphanumeric, NonRequired},
		"FirstLineAddress":              {329, 32, Alphanumeric, Required},
		"SecondLineAddress":             {361, 32, Alphanumeric, Applicable},
		"City":                          {393, 20, Alphanumeric, Required},
		"State":                         {413, 2, Alphanumeric, Required},
		"ZipCode":                       {415, 9, Alphanumeric, Required},
		"AddressIndicator":              {424, 1, Alphanumeric, NonRequired},
		"ResidenceCode":                 {425, 1, Alphanumeric, NonRequired},
	}
)
