# QuoteData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustedPreviousClose** | Pointer to **string** |  | [optional] 
**AskPrice** | Pointer to **string** |  | [optional] 
**AskSize** | Pointer to **string** |  | [optional] 
**BidPrice** | Pointer to **string** |  | [optional] 
**BidSize** | Pointer to **string** |  | [optional] 
**LastExtendedHoursTradePrice** | Pointer to **string** |  | [optional] 
**LastTradePrice** | Pointer to **string** |  | [optional] 
**PreviousClose** | Pointer to **string** |  | [optional] 
**PreviousCloseDate** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TradingHalted** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewQuoteData

`func NewQuoteData() *QuoteData`

NewQuoteData instantiates a new QuoteData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteDataWithDefaults

`func NewQuoteDataWithDefaults() *QuoteData`

NewQuoteDataWithDefaults instantiates a new QuoteData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustedPreviousClose

`func (o *QuoteData) GetAdjustedPreviousClose() string`

GetAdjustedPreviousClose returns the AdjustedPreviousClose field if non-nil, zero value otherwise.

### GetAdjustedPreviousCloseOk

`func (o *QuoteData) GetAdjustedPreviousCloseOk() (*string, bool)`

GetAdjustedPreviousCloseOk returns a tuple with the AdjustedPreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedPreviousClose

`func (o *QuoteData) SetAdjustedPreviousClose(v string)`

SetAdjustedPreviousClose sets AdjustedPreviousClose field to given value.

### HasAdjustedPreviousClose

`func (o *QuoteData) HasAdjustedPreviousClose() bool`

HasAdjustedPreviousClose returns a boolean if a field has been set.

### GetAskPrice

`func (o *QuoteData) GetAskPrice() string`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *QuoteData) GetAskPriceOk() (*string, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *QuoteData) SetAskPrice(v string)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *QuoteData) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *QuoteData) GetAskSize() string`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *QuoteData) GetAskSizeOk() (*string, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *QuoteData) SetAskSize(v string)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *QuoteData) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBidPrice

`func (o *QuoteData) GetBidPrice() string`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *QuoteData) GetBidPriceOk() (*string, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *QuoteData) SetBidPrice(v string)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *QuoteData) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *QuoteData) GetBidSize() string`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *QuoteData) GetBidSizeOk() (*string, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *QuoteData) SetBidSize(v string)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *QuoteData) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetLastExtendedHoursTradePrice

`func (o *QuoteData) GetLastExtendedHoursTradePrice() string`

GetLastExtendedHoursTradePrice returns the LastExtendedHoursTradePrice field if non-nil, zero value otherwise.

### GetLastExtendedHoursTradePriceOk

`func (o *QuoteData) GetLastExtendedHoursTradePriceOk() (*string, bool)`

GetLastExtendedHoursTradePriceOk returns a tuple with the LastExtendedHoursTradePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExtendedHoursTradePrice

`func (o *QuoteData) SetLastExtendedHoursTradePrice(v string)`

SetLastExtendedHoursTradePrice sets LastExtendedHoursTradePrice field to given value.

### HasLastExtendedHoursTradePrice

`func (o *QuoteData) HasLastExtendedHoursTradePrice() bool`

HasLastExtendedHoursTradePrice returns a boolean if a field has been set.

### GetLastTradePrice

`func (o *QuoteData) GetLastTradePrice() string`

GetLastTradePrice returns the LastTradePrice field if non-nil, zero value otherwise.

### GetLastTradePriceOk

`func (o *QuoteData) GetLastTradePriceOk() (*string, bool)`

GetLastTradePriceOk returns a tuple with the LastTradePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTradePrice

`func (o *QuoteData) SetLastTradePrice(v string)`

SetLastTradePrice sets LastTradePrice field to given value.

### HasLastTradePrice

`func (o *QuoteData) HasLastTradePrice() bool`

HasLastTradePrice returns a boolean if a field has been set.

### GetPreviousClose

`func (o *QuoteData) GetPreviousClose() string`

GetPreviousClose returns the PreviousClose field if non-nil, zero value otherwise.

### GetPreviousCloseOk

`func (o *QuoteData) GetPreviousCloseOk() (*string, bool)`

GetPreviousCloseOk returns a tuple with the PreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousClose

`func (o *QuoteData) SetPreviousClose(v string)`

SetPreviousClose sets PreviousClose field to given value.

### HasPreviousClose

`func (o *QuoteData) HasPreviousClose() bool`

HasPreviousClose returns a boolean if a field has been set.

### GetPreviousCloseDate

`func (o *QuoteData) GetPreviousCloseDate() string`

GetPreviousCloseDate returns the PreviousCloseDate field if non-nil, zero value otherwise.

### GetPreviousCloseDateOk

`func (o *QuoteData) GetPreviousCloseDateOk() (*string, bool)`

GetPreviousCloseDateOk returns a tuple with the PreviousCloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousCloseDate

`func (o *QuoteData) SetPreviousCloseDate(v string)`

SetPreviousCloseDate sets PreviousCloseDate field to given value.

### HasPreviousCloseDate

`func (o *QuoteData) HasPreviousCloseDate() bool`

HasPreviousCloseDate returns a boolean if a field has been set.

### GetSymbol

`func (o *QuoteData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QuoteData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QuoteData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QuoteData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTradingHalted

`func (o *QuoteData) GetTradingHalted() bool`

GetTradingHalted returns the TradingHalted field if non-nil, zero value otherwise.

### GetTradingHaltedOk

`func (o *QuoteData) GetTradingHaltedOk() (*bool, bool)`

GetTradingHaltedOk returns a tuple with the TradingHalted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingHalted

`func (o *QuoteData) SetTradingHalted(v bool)`

SetTradingHalted sets TradingHalted field to given value.

### HasTradingHalted

`func (o *QuoteData) HasTradingHalted() bool`

HasTradingHalted returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *QuoteData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QuoteData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QuoteData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *QuoteData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


