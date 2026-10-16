# GetCurrentAdmin200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | [**Admin**](Admin.md) |  | 
**Tenant** | [**Tenant**](Tenant.md) |  | 

## Methods

### NewGetCurrentAdmin200Response

`func NewGetCurrentAdmin200Response(admin Admin, tenant Tenant, ) *GetCurrentAdmin200Response`

NewGetCurrentAdmin200Response instantiates a new GetCurrentAdmin200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrentAdmin200ResponseWithDefaults

`func NewGetCurrentAdmin200ResponseWithDefaults() *GetCurrentAdmin200Response`

NewGetCurrentAdmin200ResponseWithDefaults instantiates a new GetCurrentAdmin200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *GetCurrentAdmin200Response) GetAdmin() Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *GetCurrentAdmin200Response) GetAdminOk() (*Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *GetCurrentAdmin200Response) SetAdmin(v Admin)`

SetAdmin sets Admin field to given value.


### GetTenant

`func (o *GetCurrentAdmin200Response) GetTenant() Tenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *GetCurrentAdmin200Response) GetTenantOk() (*Tenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *GetCurrentAdmin200Response) SetTenant(v Tenant)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


