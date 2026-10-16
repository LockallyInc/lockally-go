# ContactLookup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Known** | **bool** |  | 
**Internal** | **bool** | True when the address is one of the tenant&#39;s own mailboxes. | 
**Contact** | Pointer to [**ContactLookupContact**](ContactLookupContact.md) |  | [optional] 
**History** | Pointer to [**ContactLookupHistory**](ContactLookupHistory.md) |  | [optional] 

## Methods

### NewContactLookup

`func NewContactLookup(email string, known bool, internal bool, ) *ContactLookup`

NewContactLookup instantiates a new ContactLookup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactLookupWithDefaults

`func NewContactLookupWithDefaults() *ContactLookup`

NewContactLookupWithDefaults instantiates a new ContactLookup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ContactLookup) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactLookup) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactLookup) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetKnown

`func (o *ContactLookup) GetKnown() bool`

GetKnown returns the Known field if non-nil, zero value otherwise.

### GetKnownOk

`func (o *ContactLookup) GetKnownOk() (*bool, bool)`

GetKnownOk returns a tuple with the Known field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnown

`func (o *ContactLookup) SetKnown(v bool)`

SetKnown sets Known field to given value.


### GetInternal

`func (o *ContactLookup) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ContactLookup) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ContactLookup) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetContact

`func (o *ContactLookup) GetContact() ContactLookupContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ContactLookup) GetContactOk() (*ContactLookupContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ContactLookup) SetContact(v ContactLookupContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ContactLookup) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetHistory

`func (o *ContactLookup) GetHistory() ContactLookupHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *ContactLookup) GetHistoryOk() (*ContactLookupHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *ContactLookup) SetHistory(v ContactLookupHistory)`

SetHistory sets History field to given value.

### HasHistory

`func (o *ContactLookup) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


