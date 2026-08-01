# CreateEncryptionRecoveryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailboxEmail** | **string** |  | 
**RecoveryBlob** | **string** |  | 

## Methods

### NewCreateEncryptionRecoveryRequest

`func NewCreateEncryptionRecoveryRequest(mailboxEmail string, recoveryBlob string, ) *CreateEncryptionRecoveryRequest`

NewCreateEncryptionRecoveryRequest instantiates a new CreateEncryptionRecoveryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEncryptionRecoveryRequestWithDefaults

`func NewCreateEncryptionRecoveryRequestWithDefaults() *CreateEncryptionRecoveryRequest`

NewCreateEncryptionRecoveryRequestWithDefaults instantiates a new CreateEncryptionRecoveryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxEmail

`func (o *CreateEncryptionRecoveryRequest) GetMailboxEmail() string`

GetMailboxEmail returns the MailboxEmail field if non-nil, zero value otherwise.

### GetMailboxEmailOk

`func (o *CreateEncryptionRecoveryRequest) GetMailboxEmailOk() (*string, bool)`

GetMailboxEmailOk returns a tuple with the MailboxEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxEmail

`func (o *CreateEncryptionRecoveryRequest) SetMailboxEmail(v string)`

SetMailboxEmail sets MailboxEmail field to given value.


### GetRecoveryBlob

`func (o *CreateEncryptionRecoveryRequest) GetRecoveryBlob() string`

GetRecoveryBlob returns the RecoveryBlob field if non-nil, zero value otherwise.

### GetRecoveryBlobOk

`func (o *CreateEncryptionRecoveryRequest) GetRecoveryBlobOk() (*string, bool)`

GetRecoveryBlobOk returns a tuple with the RecoveryBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryBlob

`func (o *CreateEncryptionRecoveryRequest) SetRecoveryBlob(v string)`

SetRecoveryBlob sets RecoveryBlob field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


