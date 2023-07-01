# AccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** |  | [optional] 
**BuyingPower** | Pointer to **string** |  | [optional] 
**Cash** | Pointer to **string** |  | [optional] 
**CashAvailableForWithdrawal** | Pointer to **string** |  | [optional] 
**CashBalances** | Pointer to [**CashBalances**](CashBalances.md) |  | [optional] 
**CashHeldForOrders** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deactivated** | Pointer to **bool** |  | [optional] 
**DepositHalted** | Pointer to **bool** |  | [optional] 
**MarginBalances** | Pointer to [**MarginBalances**](MarginBalances.md) |  | [optional] 
**MaxAchEarlyAccessAmount** | Pointer to **string** |  | [optional] 
**OnlyPositionClosingTrades** | Pointer to **bool** |  | [optional] 
**Portfolio** | Pointer to **string** |  | [optional] 
**Positions** | Pointer to **string** |  | [optional] 
**Sma** | Pointer to **string** |  | [optional] 
**SmaHeldForOrders** | Pointer to **string** |  | [optional] 
**SweepEnabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**UnclearedDeposits** | Pointer to **string** |  | [optional] 
**UnsettledFunds** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**WithdrawalHalted** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccountInfo

`func NewAccountInfo() *AccountInfo`

NewAccountInfo instantiates a new AccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInfoWithDefaults

`func NewAccountInfoWithDefaults() *AccountInfo`

NewAccountInfoWithDefaults instantiates a new AccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *AccountInfo) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountInfo) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountInfo) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountInfo) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetBuyingPower

`func (o *AccountInfo) GetBuyingPower() string`

GetBuyingPower returns the BuyingPower field if non-nil, zero value otherwise.

### GetBuyingPowerOk

`func (o *AccountInfo) GetBuyingPowerOk() (*string, bool)`

GetBuyingPowerOk returns a tuple with the BuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingPower

`func (o *AccountInfo) SetBuyingPower(v string)`

SetBuyingPower sets BuyingPower field to given value.

### HasBuyingPower

`func (o *AccountInfo) HasBuyingPower() bool`

HasBuyingPower returns a boolean if a field has been set.

### GetCash

`func (o *AccountInfo) GetCash() string`

GetCash returns the Cash field if non-nil, zero value otherwise.

### GetCashOk

`func (o *AccountInfo) GetCashOk() (*string, bool)`

GetCashOk returns a tuple with the Cash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCash

`func (o *AccountInfo) SetCash(v string)`

SetCash sets Cash field to given value.

### HasCash

`func (o *AccountInfo) HasCash() bool`

HasCash returns a boolean if a field has been set.

### GetCashAvailableForWithdrawal

`func (o *AccountInfo) GetCashAvailableForWithdrawal() string`

GetCashAvailableForWithdrawal returns the CashAvailableForWithdrawal field if non-nil, zero value otherwise.

### GetCashAvailableForWithdrawalOk

`func (o *AccountInfo) GetCashAvailableForWithdrawalOk() (*string, bool)`

GetCashAvailableForWithdrawalOk returns a tuple with the CashAvailableForWithdrawal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAvailableForWithdrawal

`func (o *AccountInfo) SetCashAvailableForWithdrawal(v string)`

SetCashAvailableForWithdrawal sets CashAvailableForWithdrawal field to given value.

### HasCashAvailableForWithdrawal

`func (o *AccountInfo) HasCashAvailableForWithdrawal() bool`

HasCashAvailableForWithdrawal returns a boolean if a field has been set.

### GetCashBalances

`func (o *AccountInfo) GetCashBalances() CashBalances`

GetCashBalances returns the CashBalances field if non-nil, zero value otherwise.

### GetCashBalancesOk

`func (o *AccountInfo) GetCashBalancesOk() (*CashBalances, bool)`

GetCashBalancesOk returns a tuple with the CashBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBalances

`func (o *AccountInfo) SetCashBalances(v CashBalances)`

SetCashBalances sets CashBalances field to given value.

### HasCashBalances

`func (o *AccountInfo) HasCashBalances() bool`

HasCashBalances returns a boolean if a field has been set.

### GetCashHeldForOrders

`func (o *AccountInfo) GetCashHeldForOrders() string`

GetCashHeldForOrders returns the CashHeldForOrders field if non-nil, zero value otherwise.

### GetCashHeldForOrdersOk

`func (o *AccountInfo) GetCashHeldForOrdersOk() (*string, bool)`

GetCashHeldForOrdersOk returns a tuple with the CashHeldForOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashHeldForOrders

`func (o *AccountInfo) SetCashHeldForOrders(v string)`

SetCashHeldForOrders sets CashHeldForOrders field to given value.

### HasCashHeldForOrders

`func (o *AccountInfo) HasCashHeldForOrders() bool`

HasCashHeldForOrders returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeactivated

`func (o *AccountInfo) GetDeactivated() bool`

GetDeactivated returns the Deactivated field if non-nil, zero value otherwise.

### GetDeactivatedOk

`func (o *AccountInfo) GetDeactivatedOk() (*bool, bool)`

GetDeactivatedOk returns a tuple with the Deactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivated

`func (o *AccountInfo) SetDeactivated(v bool)`

SetDeactivated sets Deactivated field to given value.

### HasDeactivated

`func (o *AccountInfo) HasDeactivated() bool`

HasDeactivated returns a boolean if a field has been set.

### GetDepositHalted

`func (o *AccountInfo) GetDepositHalted() bool`

GetDepositHalted returns the DepositHalted field if non-nil, zero value otherwise.

### GetDepositHaltedOk

`func (o *AccountInfo) GetDepositHaltedOk() (*bool, bool)`

GetDepositHaltedOk returns a tuple with the DepositHalted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositHalted

`func (o *AccountInfo) SetDepositHalted(v bool)`

SetDepositHalted sets DepositHalted field to given value.

### HasDepositHalted

`func (o *AccountInfo) HasDepositHalted() bool`

HasDepositHalted returns a boolean if a field has been set.

### GetMarginBalances

`func (o *AccountInfo) GetMarginBalances() MarginBalances`

GetMarginBalances returns the MarginBalances field if non-nil, zero value otherwise.

### GetMarginBalancesOk

`func (o *AccountInfo) GetMarginBalancesOk() (*MarginBalances, bool)`

GetMarginBalancesOk returns a tuple with the MarginBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBalances

`func (o *AccountInfo) SetMarginBalances(v MarginBalances)`

SetMarginBalances sets MarginBalances field to given value.

### HasMarginBalances

`func (o *AccountInfo) HasMarginBalances() bool`

HasMarginBalances returns a boolean if a field has been set.

### GetMaxAchEarlyAccessAmount

`func (o *AccountInfo) GetMaxAchEarlyAccessAmount() string`

GetMaxAchEarlyAccessAmount returns the MaxAchEarlyAccessAmount field if non-nil, zero value otherwise.

### GetMaxAchEarlyAccessAmountOk

`func (o *AccountInfo) GetMaxAchEarlyAccessAmountOk() (*string, bool)`

GetMaxAchEarlyAccessAmountOk returns a tuple with the MaxAchEarlyAccessAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAchEarlyAccessAmount

`func (o *AccountInfo) SetMaxAchEarlyAccessAmount(v string)`

SetMaxAchEarlyAccessAmount sets MaxAchEarlyAccessAmount field to given value.

### HasMaxAchEarlyAccessAmount

`func (o *AccountInfo) HasMaxAchEarlyAccessAmount() bool`

HasMaxAchEarlyAccessAmount returns a boolean if a field has been set.

### GetOnlyPositionClosingTrades

`func (o *AccountInfo) GetOnlyPositionClosingTrades() bool`

GetOnlyPositionClosingTrades returns the OnlyPositionClosingTrades field if non-nil, zero value otherwise.

### GetOnlyPositionClosingTradesOk

`func (o *AccountInfo) GetOnlyPositionClosingTradesOk() (*bool, bool)`

GetOnlyPositionClosingTradesOk returns a tuple with the OnlyPositionClosingTrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyPositionClosingTrades

`func (o *AccountInfo) SetOnlyPositionClosingTrades(v bool)`

SetOnlyPositionClosingTrades sets OnlyPositionClosingTrades field to given value.

### HasOnlyPositionClosingTrades

`func (o *AccountInfo) HasOnlyPositionClosingTrades() bool`

HasOnlyPositionClosingTrades returns a boolean if a field has been set.

### GetPortfolio

`func (o *AccountInfo) GetPortfolio() string`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *AccountInfo) GetPortfolioOk() (*string, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *AccountInfo) SetPortfolio(v string)`

SetPortfolio sets Portfolio field to given value.

### HasPortfolio

`func (o *AccountInfo) HasPortfolio() bool`

HasPortfolio returns a boolean if a field has been set.

### GetPositions

`func (o *AccountInfo) GetPositions() string`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *AccountInfo) GetPositionsOk() (*string, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *AccountInfo) SetPositions(v string)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *AccountInfo) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetSma

`func (o *AccountInfo) GetSma() string`

GetSma returns the Sma field if non-nil, zero value otherwise.

### GetSmaOk

`func (o *AccountInfo) GetSmaOk() (*string, bool)`

GetSmaOk returns a tuple with the Sma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSma

`func (o *AccountInfo) SetSma(v string)`

SetSma sets Sma field to given value.

### HasSma

`func (o *AccountInfo) HasSma() bool`

HasSma returns a boolean if a field has been set.

### GetSmaHeldForOrders

`func (o *AccountInfo) GetSmaHeldForOrders() string`

GetSmaHeldForOrders returns the SmaHeldForOrders field if non-nil, zero value otherwise.

### GetSmaHeldForOrdersOk

`func (o *AccountInfo) GetSmaHeldForOrdersOk() (*string, bool)`

GetSmaHeldForOrdersOk returns a tuple with the SmaHeldForOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmaHeldForOrders

`func (o *AccountInfo) SetSmaHeldForOrders(v string)`

SetSmaHeldForOrders sets SmaHeldForOrders field to given value.

### HasSmaHeldForOrders

`func (o *AccountInfo) HasSmaHeldForOrders() bool`

HasSmaHeldForOrders returns a boolean if a field has been set.

### GetSweepEnabled

`func (o *AccountInfo) GetSweepEnabled() bool`

GetSweepEnabled returns the SweepEnabled field if non-nil, zero value otherwise.

### GetSweepEnabledOk

`func (o *AccountInfo) GetSweepEnabledOk() (*bool, bool)`

GetSweepEnabledOk returns a tuple with the SweepEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweepEnabled

`func (o *AccountInfo) SetSweepEnabled(v bool)`

SetSweepEnabled sets SweepEnabled field to given value.

### HasSweepEnabled

`func (o *AccountInfo) HasSweepEnabled() bool`

HasSweepEnabled returns a boolean if a field has been set.

### GetType

`func (o *AccountInfo) GetType() AccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountInfo) GetTypeOk() (*AccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountInfo) SetType(v AccountType)`

SetType sets Type field to given value.

### HasType

`func (o *AccountInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnclearedDeposits

`func (o *AccountInfo) GetUnclearedDeposits() string`

GetUnclearedDeposits returns the UnclearedDeposits field if non-nil, zero value otherwise.

### GetUnclearedDepositsOk

`func (o *AccountInfo) GetUnclearedDepositsOk() (*string, bool)`

GetUnclearedDepositsOk returns a tuple with the UnclearedDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnclearedDeposits

`func (o *AccountInfo) SetUnclearedDeposits(v string)`

SetUnclearedDeposits sets UnclearedDeposits field to given value.

### HasUnclearedDeposits

`func (o *AccountInfo) HasUnclearedDeposits() bool`

HasUnclearedDeposits returns a boolean if a field has been set.

### GetUnsettledFunds

`func (o *AccountInfo) GetUnsettledFunds() string`

GetUnsettledFunds returns the UnsettledFunds field if non-nil, zero value otherwise.

### GetUnsettledFundsOk

`func (o *AccountInfo) GetUnsettledFundsOk() (*string, bool)`

GetUnsettledFundsOk returns a tuple with the UnsettledFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsettledFunds

`func (o *AccountInfo) SetUnsettledFunds(v string)`

SetUnsettledFunds sets UnsettledFunds field to given value.

### HasUnsettledFunds

`func (o *AccountInfo) HasUnsettledFunds() bool`

HasUnsettledFunds returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccountInfo) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountInfo) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountInfo) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountInfo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AccountInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AccountInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AccountInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AccountInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUser

`func (o *AccountInfo) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AccountInfo) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AccountInfo) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AccountInfo) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWithdrawalHalted

`func (o *AccountInfo) GetWithdrawalHalted() bool`

GetWithdrawalHalted returns the WithdrawalHalted field if non-nil, zero value otherwise.

### GetWithdrawalHaltedOk

`func (o *AccountInfo) GetWithdrawalHaltedOk() (*bool, bool)`

GetWithdrawalHaltedOk returns a tuple with the WithdrawalHalted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalHalted

`func (o *AccountInfo) SetWithdrawalHalted(v bool)`

SetWithdrawalHalted sets WithdrawalHalted field to given value.

### HasWithdrawalHalted

`func (o *AccountInfo) HasWithdrawalHalted() bool`

HasWithdrawalHalted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


