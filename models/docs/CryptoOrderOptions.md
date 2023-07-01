# CryptoOrderOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**AveragePrice** | Pointer to **string** |  | [optional] 
**Cancel** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CumulativeQuantity** | Pointer to **string** |  | [optional] 
**CurrencyPairId** | Pointer to **string** |  | [optional] 
**Executions** | Pointer to [**[]Execution**](Execution.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastTransactionAt** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**RejectReason** | Pointer to **string** |  | [optional] 
**Side** | Pointer to [**Side**](Side.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCryptoOrderOptions

`func NewCryptoOrderOptions() *CryptoOrderOptions`

NewCryptoOrderOptions instantiates a new CryptoOrderOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoOrderOptionsWithDefaults

`func NewCryptoOrderOptionsWithDefaults() *CryptoOrderOptions`

NewCryptoOrderOptionsWithDefaults instantiates a new CryptoOrderOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CryptoOrderOptions) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CryptoOrderOptions) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CryptoOrderOptions) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CryptoOrderOptions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAveragePrice

`func (o *CryptoOrderOptions) GetAveragePrice() string`

GetAveragePrice returns the AveragePrice field if non-nil, zero value otherwise.

### GetAveragePriceOk

`func (o *CryptoOrderOptions) GetAveragePriceOk() (*string, bool)`

GetAveragePriceOk returns a tuple with the AveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePrice

`func (o *CryptoOrderOptions) SetAveragePrice(v string)`

SetAveragePrice sets AveragePrice field to given value.

### HasAveragePrice

`func (o *CryptoOrderOptions) HasAveragePrice() bool`

HasAveragePrice returns a boolean if a field has been set.

### GetCancel

`func (o *CryptoOrderOptions) GetCancel() string`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *CryptoOrderOptions) GetCancelOk() (*string, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *CryptoOrderOptions) SetCancel(v string)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *CryptoOrderOptions) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CryptoOrderOptions) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CryptoOrderOptions) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CryptoOrderOptions) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CryptoOrderOptions) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCumulativeQuantity

`func (o *CryptoOrderOptions) GetCumulativeQuantity() string`

GetCumulativeQuantity returns the CumulativeQuantity field if non-nil, zero value otherwise.

### GetCumulativeQuantityOk

`func (o *CryptoOrderOptions) GetCumulativeQuantityOk() (*string, bool)`

GetCumulativeQuantityOk returns a tuple with the CumulativeQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeQuantity

`func (o *CryptoOrderOptions) SetCumulativeQuantity(v string)`

SetCumulativeQuantity sets CumulativeQuantity field to given value.

### HasCumulativeQuantity

`func (o *CryptoOrderOptions) HasCumulativeQuantity() bool`

HasCumulativeQuantity returns a boolean if a field has been set.

### GetCurrencyPairId

`func (o *CryptoOrderOptions) GetCurrencyPairId() string`

GetCurrencyPairId returns the CurrencyPairId field if non-nil, zero value otherwise.

### GetCurrencyPairIdOk

`func (o *CryptoOrderOptions) GetCurrencyPairIdOk() (*string, bool)`

GetCurrencyPairIdOk returns a tuple with the CurrencyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyPairId

`func (o *CryptoOrderOptions) SetCurrencyPairId(v string)`

SetCurrencyPairId sets CurrencyPairId field to given value.

### HasCurrencyPairId

`func (o *CryptoOrderOptions) HasCurrencyPairId() bool`

HasCurrencyPairId returns a boolean if a field has been set.

### GetExecutions

`func (o *CryptoOrderOptions) GetExecutions() []Execution`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *CryptoOrderOptions) GetExecutionsOk() (*[]Execution, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *CryptoOrderOptions) SetExecutions(v []Execution)`

SetExecutions sets Executions field to given value.

### HasExecutions

`func (o *CryptoOrderOptions) HasExecutions() bool`

HasExecutions returns a boolean if a field has been set.

### GetId

`func (o *CryptoOrderOptions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CryptoOrderOptions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CryptoOrderOptions) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CryptoOrderOptions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastTransactionAt

`func (o *CryptoOrderOptions) GetLastTransactionAt() string`

GetLastTransactionAt returns the LastTransactionAt field if non-nil, zero value otherwise.

### GetLastTransactionAtOk

`func (o *CryptoOrderOptions) GetLastTransactionAtOk() (*string, bool)`

GetLastTransactionAtOk returns a tuple with the LastTransactionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransactionAt

`func (o *CryptoOrderOptions) SetLastTransactionAt(v string)`

SetLastTransactionAt sets LastTransactionAt field to given value.

### HasLastTransactionAt

`func (o *CryptoOrderOptions) HasLastTransactionAt() bool`

HasLastTransactionAt returns a boolean if a field has been set.

### GetPrice

`func (o *CryptoOrderOptions) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CryptoOrderOptions) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CryptoOrderOptions) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CryptoOrderOptions) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *CryptoOrderOptions) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CryptoOrderOptions) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CryptoOrderOptions) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CryptoOrderOptions) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRejectReason

`func (o *CryptoOrderOptions) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *CryptoOrderOptions) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *CryptoOrderOptions) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *CryptoOrderOptions) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetSide

`func (o *CryptoOrderOptions) GetSide() Side`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CryptoOrderOptions) GetSideOk() (*Side, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CryptoOrderOptions) SetSide(v Side)`

SetSide sets Side field to given value.

### HasSide

`func (o *CryptoOrderOptions) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetState

`func (o *CryptoOrderOptions) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CryptoOrderOptions) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CryptoOrderOptions) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CryptoOrderOptions) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStopPrice

`func (o *CryptoOrderOptions) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *CryptoOrderOptions) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *CryptoOrderOptions) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *CryptoOrderOptions) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetTimeInForce

`func (o *CryptoOrderOptions) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *CryptoOrderOptions) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *CryptoOrderOptions) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *CryptoOrderOptions) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *CryptoOrderOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CryptoOrderOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CryptoOrderOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CryptoOrderOptions) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


