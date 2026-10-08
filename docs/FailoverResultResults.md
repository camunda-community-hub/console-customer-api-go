# FailoverResultResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupClone** | [**FailoverResultResultsBackupClone**](FailoverResultResultsBackupClone.md) |  | 
**Restore** | [**FailoverResultResultsBackupClone**](FailoverResultResultsBackupClone.md) |  | 

## Methods

### NewFailoverResultResults

`func NewFailoverResultResults(backupClone FailoverResultResultsBackupClone, restore FailoverResultResultsBackupClone, ) *FailoverResultResults`

NewFailoverResultResults instantiates a new FailoverResultResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverResultResultsWithDefaults

`func NewFailoverResultResultsWithDefaults() *FailoverResultResults`

NewFailoverResultResultsWithDefaults instantiates a new FailoverResultResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupClone

`func (o *FailoverResultResults) GetBackupClone() FailoverResultResultsBackupClone`

GetBackupClone returns the BackupClone field if non-nil, zero value otherwise.

### GetBackupCloneOk

`func (o *FailoverResultResults) GetBackupCloneOk() (*FailoverResultResultsBackupClone, bool)`

GetBackupCloneOk returns a tuple with the BackupClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupClone

`func (o *FailoverResultResults) SetBackupClone(v FailoverResultResultsBackupClone)`

SetBackupClone sets BackupClone field to given value.


### GetRestore

`func (o *FailoverResultResults) GetRestore() FailoverResultResultsBackupClone`

GetRestore returns the Restore field if non-nil, zero value otherwise.

### GetRestoreOk

`func (o *FailoverResultResults) GetRestoreOk() (*FailoverResultResultsBackupClone, bool)`

GetRestoreOk returns a tuple with the Restore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestore

`func (o *FailoverResultResults) SetRestore(v FailoverResultResultsBackupClone)`

SetRestore sets Restore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


