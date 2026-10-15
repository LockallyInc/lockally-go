# MailboxGrantList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mailboxes** | [**[]MailboxGrant**](MailboxGrant.md) |  | 

## Methods

### NewMailboxGrantList

`func NewMailboxGrantList(mailboxes []MailboxGrant, ) *MailboxGrantList`

NewMailboxGrantList instantiates a new MailboxGrantList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxGrantListWithDefaults

`func NewMailboxGrantListWithDefaults() *MailboxGrantList`

NewMailboxGrantListWithDefaults instantiates a new MailboxGrantList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxes

`func (o *MailboxGrantList) GetMailboxes() []MailboxGrant`

GetMailboxes returns the Mailboxes field if non-nil, zero value otherwise.

### GetMailboxesOk

`func (o *MailboxGrantList) GetMailboxesOk() (*[]MailboxGrant, bool)`

GetMailboxesOk returns a tuple with the Mailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxes

`func (o *MailboxGrantList) SetMailboxes(v []MailboxGrant)`

SetMailboxes sets Mailboxes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


