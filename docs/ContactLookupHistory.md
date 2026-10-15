# ContactLookupHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadCount** | Pointer to **int32** |  | [optional] 
**FirstSeen** | Pointer to **time.Time** |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewContactLookupHistory

`func NewContactLookupHistory() *ContactLookupHistory`

NewContactLookupHistory instantiates a new ContactLookupHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactLookupHistoryWithDefaults

`func NewContactLookupHistoryWithDefaults() *ContactLookupHistory`

NewContactLookupHistoryWithDefaults instantiates a new ContactLookupHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadCount

`func (o *ContactLookupHistory) GetThreadCount() int32`

GetThreadCount returns the ThreadCount field if non-nil, zero value otherwise.

### GetThreadCountOk

`func (o *ContactLookupHistory) GetThreadCountOk() (*int32, bool)`

GetThreadCountOk returns a tuple with the ThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadCount

`func (o *ContactLookupHistory) SetThreadCount(v int32)`

SetThreadCount sets ThreadCount field to given value.

### HasThreadCount

`func (o *ContactLookupHistory) HasThreadCount() bool`

HasThreadCount returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ContactLookupHistory) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ContactLookupHistory) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ContactLookupHistory) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ContactLookupHistory) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *ContactLookupHistory) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ContactLookupHistory) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ContactLookupHistory) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ContactLookupHistory) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


