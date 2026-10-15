# \SendAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelMessage**](SendAPI.md#CancelMessage) | **Delete** /v1/messages/{id} | Cancel a scheduled send
[**GetMessage**](SendAPI.md#GetMessage) | **Get** /v1/messages/{id} | Get message status
[**GetMessageStats**](SendAPI.md#GetMessageStats) | **Get** /v1/messages/stats | Aggregate delivery stats
[**ListMessages**](SendAPI.md#ListMessages) | **Get** /v1/messages | List outbound messages
[**SendEmail**](SendAPI.md#SendEmail) | **Post** /v1/send | Send an email
[**SendEmailBatch**](SendAPI.md#SendEmailBatch) | **Post** /v1/send/batch | Send a batch of emails



## CancelMessage

> CancelMessage(ctx, id).Execute()

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
	r, err := apiClient.SendAPI.CancelMessage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.CancelMessage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCancelMessageRequest struct via the builder pattern


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


## GetMessage

> MessageDetail GetMessage(ctx, id).Execute()

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
	resp, r, err := apiClient.SendAPI.GetMessage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.GetMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessage`: MessageDetail
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.GetMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageRequest struct via the builder pattern


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


## GetMessageStats

> MessageStats GetMessageStats(ctx).From(from).To(to).Domain(domain).Execute()

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
	resp, r, err := apiClient.SendAPI.GetMessageStats(context.Background()).From(from).To(to).Domain(domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.GetMessageStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageStats`: MessageStats
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.GetMessageStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageStatsRequest struct via the builder pattern


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


## ListMessages

> ListMessages200Response ListMessages(ctx).Status(status).Sender(sender).Q(q).Since(since).Cursor(cursor).Limit(limit).Execute()

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
	resp, r, err := apiClient.SendAPI.ListMessages(context.Background()).Status(status).Sender(sender).Q(q).Since(since).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.ListMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMessages`: ListMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.ListMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **sender** | **string** | Exact match against the &#x60;from&#x60; mailbox. | 
 **q** | **string** | Free-text search across subject + sender. | 
 **since** | **time.Time** | Only messages queued at or after this RFC 3339 instant. | 
 **cursor** | **string** | queued_at of the prior page boundary. Pass back the &#x60;next_cursor&#x60; returned by the previous call. | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**ListMessages200Response**](ListMessages200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEmail

> SendEmail200Response SendEmail(ctx).IdempotencyKey(idempotencyKey).SendEmailRequest(sendEmailRequest).Execute()

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
	sendEmailRequest := *openapiclient.NewSendEmailRequest("From_example", []string{"To_example"}) // SendEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendEmail(context.Background()).IdempotencyKey(idempotencyKey).SendEmailRequest(sendEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEmail`: SendEmail200Response
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** |  | 
 **sendEmailRequest** | [**SendEmailRequest**](SendEmailRequest.md) |  | 

### Return type

[**SendEmail200Response**](SendEmail200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEmailBatch

> SendEmailBatch200Response SendEmailBatch(ctx).IdempotencyKey(idempotencyKey).SendEmailBatchRequest(sendEmailBatchRequest).Execute()

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
	sendEmailBatchRequest := *openapiclient.NewSendEmailBatchRequest([]openapiclient.SendMessage{*openapiclient.NewSendMessage("From_example", []string{"To_example"})}) // SendEmailBatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendEmailBatch(context.Background()).IdempotencyKey(idempotencyKey).SendEmailBatchRequest(sendEmailBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendEmailBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEmailBatch`: SendEmailBatch200Response
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendEmailBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** |  | 
 **sendEmailBatchRequest** | [**SendEmailBatchRequest**](SendEmailBatchRequest.md) |  | 

### Return type

[**SendEmailBatch200Response**](SendEmailBatch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

