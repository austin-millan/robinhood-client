# OptionOrderLeg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Executions** | Pointer to [**[]OptionOrderLegExecutionsInner**](OptionOrderLegExecutionsInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Option** | Pointer to **string** |  | [optional] 
**PositionEffect** | Pointer to [**PositionEffect**](PositionEffect.md) |  | [optional] 
**RatioQuantity** | Pointer to **float32** |  | [optional] 
**Side** | Pointer to [**Side**](Side.md) |  | [optional] 

## Methods

### NewOptionOrderLeg

`func NewOptionOrderLeg() *OptionOrderLeg`

NewOptionOrderLeg instantiates a new OptionOrderLeg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionOrderLegWithDefaults

`func NewOptionOrderLegWithDefaults() *OptionOrderLeg`

NewOptionOrderLegWithDefaults instantiates a new OptionOrderLeg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutions

`func (o *OptionOrderLeg) GetExecutions() []OptionOrderLegExecutionsInner`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *OptionOrderLeg) GetExecutionsOk() (*[]OptionOrderLegExecutionsInner, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *OptionOrderLeg) SetExecutions(v []OptionOrderLegExecutionsInner)`

SetExecutions sets Executions field to given value.

### HasExecutions

`func (o *OptionOrderLeg) HasExecutions() bool`

HasExecutions returns a boolean if a field has been set.

### GetId

`func (o *OptionOrderLeg) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptionOrderLeg) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptionOrderLeg) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OptionOrderLeg) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOption

`func (o *OptionOrderLeg) GetOption() string`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *OptionOrderLeg) GetOptionOk() (*string, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *OptionOrderLeg) SetOption(v string)`

SetOption sets Option field to given value.

### HasOption

`func (o *OptionOrderLeg) HasOption() bool`

HasOption returns a boolean if a field has been set.

### GetPositionEffect

`func (o *OptionOrderLeg) GetPositionEffect() PositionEffect`

GetPositionEffect returns the PositionEffect field if non-nil, zero value otherwise.

### GetPositionEffectOk

`func (o *OptionOrderLeg) GetPositionEffectOk() (*PositionEffect, bool)`

GetPositionEffectOk returns a tuple with the PositionEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionEffect

`func (o *OptionOrderLeg) SetPositionEffect(v PositionEffect)`

SetPositionEffect sets PositionEffect field to given value.

### HasPositionEffect

`func (o *OptionOrderLeg) HasPositionEffect() bool`

HasPositionEffect returns a boolean if a field has been set.

### GetRatioQuantity

`func (o *OptionOrderLeg) GetRatioQuantity() float32`

GetRatioQuantity returns the RatioQuantity field if non-nil, zero value otherwise.

### GetRatioQuantityOk

`func (o *OptionOrderLeg) GetRatioQuantityOk() (*float32, bool)`

GetRatioQuantityOk returns a tuple with the RatioQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatioQuantity

`func (o *OptionOrderLeg) SetRatioQuantity(v float32)`

SetRatioQuantity sets RatioQuantity field to given value.

### HasRatioQuantity

`func (o *OptionOrderLeg) HasRatioQuantity() bool`

HasRatioQuantity returns a boolean if a field has been set.

### GetSide

`func (o *OptionOrderLeg) GetSide() Side`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OptionOrderLeg) GetSideOk() (*Side, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OptionOrderLeg) SetSide(v Side)`

SetSide sets Side field to given value.

### HasSide

`func (o *OptionOrderLeg) HasSide() bool`

HasSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


