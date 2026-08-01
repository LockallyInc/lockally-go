# GetUserInsights200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecentlyAdded** | Pointer to [**[]UserEvent**](UserEvent.md) |  | [optional] 
**RecentlySuspended** | Pointer to [**[]UserEvent**](UserEvent.md) |  | [optional] 
**Inactive30d** | Pointer to [**[]UserEvent**](UserEvent.md) |  | [optional] 
**SeatsUsed** | Pointer to **int32** |  | [optional] 
**SeatsAlloc** | Pointer to **int32** |  | [optional] 
**GeneratedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetUserInsights200Response

`func NewGetUserInsights200Response() *GetUserInsights200Response`

NewGetUserInsights200Response instantiates a new GetUserInsights200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserInsights200ResponseWithDefaults

`func NewGetUserInsights200ResponseWithDefaults() *GetUserInsights200Response`

NewGetUserInsights200ResponseWithDefaults instantiates a new GetUserInsights200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecentlyAdded

`func (o *GetUserInsights200Response) GetRecentlyAdded() []UserEvent`

GetRecentlyAdded returns the RecentlyAdded field if non-nil, zero value otherwise.

### GetRecentlyAddedOk

`func (o *GetUserInsights200Response) GetRecentlyAddedOk() (*[]UserEvent, bool)`

GetRecentlyAddedOk returns a tuple with the RecentlyAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentlyAdded

`func (o *GetUserInsights200Response) SetRecentlyAdded(v []UserEvent)`

SetRecentlyAdded sets RecentlyAdded field to given value.

### HasRecentlyAdded

`func (o *GetUserInsights200Response) HasRecentlyAdded() bool`

HasRecentlyAdded returns a boolean if a field has been set.

### GetRecentlySuspended

`func (o *GetUserInsights200Response) GetRecentlySuspended() []UserEvent`

GetRecentlySuspended returns the RecentlySuspended field if non-nil, zero value otherwise.

### GetRecentlySuspendedOk

`func (o *GetUserInsights200Response) GetRecentlySuspendedOk() (*[]UserEvent, bool)`

GetRecentlySuspendedOk returns a tuple with the RecentlySuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentlySuspended

`func (o *GetUserInsights200Response) SetRecentlySuspended(v []UserEvent)`

SetRecentlySuspended sets RecentlySuspended field to given value.

### HasRecentlySuspended

`func (o *GetUserInsights200Response) HasRecentlySuspended() bool`

HasRecentlySuspended returns a boolean if a field has been set.

### GetInactive30d

`func (o *GetUserInsights200Response) GetInactive30d() []UserEvent`

GetInactive30d returns the Inactive30d field if non-nil, zero value otherwise.

### GetInactive30dOk

`func (o *GetUserInsights200Response) GetInactive30dOk() (*[]UserEvent, bool)`

GetInactive30dOk returns a tuple with the Inactive30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive30d

`func (o *GetUserInsights200Response) SetInactive30d(v []UserEvent)`

SetInactive30d sets Inactive30d field to given value.

### HasInactive30d

`func (o *GetUserInsights200Response) HasInactive30d() bool`

HasInactive30d returns a boolean if a field has been set.

### GetSeatsUsed

`func (o *GetUserInsights200Response) GetSeatsUsed() int32`

GetSeatsUsed returns the SeatsUsed field if non-nil, zero value otherwise.

### GetSeatsUsedOk

`func (o *GetUserInsights200Response) GetSeatsUsedOk() (*int32, bool)`

GetSeatsUsedOk returns a tuple with the SeatsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatsUsed

`func (o *GetUserInsights200Response) SetSeatsUsed(v int32)`

SetSeatsUsed sets SeatsUsed field to given value.

### HasSeatsUsed

`func (o *GetUserInsights200Response) HasSeatsUsed() bool`

HasSeatsUsed returns a boolean if a field has been set.

### GetSeatsAlloc

`func (o *GetUserInsights200Response) GetSeatsAlloc() int32`

GetSeatsAlloc returns the SeatsAlloc field if non-nil, zero value otherwise.

### GetSeatsAllocOk

`func (o *GetUserInsights200Response) GetSeatsAllocOk() (*int32, bool)`

GetSeatsAllocOk returns a tuple with the SeatsAlloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatsAlloc

`func (o *GetUserInsights200Response) SetSeatsAlloc(v int32)`

SetSeatsAlloc sets SeatsAlloc field to given value.

### HasSeatsAlloc

`func (o *GetUserInsights200Response) HasSeatsAlloc() bool`

HasSeatsAlloc returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *GetUserInsights200Response) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *GetUserInsights200Response) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *GetUserInsights200Response) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *GetUserInsights200Response) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


