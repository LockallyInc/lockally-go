# MessageStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Window** | Pointer to [**MessageStatsWindow**](MessageStatsWindow.md) |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Sent** | Pointer to **int32** |  | [optional] 
**Counts** | Pointer to [**MessageStatsCounts**](MessageStatsCounts.md) |  | [optional] 
**Rates** | Pointer to [**MessageStatsRates**](MessageStatsRates.md) |  | [optional] 

## Methods

### NewMessageStats

`func NewMessageStats() *MessageStats`

NewMessageStats instantiates a new MessageStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageStatsWithDefaults

`func NewMessageStatsWithDefaults() *MessageStats`

NewMessageStatsWithDefaults instantiates a new MessageStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindow

`func (o *MessageStats) GetWindow() MessageStatsWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *MessageStats) GetWindowOk() (*MessageStatsWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *MessageStats) SetWindow(v MessageStatsWindow)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *MessageStats) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetDomain

`func (o *MessageStats) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *MessageStats) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *MessageStats) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *MessageStats) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetSent

`func (o *MessageStats) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *MessageStats) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *MessageStats) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *MessageStats) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetCounts

`func (o *MessageStats) GetCounts() MessageStatsCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *MessageStats) GetCountsOk() (*MessageStatsCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *MessageStats) SetCounts(v MessageStatsCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *MessageStats) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetRates

`func (o *MessageStats) GetRates() MessageStatsRates`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *MessageStats) GetRatesOk() (*MessageStatsRates, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *MessageStats) SetRates(v MessageStatsRates)`

SetRates sets Rates field to given value.

### HasRates

`func (o *MessageStats) HasRates() bool`

HasRates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


