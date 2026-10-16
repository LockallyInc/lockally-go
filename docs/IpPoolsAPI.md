# \IpPoolsAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDedicatedIPRequest**](IpPoolsAPI.md#CreateDedicatedIPRequest) | **Post** /v1/dedicated-ip-requests | Request a dedicated IP
[**GetIPAssignment**](IpPoolsAPI.md#GetIPAssignment) | **Get** /v1/ip-assignment | Get current IP assignment
[**ListDedicatedIPRequests**](IpPoolsAPI.md#ListDedicatedIPRequests) | **Get** /v1/dedicated-ip-requests | List dedicated IP requests



## CreateDedicatedIPRequest

> DedicatedIPRequest CreateDedicatedIPRequest(ctx).CreateDedicatedIPRequestRequest(createDedicatedIPRequestRequest).Execute()

Request a dedicated IP

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lockallyinc/lockally-go"
)

func main() {
	createDedicatedIPRequestRequest := *openapiclient.NewCreateDedicatedIPRequestRequest("Note_example") // CreateDedicatedIPRequestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpPoolsAPI.CreateDedicatedIPRequest(context.Background()).CreateDedicatedIPRequestRequest(createDedicatedIPRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpPoolsAPI.CreateDedicatedIPRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDedicatedIPRequest`: DedicatedIPRequest
	fmt.Fprintf(os.Stdout, "Response from `IpPoolsAPI.CreateDedicatedIPRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDedicatedIPRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDedicatedIPRequestRequest** | [**CreateDedicatedIPRequestRequest**](CreateDedicatedIPRequestRequest.md) |  | 

### Return type

[**DedicatedIPRequest**](DedicatedIPRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIPAssignment

> GetIPAssignment200Response GetIPAssignment(ctx).Execute()

Get current IP assignment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lockallyinc/lockally-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpPoolsAPI.GetIPAssignment(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpPoolsAPI.GetIPAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIPAssignment`: GetIPAssignment200Response
	fmt.Fprintf(os.Stdout, "Response from `IpPoolsAPI.GetIPAssignment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIPAssignmentRequest struct via the builder pattern


### Return type

[**GetIPAssignment200Response**](GetIPAssignment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDedicatedIPRequests

> ListDedicatedIPRequests200Response ListDedicatedIPRequests(ctx).Execute()

List dedicated IP requests

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lockallyinc/lockally-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpPoolsAPI.ListDedicatedIPRequests(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpPoolsAPI.ListDedicatedIPRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDedicatedIPRequests`: ListDedicatedIPRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `IpPoolsAPI.ListDedicatedIPRequests`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDedicatedIPRequestsRequest struct via the builder pattern


### Return type

[**ListDedicatedIPRequests200Response**](ListDedicatedIPRequests200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

