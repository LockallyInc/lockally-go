# SendEmail202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Lockally identifier; use with GET /v1/messages/{id}. | 
**MessageId** | **string** | RFC 5322 Message-ID (with angle brackets). | 
**Status** | **string** | \&quot;scheduled\&quot; when send_at is in the future. | 
**Warning** | Pointer to **string** |  | [optional] 

## Methods

### NewSendEmail202Response

`func NewSendEmail202Response(id string, messageId string, status string, ) *SendEmail202Response`

NewSendEmail202Response instantiates a new SendEmail202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmail202ResponseWithDefaults

`func NewSendEmail202ResponseWithDefaults() *SendEmail202Response`

NewSendEmail202ResponseWithDefaults instantiates a new SendEmail202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SendEmail202Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SendEmail202Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SendEmail202Response) SetId(v string)`

SetId sets Id field to given value.


### GetMessageId

`func (o *SendEmail202Response) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SendEmail202Response) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SendEmail202Response) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetStatus

`func (o *SendEmail202Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SendEmail202Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SendEmail202Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetWarning

`func (o *SendEmail202Response) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *SendEmail202Response) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *SendEmail202Response) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *SendEmail202Response) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


