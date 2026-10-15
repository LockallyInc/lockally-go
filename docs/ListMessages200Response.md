# ListMessages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Message**](Message.md) |  | 
**NextCursor** | Pointer to **string** | Pass as &#x60;cursor&#x60; to fetch the next page. Absent on the final page. | [optional] 

## Methods

### NewListMessages200Response

`func NewListMessages200Response(data []Message, ) *ListMessages200Response`

NewListMessages200Response instantiates a new ListMessages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMessages200ResponseWithDefaults

`func NewListMessages200ResponseWithDefaults() *ListMessages200Response`

NewListMessages200ResponseWithDefaults instantiates a new ListMessages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListMessages200Response) GetData() []Message`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListMessages200Response) GetDataOk() (*[]Message, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListMessages200Response) SetData(v []Message)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *ListMessages200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListMessages200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListMessages200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListMessages200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


