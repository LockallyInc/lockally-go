# V1MessagesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Message**](Message.md) |  | 
**NextCursor** | Pointer to **string** | Pass as &#x60;cursor&#x60; to fetch the next page. Absent on the final page. | [optional] 

## Methods

### NewV1MessagesGet200Response

`func NewV1MessagesGet200Response(data []Message, ) *V1MessagesGet200Response`

NewV1MessagesGet200Response instantiates a new V1MessagesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MessagesGet200ResponseWithDefaults

`func NewV1MessagesGet200ResponseWithDefaults() *V1MessagesGet200Response`

NewV1MessagesGet200ResponseWithDefaults instantiates a new V1MessagesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V1MessagesGet200Response) GetData() []Message`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V1MessagesGet200Response) GetDataOk() (*[]Message, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V1MessagesGet200Response) SetData(v []Message)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *V1MessagesGet200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *V1MessagesGet200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *V1MessagesGet200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *V1MessagesGet200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


