# MessageDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**MessageId** | **string** | RFC 5322 Message-ID header, including angle brackets. | 
**Sender** | **string** |  | 
**Recipients** | **[]string** |  | 
**Subject** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**QueuedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**BounceReason** | Pointer to **string** |  | [optional] 
**SizeBytes** | Pointer to **int32** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **[]string** |  | [optional] 
**Cc** | Pointer to **[]string** |  | [optional] 
**Bcc** | Pointer to **[]string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Html** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMessageDetail

`func NewMessageDetail(id string, tenantId string, messageId string, sender string, recipients []string, status string, queuedAt time.Time, updatedAt time.Time, ) *MessageDetail`

NewMessageDetail instantiates a new MessageDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDetailWithDefaults

`func NewMessageDetailWithDefaults() *MessageDetail`

NewMessageDetailWithDefaults instantiates a new MessageDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageDetail) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *MessageDetail) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MessageDetail) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MessageDetail) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetMessageId

`func (o *MessageDetail) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageDetail) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageDetail) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetSender

`func (o *MessageDetail) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MessageDetail) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *MessageDetail) SetSender(v string)`

SetSender sets Sender field to given value.


### GetRecipients

`func (o *MessageDetail) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *MessageDetail) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *MessageDetail) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.


### GetSubject

`func (o *MessageDetail) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MessageDetail) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MessageDetail) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MessageDetail) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetStatus

`func (o *MessageDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageDetail) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetQueuedAt

`func (o *MessageDetail) GetQueuedAt() time.Time`

GetQueuedAt returns the QueuedAt field if non-nil, zero value otherwise.

### GetQueuedAtOk

`func (o *MessageDetail) GetQueuedAtOk() (*time.Time, bool)`

GetQueuedAtOk returns a tuple with the QueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedAt

`func (o *MessageDetail) SetQueuedAt(v time.Time)`

SetQueuedAt sets QueuedAt field to given value.


### GetUpdatedAt

`func (o *MessageDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MessageDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MessageDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetBounceReason

`func (o *MessageDetail) GetBounceReason() string`

GetBounceReason returns the BounceReason field if non-nil, zero value otherwise.

### GetBounceReasonOk

`func (o *MessageDetail) GetBounceReasonOk() (*string, bool)`

GetBounceReasonOk returns a tuple with the BounceReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceReason

`func (o *MessageDetail) SetBounceReason(v string)`

SetBounceReason sets BounceReason field to given value.

### HasBounceReason

`func (o *MessageDetail) HasBounceReason() bool`

HasBounceReason returns a boolean if a field has been set.

### GetSizeBytes

`func (o *MessageDetail) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *MessageDetail) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *MessageDetail) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *MessageDetail) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetFrom

`func (o *MessageDetail) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MessageDetail) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MessageDetail) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MessageDetail) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *MessageDetail) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MessageDetail) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MessageDetail) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *MessageDetail) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCc

`func (o *MessageDetail) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *MessageDetail) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *MessageDetail) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *MessageDetail) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *MessageDetail) GetBcc() []string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *MessageDetail) GetBccOk() (*[]string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *MessageDetail) SetBcc(v []string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *MessageDetail) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetText

`func (o *MessageDetail) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MessageDetail) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MessageDetail) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *MessageDetail) HasText() bool`

HasText returns a boolean if a field has been set.

### GetHtml

`func (o *MessageDetail) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *MessageDetail) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *MessageDetail) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *MessageDetail) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetHeaders

`func (o *MessageDetail) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MessageDetail) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MessageDetail) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MessageDetail) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


