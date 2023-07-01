# CryptoCurrencyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetCurrency** | Pointer to [**CryptoAssetCurrency**](CryptoAssetCurrency.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaxOrderSize** | Pointer to **string** |  | [optional] 
**MinOrderPriceIncrement** | Pointer to **string** |  | [optional] 
**MinOrderSize** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**QuoteCurrency** | Pointer to [**QuoteCurrency**](QuoteCurrency.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Tradability** | Pointer to **string** |  | [optional] 

## Methods

### NewCryptoCurrencyPair

`func NewCryptoCurrencyPair() *CryptoCurrencyPair`

NewCryptoCurrencyPair instantiates a new CryptoCurrencyPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoCurrencyPairWithDefaults

`func NewCryptoCurrencyPairWithDefaults() *CryptoCurrencyPair`

NewCryptoCurrencyPairWithDefaults instantiates a new CryptoCurrencyPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetCurrency

`func (o *CryptoCurrencyPair) GetAssetCurrency() CryptoAssetCurrency`

GetAssetCurrency returns the AssetCurrency field if non-nil, zero value otherwise.

### GetAssetCurrencyOk

`func (o *CryptoCurrencyPair) GetAssetCurrencyOk() (*CryptoAssetCurrency, bool)`

GetAssetCurrencyOk returns a tuple with the AssetCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCurrency

`func (o *CryptoCurrencyPair) SetAssetCurrency(v CryptoAssetCurrency)`

SetAssetCurrency sets AssetCurrency field to given value.

### HasAssetCurrency

`func (o *CryptoCurrencyPair) HasAssetCurrency() bool`

HasAssetCurrency returns a boolean if a field has been set.

### GetId

`func (o *CryptoCurrencyPair) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CryptoCurrencyPair) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CryptoCurrencyPair) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CryptoCurrencyPair) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxOrderSize

`func (o *CryptoCurrencyPair) GetMaxOrderSize() string`

GetMaxOrderSize returns the MaxOrderSize field if non-nil, zero value otherwise.

### GetMaxOrderSizeOk

`func (o *CryptoCurrencyPair) GetMaxOrderSizeOk() (*string, bool)`

GetMaxOrderSizeOk returns a tuple with the MaxOrderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrderSize

`func (o *CryptoCurrencyPair) SetMaxOrderSize(v string)`

SetMaxOrderSize sets MaxOrderSize field to given value.

### HasMaxOrderSize

`func (o *CryptoCurrencyPair) HasMaxOrderSize() bool`

HasMaxOrderSize returns a boolean if a field has been set.

### GetMinOrderPriceIncrement

`func (o *CryptoCurrencyPair) GetMinOrderPriceIncrement() string`

GetMinOrderPriceIncrement returns the MinOrderPriceIncrement field if non-nil, zero value otherwise.

### GetMinOrderPriceIncrementOk

`func (o *CryptoCurrencyPair) GetMinOrderPriceIncrementOk() (*string, bool)`

GetMinOrderPriceIncrementOk returns a tuple with the MinOrderPriceIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrderPriceIncrement

`func (o *CryptoCurrencyPair) SetMinOrderPriceIncrement(v string)`

SetMinOrderPriceIncrement sets MinOrderPriceIncrement field to given value.

### HasMinOrderPriceIncrement

`func (o *CryptoCurrencyPair) HasMinOrderPriceIncrement() bool`

HasMinOrderPriceIncrement returns a boolean if a field has been set.

### GetMinOrderSize

`func (o *CryptoCurrencyPair) GetMinOrderSize() string`

GetMinOrderSize returns the MinOrderSize field if non-nil, zero value otherwise.

### GetMinOrderSizeOk

`func (o *CryptoCurrencyPair) GetMinOrderSizeOk() (*string, bool)`

GetMinOrderSizeOk returns a tuple with the MinOrderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrderSize

`func (o *CryptoCurrencyPair) SetMinOrderSize(v string)`

SetMinOrderSize sets MinOrderSize field to given value.

### HasMinOrderSize

`func (o *CryptoCurrencyPair) HasMinOrderSize() bool`

HasMinOrderSize returns a boolean if a field has been set.

### GetName

`func (o *CryptoCurrencyPair) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CryptoCurrencyPair) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CryptoCurrencyPair) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CryptoCurrencyPair) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuoteCurrency

`func (o *CryptoCurrencyPair) GetQuoteCurrency() QuoteCurrency`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *CryptoCurrencyPair) GetQuoteCurrencyOk() (*QuoteCurrency, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *CryptoCurrencyPair) SetQuoteCurrency(v QuoteCurrency)`

SetQuoteCurrency sets QuoteCurrency field to given value.

### HasQuoteCurrency

`func (o *CryptoCurrencyPair) HasQuoteCurrency() bool`

HasQuoteCurrency returns a boolean if a field has been set.

### GetSymbol

`func (o *CryptoCurrencyPair) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CryptoCurrencyPair) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CryptoCurrencyPair) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CryptoCurrencyPair) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTradability

`func (o *CryptoCurrencyPair) GetTradability() string`

GetTradability returns the Tradability field if non-nil, zero value otherwise.

### GetTradabilityOk

`func (o *CryptoCurrencyPair) GetTradabilityOk() (*string, bool)`

GetTradabilityOk returns a tuple with the Tradability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradability

`func (o *CryptoCurrencyPair) SetTradability(v string)`

SetTradability sets Tradability field to given value.

### HasTradability

`func (o *CryptoCurrencyPair) HasTradability() bool`

HasTradability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


