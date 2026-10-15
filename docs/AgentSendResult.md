# AgentSendResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Lockally message id. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ThreadId** | **string** | The conversation the message joined. | 

## Methods

### NewAgentSendResult

`func NewAgentSendResult(threadId string, ) *AgentSendResult`

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

### HasId

`func (o *AgentSendResult) HasId() bool`

HasId returns a boolean if a field has been set.

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

### HasStatus

`func (o *AgentSendResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

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


