# MigrationDNSCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ready** | **bool** |  | 
**Domains** | [**[]MigrationDomainDNS**](MigrationDomainDNS.md) |  | 

## Methods

### NewMigrationDNSCheck

`func NewMigrationDNSCheck(ready bool, domains []MigrationDomainDNS, ) *MigrationDNSCheck`

NewMigrationDNSCheck instantiates a new MigrationDNSCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationDNSCheckWithDefaults

`func NewMigrationDNSCheckWithDefaults() *MigrationDNSCheck`

NewMigrationDNSCheckWithDefaults instantiates a new MigrationDNSCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReady

`func (o *MigrationDNSCheck) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *MigrationDNSCheck) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *MigrationDNSCheck) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetDomains

`func (o *MigrationDNSCheck) GetDomains() []MigrationDomainDNS`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *MigrationDNSCheck) GetDomainsOk() (*[]MigrationDomainDNS, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *MigrationDNSCheck) SetDomains(v []MigrationDomainDNS)`

SetDomains sets Domains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


