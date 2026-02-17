# ByomStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | **[]interface{}** |  | 
**MetricsEndpoint** | [**ByomStatusMetricsEndpoint**](ByomStatusMetricsEndpoint.md) |  | 

## Methods

### NewByomStatus

`func NewByomStatus(conditions []interface{}, metricsEndpoint ByomStatusMetricsEndpoint, ) *ByomStatus`

NewByomStatus instantiates a new ByomStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByomStatusWithDefaults

`func NewByomStatusWithDefaults() *ByomStatus`

NewByomStatusWithDefaults instantiates a new ByomStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *ByomStatus) GetConditions() []interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ByomStatus) GetConditionsOk() (*[]interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ByomStatus) SetConditions(v []interface{})`

SetConditions sets Conditions field to given value.


### GetMetricsEndpoint

`func (o *ByomStatus) GetMetricsEndpoint() ByomStatusMetricsEndpoint`

GetMetricsEndpoint returns the MetricsEndpoint field if non-nil, zero value otherwise.

### GetMetricsEndpointOk

`func (o *ByomStatus) GetMetricsEndpointOk() (*ByomStatusMetricsEndpoint, bool)`

GetMetricsEndpointOk returns a tuple with the MetricsEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEndpoint

`func (o *ByomStatus) SetMetricsEndpoint(v ByomStatusMetricsEndpoint)`

SetMetricsEndpoint sets MetricsEndpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


