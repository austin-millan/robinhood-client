# Position

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**AverageBuyPrice** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Instrument** | Pointer to **string** |  | [optional] 
**IntradayAverageBuyPrice** | Pointer to **string** |  | [optional] 
**IntradayQuantity** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**SharesHeldForBuys** | Pointer to **string** |  | [optional] 
**SharesHeldForSells** | Pointer to **string** |  | [optional] 
**SharesHeldForStockGrants** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewPosition

`func NewPosition() *Position`

NewPosition instantiates a new Position object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionWithDefaults

`func NewPositionWithDefaults() *Position`

NewPositionWithDefaults instantiates a new Position object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Position) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Position) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Position) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Position) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAverageBuyPrice

`func (o *Position) GetAverageBuyPrice() string`

GetAverageBuyPrice returns the AverageBuyPrice field if non-nil, zero value otherwise.

### GetAverageBuyPriceOk

`func (o *Position) GetAverageBuyPriceOk() (*string, bool)`

GetAverageBuyPriceOk returns a tuple with the AverageBuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageBuyPrice

`func (o *Position) SetAverageBuyPrice(v string)`

SetAverageBuyPrice sets AverageBuyPrice field to given value.

### HasAverageBuyPrice

`func (o *Position) HasAverageBuyPrice() bool`

HasAverageBuyPrice returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Position) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Position) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Position) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Position) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetInstrument

`func (o *Position) GetInstrument() string`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *Position) GetInstrumentOk() (*string, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *Position) SetInstrument(v string)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *Position) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### GetIntradayAverageBuyPrice

`func (o *Position) GetIntradayAverageBuyPrice() string`

GetIntradayAverageBuyPrice returns the IntradayAverageBuyPrice field if non-nil, zero value otherwise.

### GetIntradayAverageBuyPriceOk

`func (o *Position) GetIntradayAverageBuyPriceOk() (*string, bool)`

GetIntradayAverageBuyPriceOk returns a tuple with the IntradayAverageBuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntradayAverageBuyPrice

`func (o *Position) SetIntradayAverageBuyPrice(v string)`

SetIntradayAverageBuyPrice sets IntradayAverageBuyPrice field to given value.

### HasIntradayAverageBuyPrice

`func (o *Position) HasIntradayAverageBuyPrice() bool`

HasIntradayAverageBuyPrice returns a boolean if a field has been set.

### GetIntradayQuantity

`func (o *Position) GetIntradayQuantity() string`

GetIntradayQuantity returns the IntradayQuantity field if non-nil, zero value otherwise.

### GetIntradayQuantityOk

`func (o *Position) GetIntradayQuantityOk() (*string, bool)`

GetIntradayQuantityOk returns a tuple with the IntradayQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntradayQuantity

`func (o *Position) SetIntradayQuantity(v string)`

SetIntradayQuantity sets IntradayQuantity field to given value.

### HasIntradayQuantity

`func (o *Position) HasIntradayQuantity() bool`

HasIntradayQuantity returns a boolean if a field has been set.

### GetQuantity

`func (o *Position) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Position) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Position) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Position) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSharesHeldForBuys

`func (o *Position) GetSharesHeldForBuys() string`

GetSharesHeldForBuys returns the SharesHeldForBuys field if non-nil, zero value otherwise.

### GetSharesHeldForBuysOk

`func (o *Position) GetSharesHeldForBuysOk() (*string, bool)`

GetSharesHeldForBuysOk returns a tuple with the SharesHeldForBuys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesHeldForBuys

`func (o *Position) SetSharesHeldForBuys(v string)`

SetSharesHeldForBuys sets SharesHeldForBuys field to given value.

### HasSharesHeldForBuys

`func (o *Position) HasSharesHeldForBuys() bool`

HasSharesHeldForBuys returns a boolean if a field has been set.

### GetSharesHeldForSells

`func (o *Position) GetSharesHeldForSells() string`

GetSharesHeldForSells returns the SharesHeldForSells field if non-nil, zero value otherwise.

### GetSharesHeldForSellsOk

`func (o *Position) GetSharesHeldForSellsOk() (*string, bool)`

GetSharesHeldForSellsOk returns a tuple with the SharesHeldForSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesHeldForSells

`func (o *Position) SetSharesHeldForSells(v string)`

SetSharesHeldForSells sets SharesHeldForSells field to given value.

### HasSharesHeldForSells

`func (o *Position) HasSharesHeldForSells() bool`

HasSharesHeldForSells returns a boolean if a field has been set.

### GetSharesHeldForStockGrants

`func (o *Position) GetSharesHeldForStockGrants() string`

GetSharesHeldForStockGrants returns the SharesHeldForStockGrants field if non-nil, zero value otherwise.

### GetSharesHeldForStockGrantsOk

`func (o *Position) GetSharesHeldForStockGrantsOk() (*string, bool)`

GetSharesHeldForStockGrantsOk returns a tuple with the SharesHeldForStockGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesHeldForStockGrants

`func (o *Position) SetSharesHeldForStockGrants(v string)`

SetSharesHeldForStockGrants sets SharesHeldForStockGrants field to given value.

### HasSharesHeldForStockGrants

`func (o *Position) HasSharesHeldForStockGrants() bool`

HasSharesHeldForStockGrants returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Position) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Position) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Position) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Position) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *Position) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Position) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Position) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Position) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


