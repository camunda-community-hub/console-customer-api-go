# RestoreDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | **string** |  | 
**CompletionTimestamp** | Pointer to **string** |  | [optional] 
**StartTimestamp** | Pointer to **string** |  | [optional] 
**State** | [**RestoreState**](RestoreState.md) |  | 
**Uuid** | **string** |  | 

## Methods

### NewRestoreDto

`func NewRestoreDto(backupId string, state RestoreState, uuid string, ) *RestoreDto`

NewRestoreDto instantiates a new RestoreDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreDtoWithDefaults

`func NewRestoreDtoWithDefaults() *RestoreDto`

NewRestoreDtoWithDefaults instantiates a new RestoreDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *RestoreDto) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *RestoreDto) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *RestoreDto) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.


### GetCompletionTimestamp

`func (o *RestoreDto) GetCompletionTimestamp() string`

GetCompletionTimestamp returns the CompletionTimestamp field if non-nil, zero value otherwise.

### GetCompletionTimestampOk

`func (o *RestoreDto) GetCompletionTimestampOk() (*string, bool)`

GetCompletionTimestampOk returns a tuple with the CompletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTimestamp

`func (o *RestoreDto) SetCompletionTimestamp(v string)`

SetCompletionTimestamp sets CompletionTimestamp field to given value.

### HasCompletionTimestamp

`func (o *RestoreDto) HasCompletionTimestamp() bool`

HasCompletionTimestamp returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *RestoreDto) GetStartTimestamp() string`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *RestoreDto) GetStartTimestampOk() (*string, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *RestoreDto) SetStartTimestamp(v string)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *RestoreDto) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetState

`func (o *RestoreDto) GetState() RestoreState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RestoreDto) GetStateOk() (*RestoreState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RestoreDto) SetState(v RestoreState)`

SetState sets State field to given value.


### GetUuid

`func (o *RestoreDto) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RestoreDto) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RestoreDto) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


