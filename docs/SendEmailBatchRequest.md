# SendEmailBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]SendMessage**](SendMessage.md) |  | 
**Variables** | Pointer to **map[string]string** | Template variables every message in the batch inherits. Each message&#39;s own &#x60;variables&#x60; override these key by key, so shared values are written once and per-recipient ones stay per-recipient.  Without this, a 500-message batch repeats every shared variable 500 times: most of the payload, and 500 places for one of them to differ by a typo.  | [optional] 

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


### GetVariables

`func (o *SendEmailBatchRequest) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *SendEmailBatchRequest) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *SendEmailBatchRequest) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *SendEmailBatchRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


