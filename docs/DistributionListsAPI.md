# \DistributionListsAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDistributionList**](DistributionListsAPI.md#CreateDistributionList) | **Post** /v1/distribution-lists | Create a distribution list
[**DeleteDistributionList**](DistributionListsAPI.md#DeleteDistributionList) | **Delete** /v1/distribution-lists/{address} | Delete a distribution list
[**GetDistributionList**](DistributionListsAPI.md#GetDistributionList) | **Get** /v1/distribution-lists/{address} | Get a distribution list
[**ListDistributionLists**](DistributionListsAPI.md#ListDistributionLists) | **Get** /v1/distribution-lists | List distribution lists
[**ReplaceDistributionListMembers**](DistributionListsAPI.md#ReplaceDistributionListMembers) | **Put** /v1/distribution-lists/{address}/members | Replace distribution list members



## CreateDistributionList

> DistributionListDetail CreateDistributionList(ctx).CreateDistributionListRequest(createDistributionListRequest).Execute()

Create a distribution list

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
	createDistributionListRequest := *openapiclient.NewCreateDistributionListRequest("ListAddress_example") // CreateDistributionListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionListsAPI.CreateDistributionList(context.Background()).CreateDistributionListRequest(createDistributionListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionListsAPI.CreateDistributionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDistributionList`: DistributionListDetail
	fmt.Fprintf(os.Stdout, "Response from `DistributionListsAPI.CreateDistributionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDistributionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDistributionListRequest** | [**CreateDistributionListRequest**](CreateDistributionListRequest.md) |  | 

### Return type

[**DistributionListDetail**](DistributionListDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDistributionList

> DeleteDistributionList(ctx, address).Execute()

Delete a distribution list

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
	address := "address_example" // string | Distribution list email address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DistributionListsAPI.DeleteDistributionList(context.Background(), address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionListsAPI.DeleteDistributionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Distribution list email address | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDistributionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDistributionList

> DistributionListDetail GetDistributionList(ctx, address).Execute()

Get a distribution list

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
	address := "address_example" // string | Distribution list email address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionListsAPI.GetDistributionList(context.Background(), address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionListsAPI.GetDistributionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDistributionList`: DistributionListDetail
	fmt.Fprintf(os.Stdout, "Response from `DistributionListsAPI.GetDistributionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Distribution list email address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDistributionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DistributionListDetail**](DistributionListDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDistributionLists

> ListDistributionLists200Response ListDistributionLists(ctx).Execute()

List distribution lists

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
	resp, r, err := apiClient.DistributionListsAPI.ListDistributionLists(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionListsAPI.ListDistributionLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDistributionLists`: ListDistributionLists200Response
	fmt.Fprintf(os.Stdout, "Response from `DistributionListsAPI.ListDistributionLists`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDistributionListsRequest struct via the builder pattern


### Return type

[**ListDistributionLists200Response**](ListDistributionLists200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceDistributionListMembers

> ReplaceDistributionListMembers200Response ReplaceDistributionListMembers(ctx, address).ReplaceDistributionListMembersRequest(replaceDistributionListMembersRequest).Execute()

Replace distribution list members

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
	address := "address_example" // string | Distribution list email address
	replaceDistributionListMembersRequest := *openapiclient.NewReplaceDistributionListMembersRequest([]string{"Members_example"}) // ReplaceDistributionListMembersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionListsAPI.ReplaceDistributionListMembers(context.Background(), address).ReplaceDistributionListMembersRequest(replaceDistributionListMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionListsAPI.ReplaceDistributionListMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceDistributionListMembers`: ReplaceDistributionListMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `DistributionListsAPI.ReplaceDistributionListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Distribution list email address | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDistributionListMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceDistributionListMembersRequest** | [**ReplaceDistributionListMembersRequest**](ReplaceDistributionListMembersRequest.md) |  | 

### Return type

[**ReplaceDistributionListMembers200Response**](ReplaceDistributionListMembers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

