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

// checks if the Watchlist type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Watchlist{}

// Watchlist struct for Watchlist
type Watchlist struct {
	Name *string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
	User *string `json:"user,omitempty"`
}

// NewWatchlist instantiates a new Watchlist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlist() *Watchlist {
	this := Watchlist{}
	return &this
}

// NewWatchlistWithDefaults instantiates a new Watchlist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistWithDefaults() *Watchlist {
	this := Watchlist{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Watchlist) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Watchlist) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Watchlist) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Watchlist) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Watchlist) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Watchlist) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Watchlist) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Watchlist) SetUrl(v string) {
	o.Url = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Watchlist) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Watchlist) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Watchlist) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *Watchlist) SetUser(v string) {
	o.User = &v
}

func (o Watchlist) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Watchlist) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableWatchlist struct {
	value *Watchlist
	isSet bool
}

func (v NullableWatchlist) Get() *Watchlist {
	return v.value
}

func (v *NullableWatchlist) Set(val *Watchlist) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlist) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlist(val *Watchlist) *NullableWatchlist {
	return &NullableWatchlist{value: val, isSet: true}
}

func (v NullableWatchlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


