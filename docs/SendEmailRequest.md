# SendEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** |  | 
**To** | **[]string** |  | 
**Cc** | Pointer to **[]string** |  | [optional] 
**Bcc** | Pointer to **[]string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** | Plain-text body. Required if &#x60;html&#x60; is absent. | [optional] 
**Html** | Pointer to **string** | HTML body. Required if &#x60;text&#x60; is absent. | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**ReplyTo** | Pointer to **string** | Where replies go. REQUIRED when &#x60;from&#x60; has no mailbox behind it, since a reply there would be delivered nowhere. A &#x60;Reply-To&#x60; entry in &#x60;headers&#x60; satisfies the same rule. The named form &#x60;Acme Support &lt;a@b.com&gt;&#x60; is also accepted. | [optional] 
**ReplyToName** | Pointer to **string** | Display name shown on &#x60;Reply-To&#x60;, e.g. &#x60;Acme Support&#x60;. Ignored without &#x60;reply_to&#x60;. Encoded and quoted for you, so accents, commas and quotes are all safe. | [optional] 
**Unsubscribe** | Pointer to **bool** | Mark as opt-in/broadcast: skips suppressed recipients and adds a managed one-click List-Unsubscribe header. | [optional] 
**TemplateId** | Pointer to **string** | Render subject/text/html from a stored template (GET /v1/templates). Mutually exclusive with inline subject/text/html. | [optional] 
**Variables** | Pointer to **map[string]string** | Values substituted into the template&#39;s {{variable}} placeholders. | [optional] 
**SendAt** | Pointer to **time.Time** | Schedule delivery for a future RFC3339 time (≤ 30 days out). Omit or past &#x3D; send now. Cancel with DELETE /v1/messages/{id} while scheduled. | [optional] 
**Attachments** | Pointer to [**[]SendEmailRequestAttachmentsInner**](SendEmailRequestAttachmentsInner.md) |  | [optional] 

## Methods

### NewSendEmailRequest

`func NewSendEmailRequest(from string, to []string, ) *SendEmailRequest`

NewSendEmailRequest instantiates a new SendEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailRequestWithDefaults

`func NewSendEmailRequestWithDefaults() *SendEmailRequest`

NewSendEmailRequestWithDefaults instantiates a new SendEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *SendEmailRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SendEmailRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SendEmailRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *SendEmailRequest) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SendEmailRequest) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SendEmailRequest) SetTo(v []string)`

SetTo sets To field to given value.


### GetCc

`func (o *SendEmailRequest) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *SendEmailRequest) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *SendEmailRequest) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *SendEmailRequest) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *SendEmailRequest) GetBcc() []string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *SendEmailRequest) GetBccOk() (*[]string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *SendEmailRequest) SetBcc(v []string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *SendEmailRequest) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetSubject

`func (o *SendEmailRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SendEmailRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SendEmailRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SendEmailRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *SendEmailRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SendEmailRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SendEmailRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SendEmailRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetHtml

`func (o *SendEmailRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *SendEmailRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *SendEmailRequest) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *SendEmailRequest) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetHeaders

`func (o *SendEmailRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SendEmailRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *SendEmailRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *SendEmailRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetReplyTo

`func (o *SendEmailRequest) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *SendEmailRequest) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *SendEmailRequest) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *SendEmailRequest) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetReplyToName

`func (o *SendEmailRequest) GetReplyToName() string`

GetReplyToName returns the ReplyToName field if non-nil, zero value otherwise.

### GetReplyToNameOk

`func (o *SendEmailRequest) GetReplyToNameOk() (*string, bool)`

GetReplyToNameOk returns a tuple with the ReplyToName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToName

`func (o *SendEmailRequest) SetReplyToName(v string)`

SetReplyToName sets ReplyToName field to given value.

### HasReplyToName

`func (o *SendEmailRequest) HasReplyToName() bool`

HasReplyToName returns a boolean if a field has been set.

### GetUnsubscribe

`func (o *SendEmailRequest) GetUnsubscribe() bool`

GetUnsubscribe returns the Unsubscribe field if non-nil, zero value otherwise.

### GetUnsubscribeOk

`func (o *SendEmailRequest) GetUnsubscribeOk() (*bool, bool)`

GetUnsubscribeOk returns a tuple with the Unsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribe

`func (o *SendEmailRequest) SetUnsubscribe(v bool)`

SetUnsubscribe sets Unsubscribe field to given value.

### HasUnsubscribe

`func (o *SendEmailRequest) HasUnsubscribe() bool`

HasUnsubscribe returns a boolean if a field has been set.

### GetTemplateId

`func (o *SendEmailRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *SendEmailRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *SendEmailRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *SendEmailRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetVariables

`func (o *SendEmailRequest) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *SendEmailRequest) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *SendEmailRequest) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *SendEmailRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetSendAt

`func (o *SendEmailRequest) GetSendAt() time.Time`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *SendEmailRequest) GetSendAtOk() (*time.Time, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *SendEmailRequest) SetSendAt(v time.Time)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *SendEmailRequest) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.

### GetAttachments

`func (o *SendEmailRequest) GetAttachments() []SendEmailRequestAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SendEmailRequest) GetAttachmentsOk() (*[]SendEmailRequestAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SendEmailRequest) SetAttachments(v []SendEmailRequestAttachmentsInner)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SendEmailRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


