# CreateApiKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Free-text identifier shown in the dashboard. | 
**Scopes** | **[]string** | Allowed scopes on this key. | 
**Test** | Pointer to **bool** | Issue a test key (&#x60;lk_test_&#x60; prefix). A test key accepts sends, records them and fires a simulated delivery webhook without sending mail, and counts against no limit or allowance.  | [optional] [default to false]

## Methods

### NewCreateApiKeyRequest

`func NewCreateApiKeyRequest(label string, scopes []string, ) *CreateApiKeyRequest`

NewCreateApiKeyRequest instantiates a new CreateApiKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKeyRequestWithDefaults

`func NewCreateApiKeyRequestWithDefaults() *CreateApiKeyRequest`

NewCreateApiKeyRequestWithDefaults instantiates a new CreateApiKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateApiKeyRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateApiKeyRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateApiKeyRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetScopes

`func (o *CreateApiKeyRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateApiKeyRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateApiKeyRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetTest

`func (o *CreateApiKeyRequest) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CreateApiKeyRequest) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CreateApiKeyRequest) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *CreateApiKeyRequest) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


