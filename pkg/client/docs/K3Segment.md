# K3Segment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentIdentifier** | **string** |  | 
**AgencyIdentifier** | Pointer to **int32** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**MortgageIdentificationNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewK3Segment

`func NewK3Segment(segmentIdentifier string, ) *K3Segment`

NewK3Segment instantiates a new K3Segment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK3SegmentWithDefaults

`func NewK3SegmentWithDefaults() *K3Segment`

NewK3SegmentWithDefaults instantiates a new K3Segment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentIdentifier

`func (o *K3Segment) GetSegmentIdentifier() string`

GetSegmentIdentifier returns the SegmentIdentifier field if non-nil, zero value otherwise.

### GetSegmentIdentifierOk

`func (o *K3Segment) GetSegmentIdentifierOk() (*string, bool)`

GetSegmentIdentifierOk returns a tuple with the SegmentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIdentifier

`func (o *K3Segment) SetSegmentIdentifier(v string)`

SetSegmentIdentifier sets SegmentIdentifier field to given value.


### GetAgencyIdentifier

`func (o *K3Segment) GetAgencyIdentifier() int32`

GetAgencyIdentifier returns the AgencyIdentifier field if non-nil, zero value otherwise.

### GetAgencyIdentifierOk

`func (o *K3Segment) GetAgencyIdentifierOk() (*int32, bool)`

GetAgencyIdentifierOk returns a tuple with the AgencyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyIdentifier

`func (o *K3Segment) SetAgencyIdentifier(v int32)`

SetAgencyIdentifier sets AgencyIdentifier field to given value.

### HasAgencyIdentifier

`func (o *K3Segment) HasAgencyIdentifier() bool`

HasAgencyIdentifier returns a boolean if a field has been set.

### GetAccountNumber

`func (o *K3Segment) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *K3Segment) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *K3Segment) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *K3Segment) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetMortgageIdentificationNumber

`func (o *K3Segment) GetMortgageIdentificationNumber() string`

GetMortgageIdentificationNumber returns the MortgageIdentificationNumber field if non-nil, zero value otherwise.

### GetMortgageIdentificationNumberOk

`func (o *K3Segment) GetMortgageIdentificationNumberOk() (*string, bool)`

GetMortgageIdentificationNumberOk returns a tuple with the MortgageIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMortgageIdentificationNumber

`func (o *K3Segment) SetMortgageIdentificationNumber(v string)`

SetMortgageIdentificationNumber sets MortgageIdentificationNumber field to given value.

### HasMortgageIdentificationNumber

`func (o *K3Segment) HasMortgageIdentificationNumber() bool`

HasMortgageIdentificationNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


