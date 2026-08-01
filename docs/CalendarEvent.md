# CalendarEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CalendarId** | **string** |  | 
**TenantId** | **string** |  | 
**Uid** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**StartsAt** | **time.Time** |  | 
**EndsAt** | **time.Time** |  | 
**AllDay** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCalendarEvent

`func NewCalendarEvent(id string, calendarId string, tenantId string, uid string, title string, startsAt time.Time, endsAt time.Time, allDay bool, createdAt time.Time, updatedAt time.Time, ) *CalendarEvent`

NewCalendarEvent instantiates a new CalendarEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarEventWithDefaults

`func NewCalendarEventWithDefaults() *CalendarEvent`

NewCalendarEventWithDefaults instantiates a new CalendarEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CalendarEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CalendarEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CalendarEvent) SetId(v string)`

SetId sets Id field to given value.


### GetCalendarId

`func (o *CalendarEvent) GetCalendarId() string`

GetCalendarId returns the CalendarId field if non-nil, zero value otherwise.

### GetCalendarIdOk

`func (o *CalendarEvent) GetCalendarIdOk() (*string, bool)`

GetCalendarIdOk returns a tuple with the CalendarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarId

`func (o *CalendarEvent) SetCalendarId(v string)`

SetCalendarId sets CalendarId field to given value.


### GetTenantId

`func (o *CalendarEvent) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CalendarEvent) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CalendarEvent) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUid

`func (o *CalendarEvent) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *CalendarEvent) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *CalendarEvent) SetUid(v string)`

SetUid sets Uid field to given value.


### GetTitle

`func (o *CalendarEvent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CalendarEvent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CalendarEvent) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CalendarEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CalendarEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CalendarEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CalendarEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *CalendarEvent) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CalendarEvent) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CalendarEvent) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CalendarEvent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStartsAt

`func (o *CalendarEvent) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *CalendarEvent) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *CalendarEvent) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetEndsAt

`func (o *CalendarEvent) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *CalendarEvent) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *CalendarEvent) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### GetAllDay

`func (o *CalendarEvent) GetAllDay() bool`

GetAllDay returns the AllDay field if non-nil, zero value otherwise.

### GetAllDayOk

`func (o *CalendarEvent) GetAllDayOk() (*bool, bool)`

GetAllDayOk returns a tuple with the AllDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDay

`func (o *CalendarEvent) SetAllDay(v bool)`

SetAllDay sets AllDay field to given value.


### GetCreatedAt

`func (o *CalendarEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CalendarEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CalendarEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CalendarEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CalendarEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CalendarEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


