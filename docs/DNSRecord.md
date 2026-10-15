# DNSRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**Value** | **string** |  | 
**Ttl** | **int32** |  | 
**Priority** | Pointer to **int32** | MX records only. | [optional] 
**Purpose** | **string** |  | 

## Methods

### NewDNSRecord

`func NewDNSRecord(type_ string, name string, value string, ttl int32, purpose string, ) *DNSRecord`

NewDNSRecord instantiates a new DNSRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordWithDefaults

`func NewDNSRecordWithDefaults() *DNSRecord`

NewDNSRecordWithDefaults instantiates a new DNSRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DNSRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSRecord) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *DNSRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSRecord) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *DNSRecord) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DNSRecord) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DNSRecord) SetValue(v string)`

SetValue sets Value field to given value.


### GetTtl

`func (o *DNSRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### GetPriority

`func (o *DNSRecord) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DNSRecord) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DNSRecord) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DNSRecord) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPurpose

`func (o *DNSRecord) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *DNSRecord) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *DNSRecord) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


