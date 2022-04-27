# J1Segment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentIdentifier** | **string** |  | 
**Surname** | **string** |  | 
**FirstName** | **string** |  | 
**MiddleName** | Pointer to **string** |  | [optional] 
**GenerationCode** | Pointer to **string** |  | [optional] 
**SocialSecurityNumber** | **int32** |  | 
**DateBirth** | **time.Time** |  | 
**TelephoneNumber** | Pointer to **int64** |  | [optional] 
**EcoaCode** | **string** |  | 
**ConsumerInformationIndicator** | Pointer to **string** |  | [optional] 

## Methods

### NewJ1Segment

`func NewJ1Segment(segmentIdentifier string, surname string, firstName string, socialSecurityNumber int32, dateBirth time.Time, ecoaCode string, ) *J1Segment`

NewJ1Segment instantiates a new J1Segment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJ1SegmentWithDefaults

`func NewJ1SegmentWithDefaults() *J1Segment`

NewJ1SegmentWithDefaults instantiates a new J1Segment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentIdentifier

`func (o *J1Segment) GetSegmentIdentifier() string`

GetSegmentIdentifier returns the SegmentIdentifier field if non-nil, zero value otherwise.

### GetSegmentIdentifierOk

`func (o *J1Segment) GetSegmentIdentifierOk() (*string, bool)`

GetSegmentIdentifierOk returns a tuple with the SegmentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIdentifier

`func (o *J1Segment) SetSegmentIdentifier(v string)`

SetSegmentIdentifier sets SegmentIdentifier field to given value.


### GetSurname

`func (o *J1Segment) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *J1Segment) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *J1Segment) SetSurname(v string)`

SetSurname sets Surname field to given value.


### GetFirstName

`func (o *J1Segment) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *J1Segment) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *J1Segment) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *J1Segment) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *J1Segment) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *J1Segment) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *J1Segment) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetGenerationCode

`func (o *J1Segment) GetGenerationCode() string`

GetGenerationCode returns the GenerationCode field if non-nil, zero value otherwise.

### GetGenerationCodeOk

`func (o *J1Segment) GetGenerationCodeOk() (*string, bool)`

GetGenerationCodeOk returns a tuple with the GenerationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationCode

`func (o *J1Segment) SetGenerationCode(v string)`

SetGenerationCode sets GenerationCode field to given value.

### HasGenerationCode

`func (o *J1Segment) HasGenerationCode() bool`

HasGenerationCode returns a boolean if a field has been set.

### GetSocialSecurityNumber

`func (o *J1Segment) GetSocialSecurityNumber() int32`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *J1Segment) GetSocialSecurityNumberOk() (*int32, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *J1Segment) SetSocialSecurityNumber(v int32)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.


### GetDateBirth

`func (o *J1Segment) GetDateBirth() time.Time`

GetDateBirth returns the DateBirth field if non-nil, zero value otherwise.

### GetDateBirthOk

`func (o *J1Segment) GetDateBirthOk() (*time.Time, bool)`

GetDateBirthOk returns a tuple with the DateBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateBirth

`func (o *J1Segment) SetDateBirth(v time.Time)`

SetDateBirth sets DateBirth field to given value.


### GetTelephoneNumber

`func (o *J1Segment) GetTelephoneNumber() int64`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *J1Segment) GetTelephoneNumberOk() (*int64, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *J1Segment) SetTelephoneNumber(v int64)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *J1Segment) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetEcoaCode

`func (o *J1Segment) GetEcoaCode() string`

GetEcoaCode returns the EcoaCode field if non-nil, zero value otherwise.

### GetEcoaCodeOk

`func (o *J1Segment) GetEcoaCodeOk() (*string, bool)`

GetEcoaCodeOk returns a tuple with the EcoaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcoaCode

`func (o *J1Segment) SetEcoaCode(v string)`

SetEcoaCode sets EcoaCode field to given value.


### GetConsumerInformationIndicator

`func (o *J1Segment) GetConsumerInformationIndicator() string`

GetConsumerInformationIndicator returns the ConsumerInformationIndicator field if non-nil, zero value otherwise.

### GetConsumerInformationIndicatorOk

`func (o *J1Segment) GetConsumerInformationIndicatorOk() (*string, bool)`

GetConsumerInformationIndicatorOk returns a tuple with the ConsumerInformationIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerInformationIndicator

`func (o *J1Segment) SetConsumerInformationIndicator(v string)`

SetConsumerInformationIndicator sets ConsumerInformationIndicator field to given value.

### HasConsumerInformationIndicator

`func (o *J1Segment) HasConsumerInformationIndicator() bool`

HasConsumerInformationIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


