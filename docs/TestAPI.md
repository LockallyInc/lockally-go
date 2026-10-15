# \TestAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SimulateInboundEmail**](TestAPI.md#SimulateInboundEmail) | **Post** /v1/test/inbound | Simulate an inbound email (test keys only)



## SimulateInboundEmail

> SimulatedInbound SimulateInboundEmail(ctx).Execute()

Simulate an inbound email (test keys only)



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
	resp, r, err := apiClient.TestAPI.SimulateInboundEmail(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestAPI.SimulateInboundEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SimulateInboundEmail`: SimulatedInbound
	fmt.Fprintf(os.Stdout, "Response from `TestAPI.SimulateInboundEmail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSimulateInboundEmailRequest struct via the builder pattern


### Return type

[**SimulatedInbound**](SimulatedInbound.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

