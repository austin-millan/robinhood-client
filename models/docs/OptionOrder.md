# OptionOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelUrl** | Pointer to **string** |  | [optional] 
**CanceledQuantity** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to [**Direction**](Direction.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Legs** | Pointer to [**[]OptionOrderLeg**](OptionOrderLeg.md) |  | [optional] 
**PendingQuantity** | Pointer to **string** |  | [optional] 
**Premium** | Pointer to **string** |  | [optional] 
**ProcessedPremium** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**ProcessedQuantity** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**OrderState**](OrderState.md) |  | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 
**Trigger** | Pointer to [**Trigger**](Trigger.md) |  | [optional] 
**Type** | Pointer to [**ExecutionType**](ExecutionType.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**ChainId** | Pointer to **string** |  | [optional] 
**ChainSymbol** | Pointer to **string** |  | [optional] 
**ResponseCategory** | Pointer to **string** |  | [optional] 
**OpeningStrategy** | Pointer to [**OpenCloseStrategy**](OpenCloseStrategy.md) |  | [optional] 
**ClosingStrategy** | Pointer to [**OpenCloseStrategy**](OpenCloseStrategy.md) |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 

## Methods

### NewOptionOrder

`func NewOptionOrder() *OptionOrder`

NewOptionOrder instantiates a new OptionOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionOrderWithDefaults

`func NewOptionOrderWithDefaults() *OptionOrder`

NewOptionOrderWithDefaults instantiates a new OptionOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelUrl

`func (o *OptionOrder) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *OptionOrder) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *OptionOrder) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *OptionOrder) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetCanceledQuantity

`func (o *OptionOrder) GetCanceledQuantity() string`

GetCanceledQuantity returns the CanceledQuantity field if non-nil, zero value otherwise.

### GetCanceledQuantityOk

`func (o *OptionOrder) GetCanceledQuantityOk() (*string, bool)`

GetCanceledQuantityOk returns a tuple with the CanceledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledQuantity

`func (o *OptionOrder) SetCanceledQuantity(v string)`

SetCanceledQuantity sets CanceledQuantity field to given value.

### HasCanceledQuantity

`func (o *OptionOrder) HasCanceledQuantity() bool`

HasCanceledQuantity returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OptionOrder) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OptionOrder) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OptionOrder) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OptionOrder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDirection

`func (o *OptionOrder) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *OptionOrder) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *OptionOrder) SetDirection(v Direction)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *OptionOrder) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *OptionOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptionOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptionOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OptionOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLegs

`func (o *OptionOrder) GetLegs() []OptionOrderLeg`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *OptionOrder) GetLegsOk() (*[]OptionOrderLeg, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *OptionOrder) SetLegs(v []OptionOrderLeg)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *OptionOrder) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetPendingQuantity

`func (o *OptionOrder) GetPendingQuantity() string`

GetPendingQuantity returns the PendingQuantity field if non-nil, zero value otherwise.

### GetPendingQuantityOk

`func (o *OptionOrder) GetPendingQuantityOk() (*string, bool)`

GetPendingQuantityOk returns a tuple with the PendingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingQuantity

`func (o *OptionOrder) SetPendingQuantity(v string)`

SetPendingQuantity sets PendingQuantity field to given value.

### HasPendingQuantity

`func (o *OptionOrder) HasPendingQuantity() bool`

HasPendingQuantity returns a boolean if a field has been set.

### GetPremium

`func (o *OptionOrder) GetPremium() string`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *OptionOrder) GetPremiumOk() (*string, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *OptionOrder) SetPremium(v string)`

SetPremium sets Premium field to given value.

### HasPremium

`func (o *OptionOrder) HasPremium() bool`

HasPremium returns a boolean if a field has been set.

### GetProcessedPremium

`func (o *OptionOrder) GetProcessedPremium() string`

GetProcessedPremium returns the ProcessedPremium field if non-nil, zero value otherwise.

### GetProcessedPremiumOk

`func (o *OptionOrder) GetProcessedPremiumOk() (*string, bool)`

GetProcessedPremiumOk returns a tuple with the ProcessedPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedPremium

`func (o *OptionOrder) SetProcessedPremium(v string)`

SetProcessedPremium sets ProcessedPremium field to given value.

### HasProcessedPremium

`func (o *OptionOrder) HasProcessedPremium() bool`

HasProcessedPremium returns a boolean if a field has been set.

### GetPrice

`func (o *OptionOrder) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OptionOrder) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OptionOrder) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OptionOrder) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetProcessedQuantity

`func (o *OptionOrder) GetProcessedQuantity() string`

GetProcessedQuantity returns the ProcessedQuantity field if non-nil, zero value otherwise.

### GetProcessedQuantityOk

`func (o *OptionOrder) GetProcessedQuantityOk() (*string, bool)`

GetProcessedQuantityOk returns a tuple with the ProcessedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedQuantity

`func (o *OptionOrder) SetProcessedQuantity(v string)`

SetProcessedQuantity sets ProcessedQuantity field to given value.

### HasProcessedQuantity

`func (o *OptionOrder) HasProcessedQuantity() bool`

HasProcessedQuantity returns a boolean if a field has been set.

### GetQuantity

`func (o *OptionOrder) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OptionOrder) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OptionOrder) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OptionOrder) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRefId

`func (o *OptionOrder) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *OptionOrder) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *OptionOrder) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *OptionOrder) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetState

`func (o *OptionOrder) GetState() OrderState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OptionOrder) GetStateOk() (*OrderState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OptionOrder) SetState(v OrderState)`

SetState sets State field to given value.

### HasState

`func (o *OptionOrder) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimeInForce

`func (o *OptionOrder) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OptionOrder) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OptionOrder) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OptionOrder) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetTrigger

`func (o *OptionOrder) GetTrigger() Trigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *OptionOrder) GetTriggerOk() (*Trigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *OptionOrder) SetTrigger(v Trigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *OptionOrder) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetType

`func (o *OptionOrder) GetType() ExecutionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OptionOrder) GetTypeOk() (*ExecutionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OptionOrder) SetType(v ExecutionType)`

SetType sets Type field to given value.

### HasType

`func (o *OptionOrder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OptionOrder) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OptionOrder) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OptionOrder) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OptionOrder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetChainId

`func (o *OptionOrder) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *OptionOrder) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *OptionOrder) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *OptionOrder) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetChainSymbol

`func (o *OptionOrder) GetChainSymbol() string`

GetChainSymbol returns the ChainSymbol field if non-nil, zero value otherwise.

### GetChainSymbolOk

`func (o *OptionOrder) GetChainSymbolOk() (*string, bool)`

GetChainSymbolOk returns a tuple with the ChainSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainSymbol

`func (o *OptionOrder) SetChainSymbol(v string)`

SetChainSymbol sets ChainSymbol field to given value.

### HasChainSymbol

`func (o *OptionOrder) HasChainSymbol() bool`

HasChainSymbol returns a boolean if a field has been set.

### GetResponseCategory

`func (o *OptionOrder) GetResponseCategory() string`

GetResponseCategory returns the ResponseCategory field if non-nil, zero value otherwise.

### GetResponseCategoryOk

`func (o *OptionOrder) GetResponseCategoryOk() (*string, bool)`

GetResponseCategoryOk returns a tuple with the ResponseCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCategory

`func (o *OptionOrder) SetResponseCategory(v string)`

SetResponseCategory sets ResponseCategory field to given value.

### HasResponseCategory

`func (o *OptionOrder) HasResponseCategory() bool`

HasResponseCategory returns a boolean if a field has been set.

### GetOpeningStrategy

`func (o *OptionOrder) GetOpeningStrategy() OpenCloseStrategy`

GetOpeningStrategy returns the OpeningStrategy field if non-nil, zero value otherwise.

### GetOpeningStrategyOk

`func (o *OptionOrder) GetOpeningStrategyOk() (*OpenCloseStrategy, bool)`

GetOpeningStrategyOk returns a tuple with the OpeningStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningStrategy

`func (o *OptionOrder) SetOpeningStrategy(v OpenCloseStrategy)`

SetOpeningStrategy sets OpeningStrategy field to given value.

### HasOpeningStrategy

`func (o *OptionOrder) HasOpeningStrategy() bool`

HasOpeningStrategy returns a boolean if a field has been set.

### GetClosingStrategy

`func (o *OptionOrder) GetClosingStrategy() OpenCloseStrategy`

GetClosingStrategy returns the ClosingStrategy field if non-nil, zero value otherwise.

### GetClosingStrategyOk

`func (o *OptionOrder) GetClosingStrategyOk() (*OpenCloseStrategy, bool)`

GetClosingStrategyOk returns a tuple with the ClosingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingStrategy

`func (o *OptionOrder) SetClosingStrategy(v OpenCloseStrategy)`

SetClosingStrategy sets ClosingStrategy field to given value.

### HasClosingStrategy

`func (o *OptionOrder) HasClosingStrategy() bool`

HasClosingStrategy returns a boolean if a field has been set.

### GetStopPrice

`func (o *OptionOrder) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *OptionOrder) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *OptionOrder) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *OptionOrder) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


