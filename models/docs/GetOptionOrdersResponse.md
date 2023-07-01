# GetOptionOrdersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** |  | [optional] 
**Previous** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]OptionOrder**](OptionOrder.md) |  | [optional] 

## Methods

### NewGetOptionOrdersResponse

`func NewGetOptionOrdersResponse() *GetOptionOrdersResponse`

NewGetOptionOrdersResponse instantiates a new GetOptionOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOptionOrdersResponseWithDefaults

`func NewGetOptionOrdersResponseWithDefaults() *GetOptionOrdersResponse`

NewGetOptionOrdersResponseWithDefaults instantiates a new GetOptionOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *GetOptionOrdersResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetOptionOrdersResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetOptionOrdersResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *GetOptionOrdersResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *GetOptionOrdersResponse) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *GetOptionOrdersResponse) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *GetOptionOrdersResponse) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *GetOptionOrdersResponse) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *GetOptionOrdersResponse) GetResults() []OptionOrder`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetOptionOrdersResponse) GetResultsOk() (*[]OptionOrder, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetOptionOrdersResponse) SetResults(v []OptionOrder)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetOptionOrdersResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


