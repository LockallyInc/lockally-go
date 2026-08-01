# V1MailboxesEmailPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] 
**QuotaBytes** | Pointer to **int64** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1MailboxesEmailPatchRequest

`func NewV1MailboxesEmailPatchRequest() *V1MailboxesEmailPatchRequest`

NewV1MailboxesEmailPatchRequest instantiates a new V1MailboxesEmailPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MailboxesEmailPatchRequestWithDefaults

`func NewV1MailboxesEmailPatchRequestWithDefaults() *V1MailboxesEmailPatchRequest`

NewV1MailboxesEmailPatchRequestWithDefaults instantiates a new V1MailboxesEmailPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *V1MailboxesEmailPatchRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V1MailboxesEmailPatchRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V1MailboxesEmailPatchRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V1MailboxesEmailPatchRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetQuotaBytes

`func (o *V1MailboxesEmailPatchRequest) GetQuotaBytes() int64`

GetQuotaBytes returns the QuotaBytes field if non-nil, zero value otherwise.

### GetQuotaBytesOk

`func (o *V1MailboxesEmailPatchRequest) GetQuotaBytesOk() (*int64, bool)`

GetQuotaBytesOk returns a tuple with the QuotaBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaBytes

`func (o *V1MailboxesEmailPatchRequest) SetQuotaBytes(v int64)`

SetQuotaBytes sets QuotaBytes field to given value.

### HasQuotaBytes

`func (o *V1MailboxesEmailPatchRequest) HasQuotaBytes() bool`

HasQuotaBytes returns a boolean if a field has been set.

### GetDisabled

`func (o *V1MailboxesEmailPatchRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V1MailboxesEmailPatchRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V1MailboxesEmailPatchRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *V1MailboxesEmailPatchRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


