# BasicInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Citizenship** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**CountryOfResidence** | Pointer to **string** |  | [optional] 
**DateOfBirth** | Pointer to **string** |  | [optional] 
**MaritalStatus** | Pointer to [**MaritalStatus**](MaritalStatus.md) |  | [optional] 
**NumberDependents** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TaxIdSsn** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Zipcode** | Pointer to **string** |  | [optional] 

## Methods

### NewBasicInfo

`func NewBasicInfo() *BasicInfo`

NewBasicInfo instantiates a new BasicInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicInfoWithDefaults

`func NewBasicInfoWithDefaults() *BasicInfo`

NewBasicInfoWithDefaults instantiates a new BasicInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BasicInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BasicInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BasicInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BasicInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCitizenship

`func (o *BasicInfo) GetCitizenship() string`

GetCitizenship returns the Citizenship field if non-nil, zero value otherwise.

### GetCitizenshipOk

`func (o *BasicInfo) GetCitizenshipOk() (*string, bool)`

GetCitizenshipOk returns a tuple with the Citizenship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenship

`func (o *BasicInfo) SetCitizenship(v string)`

SetCitizenship sets Citizenship field to given value.

### HasCitizenship

`func (o *BasicInfo) HasCitizenship() bool`

HasCitizenship returns a boolean if a field has been set.

### GetCity

`func (o *BasicInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BasicInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BasicInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BasicInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryOfResidence

`func (o *BasicInfo) GetCountryOfResidence() string`

GetCountryOfResidence returns the CountryOfResidence field if non-nil, zero value otherwise.

### GetCountryOfResidenceOk

`func (o *BasicInfo) GetCountryOfResidenceOk() (*string, bool)`

GetCountryOfResidenceOk returns a tuple with the CountryOfResidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfResidence

`func (o *BasicInfo) SetCountryOfResidence(v string)`

SetCountryOfResidence sets CountryOfResidence field to given value.

### HasCountryOfResidence

`func (o *BasicInfo) HasCountryOfResidence() bool`

HasCountryOfResidence returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *BasicInfo) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *BasicInfo) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *BasicInfo) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *BasicInfo) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetMaritalStatus

`func (o *BasicInfo) GetMaritalStatus() MaritalStatus`

GetMaritalStatus returns the MaritalStatus field if non-nil, zero value otherwise.

### GetMaritalStatusOk

`func (o *BasicInfo) GetMaritalStatusOk() (*MaritalStatus, bool)`

GetMaritalStatusOk returns a tuple with the MaritalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaritalStatus

`func (o *BasicInfo) SetMaritalStatus(v MaritalStatus)`

SetMaritalStatus sets MaritalStatus field to given value.

### HasMaritalStatus

`func (o *BasicInfo) HasMaritalStatus() bool`

HasMaritalStatus returns a boolean if a field has been set.

### GetNumberDependents

`func (o *BasicInfo) GetNumberDependents() string`

GetNumberDependents returns the NumberDependents field if non-nil, zero value otherwise.

### GetNumberDependentsOk

`func (o *BasicInfo) GetNumberDependentsOk() (*string, bool)`

GetNumberDependentsOk returns a tuple with the NumberDependents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberDependents

`func (o *BasicInfo) SetNumberDependents(v string)`

SetNumberDependents sets NumberDependents field to given value.

### HasNumberDependents

`func (o *BasicInfo) HasNumberDependents() bool`

HasNumberDependents returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *BasicInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *BasicInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *BasicInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *BasicInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetState

`func (o *BasicInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BasicInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BasicInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BasicInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTaxIdSsn

`func (o *BasicInfo) GetTaxIdSsn() string`

GetTaxIdSsn returns the TaxIdSsn field if non-nil, zero value otherwise.

### GetTaxIdSsnOk

`func (o *BasicInfo) GetTaxIdSsnOk() (*string, bool)`

GetTaxIdSsnOk returns a tuple with the TaxIdSsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdSsn

`func (o *BasicInfo) SetTaxIdSsn(v string)`

SetTaxIdSsn sets TaxIdSsn field to given value.

### HasTaxIdSsn

`func (o *BasicInfo) HasTaxIdSsn() bool`

HasTaxIdSsn returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BasicInfo) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BasicInfo) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BasicInfo) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BasicInfo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *BasicInfo) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BasicInfo) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BasicInfo) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *BasicInfo) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetZipcode

`func (o *BasicInfo) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *BasicInfo) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *BasicInfo) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *BasicInfo) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


