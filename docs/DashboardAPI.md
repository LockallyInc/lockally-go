# \DashboardAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditSummary**](DashboardAPI.md#GetAuditSummary) | **Get** /v1/audit-summary | Audit summary for the dashboard
[**GetDomainsStatus**](DashboardAPI.md#GetDomainsStatus) | **Get** /v1/domains/status | Domain health status for the dashboard
[**GetIntegrationsSummary**](DashboardAPI.md#GetIntegrationsSummary) | **Get** /v1/integrations-summary | Integrations summary for the dashboard
[**GetSecurity**](DashboardAPI.md#GetSecurity) | **Get** /v1/security | Security overview for the dashboard
[**GetStorage**](DashboardAPI.md#GetStorage) | **Get** /v1/storage | Storage usage for the dashboard
[**GetTenantHealth**](DashboardAPI.md#GetTenantHealth) | **Get** /v1/health | Full tenant health report
[**GetUserInsights**](DashboardAPI.md#GetUserInsights) | **Get** /v1/user-insights | User insights for the dashboard



## GetAuditSummary

> GetAuditSummary200Response GetAuditSummary(ctx).Execute()

Audit summary for the dashboard

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
	resp, r, err := apiClient.DashboardAPI.GetAuditSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetAuditSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditSummary`: GetAuditSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetAuditSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditSummaryRequest struct via the builder pattern


### Return type

[**GetAuditSummary200Response**](GetAuditSummary200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainsStatus

> GetDomainsStatus200Response GetDomainsStatus(ctx).Execute()

Domain health status for the dashboard

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
	resp, r, err := apiClient.DashboardAPI.GetDomainsStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetDomainsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainsStatus`: GetDomainsStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetDomainsStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainsStatusRequest struct via the builder pattern


### Return type

[**GetDomainsStatus200Response**](GetDomainsStatus200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationsSummary

> GetIntegrationsSummary200Response GetIntegrationsSummary(ctx).Execute()

Integrations summary for the dashboard

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
	resp, r, err := apiClient.DashboardAPI.GetIntegrationsSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetIntegrationsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationsSummary`: GetIntegrationsSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetIntegrationsSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsSummaryRequest struct via the builder pattern


### Return type

[**GetIntegrationsSummary200Response**](GetIntegrationsSummary200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurity

> GetSecurity200Response GetSecurity(ctx).Execute()

Security overview for the dashboard

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
	resp, r, err := apiClient.DashboardAPI.GetSecurity(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetSecurity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurity`: GetSecurity200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetSecurity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityRequest struct via the builder pattern


### Return type

[**GetSecurity200Response**](GetSecurity200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorage

> GetStorage200Response GetStorage(ctx).Execute()

Storage usage for the dashboard

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
	resp, r, err := apiClient.DashboardAPI.GetStorage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorage`: GetStorage200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageRequest struct via the builder pattern


### Return type

[**GetStorage200Response**](GetStorage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantHealth

> map[string]interface{} GetTenantHealth(ctx).Execute()

Full tenant health report

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
	resp, r, err := apiClient.DashboardAPI.GetTenantHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetTenantHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantHealth`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetTenantHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantHealthRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserInsights

> GetUserInsights200Response GetUserInsights(ctx).Execute()

User insights for the dashboard

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
	resp, r, err := apiClient.DashboardAPI.GetUserInsights(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetUserInsights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInsights`: GetUserInsights200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetUserInsights`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInsightsRequest struct via the builder pattern


### Return type

[**GetUserInsights200Response**](GetUserInsights200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

