# AddCalendarMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserEmail** | **string** |  | 
**Role** | Pointer to **string** |  | [optional] [default to "viewer"]

## Methods

### NewAddCalendarMemberRequest

`func NewAddCalendarMemberRequest(userEmail string, ) *AddCalendarMemberRequest`

NewAddCalendarMemberRequest instantiates a new AddCalendarMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCalendarMemberRequestWithDefaults

`func NewAddCalendarMemberRequestWithDefaults() *AddCalendarMemberRequest`

NewAddCalendarMemberRequestWithDefaults instantiates a new AddCalendarMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserEmail

`func (o *AddCalendarMemberRequest) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *AddCalendarMemberRequest) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *AddCalendarMemberRequest) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetRole

`func (o *AddCalendarMemberRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddCalendarMemberRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddCalendarMemberRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AddCalendarMemberRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


