# MigrationMailbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MigrationId** | **string** |  | 
**TenantId** | **string** |  | 
**SourceEmail** | **string** |  | 
**DestEmail** | Pointer to **string** |  | [optional] 
**DestMailboxId** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**SourceMessageCount** | Pointer to **int32** |  | [optional] 
**SyncedMessageCount** | **int32** |  | 
**FailedMessageCount** | **int32** |  | 
**SourceSizeBytes** | Pointer to **int64** |  | [optional] 
**SyncedSizeBytes** | Pointer to **int64** |  | [optional] 
**LastSyncedUid** | Pointer to **string** |  | [optional] 
**LastSyncedAt** | Pointer to **time.Time** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewMigrationMailbox

`func NewMigrationMailbox(id string, migrationId string, tenantId string, sourceEmail string, status string, syncedMessageCount int32, failedMessageCount int32, createdAt time.Time, updatedAt time.Time, ) *MigrationMailbox`

NewMigrationMailbox instantiates a new MigrationMailbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationMailboxWithDefaults

`func NewMigrationMailboxWithDefaults() *MigrationMailbox`

NewMigrationMailboxWithDefaults instantiates a new MigrationMailbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MigrationMailbox) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MigrationMailbox) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MigrationMailbox) SetId(v string)`

SetId sets Id field to given value.


### GetMigrationId

`func (o *MigrationMailbox) GetMigrationId() string`

GetMigrationId returns the MigrationId field if non-nil, zero value otherwise.

### GetMigrationIdOk

`func (o *MigrationMailbox) GetMigrationIdOk() (*string, bool)`

GetMigrationIdOk returns a tuple with the MigrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationId

`func (o *MigrationMailbox) SetMigrationId(v string)`

SetMigrationId sets MigrationId field to given value.


### GetTenantId

`func (o *MigrationMailbox) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MigrationMailbox) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MigrationMailbox) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSourceEmail

`func (o *MigrationMailbox) GetSourceEmail() string`

GetSourceEmail returns the SourceEmail field if non-nil, zero value otherwise.

### GetSourceEmailOk

`func (o *MigrationMailbox) GetSourceEmailOk() (*string, bool)`

GetSourceEmailOk returns a tuple with the SourceEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEmail

`func (o *MigrationMailbox) SetSourceEmail(v string)`

SetSourceEmail sets SourceEmail field to given value.


### GetDestEmail

`func (o *MigrationMailbox) GetDestEmail() string`

GetDestEmail returns the DestEmail field if non-nil, zero value otherwise.

### GetDestEmailOk

`func (o *MigrationMailbox) GetDestEmailOk() (*string, bool)`

GetDestEmailOk returns a tuple with the DestEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestEmail

`func (o *MigrationMailbox) SetDestEmail(v string)`

SetDestEmail sets DestEmail field to given value.

### HasDestEmail

`func (o *MigrationMailbox) HasDestEmail() bool`

HasDestEmail returns a boolean if a field has been set.

### GetDestMailboxId

`func (o *MigrationMailbox) GetDestMailboxId() string`

GetDestMailboxId returns the DestMailboxId field if non-nil, zero value otherwise.

### GetDestMailboxIdOk

`func (o *MigrationMailbox) GetDestMailboxIdOk() (*string, bool)`

GetDestMailboxIdOk returns a tuple with the DestMailboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestMailboxId

`func (o *MigrationMailbox) SetDestMailboxId(v string)`

SetDestMailboxId sets DestMailboxId field to given value.

### HasDestMailboxId

`func (o *MigrationMailbox) HasDestMailboxId() bool`

HasDestMailboxId returns a boolean if a field has been set.

### GetStatus

`func (o *MigrationMailbox) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MigrationMailbox) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MigrationMailbox) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSourceMessageCount

`func (o *MigrationMailbox) GetSourceMessageCount() int32`

GetSourceMessageCount returns the SourceMessageCount field if non-nil, zero value otherwise.

### GetSourceMessageCountOk

`func (o *MigrationMailbox) GetSourceMessageCountOk() (*int32, bool)`

GetSourceMessageCountOk returns a tuple with the SourceMessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMessageCount

`func (o *MigrationMailbox) SetSourceMessageCount(v int32)`

SetSourceMessageCount sets SourceMessageCount field to given value.

### HasSourceMessageCount

`func (o *MigrationMailbox) HasSourceMessageCount() bool`

HasSourceMessageCount returns a boolean if a field has been set.

### GetSyncedMessageCount

`func (o *MigrationMailbox) GetSyncedMessageCount() int32`

GetSyncedMessageCount returns the SyncedMessageCount field if non-nil, zero value otherwise.

### GetSyncedMessageCountOk

`func (o *MigrationMailbox) GetSyncedMessageCountOk() (*int32, bool)`

GetSyncedMessageCountOk returns a tuple with the SyncedMessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedMessageCount

`func (o *MigrationMailbox) SetSyncedMessageCount(v int32)`

SetSyncedMessageCount sets SyncedMessageCount field to given value.


### GetFailedMessageCount

`func (o *MigrationMailbox) GetFailedMessageCount() int32`

GetFailedMessageCount returns the FailedMessageCount field if non-nil, zero value otherwise.

### GetFailedMessageCountOk

`func (o *MigrationMailbox) GetFailedMessageCountOk() (*int32, bool)`

GetFailedMessageCountOk returns a tuple with the FailedMessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMessageCount

`func (o *MigrationMailbox) SetFailedMessageCount(v int32)`

SetFailedMessageCount sets FailedMessageCount field to given value.


### GetSourceSizeBytes

`func (o *MigrationMailbox) GetSourceSizeBytes() int64`

GetSourceSizeBytes returns the SourceSizeBytes field if non-nil, zero value otherwise.

### GetSourceSizeBytesOk

`func (o *MigrationMailbox) GetSourceSizeBytesOk() (*int64, bool)`

GetSourceSizeBytesOk returns a tuple with the SourceSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSizeBytes

`func (o *MigrationMailbox) SetSourceSizeBytes(v int64)`

SetSourceSizeBytes sets SourceSizeBytes field to given value.

### HasSourceSizeBytes

`func (o *MigrationMailbox) HasSourceSizeBytes() bool`

HasSourceSizeBytes returns a boolean if a field has been set.

### GetSyncedSizeBytes

`func (o *MigrationMailbox) GetSyncedSizeBytes() int64`

GetSyncedSizeBytes returns the SyncedSizeBytes field if non-nil, zero value otherwise.

### GetSyncedSizeBytesOk

`func (o *MigrationMailbox) GetSyncedSizeBytesOk() (*int64, bool)`

GetSyncedSizeBytesOk returns a tuple with the SyncedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedSizeBytes

`func (o *MigrationMailbox) SetSyncedSizeBytes(v int64)`

SetSyncedSizeBytes sets SyncedSizeBytes field to given value.

### HasSyncedSizeBytes

`func (o *MigrationMailbox) HasSyncedSizeBytes() bool`

HasSyncedSizeBytes returns a boolean if a field has been set.

### GetLastSyncedUid

`func (o *MigrationMailbox) GetLastSyncedUid() string`

GetLastSyncedUid returns the LastSyncedUid field if non-nil, zero value otherwise.

### GetLastSyncedUidOk

`func (o *MigrationMailbox) GetLastSyncedUidOk() (*string, bool)`

GetLastSyncedUidOk returns a tuple with the LastSyncedUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedUid

`func (o *MigrationMailbox) SetLastSyncedUid(v string)`

SetLastSyncedUid sets LastSyncedUid field to given value.

### HasLastSyncedUid

`func (o *MigrationMailbox) HasLastSyncedUid() bool`

HasLastSyncedUid returns a boolean if a field has been set.

### GetLastSyncedAt

`func (o *MigrationMailbox) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *MigrationMailbox) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *MigrationMailbox) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *MigrationMailbox) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### GetErrorMessage

`func (o *MigrationMailbox) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *MigrationMailbox) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *MigrationMailbox) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *MigrationMailbox) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStartedAt

`func (o *MigrationMailbox) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MigrationMailbox) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MigrationMailbox) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MigrationMailbox) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *MigrationMailbox) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *MigrationMailbox) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *MigrationMailbox) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *MigrationMailbox) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MigrationMailbox) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MigrationMailbox) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MigrationMailbox) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MigrationMailbox) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MigrationMailbox) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MigrationMailbox) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


