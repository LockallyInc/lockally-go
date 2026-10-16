# ExportMailbox201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Relative download URL. Caller appends to the API host. | 
**ExpiresAt** | **time.Time** |  | 
**Format** | **string** |  | 

## Methods

### NewExportMailbox201Response

`func NewExportMailbox201Response(url string, expiresAt time.Time, format string, ) *ExportMailbox201Response`

NewExportMailbox201Response instantiates a new ExportMailbox201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportMailbox201ResponseWithDefaults

`func NewExportMailbox201ResponseWithDefaults() *ExportMailbox201Response`

NewExportMailbox201ResponseWithDefaults instantiates a new ExportMailbox201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ExportMailbox201Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ExportMailbox201Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ExportMailbox201Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetExpiresAt

`func (o *ExportMailbox201Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ExportMailbox201Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ExportMailbox201Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetFormat

`func (o *ExportMailbox201Response) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ExportMailbox201Response) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ExportMailbox201Response) SetFormat(v string)`

SetFormat sets Format field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


