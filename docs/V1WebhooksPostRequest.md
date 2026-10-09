# V1WebhooksPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | HTTPS recommended (HTTP allowed for local development). | 
**Events** | **[]string** |  | 

## Methods

### NewV1WebhooksPostRequest

`func NewV1WebhooksPostRequest(url string, events []string, ) *V1WebhooksPostRequest`

NewV1WebhooksPostRequest instantiates a new V1WebhooksPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1WebhooksPostRequestWithDefaults

`func NewV1WebhooksPostRequestWithDefaults() *V1WebhooksPostRequest`

NewV1WebhooksPostRequestWithDefaults instantiates a new V1WebhooksPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *V1WebhooksPostRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1WebhooksPostRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1WebhooksPostRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEvents

`func (o *V1WebhooksPostRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *V1WebhooksPostRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *V1WebhooksPostRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


