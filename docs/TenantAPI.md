# \TenantAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTenant**](TenantAPI.md#GetTenant) | **Get** /v1/tenant | Get the calling tenant
[**GetUsage**](TenantAPI.md#GetUsage) | **Get** /v1/usage | Usage snapshot



## GetTenant

> Tenant GetTenant(ctx).Execute()

Get the calling tenant



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
	resp, r, err := apiClient.TenantAPI.GetTenant(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenant`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetTenant`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern


### Return type

[**Tenant**](Tenant.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsage

> GetUsage200Response GetUsage(ctx).Execute()

Usage snapshot



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
	resp, r, err := apiClient.TenantAPI.GetUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsage`: GetUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageRequest struct via the builder pattern


### Return type

[**GetUsage200Response**](GetUsage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

