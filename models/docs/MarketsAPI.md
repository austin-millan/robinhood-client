# \MarketsAPI

All URIs are relative to *http://api.robinhood.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMArketHours**](MarketsAPI.md#GetMArketHours) | **Get** /markets/{mic}/hours/{date}/ | getMArketHours
[**GetMarkets**](MarketsAPI.md#GetMarkets) | **Get** /markets | getMarkets



## GetMArketHours

> MarketHours GetMArketHours(ctx, date, mic).Execute()

getMArketHours

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
    date := "date_example" // string | date
    mic := "mic_example" // string | mic

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsAPI.GetMArketHours(context.Background(), date, mic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.GetMArketHours``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMArketHours`: MarketHours
    fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.GetMArketHours`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**date** | **string** | date | 
**mic** | **string** | mic | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMArketHoursRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MarketHours**](MarketHours.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarkets

> PaginatedMarketData GetMarkets(ctx).Execute()

getMarkets

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsAPI.GetMarkets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.GetMarkets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarkets`: PaginatedMarketData
    fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.GetMarkets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketsRequest struct via the builder pattern


### Return type

[**PaginatedMarketData**](PaginatedMarketData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

