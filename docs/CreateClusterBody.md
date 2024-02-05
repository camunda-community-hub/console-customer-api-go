# CreateClusterBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUpdate** | Pointer to **bool** | If set to false, no automatic updates will be performed on your cluster. | [optional] 
**BackupRegionId** | Pointer to **string** | The backup region to choose. Only available if the Sales Plan supports this. (Enterprise) | [optional] 
**ChannelId** | **string** | The channel (software spec) to use. | 
**GenerationId** | **string** | The generation (software version) to use. | 
**Name** | **string** | The name of your new cluster. | 
**PlanTypeId** | **string** | The planType (hardware spec) to use. | 
**RegionId** | **string** | The data center to use. | 
**StageLabel** | Pointer to [**CamundaClusterStage**](CamundaClusterStage.md) |  | [optional] 

## Methods

### NewCreateClusterBody

`func NewCreateClusterBody(channelId string, generationId string, name string, planTypeId string, regionId string, ) *CreateClusterBody`

NewCreateClusterBody instantiates a new CreateClusterBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterBodyWithDefaults

`func NewCreateClusterBodyWithDefaults() *CreateClusterBody`

NewCreateClusterBodyWithDefaults instantiates a new CreateClusterBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUpdate

`func (o *CreateClusterBody) GetAutoUpdate() bool`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *CreateClusterBody) GetAutoUpdateOk() (*bool, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *CreateClusterBody) SetAutoUpdate(v bool)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *CreateClusterBody) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### GetBackupRegionId

`func (o *CreateClusterBody) GetBackupRegionId() string`

GetBackupRegionId returns the BackupRegionId field if non-nil, zero value otherwise.

### GetBackupRegionIdOk

`func (o *CreateClusterBody) GetBackupRegionIdOk() (*string, bool)`

GetBackupRegionIdOk returns a tuple with the BackupRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRegionId

`func (o *CreateClusterBody) SetBackupRegionId(v string)`

SetBackupRegionId sets BackupRegionId field to given value.

### HasBackupRegionId

`func (o *CreateClusterBody) HasBackupRegionId() bool`

HasBackupRegionId returns a boolean if a field has been set.

### GetChannelId

`func (o *CreateClusterBody) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *CreateClusterBody) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *CreateClusterBody) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetGenerationId

`func (o *CreateClusterBody) GetGenerationId() string`

GetGenerationId returns the GenerationId field if non-nil, zero value otherwise.

### GetGenerationIdOk

`func (o *CreateClusterBody) GetGenerationIdOk() (*string, bool)`

GetGenerationIdOk returns a tuple with the GenerationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationId

`func (o *CreateClusterBody) SetGenerationId(v string)`

SetGenerationId sets GenerationId field to given value.


### GetName

`func (o *CreateClusterBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClusterBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClusterBody) SetName(v string)`

SetName sets Name field to given value.


### GetPlanTypeId

`func (o *CreateClusterBody) GetPlanTypeId() string`

GetPlanTypeId returns the PlanTypeId field if non-nil, zero value otherwise.

### GetPlanTypeIdOk

`func (o *CreateClusterBody) GetPlanTypeIdOk() (*string, bool)`

GetPlanTypeIdOk returns a tuple with the PlanTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanTypeId

`func (o *CreateClusterBody) SetPlanTypeId(v string)`

SetPlanTypeId sets PlanTypeId field to given value.


### GetRegionId

`func (o *CreateClusterBody) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CreateClusterBody) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CreateClusterBody) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetStageLabel

`func (o *CreateClusterBody) GetStageLabel() CamundaClusterStage`

GetStageLabel returns the StageLabel field if non-nil, zero value otherwise.

### GetStageLabelOk

`func (o *CreateClusterBody) GetStageLabelOk() (*CamundaClusterStage, bool)`

GetStageLabelOk returns a tuple with the StageLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageLabel

`func (o *CreateClusterBody) SetStageLabel(v CamundaClusterStage)`

SetStageLabel sets StageLabel field to given value.

### HasStageLabel

`func (o *CreateClusterBody) HasStageLabel() bool`

HasStageLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


