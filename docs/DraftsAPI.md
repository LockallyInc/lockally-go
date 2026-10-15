# \DraftsAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveDraft**](DraftsAPI.md#ApproveDraft) | **Post** /v1/drafts/{draftID}/approve | Approve a pending draft (human)
[**CancelDraft**](DraftsAPI.md#CancelDraft) | **Post** /v1/drafts/{draftID}/cancel | Withdraw a pending draft
[**CreateConversationDraft**](DraftsAPI.md#CreateConversationDraft) | **Post** /v1/inboxes/{mailbox}/drafts | Propose a new conversation as a draft
[**CreateReplyDraft**](DraftsAPI.md#CreateReplyDraft) | **Post** /v1/threads/{threadID}/drafts | Propose a reply as a draft
[**GetDraft**](DraftsAPI.md#GetDraft) | **Get** /v1/drafts/{draftID} | Get a draft
[**ListDrafts**](DraftsAPI.md#ListDrafts) | **Get** /v1/drafts | List drafts
[**RejectDraft**](DraftsAPI.md#RejectDraft) | **Post** /v1/drafts/{draftID}/reject | Reject a pending draft (human)



## ApproveDraft

> DraftResult ApproveDraft(ctx, draftID).Execute()

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
	resp, r, err := apiClient.DraftsAPI.ApproveDraft(context.Background(), draftID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.ApproveDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveDraft`: DraftResult
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.ApproveDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DraftResult**](DraftResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelDraft

> Draft CancelDraft(ctx, draftID).Execute()

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
	resp, r, err := apiClient.DraftsAPI.CancelDraft(context.Background(), draftID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.CancelDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelDraft`: Draft
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.CancelDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Draft**](Draft.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConversationDraft

> DraftResult CreateConversationDraft(ctx, mailbox).IdempotencyKey(idempotencyKey).Execute()

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
	resp, r, err := apiClient.DraftsAPI.CreateConversationDraft(context.Background(), mailbox).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.CreateConversationDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConversationDraft`: DraftResult
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.CreateConversationDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailbox** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConversationDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

### Return type

[**DraftResult**](DraftResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReplyDraft

> DraftResult CreateReplyDraft(ctx, threadID).IdempotencyKey(idempotencyKey).Execute()

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
	resp, r, err := apiClient.DraftsAPI.CreateReplyDraft(context.Background(), threadID).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.CreateReplyDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReplyDraft`: DraftResult
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.CreateReplyDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReplyDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

### Return type

[**DraftResult**](DraftResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDraft

> Draft GetDraft(ctx, draftID).Execute()

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
	resp, r, err := apiClient.DraftsAPI.GetDraft(context.Background(), draftID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.GetDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDraft`: Draft
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.GetDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Draft**](Draft.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDrafts

> DraftList ListDrafts(ctx).Status(status).Limit(limit).Execute()

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
	resp, r, err := apiClient.DraftsAPI.ListDrafts(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.ListDrafts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDrafts`: DraftList
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.ListDrafts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDraftsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**DraftList**](DraftList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectDraft

> Draft RejectDraft(ctx, draftID).Execute()

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
	resp, r, err := apiClient.DraftsAPI.RejectDraft(context.Background(), draftID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.RejectDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectDraft`: Draft
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.RejectDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Draft**](Draft.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

