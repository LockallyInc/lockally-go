# GALSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**GalEnabled** | **bool** |  | 
**HideFromDirectory** | **bool** |  | 
**DepartmentGrouping** | **bool** |  | 
**SearchVisibility** | **string** |  | 
**IncludeExternalContacts** | **bool** |  | 
**LastIndexRebuiltAt** | Pointer to **time.Time** |  | [optional] 
**LastSyncedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGALSettings

`func NewGALSettings(tenantId string, galEnabled bool, hideFromDirectory bool, departmentGrouping bool, searchVisibility string, includeExternalContacts bool, ) *GALSettings`

NewGALSettings instantiates a new GALSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGALSettingsWithDefaults

`func NewGALSettingsWithDefaults() *GALSettings`

NewGALSettingsWithDefaults instantiates a new GALSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *GALSettings) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GALSettings) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GALSettings) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetGalEnabled

`func (o *GALSettings) GetGalEnabled() bool`

GetGalEnabled returns the GalEnabled field if non-nil, zero value otherwise.

### GetGalEnabledOk

`func (o *GALSettings) GetGalEnabledOk() (*bool, bool)`

GetGalEnabledOk returns a tuple with the GalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGalEnabled

`func (o *GALSettings) SetGalEnabled(v bool)`

SetGalEnabled sets GalEnabled field to given value.


### GetHideFromDirectory

`func (o *GALSettings) GetHideFromDirectory() bool`

GetHideFromDirectory returns the HideFromDirectory field if non-nil, zero value otherwise.

### GetHideFromDirectoryOk

`func (o *GALSettings) GetHideFromDirectoryOk() (*bool, bool)`

GetHideFromDirectoryOk returns a tuple with the HideFromDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromDirectory

`func (o *GALSettings) SetHideFromDirectory(v bool)`

SetHideFromDirectory sets HideFromDirectory field to given value.


### GetDepartmentGrouping

`func (o *GALSettings) GetDepartmentGrouping() bool`

GetDepartmentGrouping returns the DepartmentGrouping field if non-nil, zero value otherwise.

### GetDepartmentGroupingOk

`func (o *GALSettings) GetDepartmentGroupingOk() (*bool, bool)`

GetDepartmentGroupingOk returns a tuple with the DepartmentGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentGrouping

`func (o *GALSettings) SetDepartmentGrouping(v bool)`

SetDepartmentGrouping sets DepartmentGrouping field to given value.


### GetSearchVisibility

`func (o *GALSettings) GetSearchVisibility() string`

GetSearchVisibility returns the SearchVisibility field if non-nil, zero value otherwise.

### GetSearchVisibilityOk

`func (o *GALSettings) GetSearchVisibilityOk() (*string, bool)`

GetSearchVisibilityOk returns a tuple with the SearchVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchVisibility

`func (o *GALSettings) SetSearchVisibility(v string)`

SetSearchVisibility sets SearchVisibility field to given value.


### GetIncludeExternalContacts

`func (o *GALSettings) GetIncludeExternalContacts() bool`

GetIncludeExternalContacts returns the IncludeExternalContacts field if non-nil, zero value otherwise.

### GetIncludeExternalContactsOk

`func (o *GALSettings) GetIncludeExternalContactsOk() (*bool, bool)`

GetIncludeExternalContactsOk returns a tuple with the IncludeExternalContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExternalContacts

`func (o *GALSettings) SetIncludeExternalContacts(v bool)`

SetIncludeExternalContacts sets IncludeExternalContacts field to given value.


### GetLastIndexRebuiltAt

`func (o *GALSettings) GetLastIndexRebuiltAt() time.Time`

GetLastIndexRebuiltAt returns the LastIndexRebuiltAt field if non-nil, zero value otherwise.

### GetLastIndexRebuiltAtOk

`func (o *GALSettings) GetLastIndexRebuiltAtOk() (*time.Time, bool)`

GetLastIndexRebuiltAtOk returns a tuple with the LastIndexRebuiltAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexRebuiltAt

`func (o *GALSettings) SetLastIndexRebuiltAt(v time.Time)`

SetLastIndexRebuiltAt sets LastIndexRebuiltAt field to given value.

### HasLastIndexRebuiltAt

`func (o *GALSettings) HasLastIndexRebuiltAt() bool`

HasLastIndexRebuiltAt returns a boolean if a field has been set.

### GetLastSyncedAt

`func (o *GALSettings) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *GALSettings) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *GALSettings) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *GALSettings) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GALSettings) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GALSettings) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GALSettings) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GALSettings) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GALSettings) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GALSettings) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GALSettings) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GALSettings) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


