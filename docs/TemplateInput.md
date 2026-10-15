# TemplateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Template name (1–120 chars). | 
**Subject** | Pointer to **string** | May contain {{variable}} placeholders. | [optional] 
**Html** | Pointer to **string** | HTML body; {{variable}} values are HTML-escaped. | [optional] 
**Text** | Pointer to **string** | Plain-text body. | [optional] 

## Methods

### NewTemplateInput

`func NewTemplateInput(name string, ) *TemplateInput`

NewTemplateInput instantiates a new TemplateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateInputWithDefaults

`func NewTemplateInputWithDefaults() *TemplateInput`

NewTemplateInputWithDefaults instantiates a new TemplateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateInput) SetName(v string)`

SetName sets Name field to given value.


### GetSubject

`func (o *TemplateInput) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TemplateInput) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TemplateInput) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TemplateInput) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetHtml

`func (o *TemplateInput) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *TemplateInput) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *TemplateInput) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *TemplateInput) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetText

`func (o *TemplateInput) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TemplateInput) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TemplateInput) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TemplateInput) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


