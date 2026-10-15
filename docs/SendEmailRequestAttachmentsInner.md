# SendEmailRequestAttachmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** |  | 
**ContentType** | **string** |  | 
**ContentBase64** | **string** | Base64-encoded body, max 10 MB decoded per attachment. | 

## Methods

### NewSendEmailRequestAttachmentsInner

`func NewSendEmailRequestAttachmentsInner(filename string, contentType string, contentBase64 string, ) *SendEmailRequestAttachmentsInner`

NewSendEmailRequestAttachmentsInner instantiates a new SendEmailRequestAttachmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailRequestAttachmentsInnerWithDefaults

`func NewSendEmailRequestAttachmentsInnerWithDefaults() *SendEmailRequestAttachmentsInner`

NewSendEmailRequestAttachmentsInnerWithDefaults instantiates a new SendEmailRequestAttachmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *SendEmailRequestAttachmentsInner) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SendEmailRequestAttachmentsInner) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SendEmailRequestAttachmentsInner) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *SendEmailRequestAttachmentsInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SendEmailRequestAttachmentsInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SendEmailRequestAttachmentsInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetContentBase64

`func (o *SendEmailRequestAttachmentsInner) GetContentBase64() string`

GetContentBase64 returns the ContentBase64 field if non-nil, zero value otherwise.

### GetContentBase64Ok

`func (o *SendEmailRequestAttachmentsInner) GetContentBase64Ok() (*string, bool)`

GetContentBase64Ok returns a tuple with the ContentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBase64

`func (o *SendEmailRequestAttachmentsInner) SetContentBase64(v string)`

SetContentBase64 sets ContentBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


