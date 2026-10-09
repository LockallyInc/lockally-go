# RotateEncryptionKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailboxEmail** | **string** |  | 
**EncryptedPrivateKey** | **string** |  | 
**KdfParams** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRotateEncryptionKeyRequest

`func NewRotateEncryptionKeyRequest(mailboxEmail string, encryptedPrivateKey string, ) *RotateEncryptionKeyRequest`

NewRotateEncryptionKeyRequest instantiates a new RotateEncryptionKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateEncryptionKeyRequestWithDefaults

`func NewRotateEncryptionKeyRequestWithDefaults() *RotateEncryptionKeyRequest`

NewRotateEncryptionKeyRequestWithDefaults instantiates a new RotateEncryptionKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxEmail

`func (o *RotateEncryptionKeyRequest) GetMailboxEmail() string`

GetMailboxEmail returns the MailboxEmail field if non-nil, zero value otherwise.

### GetMailboxEmailOk

`func (o *RotateEncryptionKeyRequest) GetMailboxEmailOk() (*string, bool)`

GetMailboxEmailOk returns a tuple with the MailboxEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxEmail

`func (o *RotateEncryptionKeyRequest) SetMailboxEmail(v string)`

SetMailboxEmail sets MailboxEmail field to given value.


### GetEncryptedPrivateKey

`func (o *RotateEncryptionKeyRequest) GetEncryptedPrivateKey() string`

GetEncryptedPrivateKey returns the EncryptedPrivateKey field if non-nil, zero value otherwise.

### GetEncryptedPrivateKeyOk

`func (o *RotateEncryptionKeyRequest) GetEncryptedPrivateKeyOk() (*string, bool)`

GetEncryptedPrivateKeyOk returns a tuple with the EncryptedPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPrivateKey

`func (o *RotateEncryptionKeyRequest) SetEncryptedPrivateKey(v string)`

SetEncryptedPrivateKey sets EncryptedPrivateKey field to given value.


### GetKdfParams

`func (o *RotateEncryptionKeyRequest) GetKdfParams() map[string]interface{}`

GetKdfParams returns the KdfParams field if non-nil, zero value otherwise.

### GetKdfParamsOk

`func (o *RotateEncryptionKeyRequest) GetKdfParamsOk() (*map[string]interface{}, bool)`

GetKdfParamsOk returns a tuple with the KdfParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdfParams

`func (o *RotateEncryptionKeyRequest) SetKdfParams(v map[string]interface{})`

SetKdfParams sets KdfParams field to given value.

### HasKdfParams

`func (o *RotateEncryptionKeyRequest) HasKdfParams() bool`

HasKdfParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


