# BackupScheduleBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | **string** | Standard five-field cron expression in UTC (&#x60;minute hour day-of-month month day-of-week&#x60;), for example &#x60;0 2 * * *&#x60; for a daily backup at 02:00 UTC. Backups can be scheduled at most once every 15 minutes — the same floor that applies to manual backups. | 

## Methods

### NewBackupScheduleBody

`func NewBackupScheduleBody(schedule string, ) *BackupScheduleBody`

NewBackupScheduleBody instantiates a new BackupScheduleBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleBodyWithDefaults

`func NewBackupScheduleBodyWithDefaults() *BackupScheduleBody`

NewBackupScheduleBodyWithDefaults instantiates a new BackupScheduleBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *BackupScheduleBody) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *BackupScheduleBody) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *BackupScheduleBody) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


