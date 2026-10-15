# StartConversationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **[]string** |  | 
**Cc** | Pointer to **[]string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Text** | **string** |  | 
**Html** | Pointer to **string** |  | [optional] 

## Methods

### NewStartConversationRequest

`func NewStartConversationRequest(to []string, text string, ) *StartConversationRequest`

NewStartConversationRequest instantiates a new StartConversationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartConversationRequestWithDefaults

`func NewStartConversationRequestWithDefaults() *StartConversationRequest`

NewStartConversationRequestWithDefaults instantiates a new StartConversationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *StartConversationRequest) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *StartConversationRequest) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *StartConversationRequest) SetTo(v []string)`

SetTo sets To field to given value.


### GetCc

`func (o *StartConversationRequest) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *StartConversationRequest) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *StartConversationRequest) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *StartConversationRequest) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetSubject

`func (o *StartConversationRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *StartConversationRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *StartConversationRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *StartConversationRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *StartConversationRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StartConversationRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StartConversationRequest) SetText(v string)`

SetText sets Text field to given value.


### GetHtml

`func (o *StartConversationRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *StartConversationRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *StartConversationRequest) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *StartConversationRequest) HasHtml() bool`

HasHtml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


