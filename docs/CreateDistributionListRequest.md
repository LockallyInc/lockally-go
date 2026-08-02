# CreateDistributionListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListAddress** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Members** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateDistributionListRequest

`func NewCreateDistributionListRequest(listAddress string, ) *CreateDistributionListRequest`

NewCreateDistributionListRequest instantiates a new CreateDistributionListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDistributionListRequestWithDefaults

`func NewCreateDistributionListRequestWithDefaults() *CreateDistributionListRequest`

NewCreateDistributionListRequestWithDefaults instantiates a new CreateDistributionListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListAddress

`func (o *CreateDistributionListRequest) GetListAddress() string`

GetListAddress returns the ListAddress field if non-nil, zero value otherwise.

### GetListAddressOk

`func (o *CreateDistributionListRequest) GetListAddressOk() (*string, bool)`

GetListAddressOk returns a tuple with the ListAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListAddress

`func (o *CreateDistributionListRequest) SetListAddress(v string)`

SetListAddress sets ListAddress field to given value.


### GetName

`func (o *CreateDistributionListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDistributionListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDistributionListRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateDistributionListRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMembers

`func (o *CreateDistributionListRequest) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *CreateDistributionListRequest) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *CreateDistributionListRequest) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *CreateDistributionListRequest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


