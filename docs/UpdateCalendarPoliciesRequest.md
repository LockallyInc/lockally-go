# UpdateCalendarPoliciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxMeetingDurationMins** | Pointer to **int32** |  | [optional] 
**WorkingHoursStart** | Pointer to **string** |  | [optional] 
**WorkingHoursEnd** | Pointer to **string** |  | [optional] 
**BookingWindowDays** | Pointer to **int32** |  | [optional] 
**RecurringMeetingLimit** | Pointer to **int32** |  | [optional] 
**ResourceApprovalMode** | Pointer to **string** |  | [optional] 
**ExternalInvitesAllowed** | Pointer to **bool** |  | [optional] 
**ExternalSharingAllowed** | Pointer to **bool** |  | [optional] 
**PublicLinksEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateCalendarPoliciesRequest

`func NewUpdateCalendarPoliciesRequest() *UpdateCalendarPoliciesRequest`

NewUpdateCalendarPoliciesRequest instantiates a new UpdateCalendarPoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCalendarPoliciesRequestWithDefaults

`func NewUpdateCalendarPoliciesRequestWithDefaults() *UpdateCalendarPoliciesRequest`

NewUpdateCalendarPoliciesRequestWithDefaults instantiates a new UpdateCalendarPoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxMeetingDurationMins

`func (o *UpdateCalendarPoliciesRequest) GetMaxMeetingDurationMins() int32`

GetMaxMeetingDurationMins returns the MaxMeetingDurationMins field if non-nil, zero value otherwise.

### GetMaxMeetingDurationMinsOk

`func (o *UpdateCalendarPoliciesRequest) GetMaxMeetingDurationMinsOk() (*int32, bool)`

GetMaxMeetingDurationMinsOk returns a tuple with the MaxMeetingDurationMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMeetingDurationMins

`func (o *UpdateCalendarPoliciesRequest) SetMaxMeetingDurationMins(v int32)`

SetMaxMeetingDurationMins sets MaxMeetingDurationMins field to given value.

### HasMaxMeetingDurationMins

`func (o *UpdateCalendarPoliciesRequest) HasMaxMeetingDurationMins() bool`

HasMaxMeetingDurationMins returns a boolean if a field has been set.

### GetWorkingHoursStart

`func (o *UpdateCalendarPoliciesRequest) GetWorkingHoursStart() string`

GetWorkingHoursStart returns the WorkingHoursStart field if non-nil, zero value otherwise.

### GetWorkingHoursStartOk

`func (o *UpdateCalendarPoliciesRequest) GetWorkingHoursStartOk() (*string, bool)`

GetWorkingHoursStartOk returns a tuple with the WorkingHoursStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHoursStart

`func (o *UpdateCalendarPoliciesRequest) SetWorkingHoursStart(v string)`

SetWorkingHoursStart sets WorkingHoursStart field to given value.

### HasWorkingHoursStart

`func (o *UpdateCalendarPoliciesRequest) HasWorkingHoursStart() bool`

HasWorkingHoursStart returns a boolean if a field has been set.

### GetWorkingHoursEnd

`func (o *UpdateCalendarPoliciesRequest) GetWorkingHoursEnd() string`

GetWorkingHoursEnd returns the WorkingHoursEnd field if non-nil, zero value otherwise.

### GetWorkingHoursEndOk

`func (o *UpdateCalendarPoliciesRequest) GetWorkingHoursEndOk() (*string, bool)`

GetWorkingHoursEndOk returns a tuple with the WorkingHoursEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHoursEnd

`func (o *UpdateCalendarPoliciesRequest) SetWorkingHoursEnd(v string)`

SetWorkingHoursEnd sets WorkingHoursEnd field to given value.

### HasWorkingHoursEnd

`func (o *UpdateCalendarPoliciesRequest) HasWorkingHoursEnd() bool`

HasWorkingHoursEnd returns a boolean if a field has been set.

### GetBookingWindowDays

`func (o *UpdateCalendarPoliciesRequest) GetBookingWindowDays() int32`

GetBookingWindowDays returns the BookingWindowDays field if non-nil, zero value otherwise.

### GetBookingWindowDaysOk

`func (o *UpdateCalendarPoliciesRequest) GetBookingWindowDaysOk() (*int32, bool)`

GetBookingWindowDaysOk returns a tuple with the BookingWindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingWindowDays

`func (o *UpdateCalendarPoliciesRequest) SetBookingWindowDays(v int32)`

SetBookingWindowDays sets BookingWindowDays field to given value.

### HasBookingWindowDays

`func (o *UpdateCalendarPoliciesRequest) HasBookingWindowDays() bool`

HasBookingWindowDays returns a boolean if a field has been set.

### GetRecurringMeetingLimit

`func (o *UpdateCalendarPoliciesRequest) GetRecurringMeetingLimit() int32`

GetRecurringMeetingLimit returns the RecurringMeetingLimit field if non-nil, zero value otherwise.

### GetRecurringMeetingLimitOk

`func (o *UpdateCalendarPoliciesRequest) GetRecurringMeetingLimitOk() (*int32, bool)`

GetRecurringMeetingLimitOk returns a tuple with the RecurringMeetingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringMeetingLimit

`func (o *UpdateCalendarPoliciesRequest) SetRecurringMeetingLimit(v int32)`

SetRecurringMeetingLimit sets RecurringMeetingLimit field to given value.

### HasRecurringMeetingLimit

`func (o *UpdateCalendarPoliciesRequest) HasRecurringMeetingLimit() bool`

HasRecurringMeetingLimit returns a boolean if a field has been set.

### GetResourceApprovalMode

`func (o *UpdateCalendarPoliciesRequest) GetResourceApprovalMode() string`

GetResourceApprovalMode returns the ResourceApprovalMode field if non-nil, zero value otherwise.

### GetResourceApprovalModeOk

`func (o *UpdateCalendarPoliciesRequest) GetResourceApprovalModeOk() (*string, bool)`

GetResourceApprovalModeOk returns a tuple with the ResourceApprovalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceApprovalMode

`func (o *UpdateCalendarPoliciesRequest) SetResourceApprovalMode(v string)`

SetResourceApprovalMode sets ResourceApprovalMode field to given value.

### HasResourceApprovalMode

`func (o *UpdateCalendarPoliciesRequest) HasResourceApprovalMode() bool`

HasResourceApprovalMode returns a boolean if a field has been set.

### GetExternalInvitesAllowed

`func (o *UpdateCalendarPoliciesRequest) GetExternalInvitesAllowed() bool`

GetExternalInvitesAllowed returns the ExternalInvitesAllowed field if non-nil, zero value otherwise.

### GetExternalInvitesAllowedOk

`func (o *UpdateCalendarPoliciesRequest) GetExternalInvitesAllowedOk() (*bool, bool)`

GetExternalInvitesAllowedOk returns a tuple with the ExternalInvitesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInvitesAllowed

`func (o *UpdateCalendarPoliciesRequest) SetExternalInvitesAllowed(v bool)`

SetExternalInvitesAllowed sets ExternalInvitesAllowed field to given value.

### HasExternalInvitesAllowed

`func (o *UpdateCalendarPoliciesRequest) HasExternalInvitesAllowed() bool`

HasExternalInvitesAllowed returns a boolean if a field has been set.

### GetExternalSharingAllowed

`func (o *UpdateCalendarPoliciesRequest) GetExternalSharingAllowed() bool`

GetExternalSharingAllowed returns the ExternalSharingAllowed field if non-nil, zero value otherwise.

### GetExternalSharingAllowedOk

`func (o *UpdateCalendarPoliciesRequest) GetExternalSharingAllowedOk() (*bool, bool)`

GetExternalSharingAllowedOk returns a tuple with the ExternalSharingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSharingAllowed

`func (o *UpdateCalendarPoliciesRequest) SetExternalSharingAllowed(v bool)`

SetExternalSharingAllowed sets ExternalSharingAllowed field to given value.

### HasExternalSharingAllowed

`func (o *UpdateCalendarPoliciesRequest) HasExternalSharingAllowed() bool`

HasExternalSharingAllowed returns a boolean if a field has been set.

### GetPublicLinksEnabled

`func (o *UpdateCalendarPoliciesRequest) GetPublicLinksEnabled() bool`

GetPublicLinksEnabled returns the PublicLinksEnabled field if non-nil, zero value otherwise.

### GetPublicLinksEnabledOk

`func (o *UpdateCalendarPoliciesRequest) GetPublicLinksEnabledOk() (*bool, bool)`

GetPublicLinksEnabledOk returns a tuple with the PublicLinksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicLinksEnabled

`func (o *UpdateCalendarPoliciesRequest) SetPublicLinksEnabled(v bool)`

SetPublicLinksEnabled sets PublicLinksEnabled field to given value.

### HasPublicLinksEnabled

`func (o *UpdateCalendarPoliciesRequest) HasPublicLinksEnabled() bool`

HasPublicLinksEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


