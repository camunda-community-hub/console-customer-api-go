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

// OrganizationRoleVISITOR the model 'OrganizationRoleVISITOR'
type OrganizationRoleVISITOR string

// All allowed values of OrganizationRoleVISITOR enum
var AllowedOrganizationRoleVISITOREnumValues = []OrganizationRoleVISITOR{
	"visitor",
}

func (v *OrganizationRoleVISITOR) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrganizationRoleVISITOR(value)
	for _, existing := range AllowedOrganizationRoleVISITOREnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrganizationRoleVISITOR", value)
}

// NewOrganizationRoleVISITORFromValue returns a pointer to a valid OrganizationRoleVISITOR
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationRoleVISITORFromValue(v string) (*OrganizationRoleVISITOR, error) {
	ev := OrganizationRoleVISITOR(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationRoleVISITOR: valid values are %v", v, AllowedOrganizationRoleVISITOREnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationRoleVISITOR) IsValid() bool {
	for _, existing := range AllowedOrganizationRoleVISITOREnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationRole.VISITOR value
func (v OrganizationRoleVISITOR) Ptr() *OrganizationRoleVISITOR {
	return &v
}

type NullableOrganizationRoleVISITOR struct {
	value *OrganizationRoleVISITOR
	isSet bool
}

func (v NullableOrganizationRoleVISITOR) Get() *OrganizationRoleVISITOR {
	return v.value
}

func (v *NullableOrganizationRoleVISITOR) Set(val *OrganizationRoleVISITOR) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationRoleVISITOR) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationRoleVISITOR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationRoleVISITOR(val *OrganizationRoleVISITOR) *NullableOrganizationRoleVISITOR {
	return &NullableOrganizationRoleVISITOR{value: val, isSet: true}
}

func (v NullableOrganizationRoleVISITOR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationRoleVISITOR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
