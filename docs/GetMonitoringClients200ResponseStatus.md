# GetMonitoringClients200ResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | **[]map[string]interface{}** |  | 
**MetricsEndpoint** | [**ByomStatusMetricsEndpoint**](ByomStatusMetricsEndpoint.md) |  | 

## Methods

### NewGetMonitoringClients200ResponseStatus

`func NewGetMonitoringClients200ResponseStatus(conditions []map[string]interface{}, metricsEndpoint ByomStatusMetricsEndpoint, ) *GetMonitoringClients200ResponseStatus`

NewGetMonitoringClients200ResponseStatus instantiates a new GetMonitoringClients200ResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMonitoringClients200ResponseStatusWithDefaults

`func NewGetMonitoringClients200ResponseStatusWithDefaults() *GetMonitoringClients200ResponseStatus`

NewGetMonitoringClients200ResponseStatusWithDefaults instantiates a new GetMonitoringClients200ResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *GetMonitoringClients200ResponseStatus) GetConditions() []map[string]interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *GetMonitoringClients200ResponseStatus) GetConditionsOk() (*[]map[string]interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *GetMonitoringClients200ResponseStatus) SetConditions(v []map[string]interface{})`

SetConditions sets Conditions field to given value.


### GetMetricsEndpoint

`func (o *GetMonitoringClients200ResponseStatus) GetMetricsEndpoint() ByomStatusMetricsEndpoint`

GetMetricsEndpoint returns the MetricsEndpoint field if non-nil, zero value otherwise.

### GetMetricsEndpointOk

`func (o *GetMonitoringClients200ResponseStatus) GetMetricsEndpointOk() (*ByomStatusMetricsEndpoint, bool)`

GetMetricsEndpointOk returns a tuple with the MetricsEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEndpoint

`func (o *GetMonitoringClients200ResponseStatus) SetMetricsEndpoint(v ByomStatusMetricsEndpoint)`

SetMetricsEndpoint sets MetricsEndpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


