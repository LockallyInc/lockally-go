# MigrationProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MigrationId** | **string** |  | 
**Status** | **string** |  | 
**TotalMailboxes** | **int32** |  | 
**CompletedMailboxes** | **int32** |  | 
**FailedMailboxes** | **int32** |  | 
**TotalMessages** | **int32** |  | 
**SyncedMessages** | **int32** |  | 
**FailedMessages** | **int32** |  | 
**PercentComplete** | **float32** |  | 
**Mailboxes** | [**[]MigrationProgressMailboxesInner**](MigrationProgressMailboxesInner.md) |  | 

## Methods

### NewMigrationProgress

`func NewMigrationProgress(migrationId string, status string, totalMailboxes int32, completedMailboxes int32, failedMailboxes int32, totalMessages int32, syncedMessages int32, failedMessages int32, percentComplete float32, mailboxes []MigrationProgressMailboxesInner, ) *MigrationProgress`

NewMigrationProgress instantiates a new MigrationProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationProgressWithDefaults

`func NewMigrationProgressWithDefaults() *MigrationProgress`

NewMigrationProgressWithDefaults instantiates a new MigrationProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigrationId

`func (o *MigrationProgress) GetMigrationId() string`

GetMigrationId returns the MigrationId field if non-nil, zero value otherwise.

### GetMigrationIdOk

`func (o *MigrationProgress) GetMigrationIdOk() (*string, bool)`

GetMigrationIdOk returns a tuple with the MigrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationId

`func (o *MigrationProgress) SetMigrationId(v string)`

SetMigrationId sets MigrationId field to given value.


### GetStatus

`func (o *MigrationProgress) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MigrationProgress) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MigrationProgress) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalMailboxes

`func (o *MigrationProgress) GetTotalMailboxes() int32`

GetTotalMailboxes returns the TotalMailboxes field if non-nil, zero value otherwise.

### GetTotalMailboxesOk

`func (o *MigrationProgress) GetTotalMailboxesOk() (*int32, bool)`

GetTotalMailboxesOk returns a tuple with the TotalMailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMailboxes

`func (o *MigrationProgress) SetTotalMailboxes(v int32)`

SetTotalMailboxes sets TotalMailboxes field to given value.


### GetCompletedMailboxes

`func (o *MigrationProgress) GetCompletedMailboxes() int32`

GetCompletedMailboxes returns the CompletedMailboxes field if non-nil, zero value otherwise.

### GetCompletedMailboxesOk

`func (o *MigrationProgress) GetCompletedMailboxesOk() (*int32, bool)`

GetCompletedMailboxesOk returns a tuple with the CompletedMailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedMailboxes

`func (o *MigrationProgress) SetCompletedMailboxes(v int32)`

SetCompletedMailboxes sets CompletedMailboxes field to given value.


### GetFailedMailboxes

`func (o *MigrationProgress) GetFailedMailboxes() int32`

GetFailedMailboxes returns the FailedMailboxes field if non-nil, zero value otherwise.

### GetFailedMailboxesOk

`func (o *MigrationProgress) GetFailedMailboxesOk() (*int32, bool)`

GetFailedMailboxesOk returns a tuple with the FailedMailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMailboxes

`func (o *MigrationProgress) SetFailedMailboxes(v int32)`

SetFailedMailboxes sets FailedMailboxes field to given value.


### GetTotalMessages

`func (o *MigrationProgress) GetTotalMessages() int32`

GetTotalMessages returns the TotalMessages field if non-nil, zero value otherwise.

### GetTotalMessagesOk

`func (o *MigrationProgress) GetTotalMessagesOk() (*int32, bool)`

GetTotalMessagesOk returns a tuple with the TotalMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMessages

`func (o *MigrationProgress) SetTotalMessages(v int32)`

SetTotalMessages sets TotalMessages field to given value.


### GetSyncedMessages

`func (o *MigrationProgress) GetSyncedMessages() int32`

GetSyncedMessages returns the SyncedMessages field if non-nil, zero value otherwise.

### GetSyncedMessagesOk

`func (o *MigrationProgress) GetSyncedMessagesOk() (*int32, bool)`

GetSyncedMessagesOk returns a tuple with the SyncedMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedMessages

`func (o *MigrationProgress) SetSyncedMessages(v int32)`

SetSyncedMessages sets SyncedMessages field to given value.


### GetFailedMessages

`func (o *MigrationProgress) GetFailedMessages() int32`

GetFailedMessages returns the FailedMessages field if non-nil, zero value otherwise.

### GetFailedMessagesOk

`func (o *MigrationProgress) GetFailedMessagesOk() (*int32, bool)`

GetFailedMessagesOk returns a tuple with the FailedMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMessages

`func (o *MigrationProgress) SetFailedMessages(v int32)`

SetFailedMessages sets FailedMessages field to given value.


### GetPercentComplete

`func (o *MigrationProgress) GetPercentComplete() float32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *MigrationProgress) GetPercentCompleteOk() (*float32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *MigrationProgress) SetPercentComplete(v float32)`

SetPercentComplete sets PercentComplete field to given value.


### GetMailboxes

`func (o *MigrationProgress) GetMailboxes() []MigrationProgressMailboxesInner`

GetMailboxes returns the Mailboxes field if non-nil, zero value otherwise.

### GetMailboxesOk

`func (o *MigrationProgress) GetMailboxesOk() (*[]MigrationProgressMailboxesInner, bool)`

GetMailboxesOk returns a tuple with the Mailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxes

`func (o *MigrationProgress) SetMailboxes(v []MigrationProgressMailboxesInner)`

SetMailboxes sets Mailboxes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


