# \MoversAPI

All URIs are relative to *http://api.robinhood.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMovers**](MoversAPI.md#GetMovers) | **Get** /midlands/movers/sp500/ | getMovers



## GetMovers

> PaginatedMovers GetMovers(ctx).Direction(direction).Execute()

getMovers

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
    direction := "direction_example" // string | direction (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoversAPI.GetMovers(context.Background()).Direction(direction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoversAPI.GetMovers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMovers`: PaginatedMovers
    fmt.Fprintf(os.Stdout, "Response from `MoversAPI.GetMovers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMoversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **direction** | **string** | direction | 

### Return type

[**PaginatedMovers**](PaginatedMovers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

