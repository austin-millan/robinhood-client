# \FundamentalsAPI

All URIs are relative to *http://api.robinhood.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFundamentals**](FundamentalsAPI.md#GetFundamentals) | **Get** /fundamentals/ | getFundamentals
[**GetSymbolFundamentals**](FundamentalsAPI.md#GetSymbolFundamentals) | **Get** /fundamentals/{symbol}/ | getSymbolFundamentals



## GetFundamentals

> PaginatedFundamentalsData GetFundamentals(ctx).Symbols(symbols).Execute()

getFundamentals

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
    symbols := "symbols_example" // string | symbols (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundamentalsAPI.GetFundamentals(context.Background()).Symbols(symbols).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetFundamentals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundamentals`: PaginatedFundamentalsData
    fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetFundamentals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFundamentalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbols** | **string** | symbols | 

### Return type

[**PaginatedFundamentalsData**](PaginatedFundamentalsData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSymbolFundamentals

> FundamentalsData GetSymbolFundamentals(ctx, symbol).Execute()

getSymbolFundamentals

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
    symbol := "symbol_example" // string | symbol

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundamentalsAPI.GetSymbolFundamentals(context.Background(), symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetSymbolFundamentals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSymbolFundamentals`: FundamentalsData
    fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetSymbolFundamentals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string** | symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSymbolFundamentalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FundamentalsData**](FundamentalsData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

