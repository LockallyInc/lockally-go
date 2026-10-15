# CreateCalendarRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Color** | Pointer to **string** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] [default to "tenant"]

## Methods

### NewCreateCalendarRequest

`func NewCreateCalendarRequest(name string, ) *CreateCalendarRequest`

NewCreateCalendarRequest instantiates a new CreateCalendarRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCalendarRequestWithDefaults

`func NewCreateCalendarRequestWithDefaults() *CreateCalendarRequest`

NewCreateCalendarRequestWithDefaults instantiates a new CreateCalendarRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCalendarRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCalendarRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCalendarRequest) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *CreateCalendarRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateCalendarRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateCalendarRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateCalendarRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *CreateCalendarRequest) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *CreateCalendarRequest) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *CreateCalendarRequest) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *CreateCalendarRequest) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetDescription

`func (o *CreateCalendarRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCalendarRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCalendarRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCalendarRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateCalendarRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateCalendarRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateCalendarRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateCalendarRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


