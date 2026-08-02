# PlanCatalogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DisplayName** | **string** |  | 
**Description** | **string** |  | 
**PriceNairaPerSeat** | **int32** |  | 
**RateCapPerMin** | **int32** |  | 
**MonthlyIncludedSends** | **int32** |  | 
**HasSharedMailboxes** | **bool** |  | 
**HasSendUnits** | **bool** |  | 
**HasAiUnits** | **bool** |  | 
**HasE2eEncryption** | **bool** |  | 

## Methods

### NewPlanCatalogEntry

`func NewPlanCatalogEntry(name string, displayName string, description string, priceNairaPerSeat int32, rateCapPerMin int32, monthlyIncludedSends int32, hasSharedMailboxes bool, hasSendUnits bool, hasAiUnits bool, hasE2eEncryption bool, ) *PlanCatalogEntry`

NewPlanCatalogEntry instantiates a new PlanCatalogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanCatalogEntryWithDefaults

`func NewPlanCatalogEntryWithDefaults() *PlanCatalogEntry`

NewPlanCatalogEntryWithDefaults instantiates a new PlanCatalogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlanCatalogEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanCatalogEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanCatalogEntry) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *PlanCatalogEntry) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PlanCatalogEntry) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PlanCatalogEntry) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *PlanCatalogEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanCatalogEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanCatalogEntry) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPriceNairaPerSeat

`func (o *PlanCatalogEntry) GetPriceNairaPerSeat() int32`

GetPriceNairaPerSeat returns the PriceNairaPerSeat field if non-nil, zero value otherwise.

### GetPriceNairaPerSeatOk

`func (o *PlanCatalogEntry) GetPriceNairaPerSeatOk() (*int32, bool)`

GetPriceNairaPerSeatOk returns a tuple with the PriceNairaPerSeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNairaPerSeat

`func (o *PlanCatalogEntry) SetPriceNairaPerSeat(v int32)`

SetPriceNairaPerSeat sets PriceNairaPerSeat field to given value.


### GetRateCapPerMin

`func (o *PlanCatalogEntry) GetRateCapPerMin() int32`

GetRateCapPerMin returns the RateCapPerMin field if non-nil, zero value otherwise.

### GetRateCapPerMinOk

`func (o *PlanCatalogEntry) GetRateCapPerMinOk() (*int32, bool)`

GetRateCapPerMinOk returns a tuple with the RateCapPerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCapPerMin

`func (o *PlanCatalogEntry) SetRateCapPerMin(v int32)`

SetRateCapPerMin sets RateCapPerMin field to given value.


### GetMonthlyIncludedSends

`func (o *PlanCatalogEntry) GetMonthlyIncludedSends() int32`

GetMonthlyIncludedSends returns the MonthlyIncludedSends field if non-nil, zero value otherwise.

### GetMonthlyIncludedSendsOk

`func (o *PlanCatalogEntry) GetMonthlyIncludedSendsOk() (*int32, bool)`

GetMonthlyIncludedSendsOk returns a tuple with the MonthlyIncludedSends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyIncludedSends

`func (o *PlanCatalogEntry) SetMonthlyIncludedSends(v int32)`

SetMonthlyIncludedSends sets MonthlyIncludedSends field to given value.


### GetHasSharedMailboxes

`func (o *PlanCatalogEntry) GetHasSharedMailboxes() bool`

GetHasSharedMailboxes returns the HasSharedMailboxes field if non-nil, zero value otherwise.

### GetHasSharedMailboxesOk

`func (o *PlanCatalogEntry) GetHasSharedMailboxesOk() (*bool, bool)`

GetHasSharedMailboxesOk returns a tuple with the HasSharedMailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSharedMailboxes

`func (o *PlanCatalogEntry) SetHasSharedMailboxes(v bool)`

SetHasSharedMailboxes sets HasSharedMailboxes field to given value.


### GetHasSendUnits

`func (o *PlanCatalogEntry) GetHasSendUnits() bool`

GetHasSendUnits returns the HasSendUnits field if non-nil, zero value otherwise.

### GetHasSendUnitsOk

`func (o *PlanCatalogEntry) GetHasSendUnitsOk() (*bool, bool)`

GetHasSendUnitsOk returns a tuple with the HasSendUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSendUnits

`func (o *PlanCatalogEntry) SetHasSendUnits(v bool)`

SetHasSendUnits sets HasSendUnits field to given value.


### GetHasAiUnits

`func (o *PlanCatalogEntry) GetHasAiUnits() bool`

GetHasAiUnits returns the HasAiUnits field if non-nil, zero value otherwise.

### GetHasAiUnitsOk

`func (o *PlanCatalogEntry) GetHasAiUnitsOk() (*bool, bool)`

GetHasAiUnitsOk returns a tuple with the HasAiUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAiUnits

`func (o *PlanCatalogEntry) SetHasAiUnits(v bool)`

SetHasAiUnits sets HasAiUnits field to given value.


### GetHasE2eEncryption

`func (o *PlanCatalogEntry) GetHasE2eEncryption() bool`

GetHasE2eEncryption returns the HasE2eEncryption field if non-nil, zero value otherwise.

### GetHasE2eEncryptionOk

`func (o *PlanCatalogEntry) GetHasE2eEncryptionOk() (*bool, bool)`

GetHasE2eEncryptionOk returns a tuple with the HasE2eEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasE2eEncryption

`func (o *PlanCatalogEntry) SetHasE2eEncryption(v bool)`

SetHasE2eEncryption sets HasE2eEncryption field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


