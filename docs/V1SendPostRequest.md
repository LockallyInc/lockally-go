# V1SendPostRequest

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
**Unsubscribe** | Pointer to **bool** | Mark as opt-in/broadcast: skips suppressed recipients and adds a managed one-click List-Unsubscribe header. | [optional] 
**TemplateId** | Pointer to **string** | Render subject/text/html from a stored template (GET /v1/templates). Mutually exclusive with inline subject/text/html. | [optional] 
**Variables** | Pointer to **map[string]string** | Values substituted into the template&#39;s {{variable}} placeholders. | [optional] 
**SendAt** | Pointer to **time.Time** | Schedule delivery for a future RFC3339 time (≤ 30 days out). Omit or past &#x3D; send now. Cancel with DELETE /v1/messages/{id} while scheduled. | [optional] 
**Attachments** | Pointer to [**[]V1SendPostRequestAttachmentsInner**](V1SendPostRequestAttachmentsInner.md) |  | [optional] 

## Methods

### NewV1SendPostRequest

`func NewV1SendPostRequest(from string, to []string, ) *V1SendPostRequest`

NewV1SendPostRequest instantiates a new V1SendPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SendPostRequestWithDefaults

`func NewV1SendPostRequestWithDefaults() *V1SendPostRequest`

NewV1SendPostRequestWithDefaults instantiates a new V1SendPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *V1SendPostRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1SendPostRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1SendPostRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *V1SendPostRequest) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V1SendPostRequest) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V1SendPostRequest) SetTo(v []string)`

SetTo sets To field to given value.


### GetCc

`func (o *V1SendPostRequest) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *V1SendPostRequest) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *V1SendPostRequest) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *V1SendPostRequest) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *V1SendPostRequest) GetBcc() []string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *V1SendPostRequest) GetBccOk() (*[]string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *V1SendPostRequest) SetBcc(v []string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *V1SendPostRequest) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetSubject

`func (o *V1SendPostRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *V1SendPostRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *V1SendPostRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *V1SendPostRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *V1SendPostRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *V1SendPostRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *V1SendPostRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *V1SendPostRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetHtml

`func (o *V1SendPostRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *V1SendPostRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *V1SendPostRequest) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *V1SendPostRequest) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetHeaders

`func (o *V1SendPostRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *V1SendPostRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *V1SendPostRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *V1SendPostRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetUnsubscribe

`func (o *V1SendPostRequest) GetUnsubscribe() bool`

GetUnsubscribe returns the Unsubscribe field if non-nil, zero value otherwise.

### GetUnsubscribeOk

`func (o *V1SendPostRequest) GetUnsubscribeOk() (*bool, bool)`

GetUnsubscribeOk returns a tuple with the Unsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribe

`func (o *V1SendPostRequest) SetUnsubscribe(v bool)`

SetUnsubscribe sets Unsubscribe field to given value.

### HasUnsubscribe

`func (o *V1SendPostRequest) HasUnsubscribe() bool`

HasUnsubscribe returns a boolean if a field has been set.

### GetTemplateId

`func (o *V1SendPostRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *V1SendPostRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *V1SendPostRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *V1SendPostRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetVariables

`func (o *V1SendPostRequest) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *V1SendPostRequest) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *V1SendPostRequest) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *V1SendPostRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetSendAt

`func (o *V1SendPostRequest) GetSendAt() time.Time`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *V1SendPostRequest) GetSendAtOk() (*time.Time, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *V1SendPostRequest) SetSendAt(v time.Time)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *V1SendPostRequest) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.

### GetAttachments

`func (o *V1SendPostRequest) GetAttachments() []V1SendPostRequestAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *V1SendPostRequest) GetAttachmentsOk() (*[]V1SendPostRequestAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *V1SendPostRequest) SetAttachments(v []V1SendPostRequestAttachmentsInner)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *V1SendPostRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


