# \AliasesAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlias**](AliasesAPI.md#CreateAlias) | **Post** /v1/aliases | Create an alias
[**DeleteAlias**](AliasesAPI.md#DeleteAlias) | **Delete** /v1/aliases/{address} | Delete an alias
[**ListAliases**](AliasesAPI.md#ListAliases) | **Get** /v1/aliases | List aliases



## CreateAlias

> Alias CreateAlias(ctx).CreateAliasRequest(createAliasRequest).Execute()

Create an alias



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
	createAliasRequest := *openapiclient.NewCreateAliasRequest("support@acme.com", "alice@acme.com") // CreateAliasRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AliasesAPI.CreateAlias(context.Background()).CreateAliasRequest(createAliasRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.CreateAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlias`: Alias
	fmt.Fprintf(os.Stdout, "Response from `AliasesAPI.CreateAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAliasRequest** | [**CreateAliasRequest**](CreateAliasRequest.md) |  | 

### Return type

[**Alias**](Alias.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlias

> DeleteAlias(ctx, address).Execute()

Delete an alias



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
	address := "address_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AliasesAPI.DeleteAlias(context.Background(), address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.DeleteAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAliasRequest struct via the builder pattern


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


## ListAliases

> ListAliases200Response ListAliases(ctx).Execute()

List aliases

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
	resp, r, err := apiClient.AliasesAPI.ListAliases(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.ListAliases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAliases`: ListAliases200Response
	fmt.Fprintf(os.Stdout, "Response from `AliasesAPI.ListAliases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAliasesRequest struct via the builder pattern


### Return type

[**ListAliases200Response**](ListAliases200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

