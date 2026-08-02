# CreateEncryptionKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailboxEmail** | **string** |  | 
**PublicKey** | **string** |  | 
**EncryptedPrivateKey** | **string** |  | 
**KdfParams** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateEncryptionKeyRequest

`func NewCreateEncryptionKeyRequest(mailboxEmail string, publicKey string, encryptedPrivateKey string, ) *CreateEncryptionKeyRequest`

NewCreateEncryptionKeyRequest instantiates a new CreateEncryptionKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEncryptionKeyRequestWithDefaults

`func NewCreateEncryptionKeyRequestWithDefaults() *CreateEncryptionKeyRequest`

NewCreateEncryptionKeyRequestWithDefaults instantiates a new CreateEncryptionKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxEmail

`func (o *CreateEncryptionKeyRequest) GetMailboxEmail() string`

GetMailboxEmail returns the MailboxEmail field if non-nil, zero value otherwise.

### GetMailboxEmailOk

`func (o *CreateEncryptionKeyRequest) GetMailboxEmailOk() (*string, bool)`

GetMailboxEmailOk returns a tuple with the MailboxEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxEmail

`func (o *CreateEncryptionKeyRequest) SetMailboxEmail(v string)`

SetMailboxEmail sets MailboxEmail field to given value.


### GetPublicKey

`func (o *CreateEncryptionKeyRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CreateEncryptionKeyRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CreateEncryptionKeyRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetEncryptedPrivateKey

`func (o *CreateEncryptionKeyRequest) GetEncryptedPrivateKey() string`

GetEncryptedPrivateKey returns the EncryptedPrivateKey field if non-nil, zero value otherwise.

### GetEncryptedPrivateKeyOk

`func (o *CreateEncryptionKeyRequest) GetEncryptedPrivateKeyOk() (*string, bool)`

GetEncryptedPrivateKeyOk returns a tuple with the EncryptedPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPrivateKey

`func (o *CreateEncryptionKeyRequest) SetEncryptedPrivateKey(v string)`

SetEncryptedPrivateKey sets EncryptedPrivateKey field to given value.


### GetKdfParams

`func (o *CreateEncryptionKeyRequest) GetKdfParams() map[string]interface{}`

GetKdfParams returns the KdfParams field if non-nil, zero value otherwise.

### GetKdfParamsOk

`func (o *CreateEncryptionKeyRequest) GetKdfParamsOk() (*map[string]interface{}, bool)`

GetKdfParamsOk returns a tuple with the KdfParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdfParams

`func (o *CreateEncryptionKeyRequest) SetKdfParams(v map[string]interface{})`

SetKdfParams sets KdfParams field to given value.

### HasKdfParams

`func (o *CreateEncryptionKeyRequest) HasKdfParams() bool`

HasKdfParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


