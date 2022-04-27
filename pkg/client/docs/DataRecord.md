# DataRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | [**BaseSegment**](BaseSegment.md) |  | 
**J1** | Pointer to [**[]J1Segment**](J1Segment.md) |  | [optional] 
**J2** | Pointer to [**[]J2Segment**](J2Segment.md) |  | [optional] 
**K1** | Pointer to [**K1Segment**](K1Segment.md) |  | [optional] 
**K2** | Pointer to [**K2Segment**](K2Segment.md) |  | [optional] 
**K3** | Pointer to [**K3Segment**](K3Segment.md) |  | [optional] 
**K4** | Pointer to [**K4Segment**](K4Segment.md) |  | [optional] 
**L1** | Pointer to [**L1Segment**](L1Segment.md) |  | [optional] 
**N1** | Pointer to [**N1Segment**](N1Segment.md) |  | [optional] 

## Methods

### NewDataRecord

`func NewDataRecord(base BaseSegment, ) *DataRecord`

NewDataRecord instantiates a new DataRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataRecordWithDefaults

`func NewDataRecordWithDefaults() *DataRecord`

NewDataRecordWithDefaults instantiates a new DataRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *DataRecord) GetBase() BaseSegment`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *DataRecord) GetBaseOk() (*BaseSegment, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *DataRecord) SetBase(v BaseSegment)`

SetBase sets Base field to given value.


### GetJ1

`func (o *DataRecord) GetJ1() []J1Segment`

GetJ1 returns the J1 field if non-nil, zero value otherwise.

### GetJ1Ok

`func (o *DataRecord) GetJ1Ok() (*[]J1Segment, bool)`

GetJ1Ok returns a tuple with the J1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJ1

`func (o *DataRecord) SetJ1(v []J1Segment)`

SetJ1 sets J1 field to given value.

### HasJ1

`func (o *DataRecord) HasJ1() bool`

HasJ1 returns a boolean if a field has been set.

### GetJ2

`func (o *DataRecord) GetJ2() []J2Segment`

GetJ2 returns the J2 field if non-nil, zero value otherwise.

### GetJ2Ok

`func (o *DataRecord) GetJ2Ok() (*[]J2Segment, bool)`

GetJ2Ok returns a tuple with the J2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJ2

`func (o *DataRecord) SetJ2(v []J2Segment)`

SetJ2 sets J2 field to given value.

### HasJ2

`func (o *DataRecord) HasJ2() bool`

HasJ2 returns a boolean if a field has been set.

### GetK1

`func (o *DataRecord) GetK1() K1Segment`

GetK1 returns the K1 field if non-nil, zero value otherwise.

### GetK1Ok

`func (o *DataRecord) GetK1Ok() (*K1Segment, bool)`

GetK1Ok returns a tuple with the K1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK1

`func (o *DataRecord) SetK1(v K1Segment)`

SetK1 sets K1 field to given value.

### HasK1

`func (o *DataRecord) HasK1() bool`

HasK1 returns a boolean if a field has been set.

### GetK2

`func (o *DataRecord) GetK2() K2Segment`

GetK2 returns the K2 field if non-nil, zero value otherwise.

### GetK2Ok

`func (o *DataRecord) GetK2Ok() (*K2Segment, bool)`

GetK2Ok returns a tuple with the K2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK2

`func (o *DataRecord) SetK2(v K2Segment)`

SetK2 sets K2 field to given value.

### HasK2

`func (o *DataRecord) HasK2() bool`

HasK2 returns a boolean if a field has been set.

### GetK3

`func (o *DataRecord) GetK3() K3Segment`

GetK3 returns the K3 field if non-nil, zero value otherwise.

### GetK3Ok

`func (o *DataRecord) GetK3Ok() (*K3Segment, bool)`

GetK3Ok returns a tuple with the K3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK3

`func (o *DataRecord) SetK3(v K3Segment)`

SetK3 sets K3 field to given value.

### HasK3

`func (o *DataRecord) HasK3() bool`

HasK3 returns a boolean if a field has been set.

### GetK4

`func (o *DataRecord) GetK4() K4Segment`

GetK4 returns the K4 field if non-nil, zero value otherwise.

### GetK4Ok

`func (o *DataRecord) GetK4Ok() (*K4Segment, bool)`

GetK4Ok returns a tuple with the K4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK4

`func (o *DataRecord) SetK4(v K4Segment)`

SetK4 sets K4 field to given value.

### HasK4

`func (o *DataRecord) HasK4() bool`

HasK4 returns a boolean if a field has been set.

### GetL1

`func (o *DataRecord) GetL1() L1Segment`

GetL1 returns the L1 field if non-nil, zero value otherwise.

### GetL1Ok

`func (o *DataRecord) GetL1Ok() (*L1Segment, bool)`

GetL1Ok returns a tuple with the L1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL1

`func (o *DataRecord) SetL1(v L1Segment)`

SetL1 sets L1 field to given value.

### HasL1

`func (o *DataRecord) HasL1() bool`

HasL1 returns a boolean if a field has been set.

### GetN1

`func (o *DataRecord) GetN1() N1Segment`

GetN1 returns the N1 field if non-nil, zero value otherwise.

### GetN1Ok

`func (o *DataRecord) GetN1Ok() (*N1Segment, bool)`

GetN1Ok returns a tuple with the N1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1

`func (o *DataRecord) SetN1(v N1Segment)`

SetN1 sets N1 field to given value.

### HasN1

`func (o *DataRecord) HasN1() bool`

HasN1 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


