# BulkUpsertExternalContactsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contacts** | [**[]BulkUpsertExternalContactsRequestContactsInner**](BulkUpsertExternalContactsRequestContactsInner.md) |  | 

## Methods

### NewBulkUpsertExternalContactsRequest

`func NewBulkUpsertExternalContactsRequest(contacts []BulkUpsertExternalContactsRequestContactsInner, ) *BulkUpsertExternalContactsRequest`

NewBulkUpsertExternalContactsRequest instantiates a new BulkUpsertExternalContactsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpsertExternalContactsRequestWithDefaults

`func NewBulkUpsertExternalContactsRequestWithDefaults() *BulkUpsertExternalContactsRequest`

NewBulkUpsertExternalContactsRequestWithDefaults instantiates a new BulkUpsertExternalContactsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContacts

`func (o *BulkUpsertExternalContactsRequest) GetContacts() []BulkUpsertExternalContactsRequestContactsInner`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *BulkUpsertExternalContactsRequest) GetContactsOk() (*[]BulkUpsertExternalContactsRequestContactsInner, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *BulkUpsertExternalContactsRequest) SetContacts(v []BulkUpsertExternalContactsRequestContactsInner)`

SetContacts sets Contacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


