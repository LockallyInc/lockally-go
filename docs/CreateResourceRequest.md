# CreateResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "meeting_room"]
**Capacity** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateResourceRequest

`func NewCreateResourceRequest(name string, ) *CreateResourceRequest`

NewCreateResourceRequest instantiates a new CreateResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResourceRequestWithDefaults

`func NewCreateResourceRequestWithDefaults() *CreateResourceRequest`

NewCreateResourceRequestWithDefaults instantiates a new CreateResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateResourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateResourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateResourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateResourceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateResourceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateResourceRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateResourceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCapacity

`func (o *CreateResourceRequest) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *CreateResourceRequest) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *CreateResourceRequest) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *CreateResourceRequest) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


