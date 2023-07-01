# PaginatedOptionChain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **string** |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Previous** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]OptionChain**](OptionChain.md) |  | [optional] 

## Methods

### NewPaginatedOptionChain

`func NewPaginatedOptionChain() *PaginatedOptionChain`

NewPaginatedOptionChain instantiates a new PaginatedOptionChain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedOptionChainWithDefaults

`func NewPaginatedOptionChainWithDefaults() *PaginatedOptionChain`

NewPaginatedOptionChainWithDefaults instantiates a new PaginatedOptionChain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedOptionChain) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedOptionChain) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedOptionChain) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedOptionChain) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedOptionChain) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedOptionChain) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedOptionChain) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedOptionChain) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedOptionChain) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedOptionChain) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedOptionChain) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedOptionChain) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedOptionChain) GetResults() []OptionChain`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedOptionChain) GetResultsOk() (*[]OptionChain, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedOptionChain) SetResults(v []OptionChain)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedOptionChain) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


