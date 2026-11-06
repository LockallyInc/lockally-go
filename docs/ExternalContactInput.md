# ExternalContactInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Name** | Pointer to **string** | Defaults to the email address when empty. | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ContactType** | Pointer to **string** |  | [optional] [default to "external"]
**Department** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "active"]

## Methods

### NewExternalContactInput

`func NewExternalContactInput(email string, ) *ExternalContactInput`

NewExternalContactInput instantiates a new ExternalContactInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalContactInputWithDefaults

`func NewExternalContactInputWithDefaults() *ExternalContactInput`

NewExternalContactInputWithDefaults instantiates a new ExternalContactInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ExternalContactInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ExternalContactInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ExternalContactInput) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *ExternalContactInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalContactInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalContactInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalContactInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *ExternalContactInput) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ExternalContactInput) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ExternalContactInput) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ExternalContactInput) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCompany

`func (o *ExternalContactInput) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ExternalContactInput) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ExternalContactInput) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ExternalContactInput) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetNotes

`func (o *ExternalContactInput) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ExternalContactInput) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ExternalContactInput) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ExternalContactInput) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetContactType

`func (o *ExternalContactInput) GetContactType() string`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *ExternalContactInput) GetContactTypeOk() (*string, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *ExternalContactInput) SetContactType(v string)`

SetContactType sets ContactType field to given value.

### HasContactType

`func (o *ExternalContactInput) HasContactType() bool`

HasContactType returns a boolean if a field has been set.

### GetDepartment

`func (o *ExternalContactInput) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ExternalContactInput) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ExternalContactInput) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ExternalContactInput) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetRole

`func (o *ExternalContactInput) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ExternalContactInput) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ExternalContactInput) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ExternalContactInput) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *ExternalContactInput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalContactInput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalContactInput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExternalContactInput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


