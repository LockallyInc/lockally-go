# GetEncryptionKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**EncryptedPrivateKey** | Pointer to **string** |  | [optional] 
**KdfParams** | Pointer to **map[string]interface{}** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**RotatedAt** | Pointer to **time.Time** |  | [optional] 
**HasRecovery** | Pointer to **bool** |  | [optional] 
**RecoverySetAt** | Pointer to **time.Time** |  | [optional] 
**RecoveryBlob** | Pointer to **string** |  | [optional] 

## Methods

### NewGetEncryptionKey200Response

`func NewGetEncryptionKey200Response() *GetEncryptionKey200Response`

NewGetEncryptionKey200Response instantiates a new GetEncryptionKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEncryptionKey200ResponseWithDefaults

`func NewGetEncryptionKey200ResponseWithDefaults() *GetEncryptionKey200Response`

NewGetEncryptionKey200ResponseWithDefaults instantiates a new GetEncryptionKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEncryptionKey200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEncryptionKey200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEncryptionKey200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetEncryptionKey200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *GetEncryptionKey200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetEncryptionKey200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetEncryptionKey200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetEncryptionKey200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPublicKey

`func (o *GetEncryptionKey200Response) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GetEncryptionKey200Response) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GetEncryptionKey200Response) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *GetEncryptionKey200Response) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetEncryptedPrivateKey

`func (o *GetEncryptionKey200Response) GetEncryptedPrivateKey() string`

GetEncryptedPrivateKey returns the EncryptedPrivateKey field if non-nil, zero value otherwise.

### GetEncryptedPrivateKeyOk

`func (o *GetEncryptionKey200Response) GetEncryptedPrivateKeyOk() (*string, bool)`

GetEncryptedPrivateKeyOk returns a tuple with the EncryptedPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPrivateKey

`func (o *GetEncryptionKey200Response) SetEncryptedPrivateKey(v string)`

SetEncryptedPrivateKey sets EncryptedPrivateKey field to given value.

### HasEncryptedPrivateKey

`func (o *GetEncryptionKey200Response) HasEncryptedPrivateKey() bool`

HasEncryptedPrivateKey returns a boolean if a field has been set.

### GetKdfParams

`func (o *GetEncryptionKey200Response) GetKdfParams() map[string]interface{}`

GetKdfParams returns the KdfParams field if non-nil, zero value otherwise.

### GetKdfParamsOk

`func (o *GetEncryptionKey200Response) GetKdfParamsOk() (*map[string]interface{}, bool)`

GetKdfParamsOk returns a tuple with the KdfParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdfParams

`func (o *GetEncryptionKey200Response) SetKdfParams(v map[string]interface{})`

SetKdfParams sets KdfParams field to given value.

### HasKdfParams

`func (o *GetEncryptionKey200Response) HasKdfParams() bool`

HasKdfParams returns a boolean if a field has been set.

### GetVersion

`func (o *GetEncryptionKey200Response) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetEncryptionKey200Response) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetEncryptionKey200Response) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetEncryptionKey200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetEncryptionKey200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetEncryptionKey200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetEncryptionKey200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetEncryptionKey200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRotatedAt

`func (o *GetEncryptionKey200Response) GetRotatedAt() time.Time`

GetRotatedAt returns the RotatedAt field if non-nil, zero value otherwise.

### GetRotatedAtOk

`func (o *GetEncryptionKey200Response) GetRotatedAtOk() (*time.Time, bool)`

GetRotatedAtOk returns a tuple with the RotatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedAt

`func (o *GetEncryptionKey200Response) SetRotatedAt(v time.Time)`

SetRotatedAt sets RotatedAt field to given value.

### HasRotatedAt

`func (o *GetEncryptionKey200Response) HasRotatedAt() bool`

HasRotatedAt returns a boolean if a field has been set.

### GetHasRecovery

`func (o *GetEncryptionKey200Response) GetHasRecovery() bool`

GetHasRecovery returns the HasRecovery field if non-nil, zero value otherwise.

### GetHasRecoveryOk

`func (o *GetEncryptionKey200Response) GetHasRecoveryOk() (*bool, bool)`

GetHasRecoveryOk returns a tuple with the HasRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRecovery

`func (o *GetEncryptionKey200Response) SetHasRecovery(v bool)`

SetHasRecovery sets HasRecovery field to given value.

### HasHasRecovery

`func (o *GetEncryptionKey200Response) HasHasRecovery() bool`

HasHasRecovery returns a boolean if a field has been set.

### GetRecoverySetAt

`func (o *GetEncryptionKey200Response) GetRecoverySetAt() time.Time`

GetRecoverySetAt returns the RecoverySetAt field if non-nil, zero value otherwise.

### GetRecoverySetAtOk

`func (o *GetEncryptionKey200Response) GetRecoverySetAtOk() (*time.Time, bool)`

GetRecoverySetAtOk returns a tuple with the RecoverySetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverySetAt

`func (o *GetEncryptionKey200Response) SetRecoverySetAt(v time.Time)`

SetRecoverySetAt sets RecoverySetAt field to given value.

### HasRecoverySetAt

`func (o *GetEncryptionKey200Response) HasRecoverySetAt() bool`

HasRecoverySetAt returns a boolean if a field has been set.

### GetRecoveryBlob

`func (o *GetEncryptionKey200Response) GetRecoveryBlob() string`

GetRecoveryBlob returns the RecoveryBlob field if non-nil, zero value otherwise.

### GetRecoveryBlobOk

`func (o *GetEncryptionKey200Response) GetRecoveryBlobOk() (*string, bool)`

GetRecoveryBlobOk returns a tuple with the RecoveryBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryBlob

`func (o *GetEncryptionKey200Response) SetRecoveryBlob(v string)`

SetRecoveryBlob sets RecoveryBlob field to given value.

### HasRecoveryBlob

`func (o *GetEncryptionKey200Response) HasRecoveryBlob() bool`

HasRecoveryBlob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


