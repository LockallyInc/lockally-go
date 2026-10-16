# AgentSendResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Lockally message id. | 
**MessageId** | **string** | The Internet Message-ID. | 
**Status** | **string** |  | 
**Warning** | Pointer to **string** | Set when the send is accepted but worth noting | [optional] 
**ThreadId** | **string** | The conversation the message joined. | 

## Methods

### NewAgentSendResult

`func NewAgentSendResult(id string, messageId string, status string, threadId string, ) *AgentSendResult`

NewAgentSendResult instantiates a new AgentSendResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSendResultWithDefaults

`func NewAgentSendResultWithDefaults() *AgentSendResult`

NewAgentSendResultWithDefaults instantiates a new AgentSendResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentSendResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentSendResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentSendResult) SetId(v string)`

SetId sets Id field to given value.


### GetMessageId

`func (o *AgentSendResult) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AgentSendResult) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AgentSendResult) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetStatus

`func (o *AgentSendResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentSendResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentSendResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetWarning

`func (o *AgentSendResult) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *AgentSendResult) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *AgentSendResult) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *AgentSendResult) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetThreadId

`func (o *AgentSendResult) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AgentSendResult) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AgentSendResult) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


