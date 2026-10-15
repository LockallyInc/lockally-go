# Classification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Intent** | Pointer to **string** | For example support_request, sales_lead, billing, spam. | [optional] 
**Urgency** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** | One line. | [optional] 
**SuggestedAction** | Pointer to **string** | One line. | [optional] 

## Methods

### NewClassification

`func NewClassification() *Classification`

NewClassification instantiates a new Classification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassificationWithDefaults

`func NewClassificationWithDefaults() *Classification`

NewClassificationWithDefaults instantiates a new Classification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntent

`func (o *Classification) GetIntent() string`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *Classification) GetIntentOk() (*string, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *Classification) SetIntent(v string)`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *Classification) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

### GetUrgency

`func (o *Classification) GetUrgency() string`

GetUrgency returns the Urgency field if non-nil, zero value otherwise.

### GetUrgencyOk

`func (o *Classification) GetUrgencyOk() (*string, bool)`

GetUrgencyOk returns a tuple with the Urgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgency

`func (o *Classification) SetUrgency(v string)`

SetUrgency sets Urgency field to given value.

### HasUrgency

`func (o *Classification) HasUrgency() bool`

HasUrgency returns a boolean if a field has been set.

### GetSummary

`func (o *Classification) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Classification) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Classification) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Classification) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSuggestedAction

`func (o *Classification) GetSuggestedAction() string`

GetSuggestedAction returns the SuggestedAction field if non-nil, zero value otherwise.

### GetSuggestedActionOk

`func (o *Classification) GetSuggestedActionOk() (*string, bool)`

GetSuggestedActionOk returns a tuple with the SuggestedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedAction

`func (o *Classification) SetSuggestedAction(v string)`

SetSuggestedAction sets SuggestedAction field to given value.

### HasSuggestedAction

`func (o *Classification) HasSuggestedAction() bool`

HasSuggestedAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


