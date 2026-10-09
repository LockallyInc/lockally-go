# V1WebhooksIdPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **[]string** |  | [optional] 
**Paused** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1WebhooksIdPatchRequest

`func NewV1WebhooksIdPatchRequest() *V1WebhooksIdPatchRequest`

NewV1WebhooksIdPatchRequest instantiates a new V1WebhooksIdPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1WebhooksIdPatchRequestWithDefaults

`func NewV1WebhooksIdPatchRequestWithDefaults() *V1WebhooksIdPatchRequest`

NewV1WebhooksIdPatchRequestWithDefaults instantiates a new V1WebhooksIdPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *V1WebhooksIdPatchRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1WebhooksIdPatchRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1WebhooksIdPatchRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V1WebhooksIdPatchRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEvents

`func (o *V1WebhooksIdPatchRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *V1WebhooksIdPatchRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *V1WebhooksIdPatchRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *V1WebhooksIdPatchRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetPaused

`func (o *V1WebhooksIdPatchRequest) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *V1WebhooksIdPatchRequest) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *V1WebhooksIdPatchRequest) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *V1WebhooksIdPatchRequest) HasPaused() bool`

HasPaused returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


