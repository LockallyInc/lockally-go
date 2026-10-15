# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Slug** | **string** |  | 
**DisplayName** | **string** |  | 
**Status** | **string** |  | 
**Plan** | **string** |  | 
**RateCapPerMin** | **int32** | Per-tenant share of the per-VPS 5/min outbound cap (L6). | 
**DailyMsgQuota** | **int32** |  | 
**AdminEmail** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**SuspendedAt** | Pointer to **time.Time** |  | [optional] 
**ClosedAt** | Pointer to **time.Time** |  | [optional] 
**HardDeleteAfter** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTenant

`func NewTenant(id string, slug string, displayName string, status string, plan string, rateCapPerMin int32, dailyMsgQuota int32, adminEmail string, createdAt time.Time, ) *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tenant) SetId(v string)`

SetId sets Id field to given value.


### GetSlug

`func (o *Tenant) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Tenant) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Tenant) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDisplayName

`func (o *Tenant) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Tenant) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Tenant) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetStatus

`func (o *Tenant) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Tenant) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Tenant) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlan

`func (o *Tenant) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Tenant) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Tenant) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetRateCapPerMin

`func (o *Tenant) GetRateCapPerMin() int32`

GetRateCapPerMin returns the RateCapPerMin field if non-nil, zero value otherwise.

### GetRateCapPerMinOk

`func (o *Tenant) GetRateCapPerMinOk() (*int32, bool)`

GetRateCapPerMinOk returns a tuple with the RateCapPerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCapPerMin

`func (o *Tenant) SetRateCapPerMin(v int32)`

SetRateCapPerMin sets RateCapPerMin field to given value.


### GetDailyMsgQuota

`func (o *Tenant) GetDailyMsgQuota() int32`

GetDailyMsgQuota returns the DailyMsgQuota field if non-nil, zero value otherwise.

### GetDailyMsgQuotaOk

`func (o *Tenant) GetDailyMsgQuotaOk() (*int32, bool)`

GetDailyMsgQuotaOk returns a tuple with the DailyMsgQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyMsgQuota

`func (o *Tenant) SetDailyMsgQuota(v int32)`

SetDailyMsgQuota sets DailyMsgQuota field to given value.


### GetAdminEmail

`func (o *Tenant) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *Tenant) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *Tenant) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.


### GetCreatedAt

`func (o *Tenant) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Tenant) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Tenant) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSuspendedAt

`func (o *Tenant) GetSuspendedAt() time.Time`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *Tenant) GetSuspendedAtOk() (*time.Time, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *Tenant) SetSuspendedAt(v time.Time)`

SetSuspendedAt sets SuspendedAt field to given value.

### HasSuspendedAt

`func (o *Tenant) HasSuspendedAt() bool`

HasSuspendedAt returns a boolean if a field has been set.

### GetClosedAt

`func (o *Tenant) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *Tenant) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *Tenant) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *Tenant) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### GetHardDeleteAfter

`func (o *Tenant) GetHardDeleteAfter() time.Time`

GetHardDeleteAfter returns the HardDeleteAfter field if non-nil, zero value otherwise.

### GetHardDeleteAfterOk

`func (o *Tenant) GetHardDeleteAfterOk() (*time.Time, bool)`

GetHardDeleteAfterOk returns a tuple with the HardDeleteAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDeleteAfter

`func (o *Tenant) SetHardDeleteAfter(v time.Time)`

SetHardDeleteAfter sets HardDeleteAfter field to given value.

### HasHardDeleteAfter

`func (o *Tenant) HasHardDeleteAfter() bool`

HasHardDeleteAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


