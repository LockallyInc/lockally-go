# V1SendBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]SendMessage**](SendMessage.md) |  | 

## Methods

### NewV1SendBatchPostRequest

`func NewV1SendBatchPostRequest(messages []SendMessage, ) *V1SendBatchPostRequest`

NewV1SendBatchPostRequest instantiates a new V1SendBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SendBatchPostRequestWithDefaults

`func NewV1SendBatchPostRequestWithDefaults() *V1SendBatchPostRequest`

NewV1SendBatchPostRequestWithDefaults instantiates a new V1SendBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *V1SendBatchPostRequest) GetMessages() []SendMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *V1SendBatchPostRequest) GetMessagesOk() (*[]SendMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *V1SendBatchPostRequest) SetMessages(v []SendMessage)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


