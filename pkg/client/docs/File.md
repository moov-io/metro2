# File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | [**HeaderRecord**](HeaderRecord.md) |  | 
**Data** | Pointer to [**[]DataRecord**](DataRecord.md) |  | [optional] 
**Trailer** | [**TrailerRecord**](TrailerRecord.md) |  | 

## Methods

### NewFile

`func NewFile(header HeaderRecord, trailer TrailerRecord, ) *File`

NewFile instantiates a new File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileWithDefaults

`func NewFileWithDefaults() *File`

NewFileWithDefaults instantiates a new File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *File) GetHeader() HeaderRecord`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *File) GetHeaderOk() (*HeaderRecord, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *File) SetHeader(v HeaderRecord)`

SetHeader sets Header field to given value.


### GetData

`func (o *File) GetData() []DataRecord`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *File) GetDataOk() (*[]DataRecord, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *File) SetData(v []DataRecord)`

SetData sets Data field to given value.

### HasData

`func (o *File) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTrailer

`func (o *File) GetTrailer() TrailerRecord`

GetTrailer returns the Trailer field if non-nil, zero value otherwise.

### GetTrailerOk

`func (o *File) GetTrailerOk() (*TrailerRecord, bool)`

GetTrailerOk returns a tuple with the Trailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailer

`func (o *File) SetTrailer(v TrailerRecord)`

SetTrailer sets Trailer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


