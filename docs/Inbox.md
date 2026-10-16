# Inbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Email** | **string** |  | 
**LastMessageAt** | Pointer to **time.Time** |  | [optional] 
**ThreadCount** | **int32** |  | 

## Methods

### NewInbox

`func NewInbox(id string, email string, threadCount int32, ) *Inbox`

NewInbox instantiates a new Inbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboxWithDefaults

`func NewInboxWithDefaults() *Inbox`

NewInboxWithDefaults instantiates a new Inbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Inbox) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Inbox) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Inbox) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *Inbox) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Inbox) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Inbox) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLastMessageAt

`func (o *Inbox) GetLastMessageAt() time.Time`

GetLastMessageAt returns the LastMessageAt field if non-nil, zero value otherwise.

### GetLastMessageAtOk

`func (o *Inbox) GetLastMessageAtOk() (*time.Time, bool)`

GetLastMessageAtOk returns a tuple with the LastMessageAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageAt

`func (o *Inbox) SetLastMessageAt(v time.Time)`

SetLastMessageAt sets LastMessageAt field to given value.

### HasLastMessageAt

`func (o *Inbox) HasLastMessageAt() bool`

HasLastMessageAt returns a boolean if a field has been set.

### GetThreadCount

`func (o *Inbox) GetThreadCount() int32`

GetThreadCount returns the ThreadCount field if non-nil, zero value otherwise.

### GetThreadCountOk

`func (o *Inbox) GetThreadCountOk() (*int32, bool)`

GetThreadCountOk returns a tuple with the ThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadCount

`func (o *Inbox) SetThreadCount(v int32)`

SetThreadCount sets ThreadCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


