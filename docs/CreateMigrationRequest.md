# CreateMigrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CredentialId** | **string** |  | 
**SourceProvider** | **string** |  | 
**Settings** | Pointer to [**MigrationSettings**](MigrationSettings.md) |  | [optional] 

## Methods

### NewCreateMigrationRequest

`func NewCreateMigrationRequest(name string, credentialId string, sourceProvider string, ) *CreateMigrationRequest`

NewCreateMigrationRequest instantiates a new CreateMigrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMigrationRequestWithDefaults

`func NewCreateMigrationRequestWithDefaults() *CreateMigrationRequest`

NewCreateMigrationRequestWithDefaults instantiates a new CreateMigrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMigrationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMigrationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMigrationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCredentialId

`func (o *CreateMigrationRequest) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *CreateMigrationRequest) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *CreateMigrationRequest) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.


### GetSourceProvider

`func (o *CreateMigrationRequest) GetSourceProvider() string`

GetSourceProvider returns the SourceProvider field if non-nil, zero value otherwise.

### GetSourceProviderOk

`func (o *CreateMigrationRequest) GetSourceProviderOk() (*string, bool)`

GetSourceProviderOk returns a tuple with the SourceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProvider

`func (o *CreateMigrationRequest) SetSourceProvider(v string)`

SetSourceProvider sets SourceProvider field to given value.


### GetSettings

`func (o *CreateMigrationRequest) GetSettings() MigrationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateMigrationRequest) GetSettingsOk() (*MigrationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateMigrationRequest) SetSettings(v MigrationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateMigrationRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


