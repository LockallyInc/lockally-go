# V1ApiKeysPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Free-text identifier shown in the dashboard. | 
**Scopes** | **[]string** | Allowed scopes on this key. | 

## Methods

### NewV1ApiKeysPostRequest

`func NewV1ApiKeysPostRequest(label string, scopes []string, ) *V1ApiKeysPostRequest`

NewV1ApiKeysPostRequest instantiates a new V1ApiKeysPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ApiKeysPostRequestWithDefaults

`func NewV1ApiKeysPostRequestWithDefaults() *V1ApiKeysPostRequest`

NewV1ApiKeysPostRequestWithDefaults instantiates a new V1ApiKeysPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *V1ApiKeysPostRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *V1ApiKeysPostRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *V1ApiKeysPostRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetScopes

`func (o *V1ApiKeysPostRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *V1ApiKeysPostRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *V1ApiKeysPostRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


