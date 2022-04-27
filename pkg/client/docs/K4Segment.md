# K4Segment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentIdentifier** | **string** |  | 
**SpecializedPaymentIndicator** | **int32** |  | 
**DeferredPaymentStartDate** | Pointer to **time.Time** |  | [optional] 
**BalloonPaymentDueDate** | Pointer to **time.Time** |  | [optional] 
**BalloonPaymentAmount** | Pointer to **int32** |  | [optional] 

## Methods

### NewK4Segment

`func NewK4Segment(segmentIdentifier string, specializedPaymentIndicator int32, ) *K4Segment`

NewK4Segment instantiates a new K4Segment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK4SegmentWithDefaults

`func NewK4SegmentWithDefaults() *K4Segment`

NewK4SegmentWithDefaults instantiates a new K4Segment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentIdentifier

`func (o *K4Segment) GetSegmentIdentifier() string`

GetSegmentIdentifier returns the SegmentIdentifier field if non-nil, zero value otherwise.

### GetSegmentIdentifierOk

`func (o *K4Segment) GetSegmentIdentifierOk() (*string, bool)`

GetSegmentIdentifierOk returns a tuple with the SegmentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIdentifier

`func (o *K4Segment) SetSegmentIdentifier(v string)`

SetSegmentIdentifier sets SegmentIdentifier field to given value.


### GetSpecializedPaymentIndicator

`func (o *K4Segment) GetSpecializedPaymentIndicator() int32`

GetSpecializedPaymentIndicator returns the SpecializedPaymentIndicator field if non-nil, zero value otherwise.

### GetSpecializedPaymentIndicatorOk

`func (o *K4Segment) GetSpecializedPaymentIndicatorOk() (*int32, bool)`

GetSpecializedPaymentIndicatorOk returns a tuple with the SpecializedPaymentIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecializedPaymentIndicator

`func (o *K4Segment) SetSpecializedPaymentIndicator(v int32)`

SetSpecializedPaymentIndicator sets SpecializedPaymentIndicator field to given value.


### GetDeferredPaymentStartDate

`func (o *K4Segment) GetDeferredPaymentStartDate() time.Time`

GetDeferredPaymentStartDate returns the DeferredPaymentStartDate field if non-nil, zero value otherwise.

### GetDeferredPaymentStartDateOk

`func (o *K4Segment) GetDeferredPaymentStartDateOk() (*time.Time, bool)`

GetDeferredPaymentStartDateOk returns a tuple with the DeferredPaymentStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredPaymentStartDate

`func (o *K4Segment) SetDeferredPaymentStartDate(v time.Time)`

SetDeferredPaymentStartDate sets DeferredPaymentStartDate field to given value.

### HasDeferredPaymentStartDate

`func (o *K4Segment) HasDeferredPaymentStartDate() bool`

HasDeferredPaymentStartDate returns a boolean if a field has been set.

### GetBalloonPaymentDueDate

`func (o *K4Segment) GetBalloonPaymentDueDate() time.Time`

GetBalloonPaymentDueDate returns the BalloonPaymentDueDate field if non-nil, zero value otherwise.

### GetBalloonPaymentDueDateOk

`func (o *K4Segment) GetBalloonPaymentDueDateOk() (*time.Time, bool)`

GetBalloonPaymentDueDateOk returns a tuple with the BalloonPaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalloonPaymentDueDate

`func (o *K4Segment) SetBalloonPaymentDueDate(v time.Time)`

SetBalloonPaymentDueDate sets BalloonPaymentDueDate field to given value.

### HasBalloonPaymentDueDate

`func (o *K4Segment) HasBalloonPaymentDueDate() bool`

HasBalloonPaymentDueDate returns a boolean if a field has been set.

### GetBalloonPaymentAmount

`func (o *K4Segment) GetBalloonPaymentAmount() int32`

GetBalloonPaymentAmount returns the BalloonPaymentAmount field if non-nil, zero value otherwise.

### GetBalloonPaymentAmountOk

`func (o *K4Segment) GetBalloonPaymentAmountOk() (*int32, bool)`

GetBalloonPaymentAmountOk returns a tuple with the BalloonPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalloonPaymentAmount

`func (o *K4Segment) SetBalloonPaymentAmount(v int32)`

SetBalloonPaymentAmount sets BalloonPaymentAmount field to given value.

### HasBalloonPaymentAmount

`func (o *K4Segment) HasBalloonPaymentAmount() bool`

HasBalloonPaymentAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


