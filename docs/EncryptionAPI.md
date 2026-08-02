# \EncryptionAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchLookupPublicKeys**](EncryptionAPI.md#BatchLookupPublicKeys) | **Get** /v1/encryption/keys/lookup | Batch-lookup public keys by email
[**CreateEncryptionKey**](EncryptionAPI.md#CreateEncryptionKey) | **Post** /v1/encryption/keys | Upload an encryption key pair
[**CreateEncryptionRecovery**](EncryptionAPI.md#CreateEncryptionRecovery) | **Post** /v1/encryption/recovery | Store an encryption recovery blob
[**GetEncryptionKey**](EncryptionAPI.md#GetEncryptionKey) | **Get** /v1/encryption/keys/{email} | Get encryption key for a mailbox
[**RotateEncryptionKey**](EncryptionAPI.md#RotateEncryptionKey) | **Post** /v1/encryption/keys/rotate | Rotate an encryption key



## BatchLookupPublicKeys

> BatchLookupPublicKeys200Response BatchLookupPublicKeys(ctx).Emails(emails).Execute()

Batch-lookup public keys by email

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
	emails := "emails_example" // string | Comma-separated list of email addresses

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EncryptionAPI.BatchLookupPublicKeys(context.Background()).Emails(emails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAPI.BatchLookupPublicKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchLookupPublicKeys`: BatchLookupPublicKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAPI.BatchLookupPublicKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchLookupPublicKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emails** | **string** | Comma-separated list of email addresses | 

### Return type

[**BatchLookupPublicKeys200Response**](BatchLookupPublicKeys200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEncryptionKey

> CreateEncryptionKey201Response CreateEncryptionKey(ctx).CreateEncryptionKeyRequest(createEncryptionKeyRequest).Execute()

Upload an encryption key pair

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
	createEncryptionKeyRequest := *openapiclient.NewCreateEncryptionKeyRequest("MailboxEmail_example", "PublicKey_example", "EncryptedPrivateKey_example") // CreateEncryptionKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EncryptionAPI.CreateEncryptionKey(context.Background()).CreateEncryptionKeyRequest(createEncryptionKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAPI.CreateEncryptionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEncryptionKey`: CreateEncryptionKey201Response
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAPI.CreateEncryptionKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEncryptionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEncryptionKeyRequest** | [**CreateEncryptionKeyRequest**](CreateEncryptionKeyRequest.md) |  | 

### Return type

[**CreateEncryptionKey201Response**](CreateEncryptionKey201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEncryptionRecovery

> CreateEncryptionRecovery(ctx).CreateEncryptionRecoveryRequest(createEncryptionRecoveryRequest).Execute()

Store an encryption recovery blob

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
	createEncryptionRecoveryRequest := *openapiclient.NewCreateEncryptionRecoveryRequest("MailboxEmail_example", "RecoveryBlob_example") // CreateEncryptionRecoveryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EncryptionAPI.CreateEncryptionRecovery(context.Background()).CreateEncryptionRecoveryRequest(createEncryptionRecoveryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAPI.CreateEncryptionRecovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEncryptionRecoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEncryptionRecoveryRequest** | [**CreateEncryptionRecoveryRequest**](CreateEncryptionRecoveryRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncryptionKey

> GetEncryptionKey200Response GetEncryptionKey(ctx, email).Execute()

Get encryption key for a mailbox

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
	resp, r, err := apiClient.EncryptionAPI.GetEncryptionKey(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAPI.GetEncryptionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncryptionKey`: GetEncryptionKey200Response
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAPI.GetEncryptionKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncryptionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEncryptionKey200Response**](GetEncryptionKey200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateEncryptionKey

> RotateEncryptionKey(ctx).RotateEncryptionKeyRequest(rotateEncryptionKeyRequest).Execute()

Rotate an encryption key

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
	rotateEncryptionKeyRequest := *openapiclient.NewRotateEncryptionKeyRequest("MailboxEmail_example", "EncryptedPrivateKey_example") // RotateEncryptionKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EncryptionAPI.RotateEncryptionKey(context.Background()).RotateEncryptionKeyRequest(rotateEncryptionKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAPI.RotateEncryptionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRotateEncryptionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rotateEncryptionKeyRequest** | [**RotateEncryptionKeyRequest**](RotateEncryptionKeyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

