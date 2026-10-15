# ThreadClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | **string** |  | 
**Cached** | Pointer to **bool** | True when a previous classification was reused. | [optional] 
**Billed** | Pointer to **bool** | True when this call consumed an AI unit. | [optional] 
**Classification** | Pointer to [**Classification**](Classification.md) |  | [optional] 

## Methods

### NewThreadClassification

`func NewThreadClassification(threadId string, ) *ThreadClassification`

NewThreadClassification instantiates a new ThreadClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadClassificationWithDefaults

`func NewThreadClassificationWithDefaults() *ThreadClassification`

NewThreadClassificationWithDefaults instantiates a new ThreadClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *ThreadClassification) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *ThreadClassification) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *ThreadClassification) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetCached

`func (o *ThreadClassification) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *ThreadClassification) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *ThreadClassification) SetCached(v bool)`

SetCached sets Cached field to given value.

### HasCached

`func (o *ThreadClassification) HasCached() bool`

HasCached returns a boolean if a field has been set.

### GetBilled

`func (o *ThreadClassification) GetBilled() bool`

GetBilled returns the Billed field if non-nil, zero value otherwise.

### GetBilledOk

`func (o *ThreadClassification) GetBilledOk() (*bool, bool)`

GetBilledOk returns a tuple with the Billed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilled

`func (o *ThreadClassification) SetBilled(v bool)`

SetBilled sets Billed field to given value.

### HasBilled

`func (o *ThreadClassification) HasBilled() bool`

HasBilled returns a boolean if a field has been set.

### GetClassification

`func (o *ThreadClassification) GetClassification() Classification`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *ThreadClassification) GetClassificationOk() (*Classification, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *ThreadClassification) SetClassification(v Classification)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *ThreadClassification) HasClassification() bool`

HasClassification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


