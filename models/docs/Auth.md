# Auth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MfaRequired** | Pointer to **string** |  | [optional] 
**MfaType** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewAuth

`func NewAuth() *Auth`

NewAuth instantiates a new Auth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthWithDefaults

`func NewAuthWithDefaults() *Auth`

NewAuthWithDefaults instantiates a new Auth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMfaRequired

`func (o *Auth) GetMfaRequired() string`

GetMfaRequired returns the MfaRequired field if non-nil, zero value otherwise.

### GetMfaRequiredOk

`func (o *Auth) GetMfaRequiredOk() (*string, bool)`

GetMfaRequiredOk returns a tuple with the MfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRequired

`func (o *Auth) SetMfaRequired(v string)`

SetMfaRequired sets MfaRequired field to given value.

### HasMfaRequired

`func (o *Auth) HasMfaRequired() bool`

HasMfaRequired returns a boolean if a field has been set.

### GetMfaType

`func (o *Auth) GetMfaType() string`

GetMfaType returns the MfaType field if non-nil, zero value otherwise.

### GetMfaTypeOk

`func (o *Auth) GetMfaTypeOk() (*string, bool)`

GetMfaTypeOk returns a tuple with the MfaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaType

`func (o *Auth) SetMfaType(v string)`

SetMfaType sets MfaType field to given value.

### HasMfaType

`func (o *Auth) HasMfaType() bool`

HasMfaType returns a boolean if a field has been set.

### GetToken

`func (o *Auth) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Auth) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Auth) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Auth) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


