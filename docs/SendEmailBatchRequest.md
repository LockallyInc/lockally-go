# SendEmailBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]SendMessage**](SendMessage.md) |  | 

## Methods

### NewSendEmailBatchRequest

`func NewSendEmailBatchRequest(messages []SendMessage, ) *SendEmailBatchRequest`

NewSendEmailBatchRequest instantiates a new SendEmailBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailBatchRequestWithDefaults

`func NewSendEmailBatchRequestWithDefaults() *SendEmailBatchRequest`

NewSendEmailBatchRequestWithDefaults instantiates a new SendEmailBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *SendEmailBatchRequest) GetMessages() []SendMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *SendEmailBatchRequest) GetMessagesOk() (*[]SendMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *SendEmailBatchRequest) SetMessages(v []SendMessage)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


