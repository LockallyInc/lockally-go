# ListCalendars200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Calendar**](Calendar.md) |  | [optional] 

## Methods

### NewListCalendars200Response

`func NewListCalendars200Response() *ListCalendars200Response`

NewListCalendars200Response instantiates a new ListCalendars200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCalendars200ResponseWithDefaults

`func NewListCalendars200ResponseWithDefaults() *ListCalendars200Response`

NewListCalendars200ResponseWithDefaults instantiates a new ListCalendars200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCalendars200Response) GetData() []Calendar`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCalendars200Response) GetDataOk() (*[]Calendar, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCalendars200Response) SetData(v []Calendar)`

SetData sets Data field to given value.

### HasData

`func (o *ListCalendars200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


