# GetDomainsStatus200ResponseDomainsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**Checks** | Pointer to [**[]GetDomainsStatus200ResponseDomainsInnerChecksInner**](GetDomainsStatus200ResponseDomainsInnerChecksInner.md) |  | [optional] 

## Methods

### NewGetDomainsStatus200ResponseDomainsInner

`func NewGetDomainsStatus200ResponseDomainsInner() *GetDomainsStatus200ResponseDomainsInner`

NewGetDomainsStatus200ResponseDomainsInner instantiates a new GetDomainsStatus200ResponseDomainsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainsStatus200ResponseDomainsInnerWithDefaults

`func NewGetDomainsStatus200ResponseDomainsInnerWithDefaults() *GetDomainsStatus200ResponseDomainsInner`

NewGetDomainsStatus200ResponseDomainsInnerWithDefaults instantiates a new GetDomainsStatus200ResponseDomainsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *GetDomainsStatus200ResponseDomainsInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetDomainsStatus200ResponseDomainsInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetDomainsStatus200ResponseDomainsInner) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetDomainsStatus200ResponseDomainsInner) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetVerified

`func (o *GetDomainsStatus200ResponseDomainsInner) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *GetDomainsStatus200ResponseDomainsInner) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *GetDomainsStatus200ResponseDomainsInner) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *GetDomainsStatus200ResponseDomainsInner) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetChecks

`func (o *GetDomainsStatus200ResponseDomainsInner) GetChecks() []GetDomainsStatus200ResponseDomainsInnerChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *GetDomainsStatus200ResponseDomainsInner) GetChecksOk() (*[]GetDomainsStatus200ResponseDomainsInnerChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *GetDomainsStatus200ResponseDomainsInner) SetChecks(v []GetDomainsStatus200ResponseDomainsInnerChecksInner)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *GetDomainsStatus200ResponseDomainsInner) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


