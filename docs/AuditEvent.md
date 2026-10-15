# AuditEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAuditEvent

`func NewAuditEvent() *AuditEvent`

NewAuditEvent instantiates a new AuditEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventWithDefaults

`func NewAuditEventWithDefaults() *AuditEvent`

NewAuditEventWithDefaults instantiates a new AuditEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AuditEvent) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuditEvent) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuditEvent) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuditEvent) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEventType

`func (o *AuditEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AuditEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AuditEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *AuditEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetDetail

`func (o *AuditEvent) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AuditEvent) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AuditEvent) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AuditEvent) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetIp

`func (o *AuditEvent) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AuditEvent) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AuditEvent) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AuditEvent) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetTime

`func (o *AuditEvent) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AuditEvent) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AuditEvent) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *AuditEvent) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


