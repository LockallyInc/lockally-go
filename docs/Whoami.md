# Whoami

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenant** | [**WhoamiTenant**](WhoamiTenant.md) |  | 
**Scopes** | **[]string** |  | 
**Auth** | **string** |  | 
**Key** | Pointer to [**WhoamiKey**](WhoamiKey.md) |  | [optional] 

## Methods

### NewWhoami

`func NewWhoami(tenant WhoamiTenant, scopes []string, auth string, ) *Whoami`

NewWhoami instantiates a new Whoami object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoamiWithDefaults

`func NewWhoamiWithDefaults() *Whoami`

NewWhoamiWithDefaults instantiates a new Whoami object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenant

`func (o *Whoami) GetTenant() WhoamiTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Whoami) GetTenantOk() (*WhoamiTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Whoami) SetTenant(v WhoamiTenant)`

SetTenant sets Tenant field to given value.


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


### GetAuth

`func (o *Whoami) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *Whoami) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *Whoami) SetAuth(v string)`

SetAuth sets Auth field to given value.


### GetKey

`func (o *Whoami) GetKey() WhoamiKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Whoami) GetKeyOk() (*WhoamiKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Whoami) SetKey(v WhoamiKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *Whoami) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


