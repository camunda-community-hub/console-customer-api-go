# SecureConnectivityDtoStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]SecureConnectivityDtoStatusConditionsInner**](SecureConnectivityDtoStatusConditionsInner.md) |  | 
**Endpoint** | [**SecureConnectivityDtoStatusEndpoint**](SecureConnectivityDtoStatusEndpoint.md) |  | 
**EndpointConnectionCount** | **float64** |  | 
**EndpointConnections** | Pointer to [**[]SecureConnectivityDtoStatusEndpointConnectionsInner**](SecureConnectivityDtoStatusEndpointConnectionsInner.md) |  | [optional] 
**ObservedGeneration** | **float64** |  | 
**Urls** | **interface{}** |  | 

## Methods

### NewSecureConnectivityDtoStatus

`func NewSecureConnectivityDtoStatus(conditions []SecureConnectivityDtoStatusConditionsInner, endpoint SecureConnectivityDtoStatusEndpoint, endpointConnectionCount float64, observedGeneration float64, urls interface{}, ) *SecureConnectivityDtoStatus`

NewSecureConnectivityDtoStatus instantiates a new SecureConnectivityDtoStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecureConnectivityDtoStatusWithDefaults

`func NewSecureConnectivityDtoStatusWithDefaults() *SecureConnectivityDtoStatus`

NewSecureConnectivityDtoStatusWithDefaults instantiates a new SecureConnectivityDtoStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *SecureConnectivityDtoStatus) GetConditions() []SecureConnectivityDtoStatusConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SecureConnectivityDtoStatus) GetConditionsOk() (*[]SecureConnectivityDtoStatusConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SecureConnectivityDtoStatus) SetConditions(v []SecureConnectivityDtoStatusConditionsInner)`

SetConditions sets Conditions field to given value.


### GetEndpoint

`func (o *SecureConnectivityDtoStatus) GetEndpoint() SecureConnectivityDtoStatusEndpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SecureConnectivityDtoStatus) GetEndpointOk() (*SecureConnectivityDtoStatusEndpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SecureConnectivityDtoStatus) SetEndpoint(v SecureConnectivityDtoStatusEndpoint)`

SetEndpoint sets Endpoint field to given value.


### GetEndpointConnectionCount

`func (o *SecureConnectivityDtoStatus) GetEndpointConnectionCount() float64`

GetEndpointConnectionCount returns the EndpointConnectionCount field if non-nil, zero value otherwise.

### GetEndpointConnectionCountOk

`func (o *SecureConnectivityDtoStatus) GetEndpointConnectionCountOk() (*float64, bool)`

GetEndpointConnectionCountOk returns a tuple with the EndpointConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointConnectionCount

`func (o *SecureConnectivityDtoStatus) SetEndpointConnectionCount(v float64)`

SetEndpointConnectionCount sets EndpointConnectionCount field to given value.


### GetEndpointConnections

`func (o *SecureConnectivityDtoStatus) GetEndpointConnections() []SecureConnectivityDtoStatusEndpointConnectionsInner`

GetEndpointConnections returns the EndpointConnections field if non-nil, zero value otherwise.

### GetEndpointConnectionsOk

`func (o *SecureConnectivityDtoStatus) GetEndpointConnectionsOk() (*[]SecureConnectivityDtoStatusEndpointConnectionsInner, bool)`

GetEndpointConnectionsOk returns a tuple with the EndpointConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointConnections

`func (o *SecureConnectivityDtoStatus) SetEndpointConnections(v []SecureConnectivityDtoStatusEndpointConnectionsInner)`

SetEndpointConnections sets EndpointConnections field to given value.

### HasEndpointConnections

`func (o *SecureConnectivityDtoStatus) HasEndpointConnections() bool`

HasEndpointConnections returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *SecureConnectivityDtoStatus) GetObservedGeneration() float64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *SecureConnectivityDtoStatus) GetObservedGenerationOk() (*float64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *SecureConnectivityDtoStatus) SetObservedGeneration(v float64)`

SetObservedGeneration sets ObservedGeneration field to given value.


### GetUrls

`func (o *SecureConnectivityDtoStatus) GetUrls() interface{}`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *SecureConnectivityDtoStatus) GetUrlsOk() (*interface{}, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *SecureConnectivityDtoStatus) SetUrls(v interface{})`

SetUrls sets Urls field to given value.


### SetUrlsNil

`func (o *SecureConnectivityDtoStatus) SetUrlsNil(b bool)`

 SetUrlsNil sets the value for Urls to be an explicit nil

### UnsetUrls
`func (o *SecureConnectivityDtoStatus) UnsetUrls()`

UnsetUrls ensures that no value is present for Urls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


