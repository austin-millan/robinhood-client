# OptionOrderInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to [**Direction**](Direction.md) |  | [optional] 
**Legs** | Pointer to [**[]Leg**](Leg.md) |  | [optional] 
**OverrideDayTradeChecks** | Pointer to **bool** |  | [optional] 
**OverrideDtbpChecks** | Pointer to **bool** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 
**Trigger** | Pointer to [**Trigger**](Trigger.md) |  | [optional] 
**Type** | Pointer to [**ExecutionType**](ExecutionType.md) |  | [optional] 

## Methods

### NewOptionOrderInput

`func NewOptionOrderInput() *OptionOrderInput`

NewOptionOrderInput instantiates a new OptionOrderInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionOrderInputWithDefaults

`func NewOptionOrderInputWithDefaults() *OptionOrderInput`

NewOptionOrderInputWithDefaults instantiates a new OptionOrderInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *OptionOrderInput) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OptionOrderInput) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OptionOrderInput) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OptionOrderInput) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDirection

`func (o *OptionOrderInput) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *OptionOrderInput) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *OptionOrderInput) SetDirection(v Direction)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *OptionOrderInput) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetLegs

`func (o *OptionOrderInput) GetLegs() []Leg`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *OptionOrderInput) GetLegsOk() (*[]Leg, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *OptionOrderInput) SetLegs(v []Leg)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *OptionOrderInput) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetOverrideDayTradeChecks

`func (o *OptionOrderInput) GetOverrideDayTradeChecks() bool`

GetOverrideDayTradeChecks returns the OverrideDayTradeChecks field if non-nil, zero value otherwise.

### GetOverrideDayTradeChecksOk

`func (o *OptionOrderInput) GetOverrideDayTradeChecksOk() (*bool, bool)`

GetOverrideDayTradeChecksOk returns a tuple with the OverrideDayTradeChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDayTradeChecks

`func (o *OptionOrderInput) SetOverrideDayTradeChecks(v bool)`

SetOverrideDayTradeChecks sets OverrideDayTradeChecks field to given value.

### HasOverrideDayTradeChecks

`func (o *OptionOrderInput) HasOverrideDayTradeChecks() bool`

HasOverrideDayTradeChecks returns a boolean if a field has been set.

### GetOverrideDtbpChecks

`func (o *OptionOrderInput) GetOverrideDtbpChecks() bool`

GetOverrideDtbpChecks returns the OverrideDtbpChecks field if non-nil, zero value otherwise.

### GetOverrideDtbpChecksOk

`func (o *OptionOrderInput) GetOverrideDtbpChecksOk() (*bool, bool)`

GetOverrideDtbpChecksOk returns a tuple with the OverrideDtbpChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDtbpChecks

`func (o *OptionOrderInput) SetOverrideDtbpChecks(v bool)`

SetOverrideDtbpChecks sets OverrideDtbpChecks field to given value.

### HasOverrideDtbpChecks

`func (o *OptionOrderInput) HasOverrideDtbpChecks() bool`

HasOverrideDtbpChecks returns a boolean if a field has been set.

### GetPrice

`func (o *OptionOrderInput) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OptionOrderInput) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OptionOrderInput) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OptionOrderInput) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *OptionOrderInput) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OptionOrderInput) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OptionOrderInput) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OptionOrderInput) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRefId

`func (o *OptionOrderInput) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *OptionOrderInput) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *OptionOrderInput) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *OptionOrderInput) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetTimeInForce

`func (o *OptionOrderInput) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OptionOrderInput) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OptionOrderInput) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OptionOrderInput) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetTrigger

`func (o *OptionOrderInput) GetTrigger() Trigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *OptionOrderInput) GetTriggerOk() (*Trigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *OptionOrderInput) SetTrigger(v Trigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *OptionOrderInput) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetType

`func (o *OptionOrderInput) GetType() ExecutionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OptionOrderInput) GetTypeOk() (*ExecutionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OptionOrderInput) SetType(v ExecutionType)`

SetType sets Type field to given value.

### HasType

`func (o *OptionOrderInput) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


