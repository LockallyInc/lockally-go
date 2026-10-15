# ListSuppressions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Suppression**](Suppression.md) |  | [optional] 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListSuppressions200Response

`func NewListSuppressions200Response() *ListSuppressions200Response`

NewListSuppressions200Response instantiates a new ListSuppressions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSuppressions200ResponseWithDefaults

`func NewListSuppressions200ResponseWithDefaults() *ListSuppressions200Response`

NewListSuppressions200ResponseWithDefaults instantiates a new ListSuppressions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListSuppressions200Response) GetData() []Suppression`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListSuppressions200Response) GetDataOk() (*[]Suppression, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListSuppressions200Response) SetData(v []Suppression)`

SetData sets Data field to given value.

### HasData

`func (o *ListSuppressions200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextCursor

`func (o *ListSuppressions200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListSuppressions200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListSuppressions200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListSuppressions200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


