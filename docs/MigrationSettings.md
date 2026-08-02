# MigrationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxConcurrentMailboxes** | Pointer to **int32** |  | [optional] 
**MaxConcurrentMessages** | Pointer to **int32** |  | [optional] 
**SourceRateLimit** | Pointer to **int32** |  | [optional] 
**BatchSize** | Pointer to **int32** |  | [optional] 
**SkipFolders** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMigrationSettings

`func NewMigrationSettings() *MigrationSettings`

NewMigrationSettings instantiates a new MigrationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationSettingsWithDefaults

`func NewMigrationSettingsWithDefaults() *MigrationSettings`

NewMigrationSettingsWithDefaults instantiates a new MigrationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxConcurrentMailboxes

`func (o *MigrationSettings) GetMaxConcurrentMailboxes() int32`

GetMaxConcurrentMailboxes returns the MaxConcurrentMailboxes field if non-nil, zero value otherwise.

### GetMaxConcurrentMailboxesOk

`func (o *MigrationSettings) GetMaxConcurrentMailboxesOk() (*int32, bool)`

GetMaxConcurrentMailboxesOk returns a tuple with the MaxConcurrentMailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentMailboxes

`func (o *MigrationSettings) SetMaxConcurrentMailboxes(v int32)`

SetMaxConcurrentMailboxes sets MaxConcurrentMailboxes field to given value.

### HasMaxConcurrentMailboxes

`func (o *MigrationSettings) HasMaxConcurrentMailboxes() bool`

HasMaxConcurrentMailboxes returns a boolean if a field has been set.

### GetMaxConcurrentMessages

`func (o *MigrationSettings) GetMaxConcurrentMessages() int32`

GetMaxConcurrentMessages returns the MaxConcurrentMessages field if non-nil, zero value otherwise.

### GetMaxConcurrentMessagesOk

`func (o *MigrationSettings) GetMaxConcurrentMessagesOk() (*int32, bool)`

GetMaxConcurrentMessagesOk returns a tuple with the MaxConcurrentMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentMessages

`func (o *MigrationSettings) SetMaxConcurrentMessages(v int32)`

SetMaxConcurrentMessages sets MaxConcurrentMessages field to given value.

### HasMaxConcurrentMessages

`func (o *MigrationSettings) HasMaxConcurrentMessages() bool`

HasMaxConcurrentMessages returns a boolean if a field has been set.

### GetSourceRateLimit

`func (o *MigrationSettings) GetSourceRateLimit() int32`

GetSourceRateLimit returns the SourceRateLimit field if non-nil, zero value otherwise.

### GetSourceRateLimitOk

`func (o *MigrationSettings) GetSourceRateLimitOk() (*int32, bool)`

GetSourceRateLimitOk returns a tuple with the SourceRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRateLimit

`func (o *MigrationSettings) SetSourceRateLimit(v int32)`

SetSourceRateLimit sets SourceRateLimit field to given value.

### HasSourceRateLimit

`func (o *MigrationSettings) HasSourceRateLimit() bool`

HasSourceRateLimit returns a boolean if a field has been set.

### GetBatchSize

`func (o *MigrationSettings) GetBatchSize() int32`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *MigrationSettings) GetBatchSizeOk() (*int32, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *MigrationSettings) SetBatchSize(v int32)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *MigrationSettings) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetSkipFolders

`func (o *MigrationSettings) GetSkipFolders() []string`

GetSkipFolders returns the SkipFolders field if non-nil, zero value otherwise.

### GetSkipFoldersOk

`func (o *MigrationSettings) GetSkipFoldersOk() (*[]string, bool)`

GetSkipFoldersOk returns a tuple with the SkipFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipFolders

`func (o *MigrationSettings) SetSkipFolders(v []string)`

SetSkipFolders sets SkipFolders field to given value.

### HasSkipFolders

`func (o *MigrationSettings) HasSkipFolders() bool`

HasSkipFolders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


