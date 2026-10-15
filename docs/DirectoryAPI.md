# \DirectoryAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDirectoryActivity**](DirectoryAPI.md#GetDirectoryActivity) | **Get** /v1/directory-activity | Get recent directory activity
[**GetDirectoryPermissions**](DirectoryAPI.md#GetDirectoryPermissions) | **Get** /v1/directory-permissions | Get directory permission settings
[**GetDirectoryStats**](DirectoryAPI.md#GetDirectoryStats) | **Get** /v1/directory-stats | Get directory statistics
[**GetGALSettings**](DirectoryAPI.md#GetGALSettings) | **Get** /v1/gal-settings | Get Global Address List settings
[**RebuildGALIndex**](DirectoryAPI.md#RebuildGALIndex) | **Post** /v1/gal-settings/rebuild-index | Rebuild the GAL search index
[**SyncGAL**](DirectoryAPI.md#SyncGAL) | **Post** /v1/gal-settings/sync | Sync GAL with external directory sources
[**UpdateDirectoryPermissions**](DirectoryAPI.md#UpdateDirectoryPermissions) | **Patch** /v1/directory-permissions | Update directory permission settings
[**UpdateGALSettings**](DirectoryAPI.md#UpdateGALSettings) | **Patch** /v1/gal-settings | Update GAL settings



## GetDirectoryActivity

> GetDirectoryActivity200Response GetDirectoryActivity(ctx).Execute()

Get recent directory activity

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
	resp, r, err := apiClient.DirectoryAPI.GetDirectoryActivity(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAPI.GetDirectoryActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDirectoryActivity`: GetDirectoryActivity200Response
	fmt.Fprintf(os.Stdout, "Response from `DirectoryAPI.GetDirectoryActivity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectoryActivityRequest struct via the builder pattern


### Return type

[**GetDirectoryActivity200Response**](GetDirectoryActivity200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectoryPermissions

> DirectoryPermissions GetDirectoryPermissions(ctx).Execute()

Get directory permission settings

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
	resp, r, err := apiClient.DirectoryAPI.GetDirectoryPermissions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAPI.GetDirectoryPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDirectoryPermissions`: DirectoryPermissions
	fmt.Fprintf(os.Stdout, "Response from `DirectoryAPI.GetDirectoryPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectoryPermissionsRequest struct via the builder pattern


### Return type

[**DirectoryPermissions**](DirectoryPermissions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectoryStats

> GetDirectoryStats200Response GetDirectoryStats(ctx).Execute()

Get directory statistics

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
	resp, r, err := apiClient.DirectoryAPI.GetDirectoryStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAPI.GetDirectoryStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDirectoryStats`: GetDirectoryStats200Response
	fmt.Fprintf(os.Stdout, "Response from `DirectoryAPI.GetDirectoryStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectoryStatsRequest struct via the builder pattern


### Return type

[**GetDirectoryStats200Response**](GetDirectoryStats200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGALSettings

> GALSettings GetGALSettings(ctx).Execute()

Get Global Address List settings

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
	resp, r, err := apiClient.DirectoryAPI.GetGALSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAPI.GetGALSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGALSettings`: GALSettings
	fmt.Fprintf(os.Stdout, "Response from `DirectoryAPI.GetGALSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGALSettingsRequest struct via the builder pattern


### Return type

[**GALSettings**](GALSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebuildGALIndex

> GALSettings RebuildGALIndex(ctx).Execute()

Rebuild the GAL search index

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
	resp, r, err := apiClient.DirectoryAPI.RebuildGALIndex(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAPI.RebuildGALIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebuildGALIndex`: GALSettings
	fmt.Fprintf(os.Stdout, "Response from `DirectoryAPI.RebuildGALIndex`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRebuildGALIndexRequest struct via the builder pattern


### Return type

[**GALSettings**](GALSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncGAL

> GALSettings SyncGAL(ctx).Execute()

Sync GAL with external directory sources

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
	resp, r, err := apiClient.DirectoryAPI.SyncGAL(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAPI.SyncGAL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncGAL`: GALSettings
	fmt.Fprintf(os.Stdout, "Response from `DirectoryAPI.SyncGAL`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncGALRequest struct via the builder pattern


### Return type

[**GALSettings**](GALSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDirectoryPermissions

> DirectoryPermissions UpdateDirectoryPermissions(ctx).UpdateDirectoryPermissionsRequest(updateDirectoryPermissionsRequest).Execute()

Update directory permission settings

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
	updateDirectoryPermissionsRequest := *openapiclient.NewUpdateDirectoryPermissionsRequest() // UpdateDirectoryPermissionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectoryAPI.UpdateDirectoryPermissions(context.Background()).UpdateDirectoryPermissionsRequest(updateDirectoryPermissionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAPI.UpdateDirectoryPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDirectoryPermissions`: DirectoryPermissions
	fmt.Fprintf(os.Stdout, "Response from `DirectoryAPI.UpdateDirectoryPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDirectoryPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDirectoryPermissionsRequest** | [**UpdateDirectoryPermissionsRequest**](UpdateDirectoryPermissionsRequest.md) |  | 

### Return type

[**DirectoryPermissions**](DirectoryPermissions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGALSettings

> GALSettings UpdateGALSettings(ctx).UpdateGALSettingsRequest(updateGALSettingsRequest).Execute()

Update GAL settings

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
	updateGALSettingsRequest := *openapiclient.NewUpdateGALSettingsRequest() // UpdateGALSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectoryAPI.UpdateGALSettings(context.Background()).UpdateGALSettingsRequest(updateGALSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAPI.UpdateGALSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGALSettings`: GALSettings
	fmt.Fprintf(os.Stdout, "Response from `DirectoryAPI.UpdateGALSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGALSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateGALSettingsRequest** | [**UpdateGALSettingsRequest**](UpdateGALSettingsRequest.md) |  | 

### Return type

[**GALSettings**](GALSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

