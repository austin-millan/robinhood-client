# \QuoteAPI

All URIs are relative to *http://api.robinhood.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQuotes**](QuoteAPI.md#GetQuotes) | **Get** /quotes/ | getQuotes
[**GetSymbolQuote**](QuoteAPI.md#GetSymbolQuote) | **Get** /quotes/{symbol}/ | getSymbolQuote



## GetQuotes

> PaginatedQuoteData GetQuotes(ctx).Symbols(symbols).Execute()

getQuotes

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
    symbols := "TSLA,AMD" // string | symbols (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuoteAPI.GetQuotes(context.Background()).Symbols(symbols).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuoteAPI.GetQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuotes`: PaginatedQuoteData
    fmt.Fprintf(os.Stdout, "Response from `QuoteAPI.GetQuotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbols** | **string** | symbols | 

### Return type

[**PaginatedQuoteData**](PaginatedQuoteData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSymbolQuote

> QuoteData GetSymbolQuote(ctx, symbol).Execute()

getSymbolQuote

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
    resp, r, err := apiClient.QuoteAPI.GetSymbolQuote(context.Background(), symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuoteAPI.GetSymbolQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSymbolQuote`: QuoteData
    fmt.Fprintf(os.Stdout, "Response from `QuoteAPI.GetSymbolQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string** | symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSymbolQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuoteData**](QuoteData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

