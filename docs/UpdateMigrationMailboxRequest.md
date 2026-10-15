# UpdateMigrationMailboxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestEmail** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Can only be set to \&quot;skipped\&quot;. | [optional] 

## Methods

### NewUpdateMigrationMailboxRequest

`func NewUpdateMigrationMailboxRequest() *UpdateMigrationMailboxRequest`

NewUpdateMigrationMailboxRequest instantiates a new UpdateMigrationMailboxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMigrationMailboxRequestWithDefaults

`func NewUpdateMigrationMailboxRequestWithDefaults() *UpdateMigrationMailboxRequest`

NewUpdateMigrationMailboxRequestWithDefaults instantiates a new UpdateMigrationMailboxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestEmail

`func (o *UpdateMigrationMailboxRequest) GetDestEmail() string`

GetDestEmail returns the DestEmail field if non-nil, zero value otherwise.

### GetDestEmailOk

`func (o *UpdateMigrationMailboxRequest) GetDestEmailOk() (*string, bool)`

GetDestEmailOk returns a tuple with the DestEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestEmail

`func (o *UpdateMigrationMailboxRequest) SetDestEmail(v string)`

SetDestEmail sets DestEmail field to given value.

### HasDestEmail

`func (o *UpdateMigrationMailboxRequest) HasDestEmail() bool`

HasDestEmail returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateMigrationMailboxRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateMigrationMailboxRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateMigrationMailboxRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateMigrationMailboxRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


