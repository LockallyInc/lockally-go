# V1SendPostRequestAttachmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** |  | 
**ContentType** | **string** |  | 
**ContentBase64** | **string** | Base64-encoded body, max 10 MB decoded per attachment. | 

## Methods

### NewV1SendPostRequestAttachmentsInner

`func NewV1SendPostRequestAttachmentsInner(filename string, contentType string, contentBase64 string, ) *V1SendPostRequestAttachmentsInner`

NewV1SendPostRequestAttachmentsInner instantiates a new V1SendPostRequestAttachmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SendPostRequestAttachmentsInnerWithDefaults

`func NewV1SendPostRequestAttachmentsInnerWithDefaults() *V1SendPostRequestAttachmentsInner`

NewV1SendPostRequestAttachmentsInnerWithDefaults instantiates a new V1SendPostRequestAttachmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *V1SendPostRequestAttachmentsInner) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *V1SendPostRequestAttachmentsInner) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *V1SendPostRequestAttachmentsInner) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *V1SendPostRequestAttachmentsInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *V1SendPostRequestAttachmentsInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *V1SendPostRequestAttachmentsInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetContentBase64

`func (o *V1SendPostRequestAttachmentsInner) GetContentBase64() string`

GetContentBase64 returns the ContentBase64 field if non-nil, zero value otherwise.

### GetContentBase64Ok

`func (o *V1SendPostRequestAttachmentsInner) GetContentBase64Ok() (*string, bool)`

GetContentBase64Ok returns a tuple with the ContentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBase64

`func (o *V1SendPostRequestAttachmentsInner) SetContentBase64(v string)`

SetContentBase64 sets ContentBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


