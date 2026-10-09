# V1AdminLoginPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 
**Admin** | [**Admin**](Admin.md) |  | 
**Tenant** | [**Tenant**](Tenant.md) |  | 

## Methods

### NewV1AdminLoginPost200Response

`func NewV1AdminLoginPost200Response(token string, expiresAt time.Time, admin Admin, tenant Tenant, ) *V1AdminLoginPost200Response`

NewV1AdminLoginPost200Response instantiates a new V1AdminLoginPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AdminLoginPost200ResponseWithDefaults

`func NewV1AdminLoginPost200ResponseWithDefaults() *V1AdminLoginPost200Response`

NewV1AdminLoginPost200ResponseWithDefaults instantiates a new V1AdminLoginPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *V1AdminLoginPost200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *V1AdminLoginPost200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *V1AdminLoginPost200Response) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpiresAt

`func (o *V1AdminLoginPost200Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *V1AdminLoginPost200Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *V1AdminLoginPost200Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetAdmin

`func (o *V1AdminLoginPost200Response) GetAdmin() Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *V1AdminLoginPost200Response) GetAdminOk() (*Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *V1AdminLoginPost200Response) SetAdmin(v Admin)`

SetAdmin sets Admin field to given value.


### GetTenant

`func (o *V1AdminLoginPost200Response) GetTenant() Tenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *V1AdminLoginPost200Response) GetTenantOk() (*Tenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *V1AdminLoginPost200Response) SetTenant(v Tenant)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


