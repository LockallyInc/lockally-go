# GetIPAssignment200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolName** | Pointer to **string** |  | [optional] 
**PoolKind** | Pointer to **string** |  | [optional] 
**DedicatedIp** | Pointer to **string** |  | [optional] 
**AssignedAt** | Pointer to **time.Time** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewGetIPAssignment200Response

`func NewGetIPAssignment200Response() *GetIPAssignment200Response`

NewGetIPAssignment200Response instantiates a new GetIPAssignment200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIPAssignment200ResponseWithDefaults

`func NewGetIPAssignment200ResponseWithDefaults() *GetIPAssignment200Response`

NewGetIPAssignment200ResponseWithDefaults instantiates a new GetIPAssignment200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolName

`func (o *GetIPAssignment200Response) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *GetIPAssignment200Response) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *GetIPAssignment200Response) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *GetIPAssignment200Response) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolKind

`func (o *GetIPAssignment200Response) GetPoolKind() string`

GetPoolKind returns the PoolKind field if non-nil, zero value otherwise.

### GetPoolKindOk

`func (o *GetIPAssignment200Response) GetPoolKindOk() (*string, bool)`

GetPoolKindOk returns a tuple with the PoolKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolKind

`func (o *GetIPAssignment200Response) SetPoolKind(v string)`

SetPoolKind sets PoolKind field to given value.

### HasPoolKind

`func (o *GetIPAssignment200Response) HasPoolKind() bool`

HasPoolKind returns a boolean if a field has been set.

### GetDedicatedIp

`func (o *GetIPAssignment200Response) GetDedicatedIp() string`

GetDedicatedIp returns the DedicatedIp field if non-nil, zero value otherwise.

### GetDedicatedIpOk

`func (o *GetIPAssignment200Response) GetDedicatedIpOk() (*string, bool)`

GetDedicatedIpOk returns a tuple with the DedicatedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedIp

`func (o *GetIPAssignment200Response) SetDedicatedIp(v string)`

SetDedicatedIp sets DedicatedIp field to given value.

### HasDedicatedIp

`func (o *GetIPAssignment200Response) HasDedicatedIp() bool`

HasDedicatedIp returns a boolean if a field has been set.

### GetAssignedAt

`func (o *GetIPAssignment200Response) GetAssignedAt() time.Time`

GetAssignedAt returns the AssignedAt field if non-nil, zero value otherwise.

### GetAssignedAtOk

`func (o *GetIPAssignment200Response) GetAssignedAtOk() (*time.Time, bool)`

GetAssignedAtOk returns a tuple with the AssignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedAt

`func (o *GetIPAssignment200Response) SetAssignedAt(v time.Time)`

SetAssignedAt sets AssignedAt field to given value.

### HasAssignedAt

`func (o *GetIPAssignment200Response) HasAssignedAt() bool`

HasAssignedAt returns a boolean if a field has been set.

### GetReason

`func (o *GetIPAssignment200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetIPAssignment200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetIPAssignment200Response) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetIPAssignment200Response) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


