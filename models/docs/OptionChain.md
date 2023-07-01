# OptionChain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanOpenPosition** | Pointer to **bool** |  | [optional] 
**CashComponent** | Pointer to **map[string]interface{}** |  | [optional] 
**ExpirationDates** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MinTicks** | Pointer to [**MinTicks**](MinTicks.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TradeValueMultiplier** | Pointer to **string** |  | [optional] 
**UnderlyingInstruments** | Pointer to [**[]UnderlyingInstrument**](UnderlyingInstrument.md) |  | [optional] 

## Methods

### NewOptionChain

`func NewOptionChain() *OptionChain`

NewOptionChain instantiates a new OptionChain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionChainWithDefaults

`func NewOptionChainWithDefaults() *OptionChain`

NewOptionChainWithDefaults instantiates a new OptionChain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanOpenPosition

`func (o *OptionChain) GetCanOpenPosition() bool`

GetCanOpenPosition returns the CanOpenPosition field if non-nil, zero value otherwise.

### GetCanOpenPositionOk

`func (o *OptionChain) GetCanOpenPositionOk() (*bool, bool)`

GetCanOpenPositionOk returns a tuple with the CanOpenPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanOpenPosition

`func (o *OptionChain) SetCanOpenPosition(v bool)`

SetCanOpenPosition sets CanOpenPosition field to given value.

### HasCanOpenPosition

`func (o *OptionChain) HasCanOpenPosition() bool`

HasCanOpenPosition returns a boolean if a field has been set.

### GetCashComponent

`func (o *OptionChain) GetCashComponent() map[string]interface{}`

GetCashComponent returns the CashComponent field if non-nil, zero value otherwise.

### GetCashComponentOk

`func (o *OptionChain) GetCashComponentOk() (*map[string]interface{}, bool)`

GetCashComponentOk returns a tuple with the CashComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashComponent

`func (o *OptionChain) SetCashComponent(v map[string]interface{})`

SetCashComponent sets CashComponent field to given value.

### HasCashComponent

`func (o *OptionChain) HasCashComponent() bool`

HasCashComponent returns a boolean if a field has been set.

### GetExpirationDates

`func (o *OptionChain) GetExpirationDates() []string`

GetExpirationDates returns the ExpirationDates field if non-nil, zero value otherwise.

### GetExpirationDatesOk

`func (o *OptionChain) GetExpirationDatesOk() (*[]string, bool)`

GetExpirationDatesOk returns a tuple with the ExpirationDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDates

`func (o *OptionChain) SetExpirationDates(v []string)`

SetExpirationDates sets ExpirationDates field to given value.

### HasExpirationDates

`func (o *OptionChain) HasExpirationDates() bool`

HasExpirationDates returns a boolean if a field has been set.

### GetId

`func (o *OptionChain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptionChain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptionChain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OptionChain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMinTicks

`func (o *OptionChain) GetMinTicks() MinTicks`

GetMinTicks returns the MinTicks field if non-nil, zero value otherwise.

### GetMinTicksOk

`func (o *OptionChain) GetMinTicksOk() (*MinTicks, bool)`

GetMinTicksOk returns a tuple with the MinTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTicks

`func (o *OptionChain) SetMinTicks(v MinTicks)`

SetMinTicks sets MinTicks field to given value.

### HasMinTicks

`func (o *OptionChain) HasMinTicks() bool`

HasMinTicks returns a boolean if a field has been set.

### GetSymbol

`func (o *OptionChain) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OptionChain) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OptionChain) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OptionChain) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTradeValueMultiplier

`func (o *OptionChain) GetTradeValueMultiplier() string`

GetTradeValueMultiplier returns the TradeValueMultiplier field if non-nil, zero value otherwise.

### GetTradeValueMultiplierOk

`func (o *OptionChain) GetTradeValueMultiplierOk() (*string, bool)`

GetTradeValueMultiplierOk returns a tuple with the TradeValueMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeValueMultiplier

`func (o *OptionChain) SetTradeValueMultiplier(v string)`

SetTradeValueMultiplier sets TradeValueMultiplier field to given value.

### HasTradeValueMultiplier

`func (o *OptionChain) HasTradeValueMultiplier() bool`

HasTradeValueMultiplier returns a boolean if a field has been set.

### GetUnderlyingInstruments

`func (o *OptionChain) GetUnderlyingInstruments() []UnderlyingInstrument`

GetUnderlyingInstruments returns the UnderlyingInstruments field if non-nil, zero value otherwise.

### GetUnderlyingInstrumentsOk

`func (o *OptionChain) GetUnderlyingInstrumentsOk() (*[]UnderlyingInstrument, bool)`

GetUnderlyingInstrumentsOk returns a tuple with the UnderlyingInstruments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingInstruments

`func (o *OptionChain) SetUnderlyingInstruments(v []UnderlyingInstrument)`

SetUnderlyingInstruments sets UnderlyingInstruments field to given value.

### HasUnderlyingInstruments

`func (o *OptionChain) HasUnderlyingInstruments() bool`

HasUnderlyingInstruments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


