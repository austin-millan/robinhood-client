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

// checks if the WatchListsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchListsData{}

// WatchListsData struct for WatchListsData
type WatchListsData struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Instrument *string `json:"instrument,omitempty"`
	Url *string `json:"url,omitempty"`
	Watchlist *string `json:"watchlist,omitempty"`
}

// NewWatchListsData instantiates a new WatchListsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchListsData() *WatchListsData {
	this := WatchListsData{}
	return &this
}

// NewWatchListsDataWithDefaults instantiates a new WatchListsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchListsDataWithDefaults() *WatchListsData {
	this := WatchListsData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WatchListsData) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchListsData) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WatchListsData) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *WatchListsData) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetInstrument returns the Instrument field value if set, zero value otherwise.
func (o *WatchListsData) GetInstrument() string {
	if o == nil || IsNil(o.Instrument) {
		var ret string
		return ret
	}
	return *o.Instrument
}

// GetInstrumentOk returns a tuple with the Instrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchListsData) GetInstrumentOk() (*string, bool) {
	if o == nil || IsNil(o.Instrument) {
		return nil, false
	}
	return o.Instrument, true
}

// HasInstrument returns a boolean if a field has been set.
func (o *WatchListsData) HasInstrument() bool {
	if o != nil && !IsNil(o.Instrument) {
		return true
	}

	return false
}

// SetInstrument gets a reference to the given string and assigns it to the Instrument field.
func (o *WatchListsData) SetInstrument(v string) {
	o.Instrument = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WatchListsData) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchListsData) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WatchListsData) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WatchListsData) SetUrl(v string) {
	o.Url = &v
}

// GetWatchlist returns the Watchlist field value if set, zero value otherwise.
func (o *WatchListsData) GetWatchlist() string {
	if o == nil || IsNil(o.Watchlist) {
		var ret string
		return ret
	}
	return *o.Watchlist
}

// GetWatchlistOk returns a tuple with the Watchlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchListsData) GetWatchlistOk() (*string, bool) {
	if o == nil || IsNil(o.Watchlist) {
		return nil, false
	}
	return o.Watchlist, true
}

// HasWatchlist returns a boolean if a field has been set.
func (o *WatchListsData) HasWatchlist() bool {
	if o != nil && !IsNil(o.Watchlist) {
		return true
	}

	return false
}

// SetWatchlist gets a reference to the given string and assigns it to the Watchlist field.
func (o *WatchListsData) SetWatchlist(v string) {
	o.Watchlist = &v
}

func (o WatchListsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchListsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Instrument) {
		toSerialize["instrument"] = o.Instrument
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Watchlist) {
		toSerialize["watchlist"] = o.Watchlist
	}
	return toSerialize, nil
}

type NullableWatchListsData struct {
	value *WatchListsData
	isSet bool
}

func (v NullableWatchListsData) Get() *WatchListsData {
	return v.value
}

func (v *NullableWatchListsData) Set(val *WatchListsData) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchListsData) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchListsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchListsData(val *WatchListsData) *NullableWatchListsData {
	return &NullableWatchListsData{value: val, isSet: true}
}

func (v NullableWatchListsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchListsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


