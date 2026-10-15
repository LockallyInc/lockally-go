# \MailboxesAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSharedMember**](MailboxesAPI.md#AddSharedMember) | **Post** /v1/mailboxes/{email}/members | Add a shared mailbox member
[**CreateMailbox**](MailboxesAPI.md#CreateMailbox) | **Post** /v1/mailboxes | Create a mailbox
[**DeleteMailbox**](MailboxesAPI.md#DeleteMailbox) | **Delete** /v1/mailboxes/{email} | Soft-delete a mailbox
[**DeleteVacationResponder**](MailboxesAPI.md#DeleteVacationResponder) | **Delete** /v1/mailboxes/{email}/vacation | Remove the vacation responder
[**DownloadMailboxExport**](MailboxesAPI.md#DownloadMailboxExport) | **Get** /v1/mailboxes/{email}/export/download | Download a previously-issued mailbox export
[**ExportMailbox**](MailboxesAPI.md#ExportMailbox) | **Post** /v1/mailboxes/{email}/export | Request a mailbox export
[**GetMailbox**](MailboxesAPI.md#GetMailbox) | **Get** /v1/mailboxes/{email} | Get a mailbox
[**GetVacationResponder**](MailboxesAPI.md#GetVacationResponder) | **Get** /v1/mailboxes/{email}/vacation | Get the vacation responder
[**ListMailboxes**](MailboxesAPI.md#ListMailboxes) | **Get** /v1/mailboxes | List mailboxes
[**ListSharedMembers**](MailboxesAPI.md#ListSharedMembers) | **Get** /v1/mailboxes/{email}/members | List shared mailbox members
[**ListVacationResponders**](MailboxesAPI.md#ListVacationResponders) | **Get** /v1/vacation | List all vacation responders
[**RemoveSharedMember**](MailboxesAPI.md#RemoveSharedMember) | **Delete** /v1/mailboxes/{email}/members/{memberEmail} | Remove a shared mailbox member
[**SetVacationResponder**](MailboxesAPI.md#SetVacationResponder) | **Put** /v1/mailboxes/{email}/vacation | Set the vacation responder
[**UpdateMailbox**](MailboxesAPI.md#UpdateMailbox) | **Patch** /v1/mailboxes/{email} | Update a mailbox



## AddSharedMember

> SharedMember AddSharedMember(ctx, email).AddSharedMemberRequest(addSharedMemberRequest).Execute()

Add a shared mailbox member

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
	addSharedMemberRequest := *openapiclient.NewAddSharedMemberRequest("MemberEmail_example") // AddSharedMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailboxesAPI.AddSharedMember(context.Background(), email).AddSharedMemberRequest(addSharedMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.AddSharedMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSharedMember`: SharedMember
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.AddSharedMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSharedMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addSharedMemberRequest** | [**AddSharedMemberRequest**](AddSharedMemberRequest.md) |  | 

### Return type

[**SharedMember**](SharedMember.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMailbox

> Mailbox CreateMailbox(ctx).CreateMailboxRequest(createMailboxRequest).Execute()

Create a mailbox



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
	createMailboxRequest := *openapiclient.NewCreateMailboxRequest("alice@acme.com") // CreateMailboxRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailboxesAPI.CreateMailbox(context.Background()).CreateMailboxRequest(createMailboxRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.CreateMailbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMailbox`: Mailbox
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.CreateMailbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMailboxRequest** | [**CreateMailboxRequest**](CreateMailboxRequest.md) |  | 

### Return type

[**Mailbox**](Mailbox.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMailbox

> DeleteMailbox(ctx, email).Execute()

Soft-delete a mailbox



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
	r, err := apiClient.MailboxesAPI.DeleteMailbox(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.DeleteMailbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMailboxRequest struct via the builder pattern


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


## DeleteVacationResponder

> DeleteVacationResponder(ctx, email).Execute()

Remove the vacation responder



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
	r, err := apiClient.MailboxesAPI.DeleteVacationResponder(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.DeleteVacationResponder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVacationResponderRequest struct via the builder pattern


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


## DownloadMailboxExport

> *os.File DownloadMailboxExport(ctx, email).Token(token).Execute()

Download a previously-issued mailbox export



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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailboxesAPI.DownloadMailboxExport(context.Background(), email).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.DownloadMailboxExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadMailboxExport`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.DownloadMailboxExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadMailboxExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/mbox, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportMailbox

> ExportMailbox201Response ExportMailbox(ctx, email).Execute()

Request a mailbox export



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
	resp, r, err := apiClient.MailboxesAPI.ExportMailbox(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.ExportMailbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportMailbox`: ExportMailbox201Response
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.ExportMailbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExportMailbox201Response**](ExportMailbox201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMailbox

> Mailbox GetMailbox(ctx, email).Execute()

Get a mailbox

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
	resp, r, err := apiClient.MailboxesAPI.GetMailbox(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.GetMailbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMailbox`: Mailbox
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.GetMailbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Mailbox**](Mailbox.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacationResponder

> VacationResponder GetVacationResponder(ctx, email).Execute()

Get the vacation responder



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
	resp, r, err := apiClient.MailboxesAPI.GetVacationResponder(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.GetVacationResponder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacationResponder`: VacationResponder
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.GetVacationResponder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVacationResponderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VacationResponder**](VacationResponder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMailboxes

> ListMailboxes200Response ListMailboxes(ctx).Limit(limit).Execute()

List mailboxes



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
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailboxesAPI.ListMailboxes(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.ListMailboxes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMailboxes`: ListMailboxes200Response
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.ListMailboxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMailboxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]

### Return type

[**ListMailboxes200Response**](ListMailboxes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharedMembers

> ListSharedMembers200Response ListSharedMembers(ctx, email).Execute()

List shared mailbox members

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
	resp, r, err := apiClient.MailboxesAPI.ListSharedMembers(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.ListSharedMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSharedMembers`: ListSharedMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.ListSharedMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSharedMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSharedMembers200Response**](ListSharedMembers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVacationResponders

> ListVacationResponders200Response ListVacationResponders(ctx).Execute()

List all vacation responders



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
	resp, r, err := apiClient.MailboxesAPI.ListVacationResponders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.ListVacationResponders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVacationResponders`: ListVacationResponders200Response
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.ListVacationResponders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVacationRespondersRequest struct via the builder pattern


### Return type

[**ListVacationResponders200Response**](ListVacationResponders200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSharedMember

> RemoveSharedMember(ctx, email, memberEmail).Execute()

Remove a shared mailbox member

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
	memberEmail := "memberEmail_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MailboxesAPI.RemoveSharedMember(context.Background(), email, memberEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.RemoveSharedMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 
**memberEmail** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSharedMemberRequest struct via the builder pattern


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


## SetVacationResponder

> VacationResponder SetVacationResponder(ctx, email).SetVacationResponderRequest(setVacationResponderRequest).Execute()

Set the vacation responder



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
	setVacationResponderRequest := *openapiclient.NewSetVacationResponderRequest(*openapiclient.NewVacationParams("Out of office until June 5", "Hi! I'm away until June 5. For urgent matters please contact ...")) // SetVacationResponderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailboxesAPI.SetVacationResponder(context.Background(), email).SetVacationResponderRequest(setVacationResponderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.SetVacationResponder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVacationResponder`: VacationResponder
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.SetVacationResponder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVacationResponderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setVacationResponderRequest** | [**SetVacationResponderRequest**](SetVacationResponderRequest.md) |  | 

### Return type

[**VacationResponder**](VacationResponder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMailbox

> Mailbox UpdateMailbox(ctx, email).UpdateMailboxRequest(updateMailboxRequest).Execute()

Update a mailbox



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
	updateMailboxRequest := *openapiclient.NewUpdateMailboxRequest() // UpdateMailboxRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailboxesAPI.UpdateMailbox(context.Background(), email).UpdateMailboxRequest(updateMailboxRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.UpdateMailbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMailbox`: Mailbox
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.UpdateMailbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMailboxRequest** | [**UpdateMailboxRequest**](UpdateMailboxRequest.md) |  | 

### Return type

[**Mailbox**](Mailbox.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

