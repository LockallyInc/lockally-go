# DirectoryPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**ContactViewAccess** | **string** |  | 
**ContactEditAccess** | **string** |  | 
**ListManageAccess** | **string** |  | 
**ExternalSharing** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDirectoryPermissions

`func NewDirectoryPermissions(tenantId string, contactViewAccess string, contactEditAccess string, listManageAccess string, externalSharing string, ) *DirectoryPermissions`

NewDirectoryPermissions instantiates a new DirectoryPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryPermissionsWithDefaults

`func NewDirectoryPermissionsWithDefaults() *DirectoryPermissions`

NewDirectoryPermissionsWithDefaults instantiates a new DirectoryPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DirectoryPermissions) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DirectoryPermissions) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DirectoryPermissions) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetContactViewAccess

`func (o *DirectoryPermissions) GetContactViewAccess() string`

GetContactViewAccess returns the ContactViewAccess field if non-nil, zero value otherwise.

### GetContactViewAccessOk

`func (o *DirectoryPermissions) GetContactViewAccessOk() (*string, bool)`

GetContactViewAccessOk returns a tuple with the ContactViewAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactViewAccess

`func (o *DirectoryPermissions) SetContactViewAccess(v string)`

SetContactViewAccess sets ContactViewAccess field to given value.


### GetContactEditAccess

`func (o *DirectoryPermissions) GetContactEditAccess() string`

GetContactEditAccess returns the ContactEditAccess field if non-nil, zero value otherwise.

### GetContactEditAccessOk

`func (o *DirectoryPermissions) GetContactEditAccessOk() (*string, bool)`

GetContactEditAccessOk returns a tuple with the ContactEditAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEditAccess

`func (o *DirectoryPermissions) SetContactEditAccess(v string)`

SetContactEditAccess sets ContactEditAccess field to given value.


### GetListManageAccess

`func (o *DirectoryPermissions) GetListManageAccess() string`

GetListManageAccess returns the ListManageAccess field if non-nil, zero value otherwise.

### GetListManageAccessOk

`func (o *DirectoryPermissions) GetListManageAccessOk() (*string, bool)`

GetListManageAccessOk returns a tuple with the ListManageAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListManageAccess

`func (o *DirectoryPermissions) SetListManageAccess(v string)`

SetListManageAccess sets ListManageAccess field to given value.


### GetExternalSharing

`func (o *DirectoryPermissions) GetExternalSharing() string`

GetExternalSharing returns the ExternalSharing field if non-nil, zero value otherwise.

### GetExternalSharingOk

`func (o *DirectoryPermissions) GetExternalSharingOk() (*string, bool)`

GetExternalSharingOk returns a tuple with the ExternalSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSharing

`func (o *DirectoryPermissions) SetExternalSharing(v string)`

SetExternalSharing sets ExternalSharing field to given value.


### GetCreatedAt

`func (o *DirectoryPermissions) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DirectoryPermissions) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DirectoryPermissions) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DirectoryPermissions) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DirectoryPermissions) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DirectoryPermissions) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DirectoryPermissions) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DirectoryPermissions) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


