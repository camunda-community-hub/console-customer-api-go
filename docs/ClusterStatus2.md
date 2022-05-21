# ClusterStatus2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptimizeStatus** | Pointer to [**ClusterStatus**](ClusterStatus.md) |  | [optional] 
**TasklistStatus** | Pointer to [**ClusterStatus**](ClusterStatus.md) |  | [optional] 
**OperateStatus** | Pointer to [**ClusterStatus**](ClusterStatus.md) |  | [optional] 
**ZeebeStatus** | Pointer to [**ClusterStatus**](ClusterStatus.md) |  | [optional] 
**Ready** | [**ClusterStatus**](ClusterStatus.md) |  | 

## Methods

### NewClusterStatus2

`func NewClusterStatus2(ready ClusterStatus, ) *ClusterStatus2`

NewClusterStatus2 instantiates a new ClusterStatus2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatus2WithDefaults

`func NewClusterStatus2WithDefaults() *ClusterStatus2`

NewClusterStatus2WithDefaults instantiates a new ClusterStatus2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptimizeStatus

`func (o *ClusterStatus2) GetOptimizeStatus() ClusterStatus`

GetOptimizeStatus returns the OptimizeStatus field if non-nil, zero value otherwise.

### GetOptimizeStatusOk

`func (o *ClusterStatus2) GetOptimizeStatusOk() (*ClusterStatus, bool)`

GetOptimizeStatusOk returns a tuple with the OptimizeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeStatus

`func (o *ClusterStatus2) SetOptimizeStatus(v ClusterStatus)`

SetOptimizeStatus sets OptimizeStatus field to given value.

### HasOptimizeStatus

`func (o *ClusterStatus2) HasOptimizeStatus() bool`

HasOptimizeStatus returns a boolean if a field has been set.

### GetTasklistStatus

`func (o *ClusterStatus2) GetTasklistStatus() ClusterStatus`

GetTasklistStatus returns the TasklistStatus field if non-nil, zero value otherwise.

### GetTasklistStatusOk

`func (o *ClusterStatus2) GetTasklistStatusOk() (*ClusterStatus, bool)`

GetTasklistStatusOk returns a tuple with the TasklistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistStatus

`func (o *ClusterStatus2) SetTasklistStatus(v ClusterStatus)`

SetTasklistStatus sets TasklistStatus field to given value.

### HasTasklistStatus

`func (o *ClusterStatus2) HasTasklistStatus() bool`

HasTasklistStatus returns a boolean if a field has been set.

### GetOperateStatus

`func (o *ClusterStatus2) GetOperateStatus() ClusterStatus`

GetOperateStatus returns the OperateStatus field if non-nil, zero value otherwise.

### GetOperateStatusOk

`func (o *ClusterStatus2) GetOperateStatusOk() (*ClusterStatus, bool)`

GetOperateStatusOk returns a tuple with the OperateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperateStatus

`func (o *ClusterStatus2) SetOperateStatus(v ClusterStatus)`

SetOperateStatus sets OperateStatus field to given value.

### HasOperateStatus

`func (o *ClusterStatus2) HasOperateStatus() bool`

HasOperateStatus returns a boolean if a field has been set.

### GetZeebeStatus

`func (o *ClusterStatus2) GetZeebeStatus() ClusterStatus`

GetZeebeStatus returns the ZeebeStatus field if non-nil, zero value otherwise.

### GetZeebeStatusOk

`func (o *ClusterStatus2) GetZeebeStatusOk() (*ClusterStatus, bool)`

GetZeebeStatusOk returns a tuple with the ZeebeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeebeStatus

`func (o *ClusterStatus2) SetZeebeStatus(v ClusterStatus)`

SetZeebeStatus sets ZeebeStatus field to given value.

### HasZeebeStatus

`func (o *ClusterStatus2) HasZeebeStatus() bool`

HasZeebeStatus returns a boolean if a field has been set.

### GetReady

`func (o *ClusterStatus2) GetReady() ClusterStatus`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *ClusterStatus2) GetReadyOk() (*ClusterStatus, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *ClusterStatus2) SetReady(v ClusterStatus)`

SetReady sets Ready field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


