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
)

// checks if the PaginatedInstrumentData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedInstrumentData{}

// PaginatedInstrumentData struct for PaginatedInstrumentData
type PaginatedInstrumentData struct {
	Count *string `json:"count,omitempty"`
	Next *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results []InstrumentData `json:"results,omitempty"`
}

// NewPaginatedInstrumentData instantiates a new PaginatedInstrumentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedInstrumentData() *PaginatedInstrumentData {
	this := PaginatedInstrumentData{}
	return &this
}

// NewPaginatedInstrumentDataWithDefaults instantiates a new PaginatedInstrumentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedInstrumentDataWithDefaults() *PaginatedInstrumentData {
	this := PaginatedInstrumentData{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedInstrumentData) GetCount() string {
	if o == nil || IsNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedInstrumentData) GetCountOk() (*string, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedInstrumentData) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *PaginatedInstrumentData) SetCount(v string) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginatedInstrumentData) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedInstrumentData) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedInstrumentData) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *PaginatedInstrumentData) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *PaginatedInstrumentData) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedInstrumentData) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedInstrumentData) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *PaginatedInstrumentData) SetPrevious(v string) {
	o.Previous = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedInstrumentData) GetResults() []InstrumentData {
	if o == nil || IsNil(o.Results) {
		var ret []InstrumentData
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedInstrumentData) GetResultsOk() ([]InstrumentData, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedInstrumentData) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []InstrumentData and assigns it to the Results field.
func (o *PaginatedInstrumentData) SetResults(v []InstrumentData) {
	o.Results = v
}

func (o PaginatedInstrumentData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedInstrumentData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePaginatedInstrumentData struct {
	value *PaginatedInstrumentData
	isSet bool
}

func (v NullablePaginatedInstrumentData) Get() *PaginatedInstrumentData {
	return v.value
}

func (v *NullablePaginatedInstrumentData) Set(val *PaginatedInstrumentData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedInstrumentData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedInstrumentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedInstrumentData(val *PaginatedInstrumentData) *NullablePaginatedInstrumentData {
	return &NullablePaginatedInstrumentData{value: val, isSet: true}
}

func (v NullablePaginatedInstrumentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedInstrumentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


