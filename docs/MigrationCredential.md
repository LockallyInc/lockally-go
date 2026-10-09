# MigrationCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Provider** | **string** |  | 
**EncryptionKeyId** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewMigrationCredential

`func NewMigrationCredential(id string, tenantId string, provider string, createdAt time.Time, updatedAt time.Time, ) *MigrationCredential`

NewMigrationCredential instantiates a new MigrationCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationCredentialWithDefaults

`func NewMigrationCredentialWithDefaults() *MigrationCredential`

NewMigrationCredentialWithDefaults instantiates a new MigrationCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MigrationCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MigrationCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MigrationCredential) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *MigrationCredential) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MigrationCredential) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MigrationCredential) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetProvider

`func (o *MigrationCredential) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *MigrationCredential) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *MigrationCredential) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetEncryptionKeyId

`func (o *MigrationCredential) GetEncryptionKeyId() string`

GetEncryptionKeyId returns the EncryptionKeyId field if non-nil, zero value otherwise.

### GetEncryptionKeyIdOk

`func (o *MigrationCredential) GetEncryptionKeyIdOk() (*string, bool)`

GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyId

`func (o *MigrationCredential) SetEncryptionKeyId(v string)`

SetEncryptionKeyId sets EncryptionKeyId field to given value.

### HasEncryptionKeyId

`func (o *MigrationCredential) HasEncryptionKeyId() bool`

HasEncryptionKeyId returns a boolean if a field has been set.

### GetLabel

`func (o *MigrationCredential) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MigrationCredential) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MigrationCredential) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MigrationCredential) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MigrationCredential) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MigrationCredential) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MigrationCredential) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MigrationCredential) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MigrationCredential) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MigrationCredential) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


