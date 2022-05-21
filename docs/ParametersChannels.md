# ParametersChannels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedGenerations** | [**[]ParametersAllowedGenerations**](ParametersAllowedGenerations.md) |  | 
**DefaultGeneration** | [**ParametersAllowedGenerations**](ParametersAllowedGenerations.md) |  | 
**Name** | **string** |  | 
**Uuid** | **string** |  | 

## Methods

### NewParametersChannels

`func NewParametersChannels(allowedGenerations []ParametersAllowedGenerations, defaultGeneration ParametersAllowedGenerations, name string, uuid string, ) *ParametersChannels`

NewParametersChannels instantiates a new ParametersChannels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParametersChannelsWithDefaults

`func NewParametersChannelsWithDefaults() *ParametersChannels`

NewParametersChannelsWithDefaults instantiates a new ParametersChannels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedGenerations

`func (o *ParametersChannels) GetAllowedGenerations() []ParametersAllowedGenerations`

GetAllowedGenerations returns the AllowedGenerations field if non-nil, zero value otherwise.

### GetAllowedGenerationsOk

`func (o *ParametersChannels) GetAllowedGenerationsOk() (*[]ParametersAllowedGenerations, bool)`

GetAllowedGenerationsOk returns a tuple with the AllowedGenerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGenerations

`func (o *ParametersChannels) SetAllowedGenerations(v []ParametersAllowedGenerations)`

SetAllowedGenerations sets AllowedGenerations field to given value.


### GetDefaultGeneration

`func (o *ParametersChannels) GetDefaultGeneration() ParametersAllowedGenerations`

GetDefaultGeneration returns the DefaultGeneration field if non-nil, zero value otherwise.

### GetDefaultGenerationOk

`func (o *ParametersChannels) GetDefaultGenerationOk() (*ParametersAllowedGenerations, bool)`

GetDefaultGenerationOk returns a tuple with the DefaultGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGeneration

`func (o *ParametersChannels) SetDefaultGeneration(v ParametersAllowedGenerations)`

SetDefaultGeneration sets DefaultGeneration field to given value.


### GetName

`func (o *ParametersChannels) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParametersChannels) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParametersChannels) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *ParametersChannels) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ParametersChannels) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ParametersChannels) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


