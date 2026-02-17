# GetMonitoringClients200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | [**[]ByomClientDto**](ByomClientDto.md) |  | 
**Status** | [**GetMonitoringClients200ResponseStatus**](GetMonitoringClients200ResponseStatus.md) |  | 

## Methods

### NewGetMonitoringClients200Response

`func NewGetMonitoringClients200Response(clients []ByomClientDto, status GetMonitoringClients200ResponseStatus, ) *GetMonitoringClients200Response`

NewGetMonitoringClients200Response instantiates a new GetMonitoringClients200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMonitoringClients200ResponseWithDefaults

`func NewGetMonitoringClients200ResponseWithDefaults() *GetMonitoringClients200Response`

NewGetMonitoringClients200ResponseWithDefaults instantiates a new GetMonitoringClients200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *GetMonitoringClients200Response) GetClients() []ByomClientDto`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *GetMonitoringClients200Response) GetClientsOk() (*[]ByomClientDto, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *GetMonitoringClients200Response) SetClients(v []ByomClientDto)`

SetClients sets Clients field to given value.


### GetStatus

`func (o *GetMonitoringClients200Response) GetStatus() GetMonitoringClients200ResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMonitoringClients200Response) GetStatusOk() (*GetMonitoringClients200ResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMonitoringClients200Response) SetStatus(v GetMonitoringClients200ResponseStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


