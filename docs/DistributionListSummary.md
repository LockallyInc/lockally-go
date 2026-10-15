# DistributionListSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**ListAddress** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**MemberCount** | **int32** |  | 

## Methods

### NewDistributionListSummary

`func NewDistributionListSummary(id string, tenantId string, listAddress string, createdAt time.Time, memberCount int32, ) *DistributionListSummary`

NewDistributionListSummary instantiates a new DistributionListSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionListSummaryWithDefaults

`func NewDistributionListSummaryWithDefaults() *DistributionListSummary`

NewDistributionListSummaryWithDefaults instantiates a new DistributionListSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DistributionListSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DistributionListSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DistributionListSummary) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *DistributionListSummary) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DistributionListSummary) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DistributionListSummary) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetListAddress

`func (o *DistributionListSummary) GetListAddress() string`

GetListAddress returns the ListAddress field if non-nil, zero value otherwise.

### GetListAddressOk

`func (o *DistributionListSummary) GetListAddressOk() (*string, bool)`

GetListAddressOk returns a tuple with the ListAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListAddress

`func (o *DistributionListSummary) SetListAddress(v string)`

SetListAddress sets ListAddress field to given value.


### GetName

`func (o *DistributionListSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DistributionListSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DistributionListSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DistributionListSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DistributionListSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DistributionListSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DistributionListSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMemberCount

`func (o *DistributionListSummary) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *DistributionListSummary) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *DistributionListSummary) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


