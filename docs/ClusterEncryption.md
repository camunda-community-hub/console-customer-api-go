# ClusterEncryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryKeyArn** | Pointer to **string** |  | [optional] 
**SecondaryKeyArn** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**EncryptionStatus**](EncryptionStatus.md) |  | [optional] 
**Type** | [**ClusterEncryptionKey**](ClusterEncryptionKey.md) |  | 

## Methods

### NewClusterEncryption

`func NewClusterEncryption(type_ ClusterEncryptionKey, ) *ClusterEncryption`

NewClusterEncryption instantiates a new ClusterEncryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEncryptionWithDefaults

`func NewClusterEncryptionWithDefaults() *ClusterEncryption`

NewClusterEncryptionWithDefaults instantiates a new ClusterEncryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryKeyArn

`func (o *ClusterEncryption) GetPrimaryKeyArn() string`

GetPrimaryKeyArn returns the PrimaryKeyArn field if non-nil, zero value otherwise.

### GetPrimaryKeyArnOk

`func (o *ClusterEncryption) GetPrimaryKeyArnOk() (*string, bool)`

GetPrimaryKeyArnOk returns a tuple with the PrimaryKeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKeyArn

`func (o *ClusterEncryption) SetPrimaryKeyArn(v string)`

SetPrimaryKeyArn sets PrimaryKeyArn field to given value.

### HasPrimaryKeyArn

`func (o *ClusterEncryption) HasPrimaryKeyArn() bool`

HasPrimaryKeyArn returns a boolean if a field has been set.

### GetSecondaryKeyArn

`func (o *ClusterEncryption) GetSecondaryKeyArn() string`

GetSecondaryKeyArn returns the SecondaryKeyArn field if non-nil, zero value otherwise.

### GetSecondaryKeyArnOk

`func (o *ClusterEncryption) GetSecondaryKeyArnOk() (*string, bool)`

GetSecondaryKeyArnOk returns a tuple with the SecondaryKeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryKeyArn

`func (o *ClusterEncryption) SetSecondaryKeyArn(v string)`

SetSecondaryKeyArn sets SecondaryKeyArn field to given value.

### HasSecondaryKeyArn

`func (o *ClusterEncryption) HasSecondaryKeyArn() bool`

HasSecondaryKeyArn returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterEncryption) GetStatus() EncryptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterEncryption) GetStatusOk() (*EncryptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterEncryption) SetStatus(v EncryptionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterEncryption) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ClusterEncryption) GetType() ClusterEncryptionKey`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterEncryption) GetTypeOk() (*ClusterEncryptionKey, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterEncryption) SetType(v ClusterEncryptionKey)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


