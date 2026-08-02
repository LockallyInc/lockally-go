# MessageStatsWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **time.Time** |  | [optional] 
**To** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMessageStatsWindow

`func NewMessageStatsWindow() *MessageStatsWindow`

NewMessageStatsWindow instantiates a new MessageStatsWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageStatsWindowWithDefaults

`func NewMessageStatsWindowWithDefaults() *MessageStatsWindow`

NewMessageStatsWindowWithDefaults instantiates a new MessageStatsWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *MessageStatsWindow) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MessageStatsWindow) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MessageStatsWindow) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MessageStatsWindow) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *MessageStatsWindow) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MessageStatsWindow) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MessageStatsWindow) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *MessageStatsWindow) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


