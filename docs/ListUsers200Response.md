# ListUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]User**](User.md) |  | 
**Limit** | **int32** |  | 

## Methods

### NewListUsers200Response

`func NewListUsers200Response(data []User, limit int32, ) *ListUsers200Response`

NewListUsers200Response instantiates a new ListUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsers200ResponseWithDefaults

`func NewListUsers200ResponseWithDefaults() *ListUsers200Response`

NewListUsers200ResponseWithDefaults instantiates a new ListUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListUsers200Response) GetData() []User`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListUsers200Response) GetDataOk() (*[]User, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListUsers200Response) SetData(v []User)`

SetData sets Data field to given value.


### GetLimit

`func (o *ListUsers200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListUsers200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListUsers200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


