# AdminLogin200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 
**Admin** | [**Admin**](Admin.md) |  | 
**Tenant** | [**Tenant**](Tenant.md) |  | 

## Methods

### NewAdminLogin200Response

`func NewAdminLogin200Response(token string, expiresAt time.Time, admin Admin, tenant Tenant, ) *AdminLogin200Response`

NewAdminLogin200Response instantiates a new AdminLogin200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminLogin200ResponseWithDefaults

`func NewAdminLogin200ResponseWithDefaults() *AdminLogin200Response`

NewAdminLogin200ResponseWithDefaults instantiates a new AdminLogin200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AdminLogin200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AdminLogin200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AdminLogin200Response) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpiresAt

`func (o *AdminLogin200Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AdminLogin200Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AdminLogin200Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetAdmin

`func (o *AdminLogin200Response) GetAdmin() Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *AdminLogin200Response) GetAdminOk() (*Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *AdminLogin200Response) SetAdmin(v Admin)`

SetAdmin sets Admin field to given value.


### GetTenant

`func (o *AdminLogin200Response) GetTenant() Tenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AdminLogin200Response) GetTenantOk() (*Tenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AdminLogin200Response) SetTenant(v Tenant)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


