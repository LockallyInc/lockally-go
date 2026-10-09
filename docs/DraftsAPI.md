# \DraftsAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DraftsDraftIDApprovePost**](DraftsAPI.md#V1DraftsDraftIDApprovePost) | **Post** /v1/drafts/{draftID}/approve | Approve a pending draft (human)
[**V1DraftsDraftIDCancelPost**](DraftsAPI.md#V1DraftsDraftIDCancelPost) | **Post** /v1/drafts/{draftID}/cancel | Withdraw a pending draft
[**V1DraftsDraftIDGet**](DraftsAPI.md#V1DraftsDraftIDGet) | **Get** /v1/drafts/{draftID} | Get a draft
[**V1DraftsDraftIDRejectPost**](DraftsAPI.md#V1DraftsDraftIDRejectPost) | **Post** /v1/drafts/{draftID}/reject | Reject a pending draft (human)
[**V1DraftsGet**](DraftsAPI.md#V1DraftsGet) | **Get** /v1/drafts | List drafts
[**V1InboxesMailboxDraftsPost**](DraftsAPI.md#V1InboxesMailboxDraftsPost) | **Post** /v1/inboxes/{mailbox}/drafts | Propose a new conversation as a draft
[**V1ThreadsThreadIDDraftsPost**](DraftsAPI.md#V1ThreadsThreadIDDraftsPost) | **Post** /v1/threads/{threadID}/drafts | Propose a reply as a draft



## V1DraftsDraftIDApprovePost

> map[string]interface{} V1DraftsDraftIDApprovePost(ctx, draftID).Execute()

Approve a pending draft (human)



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
	draftID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DraftsAPI.V1DraftsDraftIDApprovePost(context.Background(), draftID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.V1DraftsDraftIDApprovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DraftsDraftIDApprovePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.V1DraftsDraftIDApprovePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DraftsDraftIDApprovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## V1DraftsDraftIDCancelPost

> map[string]interface{} V1DraftsDraftIDCancelPost(ctx, draftID).Execute()

Withdraw a pending draft



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
	draftID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DraftsAPI.V1DraftsDraftIDCancelPost(context.Background(), draftID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.V1DraftsDraftIDCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DraftsDraftIDCancelPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.V1DraftsDraftIDCancelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DraftsDraftIDCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## V1DraftsDraftIDGet

> map[string]interface{} V1DraftsDraftIDGet(ctx, draftID).Execute()

Get a draft

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
	draftID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DraftsAPI.V1DraftsDraftIDGet(context.Background(), draftID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.V1DraftsDraftIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DraftsDraftIDGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.V1DraftsDraftIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DraftsDraftIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## V1DraftsDraftIDRejectPost

> map[string]interface{} V1DraftsDraftIDRejectPost(ctx, draftID).Execute()

Reject a pending draft (human)



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
	draftID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DraftsAPI.V1DraftsDraftIDRejectPost(context.Background(), draftID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.V1DraftsDraftIDRejectPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DraftsDraftIDRejectPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.V1DraftsDraftIDRejectPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DraftsDraftIDRejectPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## V1DraftsGet

> map[string]interface{} V1DraftsGet(ctx).Status(status).Limit(limit).Execute()

List drafts



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
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DraftsAPI.V1DraftsGet(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.V1DraftsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DraftsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.V1DraftsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DraftsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

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


## V1InboxesMailboxDraftsPost

> map[string]interface{} V1InboxesMailboxDraftsPost(ctx, mailbox).IdempotencyKey(idempotencyKey).Execute()

Propose a new conversation as a draft



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
	mailbox := "mailbox_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DraftsAPI.V1InboxesMailboxDraftsPost(context.Background(), mailbox).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.V1InboxesMailboxDraftsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InboxesMailboxDraftsPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.V1InboxesMailboxDraftsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailbox** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InboxesMailboxDraftsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

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


## V1ThreadsThreadIDDraftsPost

> map[string]interface{} V1ThreadsThreadIDDraftsPost(ctx, threadID).IdempotencyKey(idempotencyKey).Execute()

Propose a reply as a draft



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
	idempotencyKey := "idempotencyKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DraftsAPI.V1ThreadsThreadIDDraftsPost(context.Background(), threadID).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.V1ThreadsThreadIDDraftsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ThreadsThreadIDDraftsPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.V1ThreadsThreadIDDraftsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ThreadsThreadIDDraftsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

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

