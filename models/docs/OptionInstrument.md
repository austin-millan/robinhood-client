# OptionInstrument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **string** |  | [optional] 
**ChainSymbol** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IssueDate** | Pointer to **string** |  | [optional] 
**MinTicks** | Pointer to [**MinTicks**](MinTicks.md) |  | [optional] 
**RhsTradability** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StrikePrice** | Pointer to **string** |  | [optional] 
**Tradability** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewOptionInstrument

`func NewOptionInstrument() *OptionInstrument`

NewOptionInstrument instantiates a new OptionInstrument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionInstrumentWithDefaults

`func NewOptionInstrumentWithDefaults() *OptionInstrument`

NewOptionInstrumentWithDefaults instantiates a new OptionInstrument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *OptionInstrument) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *OptionInstrument) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *OptionInstrument) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *OptionInstrument) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetChainSymbol

`func (o *OptionInstrument) GetChainSymbol() string`

GetChainSymbol returns the ChainSymbol field if non-nil, zero value otherwise.

### GetChainSymbolOk

`func (o *OptionInstrument) GetChainSymbolOk() (*string, bool)`

GetChainSymbolOk returns a tuple with the ChainSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainSymbol

`func (o *OptionInstrument) SetChainSymbol(v string)`

SetChainSymbol sets ChainSymbol field to given value.

### HasChainSymbol

`func (o *OptionInstrument) HasChainSymbol() bool`

HasChainSymbol returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OptionInstrument) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OptionInstrument) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OptionInstrument) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OptionInstrument) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpirationDate

`func (o *OptionInstrument) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *OptionInstrument) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *OptionInstrument) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *OptionInstrument) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetId

`func (o *OptionInstrument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptionInstrument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptionInstrument) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OptionInstrument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *OptionInstrument) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *OptionInstrument) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *OptionInstrument) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *OptionInstrument) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetMinTicks

`func (o *OptionInstrument) GetMinTicks() MinTicks`

GetMinTicks returns the MinTicks field if non-nil, zero value otherwise.

### GetMinTicksOk

`func (o *OptionInstrument) GetMinTicksOk() (*MinTicks, bool)`

GetMinTicksOk returns a tuple with the MinTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTicks

`func (o *OptionInstrument) SetMinTicks(v MinTicks)`

SetMinTicks sets MinTicks field to given value.

### HasMinTicks

`func (o *OptionInstrument) HasMinTicks() bool`

HasMinTicks returns a boolean if a field has been set.

### GetRhsTradability

`func (o *OptionInstrument) GetRhsTradability() string`

GetRhsTradability returns the RhsTradability field if non-nil, zero value otherwise.

### GetRhsTradabilityOk

`func (o *OptionInstrument) GetRhsTradabilityOk() (*string, bool)`

GetRhsTradabilityOk returns a tuple with the RhsTradability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhsTradability

`func (o *OptionInstrument) SetRhsTradability(v string)`

SetRhsTradability sets RhsTradability field to given value.

### HasRhsTradability

`func (o *OptionInstrument) HasRhsTradability() bool`

HasRhsTradability returns a boolean if a field has been set.

### GetState

`func (o *OptionInstrument) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OptionInstrument) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OptionInstrument) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OptionInstrument) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStrikePrice

`func (o *OptionInstrument) GetStrikePrice() string`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *OptionInstrument) GetStrikePriceOk() (*string, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *OptionInstrument) SetStrikePrice(v string)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *OptionInstrument) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.

### GetTradability

`func (o *OptionInstrument) GetTradability() string`

GetTradability returns the Tradability field if non-nil, zero value otherwise.

### GetTradabilityOk

`func (o *OptionInstrument) GetTradabilityOk() (*string, bool)`

GetTradabilityOk returns a tuple with the Tradability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradability

`func (o *OptionInstrument) SetTradability(v string)`

SetTradability sets Tradability field to given value.

### HasTradability

`func (o *OptionInstrument) HasTradability() bool`

HasTradability returns a boolean if a field has been set.

### GetType

`func (o *OptionInstrument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OptionInstrument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OptionInstrument) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OptionInstrument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OptionInstrument) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OptionInstrument) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OptionInstrument) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OptionInstrument) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *OptionInstrument) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OptionInstrument) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OptionInstrument) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OptionInstrument) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


