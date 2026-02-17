# GetSecureConnectivityStatus200ResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectivity** | [**SecureConnectivityDtoConnectivity**](SecureConnectivityDtoConnectivity.md) |  | 
**Status** | [**SecureConnectivityDtoStatus**](SecureConnectivityDtoStatus.md) |  | 

## Methods

### NewGetSecureConnectivityStatus200ResponseStatus

`func NewGetSecureConnectivityStatus200ResponseStatus(connectivity SecureConnectivityDtoConnectivity, status SecureConnectivityDtoStatus, ) *GetSecureConnectivityStatus200ResponseStatus`

NewGetSecureConnectivityStatus200ResponseStatus instantiates a new GetSecureConnectivityStatus200ResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecureConnectivityStatus200ResponseStatusWithDefaults

`func NewGetSecureConnectivityStatus200ResponseStatusWithDefaults() *GetSecureConnectivityStatus200ResponseStatus`

NewGetSecureConnectivityStatus200ResponseStatusWithDefaults instantiates a new GetSecureConnectivityStatus200ResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectivity

`func (o *GetSecureConnectivityStatus200ResponseStatus) GetConnectivity() SecureConnectivityDtoConnectivity`

GetConnectivity returns the Connectivity field if non-nil, zero value otherwise.

### GetConnectivityOk

`func (o *GetSecureConnectivityStatus200ResponseStatus) GetConnectivityOk() (*SecureConnectivityDtoConnectivity, bool)`

GetConnectivityOk returns a tuple with the Connectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivity

`func (o *GetSecureConnectivityStatus200ResponseStatus) SetConnectivity(v SecureConnectivityDtoConnectivity)`

SetConnectivity sets Connectivity field to given value.


### GetStatus

`func (o *GetSecureConnectivityStatus200ResponseStatus) GetStatus() SecureConnectivityDtoStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSecureConnectivityStatus200ResponseStatus) GetStatusOk() (*SecureConnectivityDtoStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSecureConnectivityStatus200ResponseStatus) SetStatus(v SecureConnectivityDtoStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


