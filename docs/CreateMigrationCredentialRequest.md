# CreateMigrationCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**Credentials** | [**CreateMigrationCredentialRequestCredentials**](CreateMigrationCredentialRequestCredentials.md) |  | 

## Methods

### NewCreateMigrationCredentialRequest

`func NewCreateMigrationCredentialRequest(provider string, credentials CreateMigrationCredentialRequestCredentials, ) *CreateMigrationCredentialRequest`

NewCreateMigrationCredentialRequest instantiates a new CreateMigrationCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMigrationCredentialRequestWithDefaults

`func NewCreateMigrationCredentialRequestWithDefaults() *CreateMigrationCredentialRequest`

NewCreateMigrationCredentialRequestWithDefaults instantiates a new CreateMigrationCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *CreateMigrationCredentialRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateMigrationCredentialRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateMigrationCredentialRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetLabel

`func (o *CreateMigrationCredentialRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateMigrationCredentialRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateMigrationCredentialRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateMigrationCredentialRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCredentials

`func (o *CreateMigrationCredentialRequest) GetCredentials() CreateMigrationCredentialRequestCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CreateMigrationCredentialRequest) GetCredentialsOk() (*CreateMigrationCredentialRequestCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CreateMigrationCredentialRequest) SetCredentials(v CreateMigrationCredentialRequestCredentials)`

SetCredentials sets Credentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


