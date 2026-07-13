# SecureConnectivityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**SecureConnectivityDtoMetadata**](SecureConnectivityDtoMetadata.md) |  | 
**Spec** | [**SecureConnectivityDtoSpec**](SecureConnectivityDtoSpec.md) |  | 
**Status** | [**SecureConnectivityDtoStatus**](SecureConnectivityDtoStatus.md) |  | 

## Methods

### NewSecureConnectivityDto

`func NewSecureConnectivityDto(metadata SecureConnectivityDtoMetadata, spec SecureConnectivityDtoSpec, status SecureConnectivityDtoStatus, ) *SecureConnectivityDto`

NewSecureConnectivityDto instantiates a new SecureConnectivityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecureConnectivityDtoWithDefaults

`func NewSecureConnectivityDtoWithDefaults() *SecureConnectivityDto`

NewSecureConnectivityDtoWithDefaults instantiates a new SecureConnectivityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *SecureConnectivityDto) GetMetadata() SecureConnectivityDtoMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecureConnectivityDto) GetMetadataOk() (*SecureConnectivityDtoMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecureConnectivityDto) SetMetadata(v SecureConnectivityDtoMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *SecureConnectivityDto) GetSpec() SecureConnectivityDtoSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SecureConnectivityDto) GetSpecOk() (*SecureConnectivityDtoSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SecureConnectivityDto) SetSpec(v SecureConnectivityDtoSpec)`

SetSpec sets Spec field to given value.


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


