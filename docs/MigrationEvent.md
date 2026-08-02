# MigrationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MigrationId** | **string** |  | 
**TenantId** | **string** |  | 
**MailboxId** | Pointer to **string** |  | [optional] 
**EventType** | **string** |  | 
**Actor** | **string** |  | 
**OldStatus** | Pointer to **string** |  | [optional] 
**NewStatus** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewMigrationEvent

`func NewMigrationEvent(id string, migrationId string, tenantId string, eventType string, actor string, createdAt time.Time, ) *MigrationEvent`

NewMigrationEvent instantiates a new MigrationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationEventWithDefaults

`func NewMigrationEventWithDefaults() *MigrationEvent`

NewMigrationEventWithDefaults instantiates a new MigrationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MigrationEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MigrationEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MigrationEvent) SetId(v string)`

SetId sets Id field to given value.


### GetMigrationId

`func (o *MigrationEvent) GetMigrationId() string`

GetMigrationId returns the MigrationId field if non-nil, zero value otherwise.

### GetMigrationIdOk

`func (o *MigrationEvent) GetMigrationIdOk() (*string, bool)`

GetMigrationIdOk returns a tuple with the MigrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationId

`func (o *MigrationEvent) SetMigrationId(v string)`

SetMigrationId sets MigrationId field to given value.


### GetTenantId

`func (o *MigrationEvent) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MigrationEvent) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MigrationEvent) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetMailboxId

`func (o *MigrationEvent) GetMailboxId() string`

GetMailboxId returns the MailboxId field if non-nil, zero value otherwise.

### GetMailboxIdOk

`func (o *MigrationEvent) GetMailboxIdOk() (*string, bool)`

GetMailboxIdOk returns a tuple with the MailboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxId

`func (o *MigrationEvent) SetMailboxId(v string)`

SetMailboxId sets MailboxId field to given value.

### HasMailboxId

`func (o *MigrationEvent) HasMailboxId() bool`

HasMailboxId returns a boolean if a field has been set.

### GetEventType

`func (o *MigrationEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MigrationEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MigrationEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetActor

`func (o *MigrationEvent) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *MigrationEvent) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *MigrationEvent) SetActor(v string)`

SetActor sets Actor field to given value.


### GetOldStatus

`func (o *MigrationEvent) GetOldStatus() string`

GetOldStatus returns the OldStatus field if non-nil, zero value otherwise.

### GetOldStatusOk

`func (o *MigrationEvent) GetOldStatusOk() (*string, bool)`

GetOldStatusOk returns a tuple with the OldStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldStatus

`func (o *MigrationEvent) SetOldStatus(v string)`

SetOldStatus sets OldStatus field to given value.

### HasOldStatus

`func (o *MigrationEvent) HasOldStatus() bool`

HasOldStatus returns a boolean if a field has been set.

### GetNewStatus

`func (o *MigrationEvent) GetNewStatus() string`

GetNewStatus returns the NewStatus field if non-nil, zero value otherwise.

### GetNewStatusOk

`func (o *MigrationEvent) GetNewStatusOk() (*string, bool)`

GetNewStatusOk returns a tuple with the NewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatus

`func (o *MigrationEvent) SetNewStatus(v string)`

SetNewStatus sets NewStatus field to given value.

### HasNewStatus

`func (o *MigrationEvent) HasNewStatus() bool`

HasNewStatus returns a boolean if a field has been set.

### GetDetail

`func (o *MigrationEvent) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *MigrationEvent) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *MigrationEvent) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *MigrationEvent) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MigrationEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MigrationEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MigrationEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


