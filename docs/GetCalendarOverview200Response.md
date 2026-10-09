# GetCalendarOverview200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCalendars** | Pointer to **int32** |  | [optional] 
**ActiveCalendars** | Pointer to **int32** | Calendars with an event created, edited or held in the last 30 days. | [optional] 
**SharedCalendars** | Pointer to **int32** | Calendars visible to the whole organization. | [optional] 
**EventsThisMonth** | Pointer to **int32** | Events on every calendar, in the account&#39;s own timezone. | [optional] 
**UpcomingOrgEvents** | Pointer to **int32** | Future events on organization-visible calendars. | [optional] 
**BookableResources** | Pointer to **int32** | Resources that are not disabled. | [optional] 
**AccessGranted** | Pointer to **int32** | People who own a calendar or are named on one. | [optional] 
**AccessOrgWide** | Pointer to **int32** | Everyone who can see an organization calendar; 0 when none exists. | [optional] 
**Activity** | Pointer to [**[]GetCalendarOverview200ResponseActivityInner**](GetCalendarOverview200ResponseActivityInner.md) |  | [optional] 

## Methods

### NewGetCalendarOverview200Response

`func NewGetCalendarOverview200Response() *GetCalendarOverview200Response`

NewGetCalendarOverview200Response instantiates a new GetCalendarOverview200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCalendarOverview200ResponseWithDefaults

`func NewGetCalendarOverview200ResponseWithDefaults() *GetCalendarOverview200Response`

NewGetCalendarOverview200ResponseWithDefaults instantiates a new GetCalendarOverview200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCalendars

`func (o *GetCalendarOverview200Response) GetTotalCalendars() int32`

GetTotalCalendars returns the TotalCalendars field if non-nil, zero value otherwise.

### GetTotalCalendarsOk

`func (o *GetCalendarOverview200Response) GetTotalCalendarsOk() (*int32, bool)`

GetTotalCalendarsOk returns a tuple with the TotalCalendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCalendars

`func (o *GetCalendarOverview200Response) SetTotalCalendars(v int32)`

SetTotalCalendars sets TotalCalendars field to given value.

### HasTotalCalendars

`func (o *GetCalendarOverview200Response) HasTotalCalendars() bool`

HasTotalCalendars returns a boolean if a field has been set.

### GetActiveCalendars

`func (o *GetCalendarOverview200Response) GetActiveCalendars() int32`

GetActiveCalendars returns the ActiveCalendars field if non-nil, zero value otherwise.

### GetActiveCalendarsOk

`func (o *GetCalendarOverview200Response) GetActiveCalendarsOk() (*int32, bool)`

GetActiveCalendarsOk returns a tuple with the ActiveCalendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCalendars

`func (o *GetCalendarOverview200Response) SetActiveCalendars(v int32)`

SetActiveCalendars sets ActiveCalendars field to given value.

### HasActiveCalendars

`func (o *GetCalendarOverview200Response) HasActiveCalendars() bool`

HasActiveCalendars returns a boolean if a field has been set.

### GetSharedCalendars

`func (o *GetCalendarOverview200Response) GetSharedCalendars() int32`

GetSharedCalendars returns the SharedCalendars field if non-nil, zero value otherwise.

### GetSharedCalendarsOk

`func (o *GetCalendarOverview200Response) GetSharedCalendarsOk() (*int32, bool)`

GetSharedCalendarsOk returns a tuple with the SharedCalendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedCalendars

`func (o *GetCalendarOverview200Response) SetSharedCalendars(v int32)`

SetSharedCalendars sets SharedCalendars field to given value.

### HasSharedCalendars

`func (o *GetCalendarOverview200Response) HasSharedCalendars() bool`

HasSharedCalendars returns a boolean if a field has been set.

### GetEventsThisMonth

`func (o *GetCalendarOverview200Response) GetEventsThisMonth() int32`

GetEventsThisMonth returns the EventsThisMonth field if non-nil, zero value otherwise.

### GetEventsThisMonthOk

`func (o *GetCalendarOverview200Response) GetEventsThisMonthOk() (*int32, bool)`

GetEventsThisMonthOk returns a tuple with the EventsThisMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsThisMonth

`func (o *GetCalendarOverview200Response) SetEventsThisMonth(v int32)`

SetEventsThisMonth sets EventsThisMonth field to given value.

### HasEventsThisMonth

`func (o *GetCalendarOverview200Response) HasEventsThisMonth() bool`

HasEventsThisMonth returns a boolean if a field has been set.

### GetUpcomingOrgEvents

`func (o *GetCalendarOverview200Response) GetUpcomingOrgEvents() int32`

GetUpcomingOrgEvents returns the UpcomingOrgEvents field if non-nil, zero value otherwise.

### GetUpcomingOrgEventsOk

`func (o *GetCalendarOverview200Response) GetUpcomingOrgEventsOk() (*int32, bool)`

GetUpcomingOrgEventsOk returns a tuple with the UpcomingOrgEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcomingOrgEvents

`func (o *GetCalendarOverview200Response) SetUpcomingOrgEvents(v int32)`

SetUpcomingOrgEvents sets UpcomingOrgEvents field to given value.

### HasUpcomingOrgEvents

`func (o *GetCalendarOverview200Response) HasUpcomingOrgEvents() bool`

HasUpcomingOrgEvents returns a boolean if a field has been set.

### GetBookableResources

`func (o *GetCalendarOverview200Response) GetBookableResources() int32`

GetBookableResources returns the BookableResources field if non-nil, zero value otherwise.

### GetBookableResourcesOk

`func (o *GetCalendarOverview200Response) GetBookableResourcesOk() (*int32, bool)`

GetBookableResourcesOk returns a tuple with the BookableResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookableResources

`func (o *GetCalendarOverview200Response) SetBookableResources(v int32)`

SetBookableResources sets BookableResources field to given value.

### HasBookableResources

`func (o *GetCalendarOverview200Response) HasBookableResources() bool`

HasBookableResources returns a boolean if a field has been set.

### GetAccessGranted

`func (o *GetCalendarOverview200Response) GetAccessGranted() int32`

GetAccessGranted returns the AccessGranted field if non-nil, zero value otherwise.

### GetAccessGrantedOk

`func (o *GetCalendarOverview200Response) GetAccessGrantedOk() (*int32, bool)`

GetAccessGrantedOk returns a tuple with the AccessGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGranted

`func (o *GetCalendarOverview200Response) SetAccessGranted(v int32)`

SetAccessGranted sets AccessGranted field to given value.

### HasAccessGranted

`func (o *GetCalendarOverview200Response) HasAccessGranted() bool`

HasAccessGranted returns a boolean if a field has been set.

### GetAccessOrgWide

`func (o *GetCalendarOverview200Response) GetAccessOrgWide() int32`

GetAccessOrgWide returns the AccessOrgWide field if non-nil, zero value otherwise.

### GetAccessOrgWideOk

`func (o *GetCalendarOverview200Response) GetAccessOrgWideOk() (*int32, bool)`

GetAccessOrgWideOk returns a tuple with the AccessOrgWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessOrgWide

`func (o *GetCalendarOverview200Response) SetAccessOrgWide(v int32)`

SetAccessOrgWide sets AccessOrgWide field to given value.

### HasAccessOrgWide

`func (o *GetCalendarOverview200Response) HasAccessOrgWide() bool`

HasAccessOrgWide returns a boolean if a field has been set.

### GetActivity

`func (o *GetCalendarOverview200Response) GetActivity() []GetCalendarOverview200ResponseActivityInner`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *GetCalendarOverview200Response) GetActivityOk() (*[]GetCalendarOverview200ResponseActivityInner, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *GetCalendarOverview200Response) SetActivity(v []GetCalendarOverview200ResponseActivityInner)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *GetCalendarOverview200Response) HasActivity() bool`

HasActivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


