# V1InboxesMailboxMessagesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **[]string** |  | 
**Cc** | Pointer to **[]string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Text** | **string** |  | 
**Html** | Pointer to **string** |  | [optional] 

## Methods

### NewV1InboxesMailboxMessagesPostRequest

`func NewV1InboxesMailboxMessagesPostRequest(to []string, text string, ) *V1InboxesMailboxMessagesPostRequest`

NewV1InboxesMailboxMessagesPostRequest instantiates a new V1InboxesMailboxMessagesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InboxesMailboxMessagesPostRequestWithDefaults

`func NewV1InboxesMailboxMessagesPostRequestWithDefaults() *V1InboxesMailboxMessagesPostRequest`

NewV1InboxesMailboxMessagesPostRequestWithDefaults instantiates a new V1InboxesMailboxMessagesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *V1InboxesMailboxMessagesPostRequest) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V1InboxesMailboxMessagesPostRequest) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V1InboxesMailboxMessagesPostRequest) SetTo(v []string)`

SetTo sets To field to given value.


### GetCc

`func (o *V1InboxesMailboxMessagesPostRequest) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *V1InboxesMailboxMessagesPostRequest) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *V1InboxesMailboxMessagesPostRequest) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *V1InboxesMailboxMessagesPostRequest) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetSubject

`func (o *V1InboxesMailboxMessagesPostRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *V1InboxesMailboxMessagesPostRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *V1InboxesMailboxMessagesPostRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *V1InboxesMailboxMessagesPostRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *V1InboxesMailboxMessagesPostRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *V1InboxesMailboxMessagesPostRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *V1InboxesMailboxMessagesPostRequest) SetText(v string)`

SetText sets Text field to given value.


### GetHtml

`func (o *V1InboxesMailboxMessagesPostRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *V1InboxesMailboxMessagesPostRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *V1InboxesMailboxMessagesPostRequest) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *V1InboxesMailboxMessagesPostRequest) HasHtml() bool`

HasHtml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


