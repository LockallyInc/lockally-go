# MigrationDomainDNS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** |  | 
**Mx** | **string** |  | 
**Spf** | **string** |  | 
**Dkim** | **string** |  | 
**Dmarc** | **string** |  | 

## Methods

### NewMigrationDomainDNS

`func NewMigrationDomainDNS(domain string, mx string, spf string, dkim string, dmarc string, ) *MigrationDomainDNS`

NewMigrationDomainDNS instantiates a new MigrationDomainDNS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationDomainDNSWithDefaults

`func NewMigrationDomainDNSWithDefaults() *MigrationDomainDNS`

NewMigrationDomainDNSWithDefaults instantiates a new MigrationDomainDNS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *MigrationDomainDNS) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *MigrationDomainDNS) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *MigrationDomainDNS) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetMx

`func (o *MigrationDomainDNS) GetMx() string`

GetMx returns the Mx field if non-nil, zero value otherwise.

### GetMxOk

`func (o *MigrationDomainDNS) GetMxOk() (*string, bool)`

GetMxOk returns a tuple with the Mx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMx

`func (o *MigrationDomainDNS) SetMx(v string)`

SetMx sets Mx field to given value.


### GetSpf

`func (o *MigrationDomainDNS) GetSpf() string`

GetSpf returns the Spf field if non-nil, zero value otherwise.

### GetSpfOk

`func (o *MigrationDomainDNS) GetSpfOk() (*string, bool)`

GetSpfOk returns a tuple with the Spf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpf

`func (o *MigrationDomainDNS) SetSpf(v string)`

SetSpf sets Spf field to given value.


### GetDkim

`func (o *MigrationDomainDNS) GetDkim() string`

GetDkim returns the Dkim field if non-nil, zero value otherwise.

### GetDkimOk

`func (o *MigrationDomainDNS) GetDkimOk() (*string, bool)`

GetDkimOk returns a tuple with the Dkim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkim

`func (o *MigrationDomainDNS) SetDkim(v string)`

SetDkim sets Dkim field to given value.


### GetDmarc

`func (o *MigrationDomainDNS) GetDmarc() string`

GetDmarc returns the Dmarc field if non-nil, zero value otherwise.

### GetDmarcOk

`func (o *MigrationDomainDNS) GetDmarcOk() (*string, bool)`

GetDmarcOk returns a tuple with the Dmarc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarc

`func (o *MigrationDomainDNS) SetDmarc(v string)`

SetDmarc sets Dmarc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


