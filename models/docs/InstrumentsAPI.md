# \InstrumentsAPI

All URIs are relative to *http://api.robinhood.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstrument**](InstrumentsAPI.md#GetInstrument) | **Get** /instruments/{instrument_id}/ | getInstrument
[**GetInstrumentSplits**](InstrumentsAPI.md#GetInstrumentSplits) | **Get** /instruments/{instrument_id}/splits/ | getInstrumentSplits
[**GetInstruments**](InstrumentsAPI.md#GetInstruments) | **Get** /instruments/ | getInstruments



## GetInstrument

> InstrumentData GetInstrument(ctx, instrumentId).Execute()

getInstrument

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    instrumentId := "instrumentId_example" // string | instrument_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstrumentsAPI.GetInstrument(context.Background(), instrumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstrumentsAPI.GetInstrument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstrument`: InstrumentData
    fmt.Fprintf(os.Stdout, "Response from `InstrumentsAPI.GetInstrument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instrumentId** | **string** | instrument_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstrumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstrumentData**](InstrumentData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstrumentSplits

> PaginatedInstrumentSplit GetInstrumentSplits(ctx, instrumentId).Execute()

getInstrumentSplits

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    instrumentId := "instrumentId_example" // string | instrument_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstrumentsAPI.GetInstrumentSplits(context.Background(), instrumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstrumentsAPI.GetInstrumentSplits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstrumentSplits`: PaginatedInstrumentSplit
    fmt.Fprintf(os.Stdout, "Response from `InstrumentsAPI.GetInstrumentSplits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instrumentId** | **string** | instrument_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstrumentSplitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedInstrumentSplit**](PaginatedInstrumentSplit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstruments

> PaginatedInstrumentData GetInstruments(ctx).Query(query).Execute()

getInstruments

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    query := "query_example" // string | query (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstrumentsAPI.GetInstruments(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstrumentsAPI.GetInstruments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstruments`: PaginatedInstrumentData
    fmt.Fprintf(os.Stdout, "Response from `InstrumentsAPI.GetInstruments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstrumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | query | 

### Return type

[**PaginatedInstrumentData**](PaginatedInstrumentData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

