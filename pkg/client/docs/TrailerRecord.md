# TrailerRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockDescriptorWord** | Pointer to **int32** |  | [optional] 
**RecordDescriptorWord** | **int32** |  | 
**RecordIdentifier** | **string** |  | 
**TotalBaseRecords** | Pointer to **int32** |  | [optional] 
**TotalStatusCodeDF** | Pointer to **int32** |  | [optional] 
**TotalConsumerSegmentsJ1** | Pointer to **int32** |  | [optional] 
**TotalConsumerSegmentsJ2** | Pointer to **int32** |  | [optional] 
**BlockCount** | Pointer to **int32** |  | [optional] 
**TotalStatusCodeDA** | Pointer to **int32** |  | [optional] 
**TotalStatusCode05** | Pointer to **int32** |  | [optional] 
**TotalStatusCode11** | Pointer to **int32** |  | [optional] 
**TotalStatusCode13** | Pointer to **int32** |  | [optional] 
**TotalStatusCode61** | Pointer to **int32** |  | [optional] 
**TotalStatusCode62** | Pointer to **int32** |  | [optional] 
**TotalStatusCode63** | Pointer to **int32** |  | [optional] 
**TotalStatusCode64** | Pointer to **int32** |  | [optional] 
**TotalStatusCode65** | Pointer to **int32** |  | [optional] 
**TotalStatusCode71** | Pointer to **int32** |  | [optional] 
**TotalStatusCode78** | Pointer to **int32** |  | [optional] 
**TotalStatusCode80** | Pointer to **int32** |  | [optional] 
**TotalStatusCode82** | Pointer to **int32** |  | [optional] 
**TotalStatusCode83** | Pointer to **int32** |  | [optional] 
**TotalStatusCode84** | Pointer to **int32** |  | [optional] 
**TotalStatusCode88** | Pointer to **int32** |  | [optional] 
**TotalStatusCode89** | Pointer to **int32** |  | [optional] 
**TotalStatusCode93** | Pointer to **int32** |  | [optional] 
**TotalStatusCode94** | Pointer to **int32** |  | [optional] 
**TotalStatusCode95** | Pointer to **int32** |  | [optional] 
**TotalStatusCode96** | Pointer to **int32** |  | [optional] 
**TotalStatusCode97** | Pointer to **int32** |  | [optional] 
**TotalECOACodeZ** | Pointer to **int32** |  | [optional] 
**TotalEmploymentSegments** | Pointer to **int32** |  | [optional] 
**TotalOriginalCreditorSegments** | Pointer to **int32** |  | [optional] 
**TotalPurchasedToSegments** | Pointer to **int32** |  | [optional] 
**TotalMortgageInformationSegments** | Pointer to **int32** |  | [optional] 
**TotalPaymentInformationSegments** | Pointer to **int32** |  | [optional] 
**TotalChangeSegments** | Pointer to **int32** |  | [optional] 
**TotalSocialNumbersAllSegments** | Pointer to **int32** |  | [optional] 
**TotalSocialNumbersBaseSegments** | Pointer to **int32** |  | [optional] 
**TotalSocialNumbersJ1Segments** | Pointer to **int32** |  | [optional] 
**TotalSocialNumbersJ2Segments** | Pointer to **int32** |  | [optional] 
**TotalDatesBirthAllSegments** | Pointer to **int32** |  | [optional] 
**TotalDatesBirthBaseSegments** | Pointer to **int32** |  | [optional] 
**TotalDatesBirthJ1Segments** | Pointer to **int32** |  | [optional] 
**TotalDatesBirthJ2Segments** | Pointer to **int32** |  | [optional] 
**TotalTelephoneNumbersAllSegments** | Pointer to **int32** |  | [optional] 

## Methods

### NewTrailerRecord

`func NewTrailerRecord(recordDescriptorWord int32, recordIdentifier string, ) *TrailerRecord`

NewTrailerRecord instantiates a new TrailerRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailerRecordWithDefaults

`func NewTrailerRecordWithDefaults() *TrailerRecord`

NewTrailerRecordWithDefaults instantiates a new TrailerRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockDescriptorWord

`func (o *TrailerRecord) GetBlockDescriptorWord() int32`

GetBlockDescriptorWord returns the BlockDescriptorWord field if non-nil, zero value otherwise.

### GetBlockDescriptorWordOk

`func (o *TrailerRecord) GetBlockDescriptorWordOk() (*int32, bool)`

GetBlockDescriptorWordOk returns a tuple with the BlockDescriptorWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDescriptorWord

`func (o *TrailerRecord) SetBlockDescriptorWord(v int32)`

SetBlockDescriptorWord sets BlockDescriptorWord field to given value.

### HasBlockDescriptorWord

`func (o *TrailerRecord) HasBlockDescriptorWord() bool`

HasBlockDescriptorWord returns a boolean if a field has been set.

### GetRecordDescriptorWord

`func (o *TrailerRecord) GetRecordDescriptorWord() int32`

GetRecordDescriptorWord returns the RecordDescriptorWord field if non-nil, zero value otherwise.

### GetRecordDescriptorWordOk

`func (o *TrailerRecord) GetRecordDescriptorWordOk() (*int32, bool)`

GetRecordDescriptorWordOk returns a tuple with the RecordDescriptorWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordDescriptorWord

`func (o *TrailerRecord) SetRecordDescriptorWord(v int32)`

SetRecordDescriptorWord sets RecordDescriptorWord field to given value.


### GetRecordIdentifier

`func (o *TrailerRecord) GetRecordIdentifier() string`

GetRecordIdentifier returns the RecordIdentifier field if non-nil, zero value otherwise.

### GetRecordIdentifierOk

`func (o *TrailerRecord) GetRecordIdentifierOk() (*string, bool)`

GetRecordIdentifierOk returns a tuple with the RecordIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordIdentifier

`func (o *TrailerRecord) SetRecordIdentifier(v string)`

SetRecordIdentifier sets RecordIdentifier field to given value.


### GetTotalBaseRecords

`func (o *TrailerRecord) GetTotalBaseRecords() int32`

GetTotalBaseRecords returns the TotalBaseRecords field if non-nil, zero value otherwise.

### GetTotalBaseRecordsOk

`func (o *TrailerRecord) GetTotalBaseRecordsOk() (*int32, bool)`

GetTotalBaseRecordsOk returns a tuple with the TotalBaseRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBaseRecords

`func (o *TrailerRecord) SetTotalBaseRecords(v int32)`

SetTotalBaseRecords sets TotalBaseRecords field to given value.

### HasTotalBaseRecords

`func (o *TrailerRecord) HasTotalBaseRecords() bool`

HasTotalBaseRecords returns a boolean if a field has been set.

### GetTotalStatusCodeDF

`func (o *TrailerRecord) GetTotalStatusCodeDF() int32`

GetTotalStatusCodeDF returns the TotalStatusCodeDF field if non-nil, zero value otherwise.

### GetTotalStatusCodeDFOk

`func (o *TrailerRecord) GetTotalStatusCodeDFOk() (*int32, bool)`

GetTotalStatusCodeDFOk returns a tuple with the TotalStatusCodeDF field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCodeDF

`func (o *TrailerRecord) SetTotalStatusCodeDF(v int32)`

SetTotalStatusCodeDF sets TotalStatusCodeDF field to given value.

### HasTotalStatusCodeDF

`func (o *TrailerRecord) HasTotalStatusCodeDF() bool`

HasTotalStatusCodeDF returns a boolean if a field has been set.

### GetTotalConsumerSegmentsJ1

`func (o *TrailerRecord) GetTotalConsumerSegmentsJ1() int32`

GetTotalConsumerSegmentsJ1 returns the TotalConsumerSegmentsJ1 field if non-nil, zero value otherwise.

### GetTotalConsumerSegmentsJ1Ok

`func (o *TrailerRecord) GetTotalConsumerSegmentsJ1Ok() (*int32, bool)`

GetTotalConsumerSegmentsJ1Ok returns a tuple with the TotalConsumerSegmentsJ1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalConsumerSegmentsJ1

`func (o *TrailerRecord) SetTotalConsumerSegmentsJ1(v int32)`

SetTotalConsumerSegmentsJ1 sets TotalConsumerSegmentsJ1 field to given value.

### HasTotalConsumerSegmentsJ1

`func (o *TrailerRecord) HasTotalConsumerSegmentsJ1() bool`

HasTotalConsumerSegmentsJ1 returns a boolean if a field has been set.

### GetTotalConsumerSegmentsJ2

`func (o *TrailerRecord) GetTotalConsumerSegmentsJ2() int32`

GetTotalConsumerSegmentsJ2 returns the TotalConsumerSegmentsJ2 field if non-nil, zero value otherwise.

### GetTotalConsumerSegmentsJ2Ok

`func (o *TrailerRecord) GetTotalConsumerSegmentsJ2Ok() (*int32, bool)`

GetTotalConsumerSegmentsJ2Ok returns a tuple with the TotalConsumerSegmentsJ2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalConsumerSegmentsJ2

`func (o *TrailerRecord) SetTotalConsumerSegmentsJ2(v int32)`

SetTotalConsumerSegmentsJ2 sets TotalConsumerSegmentsJ2 field to given value.

### HasTotalConsumerSegmentsJ2

`func (o *TrailerRecord) HasTotalConsumerSegmentsJ2() bool`

HasTotalConsumerSegmentsJ2 returns a boolean if a field has been set.

### GetBlockCount

`func (o *TrailerRecord) GetBlockCount() int32`

GetBlockCount returns the BlockCount field if non-nil, zero value otherwise.

### GetBlockCountOk

`func (o *TrailerRecord) GetBlockCountOk() (*int32, bool)`

GetBlockCountOk returns a tuple with the BlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCount

`func (o *TrailerRecord) SetBlockCount(v int32)`

SetBlockCount sets BlockCount field to given value.

### HasBlockCount

`func (o *TrailerRecord) HasBlockCount() bool`

HasBlockCount returns a boolean if a field has been set.

### GetTotalStatusCodeDA

`func (o *TrailerRecord) GetTotalStatusCodeDA() int32`

GetTotalStatusCodeDA returns the TotalStatusCodeDA field if non-nil, zero value otherwise.

### GetTotalStatusCodeDAOk

`func (o *TrailerRecord) GetTotalStatusCodeDAOk() (*int32, bool)`

GetTotalStatusCodeDAOk returns a tuple with the TotalStatusCodeDA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCodeDA

`func (o *TrailerRecord) SetTotalStatusCodeDA(v int32)`

SetTotalStatusCodeDA sets TotalStatusCodeDA field to given value.

### HasTotalStatusCodeDA

`func (o *TrailerRecord) HasTotalStatusCodeDA() bool`

HasTotalStatusCodeDA returns a boolean if a field has been set.

### GetTotalStatusCode05

`func (o *TrailerRecord) GetTotalStatusCode05() int32`

GetTotalStatusCode05 returns the TotalStatusCode05 field if non-nil, zero value otherwise.

### GetTotalStatusCode05Ok

`func (o *TrailerRecord) GetTotalStatusCode05Ok() (*int32, bool)`

GetTotalStatusCode05Ok returns a tuple with the TotalStatusCode05 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode05

`func (o *TrailerRecord) SetTotalStatusCode05(v int32)`

SetTotalStatusCode05 sets TotalStatusCode05 field to given value.

### HasTotalStatusCode05

`func (o *TrailerRecord) HasTotalStatusCode05() bool`

HasTotalStatusCode05 returns a boolean if a field has been set.

### GetTotalStatusCode11

`func (o *TrailerRecord) GetTotalStatusCode11() int32`

GetTotalStatusCode11 returns the TotalStatusCode11 field if non-nil, zero value otherwise.

### GetTotalStatusCode11Ok

`func (o *TrailerRecord) GetTotalStatusCode11Ok() (*int32, bool)`

GetTotalStatusCode11Ok returns a tuple with the TotalStatusCode11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode11

`func (o *TrailerRecord) SetTotalStatusCode11(v int32)`

SetTotalStatusCode11 sets TotalStatusCode11 field to given value.

### HasTotalStatusCode11

`func (o *TrailerRecord) HasTotalStatusCode11() bool`

HasTotalStatusCode11 returns a boolean if a field has been set.

### GetTotalStatusCode13

`func (o *TrailerRecord) GetTotalStatusCode13() int32`

GetTotalStatusCode13 returns the TotalStatusCode13 field if non-nil, zero value otherwise.

### GetTotalStatusCode13Ok

`func (o *TrailerRecord) GetTotalStatusCode13Ok() (*int32, bool)`

GetTotalStatusCode13Ok returns a tuple with the TotalStatusCode13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode13

`func (o *TrailerRecord) SetTotalStatusCode13(v int32)`

SetTotalStatusCode13 sets TotalStatusCode13 field to given value.

### HasTotalStatusCode13

`func (o *TrailerRecord) HasTotalStatusCode13() bool`

HasTotalStatusCode13 returns a boolean if a field has been set.

### GetTotalStatusCode61

`func (o *TrailerRecord) GetTotalStatusCode61() int32`

GetTotalStatusCode61 returns the TotalStatusCode61 field if non-nil, zero value otherwise.

### GetTotalStatusCode61Ok

`func (o *TrailerRecord) GetTotalStatusCode61Ok() (*int32, bool)`

GetTotalStatusCode61Ok returns a tuple with the TotalStatusCode61 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode61

`func (o *TrailerRecord) SetTotalStatusCode61(v int32)`

SetTotalStatusCode61 sets TotalStatusCode61 field to given value.

### HasTotalStatusCode61

`func (o *TrailerRecord) HasTotalStatusCode61() bool`

HasTotalStatusCode61 returns a boolean if a field has been set.

### GetTotalStatusCode62

`func (o *TrailerRecord) GetTotalStatusCode62() int32`

GetTotalStatusCode62 returns the TotalStatusCode62 field if non-nil, zero value otherwise.

### GetTotalStatusCode62Ok

`func (o *TrailerRecord) GetTotalStatusCode62Ok() (*int32, bool)`

GetTotalStatusCode62Ok returns a tuple with the TotalStatusCode62 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode62

`func (o *TrailerRecord) SetTotalStatusCode62(v int32)`

SetTotalStatusCode62 sets TotalStatusCode62 field to given value.

### HasTotalStatusCode62

`func (o *TrailerRecord) HasTotalStatusCode62() bool`

HasTotalStatusCode62 returns a boolean if a field has been set.

### GetTotalStatusCode63

`func (o *TrailerRecord) GetTotalStatusCode63() int32`

GetTotalStatusCode63 returns the TotalStatusCode63 field if non-nil, zero value otherwise.

### GetTotalStatusCode63Ok

`func (o *TrailerRecord) GetTotalStatusCode63Ok() (*int32, bool)`

GetTotalStatusCode63Ok returns a tuple with the TotalStatusCode63 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode63

`func (o *TrailerRecord) SetTotalStatusCode63(v int32)`

SetTotalStatusCode63 sets TotalStatusCode63 field to given value.

### HasTotalStatusCode63

`func (o *TrailerRecord) HasTotalStatusCode63() bool`

HasTotalStatusCode63 returns a boolean if a field has been set.

### GetTotalStatusCode64

`func (o *TrailerRecord) GetTotalStatusCode64() int32`

GetTotalStatusCode64 returns the TotalStatusCode64 field if non-nil, zero value otherwise.

### GetTotalStatusCode64Ok

`func (o *TrailerRecord) GetTotalStatusCode64Ok() (*int32, bool)`

GetTotalStatusCode64Ok returns a tuple with the TotalStatusCode64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode64

`func (o *TrailerRecord) SetTotalStatusCode64(v int32)`

SetTotalStatusCode64 sets TotalStatusCode64 field to given value.

### HasTotalStatusCode64

`func (o *TrailerRecord) HasTotalStatusCode64() bool`

HasTotalStatusCode64 returns a boolean if a field has been set.

### GetTotalStatusCode65

`func (o *TrailerRecord) GetTotalStatusCode65() int32`

GetTotalStatusCode65 returns the TotalStatusCode65 field if non-nil, zero value otherwise.

### GetTotalStatusCode65Ok

`func (o *TrailerRecord) GetTotalStatusCode65Ok() (*int32, bool)`

GetTotalStatusCode65Ok returns a tuple with the TotalStatusCode65 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode65

`func (o *TrailerRecord) SetTotalStatusCode65(v int32)`

SetTotalStatusCode65 sets TotalStatusCode65 field to given value.

### HasTotalStatusCode65

`func (o *TrailerRecord) HasTotalStatusCode65() bool`

HasTotalStatusCode65 returns a boolean if a field has been set.

### GetTotalStatusCode71

`func (o *TrailerRecord) GetTotalStatusCode71() int32`

GetTotalStatusCode71 returns the TotalStatusCode71 field if non-nil, zero value otherwise.

### GetTotalStatusCode71Ok

`func (o *TrailerRecord) GetTotalStatusCode71Ok() (*int32, bool)`

GetTotalStatusCode71Ok returns a tuple with the TotalStatusCode71 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode71

`func (o *TrailerRecord) SetTotalStatusCode71(v int32)`

SetTotalStatusCode71 sets TotalStatusCode71 field to given value.

### HasTotalStatusCode71

`func (o *TrailerRecord) HasTotalStatusCode71() bool`

HasTotalStatusCode71 returns a boolean if a field has been set.

### GetTotalStatusCode78

`func (o *TrailerRecord) GetTotalStatusCode78() int32`

GetTotalStatusCode78 returns the TotalStatusCode78 field if non-nil, zero value otherwise.

### GetTotalStatusCode78Ok

`func (o *TrailerRecord) GetTotalStatusCode78Ok() (*int32, bool)`

GetTotalStatusCode78Ok returns a tuple with the TotalStatusCode78 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode78

`func (o *TrailerRecord) SetTotalStatusCode78(v int32)`

SetTotalStatusCode78 sets TotalStatusCode78 field to given value.

### HasTotalStatusCode78

`func (o *TrailerRecord) HasTotalStatusCode78() bool`

HasTotalStatusCode78 returns a boolean if a field has been set.

### GetTotalStatusCode80

`func (o *TrailerRecord) GetTotalStatusCode80() int32`

GetTotalStatusCode80 returns the TotalStatusCode80 field if non-nil, zero value otherwise.

### GetTotalStatusCode80Ok

`func (o *TrailerRecord) GetTotalStatusCode80Ok() (*int32, bool)`

GetTotalStatusCode80Ok returns a tuple with the TotalStatusCode80 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode80

`func (o *TrailerRecord) SetTotalStatusCode80(v int32)`

SetTotalStatusCode80 sets TotalStatusCode80 field to given value.

### HasTotalStatusCode80

`func (o *TrailerRecord) HasTotalStatusCode80() bool`

HasTotalStatusCode80 returns a boolean if a field has been set.

### GetTotalStatusCode82

`func (o *TrailerRecord) GetTotalStatusCode82() int32`

GetTotalStatusCode82 returns the TotalStatusCode82 field if non-nil, zero value otherwise.

### GetTotalStatusCode82Ok

`func (o *TrailerRecord) GetTotalStatusCode82Ok() (*int32, bool)`

GetTotalStatusCode82Ok returns a tuple with the TotalStatusCode82 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode82

`func (o *TrailerRecord) SetTotalStatusCode82(v int32)`

SetTotalStatusCode82 sets TotalStatusCode82 field to given value.

### HasTotalStatusCode82

`func (o *TrailerRecord) HasTotalStatusCode82() bool`

HasTotalStatusCode82 returns a boolean if a field has been set.

### GetTotalStatusCode83

`func (o *TrailerRecord) GetTotalStatusCode83() int32`

GetTotalStatusCode83 returns the TotalStatusCode83 field if non-nil, zero value otherwise.

### GetTotalStatusCode83Ok

`func (o *TrailerRecord) GetTotalStatusCode83Ok() (*int32, bool)`

GetTotalStatusCode83Ok returns a tuple with the TotalStatusCode83 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode83

`func (o *TrailerRecord) SetTotalStatusCode83(v int32)`

SetTotalStatusCode83 sets TotalStatusCode83 field to given value.

### HasTotalStatusCode83

`func (o *TrailerRecord) HasTotalStatusCode83() bool`

HasTotalStatusCode83 returns a boolean if a field has been set.

### GetTotalStatusCode84

`func (o *TrailerRecord) GetTotalStatusCode84() int32`

GetTotalStatusCode84 returns the TotalStatusCode84 field if non-nil, zero value otherwise.

### GetTotalStatusCode84Ok

`func (o *TrailerRecord) GetTotalStatusCode84Ok() (*int32, bool)`

GetTotalStatusCode84Ok returns a tuple with the TotalStatusCode84 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode84

`func (o *TrailerRecord) SetTotalStatusCode84(v int32)`

SetTotalStatusCode84 sets TotalStatusCode84 field to given value.

### HasTotalStatusCode84

`func (o *TrailerRecord) HasTotalStatusCode84() bool`

HasTotalStatusCode84 returns a boolean if a field has been set.

### GetTotalStatusCode88

`func (o *TrailerRecord) GetTotalStatusCode88() int32`

GetTotalStatusCode88 returns the TotalStatusCode88 field if non-nil, zero value otherwise.

### GetTotalStatusCode88Ok

`func (o *TrailerRecord) GetTotalStatusCode88Ok() (*int32, bool)`

GetTotalStatusCode88Ok returns a tuple with the TotalStatusCode88 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode88

`func (o *TrailerRecord) SetTotalStatusCode88(v int32)`

SetTotalStatusCode88 sets TotalStatusCode88 field to given value.

### HasTotalStatusCode88

`func (o *TrailerRecord) HasTotalStatusCode88() bool`

HasTotalStatusCode88 returns a boolean if a field has been set.

### GetTotalStatusCode89

`func (o *TrailerRecord) GetTotalStatusCode89() int32`

GetTotalStatusCode89 returns the TotalStatusCode89 field if non-nil, zero value otherwise.

### GetTotalStatusCode89Ok

`func (o *TrailerRecord) GetTotalStatusCode89Ok() (*int32, bool)`

GetTotalStatusCode89Ok returns a tuple with the TotalStatusCode89 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode89

`func (o *TrailerRecord) SetTotalStatusCode89(v int32)`

SetTotalStatusCode89 sets TotalStatusCode89 field to given value.

### HasTotalStatusCode89

`func (o *TrailerRecord) HasTotalStatusCode89() bool`

HasTotalStatusCode89 returns a boolean if a field has been set.

### GetTotalStatusCode93

`func (o *TrailerRecord) GetTotalStatusCode93() int32`

GetTotalStatusCode93 returns the TotalStatusCode93 field if non-nil, zero value otherwise.

### GetTotalStatusCode93Ok

`func (o *TrailerRecord) GetTotalStatusCode93Ok() (*int32, bool)`

GetTotalStatusCode93Ok returns a tuple with the TotalStatusCode93 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode93

`func (o *TrailerRecord) SetTotalStatusCode93(v int32)`

SetTotalStatusCode93 sets TotalStatusCode93 field to given value.

### HasTotalStatusCode93

`func (o *TrailerRecord) HasTotalStatusCode93() bool`

HasTotalStatusCode93 returns a boolean if a field has been set.

### GetTotalStatusCode94

`func (o *TrailerRecord) GetTotalStatusCode94() int32`

GetTotalStatusCode94 returns the TotalStatusCode94 field if non-nil, zero value otherwise.

### GetTotalStatusCode94Ok

`func (o *TrailerRecord) GetTotalStatusCode94Ok() (*int32, bool)`

GetTotalStatusCode94Ok returns a tuple with the TotalStatusCode94 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode94

`func (o *TrailerRecord) SetTotalStatusCode94(v int32)`

SetTotalStatusCode94 sets TotalStatusCode94 field to given value.

### HasTotalStatusCode94

`func (o *TrailerRecord) HasTotalStatusCode94() bool`

HasTotalStatusCode94 returns a boolean if a field has been set.

### GetTotalStatusCode95

`func (o *TrailerRecord) GetTotalStatusCode95() int32`

GetTotalStatusCode95 returns the TotalStatusCode95 field if non-nil, zero value otherwise.

### GetTotalStatusCode95Ok

`func (o *TrailerRecord) GetTotalStatusCode95Ok() (*int32, bool)`

GetTotalStatusCode95Ok returns a tuple with the TotalStatusCode95 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode95

`func (o *TrailerRecord) SetTotalStatusCode95(v int32)`

SetTotalStatusCode95 sets TotalStatusCode95 field to given value.

### HasTotalStatusCode95

`func (o *TrailerRecord) HasTotalStatusCode95() bool`

HasTotalStatusCode95 returns a boolean if a field has been set.

### GetTotalStatusCode96

`func (o *TrailerRecord) GetTotalStatusCode96() int32`

GetTotalStatusCode96 returns the TotalStatusCode96 field if non-nil, zero value otherwise.

### GetTotalStatusCode96Ok

`func (o *TrailerRecord) GetTotalStatusCode96Ok() (*int32, bool)`

GetTotalStatusCode96Ok returns a tuple with the TotalStatusCode96 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode96

`func (o *TrailerRecord) SetTotalStatusCode96(v int32)`

SetTotalStatusCode96 sets TotalStatusCode96 field to given value.

### HasTotalStatusCode96

`func (o *TrailerRecord) HasTotalStatusCode96() bool`

HasTotalStatusCode96 returns a boolean if a field has been set.

### GetTotalStatusCode97

`func (o *TrailerRecord) GetTotalStatusCode97() int32`

GetTotalStatusCode97 returns the TotalStatusCode97 field if non-nil, zero value otherwise.

### GetTotalStatusCode97Ok

`func (o *TrailerRecord) GetTotalStatusCode97Ok() (*int32, bool)`

GetTotalStatusCode97Ok returns a tuple with the TotalStatusCode97 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStatusCode97

`func (o *TrailerRecord) SetTotalStatusCode97(v int32)`

SetTotalStatusCode97 sets TotalStatusCode97 field to given value.

### HasTotalStatusCode97

`func (o *TrailerRecord) HasTotalStatusCode97() bool`

HasTotalStatusCode97 returns a boolean if a field has been set.

### GetTotalECOACodeZ

`func (o *TrailerRecord) GetTotalECOACodeZ() int32`

GetTotalECOACodeZ returns the TotalECOACodeZ field if non-nil, zero value otherwise.

### GetTotalECOACodeZOk

`func (o *TrailerRecord) GetTotalECOACodeZOk() (*int32, bool)`

GetTotalECOACodeZOk returns a tuple with the TotalECOACodeZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalECOACodeZ

`func (o *TrailerRecord) SetTotalECOACodeZ(v int32)`

SetTotalECOACodeZ sets TotalECOACodeZ field to given value.

### HasTotalECOACodeZ

`func (o *TrailerRecord) HasTotalECOACodeZ() bool`

HasTotalECOACodeZ returns a boolean if a field has been set.

### GetTotalEmploymentSegments

`func (o *TrailerRecord) GetTotalEmploymentSegments() int32`

GetTotalEmploymentSegments returns the TotalEmploymentSegments field if non-nil, zero value otherwise.

### GetTotalEmploymentSegmentsOk

`func (o *TrailerRecord) GetTotalEmploymentSegmentsOk() (*int32, bool)`

GetTotalEmploymentSegmentsOk returns a tuple with the TotalEmploymentSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEmploymentSegments

`func (o *TrailerRecord) SetTotalEmploymentSegments(v int32)`

SetTotalEmploymentSegments sets TotalEmploymentSegments field to given value.

### HasTotalEmploymentSegments

`func (o *TrailerRecord) HasTotalEmploymentSegments() bool`

HasTotalEmploymentSegments returns a boolean if a field has been set.

### GetTotalOriginalCreditorSegments

`func (o *TrailerRecord) GetTotalOriginalCreditorSegments() int32`

GetTotalOriginalCreditorSegments returns the TotalOriginalCreditorSegments field if non-nil, zero value otherwise.

### GetTotalOriginalCreditorSegmentsOk

`func (o *TrailerRecord) GetTotalOriginalCreditorSegmentsOk() (*int32, bool)`

GetTotalOriginalCreditorSegmentsOk returns a tuple with the TotalOriginalCreditorSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOriginalCreditorSegments

`func (o *TrailerRecord) SetTotalOriginalCreditorSegments(v int32)`

SetTotalOriginalCreditorSegments sets TotalOriginalCreditorSegments field to given value.

### HasTotalOriginalCreditorSegments

`func (o *TrailerRecord) HasTotalOriginalCreditorSegments() bool`

HasTotalOriginalCreditorSegments returns a boolean if a field has been set.

### GetTotalPurchasedToSegments

`func (o *TrailerRecord) GetTotalPurchasedToSegments() int32`

GetTotalPurchasedToSegments returns the TotalPurchasedToSegments field if non-nil, zero value otherwise.

### GetTotalPurchasedToSegmentsOk

`func (o *TrailerRecord) GetTotalPurchasedToSegmentsOk() (*int32, bool)`

GetTotalPurchasedToSegmentsOk returns a tuple with the TotalPurchasedToSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPurchasedToSegments

`func (o *TrailerRecord) SetTotalPurchasedToSegments(v int32)`

SetTotalPurchasedToSegments sets TotalPurchasedToSegments field to given value.

### HasTotalPurchasedToSegments

`func (o *TrailerRecord) HasTotalPurchasedToSegments() bool`

HasTotalPurchasedToSegments returns a boolean if a field has been set.

### GetTotalMortgageInformationSegments

`func (o *TrailerRecord) GetTotalMortgageInformationSegments() int32`

GetTotalMortgageInformationSegments returns the TotalMortgageInformationSegments field if non-nil, zero value otherwise.

### GetTotalMortgageInformationSegmentsOk

`func (o *TrailerRecord) GetTotalMortgageInformationSegmentsOk() (*int32, bool)`

GetTotalMortgageInformationSegmentsOk returns a tuple with the TotalMortgageInformationSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMortgageInformationSegments

`func (o *TrailerRecord) SetTotalMortgageInformationSegments(v int32)`

SetTotalMortgageInformationSegments sets TotalMortgageInformationSegments field to given value.

### HasTotalMortgageInformationSegments

`func (o *TrailerRecord) HasTotalMortgageInformationSegments() bool`

HasTotalMortgageInformationSegments returns a boolean if a field has been set.

### GetTotalPaymentInformationSegments

`func (o *TrailerRecord) GetTotalPaymentInformationSegments() int32`

GetTotalPaymentInformationSegments returns the TotalPaymentInformationSegments field if non-nil, zero value otherwise.

### GetTotalPaymentInformationSegmentsOk

`func (o *TrailerRecord) GetTotalPaymentInformationSegmentsOk() (*int32, bool)`

GetTotalPaymentInformationSegmentsOk returns a tuple with the TotalPaymentInformationSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaymentInformationSegments

`func (o *TrailerRecord) SetTotalPaymentInformationSegments(v int32)`

SetTotalPaymentInformationSegments sets TotalPaymentInformationSegments field to given value.

### HasTotalPaymentInformationSegments

`func (o *TrailerRecord) HasTotalPaymentInformationSegments() bool`

HasTotalPaymentInformationSegments returns a boolean if a field has been set.

### GetTotalChangeSegments

`func (o *TrailerRecord) GetTotalChangeSegments() int32`

GetTotalChangeSegments returns the TotalChangeSegments field if non-nil, zero value otherwise.

### GetTotalChangeSegmentsOk

`func (o *TrailerRecord) GetTotalChangeSegmentsOk() (*int32, bool)`

GetTotalChangeSegmentsOk returns a tuple with the TotalChangeSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChangeSegments

`func (o *TrailerRecord) SetTotalChangeSegments(v int32)`

SetTotalChangeSegments sets TotalChangeSegments field to given value.

### HasTotalChangeSegments

`func (o *TrailerRecord) HasTotalChangeSegments() bool`

HasTotalChangeSegments returns a boolean if a field has been set.

### GetTotalSocialNumbersAllSegments

`func (o *TrailerRecord) GetTotalSocialNumbersAllSegments() int32`

GetTotalSocialNumbersAllSegments returns the TotalSocialNumbersAllSegments field if non-nil, zero value otherwise.

### GetTotalSocialNumbersAllSegmentsOk

`func (o *TrailerRecord) GetTotalSocialNumbersAllSegmentsOk() (*int32, bool)`

GetTotalSocialNumbersAllSegmentsOk returns a tuple with the TotalSocialNumbersAllSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSocialNumbersAllSegments

`func (o *TrailerRecord) SetTotalSocialNumbersAllSegments(v int32)`

SetTotalSocialNumbersAllSegments sets TotalSocialNumbersAllSegments field to given value.

### HasTotalSocialNumbersAllSegments

`func (o *TrailerRecord) HasTotalSocialNumbersAllSegments() bool`

HasTotalSocialNumbersAllSegments returns a boolean if a field has been set.

### GetTotalSocialNumbersBaseSegments

`func (o *TrailerRecord) GetTotalSocialNumbersBaseSegments() int32`

GetTotalSocialNumbersBaseSegments returns the TotalSocialNumbersBaseSegments field if non-nil, zero value otherwise.

### GetTotalSocialNumbersBaseSegmentsOk

`func (o *TrailerRecord) GetTotalSocialNumbersBaseSegmentsOk() (*int32, bool)`

GetTotalSocialNumbersBaseSegmentsOk returns a tuple with the TotalSocialNumbersBaseSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSocialNumbersBaseSegments

`func (o *TrailerRecord) SetTotalSocialNumbersBaseSegments(v int32)`

SetTotalSocialNumbersBaseSegments sets TotalSocialNumbersBaseSegments field to given value.

### HasTotalSocialNumbersBaseSegments

`func (o *TrailerRecord) HasTotalSocialNumbersBaseSegments() bool`

HasTotalSocialNumbersBaseSegments returns a boolean if a field has been set.

### GetTotalSocialNumbersJ1Segments

`func (o *TrailerRecord) GetTotalSocialNumbersJ1Segments() int32`

GetTotalSocialNumbersJ1Segments returns the TotalSocialNumbersJ1Segments field if non-nil, zero value otherwise.

### GetTotalSocialNumbersJ1SegmentsOk

`func (o *TrailerRecord) GetTotalSocialNumbersJ1SegmentsOk() (*int32, bool)`

GetTotalSocialNumbersJ1SegmentsOk returns a tuple with the TotalSocialNumbersJ1Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSocialNumbersJ1Segments

`func (o *TrailerRecord) SetTotalSocialNumbersJ1Segments(v int32)`

SetTotalSocialNumbersJ1Segments sets TotalSocialNumbersJ1Segments field to given value.

### HasTotalSocialNumbersJ1Segments

`func (o *TrailerRecord) HasTotalSocialNumbersJ1Segments() bool`

HasTotalSocialNumbersJ1Segments returns a boolean if a field has been set.

### GetTotalSocialNumbersJ2Segments

`func (o *TrailerRecord) GetTotalSocialNumbersJ2Segments() int32`

GetTotalSocialNumbersJ2Segments returns the TotalSocialNumbersJ2Segments field if non-nil, zero value otherwise.

### GetTotalSocialNumbersJ2SegmentsOk

`func (o *TrailerRecord) GetTotalSocialNumbersJ2SegmentsOk() (*int32, bool)`

GetTotalSocialNumbersJ2SegmentsOk returns a tuple with the TotalSocialNumbersJ2Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSocialNumbersJ2Segments

`func (o *TrailerRecord) SetTotalSocialNumbersJ2Segments(v int32)`

SetTotalSocialNumbersJ2Segments sets TotalSocialNumbersJ2Segments field to given value.

### HasTotalSocialNumbersJ2Segments

`func (o *TrailerRecord) HasTotalSocialNumbersJ2Segments() bool`

HasTotalSocialNumbersJ2Segments returns a boolean if a field has been set.

### GetTotalDatesBirthAllSegments

`func (o *TrailerRecord) GetTotalDatesBirthAllSegments() int32`

GetTotalDatesBirthAllSegments returns the TotalDatesBirthAllSegments field if non-nil, zero value otherwise.

### GetTotalDatesBirthAllSegmentsOk

`func (o *TrailerRecord) GetTotalDatesBirthAllSegmentsOk() (*int32, bool)`

GetTotalDatesBirthAllSegmentsOk returns a tuple with the TotalDatesBirthAllSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDatesBirthAllSegments

`func (o *TrailerRecord) SetTotalDatesBirthAllSegments(v int32)`

SetTotalDatesBirthAllSegments sets TotalDatesBirthAllSegments field to given value.

### HasTotalDatesBirthAllSegments

`func (o *TrailerRecord) HasTotalDatesBirthAllSegments() bool`

HasTotalDatesBirthAllSegments returns a boolean if a field has been set.

### GetTotalDatesBirthBaseSegments

`func (o *TrailerRecord) GetTotalDatesBirthBaseSegments() int32`

GetTotalDatesBirthBaseSegments returns the TotalDatesBirthBaseSegments field if non-nil, zero value otherwise.

### GetTotalDatesBirthBaseSegmentsOk

`func (o *TrailerRecord) GetTotalDatesBirthBaseSegmentsOk() (*int32, bool)`

GetTotalDatesBirthBaseSegmentsOk returns a tuple with the TotalDatesBirthBaseSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDatesBirthBaseSegments

`func (o *TrailerRecord) SetTotalDatesBirthBaseSegments(v int32)`

SetTotalDatesBirthBaseSegments sets TotalDatesBirthBaseSegments field to given value.

### HasTotalDatesBirthBaseSegments

`func (o *TrailerRecord) HasTotalDatesBirthBaseSegments() bool`

HasTotalDatesBirthBaseSegments returns a boolean if a field has been set.

### GetTotalDatesBirthJ1Segments

`func (o *TrailerRecord) GetTotalDatesBirthJ1Segments() int32`

GetTotalDatesBirthJ1Segments returns the TotalDatesBirthJ1Segments field if non-nil, zero value otherwise.

### GetTotalDatesBirthJ1SegmentsOk

`func (o *TrailerRecord) GetTotalDatesBirthJ1SegmentsOk() (*int32, bool)`

GetTotalDatesBirthJ1SegmentsOk returns a tuple with the TotalDatesBirthJ1Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDatesBirthJ1Segments

`func (o *TrailerRecord) SetTotalDatesBirthJ1Segments(v int32)`

SetTotalDatesBirthJ1Segments sets TotalDatesBirthJ1Segments field to given value.

### HasTotalDatesBirthJ1Segments

`func (o *TrailerRecord) HasTotalDatesBirthJ1Segments() bool`

HasTotalDatesBirthJ1Segments returns a boolean if a field has been set.

### GetTotalDatesBirthJ2Segments

`func (o *TrailerRecord) GetTotalDatesBirthJ2Segments() int32`

GetTotalDatesBirthJ2Segments returns the TotalDatesBirthJ2Segments field if non-nil, zero value otherwise.

### GetTotalDatesBirthJ2SegmentsOk

`func (o *TrailerRecord) GetTotalDatesBirthJ2SegmentsOk() (*int32, bool)`

GetTotalDatesBirthJ2SegmentsOk returns a tuple with the TotalDatesBirthJ2Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDatesBirthJ2Segments

`func (o *TrailerRecord) SetTotalDatesBirthJ2Segments(v int32)`

SetTotalDatesBirthJ2Segments sets TotalDatesBirthJ2Segments field to given value.

### HasTotalDatesBirthJ2Segments

`func (o *TrailerRecord) HasTotalDatesBirthJ2Segments() bool`

HasTotalDatesBirthJ2Segments returns a boolean if a field has been set.

### GetTotalTelephoneNumbersAllSegments

`func (o *TrailerRecord) GetTotalTelephoneNumbersAllSegments() int32`

GetTotalTelephoneNumbersAllSegments returns the TotalTelephoneNumbersAllSegments field if non-nil, zero value otherwise.

### GetTotalTelephoneNumbersAllSegmentsOk

`func (o *TrailerRecord) GetTotalTelephoneNumbersAllSegmentsOk() (*int32, bool)`

GetTotalTelephoneNumbersAllSegmentsOk returns a tuple with the TotalTelephoneNumbersAllSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTelephoneNumbersAllSegments

`func (o *TrailerRecord) SetTotalTelephoneNumbersAllSegments(v int32)`

SetTotalTelephoneNumbersAllSegments sets TotalTelephoneNumbersAllSegments field to given value.

### HasTotalTelephoneNumbersAllSegments

`func (o *TrailerRecord) HasTotalTelephoneNumbersAllSegments() bool`

HasTotalTelephoneNumbersAllSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


