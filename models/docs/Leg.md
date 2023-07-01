# Leg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Option** | Pointer to **string** |  | [optional] 
**PositionEffect** | Pointer to **string** |  | [optional] 
**RatioQuantity** | Pointer to **string** |  | [optional] 
**Side** | Pointer to [**Side**](Side.md) |  | [optional] 

## Methods

### NewLeg

`func NewLeg() *Leg`

NewLeg instantiates a new Leg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegWithDefaults

`func NewLegWithDefaults() *Leg`

NewLegWithDefaults instantiates a new Leg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOption

`func (o *Leg) GetOption() string`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *Leg) GetOptionOk() (*string, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *Leg) SetOption(v string)`

SetOption sets Option field to given value.

### HasOption

`func (o *Leg) HasOption() bool`

HasOption returns a boolean if a field has been set.

### GetPositionEffect

`func (o *Leg) GetPositionEffect() string`

GetPositionEffect returns the PositionEffect field if non-nil, zero value otherwise.

### GetPositionEffectOk

`func (o *Leg) GetPositionEffectOk() (*string, bool)`

GetPositionEffectOk returns a tuple with the PositionEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionEffect

`func (o *Leg) SetPositionEffect(v string)`

SetPositionEffect sets PositionEffect field to given value.

### HasPositionEffect

`func (o *Leg) HasPositionEffect() bool`

HasPositionEffect returns a boolean if a field has been set.

### GetRatioQuantity

`func (o *Leg) GetRatioQuantity() string`

GetRatioQuantity returns the RatioQuantity field if non-nil, zero value otherwise.

### GetRatioQuantityOk

`func (o *Leg) GetRatioQuantityOk() (*string, bool)`

GetRatioQuantityOk returns a tuple with the RatioQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatioQuantity

`func (o *Leg) SetRatioQuantity(v string)`

SetRatioQuantity sets RatioQuantity field to given value.

### HasRatioQuantity

`func (o *Leg) HasRatioQuantity() bool`

HasRatioQuantity returns a boolean if a field has been set.

### GetSide

`func (o *Leg) GetSide() Side`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *Leg) GetSideOk() (*Side, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *Leg) SetSide(v Side)`

SetSide sets Side field to given value.

### HasSide

`func (o *Leg) HasSide() bool`

HasSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


