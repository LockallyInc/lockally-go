# Whoami

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | The API key&#39;s label. | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWhoami

`func NewWhoami() *Whoami`

NewWhoami instantiates a new Whoami object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoamiWithDefaults

`func NewWhoamiWithDefaults() *Whoami`

NewWhoamiWithDefaults instantiates a new Whoami object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *Whoami) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Whoami) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Whoami) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Whoami) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTenant

`func (o *Whoami) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Whoami) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Whoami) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Whoami) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetLabel

`func (o *Whoami) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Whoami) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Whoami) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Whoami) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetScopes

`func (o *Whoami) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Whoami) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Whoami) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *Whoami) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


