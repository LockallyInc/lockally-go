# GetStorage200ResponseTopMessagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**SizeBytes** | Pointer to **int32** |  | [optional] 
**Filenames** | Pointer to **[]string** |  | [optional] 
**QueuedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetStorage200ResponseTopMessagesInner

`func NewGetStorage200ResponseTopMessagesInner() *GetStorage200ResponseTopMessagesInner`

NewGetStorage200ResponseTopMessagesInner instantiates a new GetStorage200ResponseTopMessagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorage200ResponseTopMessagesInnerWithDefaults

`func NewGetStorage200ResponseTopMessagesInnerWithDefaults() *GetStorage200ResponseTopMessagesInner`

NewGetStorage200ResponseTopMessagesInnerWithDefaults instantiates a new GetStorage200ResponseTopMessagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *GetStorage200ResponseTopMessagesInner) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *GetStorage200ResponseTopMessagesInner) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *GetStorage200ResponseTopMessagesInner) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *GetStorage200ResponseTopMessagesInner) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetSubject

`func (o *GetStorage200ResponseTopMessagesInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetStorage200ResponseTopMessagesInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetStorage200ResponseTopMessagesInner) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetStorage200ResponseTopMessagesInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSizeBytes

`func (o *GetStorage200ResponseTopMessagesInner) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *GetStorage200ResponseTopMessagesInner) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *GetStorage200ResponseTopMessagesInner) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *GetStorage200ResponseTopMessagesInner) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetFilenames

`func (o *GetStorage200ResponseTopMessagesInner) GetFilenames() []string`

GetFilenames returns the Filenames field if non-nil, zero value otherwise.

### GetFilenamesOk

`func (o *GetStorage200ResponseTopMessagesInner) GetFilenamesOk() (*[]string, bool)`

GetFilenamesOk returns a tuple with the Filenames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenames

`func (o *GetStorage200ResponseTopMessagesInner) SetFilenames(v []string)`

SetFilenames sets Filenames field to given value.

### HasFilenames

`func (o *GetStorage200ResponseTopMessagesInner) HasFilenames() bool`

HasFilenames returns a boolean if a field has been set.

### GetQueuedAt

`func (o *GetStorage200ResponseTopMessagesInner) GetQueuedAt() time.Time`

GetQueuedAt returns the QueuedAt field if non-nil, zero value otherwise.

### GetQueuedAtOk

`func (o *GetStorage200ResponseTopMessagesInner) GetQueuedAtOk() (*time.Time, bool)`

GetQueuedAtOk returns a tuple with the QueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedAt

`func (o *GetStorage200ResponseTopMessagesInner) SetQueuedAt(v time.Time)`

SetQueuedAt sets QueuedAt field to given value.

### HasQueuedAt

`func (o *GetStorage200ResponseTopMessagesInner) HasQueuedAt() bool`

HasQueuedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


