/*
Robinhood API

Robinhood API Documentation

API version: 3.0.1
Contact: austin.millan@protonmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// MaritalStatus the model 'MaritalStatus'
type MaritalStatus string

// List of MaritalStatus
const (
	SINGLE MaritalStatus = "single"
	MARRIED MaritalStatus = "married"
)

// All allowed values of MaritalStatus enum
var AllowedMaritalStatusEnumValues = []MaritalStatus{
	"single",
	"married",
}

func (v *MaritalStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MaritalStatus(value)
	for _, existing := range AllowedMaritalStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MaritalStatus", value)
}

// NewMaritalStatusFromValue returns a pointer to a valid MaritalStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMaritalStatusFromValue(v string) (*MaritalStatus, error) {
	ev := MaritalStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MaritalStatus: valid values are %v", v, AllowedMaritalStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MaritalStatus) IsValid() bool {
	for _, existing := range AllowedMaritalStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MaritalStatus value
func (v MaritalStatus) Ptr() *MaritalStatus {
	return &v
}

type NullableMaritalStatus struct {
	value *MaritalStatus
	isSet bool
}

func (v NullableMaritalStatus) Get() *MaritalStatus {
	return v.value
}

func (v *NullableMaritalStatus) Set(val *MaritalStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMaritalStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMaritalStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaritalStatus(val *MaritalStatus) *NullableMaritalStatus {
	return &NullableMaritalStatus{value: val, isSet: true}
}

func (v NullableMaritalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaritalStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

