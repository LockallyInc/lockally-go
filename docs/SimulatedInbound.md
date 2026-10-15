# SimulatedInbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSimulatedInbound

`func NewSimulatedInbound() *SimulatedInbound`

NewSimulatedInbound instantiates a new SimulatedInbound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulatedInboundWithDefaults

`func NewSimulatedInboundWithDefaults() *SimulatedInbound`

NewSimulatedInboundWithDefaults instantiates a new SimulatedInbound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *SimulatedInbound) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *SimulatedInbound) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *SimulatedInbound) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *SimulatedInbound) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetMessageId

`func (o *SimulatedInbound) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SimulatedInbound) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SimulatedInbound) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *SimulatedInbound) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetAnnotations

`func (o *SimulatedInbound) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *SimulatedInbound) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *SimulatedInbound) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *SimulatedInbound) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


