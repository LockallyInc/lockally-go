# V1MailboxesEmailVacationPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Params** | [**VacationParams**](VacationParams.md) |  | 

## Methods

### NewV1MailboxesEmailVacationPutRequest

`func NewV1MailboxesEmailVacationPutRequest(params VacationParams, ) *V1MailboxesEmailVacationPutRequest`

NewV1MailboxesEmailVacationPutRequest instantiates a new V1MailboxesEmailVacationPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MailboxesEmailVacationPutRequestWithDefaults

`func NewV1MailboxesEmailVacationPutRequestWithDefaults() *V1MailboxesEmailVacationPutRequest`

NewV1MailboxesEmailVacationPutRequestWithDefaults instantiates a new V1MailboxesEmailVacationPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *V1MailboxesEmailVacationPutRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V1MailboxesEmailVacationPutRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V1MailboxesEmailVacationPutRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V1MailboxesEmailVacationPutRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetParams

`func (o *V1MailboxesEmailVacationPutRequest) GetParams() VacationParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *V1MailboxesEmailVacationPutRequest) GetParamsOk() (*VacationParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *V1MailboxesEmailVacationPutRequest) SetParams(v VacationParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


