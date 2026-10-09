# \AdminAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AdminLoginPost**](AdminAPI.md#V1AdminLoginPost) | **Post** /v1/admin/login | Tenant-admin email+password login
[**V1AdminLogoutPost**](AdminAPI.md#V1AdminLogoutPost) | **Post** /v1/admin/logout | Invalidate the current admin session
[**V1AdminMeGet**](AdminAPI.md#V1AdminMeGet) | **Get** /v1/admin/me | Get the current admin + tenant



## V1AdminLoginPost

> V1AdminLoginPost200Response V1AdminLoginPost(ctx).V1AdminLoginPostRequest(v1AdminLoginPostRequest).Execute()

Tenant-admin email+password login



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
	v1AdminLoginPostRequest := *openapiclient.NewV1AdminLoginPostRequest("Email_example", "Password_example") // V1AdminLoginPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.V1AdminLoginPost(context.Background()).V1AdminLoginPostRequest(v1AdminLoginPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.V1AdminLoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AdminLoginPost`: V1AdminLoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.V1AdminLoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AdminLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1AdminLoginPostRequest** | [**V1AdminLoginPostRequest**](V1AdminLoginPostRequest.md) |  | 

### Return type

[**V1AdminLoginPost200Response**](V1AdminLoginPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AdminLogoutPost

> V1AdminLogoutPost(ctx).Execute()

Invalidate the current admin session



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
	r, err := apiClient.AdminAPI.V1AdminLogoutPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.V1AdminLogoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AdminLogoutPostRequest struct via the builder pattern


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


## V1AdminMeGet

> V1AdminMeGet200Response V1AdminMeGet(ctx).Execute()

Get the current admin + tenant



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
	resp, r, err := apiClient.AdminAPI.V1AdminMeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.V1AdminMeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AdminMeGet`: V1AdminMeGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.V1AdminMeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AdminMeGetRequest struct via the builder pattern


### Return type

[**V1AdminMeGet200Response**](V1AdminMeGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

