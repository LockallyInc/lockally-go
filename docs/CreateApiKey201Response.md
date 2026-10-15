# CreateApiKey201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Prefix** | **string** | 8-char public prefix; safe to store and display. | 
**Scopes** | **[]string** |  | 
**Label** | **string** |  | 
**LastUsedAt** | Pointer to **time.Time** |  | [optional] 
**RevokedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Secret** | **string** | The full &#x60;lk_live_&lt;prefix&gt;_&lt;secret&gt;&#x60; token. Shown ONCE. | 

## Methods

### NewCreateApiKey201Response

`func NewCreateApiKey201Response(id string, tenantId string, prefix string, scopes []string, label string, createdAt time.Time, secret string, ) *CreateApiKey201Response`

NewCreateApiKey201Response instantiates a new CreateApiKey201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKey201ResponseWithDefaults

`func NewCreateApiKey201ResponseWithDefaults() *CreateApiKey201Response`

NewCreateApiKey201ResponseWithDefaults instantiates a new CreateApiKey201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateApiKey201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateApiKey201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateApiKey201Response) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *CreateApiKey201Response) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateApiKey201Response) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateApiKey201Response) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPrefix

`func (o *CreateApiKey201Response) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CreateApiKey201Response) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CreateApiKey201Response) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetScopes

`func (o *CreateApiKey201Response) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateApiKey201Response) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateApiKey201Response) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetLabel

`func (o *CreateApiKey201Response) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateApiKey201Response) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateApiKey201Response) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLastUsedAt

`func (o *CreateApiKey201Response) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *CreateApiKey201Response) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *CreateApiKey201Response) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *CreateApiKey201Response) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetRevokedAt

`func (o *CreateApiKey201Response) GetRevokedAt() time.Time`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *CreateApiKey201Response) GetRevokedAtOk() (*time.Time, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *CreateApiKey201Response) SetRevokedAt(v time.Time)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *CreateApiKey201Response) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateApiKey201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateApiKey201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateApiKey201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSecret

`func (o *CreateApiKey201Response) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateApiKey201Response) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateApiKey201Response) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


