# \AccountsAPI

All URIs are relative to *http://api.robinhood.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccounts**](AccountsAPI.md#GetAccounts) | **Get** /accounts/ | getAccounts
[**GetPortfolio**](AccountsAPI.md#GetPortfolio) | **Get** /accounts/{accountId}/portfolio/ | getPortfolio
[**GetPosition**](AccountsAPI.md#GetPosition) | **Get** /accounts/{accountId}/positions/{positionId}/ | getPosition
[**GetPositions**](AccountsAPI.md#GetPositions) | **Get** /accounts/{accountId}/positions/ | getPositions



## GetAccounts

> PaginatedAccountInfo GetAccounts(ctx).Execute()

getAccounts

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
    resp, r, err := apiClient.AccountsAPI.GetAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccounts`: PaginatedAccountInfo
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsRequest struct via the builder pattern


### Return type

[**PaginatedAccountInfo**](PaginatedAccountInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolio

> Portfolio GetPortfolio(ctx, accountId).Execute()

getPortfolio

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
    accountId := "accountId_example" // string | accountId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.GetPortfolio(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetPortfolio``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPortfolio`: Portfolio
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetPortfolio`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | accountId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Portfolio**](Portfolio.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPosition

> Position GetPosition(ctx, accountId, positionId).Execute()

getPosition

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
    accountId := "accountId_example" // string | accountId
    positionId := "positionId_example" // string | positionId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.GetPosition(context.Background(), accountId, positionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetPosition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPosition`: Position
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetPosition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | accountId | 
**positionId** | **string** | positionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Position**](Position.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPositions

> PaginatedPosition GetPositions(ctx, accountId).Nonzero(nonzero).Execute()

getPositions

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
    accountId := "accountId_example" // string | accountId
    nonzero := true // bool | nonzero (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.GetPositions(context.Background(), accountId).Nonzero(nonzero).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPositions`: PaginatedPosition
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetPositions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | accountId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nonzero** | **bool** | nonzero | 

### Return type

[**PaginatedPosition**](PaginatedPosition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

