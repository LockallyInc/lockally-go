# ThreadMessageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | **string** |  | 
**Mailbox** | **string** |  | 
**Message** | [**ThreadMessage**](ThreadMessage.md) |  | 
**Body** | Pointer to [**ThreadMessageBodyBody**](ThreadMessageBodyBody.md) |  | [optional] 
**BodyUnavailable** | Pointer to **string** | Present instead of body when the message could not be located in the mailbox. | [optional] 

## Methods

### NewThreadMessageBody

`func NewThreadMessageBody(threadId string, mailbox string, message ThreadMessage, ) *ThreadMessageBody`

NewThreadMessageBody instantiates a new ThreadMessageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadMessageBodyWithDefaults

`func NewThreadMessageBodyWithDefaults() *ThreadMessageBody`

NewThreadMessageBodyWithDefaults instantiates a new ThreadMessageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *ThreadMessageBody) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *ThreadMessageBody) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *ThreadMessageBody) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetMailbox

`func (o *ThreadMessageBody) GetMailbox() string`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *ThreadMessageBody) GetMailboxOk() (*string, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *ThreadMessageBody) SetMailbox(v string)`

SetMailbox sets Mailbox field to given value.


### GetMessage

`func (o *ThreadMessageBody) GetMessage() ThreadMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ThreadMessageBody) GetMessageOk() (*ThreadMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ThreadMessageBody) SetMessage(v ThreadMessage)`

SetMessage sets Message field to given value.


### GetBody

`func (o *ThreadMessageBody) GetBody() ThreadMessageBodyBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ThreadMessageBody) GetBodyOk() (*ThreadMessageBodyBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ThreadMessageBody) SetBody(v ThreadMessageBodyBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *ThreadMessageBody) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBodyUnavailable

`func (o *ThreadMessageBody) GetBodyUnavailable() string`

GetBodyUnavailable returns the BodyUnavailable field if non-nil, zero value otherwise.

### GetBodyUnavailableOk

`func (o *ThreadMessageBody) GetBodyUnavailableOk() (*string, bool)`

GetBodyUnavailableOk returns a tuple with the BodyUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyUnavailable

`func (o *ThreadMessageBody) SetBodyUnavailable(v string)`

SetBodyUnavailable sets BodyUnavailable field to given value.

### HasBodyUnavailable

`func (o *ThreadMessageBody) HasBodyUnavailable() bool`

HasBodyUnavailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


