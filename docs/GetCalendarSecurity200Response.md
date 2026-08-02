# GetCalendarSecurity200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCalendars** | Pointer to **int32** |  | [optional] 
**PublicCalendars** | Pointer to **int32** |  | [optional] 
**PrivateCalendars** | Pointer to **int32** |  | [optional] 
**TotalMembers** | Pointer to **int32** |  | [optional] 
**DelegatedAccess** | Pointer to [**[]GetCalendarSecurity200ResponseDelegatedAccessInner**](GetCalendarSecurity200ResponseDelegatedAccessInner.md) |  | [optional] 
**PublicCalendarList** | Pointer to [**[]GetCalendarSecurity200ResponsePublicCalendarListInner**](GetCalendarSecurity200ResponsePublicCalendarListInner.md) |  | [optional] 
**Alerts** | Pointer to [**[]GetCalendarSecurity200ResponseAlertsInner**](GetCalendarSecurity200ResponseAlertsInner.md) |  | [optional] 
**ExternalSharing** | Pointer to [**GetCalendarSecurity200ResponseExternalSharing**](GetCalendarSecurity200ResponseExternalSharing.md) |  | [optional] 

## Methods

### NewGetCalendarSecurity200Response

`func NewGetCalendarSecurity200Response() *GetCalendarSecurity200Response`

NewGetCalendarSecurity200Response instantiates a new GetCalendarSecurity200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCalendarSecurity200ResponseWithDefaults

`func NewGetCalendarSecurity200ResponseWithDefaults() *GetCalendarSecurity200Response`

NewGetCalendarSecurity200ResponseWithDefaults instantiates a new GetCalendarSecurity200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCalendars

`func (o *GetCalendarSecurity200Response) GetTotalCalendars() int32`

GetTotalCalendars returns the TotalCalendars field if non-nil, zero value otherwise.

### GetTotalCalendarsOk

`func (o *GetCalendarSecurity200Response) GetTotalCalendarsOk() (*int32, bool)`

GetTotalCalendarsOk returns a tuple with the TotalCalendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCalendars

`func (o *GetCalendarSecurity200Response) SetTotalCalendars(v int32)`

SetTotalCalendars sets TotalCalendars field to given value.

### HasTotalCalendars

`func (o *GetCalendarSecurity200Response) HasTotalCalendars() bool`

HasTotalCalendars returns a boolean if a field has been set.

### GetPublicCalendars

`func (o *GetCalendarSecurity200Response) GetPublicCalendars() int32`

GetPublicCalendars returns the PublicCalendars field if non-nil, zero value otherwise.

### GetPublicCalendarsOk

`func (o *GetCalendarSecurity200Response) GetPublicCalendarsOk() (*int32, bool)`

GetPublicCalendarsOk returns a tuple with the PublicCalendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicCalendars

`func (o *GetCalendarSecurity200Response) SetPublicCalendars(v int32)`

SetPublicCalendars sets PublicCalendars field to given value.

### HasPublicCalendars

`func (o *GetCalendarSecurity200Response) HasPublicCalendars() bool`

HasPublicCalendars returns a boolean if a field has been set.

### GetPrivateCalendars

`func (o *GetCalendarSecurity200Response) GetPrivateCalendars() int32`

GetPrivateCalendars returns the PrivateCalendars field if non-nil, zero value otherwise.

### GetPrivateCalendarsOk

`func (o *GetCalendarSecurity200Response) GetPrivateCalendarsOk() (*int32, bool)`

GetPrivateCalendarsOk returns a tuple with the PrivateCalendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateCalendars

`func (o *GetCalendarSecurity200Response) SetPrivateCalendars(v int32)`

SetPrivateCalendars sets PrivateCalendars field to given value.

### HasPrivateCalendars

`func (o *GetCalendarSecurity200Response) HasPrivateCalendars() bool`

HasPrivateCalendars returns a boolean if a field has been set.

### GetTotalMembers

`func (o *GetCalendarSecurity200Response) GetTotalMembers() int32`

GetTotalMembers returns the TotalMembers field if non-nil, zero value otherwise.

### GetTotalMembersOk

`func (o *GetCalendarSecurity200Response) GetTotalMembersOk() (*int32, bool)`

GetTotalMembersOk returns a tuple with the TotalMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMembers

`func (o *GetCalendarSecurity200Response) SetTotalMembers(v int32)`

SetTotalMembers sets TotalMembers field to given value.

### HasTotalMembers

`func (o *GetCalendarSecurity200Response) HasTotalMembers() bool`

HasTotalMembers returns a boolean if a field has been set.

### GetDelegatedAccess

`func (o *GetCalendarSecurity200Response) GetDelegatedAccess() []GetCalendarSecurity200ResponseDelegatedAccessInner`

GetDelegatedAccess returns the DelegatedAccess field if non-nil, zero value otherwise.

### GetDelegatedAccessOk

`func (o *GetCalendarSecurity200Response) GetDelegatedAccessOk() (*[]GetCalendarSecurity200ResponseDelegatedAccessInner, bool)`

GetDelegatedAccessOk returns a tuple with the DelegatedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAccess

`func (o *GetCalendarSecurity200Response) SetDelegatedAccess(v []GetCalendarSecurity200ResponseDelegatedAccessInner)`

SetDelegatedAccess sets DelegatedAccess field to given value.

### HasDelegatedAccess

`func (o *GetCalendarSecurity200Response) HasDelegatedAccess() bool`

HasDelegatedAccess returns a boolean if a field has been set.

### GetPublicCalendarList

`func (o *GetCalendarSecurity200Response) GetPublicCalendarList() []GetCalendarSecurity200ResponsePublicCalendarListInner`

GetPublicCalendarList returns the PublicCalendarList field if non-nil, zero value otherwise.

### GetPublicCalendarListOk

`func (o *GetCalendarSecurity200Response) GetPublicCalendarListOk() (*[]GetCalendarSecurity200ResponsePublicCalendarListInner, bool)`

GetPublicCalendarListOk returns a tuple with the PublicCalendarList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicCalendarList

`func (o *GetCalendarSecurity200Response) SetPublicCalendarList(v []GetCalendarSecurity200ResponsePublicCalendarListInner)`

SetPublicCalendarList sets PublicCalendarList field to given value.

### HasPublicCalendarList

`func (o *GetCalendarSecurity200Response) HasPublicCalendarList() bool`

HasPublicCalendarList returns a boolean if a field has been set.

### GetAlerts

`func (o *GetCalendarSecurity200Response) GetAlerts() []GetCalendarSecurity200ResponseAlertsInner`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *GetCalendarSecurity200Response) GetAlertsOk() (*[]GetCalendarSecurity200ResponseAlertsInner, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *GetCalendarSecurity200Response) SetAlerts(v []GetCalendarSecurity200ResponseAlertsInner)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *GetCalendarSecurity200Response) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetExternalSharing

`func (o *GetCalendarSecurity200Response) GetExternalSharing() GetCalendarSecurity200ResponseExternalSharing`

GetExternalSharing returns the ExternalSharing field if non-nil, zero value otherwise.

### GetExternalSharingOk

`func (o *GetCalendarSecurity200Response) GetExternalSharingOk() (*GetCalendarSecurity200ResponseExternalSharing, bool)`

GetExternalSharingOk returns a tuple with the ExternalSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSharing

`func (o *GetCalendarSecurity200Response) SetExternalSharing(v GetCalendarSecurity200ResponseExternalSharing)`

SetExternalSharing sets ExternalSharing field to given value.

### HasExternalSharing

`func (o *GetCalendarSecurity200Response) HasExternalSharing() bool`

HasExternalSharing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


