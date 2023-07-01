# MarketHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClosesAt** | Pointer to **time.Time** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**ExtendedClosesAt** | Pointer to **time.Time** |  | [optional] 
**ExtendedOpensAt** | Pointer to **time.Time** |  | [optional] 
**IsOpen** | Pointer to **bool** |  | [optional] 
**NextOpenHours** | Pointer to **string** |  | [optional] 
**OpensAt** | Pointer to **time.Time** |  | [optional] 
**PreviousOpenHours** | Pointer to **string** |  | [optional] 

## Methods

### NewMarketHours

`func NewMarketHours() *MarketHours`

NewMarketHours instantiates a new MarketHours object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketHoursWithDefaults

`func NewMarketHoursWithDefaults() *MarketHours`

NewMarketHoursWithDefaults instantiates a new MarketHours object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClosesAt

`func (o *MarketHours) GetClosesAt() time.Time`

GetClosesAt returns the ClosesAt field if non-nil, zero value otherwise.

### GetClosesAtOk

`func (o *MarketHours) GetClosesAtOk() (*time.Time, bool)`

GetClosesAtOk returns a tuple with the ClosesAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosesAt

`func (o *MarketHours) SetClosesAt(v time.Time)`

SetClosesAt sets ClosesAt field to given value.

### HasClosesAt

`func (o *MarketHours) HasClosesAt() bool`

HasClosesAt returns a boolean if a field has been set.

### GetDate

`func (o *MarketHours) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MarketHours) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MarketHours) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *MarketHours) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetExtendedClosesAt

`func (o *MarketHours) GetExtendedClosesAt() time.Time`

GetExtendedClosesAt returns the ExtendedClosesAt field if non-nil, zero value otherwise.

### GetExtendedClosesAtOk

`func (o *MarketHours) GetExtendedClosesAtOk() (*time.Time, bool)`

GetExtendedClosesAtOk returns a tuple with the ExtendedClosesAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedClosesAt

`func (o *MarketHours) SetExtendedClosesAt(v time.Time)`

SetExtendedClosesAt sets ExtendedClosesAt field to given value.

### HasExtendedClosesAt

`func (o *MarketHours) HasExtendedClosesAt() bool`

HasExtendedClosesAt returns a boolean if a field has been set.

### GetExtendedOpensAt

`func (o *MarketHours) GetExtendedOpensAt() time.Time`

GetExtendedOpensAt returns the ExtendedOpensAt field if non-nil, zero value otherwise.

### GetExtendedOpensAtOk

`func (o *MarketHours) GetExtendedOpensAtOk() (*time.Time, bool)`

GetExtendedOpensAtOk returns a tuple with the ExtendedOpensAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedOpensAt

`func (o *MarketHours) SetExtendedOpensAt(v time.Time)`

SetExtendedOpensAt sets ExtendedOpensAt field to given value.

### HasExtendedOpensAt

`func (o *MarketHours) HasExtendedOpensAt() bool`

HasExtendedOpensAt returns a boolean if a field has been set.

### GetIsOpen

`func (o *MarketHours) GetIsOpen() bool`

GetIsOpen returns the IsOpen field if non-nil, zero value otherwise.

### GetIsOpenOk

`func (o *MarketHours) GetIsOpenOk() (*bool, bool)`

GetIsOpenOk returns a tuple with the IsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpen

`func (o *MarketHours) SetIsOpen(v bool)`

SetIsOpen sets IsOpen field to given value.

### HasIsOpen

`func (o *MarketHours) HasIsOpen() bool`

HasIsOpen returns a boolean if a field has been set.

### GetNextOpenHours

`func (o *MarketHours) GetNextOpenHours() string`

GetNextOpenHours returns the NextOpenHours field if non-nil, zero value otherwise.

### GetNextOpenHoursOk

`func (o *MarketHours) GetNextOpenHoursOk() (*string, bool)`

GetNextOpenHoursOk returns a tuple with the NextOpenHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOpenHours

`func (o *MarketHours) SetNextOpenHours(v string)`

SetNextOpenHours sets NextOpenHours field to given value.

### HasNextOpenHours

`func (o *MarketHours) HasNextOpenHours() bool`

HasNextOpenHours returns a boolean if a field has been set.

### GetOpensAt

`func (o *MarketHours) GetOpensAt() time.Time`

GetOpensAt returns the OpensAt field if non-nil, zero value otherwise.

### GetOpensAtOk

`func (o *MarketHours) GetOpensAtOk() (*time.Time, bool)`

GetOpensAtOk returns a tuple with the OpensAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpensAt

`func (o *MarketHours) SetOpensAt(v time.Time)`

SetOpensAt sets OpensAt field to given value.

### HasOpensAt

`func (o *MarketHours) HasOpensAt() bool`

HasOpensAt returns a boolean if a field has been set.

### GetPreviousOpenHours

`func (o *MarketHours) GetPreviousOpenHours() string`

GetPreviousOpenHours returns the PreviousOpenHours field if non-nil, zero value otherwise.

### GetPreviousOpenHoursOk

`func (o *MarketHours) GetPreviousOpenHoursOk() (*string, bool)`

GetPreviousOpenHoursOk returns a tuple with the PreviousOpenHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousOpenHours

`func (o *MarketHours) SetPreviousOpenHours(v string)`

SetPreviousOpenHours sets PreviousOpenHours field to given value.

### HasPreviousOpenHours

`func (o *MarketHours) HasPreviousOpenHours() bool`

HasPreviousOpenHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


