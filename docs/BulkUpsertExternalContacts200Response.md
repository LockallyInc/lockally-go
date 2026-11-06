# BulkUpsertExternalContacts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **int32** |  | 
**Updated** | **int32** |  | 
**Failed** | **int32** |  | 
**Results** | [**[]BulkUpsertExternalContacts200ResponseResultsInner**](BulkUpsertExternalContacts200ResponseResultsInner.md) |  | 

## Methods

### NewBulkUpsertExternalContacts200Response

`func NewBulkUpsertExternalContacts200Response(created int32, updated int32, failed int32, results []BulkUpsertExternalContacts200ResponseResultsInner, ) *BulkUpsertExternalContacts200Response`

NewBulkUpsertExternalContacts200Response instantiates a new BulkUpsertExternalContacts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpsertExternalContacts200ResponseWithDefaults

`func NewBulkUpsertExternalContacts200ResponseWithDefaults() *BulkUpsertExternalContacts200Response`

NewBulkUpsertExternalContacts200ResponseWithDefaults instantiates a new BulkUpsertExternalContacts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *BulkUpsertExternalContacts200Response) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BulkUpsertExternalContacts200Response) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BulkUpsertExternalContacts200Response) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *BulkUpsertExternalContacts200Response) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *BulkUpsertExternalContacts200Response) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *BulkUpsertExternalContacts200Response) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.


### GetFailed

`func (o *BulkUpsertExternalContacts200Response) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *BulkUpsertExternalContacts200Response) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *BulkUpsertExternalContacts200Response) SetFailed(v int32)`

SetFailed sets Failed field to given value.


### GetResults

`func (o *BulkUpsertExternalContacts200Response) GetResults() []BulkUpsertExternalContacts200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkUpsertExternalContacts200Response) GetResultsOk() (*[]BulkUpsertExternalContacts200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkUpsertExternalContacts200Response) SetResults(v []BulkUpsertExternalContacts200ResponseResultsInner)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


