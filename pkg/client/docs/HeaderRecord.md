# HeaderRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockDescriptorWord** | Pointer to **int32** |  | [optional] 
**RecordDescriptorWord** | **int32** |  | 
**RecordIdentifier** | **string** |  | 
**CycleIdentifier** | Pointer to **string** |  | [optional] 
**InnovisProgramIdentifier** | Pointer to **string** |  | [optional] 
**EquifaxProgramIdentifier** | Pointer to **string** |  | [optional] 
**ExperianProgramIdentifier** | Pointer to **string** |  | [optional] 
**TransUnionProgramIdentifier** | Pointer to **string** |  | [optional] 
**ActivityDate** | **time.Time** |  | 
**DateCreated** | **time.Time** |  | 
**ProgramDate** | Pointer to **time.Time** |  | [optional] 
**ProgramRevisionDate** | Pointer to **time.Time** |  | [optional] 
**ReporterName** | **string** |  | 
**ReporterAddress** | **string** |  | 
**ReporterTelephoneNumber** | Pointer to **int64** |  | [optional] 
**SoftwareVendorName** | Pointer to **string** |  | [optional] 
**SoftwareVersionNumber** | Pointer to **string** |  | [optional] 
**PrbcProgramIdentifier** | Pointer to **string** |  | [optional] 

## Methods

### NewHeaderRecord

`func NewHeaderRecord(recordDescriptorWord int32, recordIdentifier string, activityDate time.Time, dateCreated time.Time, reporterName string, reporterAddress string, ) *HeaderRecord`

NewHeaderRecord instantiates a new HeaderRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderRecordWithDefaults

`func NewHeaderRecordWithDefaults() *HeaderRecord`

NewHeaderRecordWithDefaults instantiates a new HeaderRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockDescriptorWord

`func (o *HeaderRecord) GetBlockDescriptorWord() int32`

GetBlockDescriptorWord returns the BlockDescriptorWord field if non-nil, zero value otherwise.

### GetBlockDescriptorWordOk

`func (o *HeaderRecord) GetBlockDescriptorWordOk() (*int32, bool)`

GetBlockDescriptorWordOk returns a tuple with the BlockDescriptorWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDescriptorWord

`func (o *HeaderRecord) SetBlockDescriptorWord(v int32)`

SetBlockDescriptorWord sets BlockDescriptorWord field to given value.

### HasBlockDescriptorWord

`func (o *HeaderRecord) HasBlockDescriptorWord() bool`

HasBlockDescriptorWord returns a boolean if a field has been set.

### GetRecordDescriptorWord

`func (o *HeaderRecord) GetRecordDescriptorWord() int32`

GetRecordDescriptorWord returns the RecordDescriptorWord field if non-nil, zero value otherwise.

### GetRecordDescriptorWordOk

`func (o *HeaderRecord) GetRecordDescriptorWordOk() (*int32, bool)`

GetRecordDescriptorWordOk returns a tuple with the RecordDescriptorWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordDescriptorWord

`func (o *HeaderRecord) SetRecordDescriptorWord(v int32)`

SetRecordDescriptorWord sets RecordDescriptorWord field to given value.


### GetRecordIdentifier

`func (o *HeaderRecord) GetRecordIdentifier() string`

GetRecordIdentifier returns the RecordIdentifier field if non-nil, zero value otherwise.

### GetRecordIdentifierOk

`func (o *HeaderRecord) GetRecordIdentifierOk() (*string, bool)`

GetRecordIdentifierOk returns a tuple with the RecordIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordIdentifier

`func (o *HeaderRecord) SetRecordIdentifier(v string)`

SetRecordIdentifier sets RecordIdentifier field to given value.


### GetCycleIdentifier

`func (o *HeaderRecord) GetCycleIdentifier() string`

GetCycleIdentifier returns the CycleIdentifier field if non-nil, zero value otherwise.

### GetCycleIdentifierOk

`func (o *HeaderRecord) GetCycleIdentifierOk() (*string, bool)`

GetCycleIdentifierOk returns a tuple with the CycleIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleIdentifier

`func (o *HeaderRecord) SetCycleIdentifier(v string)`

SetCycleIdentifier sets CycleIdentifier field to given value.

### HasCycleIdentifier

`func (o *HeaderRecord) HasCycleIdentifier() bool`

HasCycleIdentifier returns a boolean if a field has been set.

### GetInnovisProgramIdentifier

`func (o *HeaderRecord) GetInnovisProgramIdentifier() string`

GetInnovisProgramIdentifier returns the InnovisProgramIdentifier field if non-nil, zero value otherwise.

### GetInnovisProgramIdentifierOk

`func (o *HeaderRecord) GetInnovisProgramIdentifierOk() (*string, bool)`

GetInnovisProgramIdentifierOk returns a tuple with the InnovisProgramIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnovisProgramIdentifier

`func (o *HeaderRecord) SetInnovisProgramIdentifier(v string)`

SetInnovisProgramIdentifier sets InnovisProgramIdentifier field to given value.

### HasInnovisProgramIdentifier

`func (o *HeaderRecord) HasInnovisProgramIdentifier() bool`

HasInnovisProgramIdentifier returns a boolean if a field has been set.

### GetEquifaxProgramIdentifier

`func (o *HeaderRecord) GetEquifaxProgramIdentifier() string`

GetEquifaxProgramIdentifier returns the EquifaxProgramIdentifier field if non-nil, zero value otherwise.

### GetEquifaxProgramIdentifierOk

`func (o *HeaderRecord) GetEquifaxProgramIdentifierOk() (*string, bool)`

GetEquifaxProgramIdentifierOk returns a tuple with the EquifaxProgramIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquifaxProgramIdentifier

`func (o *HeaderRecord) SetEquifaxProgramIdentifier(v string)`

SetEquifaxProgramIdentifier sets EquifaxProgramIdentifier field to given value.

### HasEquifaxProgramIdentifier

`func (o *HeaderRecord) HasEquifaxProgramIdentifier() bool`

HasEquifaxProgramIdentifier returns a boolean if a field has been set.

### GetExperianProgramIdentifier

`func (o *HeaderRecord) GetExperianProgramIdentifier() string`

GetExperianProgramIdentifier returns the ExperianProgramIdentifier field if non-nil, zero value otherwise.

### GetExperianProgramIdentifierOk

`func (o *HeaderRecord) GetExperianProgramIdentifierOk() (*string, bool)`

GetExperianProgramIdentifierOk returns a tuple with the ExperianProgramIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperianProgramIdentifier

`func (o *HeaderRecord) SetExperianProgramIdentifier(v string)`

SetExperianProgramIdentifier sets ExperianProgramIdentifier field to given value.

### HasExperianProgramIdentifier

`func (o *HeaderRecord) HasExperianProgramIdentifier() bool`

HasExperianProgramIdentifier returns a boolean if a field has been set.

### GetTransUnionProgramIdentifier

`func (o *HeaderRecord) GetTransUnionProgramIdentifier() string`

GetTransUnionProgramIdentifier returns the TransUnionProgramIdentifier field if non-nil, zero value otherwise.

### GetTransUnionProgramIdentifierOk

`func (o *HeaderRecord) GetTransUnionProgramIdentifierOk() (*string, bool)`

GetTransUnionProgramIdentifierOk returns a tuple with the TransUnionProgramIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransUnionProgramIdentifier

`func (o *HeaderRecord) SetTransUnionProgramIdentifier(v string)`

SetTransUnionProgramIdentifier sets TransUnionProgramIdentifier field to given value.

### HasTransUnionProgramIdentifier

`func (o *HeaderRecord) HasTransUnionProgramIdentifier() bool`

HasTransUnionProgramIdentifier returns a boolean if a field has been set.

### GetActivityDate

`func (o *HeaderRecord) GetActivityDate() time.Time`

GetActivityDate returns the ActivityDate field if non-nil, zero value otherwise.

### GetActivityDateOk

`func (o *HeaderRecord) GetActivityDateOk() (*time.Time, bool)`

GetActivityDateOk returns a tuple with the ActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDate

`func (o *HeaderRecord) SetActivityDate(v time.Time)`

SetActivityDate sets ActivityDate field to given value.


### GetDateCreated

`func (o *HeaderRecord) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *HeaderRecord) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *HeaderRecord) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetProgramDate

`func (o *HeaderRecord) GetProgramDate() time.Time`

GetProgramDate returns the ProgramDate field if non-nil, zero value otherwise.

### GetProgramDateOk

`func (o *HeaderRecord) GetProgramDateOk() (*time.Time, bool)`

GetProgramDateOk returns a tuple with the ProgramDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramDate

`func (o *HeaderRecord) SetProgramDate(v time.Time)`

SetProgramDate sets ProgramDate field to given value.

### HasProgramDate

`func (o *HeaderRecord) HasProgramDate() bool`

HasProgramDate returns a boolean if a field has been set.

### GetProgramRevisionDate

`func (o *HeaderRecord) GetProgramRevisionDate() time.Time`

GetProgramRevisionDate returns the ProgramRevisionDate field if non-nil, zero value otherwise.

### GetProgramRevisionDateOk

`func (o *HeaderRecord) GetProgramRevisionDateOk() (*time.Time, bool)`

GetProgramRevisionDateOk returns a tuple with the ProgramRevisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramRevisionDate

`func (o *HeaderRecord) SetProgramRevisionDate(v time.Time)`

SetProgramRevisionDate sets ProgramRevisionDate field to given value.

### HasProgramRevisionDate

`func (o *HeaderRecord) HasProgramRevisionDate() bool`

HasProgramRevisionDate returns a boolean if a field has been set.

### GetReporterName

`func (o *HeaderRecord) GetReporterName() string`

GetReporterName returns the ReporterName field if non-nil, zero value otherwise.

### GetReporterNameOk

`func (o *HeaderRecord) GetReporterNameOk() (*string, bool)`

GetReporterNameOk returns a tuple with the ReporterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterName

`func (o *HeaderRecord) SetReporterName(v string)`

SetReporterName sets ReporterName field to given value.


### GetReporterAddress

`func (o *HeaderRecord) GetReporterAddress() string`

GetReporterAddress returns the ReporterAddress field if non-nil, zero value otherwise.

### GetReporterAddressOk

`func (o *HeaderRecord) GetReporterAddressOk() (*string, bool)`

GetReporterAddressOk returns a tuple with the ReporterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterAddress

`func (o *HeaderRecord) SetReporterAddress(v string)`

SetReporterAddress sets ReporterAddress field to given value.


### GetReporterTelephoneNumber

`func (o *HeaderRecord) GetReporterTelephoneNumber() int64`

GetReporterTelephoneNumber returns the ReporterTelephoneNumber field if non-nil, zero value otherwise.

### GetReporterTelephoneNumberOk

`func (o *HeaderRecord) GetReporterTelephoneNumberOk() (*int64, bool)`

GetReporterTelephoneNumberOk returns a tuple with the ReporterTelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterTelephoneNumber

`func (o *HeaderRecord) SetReporterTelephoneNumber(v int64)`

SetReporterTelephoneNumber sets ReporterTelephoneNumber field to given value.

### HasReporterTelephoneNumber

`func (o *HeaderRecord) HasReporterTelephoneNumber() bool`

HasReporterTelephoneNumber returns a boolean if a field has been set.

### GetSoftwareVendorName

`func (o *HeaderRecord) GetSoftwareVendorName() string`

GetSoftwareVendorName returns the SoftwareVendorName field if non-nil, zero value otherwise.

### GetSoftwareVendorNameOk

`func (o *HeaderRecord) GetSoftwareVendorNameOk() (*string, bool)`

GetSoftwareVendorNameOk returns a tuple with the SoftwareVendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVendorName

`func (o *HeaderRecord) SetSoftwareVendorName(v string)`

SetSoftwareVendorName sets SoftwareVendorName field to given value.

### HasSoftwareVendorName

`func (o *HeaderRecord) HasSoftwareVendorName() bool`

HasSoftwareVendorName returns a boolean if a field has been set.

### GetSoftwareVersionNumber

`func (o *HeaderRecord) GetSoftwareVersionNumber() string`

GetSoftwareVersionNumber returns the SoftwareVersionNumber field if non-nil, zero value otherwise.

### GetSoftwareVersionNumberOk

`func (o *HeaderRecord) GetSoftwareVersionNumberOk() (*string, bool)`

GetSoftwareVersionNumberOk returns a tuple with the SoftwareVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersionNumber

`func (o *HeaderRecord) SetSoftwareVersionNumber(v string)`

SetSoftwareVersionNumber sets SoftwareVersionNumber field to given value.

### HasSoftwareVersionNumber

`func (o *HeaderRecord) HasSoftwareVersionNumber() bool`

HasSoftwareVersionNumber returns a boolean if a field has been set.

### GetPrbcProgramIdentifier

`func (o *HeaderRecord) GetPrbcProgramIdentifier() string`

GetPrbcProgramIdentifier returns the PrbcProgramIdentifier field if non-nil, zero value otherwise.

### GetPrbcProgramIdentifierOk

`func (o *HeaderRecord) GetPrbcProgramIdentifierOk() (*string, bool)`

GetPrbcProgramIdentifierOk returns a tuple with the PrbcProgramIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrbcProgramIdentifier

`func (o *HeaderRecord) SetPrbcProgramIdentifier(v string)`

SetPrbcProgramIdentifier sets PrbcProgramIdentifier field to given value.

### HasPrbcProgramIdentifier

`func (o *HeaderRecord) HasPrbcProgramIdentifier() bool`

HasPrbcProgramIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


