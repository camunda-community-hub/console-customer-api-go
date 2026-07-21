# UpdateClusterEncryptionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryEncryptionKeyId** | **string** | AWS KMS key ARN for the cluster&#39;s primary region. Single-region key ARNs and multi-region key (&#x60;mrk-...&#x60;) ARNs are both accepted, as long as the ARN&#39;s region segment matches the cluster&#39;s region. | 
**SecondaryEncryptionKeyId** | Pointer to **string** | AWS KMS key ARN for the cluster&#39;s backup region. Required when the cluster&#39;s backup configuration is dual-region; must be omitted for single-region clusters. Must differ from &#x60;primaryEncryptionKeyId&#x60;. | [optional] 

## Methods

### NewUpdateClusterEncryptionBody

`func NewUpdateClusterEncryptionBody(primaryEncryptionKeyId string, ) *UpdateClusterEncryptionBody`

NewUpdateClusterEncryptionBody instantiates a new UpdateClusterEncryptionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterEncryptionBodyWithDefaults

`func NewUpdateClusterEncryptionBodyWithDefaults() *UpdateClusterEncryptionBody`

NewUpdateClusterEncryptionBodyWithDefaults instantiates a new UpdateClusterEncryptionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryEncryptionKeyId

`func (o *UpdateClusterEncryptionBody) GetPrimaryEncryptionKeyId() string`

GetPrimaryEncryptionKeyId returns the PrimaryEncryptionKeyId field if non-nil, zero value otherwise.

### GetPrimaryEncryptionKeyIdOk

`func (o *UpdateClusterEncryptionBody) GetPrimaryEncryptionKeyIdOk() (*string, bool)`

GetPrimaryEncryptionKeyIdOk returns a tuple with the PrimaryEncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEncryptionKeyId

`func (o *UpdateClusterEncryptionBody) SetPrimaryEncryptionKeyId(v string)`

SetPrimaryEncryptionKeyId sets PrimaryEncryptionKeyId field to given value.


### GetSecondaryEncryptionKeyId

`func (o *UpdateClusterEncryptionBody) GetSecondaryEncryptionKeyId() string`

GetSecondaryEncryptionKeyId returns the SecondaryEncryptionKeyId field if non-nil, zero value otherwise.

### GetSecondaryEncryptionKeyIdOk

`func (o *UpdateClusterEncryptionBody) GetSecondaryEncryptionKeyIdOk() (*string, bool)`

GetSecondaryEncryptionKeyIdOk returns a tuple with the SecondaryEncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryEncryptionKeyId

`func (o *UpdateClusterEncryptionBody) SetSecondaryEncryptionKeyId(v string)`

SetSecondaryEncryptionKeyId sets SecondaryEncryptionKeyId field to given value.

### HasSecondaryEncryptionKeyId

`func (o *UpdateClusterEncryptionBody) HasSecondaryEncryptionKeyId() bool`

HasSecondaryEncryptionKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


