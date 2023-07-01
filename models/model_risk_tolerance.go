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

// RiskTolerance the model 'RiskTolerance'
type RiskTolerance string

// List of RiskTolerance
const (
	LOW_RISK_TOLERANCE RiskTolerance = "low_risk_tolerance"
	MED_RISK_TOLERANCE RiskTolerance = "med_risk_tolerance"
	HIGH_RISK_TOLERANCE RiskTolerance = "high_risk_tolerance"
)

// All allowed values of RiskTolerance enum
var AllowedRiskToleranceEnumValues = []RiskTolerance{
	"low_risk_tolerance",
	"med_risk_tolerance",
	"high_risk_tolerance",
}

func (v *RiskTolerance) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RiskTolerance(value)
	for _, existing := range AllowedRiskToleranceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RiskTolerance", value)
}

// NewRiskToleranceFromValue returns a pointer to a valid RiskTolerance
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRiskToleranceFromValue(v string) (*RiskTolerance, error) {
	ev := RiskTolerance(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RiskTolerance: valid values are %v", v, AllowedRiskToleranceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RiskTolerance) IsValid() bool {
	for _, existing := range AllowedRiskToleranceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RiskTolerance value
func (v RiskTolerance) Ptr() *RiskTolerance {
	return &v
}

type NullableRiskTolerance struct {
	value *RiskTolerance
	isSet bool
}

func (v NullableRiskTolerance) Get() *RiskTolerance {
	return v.value
}

func (v *NullableRiskTolerance) Set(val *RiskTolerance) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskTolerance) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskTolerance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskTolerance(val *RiskTolerance) *NullableRiskTolerance {
	return &NullableRiskTolerance{value: val, isSet: true}
}

func (v NullableRiskTolerance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskTolerance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

