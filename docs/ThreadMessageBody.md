# ThreadMessageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Direction** | **string** |  | 
**RfcMessageId** | **string** | The Internet Message-ID | 
**InReplyTo** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**From** | **string** |  | 
**To** | **[]string** |  | 
**Snippet** | Pointer to **string** |  | [optional] 
**HasAttachments** | **bool** |  | 
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Signals computed at ingest | [optional] 
**ReceivedAt** | **time.Time** |  | 
**Body** | Pointer to [**ThreadMessageBodyAllOfBody**](ThreadMessageBodyAllOfBody.md) |  | [optional] 

## Methods

### NewThreadMessageBody

`func NewThreadMessageBody(id string, direction string, rfcMessageId string, from string, to []string, hasAttachments bool, receivedAt time.Time, ) *ThreadMessageBody`

NewThreadMessageBody instantiates a new ThreadMessageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadMessageBodyWithDefaults

`func NewThreadMessageBodyWithDefaults() *ThreadMessageBody`

NewThreadMessageBodyWithDefaults instantiates a new ThreadMessageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThreadMessageBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreadMessageBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreadMessageBody) SetId(v string)`

SetId sets Id field to given value.


### GetDirection

`func (o *ThreadMessageBody) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ThreadMessageBody) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ThreadMessageBody) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetRfcMessageId

`func (o *ThreadMessageBody) GetRfcMessageId() string`

GetRfcMessageId returns the RfcMessageId field if non-nil, zero value otherwise.

### GetRfcMessageIdOk

`func (o *ThreadMessageBody) GetRfcMessageIdOk() (*string, bool)`

GetRfcMessageIdOk returns a tuple with the RfcMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfcMessageId

`func (o *ThreadMessageBody) SetRfcMessageId(v string)`

SetRfcMessageId sets RfcMessageId field to given value.


### GetInReplyTo

`func (o *ThreadMessageBody) GetInReplyTo() string`

GetInReplyTo returns the InReplyTo field if non-nil, zero value otherwise.

### GetInReplyToOk

`func (o *ThreadMessageBody) GetInReplyToOk() (*string, bool)`

GetInReplyToOk returns a tuple with the InReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyTo

`func (o *ThreadMessageBody) SetInReplyTo(v string)`

SetInReplyTo sets InReplyTo field to given value.

### HasInReplyTo

`func (o *ThreadMessageBody) HasInReplyTo() bool`

HasInReplyTo returns a boolean if a field has been set.

### GetSubject

`func (o *ThreadMessageBody) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ThreadMessageBody) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ThreadMessageBody) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ThreadMessageBody) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetFrom

`func (o *ThreadMessageBody) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ThreadMessageBody) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ThreadMessageBody) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *ThreadMessageBody) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ThreadMessageBody) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ThreadMessageBody) SetTo(v []string)`

SetTo sets To field to given value.


### GetSnippet

`func (o *ThreadMessageBody) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *ThreadMessageBody) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *ThreadMessageBody) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.

### HasSnippet

`func (o *ThreadMessageBody) HasSnippet() bool`

HasSnippet returns a boolean if a field has been set.

### GetHasAttachments

`func (o *ThreadMessageBody) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *ThreadMessageBody) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *ThreadMessageBody) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.


### GetAttachments

`func (o *ThreadMessageBody) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ThreadMessageBody) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ThreadMessageBody) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ThreadMessageBody) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetAnnotations

`func (o *ThreadMessageBody) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ThreadMessageBody) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ThreadMessageBody) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ThreadMessageBody) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetReceivedAt

`func (o *ThreadMessageBody) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *ThreadMessageBody) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *ThreadMessageBody) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.


### GetBody

`func (o *ThreadMessageBody) GetBody() ThreadMessageBodyAllOfBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ThreadMessageBody) GetBodyOk() (*ThreadMessageBodyAllOfBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ThreadMessageBody) SetBody(v ThreadMessageBodyAllOfBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *ThreadMessageBody) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


