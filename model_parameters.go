/*
Camunda Management API

Manage Camunda Clusters and API Clients via API.

API version: 1.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Parameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Parameters{}

// Parameters Describes the different options to create a Camunda cluster.
type Parameters struct {
	Channels         []ParametersChannelsInner                        `json:"channels"`
	ClusterPlanTypes []ParametersChannelsInnerAllowedGenerationsInner `json:"clusterPlanTypes"`
	Regions          []ParametersChannelsInnerAllowedGenerationsInner `json:"regions"`
}

// NewParameters instantiates a new Parameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameters(channels []ParametersChannelsInner, clusterPlanTypes []ParametersChannelsInnerAllowedGenerationsInner, regions []ParametersChannelsInnerAllowedGenerationsInner) *Parameters {
	this := Parameters{}
	this.Channels = channels
	this.ClusterPlanTypes = clusterPlanTypes
	this.Regions = regions
	return &this
}

// NewParametersWithDefaults instantiates a new Parameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParametersWithDefaults() *Parameters {
	this := Parameters{}
	return &this
}

// GetChannels returns the Channels field value
func (o *Parameters) GetChannels() []ParametersChannelsInner {
	if o == nil {
		var ret []ParametersChannelsInner
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *Parameters) GetChannelsOk() ([]ParametersChannelsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channels, true
}

// SetChannels sets field value
func (o *Parameters) SetChannels(v []ParametersChannelsInner) {
	o.Channels = v
}

// GetClusterPlanTypes returns the ClusterPlanTypes field value
func (o *Parameters) GetClusterPlanTypes() []ParametersChannelsInnerAllowedGenerationsInner {
	if o == nil {
		var ret []ParametersChannelsInnerAllowedGenerationsInner
		return ret
	}

	return o.ClusterPlanTypes
}

// GetClusterPlanTypesOk returns a tuple with the ClusterPlanTypes field value
// and a boolean to check if the value has been set.
func (o *Parameters) GetClusterPlanTypesOk() ([]ParametersChannelsInnerAllowedGenerationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterPlanTypes, true
}

// SetClusterPlanTypes sets field value
func (o *Parameters) SetClusterPlanTypes(v []ParametersChannelsInnerAllowedGenerationsInner) {
	o.ClusterPlanTypes = v
}

// GetRegions returns the Regions field value
func (o *Parameters) GetRegions() []ParametersChannelsInnerAllowedGenerationsInner {
	if o == nil {
		var ret []ParametersChannelsInnerAllowedGenerationsInner
		return ret
	}

	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value
// and a boolean to check if the value has been set.
func (o *Parameters) GetRegionsOk() ([]ParametersChannelsInnerAllowedGenerationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regions, true
}

// SetRegions sets field value
func (o *Parameters) SetRegions(v []ParametersChannelsInnerAllowedGenerationsInner) {
	o.Regions = v
}

func (o Parameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Parameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channels"] = o.Channels
	toSerialize["clusterPlanTypes"] = o.ClusterPlanTypes
	toSerialize["regions"] = o.Regions
	return toSerialize, nil
}

type NullableParameters struct {
	value *Parameters
	isSet bool
}

func (v NullableParameters) Get() *Parameters {
	return v.value
}

func (v *NullableParameters) Set(val *Parameters) {
	v.value = val
	v.isSet = true
}

func (v NullableParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameters(val *Parameters) *NullableParameters {
	return &NullableParameters{value: val, isSet: true}
}

func (v NullableParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
