# \AiAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClassifyThread**](AiAPI.md#ClassifyThread) | **Post** /v1/threads/{threadID}/classify | LLM-classify a thread
[**CreateAiUnitsCheckout**](AiAPI.md#CreateAiUnitsCheckout) | **Post** /v1/billing/ai-units/checkout | Buy prepaid AI units
[**GetAiConfig**](AiAPI.md#GetAiConfig) | **Get** /v1/ai-config | Read the tenant&#39;s AI configuration
[**UpdateAiConfig**](AiAPI.md#UpdateAiConfig) | **Put** /v1/ai-config | Configure the AI tier



## ClassifyThread

> ThreadClassification ClassifyThread(ctx, threadID).Refresh(refresh).Execute()

LLM-classify a thread



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
	threadID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	refresh := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.ClassifyThread(context.Background(), threadID).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.ClassifyThread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClassifyThread`: ThreadClassification
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.ClassifyThread`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClassifyThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refresh** | **bool** |  | 

### Return type

[**ThreadClassification**](ThreadClassification.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiUnitsCheckout

> CreateAiUnitsCheckout200Response CreateAiUnitsCheckout(ctx).Execute()

Buy prepaid AI units



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
	resp, r, err := apiClient.AiAPI.CreateAiUnitsCheckout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.CreateAiUnitsCheckout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAiUnitsCheckout`: CreateAiUnitsCheckout200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.CreateAiUnitsCheckout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiUnitsCheckoutRequest struct via the builder pattern


### Return type

[**CreateAiUnitsCheckout200Response**](CreateAiUnitsCheckout200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiConfig

> AiConfig GetAiConfig(ctx).Execute()

Read the tenant's AI configuration



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
	resp, r, err := apiClient.AiAPI.GetAiConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiConfig`: AiConfig
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiConfigRequest struct via the builder pattern


### Return type

[**AiConfig**](AiConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAiConfig

> AiConfigUpdate UpdateAiConfig(ctx).Execute()

Configure the AI tier



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
	resp, r, err := apiClient.AiAPI.UpdateAiConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.UpdateAiConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAiConfig`: AiConfigUpdate
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.UpdateAiConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAiConfigRequest struct via the builder pattern


### Return type

[**AiConfigUpdate**](AiConfigUpdate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

