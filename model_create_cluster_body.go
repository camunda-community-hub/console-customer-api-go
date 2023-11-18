/*
Camunda Management API

Manage Camunda Clusters and API Clients via API.

API version: 1.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateClusterBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateClusterBody{}

// CreateClusterBody Used to create a cluster through this API. All necessary values can be retrieved from the <pre>/clusters/parameters</pre> endpoint.
type CreateClusterBody struct {
	// If set to false, no automatic updates will be performed on your cluster.
	AutoUpdate *bool `json:"autoUpdate,omitempty"`
	// The channel (software spec) to use.
	ChannelId string `json:"channelId"`
	// The generation (software version) to use.
	GenerationId string `json:"generationId"`
	// The name of your new cluster.
	Name string `json:"name"`
	// The planType (hardware spec) to use.
	PlanTypeId string `json:"planTypeId"`
	// The data center to use.
	RegionId string `json:"regionId"`
}

type _CreateClusterBody CreateClusterBody

// NewCreateClusterBody instantiates a new CreateClusterBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterBody(channelId string, generationId string, name string, planTypeId string, regionId string) *CreateClusterBody {
	this := CreateClusterBody{}
	this.ChannelId = channelId
	this.GenerationId = generationId
	this.Name = name
	this.PlanTypeId = planTypeId
	this.RegionId = regionId
	return &this
}

// NewCreateClusterBodyWithDefaults instantiates a new CreateClusterBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterBodyWithDefaults() *CreateClusterBody {
	this := CreateClusterBody{}
	return &this
}

// GetAutoUpdate returns the AutoUpdate field value if set, zero value otherwise.
func (o *CreateClusterBody) GetAutoUpdate() bool {
	if o == nil || IsNil(o.AutoUpdate) {
		var ret bool
		return ret
	}
	return *o.AutoUpdate
}

// GetAutoUpdateOk returns a tuple with the AutoUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterBody) GetAutoUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoUpdate) {
		return nil, false
	}
	return o.AutoUpdate, true
}

// HasAutoUpdate returns a boolean if a field has been set.
func (o *CreateClusterBody) HasAutoUpdate() bool {
	if o != nil && !IsNil(o.AutoUpdate) {
		return true
	}

	return false
}

// SetAutoUpdate gets a reference to the given bool and assigns it to the AutoUpdate field.
func (o *CreateClusterBody) SetAutoUpdate(v bool) {
	o.AutoUpdate = &v
}

// GetChannelId returns the ChannelId field value
func (o *CreateClusterBody) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *CreateClusterBody) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *CreateClusterBody) SetChannelId(v string) {
	o.ChannelId = v
}

// GetGenerationId returns the GenerationId field value
func (o *CreateClusterBody) GetGenerationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GenerationId
}

// GetGenerationIdOk returns a tuple with the GenerationId field value
// and a boolean to check if the value has been set.
func (o *CreateClusterBody) GetGenerationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GenerationId, true
}

// SetGenerationId sets field value
func (o *CreateClusterBody) SetGenerationId(v string) {
	o.GenerationId = v
}

// GetName returns the Name field value
func (o *CreateClusterBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateClusterBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateClusterBody) SetName(v string) {
	o.Name = v
}

// GetPlanTypeId returns the PlanTypeId field value
func (o *CreateClusterBody) GetPlanTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanTypeId
}

// GetPlanTypeIdOk returns a tuple with the PlanTypeId field value
// and a boolean to check if the value has been set.
func (o *CreateClusterBody) GetPlanTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanTypeId, true
}

// SetPlanTypeId sets field value
func (o *CreateClusterBody) SetPlanTypeId(v string) {
	o.PlanTypeId = v
}

// GetRegionId returns the RegionId field value
func (o *CreateClusterBody) GetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *CreateClusterBody) GetRegionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *CreateClusterBody) SetRegionId(v string) {
	o.RegionId = v
}

func (o CreateClusterBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateClusterBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoUpdate) {
		toSerialize["autoUpdate"] = o.AutoUpdate
	}
	toSerialize["channelId"] = o.ChannelId
	toSerialize["generationId"] = o.GenerationId
	toSerialize["name"] = o.Name
	toSerialize["planTypeId"] = o.PlanTypeId
	toSerialize["regionId"] = o.RegionId
	return toSerialize, nil
}

func (o *CreateClusterBody) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channelId",
		"generationId",
		"name",
		"planTypeId",
		"regionId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateClusterBody := _CreateClusterBody{}

	err = json.Unmarshal(bytes, &varCreateClusterBody)

	if err != nil {
		return err
	}

	*o = CreateClusterBody(varCreateClusterBody)

	return err
}

type NullableCreateClusterBody struct {
	value *CreateClusterBody
	isSet bool
}

func (v NullableCreateClusterBody) Get() *CreateClusterBody {
	return v.value
}

func (v *NullableCreateClusterBody) Set(val *CreateClusterBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterBody(val *CreateClusterBody) *NullableCreateClusterBody {
	return &NullableCreateClusterBody{value: val, isSet: true}
}

func (v NullableCreateClusterBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


