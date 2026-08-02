# ListContacts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Contact**](Contact.md) |  | [optional] 

## Methods

### NewListContacts200Response

`func NewListContacts200Response() *ListContacts200Response`

NewListContacts200Response instantiates a new ListContacts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListContacts200ResponseWithDefaults

`func NewListContacts200ResponseWithDefaults() *ListContacts200Response`

NewListContacts200ResponseWithDefaults instantiates a new ListContacts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListContacts200Response) GetData() []Contact`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListContacts200Response) GetDataOk() (*[]Contact, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListContacts200Response) SetData(v []Contact)`

SetData sets Data field to given value.

### HasData

`func (o *ListContacts200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


