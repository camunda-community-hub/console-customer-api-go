# SecureConnectivityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectivity** | [**SecureConnectivityDtoConnectivity**](SecureConnectivityDtoConnectivity.md) |  | 
**Status** | [**SecureConnectivityDtoStatus**](SecureConnectivityDtoStatus.md) |  | 

## Methods

### NewSecureConnectivityDto

`func NewSecureConnectivityDto(connectivity SecureConnectivityDtoConnectivity, status SecureConnectivityDtoStatus, ) *SecureConnectivityDto`

NewSecureConnectivityDto instantiates a new SecureConnectivityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecureConnectivityDtoWithDefaults

`func NewSecureConnectivityDtoWithDefaults() *SecureConnectivityDto`

NewSecureConnectivityDtoWithDefaults instantiates a new SecureConnectivityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectivity

`func (o *SecureConnectivityDto) GetConnectivity() SecureConnectivityDtoConnectivity`

GetConnectivity returns the Connectivity field if non-nil, zero value otherwise.

### GetConnectivityOk

`func (o *SecureConnectivityDto) GetConnectivityOk() (*SecureConnectivityDtoConnectivity, bool)`

GetConnectivityOk returns a tuple with the Connectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivity

`func (o *SecureConnectivityDto) SetConnectivity(v SecureConnectivityDtoConnectivity)`

SetConnectivity sets Connectivity field to given value.


### GetStatus

`func (o *SecureConnectivityDto) GetStatus() SecureConnectivityDtoStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecureConnectivityDto) GetStatusOk() (*SecureConnectivityDtoStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecureConnectivityDto) SetStatus(v SecureConnectivityDtoStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


