# CreateCalendarEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**StartsAt** | **time.Time** |  | 
**EndsAt** | **time.Time** |  | 
**AllDay** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewCreateCalendarEventRequest

`func NewCreateCalendarEventRequest(title string, startsAt time.Time, endsAt time.Time, ) *CreateCalendarEventRequest`

NewCreateCalendarEventRequest instantiates a new CreateCalendarEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCalendarEventRequestWithDefaults

`func NewCreateCalendarEventRequestWithDefaults() *CreateCalendarEventRequest`

NewCreateCalendarEventRequestWithDefaults instantiates a new CreateCalendarEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateCalendarEventRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateCalendarEventRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateCalendarEventRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CreateCalendarEventRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCalendarEventRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCalendarEventRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCalendarEventRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *CreateCalendarEventRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateCalendarEventRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateCalendarEventRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CreateCalendarEventRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStartsAt

`func (o *CreateCalendarEventRequest) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *CreateCalendarEventRequest) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *CreateCalendarEventRequest) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetEndsAt

`func (o *CreateCalendarEventRequest) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *CreateCalendarEventRequest) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *CreateCalendarEventRequest) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### GetAllDay

`func (o *CreateCalendarEventRequest) GetAllDay() bool`

GetAllDay returns the AllDay field if non-nil, zero value otherwise.

### GetAllDayOk

`func (o *CreateCalendarEventRequest) GetAllDayOk() (*bool, bool)`

GetAllDayOk returns a tuple with the AllDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDay

`func (o *CreateCalendarEventRequest) SetAllDay(v bool)`

SetAllDay sets AllDay field to given value.

### HasAllDay

`func (o *CreateCalendarEventRequest) HasAllDay() bool`

HasAllDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


