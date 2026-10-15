# Suppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSuppression

`func NewSuppression() *Suppression`

NewSuppression instantiates a new Suppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressionWithDefaults

`func NewSuppressionWithDefaults() *Suppression`

NewSuppressionWithDefaults instantiates a new Suppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Suppression) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Suppression) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Suppression) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Suppression) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetReason

`func (o *Suppression) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Suppression) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Suppression) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Suppression) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSource

`func (o *Suppression) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Suppression) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Suppression) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Suppression) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Suppression) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Suppression) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Suppression) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Suppression) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


