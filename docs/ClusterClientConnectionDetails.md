# ClusterClientConnectionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ZEEBE_ADDRESS** | **string** |  | 
**ZEEBE_CLIENT_ID** | **string** |  | 
**ZEEBE_AUTHORIZATION_SERVER_URL** | **string** |  | 

## Methods

### NewClusterClientConnectionDetails

`func NewClusterClientConnectionDetails(name string, zEEBEADDRESS string, zEEBECLIENTID string, zEEBEAUTHORIZATIONSERVERURL string, ) *ClusterClientConnectionDetails`

NewClusterClientConnectionDetails instantiates a new ClusterClientConnectionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterClientConnectionDetailsWithDefaults

`func NewClusterClientConnectionDetailsWithDefaults() *ClusterClientConnectionDetails`

NewClusterClientConnectionDetailsWithDefaults instantiates a new ClusterClientConnectionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterClientConnectionDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterClientConnectionDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterClientConnectionDetails) SetName(v string)`

SetName sets Name field to given value.


### GetZEEBE_ADDRESS

`func (o *ClusterClientConnectionDetails) GetZEEBE_ADDRESS() string`

GetZEEBE_ADDRESS returns the ZEEBE_ADDRESS field if non-nil, zero value otherwise.

### GetZEEBE_ADDRESSOk

`func (o *ClusterClientConnectionDetails) GetZEEBE_ADDRESSOk() (*string, bool)`

GetZEEBE_ADDRESSOk returns a tuple with the ZEEBE_ADDRESS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZEEBE_ADDRESS

`func (o *ClusterClientConnectionDetails) SetZEEBE_ADDRESS(v string)`

SetZEEBE_ADDRESS sets ZEEBE_ADDRESS field to given value.


### GetZEEBE_CLIENT_ID

`func (o *ClusterClientConnectionDetails) GetZEEBE_CLIENT_ID() string`

GetZEEBE_CLIENT_ID returns the ZEEBE_CLIENT_ID field if non-nil, zero value otherwise.

### GetZEEBE_CLIENT_IDOk

`func (o *ClusterClientConnectionDetails) GetZEEBE_CLIENT_IDOk() (*string, bool)`

GetZEEBE_CLIENT_IDOk returns a tuple with the ZEEBE_CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZEEBE_CLIENT_ID

`func (o *ClusterClientConnectionDetails) SetZEEBE_CLIENT_ID(v string)`

SetZEEBE_CLIENT_ID sets ZEEBE_CLIENT_ID field to given value.


### GetZEEBE_AUTHORIZATION_SERVER_URL

`func (o *ClusterClientConnectionDetails) GetZEEBE_AUTHORIZATION_SERVER_URL() string`

GetZEEBE_AUTHORIZATION_SERVER_URL returns the ZEEBE_AUTHORIZATION_SERVER_URL field if non-nil, zero value otherwise.

### GetZEEBE_AUTHORIZATION_SERVER_URLOk

`func (o *ClusterClientConnectionDetails) GetZEEBE_AUTHORIZATION_SERVER_URLOk() (*string, bool)`

GetZEEBE_AUTHORIZATION_SERVER_URLOk returns a tuple with the ZEEBE_AUTHORIZATION_SERVER_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZEEBE_AUTHORIZATION_SERVER_URL

`func (o *ClusterClientConnectionDetails) SetZEEBE_AUTHORIZATION_SERVER_URL(v string)`

SetZEEBE_AUTHORIZATION_SERVER_URL sets ZEEBE_AUTHORIZATION_SERVER_URL field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


