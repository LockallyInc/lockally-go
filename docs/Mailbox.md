# Mailbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**DomainId** | **string** |  | 
**Email** | **string** |  | 
**QuotaBytes** | **int64** |  | 
**Disabled** | **bool** |  | 
**DisabledAt** | Pointer to **time.Time** |  | [optional] 
**SoftDeletedAt** | Pointer to **time.Time** |  | [optional] 
**HardDeleteAfter** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Password** | Pointer to **string** | ONLY present on POST response when lockally generated the password. Shown once. | [optional] 

## Methods

### NewMailbox

`func NewMailbox(id string, tenantId string, domainId string, email string, quotaBytes int64, disabled bool, createdAt time.Time, ) *Mailbox`

NewMailbox instantiates a new Mailbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxWithDefaults

`func NewMailboxWithDefaults() *Mailbox`

NewMailboxWithDefaults instantiates a new Mailbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Mailbox) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Mailbox) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Mailbox) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *Mailbox) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Mailbox) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Mailbox) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetDomainId

`func (o *Mailbox) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *Mailbox) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *Mailbox) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetEmail

`func (o *Mailbox) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Mailbox) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Mailbox) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetQuotaBytes

`func (o *Mailbox) GetQuotaBytes() int64`

GetQuotaBytes returns the QuotaBytes field if non-nil, zero value otherwise.

### GetQuotaBytesOk

`func (o *Mailbox) GetQuotaBytesOk() (*int64, bool)`

GetQuotaBytesOk returns a tuple with the QuotaBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaBytes

`func (o *Mailbox) SetQuotaBytes(v int64)`

SetQuotaBytes sets QuotaBytes field to given value.


### GetDisabled

`func (o *Mailbox) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Mailbox) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Mailbox) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetDisabledAt

`func (o *Mailbox) GetDisabledAt() time.Time`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *Mailbox) GetDisabledAtOk() (*time.Time, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *Mailbox) SetDisabledAt(v time.Time)`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *Mailbox) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### GetSoftDeletedAt

`func (o *Mailbox) GetSoftDeletedAt() time.Time`

GetSoftDeletedAt returns the SoftDeletedAt field if non-nil, zero value otherwise.

### GetSoftDeletedAtOk

`func (o *Mailbox) GetSoftDeletedAtOk() (*time.Time, bool)`

GetSoftDeletedAtOk returns a tuple with the SoftDeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeletedAt

`func (o *Mailbox) SetSoftDeletedAt(v time.Time)`

SetSoftDeletedAt sets SoftDeletedAt field to given value.

### HasSoftDeletedAt

`func (o *Mailbox) HasSoftDeletedAt() bool`

HasSoftDeletedAt returns a boolean if a field has been set.

### GetHardDeleteAfter

`func (o *Mailbox) GetHardDeleteAfter() time.Time`

GetHardDeleteAfter returns the HardDeleteAfter field if non-nil, zero value otherwise.

### GetHardDeleteAfterOk

`func (o *Mailbox) GetHardDeleteAfterOk() (*time.Time, bool)`

GetHardDeleteAfterOk returns a tuple with the HardDeleteAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDeleteAfter

`func (o *Mailbox) SetHardDeleteAfter(v time.Time)`

SetHardDeleteAfter sets HardDeleteAfter field to given value.

### HasHardDeleteAfter

`func (o *Mailbox) HasHardDeleteAfter() bool`

HasHardDeleteAfter returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Mailbox) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Mailbox) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Mailbox) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPassword

`func (o *Mailbox) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Mailbox) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Mailbox) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Mailbox) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


