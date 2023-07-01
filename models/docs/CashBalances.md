# CashBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyingPower** | Pointer to **string** |  | [optional] 
**Cash** | Pointer to **string** |  | [optional] 
**CashAvailableForWithdrawal** | Pointer to **string** |  | [optional] 
**CashHeldForOrders** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UnclearedDeposits** | Pointer to **string** |  | [optional] 
**UnsettledFunds** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCashBalances

`func NewCashBalances() *CashBalances`

NewCashBalances instantiates a new CashBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashBalancesWithDefaults

`func NewCashBalancesWithDefaults() *CashBalances`

NewCashBalancesWithDefaults instantiates a new CashBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyingPower

`func (o *CashBalances) GetBuyingPower() string`

GetBuyingPower returns the BuyingPower field if non-nil, zero value otherwise.

### GetBuyingPowerOk

`func (o *CashBalances) GetBuyingPowerOk() (*string, bool)`

GetBuyingPowerOk returns a tuple with the BuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingPower

`func (o *CashBalances) SetBuyingPower(v string)`

SetBuyingPower sets BuyingPower field to given value.

### HasBuyingPower

`func (o *CashBalances) HasBuyingPower() bool`

HasBuyingPower returns a boolean if a field has been set.

### GetCash

`func (o *CashBalances) GetCash() string`

GetCash returns the Cash field if non-nil, zero value otherwise.

### GetCashOk

`func (o *CashBalances) GetCashOk() (*string, bool)`

GetCashOk returns a tuple with the Cash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCash

`func (o *CashBalances) SetCash(v string)`

SetCash sets Cash field to given value.

### HasCash

`func (o *CashBalances) HasCash() bool`

HasCash returns a boolean if a field has been set.

### GetCashAvailableForWithdrawal

`func (o *CashBalances) GetCashAvailableForWithdrawal() string`

GetCashAvailableForWithdrawal returns the CashAvailableForWithdrawal field if non-nil, zero value otherwise.

### GetCashAvailableForWithdrawalOk

`func (o *CashBalances) GetCashAvailableForWithdrawalOk() (*string, bool)`

GetCashAvailableForWithdrawalOk returns a tuple with the CashAvailableForWithdrawal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAvailableForWithdrawal

`func (o *CashBalances) SetCashAvailableForWithdrawal(v string)`

SetCashAvailableForWithdrawal sets CashAvailableForWithdrawal field to given value.

### HasCashAvailableForWithdrawal

`func (o *CashBalances) HasCashAvailableForWithdrawal() bool`

HasCashAvailableForWithdrawal returns a boolean if a field has been set.

### GetCashHeldForOrders

`func (o *CashBalances) GetCashHeldForOrders() string`

GetCashHeldForOrders returns the CashHeldForOrders field if non-nil, zero value otherwise.

### GetCashHeldForOrdersOk

`func (o *CashBalances) GetCashHeldForOrdersOk() (*string, bool)`

GetCashHeldForOrdersOk returns a tuple with the CashHeldForOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashHeldForOrders

`func (o *CashBalances) SetCashHeldForOrders(v string)`

SetCashHeldForOrders sets CashHeldForOrders field to given value.

### HasCashHeldForOrders

`func (o *CashBalances) HasCashHeldForOrders() bool`

HasCashHeldForOrders returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CashBalances) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CashBalances) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CashBalances) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CashBalances) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUnclearedDeposits

`func (o *CashBalances) GetUnclearedDeposits() string`

GetUnclearedDeposits returns the UnclearedDeposits field if non-nil, zero value otherwise.

### GetUnclearedDepositsOk

`func (o *CashBalances) GetUnclearedDepositsOk() (*string, bool)`

GetUnclearedDepositsOk returns a tuple with the UnclearedDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnclearedDeposits

`func (o *CashBalances) SetUnclearedDeposits(v string)`

SetUnclearedDeposits sets UnclearedDeposits field to given value.

### HasUnclearedDeposits

`func (o *CashBalances) HasUnclearedDeposits() bool`

HasUnclearedDeposits returns a boolean if a field has been set.

### GetUnsettledFunds

`func (o *CashBalances) GetUnsettledFunds() string`

GetUnsettledFunds returns the UnsettledFunds field if non-nil, zero value otherwise.

### GetUnsettledFundsOk

`func (o *CashBalances) GetUnsettledFundsOk() (*string, bool)`

GetUnsettledFundsOk returns a tuple with the UnsettledFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsettledFunds

`func (o *CashBalances) SetUnsettledFunds(v string)`

SetUnsettledFunds sets UnsettledFunds field to given value.

### HasUnsettledFunds

`func (o *CashBalances) HasUnsettledFunds() bool`

HasUnsettledFunds returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CashBalances) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CashBalances) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CashBalances) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CashBalances) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


