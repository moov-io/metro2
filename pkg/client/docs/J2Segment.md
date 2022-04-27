# J2Segment

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
**CountryCode** | Pointer to **string** |  | [optional] 
**FirstLineAddress** | **string** |  | 
**SecondLineAddress** | Pointer to **string** |  | [optional] 
**City** | **string** |  | 
**State** | **string** |  | 
**ZipCode** | **string** |  | 
**AddressIndicator** | Pointer to **string** |  | [optional] 
**ResidenceCode** | Pointer to **string** |  | [optional] 

## Methods

### NewJ2Segment

`func NewJ2Segment(segmentIdentifier string, surname string, firstName string, socialSecurityNumber int32, dateBirth time.Time, ecoaCode string, firstLineAddress string, city string, state string, zipCode string, ) *J2Segment`

NewJ2Segment instantiates a new J2Segment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJ2SegmentWithDefaults

`func NewJ2SegmentWithDefaults() *J2Segment`

NewJ2SegmentWithDefaults instantiates a new J2Segment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentIdentifier

`func (o *J2Segment) GetSegmentIdentifier() string`

GetSegmentIdentifier returns the SegmentIdentifier field if non-nil, zero value otherwise.

### GetSegmentIdentifierOk

`func (o *J2Segment) GetSegmentIdentifierOk() (*string, bool)`

GetSegmentIdentifierOk returns a tuple with the SegmentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIdentifier

`func (o *J2Segment) SetSegmentIdentifier(v string)`

SetSegmentIdentifier sets SegmentIdentifier field to given value.


### GetSurname

`func (o *J2Segment) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *J2Segment) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *J2Segment) SetSurname(v string)`

SetSurname sets Surname field to given value.


### GetFirstName

`func (o *J2Segment) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *J2Segment) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *J2Segment) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *J2Segment) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *J2Segment) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *J2Segment) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *J2Segment) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetGenerationCode

`func (o *J2Segment) GetGenerationCode() string`

GetGenerationCode returns the GenerationCode field if non-nil, zero value otherwise.

### GetGenerationCodeOk

`func (o *J2Segment) GetGenerationCodeOk() (*string, bool)`

GetGenerationCodeOk returns a tuple with the GenerationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationCode

`func (o *J2Segment) SetGenerationCode(v string)`

SetGenerationCode sets GenerationCode field to given value.

### HasGenerationCode

`func (o *J2Segment) HasGenerationCode() bool`

HasGenerationCode returns a boolean if a field has been set.

### GetSocialSecurityNumber

`func (o *J2Segment) GetSocialSecurityNumber() int32`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *J2Segment) GetSocialSecurityNumberOk() (*int32, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *J2Segment) SetSocialSecurityNumber(v int32)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.


### GetDateBirth

`func (o *J2Segment) GetDateBirth() time.Time`

GetDateBirth returns the DateBirth field if non-nil, zero value otherwise.

### GetDateBirthOk

`func (o *J2Segment) GetDateBirthOk() (*time.Time, bool)`

GetDateBirthOk returns a tuple with the DateBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateBirth

`func (o *J2Segment) SetDateBirth(v time.Time)`

SetDateBirth sets DateBirth field to given value.


### GetTelephoneNumber

`func (o *J2Segment) GetTelephoneNumber() int64`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *J2Segment) GetTelephoneNumberOk() (*int64, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *J2Segment) SetTelephoneNumber(v int64)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *J2Segment) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetEcoaCode

`func (o *J2Segment) GetEcoaCode() string`

GetEcoaCode returns the EcoaCode field if non-nil, zero value otherwise.

### GetEcoaCodeOk

`func (o *J2Segment) GetEcoaCodeOk() (*string, bool)`

GetEcoaCodeOk returns a tuple with the EcoaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcoaCode

`func (o *J2Segment) SetEcoaCode(v string)`

SetEcoaCode sets EcoaCode field to given value.


### GetConsumerInformationIndicator

`func (o *J2Segment) GetConsumerInformationIndicator() string`

GetConsumerInformationIndicator returns the ConsumerInformationIndicator field if non-nil, zero value otherwise.

### GetConsumerInformationIndicatorOk

`func (o *J2Segment) GetConsumerInformationIndicatorOk() (*string, bool)`

GetConsumerInformationIndicatorOk returns a tuple with the ConsumerInformationIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerInformationIndicator

`func (o *J2Segment) SetConsumerInformationIndicator(v string)`

SetConsumerInformationIndicator sets ConsumerInformationIndicator field to given value.

### HasConsumerInformationIndicator

`func (o *J2Segment) HasConsumerInformationIndicator() bool`

HasConsumerInformationIndicator returns a boolean if a field has been set.

### GetCountryCode

`func (o *J2Segment) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *J2Segment) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *J2Segment) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *J2Segment) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetFirstLineAddress

`func (o *J2Segment) GetFirstLineAddress() string`

GetFirstLineAddress returns the FirstLineAddress field if non-nil, zero value otherwise.

### GetFirstLineAddressOk

`func (o *J2Segment) GetFirstLineAddressOk() (*string, bool)`

GetFirstLineAddressOk returns a tuple with the FirstLineAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLineAddress

`func (o *J2Segment) SetFirstLineAddress(v string)`

SetFirstLineAddress sets FirstLineAddress field to given value.


### GetSecondLineAddress

`func (o *J2Segment) GetSecondLineAddress() string`

GetSecondLineAddress returns the SecondLineAddress field if non-nil, zero value otherwise.

### GetSecondLineAddressOk

`func (o *J2Segment) GetSecondLineAddressOk() (*string, bool)`

GetSecondLineAddressOk returns a tuple with the SecondLineAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondLineAddress

`func (o *J2Segment) SetSecondLineAddress(v string)`

SetSecondLineAddress sets SecondLineAddress field to given value.

### HasSecondLineAddress

`func (o *J2Segment) HasSecondLineAddress() bool`

HasSecondLineAddress returns a boolean if a field has been set.

### GetCity

`func (o *J2Segment) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *J2Segment) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *J2Segment) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *J2Segment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *J2Segment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *J2Segment) SetState(v string)`

SetState sets State field to given value.


### GetZipCode

`func (o *J2Segment) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *J2Segment) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *J2Segment) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.


### GetAddressIndicator

`func (o *J2Segment) GetAddressIndicator() string`

GetAddressIndicator returns the AddressIndicator field if non-nil, zero value otherwise.

### GetAddressIndicatorOk

`func (o *J2Segment) GetAddressIndicatorOk() (*string, bool)`

GetAddressIndicatorOk returns a tuple with the AddressIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressIndicator

`func (o *J2Segment) SetAddressIndicator(v string)`

SetAddressIndicator sets AddressIndicator field to given value.

### HasAddressIndicator

`func (o *J2Segment) HasAddressIndicator() bool`

HasAddressIndicator returns a boolean if a field has been set.

### GetResidenceCode

`func (o *J2Segment) GetResidenceCode() string`

GetResidenceCode returns the ResidenceCode field if non-nil, zero value otherwise.

### GetResidenceCodeOk

`func (o *J2Segment) GetResidenceCodeOk() (*string, bool)`

GetResidenceCodeOk returns a tuple with the ResidenceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidenceCode

`func (o *J2Segment) SetResidenceCode(v string)`

SetResidenceCode sets ResidenceCode field to given value.

### HasResidenceCode

`func (o *J2Segment) HasResidenceCode() bool`

HasResidenceCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


