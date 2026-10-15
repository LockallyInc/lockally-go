# ReturnPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Host** | **string** |  | 
**Records** | [**[]DNSRecord**](DNSRecord.md) |  | 
**Check** | Pointer to [**ReturnPathCheck**](ReturnPathCheck.md) |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**VerifiedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewReturnPath

`func NewReturnPath(status string, host string, records []DNSRecord, ) *ReturnPath`

NewReturnPath instantiates a new ReturnPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnPathWithDefaults

`func NewReturnPathWithDefaults() *ReturnPath`

NewReturnPathWithDefaults instantiates a new ReturnPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ReturnPath) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReturnPath) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReturnPath) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetHost

`func (o *ReturnPath) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ReturnPath) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ReturnPath) SetHost(v string)`

SetHost sets Host field to given value.


### GetRecords

`func (o *ReturnPath) GetRecords() []DNSRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ReturnPath) GetRecordsOk() (*[]DNSRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ReturnPath) SetRecords(v []DNSRecord)`

SetRecords sets Records field to given value.


### GetCheck

`func (o *ReturnPath) GetCheck() ReturnPathCheck`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *ReturnPath) GetCheckOk() (*ReturnPathCheck, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *ReturnPath) SetCheck(v ReturnPathCheck)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *ReturnPath) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### GetRequestedAt

`func (o *ReturnPath) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *ReturnPath) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *ReturnPath) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *ReturnPath) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *ReturnPath) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *ReturnPath) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *ReturnPath) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *ReturnPath) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


