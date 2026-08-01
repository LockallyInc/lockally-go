# GetDomainsStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to [**[]GetDomainsStatus200ResponseDomainsInner**](GetDomainsStatus200ResponseDomainsInner.md) |  | [optional] 
**GeneratedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetDomainsStatus200Response

`func NewGetDomainsStatus200Response() *GetDomainsStatus200Response`

NewGetDomainsStatus200Response instantiates a new GetDomainsStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainsStatus200ResponseWithDefaults

`func NewGetDomainsStatus200ResponseWithDefaults() *GetDomainsStatus200Response`

NewGetDomainsStatus200ResponseWithDefaults instantiates a new GetDomainsStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *GetDomainsStatus200Response) GetDomains() []GetDomainsStatus200ResponseDomainsInner`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *GetDomainsStatus200Response) GetDomainsOk() (*[]GetDomainsStatus200ResponseDomainsInner, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *GetDomainsStatus200Response) SetDomains(v []GetDomainsStatus200ResponseDomainsInner)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *GetDomainsStatus200Response) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *GetDomainsStatus200Response) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *GetDomainsStatus200Response) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *GetDomainsStatus200Response) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *GetDomainsStatus200Response) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


