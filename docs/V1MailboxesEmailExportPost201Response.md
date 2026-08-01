# V1MailboxesEmailExportPost201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Relative download URL. Caller appends to the API host. | 
**ExpiresAt** | **time.Time** |  | 
**Format** | **string** |  | 

## Methods

### NewV1MailboxesEmailExportPost201Response

`func NewV1MailboxesEmailExportPost201Response(url string, expiresAt time.Time, format string, ) *V1MailboxesEmailExportPost201Response`

NewV1MailboxesEmailExportPost201Response instantiates a new V1MailboxesEmailExportPost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MailboxesEmailExportPost201ResponseWithDefaults

`func NewV1MailboxesEmailExportPost201ResponseWithDefaults() *V1MailboxesEmailExportPost201Response`

NewV1MailboxesEmailExportPost201ResponseWithDefaults instantiates a new V1MailboxesEmailExportPost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *V1MailboxesEmailExportPost201Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1MailboxesEmailExportPost201Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1MailboxesEmailExportPost201Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetExpiresAt

`func (o *V1MailboxesEmailExportPost201Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *V1MailboxesEmailExportPost201Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *V1MailboxesEmailExportPost201Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetFormat

`func (o *V1MailboxesEmailExportPost201Response) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V1MailboxesEmailExportPost201Response) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V1MailboxesEmailExportPost201Response) SetFormat(v string)`

SetFormat sets Format field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


