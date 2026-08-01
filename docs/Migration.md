# Migration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**CredentialId** | **string** |  | 
**Name** | **string** |  | 
**Status** | **string** |  | 
**SourceProvider** | **string** |  | 
**SourceSummary** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**MigrationSettings**](MigrationSettings.md) |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**MailboxCount** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewMigration

`func NewMigration(id string, tenantId string, credentialId string, name string, status string, sourceProvider string, mailboxCount int32, createdAt time.Time, updatedAt time.Time, ) *Migration`

NewMigration instantiates a new Migration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationWithDefaults

`func NewMigrationWithDefaults() *Migration`

NewMigrationWithDefaults instantiates a new Migration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Migration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Migration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Migration) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *Migration) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Migration) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Migration) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCredentialId

`func (o *Migration) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *Migration) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *Migration) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.


### GetName

`func (o *Migration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Migration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Migration) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Migration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Migration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Migration) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSourceProvider

`func (o *Migration) GetSourceProvider() string`

GetSourceProvider returns the SourceProvider field if non-nil, zero value otherwise.

### GetSourceProviderOk

`func (o *Migration) GetSourceProviderOk() (*string, bool)`

GetSourceProviderOk returns a tuple with the SourceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProvider

`func (o *Migration) SetSourceProvider(v string)`

SetSourceProvider sets SourceProvider field to given value.


### GetSourceSummary

`func (o *Migration) GetSourceSummary() string`

GetSourceSummary returns the SourceSummary field if non-nil, zero value otherwise.

### GetSourceSummaryOk

`func (o *Migration) GetSourceSummaryOk() (*string, bool)`

GetSourceSummaryOk returns a tuple with the SourceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSummary

`func (o *Migration) SetSourceSummary(v string)`

SetSourceSummary sets SourceSummary field to given value.

### HasSourceSummary

`func (o *Migration) HasSourceSummary() bool`

HasSourceSummary returns a boolean if a field has been set.

### GetSettings

`func (o *Migration) GetSettings() MigrationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Migration) GetSettingsOk() (*MigrationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Migration) SetSettings(v MigrationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Migration) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Migration) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Migration) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Migration) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Migration) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStartedAt

`func (o *Migration) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Migration) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Migration) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Migration) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *Migration) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Migration) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Migration) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Migration) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetMailboxCount

`func (o *Migration) GetMailboxCount() int32`

GetMailboxCount returns the MailboxCount field if non-nil, zero value otherwise.

### GetMailboxCountOk

`func (o *Migration) GetMailboxCountOk() (*int32, bool)`

GetMailboxCountOk returns a tuple with the MailboxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxCount

`func (o *Migration) SetMailboxCount(v int32)`

SetMailboxCount sets MailboxCount field to given value.


### GetCreatedAt

`func (o *Migration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Migration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Migration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Migration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Migration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Migration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


