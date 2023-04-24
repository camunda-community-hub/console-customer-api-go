# ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptimizeStatus** | Pointer to [**ClusterHealthStatus**](ClusterHealthStatus.md) |  | [optional] 
**TasklistStatus** | Pointer to [**ClusterHealthStatus**](ClusterHealthStatus.md) |  | [optional] 
**OperateStatus** | Pointer to [**ClusterHealthStatus**](ClusterHealthStatus.md) |  | [optional] 
**ZeebeStatus** | Pointer to [**ClusterHealthStatus**](ClusterHealthStatus.md) |  | [optional] 
**Ready** | [**ClusterHealthStatus**](ClusterHealthStatus.md) |  | 

## Methods

### NewClusterStatus

`func NewClusterStatus(ready ClusterHealthStatus, ) *ClusterStatus`

NewClusterStatus instantiates a new ClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusWithDefaults

`func NewClusterStatusWithDefaults() *ClusterStatus`

NewClusterStatusWithDefaults instantiates a new ClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptimizeStatus

`func (o *ClusterStatus) GetOptimizeStatus() ClusterHealthStatus`

GetOptimizeStatus returns the OptimizeStatus field if non-nil, zero value otherwise.

### GetOptimizeStatusOk

`func (o *ClusterStatus) GetOptimizeStatusOk() (*ClusterHealthStatus, bool)`

GetOptimizeStatusOk returns a tuple with the OptimizeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeStatus

`func (o *ClusterStatus) SetOptimizeStatus(v ClusterHealthStatus)`

SetOptimizeStatus sets OptimizeStatus field to given value.

### HasOptimizeStatus

`func (o *ClusterStatus) HasOptimizeStatus() bool`

HasOptimizeStatus returns a boolean if a field has been set.

### GetTasklistStatus

`func (o *ClusterStatus) GetTasklistStatus() ClusterHealthStatus`

GetTasklistStatus returns the TasklistStatus field if non-nil, zero value otherwise.

### GetTasklistStatusOk

`func (o *ClusterStatus) GetTasklistStatusOk() (*ClusterHealthStatus, bool)`

GetTasklistStatusOk returns a tuple with the TasklistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistStatus

`func (o *ClusterStatus) SetTasklistStatus(v ClusterHealthStatus)`

SetTasklistStatus sets TasklistStatus field to given value.

### HasTasklistStatus

`func (o *ClusterStatus) HasTasklistStatus() bool`

HasTasklistStatus returns a boolean if a field has been set.

### GetOperateStatus

`func (o *ClusterStatus) GetOperateStatus() ClusterHealthStatus`

GetOperateStatus returns the OperateStatus field if non-nil, zero value otherwise.

### GetOperateStatusOk

`func (o *ClusterStatus) GetOperateStatusOk() (*ClusterHealthStatus, bool)`

GetOperateStatusOk returns a tuple with the OperateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperateStatus

`func (o *ClusterStatus) SetOperateStatus(v ClusterHealthStatus)`

SetOperateStatus sets OperateStatus field to given value.

### HasOperateStatus

`func (o *ClusterStatus) HasOperateStatus() bool`

HasOperateStatus returns a boolean if a field has been set.

### GetZeebeStatus

`func (o *ClusterStatus) GetZeebeStatus() ClusterHealthStatus`

GetZeebeStatus returns the ZeebeStatus field if non-nil, zero value otherwise.

### GetZeebeStatusOk

`func (o *ClusterStatus) GetZeebeStatusOk() (*ClusterHealthStatus, bool)`

GetZeebeStatusOk returns a tuple with the ZeebeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeebeStatus

`func (o *ClusterStatus) SetZeebeStatus(v ClusterHealthStatus)`

SetZeebeStatus sets ZeebeStatus field to given value.

### HasZeebeStatus

`func (o *ClusterStatus) HasZeebeStatus() bool`

HasZeebeStatus returns a boolean if a field has been set.

### GetReady

`func (o *ClusterStatus) GetReady() ClusterHealthStatus`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *ClusterStatus) GetReadyOk() (*ClusterHealthStatus, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *ClusterStatus) SetReady(v ClusterHealthStatus)`

SetReady sets Ready field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


