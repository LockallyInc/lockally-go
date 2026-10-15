# CreateMailboxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | Pointer to **string** | Optional; lockally generates if absent. | [optional] 
**QuotaBytes** | Pointer to **int64** | 5 GB default. | [optional] [default to 5368709120]

## Methods

### NewCreateMailboxRequest

`func NewCreateMailboxRequest(email string, ) *CreateMailboxRequest`

NewCreateMailboxRequest instantiates a new CreateMailboxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMailboxRequestWithDefaults

`func NewCreateMailboxRequestWithDefaults() *CreateMailboxRequest`

NewCreateMailboxRequestWithDefaults instantiates a new CreateMailboxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateMailboxRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateMailboxRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateMailboxRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *CreateMailboxRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateMailboxRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateMailboxRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateMailboxRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetQuotaBytes

`func (o *CreateMailboxRequest) GetQuotaBytes() int64`

GetQuotaBytes returns the QuotaBytes field if non-nil, zero value otherwise.

### GetQuotaBytesOk

`func (o *CreateMailboxRequest) GetQuotaBytesOk() (*int64, bool)`

GetQuotaBytesOk returns a tuple with the QuotaBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaBytes

`func (o *CreateMailboxRequest) SetQuotaBytes(v int64)`

SetQuotaBytes sets QuotaBytes field to given value.

### HasQuotaBytes

`func (o *CreateMailboxRequest) HasQuotaBytes() bool`

HasQuotaBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


