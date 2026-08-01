# CreateContactListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] [default to "company_wide"]

## Methods

### NewCreateContactListRequest

`func NewCreateContactListRequest(name string, ) *CreateContactListRequest`

NewCreateContactListRequest instantiates a new CreateContactListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContactListRequestWithDefaults

`func NewCreateContactListRequestWithDefaults() *CreateContactListRequest`

NewCreateContactListRequestWithDefaults instantiates a new CreateContactListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateContactListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateContactListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateContactListRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateContactListRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateContactListRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateContactListRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateContactListRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateContactListRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateContactListRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateContactListRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateContactListRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


