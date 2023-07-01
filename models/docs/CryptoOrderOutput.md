# CryptoOrderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**AveragePrice** | Pointer to **string** |  | [optional] 
**Cancel** | Pointer to **string** |  | [optional] 
**CancelUrl** | Pointer to **string** |  | [optional] 
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

### NewCryptoOrderOutput

`func NewCryptoOrderOutput() *CryptoOrderOutput`

NewCryptoOrderOutput instantiates a new CryptoOrderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoOrderOutputWithDefaults

`func NewCryptoOrderOutputWithDefaults() *CryptoOrderOutput`

NewCryptoOrderOutputWithDefaults instantiates a new CryptoOrderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CryptoOrderOutput) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CryptoOrderOutput) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CryptoOrderOutput) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CryptoOrderOutput) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAveragePrice

`func (o *CryptoOrderOutput) GetAveragePrice() string`

GetAveragePrice returns the AveragePrice field if non-nil, zero value otherwise.

### GetAveragePriceOk

`func (o *CryptoOrderOutput) GetAveragePriceOk() (*string, bool)`

GetAveragePriceOk returns a tuple with the AveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePrice

`func (o *CryptoOrderOutput) SetAveragePrice(v string)`

SetAveragePrice sets AveragePrice field to given value.

### HasAveragePrice

`func (o *CryptoOrderOutput) HasAveragePrice() bool`

HasAveragePrice returns a boolean if a field has been set.

### GetCancel

`func (o *CryptoOrderOutput) GetCancel() string`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *CryptoOrderOutput) GetCancelOk() (*string, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *CryptoOrderOutput) SetCancel(v string)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *CryptoOrderOutput) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### GetCancelUrl

`func (o *CryptoOrderOutput) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *CryptoOrderOutput) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *CryptoOrderOutput) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *CryptoOrderOutput) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CryptoOrderOutput) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CryptoOrderOutput) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CryptoOrderOutput) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CryptoOrderOutput) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCumulativeQuantity

`func (o *CryptoOrderOutput) GetCumulativeQuantity() string`

GetCumulativeQuantity returns the CumulativeQuantity field if non-nil, zero value otherwise.

### GetCumulativeQuantityOk

`func (o *CryptoOrderOutput) GetCumulativeQuantityOk() (*string, bool)`

GetCumulativeQuantityOk returns a tuple with the CumulativeQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeQuantity

`func (o *CryptoOrderOutput) SetCumulativeQuantity(v string)`

SetCumulativeQuantity sets CumulativeQuantity field to given value.

### HasCumulativeQuantity

`func (o *CryptoOrderOutput) HasCumulativeQuantity() bool`

HasCumulativeQuantity returns a boolean if a field has been set.

### GetCurrencyPairId

`func (o *CryptoOrderOutput) GetCurrencyPairId() string`

GetCurrencyPairId returns the CurrencyPairId field if non-nil, zero value otherwise.

### GetCurrencyPairIdOk

`func (o *CryptoOrderOutput) GetCurrencyPairIdOk() (*string, bool)`

GetCurrencyPairIdOk returns a tuple with the CurrencyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyPairId

`func (o *CryptoOrderOutput) SetCurrencyPairId(v string)`

SetCurrencyPairId sets CurrencyPairId field to given value.

### HasCurrencyPairId

`func (o *CryptoOrderOutput) HasCurrencyPairId() bool`

HasCurrencyPairId returns a boolean if a field has been set.

### GetExecutions

`func (o *CryptoOrderOutput) GetExecutions() []Execution`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *CryptoOrderOutput) GetExecutionsOk() (*[]Execution, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *CryptoOrderOutput) SetExecutions(v []Execution)`

SetExecutions sets Executions field to given value.

### HasExecutions

`func (o *CryptoOrderOutput) HasExecutions() bool`

HasExecutions returns a boolean if a field has been set.

### GetId

`func (o *CryptoOrderOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CryptoOrderOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CryptoOrderOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CryptoOrderOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastTransactionAt

`func (o *CryptoOrderOutput) GetLastTransactionAt() string`

GetLastTransactionAt returns the LastTransactionAt field if non-nil, zero value otherwise.

### GetLastTransactionAtOk

`func (o *CryptoOrderOutput) GetLastTransactionAtOk() (*string, bool)`

GetLastTransactionAtOk returns a tuple with the LastTransactionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransactionAt

`func (o *CryptoOrderOutput) SetLastTransactionAt(v string)`

SetLastTransactionAt sets LastTransactionAt field to given value.

### HasLastTransactionAt

`func (o *CryptoOrderOutput) HasLastTransactionAt() bool`

HasLastTransactionAt returns a boolean if a field has been set.

### GetPrice

`func (o *CryptoOrderOutput) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CryptoOrderOutput) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CryptoOrderOutput) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CryptoOrderOutput) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *CryptoOrderOutput) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CryptoOrderOutput) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CryptoOrderOutput) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CryptoOrderOutput) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRejectReason

`func (o *CryptoOrderOutput) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *CryptoOrderOutput) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *CryptoOrderOutput) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *CryptoOrderOutput) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetSide

`func (o *CryptoOrderOutput) GetSide() Side`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CryptoOrderOutput) GetSideOk() (*Side, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CryptoOrderOutput) SetSide(v Side)`

SetSide sets Side field to given value.

### HasSide

`func (o *CryptoOrderOutput) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetState

`func (o *CryptoOrderOutput) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CryptoOrderOutput) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CryptoOrderOutput) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CryptoOrderOutput) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStopPrice

`func (o *CryptoOrderOutput) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *CryptoOrderOutput) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *CryptoOrderOutput) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *CryptoOrderOutput) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetTimeInForce

`func (o *CryptoOrderOutput) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *CryptoOrderOutput) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *CryptoOrderOutput) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *CryptoOrderOutput) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *CryptoOrderOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CryptoOrderOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CryptoOrderOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CryptoOrderOutput) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


