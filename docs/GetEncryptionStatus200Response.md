# GetEncryptionStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eligible** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Mailboxes** | Pointer to [**[]GetEncryptionStatus200ResponseMailboxesInner**](GetEncryptionStatus200ResponseMailboxesInner.md) |  | [optional] 

## Methods

### NewGetEncryptionStatus200Response

`func NewGetEncryptionStatus200Response() *GetEncryptionStatus200Response`

NewGetEncryptionStatus200Response instantiates a new GetEncryptionStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEncryptionStatus200ResponseWithDefaults

`func NewGetEncryptionStatus200ResponseWithDefaults() *GetEncryptionStatus200Response`

NewGetEncryptionStatus200ResponseWithDefaults instantiates a new GetEncryptionStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligible

`func (o *GetEncryptionStatus200Response) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *GetEncryptionStatus200Response) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *GetEncryptionStatus200Response) SetEligible(v bool)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *GetEncryptionStatus200Response) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetActive

`func (o *GetEncryptionStatus200Response) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetEncryptionStatus200Response) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetEncryptionStatus200Response) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetEncryptionStatus200Response) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMailboxes

`func (o *GetEncryptionStatus200Response) GetMailboxes() []GetEncryptionStatus200ResponseMailboxesInner`

GetMailboxes returns the Mailboxes field if non-nil, zero value otherwise.

### GetMailboxesOk

`func (o *GetEncryptionStatus200Response) GetMailboxesOk() (*[]GetEncryptionStatus200ResponseMailboxesInner, bool)`

GetMailboxesOk returns a tuple with the Mailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxes

`func (o *GetEncryptionStatus200Response) SetMailboxes(v []GetEncryptionStatus200ResponseMailboxesInner)`

SetMailboxes sets Mailboxes field to given value.

### HasMailboxes

`func (o *GetEncryptionStatus200Response) HasMailboxes() bool`

HasMailboxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


