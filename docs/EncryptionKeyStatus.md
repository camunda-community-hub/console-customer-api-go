# EncryptionKeyStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **string** | When the reason last changed, in ISO 8601. | [optional] 
**Message** | Pointer to **string** | The detail of the reason, for example the request that AWS KMS denied. | [optional] 
**Reason** | Pointer to **string** | Why the key is or is not usable. Only &#x60;Ready&#x60; means that Camunda can use the key and every grant on it. &#x60;KeyAccessDenied&#x60;, &#x60;KeyNotUsable&#x60; and &#x60;KeyNotEnabled&#x60; need an action on the key in AWS KMS. &#x60;Pending&#x60;, &#x60;ProviderUnavailable&#x60; and &#x60;InternalError&#x60; clear without an action. The reason can be a value that is not named here. | [optional] 

## Methods

### NewEncryptionKeyStatus

`func NewEncryptionKeyStatus() *EncryptionKeyStatus`

NewEncryptionKeyStatus instantiates a new EncryptionKeyStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionKeyStatusWithDefaults

`func NewEncryptionKeyStatusWithDefaults() *EncryptionKeyStatus`

NewEncryptionKeyStatusWithDefaults instantiates a new EncryptionKeyStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *EncryptionKeyStatus) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *EncryptionKeyStatus) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *EncryptionKeyStatus) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *EncryptionKeyStatus) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *EncryptionKeyStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EncryptionKeyStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EncryptionKeyStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EncryptionKeyStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *EncryptionKeyStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EncryptionKeyStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EncryptionKeyStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EncryptionKeyStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


