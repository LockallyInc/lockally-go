# \BillingAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingCheckout**](BillingAPI.md#CreateBillingCheckout) | **Post** /v1/billing/checkout | Create a plan checkout session
[**CreateUnitsCheckout**](BillingAPI.md#CreateUnitsCheckout) | **Post** /v1/billing/units/checkout | Create a send-units checkout session
[**GetBilling**](BillingAPI.md#GetBilling) | **Get** /v1/billing | Get billing status



## CreateBillingCheckout

> CreateBillingCheckout200Response CreateBillingCheckout(ctx).CreateBillingCheckoutRequest(createBillingCheckoutRequest).Execute()

Create a plan checkout session

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
	createBillingCheckoutRequest := *openapiclient.NewCreateBillingCheckoutRequest("Plan_example") // CreateBillingCheckoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.CreateBillingCheckout(context.Background()).CreateBillingCheckoutRequest(createBillingCheckoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.CreateBillingCheckout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBillingCheckout`: CreateBillingCheckout200Response
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.CreateBillingCheckout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillingCheckoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBillingCheckoutRequest** | [**CreateBillingCheckoutRequest**](CreateBillingCheckoutRequest.md) |  | 

### Return type

[**CreateBillingCheckout200Response**](CreateBillingCheckout200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUnitsCheckout

> CreateUnitsCheckout200Response CreateUnitsCheckout(ctx).CreateUnitsCheckoutRequest(createUnitsCheckoutRequest).Execute()

Create a send-units checkout session

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
	createUnitsCheckoutRequest := *openapiclient.NewCreateUnitsCheckoutRequest("Bundle_example") // CreateUnitsCheckoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.CreateUnitsCheckout(context.Background()).CreateUnitsCheckoutRequest(createUnitsCheckoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.CreateUnitsCheckout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUnitsCheckout`: CreateUnitsCheckout200Response
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.CreateUnitsCheckout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUnitsCheckoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUnitsCheckoutRequest** | [**CreateUnitsCheckoutRequest**](CreateUnitsCheckoutRequest.md) |  | 

### Return type

[**CreateUnitsCheckout200Response**](CreateUnitsCheckout200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBilling

> BillingStatus GetBilling(ctx).Execute()

Get billing status

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
	resp, r, err := apiClient.BillingAPI.GetBilling(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBilling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBilling`: BillingStatus
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBilling`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingRequest struct via the builder pattern


### Return type

[**BillingStatus**](BillingStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

