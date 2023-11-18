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

// ClusterComponentStatus the model 'ClusterComponentStatus'
type ClusterComponentStatus string

// List of ClusterStatus
const (
	CLUSTERCOMPONENTSTATUS_HEALTHY   ClusterComponentStatus = "Healthy"
	CLUSTERCOMPONENTSTATUS_UNHEALTHY ClusterComponentStatus = "Unhealthy"
	CLUSTERCOMPONENTSTATUS_CREATING  ClusterComponentStatus = "Creating"
	CLUSTERCOMPONENTSTATUS_UPDATING  ClusterComponentStatus = "Updating"
)

// All allowed values of ClusterComponentStatus enum
var AllowedClusterComponentStatusEnumValues = []ClusterComponentStatus{
	"Healthy",
	"Unhealthy",
	"Creating",
	"Updating",
}

func (v *ClusterComponentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterComponentStatus(value)
	for _, existing := range AllowedClusterComponentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterComponentStatus", value)
}

// NewClusterComponentStatusFromValue returns a pointer to a valid ClusterComponentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterComponentStatusFromValue(v string) (*ClusterComponentStatus, error) {
	ev := ClusterComponentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterComponentStatus: valid values are %v", v, AllowedClusterComponentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterComponentStatus) IsValid() bool {
	for _, existing := range AllowedClusterComponentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterStatus value
func (v ClusterComponentStatus) Ptr() *ClusterComponentStatus {
	return &v
}

type NullableClusterComponentStatus struct {
	value *ClusterComponentStatus
	isSet bool
}

func (v NullableClusterComponentStatus) Get() *ClusterComponentStatus {
	return v.value
}

func (v *NullableClusterComponentStatus) Set(val *ClusterComponentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterComponentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterComponentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterComponentStatus(val *ClusterComponentStatus) *NullableClusterComponentStatus {
	return &NullableClusterComponentStatus{value: val, isSet: true}
}

func (v NullableClusterComponentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterComponentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
