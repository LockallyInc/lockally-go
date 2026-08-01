# CreateContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Email** | **string** |  | 
**Phone** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ContactType** | Pointer to **string** |  | [optional] [default to "external"]
**Department** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateContactRequest

`func NewCreateContactRequest(name string, email string, ) *CreateContactRequest`

NewCreateContactRequest instantiates a new CreateContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContactRequestWithDefaults

`func NewCreateContactRequestWithDefaults() *CreateContactRequest`

NewCreateContactRequestWithDefaults instantiates a new CreateContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateContactRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateContactRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateContactRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CreateContactRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateContactRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateContactRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *CreateContactRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CreateContactRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CreateContactRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CreateContactRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCompany

`func (o *CreateContactRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CreateContactRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CreateContactRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CreateContactRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetNotes

`func (o *CreateContactRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateContactRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateContactRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateContactRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetContactType

`func (o *CreateContactRequest) GetContactType() string`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *CreateContactRequest) GetContactTypeOk() (*string, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *CreateContactRequest) SetContactType(v string)`

SetContactType sets ContactType field to given value.

### HasContactType

`func (o *CreateContactRequest) HasContactType() bool`

HasContactType returns a boolean if a field has been set.

### GetDepartment

`func (o *CreateContactRequest) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *CreateContactRequest) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *CreateContactRequest) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *CreateContactRequest) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetRole

`func (o *CreateContactRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateContactRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateContactRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CreateContactRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


