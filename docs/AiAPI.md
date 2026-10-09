# \AiAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AiConfigGet**](AiAPI.md#V1AiConfigGet) | **Get** /v1/ai-config | Read the tenant&#39;s AI configuration
[**V1AiConfigPut**](AiAPI.md#V1AiConfigPut) | **Put** /v1/ai-config | Configure the AI tier
[**V1BillingAiUnitsCheckoutPost**](AiAPI.md#V1BillingAiUnitsCheckoutPost) | **Post** /v1/billing/ai-units/checkout | Buy prepaid AI units
[**V1ThreadsThreadIDClassifyPost**](AiAPI.md#V1ThreadsThreadIDClassifyPost) | **Post** /v1/threads/{threadID}/classify | LLM-classify a thread



## V1AiConfigGet

> map[string]interface{} V1AiConfigGet(ctx).Execute()

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
	resp, r, err := apiClient.AiAPI.V1AiConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.V1AiConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AiConfigGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.V1AiConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AiConfigGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AiConfigPut

> map[string]interface{} V1AiConfigPut(ctx).Execute()

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
	resp, r, err := apiClient.AiAPI.V1AiConfigPut(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.V1AiConfigPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AiConfigPut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.V1AiConfigPut`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AiConfigPutRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BillingAiUnitsCheckoutPost

> map[string]interface{} V1BillingAiUnitsCheckoutPost(ctx).Execute()

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
	resp, r, err := apiClient.AiAPI.V1BillingAiUnitsCheckoutPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.V1BillingAiUnitsCheckoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1BillingAiUnitsCheckoutPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.V1BillingAiUnitsCheckoutPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingAiUnitsCheckoutPostRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ThreadsThreadIDClassifyPost

> map[string]interface{} V1ThreadsThreadIDClassifyPost(ctx, threadID).Refresh(refresh).Execute()

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
	resp, r, err := apiClient.AiAPI.V1ThreadsThreadIDClassifyPost(context.Background(), threadID).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.V1ThreadsThreadIDClassifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ThreadsThreadIDClassifyPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.V1ThreadsThreadIDClassifyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ThreadsThreadIDClassifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refresh** | **bool** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

