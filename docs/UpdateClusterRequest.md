# UpdateClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**NumberOfAllocatedHwPackages** | Pointer to **float64** |  | [optional] 
**StageLabel** | Pointer to [**CamundaClusterStage**](CamundaClusterStage.md) |  | [optional] 

## Methods

### NewUpdateClusterRequest

`func NewUpdateClusterRequest() *UpdateClusterRequest`

NewUpdateClusterRequest instantiates a new UpdateClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterRequestWithDefaults

`func NewUpdateClusterRequestWithDefaults() *UpdateClusterRequest`

NewUpdateClusterRequestWithDefaults instantiates a new UpdateClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateClusterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateClusterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfAllocatedHwPackages

`func (o *UpdateClusterRequest) GetNumberOfAllocatedHwPackages() float64`

GetNumberOfAllocatedHwPackages returns the NumberOfAllocatedHwPackages field if non-nil, zero value otherwise.

### GetNumberOfAllocatedHwPackagesOk

`func (o *UpdateClusterRequest) GetNumberOfAllocatedHwPackagesOk() (*float64, bool)`

GetNumberOfAllocatedHwPackagesOk returns a tuple with the NumberOfAllocatedHwPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAllocatedHwPackages

`func (o *UpdateClusterRequest) SetNumberOfAllocatedHwPackages(v float64)`

SetNumberOfAllocatedHwPackages sets NumberOfAllocatedHwPackages field to given value.

### HasNumberOfAllocatedHwPackages

`func (o *UpdateClusterRequest) HasNumberOfAllocatedHwPackages() bool`

HasNumberOfAllocatedHwPackages returns a boolean if a field has been set.

### GetStageLabel

`func (o *UpdateClusterRequest) GetStageLabel() CamundaClusterStage`

GetStageLabel returns the StageLabel field if non-nil, zero value otherwise.

### GetStageLabelOk

`func (o *UpdateClusterRequest) GetStageLabelOk() (*CamundaClusterStage, bool)`

GetStageLabelOk returns a tuple with the StageLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageLabel

`func (o *UpdateClusterRequest) SetStageLabel(v CamundaClusterStage)`

SetStageLabel sets StageLabel field to given value.

### HasStageLabel

`func (o *UpdateClusterRequest) HasStageLabel() bool`

HasStageLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


