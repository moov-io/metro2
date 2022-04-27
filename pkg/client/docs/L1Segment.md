# L1Segment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentIdentifier** | **string** |  | 
**ChangeIndicator** | **int32** |  | 
**NewConsumerAccountNumber** | Pointer to **string** |  | [optional] 
**BalloonPaymentDueDate** | Pointer to **string** |  | [optional] 

## Methods

### NewL1Segment

`func NewL1Segment(segmentIdentifier string, changeIndicator int32, ) *L1Segment`

NewL1Segment instantiates a new L1Segment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewL1SegmentWithDefaults

`func NewL1SegmentWithDefaults() *L1Segment`

NewL1SegmentWithDefaults instantiates a new L1Segment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentIdentifier

`func (o *L1Segment) GetSegmentIdentifier() string`

GetSegmentIdentifier returns the SegmentIdentifier field if non-nil, zero value otherwise.

### GetSegmentIdentifierOk

`func (o *L1Segment) GetSegmentIdentifierOk() (*string, bool)`

GetSegmentIdentifierOk returns a tuple with the SegmentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIdentifier

`func (o *L1Segment) SetSegmentIdentifier(v string)`

SetSegmentIdentifier sets SegmentIdentifier field to given value.


### GetChangeIndicator

`func (o *L1Segment) GetChangeIndicator() int32`

GetChangeIndicator returns the ChangeIndicator field if non-nil, zero value otherwise.

### GetChangeIndicatorOk

`func (o *L1Segment) GetChangeIndicatorOk() (*int32, bool)`

GetChangeIndicatorOk returns a tuple with the ChangeIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeIndicator

`func (o *L1Segment) SetChangeIndicator(v int32)`

SetChangeIndicator sets ChangeIndicator field to given value.


### GetNewConsumerAccountNumber

`func (o *L1Segment) GetNewConsumerAccountNumber() string`

GetNewConsumerAccountNumber returns the NewConsumerAccountNumber field if non-nil, zero value otherwise.

### GetNewConsumerAccountNumberOk

`func (o *L1Segment) GetNewConsumerAccountNumberOk() (*string, bool)`

GetNewConsumerAccountNumberOk returns a tuple with the NewConsumerAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewConsumerAccountNumber

`func (o *L1Segment) SetNewConsumerAccountNumber(v string)`

SetNewConsumerAccountNumber sets NewConsumerAccountNumber field to given value.

### HasNewConsumerAccountNumber

`func (o *L1Segment) HasNewConsumerAccountNumber() bool`

HasNewConsumerAccountNumber returns a boolean if a field has been set.

### GetBalloonPaymentDueDate

`func (o *L1Segment) GetBalloonPaymentDueDate() string`

GetBalloonPaymentDueDate returns the BalloonPaymentDueDate field if non-nil, zero value otherwise.

### GetBalloonPaymentDueDateOk

`func (o *L1Segment) GetBalloonPaymentDueDateOk() (*string, bool)`

GetBalloonPaymentDueDateOk returns a tuple with the BalloonPaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalloonPaymentDueDate

`func (o *L1Segment) SetBalloonPaymentDueDate(v string)`

SetBalloonPaymentDueDate sets BalloonPaymentDueDate field to given value.

### HasBalloonPaymentDueDate

`func (o *L1Segment) HasBalloonPaymentDueDate() bool`

HasBalloonPaymentDueDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


