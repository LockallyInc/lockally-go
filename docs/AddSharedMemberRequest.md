# AddSharedMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberEmail** | **string** |  | 
**Role** | Pointer to **string** |  | [optional] [default to "member"]

## Methods

### NewAddSharedMemberRequest

`func NewAddSharedMemberRequest(memberEmail string, ) *AddSharedMemberRequest`

NewAddSharedMemberRequest instantiates a new AddSharedMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSharedMemberRequestWithDefaults

`func NewAddSharedMemberRequestWithDefaults() *AddSharedMemberRequest`

NewAddSharedMemberRequestWithDefaults instantiates a new AddSharedMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberEmail

`func (o *AddSharedMemberRequest) GetMemberEmail() string`

GetMemberEmail returns the MemberEmail field if non-nil, zero value otherwise.

### GetMemberEmailOk

`func (o *AddSharedMemberRequest) GetMemberEmailOk() (*string, bool)`

GetMemberEmailOk returns a tuple with the MemberEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberEmail

`func (o *AddSharedMemberRequest) SetMemberEmail(v string)`

SetMemberEmail sets MemberEmail field to given value.


### GetRole

`func (o *AddSharedMemberRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddSharedMemberRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddSharedMemberRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AddSharedMemberRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


