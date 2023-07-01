# MarginBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cash** | Pointer to **string** |  | [optional] 
**CashAvailableForWithdrawal** | Pointer to **string** |  | [optional] 
**CashHeldForOrders** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DayTradeBuyingPower** | Pointer to **string** |  | [optional] 
**DayTradeBuyingPowerHeldForOrders** | Pointer to **string** |  | [optional] 
**DayTradeRatio** | Pointer to **string** |  | [optional] 
**MarginLimit** | Pointer to **string** |  | [optional] 
**MarkedPatternDayTraderDate** | Pointer to **time.Time** |  | [optional] 
**OvernightBuyingPower** | Pointer to **string** |  | [optional] 
**OvernightBuyingPowerHeldForOrders** | Pointer to **string** |  | [optional] 
**OvernightRatio** | Pointer to **string** |  | [optional] 
**UnallocatedMarginCash** | Pointer to **string** |  | [optional] 
**UnclearedDeposits** | Pointer to **string** |  | [optional] 
**UnsettledFunds** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMarginBalances

`func NewMarginBalances() *MarginBalances`

NewMarginBalances instantiates a new MarginBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginBalancesWithDefaults

`func NewMarginBalancesWithDefaults() *MarginBalances`

NewMarginBalancesWithDefaults instantiates a new MarginBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCash

`func (o *MarginBalances) GetCash() string`

GetCash returns the Cash field if non-nil, zero value otherwise.

### GetCashOk

`func (o *MarginBalances) GetCashOk() (*string, bool)`

GetCashOk returns a tuple with the Cash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCash

`func (o *MarginBalances) SetCash(v string)`

SetCash sets Cash field to given value.

### HasCash

`func (o *MarginBalances) HasCash() bool`

HasCash returns a boolean if a field has been set.

### GetCashAvailableForWithdrawal

`func (o *MarginBalances) GetCashAvailableForWithdrawal() string`

GetCashAvailableForWithdrawal returns the CashAvailableForWithdrawal field if non-nil, zero value otherwise.

### GetCashAvailableForWithdrawalOk

`func (o *MarginBalances) GetCashAvailableForWithdrawalOk() (*string, bool)`

GetCashAvailableForWithdrawalOk returns a tuple with the CashAvailableForWithdrawal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAvailableForWithdrawal

`func (o *MarginBalances) SetCashAvailableForWithdrawal(v string)`

SetCashAvailableForWithdrawal sets CashAvailableForWithdrawal field to given value.

### HasCashAvailableForWithdrawal

`func (o *MarginBalances) HasCashAvailableForWithdrawal() bool`

HasCashAvailableForWithdrawal returns a boolean if a field has been set.

### GetCashHeldForOrders

`func (o *MarginBalances) GetCashHeldForOrders() string`

GetCashHeldForOrders returns the CashHeldForOrders field if non-nil, zero value otherwise.

### GetCashHeldForOrdersOk

`func (o *MarginBalances) GetCashHeldForOrdersOk() (*string, bool)`

GetCashHeldForOrdersOk returns a tuple with the CashHeldForOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashHeldForOrders

`func (o *MarginBalances) SetCashHeldForOrders(v string)`

SetCashHeldForOrders sets CashHeldForOrders field to given value.

### HasCashHeldForOrders

`func (o *MarginBalances) HasCashHeldForOrders() bool`

HasCashHeldForOrders returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarginBalances) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarginBalances) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarginBalances) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarginBalances) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDayTradeBuyingPower

`func (o *MarginBalances) GetDayTradeBuyingPower() string`

GetDayTradeBuyingPower returns the DayTradeBuyingPower field if non-nil, zero value otherwise.

### GetDayTradeBuyingPowerOk

`func (o *MarginBalances) GetDayTradeBuyingPowerOk() (*string, bool)`

GetDayTradeBuyingPowerOk returns a tuple with the DayTradeBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradeBuyingPower

`func (o *MarginBalances) SetDayTradeBuyingPower(v string)`

SetDayTradeBuyingPower sets DayTradeBuyingPower field to given value.

### HasDayTradeBuyingPower

`func (o *MarginBalances) HasDayTradeBuyingPower() bool`

HasDayTradeBuyingPower returns a boolean if a field has been set.

### GetDayTradeBuyingPowerHeldForOrders

`func (o *MarginBalances) GetDayTradeBuyingPowerHeldForOrders() string`

GetDayTradeBuyingPowerHeldForOrders returns the DayTradeBuyingPowerHeldForOrders field if non-nil, zero value otherwise.

### GetDayTradeBuyingPowerHeldForOrdersOk

`func (o *MarginBalances) GetDayTradeBuyingPowerHeldForOrdersOk() (*string, bool)`

GetDayTradeBuyingPowerHeldForOrdersOk returns a tuple with the DayTradeBuyingPowerHeldForOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradeBuyingPowerHeldForOrders

`func (o *MarginBalances) SetDayTradeBuyingPowerHeldForOrders(v string)`

SetDayTradeBuyingPowerHeldForOrders sets DayTradeBuyingPowerHeldForOrders field to given value.

### HasDayTradeBuyingPowerHeldForOrders

`func (o *MarginBalances) HasDayTradeBuyingPowerHeldForOrders() bool`

HasDayTradeBuyingPowerHeldForOrders returns a boolean if a field has been set.

### GetDayTradeRatio

`func (o *MarginBalances) GetDayTradeRatio() string`

GetDayTradeRatio returns the DayTradeRatio field if non-nil, zero value otherwise.

### GetDayTradeRatioOk

`func (o *MarginBalances) GetDayTradeRatioOk() (*string, bool)`

GetDayTradeRatioOk returns a tuple with the DayTradeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradeRatio

`func (o *MarginBalances) SetDayTradeRatio(v string)`

SetDayTradeRatio sets DayTradeRatio field to given value.

### HasDayTradeRatio

`func (o *MarginBalances) HasDayTradeRatio() bool`

HasDayTradeRatio returns a boolean if a field has been set.

### GetMarginLimit

`func (o *MarginBalances) GetMarginLimit() string`

GetMarginLimit returns the MarginLimit field if non-nil, zero value otherwise.

### GetMarginLimitOk

`func (o *MarginBalances) GetMarginLimitOk() (*string, bool)`

GetMarginLimitOk returns a tuple with the MarginLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLimit

`func (o *MarginBalances) SetMarginLimit(v string)`

SetMarginLimit sets MarginLimit field to given value.

### HasMarginLimit

`func (o *MarginBalances) HasMarginLimit() bool`

HasMarginLimit returns a boolean if a field has been set.

### GetMarkedPatternDayTraderDate

`func (o *MarginBalances) GetMarkedPatternDayTraderDate() time.Time`

GetMarkedPatternDayTraderDate returns the MarkedPatternDayTraderDate field if non-nil, zero value otherwise.

### GetMarkedPatternDayTraderDateOk

`func (o *MarginBalances) GetMarkedPatternDayTraderDateOk() (*time.Time, bool)`

GetMarkedPatternDayTraderDateOk returns a tuple with the MarkedPatternDayTraderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedPatternDayTraderDate

`func (o *MarginBalances) SetMarkedPatternDayTraderDate(v time.Time)`

SetMarkedPatternDayTraderDate sets MarkedPatternDayTraderDate field to given value.

### HasMarkedPatternDayTraderDate

`func (o *MarginBalances) HasMarkedPatternDayTraderDate() bool`

HasMarkedPatternDayTraderDate returns a boolean if a field has been set.

### GetOvernightBuyingPower

`func (o *MarginBalances) GetOvernightBuyingPower() string`

GetOvernightBuyingPower returns the OvernightBuyingPower field if non-nil, zero value otherwise.

### GetOvernightBuyingPowerOk

`func (o *MarginBalances) GetOvernightBuyingPowerOk() (*string, bool)`

GetOvernightBuyingPowerOk returns a tuple with the OvernightBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvernightBuyingPower

`func (o *MarginBalances) SetOvernightBuyingPower(v string)`

SetOvernightBuyingPower sets OvernightBuyingPower field to given value.

### HasOvernightBuyingPower

`func (o *MarginBalances) HasOvernightBuyingPower() bool`

HasOvernightBuyingPower returns a boolean if a field has been set.

### GetOvernightBuyingPowerHeldForOrders

`func (o *MarginBalances) GetOvernightBuyingPowerHeldForOrders() string`

GetOvernightBuyingPowerHeldForOrders returns the OvernightBuyingPowerHeldForOrders field if non-nil, zero value otherwise.

### GetOvernightBuyingPowerHeldForOrdersOk

`func (o *MarginBalances) GetOvernightBuyingPowerHeldForOrdersOk() (*string, bool)`

GetOvernightBuyingPowerHeldForOrdersOk returns a tuple with the OvernightBuyingPowerHeldForOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvernightBuyingPowerHeldForOrders

`func (o *MarginBalances) SetOvernightBuyingPowerHeldForOrders(v string)`

SetOvernightBuyingPowerHeldForOrders sets OvernightBuyingPowerHeldForOrders field to given value.

### HasOvernightBuyingPowerHeldForOrders

`func (o *MarginBalances) HasOvernightBuyingPowerHeldForOrders() bool`

HasOvernightBuyingPowerHeldForOrders returns a boolean if a field has been set.

### GetOvernightRatio

`func (o *MarginBalances) GetOvernightRatio() string`

GetOvernightRatio returns the OvernightRatio field if non-nil, zero value otherwise.

### GetOvernightRatioOk

`func (o *MarginBalances) GetOvernightRatioOk() (*string, bool)`

GetOvernightRatioOk returns a tuple with the OvernightRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvernightRatio

`func (o *MarginBalances) SetOvernightRatio(v string)`

SetOvernightRatio sets OvernightRatio field to given value.

### HasOvernightRatio

`func (o *MarginBalances) HasOvernightRatio() bool`

HasOvernightRatio returns a boolean if a field has been set.

### GetUnallocatedMarginCash

`func (o *MarginBalances) GetUnallocatedMarginCash() string`

GetUnallocatedMarginCash returns the UnallocatedMarginCash field if non-nil, zero value otherwise.

### GetUnallocatedMarginCashOk

`func (o *MarginBalances) GetUnallocatedMarginCashOk() (*string, bool)`

GetUnallocatedMarginCashOk returns a tuple with the UnallocatedMarginCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnallocatedMarginCash

`func (o *MarginBalances) SetUnallocatedMarginCash(v string)`

SetUnallocatedMarginCash sets UnallocatedMarginCash field to given value.

### HasUnallocatedMarginCash

`func (o *MarginBalances) HasUnallocatedMarginCash() bool`

HasUnallocatedMarginCash returns a boolean if a field has been set.

### GetUnclearedDeposits

`func (o *MarginBalances) GetUnclearedDeposits() string`

GetUnclearedDeposits returns the UnclearedDeposits field if non-nil, zero value otherwise.

### GetUnclearedDepositsOk

`func (o *MarginBalances) GetUnclearedDepositsOk() (*string, bool)`

GetUnclearedDepositsOk returns a tuple with the UnclearedDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnclearedDeposits

`func (o *MarginBalances) SetUnclearedDeposits(v string)`

SetUnclearedDeposits sets UnclearedDeposits field to given value.

### HasUnclearedDeposits

`func (o *MarginBalances) HasUnclearedDeposits() bool`

HasUnclearedDeposits returns a boolean if a field has been set.

### GetUnsettledFunds

`func (o *MarginBalances) GetUnsettledFunds() string`

GetUnsettledFunds returns the UnsettledFunds field if non-nil, zero value otherwise.

### GetUnsettledFundsOk

`func (o *MarginBalances) GetUnsettledFundsOk() (*string, bool)`

GetUnsettledFundsOk returns a tuple with the UnsettledFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsettledFunds

`func (o *MarginBalances) SetUnsettledFunds(v string)`

SetUnsettledFunds sets UnsettledFunds field to given value.

### HasUnsettledFunds

`func (o *MarginBalances) HasUnsettledFunds() bool`

HasUnsettledFunds returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MarginBalances) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarginBalances) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarginBalances) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MarginBalances) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


