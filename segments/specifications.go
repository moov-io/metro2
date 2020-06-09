// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

type field struct {
	Start    int
	Length   int
	Type     int
	Required string
}

type specification struct {
	Key   int
	Name  string
	Field field
}

const (
	nonrequired        = "N"
	required           = "Y"
	applicable         = "A"
	packedDateLongSize = 8
	packedDateSize     = 5
	int64Size          = 8
)

const (
	alphanumeric = 1 << iota
	alpha
	numeric
	binaryDescriptor
	packedDateLong
	packedDate
	packedNumber
	zeroFill
	omitted
)

var (
	baseSegmentCharacterFormat = map[string]field{
		"BlockDescriptorWord":           {0, 4, numeric | omitted, applicable},
		"RecordDescriptorWord":          {0, 4, numeric, required},
		"ProcessingIndicator":           {4, 1, numeric, nonrequired},
		"TimeStamp":                     {5, 14, numeric, nonrequired},
		"Reserved1":                     {19, 1, alphanumeric | zeroFill, nonrequired},
		"IdentificationNumber":          {20, 20, alphanumeric, required},
		"CycleIdentifier":               {40, 2, alphanumeric, applicable},
		"ConsumerAccountNumber":         {42, 30, alphanumeric, required},
		"PortfolioType":                 {72, 1, alphanumeric, required},
		"AccountType":                   {73, 2, alphanumeric, required},
		"DateOpened":                    {75, 8, numeric, required},
		"CreditLimit":                   {83, 9, numeric | zeroFill, applicable},
		"HighestCredit":                 {92, 9, numeric, required},
		"TermsDuration":                 {101, 3, alphanumeric, required},
		"TermsFrequency":                {104, 1, alphanumeric, applicable},
		"ScheduledMonthlyPaymentAmount": {105, 9, numeric, applicable},
		"ActualPaymentAmount":           {114, 9, numeric, applicable},
		"AccountStatus":                 {123, 2, alphanumeric, required},
		"PaymentRating":                 {125, 1, alphanumeric, applicable},
		"PaymentHistoryProfile":         {126, 24, alphanumeric, required},
		"SpecialComment":                {150, 2, alphanumeric, applicable},
		"ComplianceConditionCode":       {152, 2, alphanumeric, applicable},
		"CurrentBalance":                {154, 9, numeric, required},
		"AmountPastDue":                 {163, 9, numeric, applicable},
		"OriginalChargeOffAmount":       {172, 9, numeric, applicable},
		"DateAccountInformation":        {181, 8, numeric, required},
		"DateFirstDelinquency":          {189, 8, numeric, applicable},
		"DateClosed":                    {197, 8, numeric | zeroFill, applicable},
		"DateLastPayment":               {205, 8, numeric, applicable},
		"InterestTypeIndicator":         {213, 1, alphanumeric, nonrequired},
		"Reserved2":                     {214, 17, alphanumeric, nonrequired},
		"Surname":                       {231, 25, alphanumeric, required},
		"FirstName":                     {256, 20, alphanumeric, required},
		"MiddleName":                    {276, 20, alphanumeric, applicable},
		"GenerationCode":                {296, 1, alphanumeric, applicable},
		"SocialSecurityNumber":          {297, 9, numeric, required},
		"DateBirth":                     {306, 8, numeric, required},
		"TelephoneNumber":               {314, 10, numeric, nonrequired},
		"ECOACode":                      {324, 1, alphanumeric, required},
		"ConsumerInformationIndicator":  {325, 2, alphanumeric, applicable},
		"CountryCode":                   {327, 2, alphanumeric, nonrequired},
		"FirstLineAddress":              {329, 32, alphanumeric, required},
		"SecondLineAddress":             {361, 32, alphanumeric, applicable},
		"City":                          {393, 20, alphanumeric, required},
		"State":                         {413, 2, alphanumeric, required},
		"ZipCode":                       {415, 9, alphanumeric, required},
		"AddressIndicator":              {424, 1, alphanumeric, nonrequired},
		"ResidenceCode":                 {425, 1, alphanumeric, nonrequired},
	}
	baseSegmentPackedFormat = map[string]field{
		"BlockDescriptorWord":           {0, 4, binaryDescriptor | omitted, applicable},
		"RecordDescriptorWord":          {0, 4, binaryDescriptor, required},
		"ProcessingIndicator":           {4, 1, numeric, nonrequired},
		"TimeStamp":                     {5, 8, packedDateLong, nonrequired},
		"Reserved1":                     {13, 1, alphanumeric | zeroFill, nonrequired},
		"IdentificationNumber":          {14, 20, alphanumeric, required},
		"CycleIdentifier":               {34, 2, alphanumeric, applicable},
		"ConsumerAccountNumber":         {36, 30, alphanumeric, required},
		"PortfolioType":                 {66, 1, alphanumeric, required},
		"AccountType":                   {67, 2, alphanumeric, required},
		"DateOpened":                    {69, 5, packedDate, required},
		"CreditLimit":                   {74, 5, packedNumber | zeroFill, applicable},
		"HighestCredit":                 {79, 5, packedNumber, required},
		"TermsDuration":                 {84, 3, alphanumeric, required},
		"TermsFrequency":                {87, 1, alphanumeric, applicable},
		"ScheduledMonthlyPaymentAmount": {88, 5, packedNumber, applicable},
		"ActualPaymentAmount":           {93, 5, packedNumber, applicable},
		"AccountStatus":                 {98, 2, alphanumeric, required},
		"PaymentRating":                 {100, 1, alphanumeric, applicable},
		"PaymentHistoryProfile":         {101, 24, alphanumeric, required},
		"SpecialComment":                {125, 2, alphanumeric, applicable},
		"ComplianceConditionCode":       {127, 2, alphanumeric, applicable},
		"CurrentBalance":                {129, 5, packedNumber, required},
		"AmountPastDue":                 {134, 5, packedNumber, applicable},
		"OriginalChargeOffAmount":       {139, 5, packedNumber, applicable},
		"DateAccountInformation":        {144, 5, packedDate, required},
		"DateFirstDelinquency":          {149, 5, packedDate, applicable},
		"DateClosed":                    {154, 5, packedDate | zeroFill, applicable},
		"DateLastPayment":               {159, 5, packedDate, applicable},
		"InterestTypeIndicator":         {164, 1, alphanumeric, nonrequired},
		"Reserved2":                     {165, 17, alphanumeric, nonrequired},
		"Surname":                       {182, 25, alphanumeric, required},
		"FirstName":                     {207, 20, alphanumeric, required},
		"MiddleName":                    {227, 20, alphanumeric, applicable},
		"GenerationCode":                {247, 1, alphanumeric, applicable},
		"SocialSecurityNumber":          {248, 5, packedNumber, required},
		"DateBirth":                     {253, 5, packedDate, required},
		"TelephoneNumber":               {258, 6, packedNumber, nonrequired},
		"ECOACode":                      {264, 1, alphanumeric, required},
		"ConsumerInformationIndicator":  {265, 2, alphanumeric, applicable},
		"CountryCode":                   {267, 2, alphanumeric, nonrequired},
		"FirstLineAddress":              {269, 32, alphanumeric, required},
		"SecondLineAddress":             {301, 32, alphanumeric, applicable},
		"City":                          {333, 20, alphanumeric, required},
		"State":                         {353, 2, alphanumeric, required},
		"ZipCode":                       {355, 9, alphanumeric, required},
		"AddressIndicator":              {364, 1, alphanumeric, nonrequired},
		"ResidenceCode":                 {365, 1, alphanumeric, nonrequired},
	}
	headerRecordCharacterFormat = map[string]field{
		"BlockDescriptorWord":         {0, 4, numeric | omitted, applicable},
		"RecordDescriptorWord":        {0, 4, numeric, required},
		"RecordIdentifier":            {4, 6, alphanumeric, required},
		"CycleIdentifier":             {10, 2, alphanumeric, applicable},
		"InnovisProgramIdentifier":    {12, 10, alphanumeric, applicable},
		"EquifaxProgramIdentifier":    {22, 10, alphanumeric, applicable},
		"ExperianProgramIdentifier":   {32, 5, alphanumeric, applicable},
		"TransUnionProgramIdentifier": {37, 10, alphanumeric, applicable},
		"ActivityDate":                {47, 8, numeric, required},
		"DateCreated":                 {55, 8, numeric, required},
		"ProgramDate":                 {63, 8, numeric, nonrequired},
		"ProgramRevisionDate":         {71, 8, numeric, nonrequired},
		"ReporterName":                {79, 40, alphanumeric, required},
		"ReporterAddress":             {119, 96, alphanumeric, required},
		"ReporterTelephoneNumber":     {215, 10, numeric, nonrequired},
		"SoftwareVendorName":          {225, 40, alphanumeric, applicable},
		"SoftwareVersionNumber":       {265, 5, alphanumeric, applicable},
		"PRBCProgramIdentifier":       {270, 10, alphanumeric, applicable},
		"Reserved":                    {280, 146, alphanumeric, nonrequired},
	}
	headerRecordPackedFormat = map[string]field{
		"BlockDescriptorWord":         {0, 4, binaryDescriptor | omitted, applicable},
		"RecordDescriptorWord":        {0, 4, binaryDescriptor, required},
		"RecordIdentifier":            {4, 6, alphanumeric, required},
		"CycleIdentifier":             {10, 2, alphanumeric, applicable},
		"InnovisProgramIdentifier":    {12, 10, alphanumeric, applicable},
		"EquifaxProgramIdentifier":    {22, 10, alphanumeric, applicable},
		"ExperianProgramIdentifier":   {32, 5, alphanumeric, applicable},
		"TransUnionProgramIdentifier": {37, 10, alphanumeric, applicable},
		"ActivityDate":                {47, 8, numeric, required},
		"DateCreated":                 {55, 8, numeric, required},
		"ProgramDate":                 {63, 8, numeric, nonrequired},
		"ProgramRevisionDate":         {71, 8, numeric, nonrequired},
		"ReporterName":                {79, 40, alphanumeric, required},
		"ReporterAddress":             {119, 96, alphanumeric, required},
		"ReporterTelephoneNumber":     {215, 10, numeric, nonrequired},
		"SoftwareVendorName":          {225, 40, alphanumeric, applicable},
		"SoftwareVersionNumber":       {265, 5, alphanumeric, applicable},
		"PRBCProgramIdentifier":       {270, 10, alphanumeric, applicable},
		"Reserved":                    {280, 86, alphanumeric, nonrequired},
	}
)
