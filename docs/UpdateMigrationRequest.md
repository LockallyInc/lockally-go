# UpdateMigrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**MigrationSettings**](MigrationSettings.md) |  | [optional] 

## Methods

### NewUpdateMigrationRequest

`func NewUpdateMigrationRequest() *UpdateMigrationRequest`

NewUpdateMigrationRequest instantiates a new UpdateMigrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMigrationRequestWithDefaults

`func NewUpdateMigrationRequestWithDefaults() *UpdateMigrationRequest`

NewUpdateMigrationRequestWithDefaults instantiates a new UpdateMigrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMigrationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMigrationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMigrationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMigrationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateMigrationRequest) GetSettings() MigrationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateMigrationRequest) GetSettingsOk() (*MigrationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateMigrationRequest) SetSettings(v MigrationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateMigrationRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


