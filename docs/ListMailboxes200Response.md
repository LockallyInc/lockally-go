# ListMailboxes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Mailbox**](Mailbox.md) |  | 
**Limit** | **int32** |  | 

## Methods

### NewListMailboxes200Response

`func NewListMailboxes200Response(data []Mailbox, limit int32, ) *ListMailboxes200Response`

NewListMailboxes200Response instantiates a new ListMailboxes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMailboxes200ResponseWithDefaults

`func NewListMailboxes200ResponseWithDefaults() *ListMailboxes200Response`

NewListMailboxes200ResponseWithDefaults instantiates a new ListMailboxes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListMailboxes200Response) GetData() []Mailbox`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListMailboxes200Response) GetDataOk() (*[]Mailbox, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListMailboxes200Response) SetData(v []Mailbox)`

SetData sets Data field to given value.


### GetLimit

`func (o *ListMailboxes200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListMailboxes200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListMailboxes200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


