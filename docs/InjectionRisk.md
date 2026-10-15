# InjectionRisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flagged** | Pointer to **bool** |  | [optional] 
**Signals** | Pointer to **[]string** | The phrases that tripped the check. | [optional] 

## Methods

### NewInjectionRisk

`func NewInjectionRisk() *InjectionRisk`

NewInjectionRisk instantiates a new InjectionRisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInjectionRiskWithDefaults

`func NewInjectionRiskWithDefaults() *InjectionRisk`

NewInjectionRiskWithDefaults instantiates a new InjectionRisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlagged

`func (o *InjectionRisk) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *InjectionRisk) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *InjectionRisk) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *InjectionRisk) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### GetSignals

`func (o *InjectionRisk) GetSignals() []string`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *InjectionRisk) GetSignalsOk() (*[]string, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *InjectionRisk) SetSignals(v []string)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *InjectionRisk) HasSignals() bool`

HasSignals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


