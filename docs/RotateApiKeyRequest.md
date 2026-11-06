# RotateApiKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraceSeconds** | Pointer to **int32** | How long the old key keeps working. Defaults to 24 hours, which is long enough to cut over without a scramble. Capped at 30 days, because beyond that \&quot;rotated\&quot; stops meaning anything.  | [optional] [default to 86400]

## Methods

### NewRotateApiKeyRequest

`func NewRotateApiKeyRequest() *RotateApiKeyRequest`

NewRotateApiKeyRequest instantiates a new RotateApiKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateApiKeyRequestWithDefaults

`func NewRotateApiKeyRequestWithDefaults() *RotateApiKeyRequest`

NewRotateApiKeyRequestWithDefaults instantiates a new RotateApiKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraceSeconds

`func (o *RotateApiKeyRequest) GetGraceSeconds() int32`

GetGraceSeconds returns the GraceSeconds field if non-nil, zero value otherwise.

### GetGraceSecondsOk

`func (o *RotateApiKeyRequest) GetGraceSecondsOk() (*int32, bool)`

GetGraceSecondsOk returns a tuple with the GraceSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceSeconds

`func (o *RotateApiKeyRequest) SetGraceSeconds(v int32)`

SetGraceSeconds sets GraceSeconds field to given value.

### HasGraceSeconds

`func (o *RotateApiKeyRequest) HasGraceSeconds() bool`

HasGraceSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


