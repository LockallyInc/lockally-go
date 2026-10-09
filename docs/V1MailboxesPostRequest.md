# V1MailboxesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | Pointer to **string** | Optional; lockally generates if absent. | [optional] 
**QuotaBytes** | Pointer to **int64** | 5 GB default. | [optional] [default to 5368709120]

## Methods

### NewV1MailboxesPostRequest

`func NewV1MailboxesPostRequest(email string, ) *V1MailboxesPostRequest`

NewV1MailboxesPostRequest instantiates a new V1MailboxesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MailboxesPostRequestWithDefaults

`func NewV1MailboxesPostRequestWithDefaults() *V1MailboxesPostRequest`

NewV1MailboxesPostRequestWithDefaults instantiates a new V1MailboxesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *V1MailboxesPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *V1MailboxesPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *V1MailboxesPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *V1MailboxesPostRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V1MailboxesPostRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V1MailboxesPostRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V1MailboxesPostRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetQuotaBytes

`func (o *V1MailboxesPostRequest) GetQuotaBytes() int64`

GetQuotaBytes returns the QuotaBytes field if non-nil, zero value otherwise.

### GetQuotaBytesOk

`func (o *V1MailboxesPostRequest) GetQuotaBytesOk() (*int64, bool)`

GetQuotaBytesOk returns a tuple with the QuotaBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaBytes

`func (o *V1MailboxesPostRequest) SetQuotaBytes(v int64)`

SetQuotaBytes sets QuotaBytes field to given value.

### HasQuotaBytes

`func (o *V1MailboxesPostRequest) HasQuotaBytes() bool`

HasQuotaBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


