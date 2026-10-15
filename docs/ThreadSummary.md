# ThreadSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Subject** | Pointer to **string** |  | [optional] 
**Participants** | **[]string** |  | 
**MessageCount** | **int32** |  | 
**LastMessageAt** | Pointer to **time.Time** |  | [optional] 
**LastDirection** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewThreadSummary

`func NewThreadSummary(id string, participants []string, messageCount int32, ) *ThreadSummary`

NewThreadSummary instantiates a new ThreadSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadSummaryWithDefaults

`func NewThreadSummaryWithDefaults() *ThreadSummary`

NewThreadSummaryWithDefaults instantiates a new ThreadSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThreadSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreadSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreadSummary) SetId(v string)`

SetId sets Id field to given value.


### GetSubject

`func (o *ThreadSummary) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ThreadSummary) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ThreadSummary) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ThreadSummary) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetParticipants

`func (o *ThreadSummary) GetParticipants() []string`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ThreadSummary) GetParticipantsOk() (*[]string, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ThreadSummary) SetParticipants(v []string)`

SetParticipants sets Participants field to given value.


### GetMessageCount

`func (o *ThreadSummary) GetMessageCount() int32`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *ThreadSummary) GetMessageCountOk() (*int32, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *ThreadSummary) SetMessageCount(v int32)`

SetMessageCount sets MessageCount field to given value.


### GetLastMessageAt

`func (o *ThreadSummary) GetLastMessageAt() time.Time`

GetLastMessageAt returns the LastMessageAt field if non-nil, zero value otherwise.

### GetLastMessageAtOk

`func (o *ThreadSummary) GetLastMessageAtOk() (*time.Time, bool)`

GetLastMessageAtOk returns a tuple with the LastMessageAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageAt

`func (o *ThreadSummary) SetLastMessageAt(v time.Time)`

SetLastMessageAt sets LastMessageAt field to given value.

### HasLastMessageAt

`func (o *ThreadSummary) HasLastMessageAt() bool`

HasLastMessageAt returns a boolean if a field has been set.

### GetLastDirection

`func (o *ThreadSummary) GetLastDirection() string`

GetLastDirection returns the LastDirection field if non-nil, zero value otherwise.

### GetLastDirectionOk

`func (o *ThreadSummary) GetLastDirectionOk() (*string, bool)`

GetLastDirectionOk returns a tuple with the LastDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDirection

`func (o *ThreadSummary) SetLastDirection(v string)`

SetLastDirection sets LastDirection field to given value.

### HasLastDirection

`func (o *ThreadSummary) HasLastDirection() bool`

HasLastDirection returns a boolean if a field has been set.

### SetLastDirectionNil

`func (o *ThreadSummary) SetLastDirectionNil(b bool)`

 SetLastDirectionNil sets the value for LastDirection to be an explicit nil

### UnsetLastDirection
`func (o *ThreadSummary) UnsetLastDirection()`

UnsetLastDirection ensures that no value is present for LastDirection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


