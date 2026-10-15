# MailboxGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Mailbox** | Pointer to **string** |  | [optional] 

## Methods

### NewMailboxGrant

`func NewMailboxGrant() *MailboxGrant`

NewMailboxGrant instantiates a new MailboxGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxGrantWithDefaults

`func NewMailboxGrantWithDefaults() *MailboxGrant`

NewMailboxGrantWithDefaults instantiates a new MailboxGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MailboxGrant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MailboxGrant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MailboxGrant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MailboxGrant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMailbox

`func (o *MailboxGrant) GetMailbox() string`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *MailboxGrant) GetMailboxOk() (*string, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *MailboxGrant) SetMailbox(v string)`

SetMailbox sets Mailbox field to given value.

### HasMailbox

`func (o *MailboxGrant) HasMailbox() bool`

HasMailbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


