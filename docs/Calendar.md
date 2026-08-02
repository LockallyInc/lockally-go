# Calendar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Name** | **string** |  | 
**Color** | Pointer to **string** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | **string** |  | 
**FeedUrl** | Pointer to **string** |  | [optional] 
**EventCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCalendar

`func NewCalendar(id string, tenantId string, name string, visibility string, createdAt time.Time, updatedAt time.Time, ) *Calendar`

NewCalendar instantiates a new Calendar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarWithDefaults

`func NewCalendarWithDefaults() *Calendar`

NewCalendarWithDefaults instantiates a new Calendar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Calendar) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Calendar) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Calendar) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *Calendar) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Calendar) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Calendar) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetName

`func (o *Calendar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Calendar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Calendar) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *Calendar) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Calendar) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Calendar) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Calendar) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *Calendar) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *Calendar) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *Calendar) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *Calendar) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetDescription

`func (o *Calendar) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Calendar) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Calendar) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Calendar) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *Calendar) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Calendar) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Calendar) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetFeedUrl

`func (o *Calendar) GetFeedUrl() string`

GetFeedUrl returns the FeedUrl field if non-nil, zero value otherwise.

### GetFeedUrlOk

`func (o *Calendar) GetFeedUrlOk() (*string, bool)`

GetFeedUrlOk returns a tuple with the FeedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedUrl

`func (o *Calendar) SetFeedUrl(v string)`

SetFeedUrl sets FeedUrl field to given value.

### HasFeedUrl

`func (o *Calendar) HasFeedUrl() bool`

HasFeedUrl returns a boolean if a field has been set.

### GetEventCount

`func (o *Calendar) GetEventCount() int32`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *Calendar) GetEventCountOk() (*int32, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *Calendar) SetEventCount(v int32)`

SetEventCount sets EventCount field to given value.

### HasEventCount

`func (o *Calendar) HasEventCount() bool`

HasEventCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Calendar) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Calendar) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Calendar) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Calendar) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Calendar) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Calendar) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


