# AdminFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Email** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Role** | **string** |  | 
**LastLoginAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Disabled** | **bool** |  | 
**DisabledAt** | Pointer to **time.Time** |  | [optional] 
**Password** | Pointer to **string** | Present ONLY on POST response when lockally generated the password. Shown once. | [optional] 

## Methods

### NewAdminFull

`func NewAdminFull(id string, tenantId string, email string, role string, createdAt time.Time, disabled bool, ) *AdminFull`

NewAdminFull instantiates a new AdminFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminFullWithDefaults

`func NewAdminFullWithDefaults() *AdminFull`

NewAdminFullWithDefaults instantiates a new AdminFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminFull) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *AdminFull) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AdminFull) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AdminFull) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetEmail

`func (o *AdminFull) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminFull) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminFull) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDisplayName

`func (o *AdminFull) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AdminFull) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AdminFull) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AdminFull) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRole

`func (o *AdminFull) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AdminFull) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AdminFull) SetRole(v string)`

SetRole sets Role field to given value.


### GetLastLoginAt

`func (o *AdminFull) GetLastLoginAt() time.Time`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *AdminFull) GetLastLoginAtOk() (*time.Time, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *AdminFull) SetLastLoginAt(v time.Time)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *AdminFull) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminFull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminFull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminFull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDisabled

`func (o *AdminFull) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AdminFull) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AdminFull) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetDisabledAt

`func (o *AdminFull) GetDisabledAt() time.Time`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *AdminFull) GetDisabledAtOk() (*time.Time, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *AdminFull) SetDisabledAt(v time.Time)`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *AdminFull) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### GetPassword

`func (o *AdminFull) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AdminFull) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AdminFull) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AdminFull) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


