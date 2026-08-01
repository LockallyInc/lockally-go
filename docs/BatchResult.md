# BatchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** | Present when this message failed; the others are then absent. | [optional] 

## Methods

### NewBatchResult

`func NewBatchResult() *BatchResult`

NewBatchResult instantiates a new BatchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResultWithDefaults

`func NewBatchResultWithDefaults() *BatchResult`

NewBatchResultWithDefaults instantiates a new BatchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *BatchResult) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BatchResult) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BatchResult) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *BatchResult) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetId

`func (o *BatchResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BatchResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessageId

`func (o *BatchResult) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *BatchResult) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *BatchResult) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *BatchResult) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetStatus

`func (o *BatchResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BatchResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *BatchResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BatchResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BatchResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BatchResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


