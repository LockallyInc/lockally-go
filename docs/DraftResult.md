# DraftResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DraftId** | **string** |  | 
**Status** | **string** |  | 
**PolicyFlags** | **map[string]interface{}** |  | 
**ThreadId** | Pointer to **string** |  | [optional] 
**Sent** | Pointer to [**AgentSendResult**](AgentSendResult.md) |  | [optional] 

## Methods

### NewDraftResult

`func NewDraftResult(draftId string, status string, policyFlags map[string]interface{}, ) *DraftResult`

NewDraftResult instantiates a new DraftResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDraftResultWithDefaults

`func NewDraftResultWithDefaults() *DraftResult`

NewDraftResultWithDefaults instantiates a new DraftResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraftId

`func (o *DraftResult) GetDraftId() string`

GetDraftId returns the DraftId field if non-nil, zero value otherwise.

### GetDraftIdOk

`func (o *DraftResult) GetDraftIdOk() (*string, bool)`

GetDraftIdOk returns a tuple with the DraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftId

`func (o *DraftResult) SetDraftId(v string)`

SetDraftId sets DraftId field to given value.


### GetStatus

`func (o *DraftResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DraftResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DraftResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPolicyFlags

`func (o *DraftResult) GetPolicyFlags() map[string]interface{}`

GetPolicyFlags returns the PolicyFlags field if non-nil, zero value otherwise.

### GetPolicyFlagsOk

`func (o *DraftResult) GetPolicyFlagsOk() (*map[string]interface{}, bool)`

GetPolicyFlagsOk returns a tuple with the PolicyFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyFlags

`func (o *DraftResult) SetPolicyFlags(v map[string]interface{})`

SetPolicyFlags sets PolicyFlags field to given value.


### GetThreadId

`func (o *DraftResult) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *DraftResult) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *DraftResult) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *DraftResult) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetSent

`func (o *DraftResult) GetSent() AgentSendResult`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *DraftResult) GetSentOk() (*AgentSendResult, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *DraftResult) SetSent(v AgentSendResult)`

SetSent sets Sent field to given value.

### HasSent

`func (o *DraftResult) HasSent() bool`

HasSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


