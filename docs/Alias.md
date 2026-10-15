# Alias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**AliasAddress** | **string** |  | 
**AliasTarget** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewAlias

`func NewAlias(id string, tenantId string, aliasAddress string, aliasTarget string, createdAt time.Time, ) *Alias`

NewAlias instantiates a new Alias object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliasWithDefaults

`func NewAliasWithDefaults() *Alias`

NewAliasWithDefaults instantiates a new Alias object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Alias) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Alias) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Alias) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *Alias) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Alias) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Alias) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAliasAddress

`func (o *Alias) GetAliasAddress() string`

GetAliasAddress returns the AliasAddress field if non-nil, zero value otherwise.

### GetAliasAddressOk

`func (o *Alias) GetAliasAddressOk() (*string, bool)`

GetAliasAddressOk returns a tuple with the AliasAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasAddress

`func (o *Alias) SetAliasAddress(v string)`

SetAliasAddress sets AliasAddress field to given value.


### GetAliasTarget

`func (o *Alias) GetAliasTarget() string`

GetAliasTarget returns the AliasTarget field if non-nil, zero value otherwise.

### GetAliasTargetOk

`func (o *Alias) GetAliasTargetOk() (*string, bool)`

GetAliasTargetOk returns a tuple with the AliasTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTarget

`func (o *Alias) SetAliasTarget(v string)`

SetAliasTarget sets AliasTarget field to given value.


### GetCreatedAt

`func (o *Alias) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Alias) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Alias) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


