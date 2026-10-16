# ThreadList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mailbox** | **string** |  | 
**Threads** | [**[]ThreadSummary**](ThreadSummary.md) |  | 

## Methods

### NewThreadList

`func NewThreadList(mailbox string, threads []ThreadSummary, ) *ThreadList`

NewThreadList instantiates a new ThreadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadListWithDefaults

`func NewThreadListWithDefaults() *ThreadList`

NewThreadListWithDefaults instantiates a new ThreadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailbox

`func (o *ThreadList) GetMailbox() string`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *ThreadList) GetMailboxOk() (*string, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *ThreadList) SetMailbox(v string)`

SetMailbox sets Mailbox field to given value.


### GetThreads

`func (o *ThreadList) GetThreads() []ThreadSummary`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ThreadList) GetThreadsOk() (*[]ThreadSummary, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ThreadList) SetThreads(v []ThreadSummary)`

SetThreads sets Threads field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


