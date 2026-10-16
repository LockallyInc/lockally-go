# MailboxGrantList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grants** | [**[]MailboxGrant**](MailboxGrant.md) |  | 

## Methods

### NewMailboxGrantList

`func NewMailboxGrantList(grants []MailboxGrant, ) *MailboxGrantList`

NewMailboxGrantList instantiates a new MailboxGrantList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxGrantListWithDefaults

`func NewMailboxGrantListWithDefaults() *MailboxGrantList`

NewMailboxGrantListWithDefaults instantiates a new MailboxGrantList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrants

`func (o *MailboxGrantList) GetGrants() []MailboxGrant`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *MailboxGrantList) GetGrantsOk() (*[]MailboxGrant, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *MailboxGrantList) SetGrants(v []MailboxGrant)`

SetGrants sets Grants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


