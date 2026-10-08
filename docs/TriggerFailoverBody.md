# TriggerFailoverBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupIds** | **[]string** | One or more backups to clone into the failover cluster. | 
**RestoreBackupId** | **string** | The backup to restore onto the new cluster immediately. Must be one of &#x60;backupIds&#x60;. | 

## Methods

### NewTriggerFailoverBody

`func NewTriggerFailoverBody(backupIds []string, restoreBackupId string, ) *TriggerFailoverBody`

NewTriggerFailoverBody instantiates a new TriggerFailoverBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerFailoverBodyWithDefaults

`func NewTriggerFailoverBodyWithDefaults() *TriggerFailoverBody`

NewTriggerFailoverBodyWithDefaults instantiates a new TriggerFailoverBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupIds

`func (o *TriggerFailoverBody) GetBackupIds() []string`

GetBackupIds returns the BackupIds field if non-nil, zero value otherwise.

### GetBackupIdsOk

`func (o *TriggerFailoverBody) GetBackupIdsOk() (*[]string, bool)`

GetBackupIdsOk returns a tuple with the BackupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupIds

`func (o *TriggerFailoverBody) SetBackupIds(v []string)`

SetBackupIds sets BackupIds field to given value.


### GetRestoreBackupId

`func (o *TriggerFailoverBody) GetRestoreBackupId() string`

GetRestoreBackupId returns the RestoreBackupId field if non-nil, zero value otherwise.

### GetRestoreBackupIdOk

`func (o *TriggerFailoverBody) GetRestoreBackupIdOk() (*string, bool)`

GetRestoreBackupIdOk returns a tuple with the RestoreBackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreBackupId

`func (o *TriggerFailoverBody) SetRestoreBackupId(v string)`

SetRestoreBackupId sets RestoreBackupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


