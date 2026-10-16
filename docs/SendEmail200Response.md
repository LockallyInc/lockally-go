# SendEmail200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Present on a test-mode send; absent when suppressed. | [optional] 
**MessageId** | Pointer to **string** | RFC 5322 Message-ID. On a test send the domain is test.lockally.com. Absent when suppressed. | [optional] 
**Status** | **string** |  | 
**Warning** | Pointer to **string** | On a suppressed send, explains that nothing was sent. | [optional] 

## Methods

### NewSendEmail200Response

`func NewSendEmail200Response(status string, ) *SendEmail200Response`

NewSendEmail200Response instantiates a new SendEmail200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmail200ResponseWithDefaults

`func NewSendEmail200ResponseWithDefaults() *SendEmail200Response`

NewSendEmail200ResponseWithDefaults instantiates a new SendEmail200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SendEmail200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SendEmail200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SendEmail200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SendEmail200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessageId

`func (o *SendEmail200Response) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SendEmail200Response) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SendEmail200Response) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *SendEmail200Response) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetStatus

`func (o *SendEmail200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SendEmail200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SendEmail200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetWarning

`func (o *SendEmail200Response) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *SendEmail200Response) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *SendEmail200Response) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *SendEmail200Response) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


