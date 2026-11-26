# SendEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | Sender address, alone (&#x60;orders@acme.com&#x60;) or with a display name (&#x60;Acme Orders &lt;orders@acme.com&gt;&#x60;). Only the address is checked against your sender rules and used as the envelope sender. | 
**FromName** | Pointer to **string** | Display name recipients see next to &#x60;from&#x60;, e.g. &#x60;Acme Orders&#x60;. The same as giving it in &#x60;from&#x60;; if you do both, the names must match (400 &#x60;Conflicting from name&#x60;). Cannot contain &#x60;@&#x60; or control characters. Encoded and quoted for you, so accents, commas and quotes are safe. | [optional] 
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
**Metadata** | Pointer to **map[string]string** | Arbitrary key/value you attach to this message and get back on every delivery webhook for it, so a send can be matched to your own records without a lookup table.  **It never leaves Lockally.** A custom entry in &#x60;headers&#x60; goes out on the wire to the recipient, which is why that is not a substitute: order ids and internal identifiers do not belong in a header the recipient can read.  At most 20 keys, keys at most 64 characters, values at most 512. **Over a limit is refused with a &#x60;400&#x60;, never truncated**: a silently shortened order id succeeds and then fails to match weeks later.  Not indexed and not filterable.  | [optional] 
**Tags** | Pointer to **[]string** | Labels that group this send for filtering and analytics: campaign, template, environment.  **Lowercased, trimmed and deduplicated on write**, so the grouping does not depend on every caller spelling them identically. &#x60;Campaign-X&#x60; is stored as &#x60;campaign-x&#x60;, and either spelling matches it in a filter.  At most 10 per message after deduplication, each at most 32 characters. Over either is a &#x60;400&#x60; rather than a trimmed list: a send missing the tag it was tagged for is worse than a refusal.  Unlike &#x60;metadata&#x60;, tags ARE queryable: filter with &#x60;GET /v1/messages?tag&#x3D;&#x60; and group with &#x60;GET /v1/messages/stats?tag&#x3D;&#x60;.  | [optional] 
**TrackOpens** | Pointer to **bool** | Override the account&#39;s open-tracking setting for this one send. Omit to inherit it.  Three states matter here: &#x60;true&#x60;, &#x60;false&#x60;, and absent. A transactional receipt can set &#x60;false&#x60; without the account turning tracking off for its marketing.  | [optional] 
**TrackClicks** | Pointer to **bool** | As &#x60;track_opens&#x60;, for click tracking. Click tracking rewrites &#x60;http(s)&#x60; links in the **HTML** part only: the plain-text part is never rewritten, because a visible redirect URL is what recipients complain about. A text-only send is therefore not click-tracked.  | [optional] 
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


### GetFromName

`func (o *SendEmailRequest) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *SendEmailRequest) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *SendEmailRequest) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *SendEmailRequest) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

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

### GetMetadata

`func (o *SendEmailRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SendEmailRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SendEmailRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SendEmailRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTags

`func (o *SendEmailRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SendEmailRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SendEmailRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SendEmailRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTrackOpens

`func (o *SendEmailRequest) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *SendEmailRequest) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *SendEmailRequest) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.

### HasTrackOpens

`func (o *SendEmailRequest) HasTrackOpens() bool`

HasTrackOpens returns a boolean if a field has been set.

### GetTrackClicks

`func (o *SendEmailRequest) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *SendEmailRequest) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *SendEmailRequest) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.

### HasTrackClicks

`func (o *SendEmailRequest) HasTrackClicks() bool`

HasTrackClicks returns a boolean if a field has been set.

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


