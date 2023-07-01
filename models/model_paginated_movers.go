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

// checks if the PaginatedMovers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedMovers{}

// PaginatedMovers struct for PaginatedMovers
type PaginatedMovers struct {
	Count *string `json:"count,omitempty"`
	Next *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results []map[string]interface{} `json:"results,omitempty"`
}

// NewPaginatedMovers instantiates a new PaginatedMovers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedMovers() *PaginatedMovers {
	this := PaginatedMovers{}
	return &this
}

// NewPaginatedMoversWithDefaults instantiates a new PaginatedMovers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedMoversWithDefaults() *PaginatedMovers {
	this := PaginatedMovers{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedMovers) GetCount() string {
	if o == nil || IsNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedMovers) GetCountOk() (*string, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedMovers) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *PaginatedMovers) SetCount(v string) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginatedMovers) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedMovers) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedMovers) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *PaginatedMovers) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *PaginatedMovers) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedMovers) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedMovers) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *PaginatedMovers) SetPrevious(v string) {
	o.Previous = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedMovers) GetResults() []map[string]interface{} {
	if o == nil || IsNil(o.Results) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedMovers) GetResultsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedMovers) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []map[string]interface{} and assigns it to the Results field.
func (o *PaginatedMovers) SetResults(v []map[string]interface{}) {
	o.Results = v
}

func (o PaginatedMovers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedMovers) ToMap() (map[string]interface{}, error) {
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

type NullablePaginatedMovers struct {
	value *PaginatedMovers
	isSet bool
}

func (v NullablePaginatedMovers) Get() *PaginatedMovers {
	return v.value
}

func (v *NullablePaginatedMovers) Set(val *PaginatedMovers) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedMovers) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedMovers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedMovers(val *PaginatedMovers) *NullablePaginatedMovers {
	return &NullablePaginatedMovers{value: val, isSet: true}
}

func (v NullablePaginatedMovers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedMovers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


