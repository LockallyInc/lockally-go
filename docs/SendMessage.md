# SendMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** |  | 
**To** | **[]string** |  | 
**Cc** | Pointer to **[]string** |  | [optional] 
**Bcc** | Pointer to **[]string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Html** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Unsubscribe** | Pointer to **bool** |  | [optional] 
**TemplateId** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **map[string]string** |  | [optional] 
**SendAt** | Pointer to **time.Time** |  | [optional] 
**Attachments** | Pointer to [**[]V1SendPostRequestAttachmentsInner**](V1SendPostRequestAttachmentsInner.md) |  | [optional] 

## Methods

### NewSendMessage

`func NewSendMessage(from string, to []string, ) *SendMessage`

NewSendMessage instantiates a new SendMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageWithDefaults

`func NewSendMessageWithDefaults() *SendMessage`

NewSendMessageWithDefaults instantiates a new SendMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *SendMessage) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SendMessage) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SendMessage) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *SendMessage) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SendMessage) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SendMessage) SetTo(v []string)`

SetTo sets To field to given value.


### GetCc

`func (o *SendMessage) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *SendMessage) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *SendMessage) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *SendMessage) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *SendMessage) GetBcc() []string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *SendMessage) GetBccOk() (*[]string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *SendMessage) SetBcc(v []string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *SendMessage) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetSubject

`func (o *SendMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SendMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SendMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SendMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *SendMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SendMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SendMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SendMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetHtml

`func (o *SendMessage) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *SendMessage) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *SendMessage) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *SendMessage) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetHeaders

`func (o *SendMessage) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SendMessage) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *SendMessage) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *SendMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetUnsubscribe

`func (o *SendMessage) GetUnsubscribe() bool`

GetUnsubscribe returns the Unsubscribe field if non-nil, zero value otherwise.

### GetUnsubscribeOk

`func (o *SendMessage) GetUnsubscribeOk() (*bool, bool)`

GetUnsubscribeOk returns a tuple with the Unsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribe

`func (o *SendMessage) SetUnsubscribe(v bool)`

SetUnsubscribe sets Unsubscribe field to given value.

### HasUnsubscribe

`func (o *SendMessage) HasUnsubscribe() bool`

HasUnsubscribe returns a boolean if a field has been set.

### GetTemplateId

`func (o *SendMessage) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *SendMessage) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *SendMessage) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *SendMessage) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetVariables

`func (o *SendMessage) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *SendMessage) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *SendMessage) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *SendMessage) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetSendAt

`func (o *SendMessage) GetSendAt() time.Time`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *SendMessage) GetSendAtOk() (*time.Time, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *SendMessage) SetSendAt(v time.Time)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *SendMessage) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.

### GetAttachments

`func (o *SendMessage) GetAttachments() []V1SendPostRequestAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SendMessage) GetAttachmentsOk() (*[]V1SendPostRequestAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SendMessage) SetAttachments(v []V1SendPostRequestAttachmentsInner)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SendMessage) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


