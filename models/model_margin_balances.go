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

// checks if the MarginBalances type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginBalances{}

// MarginBalances struct for MarginBalances
type MarginBalances struct {
	Cash *string `json:"cash,omitempty"`
	CashAvailableForWithdrawal *string `json:"cash_available_for_withdrawal,omitempty"`
	CashHeldForOrders *string `json:"cash_held_for_orders,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	DayTradeBuyingPower *string `json:"day_trade_buying_power,omitempty"`
	DayTradeBuyingPowerHeldForOrders *string `json:"day_trade_buying_power_held_for_orders,omitempty"`
	DayTradeRatio *string `json:"day_trade_ratio,omitempty"`
	MarginLimit *string `json:"margin_limit,omitempty"`
	MarkedPatternDayTraderDate *time.Time `json:"marked_pattern_day_trader_date,omitempty"`
	OvernightBuyingPower *string `json:"overnight_buying_power,omitempty"`
	OvernightBuyingPowerHeldForOrders *string `json:"overnight_buying_power_held_for_orders,omitempty"`
	OvernightRatio *string `json:"overnight_ratio,omitempty"`
	UnallocatedMarginCash *string `json:"unallocated_margin_cash,omitempty"`
	UnclearedDeposits *string `json:"uncleared_deposits,omitempty"`
	UnsettledFunds *string `json:"unsettled_funds,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewMarginBalances instantiates a new MarginBalances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginBalances() *MarginBalances {
	this := MarginBalances{}
	return &this
}

// NewMarginBalancesWithDefaults instantiates a new MarginBalances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginBalancesWithDefaults() *MarginBalances {
	this := MarginBalances{}
	return &this
}

// GetCash returns the Cash field value if set, zero value otherwise.
func (o *MarginBalances) GetCash() string {
	if o == nil || IsNil(o.Cash) {
		var ret string
		return ret
	}
	return *o.Cash
}

// GetCashOk returns a tuple with the Cash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetCashOk() (*string, bool) {
	if o == nil || IsNil(o.Cash) {
		return nil, false
	}
	return o.Cash, true
}

// HasCash returns a boolean if a field has been set.
func (o *MarginBalances) HasCash() bool {
	if o != nil && !IsNil(o.Cash) {
		return true
	}

	return false
}

// SetCash gets a reference to the given string and assigns it to the Cash field.
func (o *MarginBalances) SetCash(v string) {
	o.Cash = &v
}

// GetCashAvailableForWithdrawal returns the CashAvailableForWithdrawal field value if set, zero value otherwise.
func (o *MarginBalances) GetCashAvailableForWithdrawal() string {
	if o == nil || IsNil(o.CashAvailableForWithdrawal) {
		var ret string
		return ret
	}
	return *o.CashAvailableForWithdrawal
}

// GetCashAvailableForWithdrawalOk returns a tuple with the CashAvailableForWithdrawal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetCashAvailableForWithdrawalOk() (*string, bool) {
	if o == nil || IsNil(o.CashAvailableForWithdrawal) {
		return nil, false
	}
	return o.CashAvailableForWithdrawal, true
}

// HasCashAvailableForWithdrawal returns a boolean if a field has been set.
func (o *MarginBalances) HasCashAvailableForWithdrawal() bool {
	if o != nil && !IsNil(o.CashAvailableForWithdrawal) {
		return true
	}

	return false
}

// SetCashAvailableForWithdrawal gets a reference to the given string and assigns it to the CashAvailableForWithdrawal field.
func (o *MarginBalances) SetCashAvailableForWithdrawal(v string) {
	o.CashAvailableForWithdrawal = &v
}

// GetCashHeldForOrders returns the CashHeldForOrders field value if set, zero value otherwise.
func (o *MarginBalances) GetCashHeldForOrders() string {
	if o == nil || IsNil(o.CashHeldForOrders) {
		var ret string
		return ret
	}
	return *o.CashHeldForOrders
}

// GetCashHeldForOrdersOk returns a tuple with the CashHeldForOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetCashHeldForOrdersOk() (*string, bool) {
	if o == nil || IsNil(o.CashHeldForOrders) {
		return nil, false
	}
	return o.CashHeldForOrders, true
}

// HasCashHeldForOrders returns a boolean if a field has been set.
func (o *MarginBalances) HasCashHeldForOrders() bool {
	if o != nil && !IsNil(o.CashHeldForOrders) {
		return true
	}

	return false
}

// SetCashHeldForOrders gets a reference to the given string and assigns it to the CashHeldForOrders field.
func (o *MarginBalances) SetCashHeldForOrders(v string) {
	o.CashHeldForOrders = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MarginBalances) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MarginBalances) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MarginBalances) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDayTradeBuyingPower returns the DayTradeBuyingPower field value if set, zero value otherwise.
func (o *MarginBalances) GetDayTradeBuyingPower() string {
	if o == nil || IsNil(o.DayTradeBuyingPower) {
		var ret string
		return ret
	}
	return *o.DayTradeBuyingPower
}

// GetDayTradeBuyingPowerOk returns a tuple with the DayTradeBuyingPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetDayTradeBuyingPowerOk() (*string, bool) {
	if o == nil || IsNil(o.DayTradeBuyingPower) {
		return nil, false
	}
	return o.DayTradeBuyingPower, true
}

// HasDayTradeBuyingPower returns a boolean if a field has been set.
func (o *MarginBalances) HasDayTradeBuyingPower() bool {
	if o != nil && !IsNil(o.DayTradeBuyingPower) {
		return true
	}

	return false
}

// SetDayTradeBuyingPower gets a reference to the given string and assigns it to the DayTradeBuyingPower field.
func (o *MarginBalances) SetDayTradeBuyingPower(v string) {
	o.DayTradeBuyingPower = &v
}

// GetDayTradeBuyingPowerHeldForOrders returns the DayTradeBuyingPowerHeldForOrders field value if set, zero value otherwise.
func (o *MarginBalances) GetDayTradeBuyingPowerHeldForOrders() string {
	if o == nil || IsNil(o.DayTradeBuyingPowerHeldForOrders) {
		var ret string
		return ret
	}
	return *o.DayTradeBuyingPowerHeldForOrders
}

// GetDayTradeBuyingPowerHeldForOrdersOk returns a tuple with the DayTradeBuyingPowerHeldForOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetDayTradeBuyingPowerHeldForOrdersOk() (*string, bool) {
	if o == nil || IsNil(o.DayTradeBuyingPowerHeldForOrders) {
		return nil, false
	}
	return o.DayTradeBuyingPowerHeldForOrders, true
}

// HasDayTradeBuyingPowerHeldForOrders returns a boolean if a field has been set.
func (o *MarginBalances) HasDayTradeBuyingPowerHeldForOrders() bool {
	if o != nil && !IsNil(o.DayTradeBuyingPowerHeldForOrders) {
		return true
	}

	return false
}

// SetDayTradeBuyingPowerHeldForOrders gets a reference to the given string and assigns it to the DayTradeBuyingPowerHeldForOrders field.
func (o *MarginBalances) SetDayTradeBuyingPowerHeldForOrders(v string) {
	o.DayTradeBuyingPowerHeldForOrders = &v
}

// GetDayTradeRatio returns the DayTradeRatio field value if set, zero value otherwise.
func (o *MarginBalances) GetDayTradeRatio() string {
	if o == nil || IsNil(o.DayTradeRatio) {
		var ret string
		return ret
	}
	return *o.DayTradeRatio
}

// GetDayTradeRatioOk returns a tuple with the DayTradeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetDayTradeRatioOk() (*string, bool) {
	if o == nil || IsNil(o.DayTradeRatio) {
		return nil, false
	}
	return o.DayTradeRatio, true
}

// HasDayTradeRatio returns a boolean if a field has been set.
func (o *MarginBalances) HasDayTradeRatio() bool {
	if o != nil && !IsNil(o.DayTradeRatio) {
		return true
	}

	return false
}

// SetDayTradeRatio gets a reference to the given string and assigns it to the DayTradeRatio field.
func (o *MarginBalances) SetDayTradeRatio(v string) {
	o.DayTradeRatio = &v
}

// GetMarginLimit returns the MarginLimit field value if set, zero value otherwise.
func (o *MarginBalances) GetMarginLimit() string {
	if o == nil || IsNil(o.MarginLimit) {
		var ret string
		return ret
	}
	return *o.MarginLimit
}

// GetMarginLimitOk returns a tuple with the MarginLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetMarginLimitOk() (*string, bool) {
	if o == nil || IsNil(o.MarginLimit) {
		return nil, false
	}
	return o.MarginLimit, true
}

// HasMarginLimit returns a boolean if a field has been set.
func (o *MarginBalances) HasMarginLimit() bool {
	if o != nil && !IsNil(o.MarginLimit) {
		return true
	}

	return false
}

// SetMarginLimit gets a reference to the given string and assigns it to the MarginLimit field.
func (o *MarginBalances) SetMarginLimit(v string) {
	o.MarginLimit = &v
}

// GetMarkedPatternDayTraderDate returns the MarkedPatternDayTraderDate field value if set, zero value otherwise.
func (o *MarginBalances) GetMarkedPatternDayTraderDate() time.Time {
	if o == nil || IsNil(o.MarkedPatternDayTraderDate) {
		var ret time.Time
		return ret
	}
	return *o.MarkedPatternDayTraderDate
}

// GetMarkedPatternDayTraderDateOk returns a tuple with the MarkedPatternDayTraderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetMarkedPatternDayTraderDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MarkedPatternDayTraderDate) {
		return nil, false
	}
	return o.MarkedPatternDayTraderDate, true
}

// HasMarkedPatternDayTraderDate returns a boolean if a field has been set.
func (o *MarginBalances) HasMarkedPatternDayTraderDate() bool {
	if o != nil && !IsNil(o.MarkedPatternDayTraderDate) {
		return true
	}

	return false
}

// SetMarkedPatternDayTraderDate gets a reference to the given time.Time and assigns it to the MarkedPatternDayTraderDate field.
func (o *MarginBalances) SetMarkedPatternDayTraderDate(v time.Time) {
	o.MarkedPatternDayTraderDate = &v
}

// GetOvernightBuyingPower returns the OvernightBuyingPower field value if set, zero value otherwise.
func (o *MarginBalances) GetOvernightBuyingPower() string {
	if o == nil || IsNil(o.OvernightBuyingPower) {
		var ret string
		return ret
	}
	return *o.OvernightBuyingPower
}

// GetOvernightBuyingPowerOk returns a tuple with the OvernightBuyingPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetOvernightBuyingPowerOk() (*string, bool) {
	if o == nil || IsNil(o.OvernightBuyingPower) {
		return nil, false
	}
	return o.OvernightBuyingPower, true
}

// HasOvernightBuyingPower returns a boolean if a field has been set.
func (o *MarginBalances) HasOvernightBuyingPower() bool {
	if o != nil && !IsNil(o.OvernightBuyingPower) {
		return true
	}

	return false
}

// SetOvernightBuyingPower gets a reference to the given string and assigns it to the OvernightBuyingPower field.
func (o *MarginBalances) SetOvernightBuyingPower(v string) {
	o.OvernightBuyingPower = &v
}

// GetOvernightBuyingPowerHeldForOrders returns the OvernightBuyingPowerHeldForOrders field value if set, zero value otherwise.
func (o *MarginBalances) GetOvernightBuyingPowerHeldForOrders() string {
	if o == nil || IsNil(o.OvernightBuyingPowerHeldForOrders) {
		var ret string
		return ret
	}
	return *o.OvernightBuyingPowerHeldForOrders
}

// GetOvernightBuyingPowerHeldForOrdersOk returns a tuple with the OvernightBuyingPowerHeldForOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetOvernightBuyingPowerHeldForOrdersOk() (*string, bool) {
	if o == nil || IsNil(o.OvernightBuyingPowerHeldForOrders) {
		return nil, false
	}
	return o.OvernightBuyingPowerHeldForOrders, true
}

// HasOvernightBuyingPowerHeldForOrders returns a boolean if a field has been set.
func (o *MarginBalances) HasOvernightBuyingPowerHeldForOrders() bool {
	if o != nil && !IsNil(o.OvernightBuyingPowerHeldForOrders) {
		return true
	}

	return false
}

// SetOvernightBuyingPowerHeldForOrders gets a reference to the given string and assigns it to the OvernightBuyingPowerHeldForOrders field.
func (o *MarginBalances) SetOvernightBuyingPowerHeldForOrders(v string) {
	o.OvernightBuyingPowerHeldForOrders = &v
}

// GetOvernightRatio returns the OvernightRatio field value if set, zero value otherwise.
func (o *MarginBalances) GetOvernightRatio() string {
	if o == nil || IsNil(o.OvernightRatio) {
		var ret string
		return ret
	}
	return *o.OvernightRatio
}

// GetOvernightRatioOk returns a tuple with the OvernightRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetOvernightRatioOk() (*string, bool) {
	if o == nil || IsNil(o.OvernightRatio) {
		return nil, false
	}
	return o.OvernightRatio, true
}

// HasOvernightRatio returns a boolean if a field has been set.
func (o *MarginBalances) HasOvernightRatio() bool {
	if o != nil && !IsNil(o.OvernightRatio) {
		return true
	}

	return false
}

// SetOvernightRatio gets a reference to the given string and assigns it to the OvernightRatio field.
func (o *MarginBalances) SetOvernightRatio(v string) {
	o.OvernightRatio = &v
}

// GetUnallocatedMarginCash returns the UnallocatedMarginCash field value if set, zero value otherwise.
func (o *MarginBalances) GetUnallocatedMarginCash() string {
	if o == nil || IsNil(o.UnallocatedMarginCash) {
		var ret string
		return ret
	}
	return *o.UnallocatedMarginCash
}

// GetUnallocatedMarginCashOk returns a tuple with the UnallocatedMarginCash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetUnallocatedMarginCashOk() (*string, bool) {
	if o == nil || IsNil(o.UnallocatedMarginCash) {
		return nil, false
	}
	return o.UnallocatedMarginCash, true
}

// HasUnallocatedMarginCash returns a boolean if a field has been set.
func (o *MarginBalances) HasUnallocatedMarginCash() bool {
	if o != nil && !IsNil(o.UnallocatedMarginCash) {
		return true
	}

	return false
}

// SetUnallocatedMarginCash gets a reference to the given string and assigns it to the UnallocatedMarginCash field.
func (o *MarginBalances) SetUnallocatedMarginCash(v string) {
	o.UnallocatedMarginCash = &v
}

// GetUnclearedDeposits returns the UnclearedDeposits field value if set, zero value otherwise.
func (o *MarginBalances) GetUnclearedDeposits() string {
	if o == nil || IsNil(o.UnclearedDeposits) {
		var ret string
		return ret
	}
	return *o.UnclearedDeposits
}

// GetUnclearedDepositsOk returns a tuple with the UnclearedDeposits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetUnclearedDepositsOk() (*string, bool) {
	if o == nil || IsNil(o.UnclearedDeposits) {
		return nil, false
	}
	return o.UnclearedDeposits, true
}

// HasUnclearedDeposits returns a boolean if a field has been set.
func (o *MarginBalances) HasUnclearedDeposits() bool {
	if o != nil && !IsNil(o.UnclearedDeposits) {
		return true
	}

	return false
}

// SetUnclearedDeposits gets a reference to the given string and assigns it to the UnclearedDeposits field.
func (o *MarginBalances) SetUnclearedDeposits(v string) {
	o.UnclearedDeposits = &v
}

// GetUnsettledFunds returns the UnsettledFunds field value if set, zero value otherwise.
func (o *MarginBalances) GetUnsettledFunds() string {
	if o == nil || IsNil(o.UnsettledFunds) {
		var ret string
		return ret
	}
	return *o.UnsettledFunds
}

// GetUnsettledFundsOk returns a tuple with the UnsettledFunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetUnsettledFundsOk() (*string, bool) {
	if o == nil || IsNil(o.UnsettledFunds) {
		return nil, false
	}
	return o.UnsettledFunds, true
}

// HasUnsettledFunds returns a boolean if a field has been set.
func (o *MarginBalances) HasUnsettledFunds() bool {
	if o != nil && !IsNil(o.UnsettledFunds) {
		return true
	}

	return false
}

// SetUnsettledFunds gets a reference to the given string and assigns it to the UnsettledFunds field.
func (o *MarginBalances) SetUnsettledFunds(v string) {
	o.UnsettledFunds = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MarginBalances) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginBalances) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MarginBalances) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *MarginBalances) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o MarginBalances) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginBalances) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.DayTradeBuyingPower) {
		toSerialize["day_trade_buying_power"] = o.DayTradeBuyingPower
	}
	if !IsNil(o.DayTradeBuyingPowerHeldForOrders) {
		toSerialize["day_trade_buying_power_held_for_orders"] = o.DayTradeBuyingPowerHeldForOrders
	}
	if !IsNil(o.DayTradeRatio) {
		toSerialize["day_trade_ratio"] = o.DayTradeRatio
	}
	if !IsNil(o.MarginLimit) {
		toSerialize["margin_limit"] = o.MarginLimit
	}
	if !IsNil(o.MarkedPatternDayTraderDate) {
		toSerialize["marked_pattern_day_trader_date"] = o.MarkedPatternDayTraderDate
	}
	if !IsNil(o.OvernightBuyingPower) {
		toSerialize["overnight_buying_power"] = o.OvernightBuyingPower
	}
	if !IsNil(o.OvernightBuyingPowerHeldForOrders) {
		toSerialize["overnight_buying_power_held_for_orders"] = o.OvernightBuyingPowerHeldForOrders
	}
	if !IsNil(o.OvernightRatio) {
		toSerialize["overnight_ratio"] = o.OvernightRatio
	}
	if !IsNil(o.UnallocatedMarginCash) {
		toSerialize["unallocated_margin_cash"] = o.UnallocatedMarginCash
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

type NullableMarginBalances struct {
	value *MarginBalances
	isSet bool
}

func (v NullableMarginBalances) Get() *MarginBalances {
	return v.value
}

func (v *NullableMarginBalances) Set(val *MarginBalances) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginBalances) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginBalances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginBalances(val *MarginBalances) *NullableMarginBalances {
	return &NullableMarginBalances{value: val, isSet: true}
}

func (v NullableMarginBalances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginBalances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


