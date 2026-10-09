# SharedMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MailboxId** | **string** |  | 
**TenantId** | **string** |  | 
**MemberEmail** | **string** |  | 
**Role** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewSharedMember

`func NewSharedMember(id string, mailboxId string, tenantId string, memberEmail string, role string, createdAt time.Time, ) *SharedMember`

NewSharedMember instantiates a new SharedMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedMemberWithDefaults

`func NewSharedMemberWithDefaults() *SharedMember`

NewSharedMemberWithDefaults instantiates a new SharedMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SharedMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedMember) SetId(v string)`

SetId sets Id field to given value.


### GetMailboxId

`func (o *SharedMember) GetMailboxId() string`

GetMailboxId returns the MailboxId field if non-nil, zero value otherwise.

### GetMailboxIdOk

`func (o *SharedMember) GetMailboxIdOk() (*string, bool)`

GetMailboxIdOk returns a tuple with the MailboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxId

`func (o *SharedMember) SetMailboxId(v string)`

SetMailboxId sets MailboxId field to given value.


### GetTenantId

`func (o *SharedMember) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SharedMember) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SharedMember) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetMemberEmail

`func (o *SharedMember) GetMemberEmail() string`

GetMemberEmail returns the MemberEmail field if non-nil, zero value otherwise.

### GetMemberEmailOk

`func (o *SharedMember) GetMemberEmailOk() (*string, bool)`

GetMemberEmailOk returns a tuple with the MemberEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberEmail

`func (o *SharedMember) SetMemberEmail(v string)`

SetMemberEmail sets MemberEmail field to given value.


### GetRole

`func (o *SharedMember) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SharedMember) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SharedMember) SetRole(v string)`

SetRole sets Role field to given value.


### GetCreatedAt

`func (o *SharedMember) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SharedMember) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SharedMember) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


