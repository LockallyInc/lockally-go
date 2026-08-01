# UserEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**EventAt** | Pointer to **time.Time** |  | [optional] 
**DaysInactive** | Pointer to **int32** |  | [optional] 

## Methods

### NewUserEvent

`func NewUserEvent() *UserEvent`

NewUserEvent instantiates a new UserEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEventWithDefaults

`func NewUserEventWithDefaults() *UserEvent`

NewUserEventWithDefaults instantiates a new UserEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserEvent) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserEvent) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserEvent) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserEvent) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEventAt

`func (o *UserEvent) GetEventAt() time.Time`

GetEventAt returns the EventAt field if non-nil, zero value otherwise.

### GetEventAtOk

`func (o *UserEvent) GetEventAtOk() (*time.Time, bool)`

GetEventAtOk returns a tuple with the EventAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAt

`func (o *UserEvent) SetEventAt(v time.Time)`

SetEventAt sets EventAt field to given value.

### HasEventAt

`func (o *UserEvent) HasEventAt() bool`

HasEventAt returns a boolean if a field has been set.

### GetDaysInactive

`func (o *UserEvent) GetDaysInactive() int32`

GetDaysInactive returns the DaysInactive field if non-nil, zero value otherwise.

### GetDaysInactiveOk

`func (o *UserEvent) GetDaysInactiveOk() (*int32, bool)`

GetDaysInactiveOk returns a tuple with the DaysInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInactive

`func (o *UserEvent) SetDaysInactive(v int32)`

SetDaysInactive sets DaysInactive field to given value.

### HasDaysInactive

`func (o *UserEvent) HasDaysInactive() bool`

HasDaysInactive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


