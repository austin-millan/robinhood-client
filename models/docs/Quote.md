# Quote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustedPreviousClose** | Pointer to **string** |  | [optional] 
**AskPrice** | Pointer to **string** |  | [optional] 
**AskSize** | Pointer to **int32** |  | [optional] 
**BidPrice** | Pointer to **string** |  | [optional] 
**BidSize** | Pointer to **int32** |  | [optional] 
**LastExtendedHoursTradePrice** | Pointer to **string** |  | [optional] 
**LastTradePrice** | Pointer to **string** |  | [optional] 
**PreviousClose** | Pointer to **string** |  | [optional] 
**PreviousCloseDate** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TradingHalted** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewQuote

`func NewQuote() *Quote`

NewQuote instantiates a new Quote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteWithDefaults

`func NewQuoteWithDefaults() *Quote`

NewQuoteWithDefaults instantiates a new Quote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustedPreviousClose

`func (o *Quote) GetAdjustedPreviousClose() string`

GetAdjustedPreviousClose returns the AdjustedPreviousClose field if non-nil, zero value otherwise.

### GetAdjustedPreviousCloseOk

`func (o *Quote) GetAdjustedPreviousCloseOk() (*string, bool)`

GetAdjustedPreviousCloseOk returns a tuple with the AdjustedPreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedPreviousClose

`func (o *Quote) SetAdjustedPreviousClose(v string)`

SetAdjustedPreviousClose sets AdjustedPreviousClose field to given value.

### HasAdjustedPreviousClose

`func (o *Quote) HasAdjustedPreviousClose() bool`

HasAdjustedPreviousClose returns a boolean if a field has been set.

### GetAskPrice

`func (o *Quote) GetAskPrice() string`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *Quote) GetAskPriceOk() (*string, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *Quote) SetAskPrice(v string)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *Quote) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *Quote) GetAskSize() int32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *Quote) GetAskSizeOk() (*int32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *Quote) SetAskSize(v int32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *Quote) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBidPrice

`func (o *Quote) GetBidPrice() string`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *Quote) GetBidPriceOk() (*string, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *Quote) SetBidPrice(v string)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *Quote) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *Quote) GetBidSize() int32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *Quote) GetBidSizeOk() (*int32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *Quote) SetBidSize(v int32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *Quote) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetLastExtendedHoursTradePrice

`func (o *Quote) GetLastExtendedHoursTradePrice() string`

GetLastExtendedHoursTradePrice returns the LastExtendedHoursTradePrice field if non-nil, zero value otherwise.

### GetLastExtendedHoursTradePriceOk

`func (o *Quote) GetLastExtendedHoursTradePriceOk() (*string, bool)`

GetLastExtendedHoursTradePriceOk returns a tuple with the LastExtendedHoursTradePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExtendedHoursTradePrice

`func (o *Quote) SetLastExtendedHoursTradePrice(v string)`

SetLastExtendedHoursTradePrice sets LastExtendedHoursTradePrice field to given value.

### HasLastExtendedHoursTradePrice

`func (o *Quote) HasLastExtendedHoursTradePrice() bool`

HasLastExtendedHoursTradePrice returns a boolean if a field has been set.

### GetLastTradePrice

`func (o *Quote) GetLastTradePrice() string`

GetLastTradePrice returns the LastTradePrice field if non-nil, zero value otherwise.

### GetLastTradePriceOk

`func (o *Quote) GetLastTradePriceOk() (*string, bool)`

GetLastTradePriceOk returns a tuple with the LastTradePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTradePrice

`func (o *Quote) SetLastTradePrice(v string)`

SetLastTradePrice sets LastTradePrice field to given value.

### HasLastTradePrice

`func (o *Quote) HasLastTradePrice() bool`

HasLastTradePrice returns a boolean if a field has been set.

### GetPreviousClose

`func (o *Quote) GetPreviousClose() string`

GetPreviousClose returns the PreviousClose field if non-nil, zero value otherwise.

### GetPreviousCloseOk

`func (o *Quote) GetPreviousCloseOk() (*string, bool)`

GetPreviousCloseOk returns a tuple with the PreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousClose

`func (o *Quote) SetPreviousClose(v string)`

SetPreviousClose sets PreviousClose field to given value.

### HasPreviousClose

`func (o *Quote) HasPreviousClose() bool`

HasPreviousClose returns a boolean if a field has been set.

### GetPreviousCloseDate

`func (o *Quote) GetPreviousCloseDate() string`

GetPreviousCloseDate returns the PreviousCloseDate field if non-nil, zero value otherwise.

### GetPreviousCloseDateOk

`func (o *Quote) GetPreviousCloseDateOk() (*string, bool)`

GetPreviousCloseDateOk returns a tuple with the PreviousCloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousCloseDate

`func (o *Quote) SetPreviousCloseDate(v string)`

SetPreviousCloseDate sets PreviousCloseDate field to given value.

### HasPreviousCloseDate

`func (o *Quote) HasPreviousCloseDate() bool`

HasPreviousCloseDate returns a boolean if a field has been set.

### GetSymbol

`func (o *Quote) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Quote) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Quote) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Quote) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTradingHalted

`func (o *Quote) GetTradingHalted() bool`

GetTradingHalted returns the TradingHalted field if non-nil, zero value otherwise.

### GetTradingHaltedOk

`func (o *Quote) GetTradingHaltedOk() (*bool, bool)`

GetTradingHaltedOk returns a tuple with the TradingHalted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingHalted

`func (o *Quote) SetTradingHalted(v bool)`

SetTradingHalted sets TradingHalted field to given value.

### HasTradingHalted

`func (o *Quote) HasTradingHalted() bool`

HasTradingHalted returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Quote) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Quote) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Quote) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Quote) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


