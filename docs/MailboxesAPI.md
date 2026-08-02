# \MailboxesAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSharedMember**](MailboxesAPI.md#AddSharedMember) | **Post** /v1/mailboxes/{email}/members | Add a shared mailbox member
[**ListSharedMembers**](MailboxesAPI.md#ListSharedMembers) | **Get** /v1/mailboxes/{email}/members | List shared mailbox members
[**RemoveSharedMember**](MailboxesAPI.md#RemoveSharedMember) | **Delete** /v1/mailboxes/{email}/members/{memberEmail} | Remove a shared mailbox member
[**V1MailboxesEmailDelete**](MailboxesAPI.md#V1MailboxesEmailDelete) | **Delete** /v1/mailboxes/{email} | Soft-delete a mailbox
[**V1MailboxesEmailExportDownloadGet**](MailboxesAPI.md#V1MailboxesEmailExportDownloadGet) | **Get** /v1/mailboxes/{email}/export/download | Download a previously-issued mailbox export
[**V1MailboxesEmailExportPost**](MailboxesAPI.md#V1MailboxesEmailExportPost) | **Post** /v1/mailboxes/{email}/export | Request a mailbox export
[**V1MailboxesEmailGet**](MailboxesAPI.md#V1MailboxesEmailGet) | **Get** /v1/mailboxes/{email} | Get a mailbox
[**V1MailboxesEmailPatch**](MailboxesAPI.md#V1MailboxesEmailPatch) | **Patch** /v1/mailboxes/{email} | Update a mailbox
[**V1MailboxesEmailVacationDelete**](MailboxesAPI.md#V1MailboxesEmailVacationDelete) | **Delete** /v1/mailboxes/{email}/vacation | Remove the vacation responder
[**V1MailboxesEmailVacationGet**](MailboxesAPI.md#V1MailboxesEmailVacationGet) | **Get** /v1/mailboxes/{email}/vacation | Get the vacation responder
[**V1MailboxesEmailVacationPut**](MailboxesAPI.md#V1MailboxesEmailVacationPut) | **Put** /v1/mailboxes/{email}/vacation | Set the vacation responder
[**V1MailboxesGet**](MailboxesAPI.md#V1MailboxesGet) | **Get** /v1/mailboxes | List mailboxes
[**V1MailboxesPost**](MailboxesAPI.md#V1MailboxesPost) | **Post** /v1/mailboxes | Create a mailbox
[**V1VacationGet**](MailboxesAPI.md#V1VacationGet) | **Get** /v1/vacation | List all vacation responders



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


## V1MailboxesEmailDelete

> V1MailboxesEmailDelete(ctx, email).Execute()

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
	r, err := apiClient.MailboxesAPI.V1MailboxesEmailDelete(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.V1MailboxesEmailDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1MailboxesEmailDeleteRequest struct via the builder pattern


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


## V1MailboxesEmailExportDownloadGet

> *os.File V1MailboxesEmailExportDownloadGet(ctx, email).Token(token).Execute()

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
	resp, r, err := apiClient.MailboxesAPI.V1MailboxesEmailExportDownloadGet(context.Background(), email).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.V1MailboxesEmailExportDownloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MailboxesEmailExportDownloadGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.V1MailboxesEmailExportDownloadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MailboxesEmailExportDownloadGetRequest struct via the builder pattern


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


## V1MailboxesEmailExportPost

> V1MailboxesEmailExportPost201Response V1MailboxesEmailExportPost(ctx, email).Execute()

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
	resp, r, err := apiClient.MailboxesAPI.V1MailboxesEmailExportPost(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.V1MailboxesEmailExportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MailboxesEmailExportPost`: V1MailboxesEmailExportPost201Response
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.V1MailboxesEmailExportPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MailboxesEmailExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1MailboxesEmailExportPost201Response**](V1MailboxesEmailExportPost201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MailboxesEmailGet

> Mailbox V1MailboxesEmailGet(ctx, email).Execute()

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
	resp, r, err := apiClient.MailboxesAPI.V1MailboxesEmailGet(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.V1MailboxesEmailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MailboxesEmailGet`: Mailbox
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.V1MailboxesEmailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MailboxesEmailGetRequest struct via the builder pattern


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


## V1MailboxesEmailPatch

> Mailbox V1MailboxesEmailPatch(ctx, email).V1MailboxesEmailPatchRequest(v1MailboxesEmailPatchRequest).Execute()

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
	v1MailboxesEmailPatchRequest := *openapiclient.NewV1MailboxesEmailPatchRequest() // V1MailboxesEmailPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailboxesAPI.V1MailboxesEmailPatch(context.Background(), email).V1MailboxesEmailPatchRequest(v1MailboxesEmailPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.V1MailboxesEmailPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MailboxesEmailPatch`: Mailbox
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.V1MailboxesEmailPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MailboxesEmailPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1MailboxesEmailPatchRequest** | [**V1MailboxesEmailPatchRequest**](V1MailboxesEmailPatchRequest.md) |  | 

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


## V1MailboxesEmailVacationDelete

> V1MailboxesEmailVacationDelete(ctx, email).Execute()

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
	r, err := apiClient.MailboxesAPI.V1MailboxesEmailVacationDelete(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.V1MailboxesEmailVacationDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1MailboxesEmailVacationDeleteRequest struct via the builder pattern


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


## V1MailboxesEmailVacationGet

> VacationResponder V1MailboxesEmailVacationGet(ctx, email).Execute()

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
	resp, r, err := apiClient.MailboxesAPI.V1MailboxesEmailVacationGet(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.V1MailboxesEmailVacationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MailboxesEmailVacationGet`: VacationResponder
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.V1MailboxesEmailVacationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MailboxesEmailVacationGetRequest struct via the builder pattern


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


## V1MailboxesEmailVacationPut

> VacationResponder V1MailboxesEmailVacationPut(ctx, email).V1MailboxesEmailVacationPutRequest(v1MailboxesEmailVacationPutRequest).Execute()

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
	v1MailboxesEmailVacationPutRequest := *openapiclient.NewV1MailboxesEmailVacationPutRequest(*openapiclient.NewVacationParams("Out of office until June 5", "Hi! I'm away until June 5. For urgent matters please contact ...")) // V1MailboxesEmailVacationPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailboxesAPI.V1MailboxesEmailVacationPut(context.Background(), email).V1MailboxesEmailVacationPutRequest(v1MailboxesEmailVacationPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.V1MailboxesEmailVacationPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MailboxesEmailVacationPut`: VacationResponder
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.V1MailboxesEmailVacationPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MailboxesEmailVacationPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1MailboxesEmailVacationPutRequest** | [**V1MailboxesEmailVacationPutRequest**](V1MailboxesEmailVacationPutRequest.md) |  | 

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


## V1MailboxesGet

> V1MailboxesGet200Response V1MailboxesGet(ctx).Limit(limit).Execute()

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
	resp, r, err := apiClient.MailboxesAPI.V1MailboxesGet(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.V1MailboxesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MailboxesGet`: V1MailboxesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.V1MailboxesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MailboxesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]

### Return type

[**V1MailboxesGet200Response**](V1MailboxesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MailboxesPost

> Mailbox V1MailboxesPost(ctx).V1MailboxesPostRequest(v1MailboxesPostRequest).Execute()

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
	v1MailboxesPostRequest := *openapiclient.NewV1MailboxesPostRequest("alice@acme.com") // V1MailboxesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailboxesAPI.V1MailboxesPost(context.Background()).V1MailboxesPostRequest(v1MailboxesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.V1MailboxesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MailboxesPost`: Mailbox
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.V1MailboxesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MailboxesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1MailboxesPostRequest** | [**V1MailboxesPostRequest**](V1MailboxesPostRequest.md) |  | 

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


## V1VacationGet

> V1VacationGet200Response V1VacationGet(ctx).Execute()

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
	resp, r, err := apiClient.MailboxesAPI.V1VacationGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxesAPI.V1VacationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1VacationGet`: V1VacationGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MailboxesAPI.V1VacationGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1VacationGetRequest struct via the builder pattern


### Return type

[**V1VacationGet200Response**](V1VacationGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

