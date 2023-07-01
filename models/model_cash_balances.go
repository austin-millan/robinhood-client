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
	"time"
)

// checks if the CashBalances type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashBalances{}

// CashBalances struct for CashBalances
type CashBalances struct {
	BuyingPower *string `json:"buying_power,omitempty"`
	Cash *string `json:"cash,omitempty"`
	CashAvailableForWithdrawal *string `json:"cash_available_for_withdrawal,omitempty"`
	CashHeldForOrders *string `json:"cash_held_for_orders,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UnclearedDeposits *string `json:"uncleared_deposits,omitempty"`
	UnsettledFunds *string `json:"unsettled_funds,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewCashBalances instantiates a new CashBalances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashBalances() *CashBalances {
	this := CashBalances{}
	return &this
}

// NewCashBalancesWithDefaults instantiates a new CashBalances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashBalancesWithDefaults() *CashBalances {
	this := CashBalances{}
	return &this
}

// GetBuyingPower returns the BuyingPower field value if set, zero value otherwise.
func (o *CashBalances) GetBuyingPower() string {
	if o == nil || IsNil(o.BuyingPower) {
		var ret string
		return ret
	}
	return *o.BuyingPower
}

// GetBuyingPowerOk returns a tuple with the BuyingPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashBalances) GetBuyingPowerOk() (*string, bool) {
	if o == nil || IsNil(o.BuyingPower) {
		return nil, false
	}
	return o.BuyingPower, true
}

// HasBuyingPower returns a boolean if a field has been set.
func (o *CashBalances) HasBuyingPower() bool {
	if o != nil && !IsNil(o.BuyingPower) {
		return true
	}

	return false
}

// SetBuyingPower gets a reference to the given string and assigns it to the BuyingPower field.
func (o *CashBalances) SetBuyingPower(v string) {
	o.BuyingPower = &v
}

// GetCash returns the Cash field value if set, zero value otherwise.
func (o *CashBalances) GetCash() string {
	if o == nil || IsNil(o.Cash) {
		var ret string
		return ret
	}
	return *o.Cash
}

// GetCashOk returns a tuple with the Cash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashBalances) GetCashOk() (*string, bool) {
	if o == nil || IsNil(o.Cash) {
		return nil, false
	}
	return o.Cash, true
}

// HasCash returns a boolean if a field has been set.
func (o *CashBalances) HasCash() bool {
	if o != nil && !IsNil(o.Cash) {
		return true
	}

	return false
}

// SetCash gets a reference to the given string and assigns it to the Cash field.
func (o *CashBalances) SetCash(v string) {
	o.Cash = &v
}

// GetCashAvailableForWithdrawal returns the CashAvailableForWithdrawal field value if set, zero value otherwise.
func (o *CashBalances) GetCashAvailableForWithdrawal() string {
	if o == nil || IsNil(o.CashAvailableForWithdrawal) {
		var ret string
		return ret
	}
	return *o.CashAvailableForWithdrawal
}

// GetCashAvailableForWithdrawalOk returns a tuple with the CashAvailableForWithdrawal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashBalances) GetCashAvailableForWithdrawalOk() (*string, bool) {
	if o == nil || IsNil(o.CashAvailableForWithdrawal) {
		return nil, false
	}
	return o.CashAvailableForWithdrawal, true
}

// HasCashAvailableForWithdrawal returns a boolean if a field has been set.
func (o *CashBalances) HasCashAvailableForWithdrawal() bool {
	if o != nil && !IsNil(o.CashAvailableForWithdrawal) {
		return true
	}

	return false
}

// SetCashAvailableForWithdrawal gets a reference to the given string and assigns it to the CashAvailableForWithdrawal field.
func (o *CashBalances) SetCashAvailableForWithdrawal(v string) {
	o.CashAvailableForWithdrawal = &v
}

// GetCashHeldForOrders returns the CashHeldForOrders field value if set, zero value otherwise.
func (o *CashBalances) GetCashHeldForOrders() string {
	if o == nil || IsNil(o.CashHeldForOrders) {
		var ret string
		return ret
	}
	return *o.CashHeldForOrders
}

// GetCashHeldForOrdersOk returns a tuple with the CashHeldForOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashBalances) GetCashHeldForOrdersOk() (*string, bool) {
	if o == nil || IsNil(o.CashHeldForOrders) {
		return nil, false
	}
	return o.CashHeldForOrders, true
}

// HasCashHeldForOrders returns a boolean if a field has been set.
func (o *CashBalances) HasCashHeldForOrders() bool {
	if o != nil && !IsNil(o.CashHeldForOrders) {
		return true
	}

	return false
}

// SetCashHeldForOrders gets a reference to the given string and assigns it to the CashHeldForOrders field.
func (o *CashBalances) SetCashHeldForOrders(v string) {
	o.CashHeldForOrders = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CashBalances) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashBalances) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CashBalances) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CashBalances) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUnclearedDeposits returns the UnclearedDeposits field value if set, zero value otherwise.
func (o *CashBalances) GetUnclearedDeposits() string {
	if o == nil || IsNil(o.UnclearedDeposits) {
		var ret string
		return ret
	}
	return *o.UnclearedDeposits
}

// GetUnclearedDepositsOk returns a tuple with the UnclearedDeposits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashBalances) GetUnclearedDepositsOk() (*string, bool) {
	if o == nil || IsNil(o.UnclearedDeposits) {
		return nil, false
	}
	return o.UnclearedDeposits, true
}

// HasUnclearedDeposits returns a boolean if a field has been set.
func (o *CashBalances) HasUnclearedDeposits() bool {
	if o != nil && !IsNil(o.UnclearedDeposits) {
		return true
	}

	return false
}

// SetUnclearedDeposits gets a reference to the given string and assigns it to the UnclearedDeposits field.
func (o *CashBalances) SetUnclearedDeposits(v string) {
	o.UnclearedDeposits = &v
}

// GetUnsettledFunds returns the UnsettledFunds field value if set, zero value otherwise.
func (o *CashBalances) GetUnsettledFunds() string {
	if o == nil || IsNil(o.UnsettledFunds) {
		var ret string
		return ret
	}
	return *o.UnsettledFunds
}

// GetUnsettledFundsOk returns a tuple with the UnsettledFunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashBalances) GetUnsettledFundsOk() (*string, bool) {
	if o == nil || IsNil(o.UnsettledFunds) {
		return nil, false
	}
	return o.UnsettledFunds, true
}

// HasUnsettledFunds returns a boolean if a field has been set.
func (o *CashBalances) HasUnsettledFunds() bool {
	if o != nil && !IsNil(o.UnsettledFunds) {
		return true
	}

	return false
}

// SetUnsettledFunds gets a reference to the given string and assigns it to the UnsettledFunds field.
func (o *CashBalances) SetUnsettledFunds(v string) {
	o.UnsettledFunds = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CashBalances) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashBalances) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CashBalances) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CashBalances) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o CashBalances) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashBalances) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuyingPower) {
		toSerialize["buying_power"] = o.BuyingPower
	}
	if !IsNil(o.Cash) {
		toSerialize["cash"] = o.Cash
	}
	if !IsNil(o.CashAvailableForWithdrawal) {
		toSerialize["cash_available_for_withdrawal"] = o.CashAvailableForWithdrawal
	}
	if !IsNil(o.CashHeldForOrders) {
		toSerialize["cash_held_for_orders"] = o.CashHeldForOrders
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UnclearedDeposits) {
		toSerialize["uncleared_deposits"] = o.UnclearedDeposits
	}
	if !IsNil(o.UnsettledFunds) {
		toSerialize["unsettled_funds"] = o.UnsettledFunds
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableCashBalances struct {
	value *CashBalances
	isSet bool
}

func (v NullableCashBalances) Get() *CashBalances {
	return v.value
}

func (v *NullableCashBalances) Set(val *CashBalances) {
	v.value = val
	v.isSet = true
}

func (v NullableCashBalances) IsSet() bool {
	return v.isSet
}

func (v *NullableCashBalances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashBalances(val *CashBalances) *NullableCashBalances {
	return &NullableCashBalances{value: val, isSet: true}
}

func (v NullableCashBalances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashBalances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


