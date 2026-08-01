# \ContactListsAPI

All URIs are relative to *https://api.lockally.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContactListMember**](ContactListsAPI.md#AddContactListMember) | **Post** /v1/contact-lists/{id}/members | Add a member to a contact list
[**CreateContactList**](ContactListsAPI.md#CreateContactList) | **Post** /v1/contact-lists | Create a contact list
[**DeleteContactList**](ContactListsAPI.md#DeleteContactList) | **Delete** /v1/contact-lists/{id} | Delete a contact list
[**GetContactList**](ContactListsAPI.md#GetContactList) | **Get** /v1/contact-lists/{id} | Get a contact list with members
[**ListContactLists**](ContactListsAPI.md#ListContactLists) | **Get** /v1/contact-lists | List contact lists
[**RemoveContactListMember**](ContactListsAPI.md#RemoveContactListMember) | **Delete** /v1/contact-lists/{id}/members/{contactId} | Remove a member from a contact list
[**UpdateContactList**](ContactListsAPI.md#UpdateContactList) | **Patch** /v1/contact-lists/{id} | Update a contact list



## AddContactListMember

> AddContactListMember(ctx, id).AddContactListMemberRequest(addContactListMemberRequest).Execute()

Add a member to a contact list

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
	addContactListMemberRequest := *openapiclient.NewAddContactListMemberRequest("ContactId_example") // AddContactListMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContactListsAPI.AddContactListMember(context.Background(), id).AddContactListMemberRequest(addContactListMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.AddContactListMember``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAddContactListMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addContactListMemberRequest** | [**AddContactListMemberRequest**](AddContactListMemberRequest.md) |  | 

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


## CreateContactList

> ContactList CreateContactList(ctx).CreateContactListRequest(createContactListRequest).Execute()

Create a contact list

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
	createContactListRequest := *openapiclient.NewCreateContactListRequest("Name_example") // CreateContactListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.CreateContactList(context.Background()).CreateContactListRequest(createContactListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.CreateContactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContactList`: ContactList
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.CreateContactList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContactListRequest** | [**CreateContactListRequest**](CreateContactListRequest.md) |  | 

### Return type

[**ContactList**](ContactList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactList

> DeleteContactList(ctx, id).Execute()

Delete a contact list

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
	r, err := apiClient.ContactListsAPI.DeleteContactList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.DeleteContactList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteContactListRequest struct via the builder pattern


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


## GetContactList

> GetContactList200Response GetContactList(ctx, id).Execute()

Get a contact list with members

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
	resp, r, err := apiClient.ContactListsAPI.GetContactList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.GetContactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactList`: GetContactList200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.GetContactList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContactList200Response**](GetContactList200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContactLists

> ListContactLists200Response ListContactLists(ctx).Execute()

List contact lists

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
	resp, r, err := apiClient.ContactListsAPI.ListContactLists(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.ListContactLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContactLists`: ListContactLists200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.ListContactLists`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListContactListsRequest struct via the builder pattern


### Return type

[**ListContactLists200Response**](ListContactLists200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveContactListMember

> RemoveContactListMember(ctx, id, contactId).Execute()

Remove a member from a contact list

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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContactListsAPI.RemoveContactListMember(context.Background(), id, contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.RemoveContactListMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveContactListMemberRequest struct via the builder pattern


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


## UpdateContactList

> ContactList UpdateContactList(ctx, id).UpdateContactListRequest(updateContactListRequest).Execute()

Update a contact list

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
	updateContactListRequest := *openapiclient.NewUpdateContactListRequest() // UpdateContactListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.UpdateContactList(context.Background(), id).UpdateContactListRequest(updateContactListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.UpdateContactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContactList`: ContactList
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.UpdateContactList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateContactListRequest** | [**UpdateContactListRequest**](UpdateContactListRequest.md) |  | 

### Return type

[**ContactList**](ContactList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

