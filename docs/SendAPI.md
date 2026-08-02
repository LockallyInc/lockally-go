# \SendAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MessagesGet**](SendAPI.md#V1MessagesGet) | **Get** /v1/messages | List outbound messages
[**V1MessagesIdDelete**](SendAPI.md#V1MessagesIdDelete) | **Delete** /v1/messages/{id} | Cancel a scheduled send
[**V1MessagesIdGet**](SendAPI.md#V1MessagesIdGet) | **Get** /v1/messages/{id} | Get message status
[**V1MessagesStatsGet**](SendAPI.md#V1MessagesStatsGet) | **Get** /v1/messages/stats | Aggregate delivery stats
[**V1SendBatchPost**](SendAPI.md#V1SendBatchPost) | **Post** /v1/send/batch | Send a batch of emails
[**V1SendPost**](SendAPI.md#V1SendPost) | **Post** /v1/send | Send an email



## V1MessagesGet

> V1MessagesGet200Response V1MessagesGet(ctx).Status(status).Sender(sender).Q(q).Since(since).Cursor(cursor).Limit(limit).Execute()

List outbound messages



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
	status := "status_example" // string |  (optional)
	sender := "sender_example" // string | Exact match against the `from` mailbox. (optional)
	q := "q_example" // string | Free-text search across subject + sender. (optional)
	since := time.Now() // time.Time | Only messages queued at or after this RFC 3339 instant. (optional)
	cursor := "cursor_example" // string | queued_at of the prior page boundary. Pass back the `next_cursor` returned by the previous call. (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.V1MessagesGet(context.Background()).Status(status).Sender(sender).Q(q).Since(since).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.V1MessagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MessagesGet`: V1MessagesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.V1MessagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MessagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **sender** | **string** | Exact match against the &#x60;from&#x60; mailbox. | 
 **q** | **string** | Free-text search across subject + sender. | 
 **since** | **time.Time** | Only messages queued at or after this RFC 3339 instant. | 
 **cursor** | **string** | queued_at of the prior page boundary. Pass back the &#x60;next_cursor&#x60; returned by the previous call. | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**V1MessagesGet200Response**](V1MessagesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MessagesIdDelete

> V1MessagesIdDelete(ctx, id).Execute()

Cancel a scheduled send



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SendAPI.V1MessagesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.V1MessagesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MessagesIdDeleteRequest struct via the builder pattern


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


## V1MessagesIdGet

> MessageDetail V1MessagesIdGet(ctx, id).Execute()

Get message status

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.V1MessagesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.V1MessagesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MessagesIdGet`: MessageDetail
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.V1MessagesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MessagesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageDetail**](MessageDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MessagesStatsGet

> MessageStats V1MessagesStatsGet(ctx).From(from).To(to).Domain(domain).Execute()

Aggregate delivery stats



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
	from := time.Now() // time.Time | Window start (default 7 days ago). (optional)
	to := time.Now() // time.Time | Window end (default now). (optional)
	domain := "domain_example" // string | Filter by sender domain. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.V1MessagesStatsGet(context.Background()).From(from).To(to).Domain(domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.V1MessagesStatsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MessagesStatsGet`: MessageStats
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.V1MessagesStatsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MessagesStatsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** | Window start (default 7 days ago). | 
 **to** | **time.Time** | Window end (default now). | 
 **domain** | **string** | Filter by sender domain. | 

### Return type

[**MessageStats**](MessageStats.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SendBatchPost

> V1SendBatchPost200Response V1SendBatchPost(ctx).IdempotencyKey(idempotencyKey).V1SendBatchPostRequest(v1SendBatchPostRequest).Execute()

Send a batch of emails



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
	idempotencyKey := "idempotencyKey_example" // string | 
	v1SendBatchPostRequest := *openapiclient.NewV1SendBatchPostRequest([]openapiclient.SendMessage{*openapiclient.NewSendMessage("From_example", []string{"To_example"})}) // V1SendBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.V1SendBatchPost(context.Background()).IdempotencyKey(idempotencyKey).V1SendBatchPostRequest(v1SendBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.V1SendBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SendBatchPost`: V1SendBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.V1SendBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SendBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** |  | 
 **v1SendBatchPostRequest** | [**V1SendBatchPostRequest**](V1SendBatchPostRequest.md) |  | 

### Return type

[**V1SendBatchPost200Response**](V1SendBatchPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SendPost

> V1SendPost202Response V1SendPost(ctx).IdempotencyKey(idempotencyKey).V1SendPostRequest(v1SendPostRequest).Execute()

Send an email



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
	idempotencyKey := "idempotencyKey_example" // string | 
	v1SendPostRequest := *openapiclient.NewV1SendPostRequest("From_example", []string{"To_example"}) // V1SendPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.V1SendPost(context.Background()).IdempotencyKey(idempotencyKey).V1SendPostRequest(v1SendPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.V1SendPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SendPost`: V1SendPost202Response
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.V1SendPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** |  | 
 **v1SendPostRequest** | [**V1SendPostRequest**](V1SendPostRequest.md) |  | 

### Return type

[**V1SendPost202Response**](V1SendPost202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

