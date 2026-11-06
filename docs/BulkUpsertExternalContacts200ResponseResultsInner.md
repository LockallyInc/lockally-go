# BulkUpsertExternalContacts200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**ExternalId** | **string** |  | 
**Status** | **string** |  | 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkUpsertExternalContacts200ResponseResultsInner

`func NewBulkUpsertExternalContacts200ResponseResultsInner(index int32, externalId string, status string, ) *BulkUpsertExternalContacts200ResponseResultsInner`

NewBulkUpsertExternalContacts200ResponseResultsInner instantiates a new BulkUpsertExternalContacts200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpsertExternalContacts200ResponseResultsInnerWithDefaults

`func NewBulkUpsertExternalContacts200ResponseResultsInnerWithDefaults() *BulkUpsertExternalContacts200ResponseResultsInner`

NewBulkUpsertExternalContacts200ResponseResultsInnerWithDefaults instantiates a new BulkUpsertExternalContacts200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetExternalId

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetStatus

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetContact

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetError

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BulkUpsertExternalContacts200ResponseResultsInner) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


