# CreateAliasRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasAddress** | **string** |  | 
**AliasTarget** | **string** |  | 

## Methods

### NewCreateAliasRequest

`func NewCreateAliasRequest(aliasAddress string, aliasTarget string, ) *CreateAliasRequest`

NewCreateAliasRequest instantiates a new CreateAliasRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAliasRequestWithDefaults

`func NewCreateAliasRequestWithDefaults() *CreateAliasRequest`

NewCreateAliasRequestWithDefaults instantiates a new CreateAliasRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasAddress

`func (o *CreateAliasRequest) GetAliasAddress() string`

GetAliasAddress returns the AliasAddress field if non-nil, zero value otherwise.

### GetAliasAddressOk

`func (o *CreateAliasRequest) GetAliasAddressOk() (*string, bool)`

GetAliasAddressOk returns a tuple with the AliasAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasAddress

`func (o *CreateAliasRequest) SetAliasAddress(v string)`

SetAliasAddress sets AliasAddress field to given value.


### GetAliasTarget

`func (o *CreateAliasRequest) GetAliasTarget() string`

GetAliasTarget returns the AliasTarget field if non-nil, zero value otherwise.

### GetAliasTargetOk

`func (o *CreateAliasRequest) GetAliasTargetOk() (*string, bool)`

GetAliasTargetOk returns a tuple with the AliasTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTarget

`func (o *CreateAliasRequest) SetAliasTarget(v string)`

SetAliasTarget sets AliasTarget field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


