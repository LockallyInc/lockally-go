# CalendarIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Provider** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**LastSyncAt** | Pointer to **time.Time** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCalendarIntegration

`func NewCalendarIntegration(id string, tenantId string, provider string, status string, createdAt time.Time, updatedAt time.Time, ) *CalendarIntegration`

NewCalendarIntegration instantiates a new CalendarIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarIntegrationWithDefaults

`func NewCalendarIntegrationWithDefaults() *CalendarIntegration`

NewCalendarIntegrationWithDefaults instantiates a new CalendarIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CalendarIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CalendarIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CalendarIntegration) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *CalendarIntegration) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CalendarIntegration) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CalendarIntegration) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetProvider

`func (o *CalendarIntegration) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CalendarIntegration) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CalendarIntegration) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetLabel

`func (o *CalendarIntegration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CalendarIntegration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CalendarIntegration) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CalendarIntegration) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetStatus

`func (o *CalendarIntegration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CalendarIntegration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CalendarIntegration) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastSyncAt

`func (o *CalendarIntegration) GetLastSyncAt() time.Time`

GetLastSyncAt returns the LastSyncAt field if non-nil, zero value otherwise.

### GetLastSyncAtOk

`func (o *CalendarIntegration) GetLastSyncAtOk() (*time.Time, bool)`

GetLastSyncAtOk returns a tuple with the LastSyncAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncAt

`func (o *CalendarIntegration) SetLastSyncAt(v time.Time)`

SetLastSyncAt sets LastSyncAt field to given value.

### HasLastSyncAt

`func (o *CalendarIntegration) HasLastSyncAt() bool`

HasLastSyncAt returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CalendarIntegration) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CalendarIntegration) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CalendarIntegration) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CalendarIntegration) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CalendarIntegration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CalendarIntegration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CalendarIntegration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CalendarIntegration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CalendarIntegration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CalendarIntegration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


