# CreateSuppressionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Reason** | Pointer to **string** |  | [optional] [default to "manual"]

## Methods

### NewCreateSuppressionRequest

`func NewCreateSuppressionRequest(email string, ) *CreateSuppressionRequest`

NewCreateSuppressionRequest instantiates a new CreateSuppressionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSuppressionRequestWithDefaults

`func NewCreateSuppressionRequestWithDefaults() *CreateSuppressionRequest`

NewCreateSuppressionRequestWithDefaults instantiates a new CreateSuppressionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateSuppressionRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateSuppressionRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateSuppressionRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetReason

`func (o *CreateSuppressionRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateSuppressionRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateSuppressionRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateSuppressionRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


