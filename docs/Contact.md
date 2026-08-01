# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Name** | **string** |  | 
**Email** | **string** |  | 
**Phone** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ContactType** | **string** |  | 
**Source** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewContact

`func NewContact(id string, tenantId string, name string, email string, contactType string, status string, createdAt time.Time, updatedAt time.Time, ) *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contact) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *Contact) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Contact) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Contact) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetName

`func (o *Contact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Contact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Contact) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *Contact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Contact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Contact) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *Contact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Contact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Contact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Contact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCompany

`func (o *Contact) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Contact) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Contact) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Contact) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetNotes

`func (o *Contact) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Contact) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Contact) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Contact) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetContactType

`func (o *Contact) GetContactType() string`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *Contact) GetContactTypeOk() (*string, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *Contact) SetContactType(v string)`

SetContactType sets ContactType field to given value.


### GetSource

`func (o *Contact) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Contact) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Contact) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Contact) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDepartment

`func (o *Contact) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Contact) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Contact) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Contact) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetRole

`func (o *Contact) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Contact) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Contact) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Contact) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *Contact) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Contact) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Contact) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *Contact) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Contact) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Contact) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Contact) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Contact) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Contact) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


