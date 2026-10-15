# MigrationProgressMailboxesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceEmail** | **string** |  | 
**DestEmail** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**SourceMessageCount** | Pointer to **int32** |  | [optional] 
**SyncedMessageCount** | **int32** |  | 
**FailedMessageCount** | **int32** |  | 
**PercentComplete** | **float32** |  | 

## Methods

### NewMigrationProgressMailboxesInner

`func NewMigrationProgressMailboxesInner(sourceEmail string, status string, syncedMessageCount int32, failedMessageCount int32, percentComplete float32, ) *MigrationProgressMailboxesInner`

NewMigrationProgressMailboxesInner instantiates a new MigrationProgressMailboxesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationProgressMailboxesInnerWithDefaults

`func NewMigrationProgressMailboxesInnerWithDefaults() *MigrationProgressMailboxesInner`

NewMigrationProgressMailboxesInnerWithDefaults instantiates a new MigrationProgressMailboxesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceEmail

`func (o *MigrationProgressMailboxesInner) GetSourceEmail() string`

GetSourceEmail returns the SourceEmail field if non-nil, zero value otherwise.

### GetSourceEmailOk

`func (o *MigrationProgressMailboxesInner) GetSourceEmailOk() (*string, bool)`

GetSourceEmailOk returns a tuple with the SourceEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEmail

`func (o *MigrationProgressMailboxesInner) SetSourceEmail(v string)`

SetSourceEmail sets SourceEmail field to given value.


### GetDestEmail

`func (o *MigrationProgressMailboxesInner) GetDestEmail() string`

GetDestEmail returns the DestEmail field if non-nil, zero value otherwise.

### GetDestEmailOk

`func (o *MigrationProgressMailboxesInner) GetDestEmailOk() (*string, bool)`

GetDestEmailOk returns a tuple with the DestEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestEmail

`func (o *MigrationProgressMailboxesInner) SetDestEmail(v string)`

SetDestEmail sets DestEmail field to given value.

### HasDestEmail

`func (o *MigrationProgressMailboxesInner) HasDestEmail() bool`

HasDestEmail returns a boolean if a field has been set.

### GetStatus

`func (o *MigrationProgressMailboxesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MigrationProgressMailboxesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MigrationProgressMailboxesInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSourceMessageCount

`func (o *MigrationProgressMailboxesInner) GetSourceMessageCount() int32`

GetSourceMessageCount returns the SourceMessageCount field if non-nil, zero value otherwise.

### GetSourceMessageCountOk

`func (o *MigrationProgressMailboxesInner) GetSourceMessageCountOk() (*int32, bool)`

GetSourceMessageCountOk returns a tuple with the SourceMessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMessageCount

`func (o *MigrationProgressMailboxesInner) SetSourceMessageCount(v int32)`

SetSourceMessageCount sets SourceMessageCount field to given value.

### HasSourceMessageCount

`func (o *MigrationProgressMailboxesInner) HasSourceMessageCount() bool`

HasSourceMessageCount returns a boolean if a field has been set.

### GetSyncedMessageCount

`func (o *MigrationProgressMailboxesInner) GetSyncedMessageCount() int32`

GetSyncedMessageCount returns the SyncedMessageCount field if non-nil, zero value otherwise.

### GetSyncedMessageCountOk

`func (o *MigrationProgressMailboxesInner) GetSyncedMessageCountOk() (*int32, bool)`

GetSyncedMessageCountOk returns a tuple with the SyncedMessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedMessageCount

`func (o *MigrationProgressMailboxesInner) SetSyncedMessageCount(v int32)`

SetSyncedMessageCount sets SyncedMessageCount field to given value.


### GetFailedMessageCount

`func (o *MigrationProgressMailboxesInner) GetFailedMessageCount() int32`

GetFailedMessageCount returns the FailedMessageCount field if non-nil, zero value otherwise.

### GetFailedMessageCountOk

`func (o *MigrationProgressMailboxesInner) GetFailedMessageCountOk() (*int32, bool)`

GetFailedMessageCountOk returns a tuple with the FailedMessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMessageCount

`func (o *MigrationProgressMailboxesInner) SetFailedMessageCount(v int32)`

SetFailedMessageCount sets FailedMessageCount field to given value.


### GetPercentComplete

`func (o *MigrationProgressMailboxesInner) GetPercentComplete() float32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *MigrationProgressMailboxesInner) GetPercentCompleteOk() (*float32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *MigrationProgressMailboxesInner) SetPercentComplete(v float32)`

SetPercentComplete sets PercentComplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


