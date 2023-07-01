# PaginatedQuoteData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **string** |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Previous** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]QuoteData**](QuoteData.md) |  | [optional] 

## Methods

### NewPaginatedQuoteData

`func NewPaginatedQuoteData() *PaginatedQuoteData`

NewPaginatedQuoteData instantiates a new PaginatedQuoteData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedQuoteDataWithDefaults

`func NewPaginatedQuoteDataWithDefaults() *PaginatedQuoteData`

NewPaginatedQuoteDataWithDefaults instantiates a new PaginatedQuoteData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedQuoteData) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedQuoteData) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedQuoteData) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedQuoteData) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedQuoteData) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedQuoteData) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedQuoteData) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedQuoteData) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedQuoteData) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedQuoteData) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedQuoteData) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedQuoteData) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedQuoteData) GetResults() []QuoteData`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedQuoteData) GetResultsOk() (*[]QuoteData, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedQuoteData) SetResults(v []QuoteData)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedQuoteData) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


