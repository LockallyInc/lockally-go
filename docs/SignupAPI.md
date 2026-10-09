# \SignupAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Signup**](SignupAPI.md#Signup) | **Post** /v1/signup | Sign up a new tenant



## Signup

> V1AdminLoginPost200Response Signup(ctx).SignupRequest(signupRequest).Execute()

Sign up a new tenant

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
	signupRequest := *openapiclient.NewSignupRequest("Slug_example", "AdminEmail_example", "Password_example") // SignupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignupAPI.Signup(context.Background()).SignupRequest(signupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignupAPI.Signup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Signup`: V1AdminLoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SignupAPI.Signup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signupRequest** | [**SignupRequest**](SignupRequest.md) |  | 

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

