# Parameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | [**[]ParametersChannelsInner**](ParametersChannelsInner.md) |  | 
**ClusterPlanTypes** | [**[]ParametersChannelsInnerAllowedGenerationsInner**](ParametersChannelsInnerAllowedGenerationsInner.md) |  | 
**Regions** | [**[]ParametersRegionsInner**](ParametersRegionsInner.md) |  | 

## Methods

### NewParameters

`func NewParameters(channels []ParametersChannelsInner, clusterPlanTypes []ParametersChannelsInnerAllowedGenerationsInner, regions []ParametersRegionsInner, ) *Parameters`

NewParameters instantiates a new Parameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParametersWithDefaults

`func NewParametersWithDefaults() *Parameters`

NewParametersWithDefaults instantiates a new Parameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *Parameters) GetChannels() []ParametersChannelsInner`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *Parameters) GetChannelsOk() (*[]ParametersChannelsInner, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *Parameters) SetChannels(v []ParametersChannelsInner)`

SetChannels sets Channels field to given value.


### GetClusterPlanTypes

`func (o *Parameters) GetClusterPlanTypes() []ParametersChannelsInnerAllowedGenerationsInner`

GetClusterPlanTypes returns the ClusterPlanTypes field if non-nil, zero value otherwise.

### GetClusterPlanTypesOk

`func (o *Parameters) GetClusterPlanTypesOk() (*[]ParametersChannelsInnerAllowedGenerationsInner, bool)`

GetClusterPlanTypesOk returns a tuple with the ClusterPlanTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPlanTypes

`func (o *Parameters) SetClusterPlanTypes(v []ParametersChannelsInnerAllowedGenerationsInner)`

SetClusterPlanTypes sets ClusterPlanTypes field to given value.


### GetRegions

`func (o *Parameters) GetRegions() []ParametersRegionsInner`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Parameters) GetRegionsOk() (*[]ParametersRegionsInner, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Parameters) SetRegions(v []ParametersRegionsInner)`

SetRegions sets Regions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


