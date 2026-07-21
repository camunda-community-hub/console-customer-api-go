# SecureConnectivityDtoSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPrincipals** | **[]string** |  | 
**AllowedRegions** | **[]string** |  | 
**Cluster** | [**SecureConnectivityDtoSpecCluster**](SecureConnectivityDtoSpecCluster.md) |  | 

## Methods

### NewSecureConnectivityDtoSpec

`func NewSecureConnectivityDtoSpec(allowedPrincipals []string, allowedRegions []string, cluster SecureConnectivityDtoSpecCluster, ) *SecureConnectivityDtoSpec`

NewSecureConnectivityDtoSpec instantiates a new SecureConnectivityDtoSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecureConnectivityDtoSpecWithDefaults

`func NewSecureConnectivityDtoSpecWithDefaults() *SecureConnectivityDtoSpec`

NewSecureConnectivityDtoSpecWithDefaults instantiates a new SecureConnectivityDtoSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPrincipals

`func (o *SecureConnectivityDtoSpec) GetAllowedPrincipals() []string`

GetAllowedPrincipals returns the AllowedPrincipals field if non-nil, zero value otherwise.

### GetAllowedPrincipalsOk

`func (o *SecureConnectivityDtoSpec) GetAllowedPrincipalsOk() (*[]string, bool)`

GetAllowedPrincipalsOk returns a tuple with the AllowedPrincipals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPrincipals

`func (o *SecureConnectivityDtoSpec) SetAllowedPrincipals(v []string)`

SetAllowedPrincipals sets AllowedPrincipals field to given value.


### GetAllowedRegions

`func (o *SecureConnectivityDtoSpec) GetAllowedRegions() []string`

GetAllowedRegions returns the AllowedRegions field if non-nil, zero value otherwise.

### GetAllowedRegionsOk

`func (o *SecureConnectivityDtoSpec) GetAllowedRegionsOk() (*[]string, bool)`

GetAllowedRegionsOk returns a tuple with the AllowedRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRegions

`func (o *SecureConnectivityDtoSpec) SetAllowedRegions(v []string)`

SetAllowedRegions sets AllowedRegions field to given value.


### GetCluster

`func (o *SecureConnectivityDtoSpec) GetCluster() SecureConnectivityDtoSpecCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *SecureConnectivityDtoSpec) GetClusterOk() (*SecureConnectivityDtoSpecCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *SecureConnectivityDtoSpec) SetCluster(v SecureConnectivityDtoSpecCluster)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


