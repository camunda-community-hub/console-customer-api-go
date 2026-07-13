# BackupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | **string** |  | 
**Created** | **string** |  | 
**GenerationName** | Pointer to **string** | Human-readable name of the Generation referenced by &#x60;generationUuid&#x60;. Resolved server-side: the Generation entity is looked up by &#x60;generationUuid&#x60; and its &#x60;name&#x60; field is propagated here. Falls back to &#x60;&#39;Unknown Generation&#39;&#x60; if the generation can&#39;t be found (e.g. it was deleted). &#x60;undefined&#x60; when the backup has no &#x60;generationUuid&#x60; at all (legacy backups). | [optional] 
**GenerationUuid** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**OperateStatus** | [**BackupStatus**](BackupStatus.md) |  | 
**OptimizeStatus** | [**BackupStatus**](BackupStatus.md) |  | 
**Status** | [**BackupStatus**](BackupStatus.md) |  | 
**TasklistStatus** | [**BackupStatus**](BackupStatus.md) |  | 
**Uuid** | **string** |  | 
**ZeebeStatus** | [**BackupStatus**](BackupStatus.md) |  | 

## Methods

### NewBackupDto

`func NewBackupDto(completed string, created string, name string, operateStatus BackupStatus, optimizeStatus BackupStatus, status BackupStatus, tasklistStatus BackupStatus, uuid string, zeebeStatus BackupStatus, ) *BackupDto`

NewBackupDto instantiates a new BackupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupDtoWithDefaults

`func NewBackupDtoWithDefaults() *BackupDto`

NewBackupDtoWithDefaults instantiates a new BackupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *BackupDto) GetCompleted() string`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *BackupDto) GetCompletedOk() (*string, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *BackupDto) SetCompleted(v string)`

SetCompleted sets Completed field to given value.


### GetCreated

`func (o *BackupDto) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BackupDto) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BackupDto) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetGenerationName

`func (o *BackupDto) GetGenerationName() string`

GetGenerationName returns the GenerationName field if non-nil, zero value otherwise.

### GetGenerationNameOk

`func (o *BackupDto) GetGenerationNameOk() (*string, bool)`

GetGenerationNameOk returns a tuple with the GenerationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationName

`func (o *BackupDto) SetGenerationName(v string)`

SetGenerationName sets GenerationName field to given value.

### HasGenerationName

`func (o *BackupDto) HasGenerationName() bool`

HasGenerationName returns a boolean if a field has been set.

### GetGenerationUuid

`func (o *BackupDto) GetGenerationUuid() string`

GetGenerationUuid returns the GenerationUuid field if non-nil, zero value otherwise.

### GetGenerationUuidOk

`func (o *BackupDto) GetGenerationUuidOk() (*string, bool)`

GetGenerationUuidOk returns a tuple with the GenerationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationUuid

`func (o *BackupDto) SetGenerationUuid(v string)`

SetGenerationUuid sets GenerationUuid field to given value.

### HasGenerationUuid

`func (o *BackupDto) HasGenerationUuid() bool`

HasGenerationUuid returns a boolean if a field has been set.

### GetName

`func (o *BackupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupDto) SetName(v string)`

SetName sets Name field to given value.


### GetOperateStatus

`func (o *BackupDto) GetOperateStatus() BackupStatus`

GetOperateStatus returns the OperateStatus field if non-nil, zero value otherwise.

### GetOperateStatusOk

`func (o *BackupDto) GetOperateStatusOk() (*BackupStatus, bool)`

GetOperateStatusOk returns a tuple with the OperateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperateStatus

`func (o *BackupDto) SetOperateStatus(v BackupStatus)`

SetOperateStatus sets OperateStatus field to given value.


### GetOptimizeStatus

`func (o *BackupDto) GetOptimizeStatus() BackupStatus`

GetOptimizeStatus returns the OptimizeStatus field if non-nil, zero value otherwise.

### GetOptimizeStatusOk

`func (o *BackupDto) GetOptimizeStatusOk() (*BackupStatus, bool)`

GetOptimizeStatusOk returns a tuple with the OptimizeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeStatus

`func (o *BackupDto) SetOptimizeStatus(v BackupStatus)`

SetOptimizeStatus sets OptimizeStatus field to given value.


### GetStatus

`func (o *BackupDto) GetStatus() BackupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupDto) GetStatusOk() (*BackupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupDto) SetStatus(v BackupStatus)`

SetStatus sets Status field to given value.


### GetTasklistStatus

`func (o *BackupDto) GetTasklistStatus() BackupStatus`

GetTasklistStatus returns the TasklistStatus field if non-nil, zero value otherwise.

### GetTasklistStatusOk

`func (o *BackupDto) GetTasklistStatusOk() (*BackupStatus, bool)`

GetTasklistStatusOk returns a tuple with the TasklistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistStatus

`func (o *BackupDto) SetTasklistStatus(v BackupStatus)`

SetTasklistStatus sets TasklistStatus field to given value.


### GetUuid

`func (o *BackupDto) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *BackupDto) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *BackupDto) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetZeebeStatus

`func (o *BackupDto) GetZeebeStatus() BackupStatus`

GetZeebeStatus returns the ZeebeStatus field if non-nil, zero value otherwise.

### GetZeebeStatusOk

`func (o *BackupDto) GetZeebeStatusOk() (*BackupStatus, bool)`

GetZeebeStatusOk returns a tuple with the ZeebeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeebeStatus

`func (o *BackupDto) SetZeebeStatus(v BackupStatus)`

SetZeebeStatus sets ZeebeStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


