# DedicatedIPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Status** | **string** |  | 
**Note** | Pointer to **string** |  | [optional] 
**AdminNote** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**ResolvedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDedicatedIPRequest

`func NewDedicatedIPRequest(id string, tenantId string, status string, createdAt time.Time, ) *DedicatedIPRequest`

NewDedicatedIPRequest instantiates a new DedicatedIPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedIPRequestWithDefaults

`func NewDedicatedIPRequestWithDefaults() *DedicatedIPRequest`

NewDedicatedIPRequestWithDefaults instantiates a new DedicatedIPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DedicatedIPRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DedicatedIPRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DedicatedIPRequest) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *DedicatedIPRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DedicatedIPRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DedicatedIPRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetStatus

`func (o *DedicatedIPRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DedicatedIPRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DedicatedIPRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNote

`func (o *DedicatedIPRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *DedicatedIPRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *DedicatedIPRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *DedicatedIPRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetAdminNote

`func (o *DedicatedIPRequest) GetAdminNote() string`

GetAdminNote returns the AdminNote field if non-nil, zero value otherwise.

### GetAdminNoteOk

`func (o *DedicatedIPRequest) GetAdminNoteOk() (*string, bool)`

GetAdminNoteOk returns a tuple with the AdminNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminNote

`func (o *DedicatedIPRequest) SetAdminNote(v string)`

SetAdminNote sets AdminNote field to given value.

### HasAdminNote

`func (o *DedicatedIPRequest) HasAdminNote() bool`

HasAdminNote returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DedicatedIPRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DedicatedIPRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DedicatedIPRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetResolvedAt

`func (o *DedicatedIPRequest) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *DedicatedIPRequest) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *DedicatedIPRequest) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *DedicatedIPRequest) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


