# \DomainsAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DomainsDomainDelete**](DomainsAPI.md#V1DomainsDomainDelete) | **Delete** /v1/domains/{domain} | Delete a domain
[**V1DomainsDomainGet**](DomainsAPI.md#V1DomainsDomainGet) | **Get** /v1/domains/{domain} | Get a domain
[**V1DomainsDomainVerifyPost**](DomainsAPI.md#V1DomainsDomainVerifyPost) | **Post** /v1/domains/{domain}/verify | Force-poll DNS verification
[**V1DomainsGet**](DomainsAPI.md#V1DomainsGet) | **Get** /v1/domains | List domains
[**V1DomainsPost**](DomainsAPI.md#V1DomainsPost) | **Post** /v1/domains | Register a domain



## V1DomainsDomainDelete

> V1DomainsDomainDelete(ctx, domain).Execute()

Delete a domain



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
	domain := "domain_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DomainsAPI.V1DomainsDomainDelete(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.V1DomainsDomainDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DomainsDomainDeleteRequest struct via the builder pattern


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


## V1DomainsDomainGet

> Domain V1DomainsDomainGet(ctx, domain).Execute()

Get a domain

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
	domain := "domain_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.V1DomainsDomainGet(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.V1DomainsDomainGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DomainsDomainGet`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.V1DomainsDomainGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DomainsDomainGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Domain**](Domain.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DomainsDomainVerifyPost

> Domain V1DomainsDomainVerifyPost(ctx, domain).Execute()

Force-poll DNS verification



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
	domain := "domain_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.V1DomainsDomainVerifyPost(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.V1DomainsDomainVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DomainsDomainVerifyPost`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.V1DomainsDomainVerifyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DomainsDomainVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Domain**](Domain.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DomainsGet

> V1DomainsGet200Response V1DomainsGet(ctx).Execute()

List domains



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
	resp, r, err := apiClient.DomainsAPI.V1DomainsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.V1DomainsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DomainsGet`: V1DomainsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.V1DomainsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1DomainsGetRequest struct via the builder pattern


### Return type

[**V1DomainsGet200Response**](V1DomainsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DomainsPost

> Domain V1DomainsPost(ctx).V1DomainsPostRequest(v1DomainsPostRequest).Execute()

Register a domain



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
	v1DomainsPostRequest := *openapiclient.NewV1DomainsPostRequest("acme.com") // V1DomainsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.V1DomainsPost(context.Background()).V1DomainsPostRequest(v1DomainsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.V1DomainsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DomainsPost`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.V1DomainsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DomainsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1DomainsPostRequest** | [**V1DomainsPostRequest**](V1DomainsPostRequest.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

