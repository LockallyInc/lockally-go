# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Domain** | **string** |  | 
**VerificationToken** | **string** |  | 
**Verified** | **bool** |  | 
**VerifiedAt** | Pointer to **time.Time** |  | [optional] 
**DkimSelector** | **string** |  | 
**DkimPublicRecord** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Records** | Pointer to [**[]DNSRecord**](DNSRecord.md) | DNS records the tenant must publish under their own DNS. | [optional] 

## Methods

### NewDomain

`func NewDomain(id string, tenantId string, domain string, verificationToken string, verified bool, dkimSelector string, dkimPublicRecord string, createdAt time.Time, ) *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Domain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *Domain) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Domain) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Domain) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetDomain

`func (o *Domain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Domain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Domain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetVerificationToken

`func (o *Domain) GetVerificationToken() string`

GetVerificationToken returns the VerificationToken field if non-nil, zero value otherwise.

### GetVerificationTokenOk

`func (o *Domain) GetVerificationTokenOk() (*string, bool)`

GetVerificationTokenOk returns a tuple with the VerificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationToken

`func (o *Domain) SetVerificationToken(v string)`

SetVerificationToken sets VerificationToken field to given value.


### GetVerified

`func (o *Domain) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *Domain) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *Domain) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetVerifiedAt

`func (o *Domain) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *Domain) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *Domain) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *Domain) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### GetDkimSelector

`func (o *Domain) GetDkimSelector() string`

GetDkimSelector returns the DkimSelector field if non-nil, zero value otherwise.

### GetDkimSelectorOk

`func (o *Domain) GetDkimSelectorOk() (*string, bool)`

GetDkimSelectorOk returns a tuple with the DkimSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimSelector

`func (o *Domain) SetDkimSelector(v string)`

SetDkimSelector sets DkimSelector field to given value.


### GetDkimPublicRecord

`func (o *Domain) GetDkimPublicRecord() string`

GetDkimPublicRecord returns the DkimPublicRecord field if non-nil, zero value otherwise.

### GetDkimPublicRecordOk

`func (o *Domain) GetDkimPublicRecordOk() (*string, bool)`

GetDkimPublicRecordOk returns a tuple with the DkimPublicRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimPublicRecord

`func (o *Domain) SetDkimPublicRecord(v string)`

SetDkimPublicRecord sets DkimPublicRecord field to given value.


### GetCreatedAt

`func (o *Domain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Domain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Domain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetRecords

`func (o *Domain) GetRecords() []DNSRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *Domain) GetRecordsOk() (*[]DNSRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *Domain) SetRecords(v []DNSRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *Domain) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


