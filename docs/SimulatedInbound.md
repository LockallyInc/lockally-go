# SimulatedInbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Simulated** | Pointer to **bool** |  | [optional] 
**Indexed** | Pointer to **bool** | True when the message was adopted into a thread. | [optional] 
**ThreadId** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**RfcMessageId** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to [**Annotations**](Annotations.md) |  | [optional] 

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

### GetSimulated

`func (o *SimulatedInbound) GetSimulated() bool`

GetSimulated returns the Simulated field if non-nil, zero value otherwise.

### GetSimulatedOk

`func (o *SimulatedInbound) GetSimulatedOk() (*bool, bool)`

GetSimulatedOk returns a tuple with the Simulated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulated

`func (o *SimulatedInbound) SetSimulated(v bool)`

SetSimulated sets Simulated field to given value.

### HasSimulated

`func (o *SimulatedInbound) HasSimulated() bool`

HasSimulated returns a boolean if a field has been set.

### GetIndexed

`func (o *SimulatedInbound) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *SimulatedInbound) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *SimulatedInbound) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *SimulatedInbound) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

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

### GetRfcMessageId

`func (o *SimulatedInbound) GetRfcMessageId() string`

GetRfcMessageId returns the RfcMessageId field if non-nil, zero value otherwise.

### GetRfcMessageIdOk

`func (o *SimulatedInbound) GetRfcMessageIdOk() (*string, bool)`

GetRfcMessageIdOk returns a tuple with the RfcMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfcMessageId

`func (o *SimulatedInbound) SetRfcMessageId(v string)`

SetRfcMessageId sets RfcMessageId field to given value.

### HasRfcMessageId

`func (o *SimulatedInbound) HasRfcMessageId() bool`

HasRfcMessageId returns a boolean if a field has been set.

### GetAnnotations

`func (o *SimulatedInbound) GetAnnotations() Annotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *SimulatedInbound) GetAnnotationsOk() (*Annotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *SimulatedInbound) SetAnnotations(v Annotations)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *SimulatedInbound) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


