# V1MailboxesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Mailbox**](Mailbox.md) |  | 
**Limit** | **int32** |  | 

## Methods

### NewV1MailboxesGet200Response

`func NewV1MailboxesGet200Response(data []Mailbox, limit int32, ) *V1MailboxesGet200Response`

NewV1MailboxesGet200Response instantiates a new V1MailboxesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MailboxesGet200ResponseWithDefaults

`func NewV1MailboxesGet200ResponseWithDefaults() *V1MailboxesGet200Response`

NewV1MailboxesGet200ResponseWithDefaults instantiates a new V1MailboxesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V1MailboxesGet200Response) GetData() []Mailbox`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V1MailboxesGet200Response) GetDataOk() (*[]Mailbox, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V1MailboxesGet200Response) SetData(v []Mailbox)`

SetData sets Data field to given value.


### GetLimit

`func (o *V1MailboxesGet200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V1MailboxesGet200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V1MailboxesGet200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


