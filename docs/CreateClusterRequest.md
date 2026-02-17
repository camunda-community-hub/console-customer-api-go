# CreateClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUpdate** | Pointer to **bool** | If set to false, no automatic updates will be performed on your cluster. | [optional] 
**BackupRegionId** | Pointer to **string** | The backup region to choose. Only available if the Sales Plan supports this. (Enterprise) | [optional] 
**ChannelId** | **string** | The channel (software spec) to use. | 
**Description** | Pointer to **string** | Optional description for the cluster (max 150 characters). | [optional] 
**Encryption** | Pointer to [**ClusterEncryptionKey**](ClusterEncryptionKey.md) |  | [optional] 
**GenerationId** | **string** | The generation (software version) to use. | 
**HardwarePackages** | Pointer to **float64** | Optional number uf hardware packages, defaults to 1. Only availabe on request and for Advanced offering cluster types. | [optional] 
**Name** | **string** | The name of your new cluster. | 
**PlanTypeId** | **string** | The planType (hardware spec) to use. | 
**RegionId** | **string** | The data center to use. | 
**StageLabel** | Pointer to [**CamundaClusterStage**](CamundaClusterStage.md) |  | [optional] 
**IdentityBackendChecksEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateClusterRequest

`func NewCreateClusterRequest(channelId string, generationId string, name string, planTypeId string, regionId string, ) *CreateClusterRequest`

NewCreateClusterRequest instantiates a new CreateClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterRequestWithDefaults

`func NewCreateClusterRequestWithDefaults() *CreateClusterRequest`

NewCreateClusterRequestWithDefaults instantiates a new CreateClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUpdate

`func (o *CreateClusterRequest) GetAutoUpdate() bool`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *CreateClusterRequest) GetAutoUpdateOk() (*bool, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *CreateClusterRequest) SetAutoUpdate(v bool)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *CreateClusterRequest) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### GetBackupRegionId

`func (o *CreateClusterRequest) GetBackupRegionId() string`

GetBackupRegionId returns the BackupRegionId field if non-nil, zero value otherwise.

### GetBackupRegionIdOk

`func (o *CreateClusterRequest) GetBackupRegionIdOk() (*string, bool)`

GetBackupRegionIdOk returns a tuple with the BackupRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRegionId

`func (o *CreateClusterRequest) SetBackupRegionId(v string)`

SetBackupRegionId sets BackupRegionId field to given value.

### HasBackupRegionId

`func (o *CreateClusterRequest) HasBackupRegionId() bool`

HasBackupRegionId returns a boolean if a field has been set.

### GetChannelId

`func (o *CreateClusterRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *CreateClusterRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *CreateClusterRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetDescription

`func (o *CreateClusterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateClusterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateClusterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateClusterRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEncryption

`func (o *CreateClusterRequest) GetEncryption() ClusterEncryptionKey`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *CreateClusterRequest) GetEncryptionOk() (*ClusterEncryptionKey, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *CreateClusterRequest) SetEncryption(v ClusterEncryptionKey)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *CreateClusterRequest) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetGenerationId

`func (o *CreateClusterRequest) GetGenerationId() string`

GetGenerationId returns the GenerationId field if non-nil, zero value otherwise.

### GetGenerationIdOk

`func (o *CreateClusterRequest) GetGenerationIdOk() (*string, bool)`

GetGenerationIdOk returns a tuple with the GenerationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationId

`func (o *CreateClusterRequest) SetGenerationId(v string)`

SetGenerationId sets GenerationId field to given value.


### GetHardwarePackages

`func (o *CreateClusterRequest) GetHardwarePackages() float64`

GetHardwarePackages returns the HardwarePackages field if non-nil, zero value otherwise.

### GetHardwarePackagesOk

`func (o *CreateClusterRequest) GetHardwarePackagesOk() (*float64, bool)`

GetHardwarePackagesOk returns a tuple with the HardwarePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwarePackages

`func (o *CreateClusterRequest) SetHardwarePackages(v float64)`

SetHardwarePackages sets HardwarePackages field to given value.

### HasHardwarePackages

`func (o *CreateClusterRequest) HasHardwarePackages() bool`

HasHardwarePackages returns a boolean if a field has been set.

### GetName

`func (o *CreateClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClusterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPlanTypeId

`func (o *CreateClusterRequest) GetPlanTypeId() string`

GetPlanTypeId returns the PlanTypeId field if non-nil, zero value otherwise.

### GetPlanTypeIdOk

`func (o *CreateClusterRequest) GetPlanTypeIdOk() (*string, bool)`

GetPlanTypeIdOk returns a tuple with the PlanTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanTypeId

`func (o *CreateClusterRequest) SetPlanTypeId(v string)`

SetPlanTypeId sets PlanTypeId field to given value.


### GetRegionId

`func (o *CreateClusterRequest) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CreateClusterRequest) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CreateClusterRequest) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetStageLabel

`func (o *CreateClusterRequest) GetStageLabel() CamundaClusterStage`

GetStageLabel returns the StageLabel field if non-nil, zero value otherwise.

### GetStageLabelOk

`func (o *CreateClusterRequest) GetStageLabelOk() (*CamundaClusterStage, bool)`

GetStageLabelOk returns a tuple with the StageLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageLabel

`func (o *CreateClusterRequest) SetStageLabel(v CamundaClusterStage)`

SetStageLabel sets StageLabel field to given value.

### HasStageLabel

`func (o *CreateClusterRequest) HasStageLabel() bool`

HasStageLabel returns a boolean if a field has been set.

### GetIdentityBackendChecksEnabled

`func (o *CreateClusterRequest) GetIdentityBackendChecksEnabled() bool`

GetIdentityBackendChecksEnabled returns the IdentityBackendChecksEnabled field if non-nil, zero value otherwise.

### GetIdentityBackendChecksEnabledOk

`func (o *CreateClusterRequest) GetIdentityBackendChecksEnabledOk() (*bool, bool)`

GetIdentityBackendChecksEnabledOk returns a tuple with the IdentityBackendChecksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityBackendChecksEnabled

`func (o *CreateClusterRequest) SetIdentityBackendChecksEnabled(v bool)`

SetIdentityBackendChecksEnabled sets IdentityBackendChecksEnabled field to given value.

### HasIdentityBackendChecksEnabled

`func (o *CreateClusterRequest) HasIdentityBackendChecksEnabled() bool`

HasIdentityBackendChecksEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


