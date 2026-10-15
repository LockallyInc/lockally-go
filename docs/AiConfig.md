# AiConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**Model** | **string** |  | 
**ByokKeyHint** | Pointer to **string** |  | [optional] 
**UnitsBalance** | Pointer to **int32** |  | [optional] 
**UnitsAvailable** | Pointer to **bool** | True when a platform model key is configured. | [optional] 

## Methods

### NewAiConfig

`func NewAiConfig(mode string, model string, ) *AiConfig`

NewAiConfig instantiates a new AiConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiConfigWithDefaults

`func NewAiConfigWithDefaults() *AiConfig`

NewAiConfigWithDefaults instantiates a new AiConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *AiConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AiConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AiConfig) SetMode(v string)`

SetMode sets Mode field to given value.


### GetModel

`func (o *AiConfig) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiConfig) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiConfig) SetModel(v string)`

SetModel sets Model field to given value.


### GetByokKeyHint

`func (o *AiConfig) GetByokKeyHint() string`

GetByokKeyHint returns the ByokKeyHint field if non-nil, zero value otherwise.

### GetByokKeyHintOk

`func (o *AiConfig) GetByokKeyHintOk() (*string, bool)`

GetByokKeyHintOk returns a tuple with the ByokKeyHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByokKeyHint

`func (o *AiConfig) SetByokKeyHint(v string)`

SetByokKeyHint sets ByokKeyHint field to given value.

### HasByokKeyHint

`func (o *AiConfig) HasByokKeyHint() bool`

HasByokKeyHint returns a boolean if a field has been set.

### GetUnitsBalance

`func (o *AiConfig) GetUnitsBalance() int32`

GetUnitsBalance returns the UnitsBalance field if non-nil, zero value otherwise.

### GetUnitsBalanceOk

`func (o *AiConfig) GetUnitsBalanceOk() (*int32, bool)`

GetUnitsBalanceOk returns a tuple with the UnitsBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsBalance

`func (o *AiConfig) SetUnitsBalance(v int32)`

SetUnitsBalance sets UnitsBalance field to given value.

### HasUnitsBalance

`func (o *AiConfig) HasUnitsBalance() bool`

HasUnitsBalance returns a boolean if a field has been set.

### GetUnitsAvailable

`func (o *AiConfig) GetUnitsAvailable() bool`

GetUnitsAvailable returns the UnitsAvailable field if non-nil, zero value otherwise.

### GetUnitsAvailableOk

`func (o *AiConfig) GetUnitsAvailableOk() (*bool, bool)`

GetUnitsAvailableOk returns a tuple with the UnitsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsAvailable

`func (o *AiConfig) SetUnitsAvailable(v bool)`

SetUnitsAvailable sets UnitsAvailable field to given value.

### HasUnitsAvailable

`func (o *AiConfig) HasUnitsAvailable() bool`

HasUnitsAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


