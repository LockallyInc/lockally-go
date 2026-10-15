# VacationResponder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailboxEmail** | **string** |  | 
**Enabled** | **bool** |  | 
**Params** | [**VacationParams**](VacationParams.md) |  | 
**Script** | **string** | Pre-rendered Sieve script (RFC 5230). | 
**SyncedAt** | Pointer to **time.Time** | Null &#x3D; stored on lockally but not yet pushed to the mail server. | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewVacationResponder

`func NewVacationResponder(mailboxEmail string, enabled bool, params VacationParams, script string, createdAt time.Time, updatedAt time.Time, ) *VacationResponder`

NewVacationResponder instantiates a new VacationResponder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacationResponderWithDefaults

`func NewVacationResponderWithDefaults() *VacationResponder`

NewVacationResponderWithDefaults instantiates a new VacationResponder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxEmail

`func (o *VacationResponder) GetMailboxEmail() string`

GetMailboxEmail returns the MailboxEmail field if non-nil, zero value otherwise.

### GetMailboxEmailOk

`func (o *VacationResponder) GetMailboxEmailOk() (*string, bool)`

GetMailboxEmailOk returns a tuple with the MailboxEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxEmail

`func (o *VacationResponder) SetMailboxEmail(v string)`

SetMailboxEmail sets MailboxEmail field to given value.


### GetEnabled

`func (o *VacationResponder) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VacationResponder) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VacationResponder) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetParams

`func (o *VacationResponder) GetParams() VacationParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *VacationResponder) GetParamsOk() (*VacationParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *VacationResponder) SetParams(v VacationParams)`

SetParams sets Params field to given value.


### GetScript

`func (o *VacationResponder) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *VacationResponder) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *VacationResponder) SetScript(v string)`

SetScript sets Script field to given value.


### GetSyncedAt

`func (o *VacationResponder) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *VacationResponder) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *VacationResponder) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *VacationResponder) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VacationResponder) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacationResponder) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacationResponder) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VacationResponder) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VacationResponder) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VacationResponder) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


