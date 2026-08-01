# GetStorage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalBytes** | Pointer to **int64** |  | [optional] 
**AllocBytes** | Pointer to **int64** |  | [optional] 
**PerSeatBytes** | Pointer to **int64** |  | [optional] 
**TopMailboxes** | Pointer to [**[]GetStorage200ResponseTopMailboxesInner**](GetStorage200ResponseTopMailboxesInner.md) |  | [optional] 
**TopMessages** | Pointer to [**[]GetStorage200ResponseTopMessagesInner**](GetStorage200ResponseTopMessagesInner.md) |  | [optional] 
**GeneratedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetStorage200Response

`func NewGetStorage200Response() *GetStorage200Response`

NewGetStorage200Response instantiates a new GetStorage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorage200ResponseWithDefaults

`func NewGetStorage200ResponseWithDefaults() *GetStorage200Response`

NewGetStorage200ResponseWithDefaults instantiates a new GetStorage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalBytes

`func (o *GetStorage200Response) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *GetStorage200Response) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *GetStorage200Response) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *GetStorage200Response) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.

### GetAllocBytes

`func (o *GetStorage200Response) GetAllocBytes() int64`

GetAllocBytes returns the AllocBytes field if non-nil, zero value otherwise.

### GetAllocBytesOk

`func (o *GetStorage200Response) GetAllocBytesOk() (*int64, bool)`

GetAllocBytesOk returns a tuple with the AllocBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocBytes

`func (o *GetStorage200Response) SetAllocBytes(v int64)`

SetAllocBytes sets AllocBytes field to given value.

### HasAllocBytes

`func (o *GetStorage200Response) HasAllocBytes() bool`

HasAllocBytes returns a boolean if a field has been set.

### GetPerSeatBytes

`func (o *GetStorage200Response) GetPerSeatBytes() int64`

GetPerSeatBytes returns the PerSeatBytes field if non-nil, zero value otherwise.

### GetPerSeatBytesOk

`func (o *GetStorage200Response) GetPerSeatBytesOk() (*int64, bool)`

GetPerSeatBytesOk returns a tuple with the PerSeatBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSeatBytes

`func (o *GetStorage200Response) SetPerSeatBytes(v int64)`

SetPerSeatBytes sets PerSeatBytes field to given value.

### HasPerSeatBytes

`func (o *GetStorage200Response) HasPerSeatBytes() bool`

HasPerSeatBytes returns a boolean if a field has been set.

### GetTopMailboxes

`func (o *GetStorage200Response) GetTopMailboxes() []GetStorage200ResponseTopMailboxesInner`

GetTopMailboxes returns the TopMailboxes field if non-nil, zero value otherwise.

### GetTopMailboxesOk

`func (o *GetStorage200Response) GetTopMailboxesOk() (*[]GetStorage200ResponseTopMailboxesInner, bool)`

GetTopMailboxesOk returns a tuple with the TopMailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopMailboxes

`func (o *GetStorage200Response) SetTopMailboxes(v []GetStorage200ResponseTopMailboxesInner)`

SetTopMailboxes sets TopMailboxes field to given value.

### HasTopMailboxes

`func (o *GetStorage200Response) HasTopMailboxes() bool`

HasTopMailboxes returns a boolean if a field has been set.

### GetTopMessages

`func (o *GetStorage200Response) GetTopMessages() []GetStorage200ResponseTopMessagesInner`

GetTopMessages returns the TopMessages field if non-nil, zero value otherwise.

### GetTopMessagesOk

`func (o *GetStorage200Response) GetTopMessagesOk() (*[]GetStorage200ResponseTopMessagesInner, bool)`

GetTopMessagesOk returns a tuple with the TopMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopMessages

`func (o *GetStorage200Response) SetTopMessages(v []GetStorage200ResponseTopMessagesInner)`

SetTopMessages sets TopMessages field to given value.

### HasTopMessages

`func (o *GetStorage200Response) HasTopMessages() bool`

HasTopMessages returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *GetStorage200Response) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *GetStorage200Response) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *GetStorage200Response) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *GetStorage200Response) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


