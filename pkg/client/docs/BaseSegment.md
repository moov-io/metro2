# BaseSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockDescriptorWord** | Pointer to **int32** |  | [optional] 
**RecordDescriptorWord** | **int32** |  | 
**TimeStamp** | Pointer to **time.Time** |  | [optional] 
**IdentificationNumber** | **string** |  | 
**CycleIdentifier** | Pointer to **string** |  | [optional] 
**ConsumerAccountNumber** | **string** |  | 
**PortfolioType** | Pointer to **string** |  | [optional] 
**AccountType** | **string** |  | 
**DateOpened** | Pointer to **time.Time** |  | [optional] 
**CreditLimit** | Pointer to **int32** |  | [optional] 
**HighestCredit** | **int32** |  | 
**TermsDuration** | **string** |  | 
**TermsFrequency** | Pointer to **string** |  | [optional] 
**ScheduledMonthlyPaymentAmount** | Pointer to **int32** |  | [optional] 
**ActualPaymentAmount** | Pointer to **int32** |  | [optional] 
**AccountStatus** | **string** |  | 
**PaymentRating** | Pointer to **string** |  | [optional] 
**PaymentHistoryProfile** | **string** |  | 
**SpecialComment** | Pointer to **string** |  | [optional] 
**ComplianceConditionCode** | Pointer to **string** |  | [optional] 
**CurrentBalance** | **int32** |  | 
**AmountPastDue** | Pointer to **int32** |  | [optional] 
**OriginalChargeOffAmount** | Pointer to **int32** |  | [optional] 
**DateAccountInformation** | **time.Time** |  | 
**DateFirstDelinquency** | Pointer to **time.Time** |  | [optional] 
**DateClosed** | Pointer to **time.Time** |  | [optional] 
**DateLastPayment** | Pointer to **time.Time** |  | [optional] 
**InterestTypeIndicator** | Pointer to **string** |  | [optional] 
**Surname** | **string** |  | 
**FirstName** | **string** |  | 
**MiddleName** | Pointer to **string** |  | [optional] 
**GenerationCode** | Pointer to **string** |  | [optional] 
**SocialSecurityNumber** | **int32** |  | 
**DateBirth** | **time.Time** |  | 
**TelephoneNumber** | Pointer to **int64** |  | [optional] 
**EcoaCode** | **string** |  | 
**ConsumerInformationIndicator** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**FirstLineAddress** | **string** |  | 
**SecondLineAddress** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | **string** |  | 
**ZipCode** | **string** |  | 
**AddressIndicator** | Pointer to **string** |  | [optional] 
**ResidenceCode** | Pointer to **string** |  | [optional] 

## Methods

### NewBaseSegment

`func NewBaseSegment(recordDescriptorWord int32, identificationNumber string, consumerAccountNumber string, accountType string, highestCredit int32, termsDuration string, accountStatus string, paymentHistoryProfile string, currentBalance int32, dateAccountInformation time.Time, surname string, firstName string, socialSecurityNumber int32, dateBirth time.Time, ecoaCode string, firstLineAddress string, state string, zipCode string, ) *BaseSegment`

NewBaseSegment instantiates a new BaseSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseSegmentWithDefaults

`func NewBaseSegmentWithDefaults() *BaseSegment`

NewBaseSegmentWithDefaults instantiates a new BaseSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockDescriptorWord

`func (o *BaseSegment) GetBlockDescriptorWord() int32`

GetBlockDescriptorWord returns the BlockDescriptorWord field if non-nil, zero value otherwise.

### GetBlockDescriptorWordOk

`func (o *BaseSegment) GetBlockDescriptorWordOk() (*int32, bool)`

GetBlockDescriptorWordOk returns a tuple with the BlockDescriptorWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDescriptorWord

`func (o *BaseSegment) SetBlockDescriptorWord(v int32)`

SetBlockDescriptorWord sets BlockDescriptorWord field to given value.

### HasBlockDescriptorWord

`func (o *BaseSegment) HasBlockDescriptorWord() bool`

HasBlockDescriptorWord returns a boolean if a field has been set.

### GetRecordDescriptorWord

`func (o *BaseSegment) GetRecordDescriptorWord() int32`

GetRecordDescriptorWord returns the RecordDescriptorWord field if non-nil, zero value otherwise.

### GetRecordDescriptorWordOk

`func (o *BaseSegment) GetRecordDescriptorWordOk() (*int32, bool)`

GetRecordDescriptorWordOk returns a tuple with the RecordDescriptorWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordDescriptorWord

`func (o *BaseSegment) SetRecordDescriptorWord(v int32)`

SetRecordDescriptorWord sets RecordDescriptorWord field to given value.


### GetTimeStamp

`func (o *BaseSegment) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *BaseSegment) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *BaseSegment) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *BaseSegment) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetIdentificationNumber

`func (o *BaseSegment) GetIdentificationNumber() string`

GetIdentificationNumber returns the IdentificationNumber field if non-nil, zero value otherwise.

### GetIdentificationNumberOk

`func (o *BaseSegment) GetIdentificationNumberOk() (*string, bool)`

GetIdentificationNumberOk returns a tuple with the IdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationNumber

`func (o *BaseSegment) SetIdentificationNumber(v string)`

SetIdentificationNumber sets IdentificationNumber field to given value.


### GetCycleIdentifier

`func (o *BaseSegment) GetCycleIdentifier() string`

GetCycleIdentifier returns the CycleIdentifier field if non-nil, zero value otherwise.

### GetCycleIdentifierOk

`func (o *BaseSegment) GetCycleIdentifierOk() (*string, bool)`

GetCycleIdentifierOk returns a tuple with the CycleIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleIdentifier

`func (o *BaseSegment) SetCycleIdentifier(v string)`

SetCycleIdentifier sets CycleIdentifier field to given value.

### HasCycleIdentifier

`func (o *BaseSegment) HasCycleIdentifier() bool`

HasCycleIdentifier returns a boolean if a field has been set.

### GetConsumerAccountNumber

`func (o *BaseSegment) GetConsumerAccountNumber() string`

GetConsumerAccountNumber returns the ConsumerAccountNumber field if non-nil, zero value otherwise.

### GetConsumerAccountNumberOk

`func (o *BaseSegment) GetConsumerAccountNumberOk() (*string, bool)`

GetConsumerAccountNumberOk returns a tuple with the ConsumerAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerAccountNumber

`func (o *BaseSegment) SetConsumerAccountNumber(v string)`

SetConsumerAccountNumber sets ConsumerAccountNumber field to given value.


### GetPortfolioType

`func (o *BaseSegment) GetPortfolioType() string`

GetPortfolioType returns the PortfolioType field if non-nil, zero value otherwise.

### GetPortfolioTypeOk

`func (o *BaseSegment) GetPortfolioTypeOk() (*string, bool)`

GetPortfolioTypeOk returns a tuple with the PortfolioType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolioType

`func (o *BaseSegment) SetPortfolioType(v string)`

SetPortfolioType sets PortfolioType field to given value.

### HasPortfolioType

`func (o *BaseSegment) HasPortfolioType() bool`

HasPortfolioType returns a boolean if a field has been set.

### GetAccountType

`func (o *BaseSegment) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BaseSegment) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BaseSegment) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetDateOpened

`func (o *BaseSegment) GetDateOpened() time.Time`

GetDateOpened returns the DateOpened field if non-nil, zero value otherwise.

### GetDateOpenedOk

`func (o *BaseSegment) GetDateOpenedOk() (*time.Time, bool)`

GetDateOpenedOk returns a tuple with the DateOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOpened

`func (o *BaseSegment) SetDateOpened(v time.Time)`

SetDateOpened sets DateOpened field to given value.

### HasDateOpened

`func (o *BaseSegment) HasDateOpened() bool`

HasDateOpened returns a boolean if a field has been set.

### GetCreditLimit

`func (o *BaseSegment) GetCreditLimit() int32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *BaseSegment) GetCreditLimitOk() (*int32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *BaseSegment) SetCreditLimit(v int32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *BaseSegment) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetHighestCredit

`func (o *BaseSegment) GetHighestCredit() int32`

GetHighestCredit returns the HighestCredit field if non-nil, zero value otherwise.

### GetHighestCreditOk

`func (o *BaseSegment) GetHighestCreditOk() (*int32, bool)`

GetHighestCreditOk returns a tuple with the HighestCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestCredit

`func (o *BaseSegment) SetHighestCredit(v int32)`

SetHighestCredit sets HighestCredit field to given value.


### GetTermsDuration

`func (o *BaseSegment) GetTermsDuration() string`

GetTermsDuration returns the TermsDuration field if non-nil, zero value otherwise.

### GetTermsDurationOk

`func (o *BaseSegment) GetTermsDurationOk() (*string, bool)`

GetTermsDurationOk returns a tuple with the TermsDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsDuration

`func (o *BaseSegment) SetTermsDuration(v string)`

SetTermsDuration sets TermsDuration field to given value.


### GetTermsFrequency

`func (o *BaseSegment) GetTermsFrequency() string`

GetTermsFrequency returns the TermsFrequency field if non-nil, zero value otherwise.

### GetTermsFrequencyOk

`func (o *BaseSegment) GetTermsFrequencyOk() (*string, bool)`

GetTermsFrequencyOk returns a tuple with the TermsFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsFrequency

`func (o *BaseSegment) SetTermsFrequency(v string)`

SetTermsFrequency sets TermsFrequency field to given value.

### HasTermsFrequency

`func (o *BaseSegment) HasTermsFrequency() bool`

HasTermsFrequency returns a boolean if a field has been set.

### GetScheduledMonthlyPaymentAmount

`func (o *BaseSegment) GetScheduledMonthlyPaymentAmount() int32`

GetScheduledMonthlyPaymentAmount returns the ScheduledMonthlyPaymentAmount field if non-nil, zero value otherwise.

### GetScheduledMonthlyPaymentAmountOk

`func (o *BaseSegment) GetScheduledMonthlyPaymentAmountOk() (*int32, bool)`

GetScheduledMonthlyPaymentAmountOk returns a tuple with the ScheduledMonthlyPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMonthlyPaymentAmount

`func (o *BaseSegment) SetScheduledMonthlyPaymentAmount(v int32)`

SetScheduledMonthlyPaymentAmount sets ScheduledMonthlyPaymentAmount field to given value.

### HasScheduledMonthlyPaymentAmount

`func (o *BaseSegment) HasScheduledMonthlyPaymentAmount() bool`

HasScheduledMonthlyPaymentAmount returns a boolean if a field has been set.

### GetActualPaymentAmount

`func (o *BaseSegment) GetActualPaymentAmount() int32`

GetActualPaymentAmount returns the ActualPaymentAmount field if non-nil, zero value otherwise.

### GetActualPaymentAmountOk

`func (o *BaseSegment) GetActualPaymentAmountOk() (*int32, bool)`

GetActualPaymentAmountOk returns a tuple with the ActualPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPaymentAmount

`func (o *BaseSegment) SetActualPaymentAmount(v int32)`

SetActualPaymentAmount sets ActualPaymentAmount field to given value.

### HasActualPaymentAmount

`func (o *BaseSegment) HasActualPaymentAmount() bool`

HasActualPaymentAmount returns a boolean if a field has been set.

### GetAccountStatus

`func (o *BaseSegment) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *BaseSegment) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *BaseSegment) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.


### GetPaymentRating

`func (o *BaseSegment) GetPaymentRating() string`

GetPaymentRating returns the PaymentRating field if non-nil, zero value otherwise.

### GetPaymentRatingOk

`func (o *BaseSegment) GetPaymentRatingOk() (*string, bool)`

GetPaymentRatingOk returns a tuple with the PaymentRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRating

`func (o *BaseSegment) SetPaymentRating(v string)`

SetPaymentRating sets PaymentRating field to given value.

### HasPaymentRating

`func (o *BaseSegment) HasPaymentRating() bool`

HasPaymentRating returns a boolean if a field has been set.

### GetPaymentHistoryProfile

`func (o *BaseSegment) GetPaymentHistoryProfile() string`

GetPaymentHistoryProfile returns the PaymentHistoryProfile field if non-nil, zero value otherwise.

### GetPaymentHistoryProfileOk

`func (o *BaseSegment) GetPaymentHistoryProfileOk() (*string, bool)`

GetPaymentHistoryProfileOk returns a tuple with the PaymentHistoryProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentHistoryProfile

`func (o *BaseSegment) SetPaymentHistoryProfile(v string)`

SetPaymentHistoryProfile sets PaymentHistoryProfile field to given value.


### GetSpecialComment

`func (o *BaseSegment) GetSpecialComment() string`

GetSpecialComment returns the SpecialComment field if non-nil, zero value otherwise.

### GetSpecialCommentOk

`func (o *BaseSegment) GetSpecialCommentOk() (*string, bool)`

GetSpecialCommentOk returns a tuple with the SpecialComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialComment

`func (o *BaseSegment) SetSpecialComment(v string)`

SetSpecialComment sets SpecialComment field to given value.

### HasSpecialComment

`func (o *BaseSegment) HasSpecialComment() bool`

HasSpecialComment returns a boolean if a field has been set.

### GetComplianceConditionCode

`func (o *BaseSegment) GetComplianceConditionCode() string`

GetComplianceConditionCode returns the ComplianceConditionCode field if non-nil, zero value otherwise.

### GetComplianceConditionCodeOk

`func (o *BaseSegment) GetComplianceConditionCodeOk() (*string, bool)`

GetComplianceConditionCodeOk returns a tuple with the ComplianceConditionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceConditionCode

`func (o *BaseSegment) SetComplianceConditionCode(v string)`

SetComplianceConditionCode sets ComplianceConditionCode field to given value.

### HasComplianceConditionCode

`func (o *BaseSegment) HasComplianceConditionCode() bool`

HasComplianceConditionCode returns a boolean if a field has been set.

### GetCurrentBalance

`func (o *BaseSegment) GetCurrentBalance() int32`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *BaseSegment) GetCurrentBalanceOk() (*int32, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalance

`func (o *BaseSegment) SetCurrentBalance(v int32)`

SetCurrentBalance sets CurrentBalance field to given value.


### GetAmountPastDue

`func (o *BaseSegment) GetAmountPastDue() int32`

GetAmountPastDue returns the AmountPastDue field if non-nil, zero value otherwise.

### GetAmountPastDueOk

`func (o *BaseSegment) GetAmountPastDueOk() (*int32, bool)`

GetAmountPastDueOk returns a tuple with the AmountPastDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPastDue

`func (o *BaseSegment) SetAmountPastDue(v int32)`

SetAmountPastDue sets AmountPastDue field to given value.

### HasAmountPastDue

`func (o *BaseSegment) HasAmountPastDue() bool`

HasAmountPastDue returns a boolean if a field has been set.

### GetOriginalChargeOffAmount

`func (o *BaseSegment) GetOriginalChargeOffAmount() int32`

GetOriginalChargeOffAmount returns the OriginalChargeOffAmount field if non-nil, zero value otherwise.

### GetOriginalChargeOffAmountOk

`func (o *BaseSegment) GetOriginalChargeOffAmountOk() (*int32, bool)`

GetOriginalChargeOffAmountOk returns a tuple with the OriginalChargeOffAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalChargeOffAmount

`func (o *BaseSegment) SetOriginalChargeOffAmount(v int32)`

SetOriginalChargeOffAmount sets OriginalChargeOffAmount field to given value.

### HasOriginalChargeOffAmount

`func (o *BaseSegment) HasOriginalChargeOffAmount() bool`

HasOriginalChargeOffAmount returns a boolean if a field has been set.

### GetDateAccountInformation

`func (o *BaseSegment) GetDateAccountInformation() time.Time`

GetDateAccountInformation returns the DateAccountInformation field if non-nil, zero value otherwise.

### GetDateAccountInformationOk

`func (o *BaseSegment) GetDateAccountInformationOk() (*time.Time, bool)`

GetDateAccountInformationOk returns a tuple with the DateAccountInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAccountInformation

`func (o *BaseSegment) SetDateAccountInformation(v time.Time)`

SetDateAccountInformation sets DateAccountInformation field to given value.


### GetDateFirstDelinquency

`func (o *BaseSegment) GetDateFirstDelinquency() time.Time`

GetDateFirstDelinquency returns the DateFirstDelinquency field if non-nil, zero value otherwise.

### GetDateFirstDelinquencyOk

`func (o *BaseSegment) GetDateFirstDelinquencyOk() (*time.Time, bool)`

GetDateFirstDelinquencyOk returns a tuple with the DateFirstDelinquency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFirstDelinquency

`func (o *BaseSegment) SetDateFirstDelinquency(v time.Time)`

SetDateFirstDelinquency sets DateFirstDelinquency field to given value.

### HasDateFirstDelinquency

`func (o *BaseSegment) HasDateFirstDelinquency() bool`

HasDateFirstDelinquency returns a boolean if a field has been set.

### GetDateClosed

`func (o *BaseSegment) GetDateClosed() time.Time`

GetDateClosed returns the DateClosed field if non-nil, zero value otherwise.

### GetDateClosedOk

`func (o *BaseSegment) GetDateClosedOk() (*time.Time, bool)`

GetDateClosedOk returns a tuple with the DateClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClosed

`func (o *BaseSegment) SetDateClosed(v time.Time)`

SetDateClosed sets DateClosed field to given value.

### HasDateClosed

`func (o *BaseSegment) HasDateClosed() bool`

HasDateClosed returns a boolean if a field has been set.

### GetDateLastPayment

`func (o *BaseSegment) GetDateLastPayment() time.Time`

GetDateLastPayment returns the DateLastPayment field if non-nil, zero value otherwise.

### GetDateLastPaymentOk

`func (o *BaseSegment) GetDateLastPaymentOk() (*time.Time, bool)`

GetDateLastPaymentOk returns a tuple with the DateLastPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastPayment

`func (o *BaseSegment) SetDateLastPayment(v time.Time)`

SetDateLastPayment sets DateLastPayment field to given value.

### HasDateLastPayment

`func (o *BaseSegment) HasDateLastPayment() bool`

HasDateLastPayment returns a boolean if a field has been set.

### GetInterestTypeIndicator

`func (o *BaseSegment) GetInterestTypeIndicator() string`

GetInterestTypeIndicator returns the InterestTypeIndicator field if non-nil, zero value otherwise.

### GetInterestTypeIndicatorOk

`func (o *BaseSegment) GetInterestTypeIndicatorOk() (*string, bool)`

GetInterestTypeIndicatorOk returns a tuple with the InterestTypeIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestTypeIndicator

`func (o *BaseSegment) SetInterestTypeIndicator(v string)`

SetInterestTypeIndicator sets InterestTypeIndicator field to given value.

### HasInterestTypeIndicator

`func (o *BaseSegment) HasInterestTypeIndicator() bool`

HasInterestTypeIndicator returns a boolean if a field has been set.

### GetSurname

`func (o *BaseSegment) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *BaseSegment) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *BaseSegment) SetSurname(v string)`

SetSurname sets Surname field to given value.


### GetFirstName

`func (o *BaseSegment) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BaseSegment) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BaseSegment) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *BaseSegment) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *BaseSegment) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *BaseSegment) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *BaseSegment) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetGenerationCode

`func (o *BaseSegment) GetGenerationCode() string`

GetGenerationCode returns the GenerationCode field if non-nil, zero value otherwise.

### GetGenerationCodeOk

`func (o *BaseSegment) GetGenerationCodeOk() (*string, bool)`

GetGenerationCodeOk returns a tuple with the GenerationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationCode

`func (o *BaseSegment) SetGenerationCode(v string)`

SetGenerationCode sets GenerationCode field to given value.

### HasGenerationCode

`func (o *BaseSegment) HasGenerationCode() bool`

HasGenerationCode returns a boolean if a field has been set.

### GetSocialSecurityNumber

`func (o *BaseSegment) GetSocialSecurityNumber() int32`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *BaseSegment) GetSocialSecurityNumberOk() (*int32, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *BaseSegment) SetSocialSecurityNumber(v int32)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.


### GetDateBirth

`func (o *BaseSegment) GetDateBirth() time.Time`

GetDateBirth returns the DateBirth field if non-nil, zero value otherwise.

### GetDateBirthOk

`func (o *BaseSegment) GetDateBirthOk() (*time.Time, bool)`

GetDateBirthOk returns a tuple with the DateBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateBirth

`func (o *BaseSegment) SetDateBirth(v time.Time)`

SetDateBirth sets DateBirth field to given value.


### GetTelephoneNumber

`func (o *BaseSegment) GetTelephoneNumber() int64`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *BaseSegment) GetTelephoneNumberOk() (*int64, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *BaseSegment) SetTelephoneNumber(v int64)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *BaseSegment) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetEcoaCode

`func (o *BaseSegment) GetEcoaCode() string`

GetEcoaCode returns the EcoaCode field if non-nil, zero value otherwise.

### GetEcoaCodeOk

`func (o *BaseSegment) GetEcoaCodeOk() (*string, bool)`

GetEcoaCodeOk returns a tuple with the EcoaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcoaCode

`func (o *BaseSegment) SetEcoaCode(v string)`

SetEcoaCode sets EcoaCode field to given value.


### GetConsumerInformationIndicator

`func (o *BaseSegment) GetConsumerInformationIndicator() string`

GetConsumerInformationIndicator returns the ConsumerInformationIndicator field if non-nil, zero value otherwise.

### GetConsumerInformationIndicatorOk

`func (o *BaseSegment) GetConsumerInformationIndicatorOk() (*string, bool)`

GetConsumerInformationIndicatorOk returns a tuple with the ConsumerInformationIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerInformationIndicator

`func (o *BaseSegment) SetConsumerInformationIndicator(v string)`

SetConsumerInformationIndicator sets ConsumerInformationIndicator field to given value.

### HasConsumerInformationIndicator

`func (o *BaseSegment) HasConsumerInformationIndicator() bool`

HasConsumerInformationIndicator returns a boolean if a field has been set.

### GetCountryCode

`func (o *BaseSegment) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *BaseSegment) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *BaseSegment) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *BaseSegment) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetFirstLineAddress

`func (o *BaseSegment) GetFirstLineAddress() string`

GetFirstLineAddress returns the FirstLineAddress field if non-nil, zero value otherwise.

### GetFirstLineAddressOk

`func (o *BaseSegment) GetFirstLineAddressOk() (*string, bool)`

GetFirstLineAddressOk returns a tuple with the FirstLineAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLineAddress

`func (o *BaseSegment) SetFirstLineAddress(v string)`

SetFirstLineAddress sets FirstLineAddress field to given value.


### GetSecondLineAddress

`func (o *BaseSegment) GetSecondLineAddress() string`

GetSecondLineAddress returns the SecondLineAddress field if non-nil, zero value otherwise.

### GetSecondLineAddressOk

`func (o *BaseSegment) GetSecondLineAddressOk() (*string, bool)`

GetSecondLineAddressOk returns a tuple with the SecondLineAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondLineAddress

`func (o *BaseSegment) SetSecondLineAddress(v string)`

SetSecondLineAddress sets SecondLineAddress field to given value.

### HasSecondLineAddress

`func (o *BaseSegment) HasSecondLineAddress() bool`

HasSecondLineAddress returns a boolean if a field has been set.

### GetCity

`func (o *BaseSegment) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BaseSegment) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BaseSegment) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BaseSegment) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *BaseSegment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BaseSegment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BaseSegment) SetState(v string)`

SetState sets State field to given value.


### GetZipCode

`func (o *BaseSegment) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *BaseSegment) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *BaseSegment) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.


### GetAddressIndicator

`func (o *BaseSegment) GetAddressIndicator() string`

GetAddressIndicator returns the AddressIndicator field if non-nil, zero value otherwise.

### GetAddressIndicatorOk

`func (o *BaseSegment) GetAddressIndicatorOk() (*string, bool)`

GetAddressIndicatorOk returns a tuple with the AddressIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressIndicator

`func (o *BaseSegment) SetAddressIndicator(v string)`

SetAddressIndicator sets AddressIndicator field to given value.

### HasAddressIndicator

`func (o *BaseSegment) HasAddressIndicator() bool`

HasAddressIndicator returns a boolean if a field has been set.

### GetResidenceCode

`func (o *BaseSegment) GetResidenceCode() string`

GetResidenceCode returns the ResidenceCode field if non-nil, zero value otherwise.

### GetResidenceCodeOk

`func (o *BaseSegment) GetResidenceCodeOk() (*string, bool)`

GetResidenceCodeOk returns a tuple with the ResidenceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidenceCode

`func (o *BaseSegment) SetResidenceCode(v string)`

SetResidenceCode sets ResidenceCode field to given value.

### HasResidenceCode

`func (o *BaseSegment) HasResidenceCode() bool`

HasResidenceCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


