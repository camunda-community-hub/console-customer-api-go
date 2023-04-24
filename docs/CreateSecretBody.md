# CreateSecretBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretName** | **string** |  | 
**SecretValue** | **string** |  | 

## Methods

### NewCreateSecretBody

`func NewCreateSecretBody(secretName string, secretValue string, ) *CreateSecretBody`

NewCreateSecretBody instantiates a new CreateSecretBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecretBodyWithDefaults

`func NewCreateSecretBodyWithDefaults() *CreateSecretBody`

NewCreateSecretBodyWithDefaults instantiates a new CreateSecretBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretName

`func (o *CreateSecretBody) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *CreateSecretBody) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *CreateSecretBody) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetSecretValue

`func (o *CreateSecretBody) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *CreateSecretBody) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *CreateSecretBody) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


