# UpdateCalendarEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**StartsAt** | Pointer to **time.Time** |  | [optional] 
**EndsAt** | Pointer to **time.Time** |  | [optional] 
**AllDay** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateCalendarEventRequest

`func NewUpdateCalendarEventRequest() *UpdateCalendarEventRequest`

NewUpdateCalendarEventRequest instantiates a new UpdateCalendarEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCalendarEventRequestWithDefaults

`func NewUpdateCalendarEventRequestWithDefaults() *UpdateCalendarEventRequest`

NewUpdateCalendarEventRequestWithDefaults instantiates a new UpdateCalendarEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateCalendarEventRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateCalendarEventRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateCalendarEventRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateCalendarEventRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCalendarEventRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCalendarEventRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCalendarEventRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCalendarEventRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *UpdateCalendarEventRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateCalendarEventRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateCalendarEventRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateCalendarEventRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStartsAt

`func (o *UpdateCalendarEventRequest) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *UpdateCalendarEventRequest) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *UpdateCalendarEventRequest) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *UpdateCalendarEventRequest) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetEndsAt

`func (o *UpdateCalendarEventRequest) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *UpdateCalendarEventRequest) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *UpdateCalendarEventRequest) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *UpdateCalendarEventRequest) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetAllDay

`func (o *UpdateCalendarEventRequest) GetAllDay() bool`

GetAllDay returns the AllDay field if non-nil, zero value otherwise.

### GetAllDayOk

`func (o *UpdateCalendarEventRequest) GetAllDayOk() (*bool, bool)`

GetAllDayOk returns a tuple with the AllDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDay

`func (o *UpdateCalendarEventRequest) SetAllDay(v bool)`

SetAllDay sets AllDay field to given value.

### HasAllDay

`func (o *UpdateCalendarEventRequest) HasAllDay() bool`

HasAllDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


