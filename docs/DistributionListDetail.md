# DistributionListDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**ListAddress** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Members** | **[]string** |  | 
**MemberCount** | **int32** |  | 

## Methods

### NewDistributionListDetail

`func NewDistributionListDetail(id string, tenantId string, listAddress string, createdAt time.Time, members []string, memberCount int32, ) *DistributionListDetail`

NewDistributionListDetail instantiates a new DistributionListDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionListDetailWithDefaults

`func NewDistributionListDetailWithDefaults() *DistributionListDetail`

NewDistributionListDetailWithDefaults instantiates a new DistributionListDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DistributionListDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DistributionListDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DistributionListDetail) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *DistributionListDetail) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DistributionListDetail) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DistributionListDetail) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetListAddress

`func (o *DistributionListDetail) GetListAddress() string`

GetListAddress returns the ListAddress field if non-nil, zero value otherwise.

### GetListAddressOk

`func (o *DistributionListDetail) GetListAddressOk() (*string, bool)`

GetListAddressOk returns a tuple with the ListAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListAddress

`func (o *DistributionListDetail) SetListAddress(v string)`

SetListAddress sets ListAddress field to given value.


### GetName

`func (o *DistributionListDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DistributionListDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DistributionListDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DistributionListDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DistributionListDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DistributionListDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DistributionListDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMembers

`func (o *DistributionListDetail) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DistributionListDetail) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *DistributionListDetail) SetMembers(v []string)`

SetMembers sets Members field to given value.


### GetMemberCount

`func (o *DistributionListDetail) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *DistributionListDetail) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *DistributionListDetail) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


