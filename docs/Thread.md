# Thread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Thread** | [**ThreadSummary**](ThreadSummary.md) |  | 
**Mailbox** | **string** |  | 
**Messages** | [**[]ThreadMessage**](ThreadMessage.md) | Chronological. Reading never marks a message read. | 

## Methods

### NewThread

`func NewThread(thread ThreadSummary, mailbox string, messages []ThreadMessage, ) *Thread`

NewThread instantiates a new Thread object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadWithDefaults

`func NewThreadWithDefaults() *Thread`

NewThreadWithDefaults instantiates a new Thread object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThread

`func (o *Thread) GetThread() ThreadSummary`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *Thread) GetThreadOk() (*ThreadSummary, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *Thread) SetThread(v ThreadSummary)`

SetThread sets Thread field to given value.


### GetMailbox

`func (o *Thread) GetMailbox() string`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *Thread) GetMailboxOk() (*string, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *Thread) SetMailbox(v string)`

SetMailbox sets Mailbox field to given value.


### GetMessages

`func (o *Thread) GetMessages() []ThreadMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Thread) GetMessagesOk() (*[]ThreadMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Thread) SetMessages(v []ThreadMessage)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


