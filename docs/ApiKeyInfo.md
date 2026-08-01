# ApiKeyInfo

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

## Methods

### NewApiKeyInfo

`func NewApiKeyInfo(id string, tenantId string, prefix string, scopes []string, label string, createdAt time.Time, ) *ApiKeyInfo`

NewApiKeyInfo instantiates a new ApiKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyInfoWithDefaults

`func NewApiKeyInfoWithDefaults() *ApiKeyInfo`

NewApiKeyInfoWithDefaults instantiates a new ApiKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiKeyInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyInfo) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *ApiKeyInfo) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ApiKeyInfo) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ApiKeyInfo) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPrefix

`func (o *ApiKeyInfo) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ApiKeyInfo) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ApiKeyInfo) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetScopes

`func (o *ApiKeyInfo) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApiKeyInfo) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApiKeyInfo) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetLabel

`func (o *ApiKeyInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApiKeyInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApiKeyInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLastUsedAt

`func (o *ApiKeyInfo) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ApiKeyInfo) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ApiKeyInfo) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ApiKeyInfo) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetRevokedAt

`func (o *ApiKeyInfo) GetRevokedAt() time.Time`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *ApiKeyInfo) GetRevokedAtOk() (*time.Time, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *ApiKeyInfo) SetRevokedAt(v time.Time)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *ApiKeyInfo) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiKeyInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiKeyInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiKeyInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


