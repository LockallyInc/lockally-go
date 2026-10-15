# CalendarPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**MaxMeetingDurationMins** | Pointer to **int32** |  | [optional] 
**WorkingHoursStart** | Pointer to **string** |  | [optional] 
**WorkingHoursEnd** | Pointer to **string** |  | [optional] 
**BookingWindowDays** | Pointer to **int32** |  | [optional] 
**RecurringMeetingLimit** | Pointer to **int32** |  | [optional] 
**ResourceApprovalMode** | Pointer to **string** |  | [optional] 
**ExternalInvitesAllowed** | Pointer to **bool** |  | [optional] 
**ExternalSharingAllowed** | Pointer to **bool** |  | [optional] 
**PublicLinksEnabled** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCalendarPolicies

`func NewCalendarPolicies(tenantId string, ) *CalendarPolicies`

NewCalendarPolicies instantiates a new CalendarPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarPoliciesWithDefaults

`func NewCalendarPoliciesWithDefaults() *CalendarPolicies`

NewCalendarPoliciesWithDefaults instantiates a new CalendarPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CalendarPolicies) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CalendarPolicies) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CalendarPolicies) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetMaxMeetingDurationMins

`func (o *CalendarPolicies) GetMaxMeetingDurationMins() int32`

GetMaxMeetingDurationMins returns the MaxMeetingDurationMins field if non-nil, zero value otherwise.

### GetMaxMeetingDurationMinsOk

`func (o *CalendarPolicies) GetMaxMeetingDurationMinsOk() (*int32, bool)`

GetMaxMeetingDurationMinsOk returns a tuple with the MaxMeetingDurationMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMeetingDurationMins

`func (o *CalendarPolicies) SetMaxMeetingDurationMins(v int32)`

SetMaxMeetingDurationMins sets MaxMeetingDurationMins field to given value.

### HasMaxMeetingDurationMins

`func (o *CalendarPolicies) HasMaxMeetingDurationMins() bool`

HasMaxMeetingDurationMins returns a boolean if a field has been set.

### GetWorkingHoursStart

`func (o *CalendarPolicies) GetWorkingHoursStart() string`

GetWorkingHoursStart returns the WorkingHoursStart field if non-nil, zero value otherwise.

### GetWorkingHoursStartOk

`func (o *CalendarPolicies) GetWorkingHoursStartOk() (*string, bool)`

GetWorkingHoursStartOk returns a tuple with the WorkingHoursStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHoursStart

`func (o *CalendarPolicies) SetWorkingHoursStart(v string)`

SetWorkingHoursStart sets WorkingHoursStart field to given value.

### HasWorkingHoursStart

`func (o *CalendarPolicies) HasWorkingHoursStart() bool`

HasWorkingHoursStart returns a boolean if a field has been set.

### GetWorkingHoursEnd

`func (o *CalendarPolicies) GetWorkingHoursEnd() string`

GetWorkingHoursEnd returns the WorkingHoursEnd field if non-nil, zero value otherwise.

### GetWorkingHoursEndOk

`func (o *CalendarPolicies) GetWorkingHoursEndOk() (*string, bool)`

GetWorkingHoursEndOk returns a tuple with the WorkingHoursEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHoursEnd

`func (o *CalendarPolicies) SetWorkingHoursEnd(v string)`

SetWorkingHoursEnd sets WorkingHoursEnd field to given value.

### HasWorkingHoursEnd

`func (o *CalendarPolicies) HasWorkingHoursEnd() bool`

HasWorkingHoursEnd returns a boolean if a field has been set.

### GetBookingWindowDays

`func (o *CalendarPolicies) GetBookingWindowDays() int32`

GetBookingWindowDays returns the BookingWindowDays field if non-nil, zero value otherwise.

### GetBookingWindowDaysOk

`func (o *CalendarPolicies) GetBookingWindowDaysOk() (*int32, bool)`

GetBookingWindowDaysOk returns a tuple with the BookingWindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingWindowDays

`func (o *CalendarPolicies) SetBookingWindowDays(v int32)`

SetBookingWindowDays sets BookingWindowDays field to given value.

### HasBookingWindowDays

`func (o *CalendarPolicies) HasBookingWindowDays() bool`

HasBookingWindowDays returns a boolean if a field has been set.

### GetRecurringMeetingLimit

`func (o *CalendarPolicies) GetRecurringMeetingLimit() int32`

GetRecurringMeetingLimit returns the RecurringMeetingLimit field if non-nil, zero value otherwise.

### GetRecurringMeetingLimitOk

`func (o *CalendarPolicies) GetRecurringMeetingLimitOk() (*int32, bool)`

GetRecurringMeetingLimitOk returns a tuple with the RecurringMeetingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringMeetingLimit

`func (o *CalendarPolicies) SetRecurringMeetingLimit(v int32)`

SetRecurringMeetingLimit sets RecurringMeetingLimit field to given value.

### HasRecurringMeetingLimit

`func (o *CalendarPolicies) HasRecurringMeetingLimit() bool`

HasRecurringMeetingLimit returns a boolean if a field has been set.

### GetResourceApprovalMode

`func (o *CalendarPolicies) GetResourceApprovalMode() string`

GetResourceApprovalMode returns the ResourceApprovalMode field if non-nil, zero value otherwise.

### GetResourceApprovalModeOk

`func (o *CalendarPolicies) GetResourceApprovalModeOk() (*string, bool)`

GetResourceApprovalModeOk returns a tuple with the ResourceApprovalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceApprovalMode

`func (o *CalendarPolicies) SetResourceApprovalMode(v string)`

SetResourceApprovalMode sets ResourceApprovalMode field to given value.

### HasResourceApprovalMode

`func (o *CalendarPolicies) HasResourceApprovalMode() bool`

HasResourceApprovalMode returns a boolean if a field has been set.

### GetExternalInvitesAllowed

`func (o *CalendarPolicies) GetExternalInvitesAllowed() bool`

GetExternalInvitesAllowed returns the ExternalInvitesAllowed field if non-nil, zero value otherwise.

### GetExternalInvitesAllowedOk

`func (o *CalendarPolicies) GetExternalInvitesAllowedOk() (*bool, bool)`

GetExternalInvitesAllowedOk returns a tuple with the ExternalInvitesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInvitesAllowed

`func (o *CalendarPolicies) SetExternalInvitesAllowed(v bool)`

SetExternalInvitesAllowed sets ExternalInvitesAllowed field to given value.

### HasExternalInvitesAllowed

`func (o *CalendarPolicies) HasExternalInvitesAllowed() bool`

HasExternalInvitesAllowed returns a boolean if a field has been set.

### GetExternalSharingAllowed

`func (o *CalendarPolicies) GetExternalSharingAllowed() bool`

GetExternalSharingAllowed returns the ExternalSharingAllowed field if non-nil, zero value otherwise.

### GetExternalSharingAllowedOk

`func (o *CalendarPolicies) GetExternalSharingAllowedOk() (*bool, bool)`

GetExternalSharingAllowedOk returns a tuple with the ExternalSharingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSharingAllowed

`func (o *CalendarPolicies) SetExternalSharingAllowed(v bool)`

SetExternalSharingAllowed sets ExternalSharingAllowed field to given value.

### HasExternalSharingAllowed

`func (o *CalendarPolicies) HasExternalSharingAllowed() bool`

HasExternalSharingAllowed returns a boolean if a field has been set.

### GetPublicLinksEnabled

`func (o *CalendarPolicies) GetPublicLinksEnabled() bool`

GetPublicLinksEnabled returns the PublicLinksEnabled field if non-nil, zero value otherwise.

### GetPublicLinksEnabledOk

`func (o *CalendarPolicies) GetPublicLinksEnabledOk() (*bool, bool)`

GetPublicLinksEnabledOk returns a tuple with the PublicLinksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicLinksEnabled

`func (o *CalendarPolicies) SetPublicLinksEnabled(v bool)`

SetPublicLinksEnabled sets PublicLinksEnabled field to given value.

### HasPublicLinksEnabled

`func (o *CalendarPolicies) HasPublicLinksEnabled() bool`

HasPublicLinksEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CalendarPolicies) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CalendarPolicies) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CalendarPolicies) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CalendarPolicies) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CalendarPolicies) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CalendarPolicies) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CalendarPolicies) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CalendarPolicies) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


