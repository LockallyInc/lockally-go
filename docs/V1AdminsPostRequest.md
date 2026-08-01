# V1AdminsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] [default to "admin"]

## Methods

### NewV1AdminsPostRequest

`func NewV1AdminsPostRequest(email string, ) *V1AdminsPostRequest`

NewV1AdminsPostRequest instantiates a new V1AdminsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AdminsPostRequestWithDefaults

`func NewV1AdminsPostRequestWithDefaults() *V1AdminsPostRequest`

NewV1AdminsPostRequestWithDefaults instantiates a new V1AdminsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *V1AdminsPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *V1AdminsPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *V1AdminsPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *V1AdminsPostRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V1AdminsPostRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V1AdminsPostRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V1AdminsPostRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDisplayName

`func (o *V1AdminsPostRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *V1AdminsPostRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *V1AdminsPostRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *V1AdminsPostRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRole

`func (o *V1AdminsPostRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *V1AdminsPostRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *V1AdminsPostRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *V1AdminsPostRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


