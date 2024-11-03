# ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityStatus** | Pointer to [**ClusterStatus**](ClusterStatus.md) |  | [optional] 
**OperateStatus** | Pointer to [**ClusterComponentStatus**](ClusterComponentStatus.md) |  | [optional] 
**OptimizeStatus** | Pointer to [**ClusterComponentStatus**](ClusterComponentStatus.md) |  | [optional] 
**Ready** | [**ClusterComponentStatus**](ClusterComponentStatus.md) |  | 
**TasklistStatus** | Pointer to [**ClusterComponentStatus**](ClusterComponentStatus.md) |  | [optional] 
**ZeebeStatus** | Pointer to [**ClusterComponentStatus**](ClusterComponentStatus.md) |  | [optional] 

## Methods

### NewClusterStatus

`func NewClusterStatus(ready ClusterComponentStatus, ) *ClusterStatus`

NewClusterStatus instantiates a new ClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusWithDefaults

`func NewClusterStatusWithDefaults() *ClusterStatus`

NewClusterStatusWithDefaults instantiates a new ClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityStatus

`func (o *ClusterStatus) GetIdentityStatus() ClusterStatus`

GetIdentityStatus returns the IdentityStatus field if non-nil, zero value otherwise.

### GetIdentityStatusOk

`func (o *ClusterStatus) GetIdentityStatusOk() (*ClusterStatus, bool)`

GetIdentityStatusOk returns a tuple with the IdentityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStatus

`func (o *ClusterStatus) SetIdentityStatus(v ClusterStatus)`

SetIdentityStatus sets IdentityStatus field to given value.

### HasIdentityStatus

`func (o *ClusterStatus) HasIdentityStatus() bool`

HasIdentityStatus returns a boolean if a field has been set.

### GetOperateStatus

`func (o *ClusterStatus) GetOperateStatus() ClusterComponentStatus`

GetOperateStatus returns the OperateStatus field if non-nil, zero value otherwise.

### GetOperateStatusOk

`func (o *ClusterStatus) GetOperateStatusOk() (*ClusterComponentStatus, bool)`

GetOperateStatusOk returns a tuple with the OperateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperateStatus

`func (o *ClusterStatus) SetOperateStatus(v ClusterComponentStatus)`

SetOperateStatus sets OperateStatus field to given value.

### HasOperateStatus

`func (o *ClusterStatus) HasOperateStatus() bool`

HasOperateStatus returns a boolean if a field has been set.

### GetOptimizeStatus

`func (o *ClusterStatus) GetOptimizeStatus() ClusterComponentStatus`

GetOptimizeStatus returns the OptimizeStatus field if non-nil, zero value otherwise.

### GetOptimizeStatusOk

`func (o *ClusterStatus) GetOptimizeStatusOk() (*ClusterComponentStatus, bool)`

GetOptimizeStatusOk returns a tuple with the OptimizeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeStatus

`func (o *ClusterStatus) SetOptimizeStatus(v ClusterComponentStatus)`

SetOptimizeStatus sets OptimizeStatus field to given value.

### HasOptimizeStatus

`func (o *ClusterStatus) HasOptimizeStatus() bool`

HasOptimizeStatus returns a boolean if a field has been set.

### GetReady

`func (o *ClusterStatus) GetReady() ClusterComponentStatus`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *ClusterStatus) GetReadyOk() (*ClusterComponentStatus, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *ClusterStatus) SetReady(v ClusterComponentStatus)`

SetReady sets Ready field to given value.


### GetTasklistStatus

`func (o *ClusterStatus) GetTasklistStatus() ClusterComponentStatus`

GetTasklistStatus returns the TasklistStatus field if non-nil, zero value otherwise.

### GetTasklistStatusOk

`func (o *ClusterStatus) GetTasklistStatusOk() (*ClusterComponentStatus, bool)`

GetTasklistStatusOk returns a tuple with the TasklistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistStatus

`func (o *ClusterStatus) SetTasklistStatus(v ClusterComponentStatus)`

SetTasklistStatus sets TasklistStatus field to given value.

### HasTasklistStatus

`func (o *ClusterStatus) HasTasklistStatus() bool`

HasTasklistStatus returns a boolean if a field has been set.

### GetZeebeStatus

`func (o *ClusterStatus) GetZeebeStatus() ClusterComponentStatus`

GetZeebeStatus returns the ZeebeStatus field if non-nil, zero value otherwise.

### GetZeebeStatusOk

`func (o *ClusterStatus) GetZeebeStatusOk() (*ClusterComponentStatus, bool)`

GetZeebeStatusOk returns a tuple with the ZeebeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeebeStatus

`func (o *ClusterStatus) SetZeebeStatus(v ClusterComponentStatus)`

SetZeebeStatus sets ZeebeStatus field to given value.

### HasZeebeStatus

`func (o *ClusterStatus) HasZeebeStatus() bool`

HasZeebeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


