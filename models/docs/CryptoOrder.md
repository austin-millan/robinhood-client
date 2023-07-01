# CryptoOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**CurrencyPairId** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**Side** | Pointer to [**Side**](Side.md) |  | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 
**Type** | Pointer to [**ExecutionType**](ExecutionType.md) |  | [optional] 

## Methods

### NewCryptoOrder

`func NewCryptoOrder() *CryptoOrder`

NewCryptoOrder instantiates a new CryptoOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoOrderWithDefaults

`func NewCryptoOrderWithDefaults() *CryptoOrder`

NewCryptoOrderWithDefaults instantiates a new CryptoOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CryptoOrder) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CryptoOrder) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CryptoOrder) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CryptoOrder) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCurrencyPairId

`func (o *CryptoOrder) GetCurrencyPairId() string`

GetCurrencyPairId returns the CurrencyPairId field if non-nil, zero value otherwise.

### GetCurrencyPairIdOk

`func (o *CryptoOrder) GetCurrencyPairIdOk() (*string, bool)`

GetCurrencyPairIdOk returns a tuple with the CurrencyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyPairId

`func (o *CryptoOrder) SetCurrencyPairId(v string)`

SetCurrencyPairId sets CurrencyPairId field to given value.

### HasCurrencyPairId

`func (o *CryptoOrder) HasCurrencyPairId() bool`

HasCurrencyPairId returns a boolean if a field has been set.

### GetPrice

`func (o *CryptoOrder) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CryptoOrder) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CryptoOrder) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CryptoOrder) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *CryptoOrder) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CryptoOrder) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CryptoOrder) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CryptoOrder) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRefId

`func (o *CryptoOrder) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CryptoOrder) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CryptoOrder) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CryptoOrder) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetSide

`func (o *CryptoOrder) GetSide() Side`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CryptoOrder) GetSideOk() (*Side, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CryptoOrder) SetSide(v Side)`

SetSide sets Side field to given value.

### HasSide

`func (o *CryptoOrder) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetTimeInForce

`func (o *CryptoOrder) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *CryptoOrder) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *CryptoOrder) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *CryptoOrder) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *CryptoOrder) GetType() ExecutionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CryptoOrder) GetTypeOk() (*ExecutionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CryptoOrder) SetType(v ExecutionType)`

SetType sets Type field to given value.

### HasType

`func (o *CryptoOrder) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


