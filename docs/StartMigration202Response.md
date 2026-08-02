# StartMigration202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**MailboxesQueued** | **int32** |  | 

## Methods

### NewStartMigration202Response

`func NewStartMigration202Response(status string, mailboxesQueued int32, ) *StartMigration202Response`

NewStartMigration202Response instantiates a new StartMigration202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartMigration202ResponseWithDefaults

`func NewStartMigration202ResponseWithDefaults() *StartMigration202Response`

NewStartMigration202ResponseWithDefaults instantiates a new StartMigration202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StartMigration202Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StartMigration202Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StartMigration202Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMailboxesQueued

`func (o *StartMigration202Response) GetMailboxesQueued() int32`

GetMailboxesQueued returns the MailboxesQueued field if non-nil, zero value otherwise.

### GetMailboxesQueuedOk

`func (o *StartMigration202Response) GetMailboxesQueuedOk() (*int32, bool)`

GetMailboxesQueuedOk returns a tuple with the MailboxesQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxesQueued

`func (o *StartMigration202Response) SetMailboxesQueued(v int32)`

SetMailboxesQueued sets MailboxesQueued field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


