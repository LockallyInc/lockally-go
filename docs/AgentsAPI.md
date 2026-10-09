# \AgentsAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ApiKeysKeyIDMailboxesGet**](AgentsAPI.md#V1ApiKeysKeyIDMailboxesGet) | **Get** /v1/api-keys/{keyID}/mailboxes | List a key&#39;s mailbox grants
[**V1ApiKeysKeyIDMailboxesMailboxIDDelete**](AgentsAPI.md#V1ApiKeysKeyIDMailboxesMailboxIDDelete) | **Delete** /v1/api-keys/{keyID}/mailboxes/{mailboxID} | Revoke a mailbox grant
[**V1ApiKeysKeyIDMailboxesPost**](AgentsAPI.md#V1ApiKeysKeyIDMailboxesPost) | **Post** /v1/api-keys/{keyID}/mailboxes | Grant a mailbox to a key
[**V1AuthWhoamiGet**](AgentsAPI.md#V1AuthWhoamiGet) | **Get** /v1/auth/whoami | Introspect the calling credentials
[**V1ContactsLookupGet**](AgentsAPI.md#V1ContactsLookupGet) | **Get** /v1/contacts/lookup | Who is this sender?
[**V1InboxesGet**](AgentsAPI.md#V1InboxesGet) | **Get** /v1/inboxes | List granted inboxes
[**V1InboxesMailboxMessagesPost**](AgentsAPI.md#V1InboxesMailboxMessagesPost) | **Post** /v1/inboxes/{mailbox}/messages | Start a new conversation (agent stream)
[**V1InboxesMailboxThreadsGet**](AgentsAPI.md#V1InboxesMailboxThreadsGet) | **Get** /v1/inboxes/{mailbox}/threads | List conversation threads
[**V1ThreadsThreadIDGet**](AgentsAPI.md#V1ThreadsThreadIDGet) | **Get** /v1/threads/{threadID} | Get a whole conversation
[**V1ThreadsThreadIDMessagesMessageIDAttachmentsIdxGet**](AgentsAPI.md#V1ThreadsThreadIDMessagesMessageIDAttachmentsIdxGet) | **Get** /v1/threads/{threadID}/messages/{messageID}/attachments/{idx} | Download an attachment
[**V1ThreadsThreadIDMessagesMessageIDGet**](AgentsAPI.md#V1ThreadsThreadIDMessagesMessageIDGet) | **Get** /v1/threads/{threadID}/messages/{messageID} | Get one message with body
[**V1ThreadsThreadIDMessagesMessageIDReadPost**](AgentsAPI.md#V1ThreadsThreadIDMessagesMessageIDReadPost) | **Post** /v1/threads/{threadID}/messages/{messageID}/read | Mark read/unread
[**V1ThreadsThreadIDReplyPost**](AgentsAPI.md#V1ThreadsThreadIDReplyPost) | **Post** /v1/threads/{threadID}/reply | Reply in-thread (agent stream)



## V1ApiKeysKeyIDMailboxesGet

> map[string]interface{} V1ApiKeysKeyIDMailboxesGet(ctx, keyID).Execute()

List a key's mailbox grants

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
	keyID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.V1ApiKeysKeyIDMailboxesGet(context.Background(), keyID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.V1ApiKeysKeyIDMailboxesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApiKeysKeyIDMailboxesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.V1ApiKeysKeyIDMailboxesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApiKeysKeyIDMailboxesGetRequest struct via the builder pattern


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


## V1ApiKeysKeyIDMailboxesMailboxIDDelete

> V1ApiKeysKeyIDMailboxesMailboxIDDelete(ctx, keyID, mailboxID).Execute()

Revoke a mailbox grant

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
	keyID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	mailboxID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentsAPI.V1ApiKeysKeyIDMailboxesMailboxIDDelete(context.Background(), keyID, mailboxID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.V1ApiKeysKeyIDMailboxesMailboxIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** |  | 
**mailboxID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApiKeysKeyIDMailboxesMailboxIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApiKeysKeyIDMailboxesPost

> map[string]interface{} V1ApiKeysKeyIDMailboxesPost(ctx, keyID).Execute()

Grant a mailbox to a key



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
	keyID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.V1ApiKeysKeyIDMailboxesPost(context.Background(), keyID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.V1ApiKeysKeyIDMailboxesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApiKeysKeyIDMailboxesPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.V1ApiKeysKeyIDMailboxesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApiKeysKeyIDMailboxesPostRequest struct via the builder pattern


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


## V1AuthWhoamiGet

> map[string]interface{} V1AuthWhoamiGet(ctx).Execute()

Introspect the calling credentials



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
	resp, r, err := apiClient.AgentsAPI.V1AuthWhoamiGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.V1AuthWhoamiGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AuthWhoamiGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.V1AuthWhoamiGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AuthWhoamiGetRequest struct via the builder pattern


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


## V1ContactsLookupGet

> map[string]interface{} V1ContactsLookupGet(ctx).Email(email).Execute()

Who is this sender?



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
	email := "email_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.V1ContactsLookupGet(context.Background()).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.V1ContactsLookupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ContactsLookupGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.V1ContactsLookupGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ContactsLookupGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 

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


## V1InboxesGet

> map[string]interface{} V1InboxesGet(ctx).Execute()

List granted inboxes



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
	resp, r, err := apiClient.AgentsAPI.V1InboxesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.V1InboxesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InboxesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.V1InboxesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1InboxesGetRequest struct via the builder pattern


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


## V1InboxesMailboxMessagesPost

> map[string]interface{} V1InboxesMailboxMessagesPost(ctx, mailbox).IdempotencyKey(idempotencyKey).V1InboxesMailboxMessagesPostRequest(v1InboxesMailboxMessagesPostRequest).Execute()

Start a new conversation (agent stream)



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
	v1InboxesMailboxMessagesPostRequest := *openapiclient.NewV1InboxesMailboxMessagesPostRequest([]string{"To_example"}, "Text_example") // V1InboxesMailboxMessagesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.V1InboxesMailboxMessagesPost(context.Background(), mailbox).IdempotencyKey(idempotencyKey).V1InboxesMailboxMessagesPostRequest(v1InboxesMailboxMessagesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.V1InboxesMailboxMessagesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InboxesMailboxMessagesPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.V1InboxesMailboxMessagesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailbox** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InboxesMailboxMessagesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 
 **v1InboxesMailboxMessagesPostRequest** | [**V1InboxesMailboxMessagesPostRequest**](V1InboxesMailboxMessagesPostRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InboxesMailboxThreadsGet

> map[string]interface{} V1InboxesMailboxThreadsGet(ctx, mailbox).Since(since).Before(before).Limit(limit).Execute()

List conversation threads



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/lockallyinc/lockally-go"
)

func main() {
	mailbox := "mailbox_example" // string | mailbox email or id
	since := time.Now() // time.Time |  (optional)
	before := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.V1InboxesMailboxThreadsGet(context.Background(), mailbox).Since(since).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.V1InboxesMailboxThreadsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InboxesMailboxThreadsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.V1InboxesMailboxThreadsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailbox** | **string** | mailbox email or id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InboxesMailboxThreadsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **time.Time** |  | 
 **before** | **time.Time** |  | 
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


## V1ThreadsThreadIDGet

> map[string]interface{} V1ThreadsThreadIDGet(ctx, threadID).Execute()

Get a whole conversation



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.V1ThreadsThreadIDGet(context.Background(), threadID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.V1ThreadsThreadIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ThreadsThreadIDGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.V1ThreadsThreadIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ThreadsThreadIDGetRequest struct via the builder pattern


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


## V1ThreadsThreadIDMessagesMessageIDAttachmentsIdxGet

> V1ThreadsThreadIDMessagesMessageIDAttachmentsIdxGet(ctx, threadID, messageID, idx).Execute()

Download an attachment

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
	messageID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idx := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentsAPI.V1ThreadsThreadIDMessagesMessageIDAttachmentsIdxGet(context.Background(), threadID, messageID, idx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.V1ThreadsThreadIDMessagesMessageIDAttachmentsIdxGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadID** | **string** |  | 
**messageID** | **string** |  | 
**idx** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ThreadsThreadIDMessagesMessageIDAttachmentsIdxGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ThreadsThreadIDMessagesMessageIDGet

> map[string]interface{} V1ThreadsThreadIDMessagesMessageIDGet(ctx, threadID, messageID).Execute()

Get one message with body



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
	messageID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.V1ThreadsThreadIDMessagesMessageIDGet(context.Background(), threadID, messageID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.V1ThreadsThreadIDMessagesMessageIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ThreadsThreadIDMessagesMessageIDGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.V1ThreadsThreadIDMessagesMessageIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadID** | **string** |  | 
**messageID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ThreadsThreadIDMessagesMessageIDGetRequest struct via the builder pattern


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


## V1ThreadsThreadIDMessagesMessageIDReadPost

> map[string]interface{} V1ThreadsThreadIDMessagesMessageIDReadPost(ctx, threadID, messageID).Execute()

Mark read/unread



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
	messageID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.V1ThreadsThreadIDMessagesMessageIDReadPost(context.Background(), threadID, messageID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.V1ThreadsThreadIDMessagesMessageIDReadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ThreadsThreadIDMessagesMessageIDReadPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.V1ThreadsThreadIDMessagesMessageIDReadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadID** | **string** |  | 
**messageID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ThreadsThreadIDMessagesMessageIDReadPostRequest struct via the builder pattern


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


## V1ThreadsThreadIDReplyPost

> map[string]interface{} V1ThreadsThreadIDReplyPost(ctx, threadID).IdempotencyKey(idempotencyKey).Execute()

Reply in-thread (agent stream)



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
	resp, r, err := apiClient.AgentsAPI.V1ThreadsThreadIDReplyPost(context.Background(), threadID).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.V1ThreadsThreadIDReplyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ThreadsThreadIDReplyPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.V1ThreadsThreadIDReplyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ThreadsThreadIDReplyPostRequest struct via the builder pattern


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

