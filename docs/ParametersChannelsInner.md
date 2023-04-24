# ParametersChannelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedGenerations** | [**[]ParametersChannelsInnerAllowedGenerationsInner**](ParametersChannelsInnerAllowedGenerationsInner.md) |  | 
**DefaultGeneration** | [**ParametersChannelsInnerAllowedGenerationsInner**](ParametersChannelsInnerAllowedGenerationsInner.md) |  | 
**Name** | **string** |  | 
**Uuid** | **string** |  | 

## Methods

### NewParametersChannelsInner

`func NewParametersChannelsInner(allowedGenerations []ParametersChannelsInnerAllowedGenerationsInner, defaultGeneration ParametersChannelsInnerAllowedGenerationsInner, name string, uuid string, ) *ParametersChannelsInner`

NewParametersChannelsInner instantiates a new ParametersChannelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParametersChannelsInnerWithDefaults

`func NewParametersChannelsInnerWithDefaults() *ParametersChannelsInner`

NewParametersChannelsInnerWithDefaults instantiates a new ParametersChannelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedGenerations

`func (o *ParametersChannelsInner) GetAllowedGenerations() []ParametersChannelsInnerAllowedGenerationsInner`

GetAllowedGenerations returns the AllowedGenerations field if non-nil, zero value otherwise.

### GetAllowedGenerationsOk

`func (o *ParametersChannelsInner) GetAllowedGenerationsOk() (*[]ParametersChannelsInnerAllowedGenerationsInner, bool)`

GetAllowedGenerationsOk returns a tuple with the AllowedGenerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGenerations

`func (o *ParametersChannelsInner) SetAllowedGenerations(v []ParametersChannelsInnerAllowedGenerationsInner)`

SetAllowedGenerations sets AllowedGenerations field to given value.


### GetDefaultGeneration

`func (o *ParametersChannelsInner) GetDefaultGeneration() ParametersChannelsInnerAllowedGenerationsInner`

GetDefaultGeneration returns the DefaultGeneration field if non-nil, zero value otherwise.

### GetDefaultGenerationOk

`func (o *ParametersChannelsInner) GetDefaultGenerationOk() (*ParametersChannelsInnerAllowedGenerationsInner, bool)`

GetDefaultGenerationOk returns a tuple with the DefaultGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGeneration

`func (o *ParametersChannelsInner) SetDefaultGeneration(v ParametersChannelsInnerAllowedGenerationsInner)`

SetDefaultGeneration sets DefaultGeneration field to given value.


### GetName

`func (o *ParametersChannelsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParametersChannelsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParametersChannelsInner) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *ParametersChannelsInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ParametersChannelsInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ParametersChannelsInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


