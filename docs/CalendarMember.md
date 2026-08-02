# CalendarMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CalendarId** | **string** |  | 
**TenantId** | **string** |  | 
**UserEmail** | **string** |  | 
**Role** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewCalendarMember

`func NewCalendarMember(id string, calendarId string, tenantId string, userEmail string, role string, createdAt time.Time, ) *CalendarMember`

NewCalendarMember instantiates a new CalendarMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarMemberWithDefaults

`func NewCalendarMemberWithDefaults() *CalendarMember`

NewCalendarMemberWithDefaults instantiates a new CalendarMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CalendarMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CalendarMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CalendarMember) SetId(v string)`

SetId sets Id field to given value.


### GetCalendarId

`func (o *CalendarMember) GetCalendarId() string`

GetCalendarId returns the CalendarId field if non-nil, zero value otherwise.

### GetCalendarIdOk

`func (o *CalendarMember) GetCalendarIdOk() (*string, bool)`

GetCalendarIdOk returns a tuple with the CalendarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarId

`func (o *CalendarMember) SetCalendarId(v string)`

SetCalendarId sets CalendarId field to given value.


### GetTenantId

`func (o *CalendarMember) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CalendarMember) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CalendarMember) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUserEmail

`func (o *CalendarMember) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *CalendarMember) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *CalendarMember) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetRole

`func (o *CalendarMember) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CalendarMember) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CalendarMember) SetRole(v string)`

SetRole sets Role field to given value.


### GetCreatedAt

`func (o *CalendarMember) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CalendarMember) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CalendarMember) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


