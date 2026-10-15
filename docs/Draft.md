# Draft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MailboxId** | **string** |  | 
**Mailbox** | Pointer to **string** |  | [optional] 
**ThreadId** | Pointer to **string** |  | [optional] 
**To** | **[]string** |  | 
**Cc** | Pointer to **[]string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**BodyText** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**PolicyFlags** | [**PolicyFlags**](PolicyFlags.md) |  | 
**ReviewedAt** | Pointer to **time.Time** |  | [optional] 
**SentMessageId** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewDraft

`func NewDraft(id string, mailboxId string, to []string, status string, policyFlags PolicyFlags, createdAt time.Time, updatedAt time.Time, ) *Draft`

NewDraft instantiates a new Draft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDraftWithDefaults

`func NewDraftWithDefaults() *Draft`

NewDraftWithDefaults instantiates a new Draft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Draft) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Draft) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Draft) SetId(v string)`

SetId sets Id field to given value.


### GetMailboxId

`func (o *Draft) GetMailboxId() string`

GetMailboxId returns the MailboxId field if non-nil, zero value otherwise.

### GetMailboxIdOk

`func (o *Draft) GetMailboxIdOk() (*string, bool)`

GetMailboxIdOk returns a tuple with the MailboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxId

`func (o *Draft) SetMailboxId(v string)`

SetMailboxId sets MailboxId field to given value.


### GetMailbox

`func (o *Draft) GetMailbox() string`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *Draft) GetMailboxOk() (*string, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *Draft) SetMailbox(v string)`

SetMailbox sets Mailbox field to given value.

### HasMailbox

`func (o *Draft) HasMailbox() bool`

HasMailbox returns a boolean if a field has been set.

### GetThreadId

`func (o *Draft) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *Draft) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *Draft) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *Draft) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetTo

`func (o *Draft) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Draft) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Draft) SetTo(v []string)`

SetTo sets To field to given value.


### GetCc

`func (o *Draft) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *Draft) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *Draft) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *Draft) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetSubject

`func (o *Draft) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Draft) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Draft) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Draft) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetBodyText

`func (o *Draft) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *Draft) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *Draft) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *Draft) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetStatus

`func (o *Draft) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Draft) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Draft) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPolicyFlags

`func (o *Draft) GetPolicyFlags() PolicyFlags`

GetPolicyFlags returns the PolicyFlags field if non-nil, zero value otherwise.

### GetPolicyFlagsOk

`func (o *Draft) GetPolicyFlagsOk() (*PolicyFlags, bool)`

GetPolicyFlagsOk returns a tuple with the PolicyFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyFlags

`func (o *Draft) SetPolicyFlags(v PolicyFlags)`

SetPolicyFlags sets PolicyFlags field to given value.


### GetReviewedAt

`func (o *Draft) GetReviewedAt() time.Time`

GetReviewedAt returns the ReviewedAt field if non-nil, zero value otherwise.

### GetReviewedAtOk

`func (o *Draft) GetReviewedAtOk() (*time.Time, bool)`

GetReviewedAtOk returns a tuple with the ReviewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedAt

`func (o *Draft) SetReviewedAt(v time.Time)`

SetReviewedAt sets ReviewedAt field to given value.

### HasReviewedAt

`func (o *Draft) HasReviewedAt() bool`

HasReviewedAt returns a boolean if a field has been set.

### GetSentMessageId

`func (o *Draft) GetSentMessageId() string`

GetSentMessageId returns the SentMessageId field if non-nil, zero value otherwise.

### GetSentMessageIdOk

`func (o *Draft) GetSentMessageIdOk() (*string, bool)`

GetSentMessageIdOk returns a tuple with the SentMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentMessageId

`func (o *Draft) SetSentMessageId(v string)`

SetSentMessageId sets SentMessageId field to given value.

### HasSentMessageId

`func (o *Draft) HasSentMessageId() bool`

HasSentMessageId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Draft) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Draft) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Draft) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Draft) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Draft) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Draft) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


