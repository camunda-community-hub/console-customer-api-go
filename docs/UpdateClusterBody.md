# UpdateClusterBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of your cluster. | [optional] 
**NumberOfAllocatedHwPackages** | Pointer to **float64** | Optional number uf hardware packages, defaults to 1. Only availabe on request and for Advanced offering cluster types. | [optional] 
**StageLabel** | Pointer to [**CamundaClusterStage**](CamundaClusterStage.md) |  | [optional] 

## Methods

### NewUpdateClusterBody

`func NewUpdateClusterBody() *UpdateClusterBody`

NewUpdateClusterBody instantiates a new UpdateClusterBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterBodyWithDefaults

`func NewUpdateClusterBodyWithDefaults() *UpdateClusterBody`

NewUpdateClusterBodyWithDefaults instantiates a new UpdateClusterBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateClusterBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateClusterBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateClusterBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateClusterBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfAllocatedHwPackages

`func (o *UpdateClusterBody) GetNumberOfAllocatedHwPackages() float64`

GetNumberOfAllocatedHwPackages returns the NumberOfAllocatedHwPackages field if non-nil, zero value otherwise.

### GetNumberOfAllocatedHwPackagesOk

`func (o *UpdateClusterBody) GetNumberOfAllocatedHwPackagesOk() (*float64, bool)`

GetNumberOfAllocatedHwPackagesOk returns a tuple with the NumberOfAllocatedHwPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAllocatedHwPackages

`func (o *UpdateClusterBody) SetNumberOfAllocatedHwPackages(v float64)`

SetNumberOfAllocatedHwPackages sets NumberOfAllocatedHwPackages field to given value.

### HasNumberOfAllocatedHwPackages

`func (o *UpdateClusterBody) HasNumberOfAllocatedHwPackages() bool`

HasNumberOfAllocatedHwPackages returns a boolean if a field has been set.

### GetStageLabel

`func (o *UpdateClusterBody) GetStageLabel() CamundaClusterStage`

GetStageLabel returns the StageLabel field if non-nil, zero value otherwise.

### GetStageLabelOk

`func (o *UpdateClusterBody) GetStageLabelOk() (*CamundaClusterStage, bool)`

GetStageLabelOk returns a tuple with the StageLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageLabel

`func (o *UpdateClusterBody) SetStageLabel(v CamundaClusterStage)`

SetStageLabel sets StageLabel field to given value.

### HasStageLabel

`func (o *UpdateClusterBody) HasStageLabel() bool`

HasStageLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


