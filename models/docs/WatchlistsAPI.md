# \WatchlistsAPI

All URIs are relative to *http://api.robinhood.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkAddWatchlists**](WatchlistsAPI.md#BulkAddWatchlists) | **Post** /watchlists/Default/bulk_add/ | bulkAddWatchlists
[**CreateWatchlist**](WatchlistsAPI.md#CreateWatchlist) | **Post** /watchlists/ | createWatchlist
[**DeleteInstrumentFromWatchlist**](WatchlistsAPI.md#DeleteInstrumentFromWatchlist) | **Delete** /watchlists/{name}/{instrumentId} | deleteInstrumentFromWatchlist
[**GetWatchlistByName**](WatchlistsAPI.md#GetWatchlistByName) | **Get** /watchlists/{name}/ | getWatchlistByName
[**GetWatchlists**](WatchlistsAPI.md#GetWatchlists) | **Get** /watchlists/ | getWatchlists



## BulkAddWatchlists

> WatchListsData BulkAddWatchlists(ctx).Symbols(symbols).Execute()

bulkAddWatchlists

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
    resp, r, err := apiClient.WatchlistsAPI.BulkAddWatchlists(context.Background()).Symbols(symbols).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistsAPI.BulkAddWatchlists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkAddWatchlists`: WatchListsData
    fmt.Fprintf(os.Stdout, "Response from `WatchlistsAPI.BulkAddWatchlists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkAddWatchlistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbols** | **string** | symbols | 

### Return type

[**WatchListsData**](WatchListsData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWatchlist

> WatchListCreateResponse CreateWatchlist(ctx).Name(name).Execute()

createWatchlist

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
    name := "name_example" // string | name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WatchlistsAPI.CreateWatchlist(context.Background()).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistsAPI.CreateWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWatchlist`: WatchListCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `WatchlistsAPI.CreateWatchlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWatchlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | name | 

### Return type

[**WatchListCreateResponse**](WatchListCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstrumentFromWatchlist

> PaginatedWatchListsData DeleteInstrumentFromWatchlist(ctx, instrumentId, name).Execute()

deleteInstrumentFromWatchlist

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
    instrumentId := "instrumentId_example" // string | instrumentId
    name := "name_example" // string | name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WatchlistsAPI.DeleteInstrumentFromWatchlist(context.Background(), instrumentId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistsAPI.DeleteInstrumentFromWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInstrumentFromWatchlist`: PaginatedWatchListsData
    fmt.Fprintf(os.Stdout, "Response from `WatchlistsAPI.DeleteInstrumentFromWatchlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instrumentId** | **string** | instrumentId | 
**name** | **string** | name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstrumentFromWatchlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedWatchListsData**](PaginatedWatchListsData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWatchlistByName

> PaginatedWatchListsData GetWatchlistByName(ctx, name).Execute()

getWatchlistByName

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
    name := "name_example" // string | name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WatchlistsAPI.GetWatchlistByName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistsAPI.GetWatchlistByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWatchlistByName`: PaginatedWatchListsData
    fmt.Fprintf(os.Stdout, "Response from `WatchlistsAPI.GetWatchlistByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWatchlistByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedWatchListsData**](PaginatedWatchListsData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWatchlists

> PaginatedWatchListCreateResponse GetWatchlists(ctx).Execute()

getWatchlists

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
    resp, r, err := apiClient.WatchlistsAPI.GetWatchlists(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistsAPI.GetWatchlists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWatchlists`: PaginatedWatchListCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `WatchlistsAPI.GetWatchlists`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWatchlistsRequest struct via the builder pattern


### Return type

[**PaginatedWatchListCreateResponse**](PaginatedWatchListCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

