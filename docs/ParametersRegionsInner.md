# ParametersRegionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backups** | Pointer to [**[]ParametersRegionsInnerBackupsInner**](ParametersRegionsInnerBackupsInner.md) |  | [optional] 
**Name** | **string** |  | 
**Uuid** | **string** |  | 

## Methods

### NewParametersRegionsInner

`func NewParametersRegionsInner(name string, uuid string, ) *ParametersRegionsInner`

NewParametersRegionsInner instantiates a new ParametersRegionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParametersRegionsInnerWithDefaults

`func NewParametersRegionsInnerWithDefaults() *ParametersRegionsInner`

NewParametersRegionsInnerWithDefaults instantiates a new ParametersRegionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackups

`func (o *ParametersRegionsInner) GetBackups() []ParametersRegionsInnerBackupsInner`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *ParametersRegionsInner) GetBackupsOk() (*[]ParametersRegionsInnerBackupsInner, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *ParametersRegionsInner) SetBackups(v []ParametersRegionsInnerBackupsInner)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *ParametersRegionsInner) HasBackups() bool`

HasBackups returns a boolean if a field has been set.

### GetName

`func (o *ParametersRegionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParametersRegionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParametersRegionsInner) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *ParametersRegionsInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ParametersRegionsInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ParametersRegionsInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


