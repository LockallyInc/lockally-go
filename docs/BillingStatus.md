# BillingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | **string** |  | 
**DisplayName** | **string** |  | 
**Mode** | **string** |  | 
**RateCapPerMin** | **int32** |  | 
**MonthlyIncludedSends** | **int32** |  | 
**MsgsThisPeriod** | **int32** |  | 
**Status** | **string** |  | 
**PriceNairaPerSeat** | **int32** |  | 
**SubscribedAt** | Pointer to **time.Time** |  | [optional] 
**CurrentPeriodEnd** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**SendUnitsBalance** | **int32** |  | 
**SendUnitsNextExpiry** | Pointer to **time.Time** |  | [optional] 
**UnitBundles** | [**[]UnitBundle**](UnitBundle.md) |  | 
**Catalog** | [**[]PlanCatalogEntry**](PlanCatalogEntry.md) |  | 

## Methods

### NewBillingStatus

`func NewBillingStatus(plan string, displayName string, mode string, rateCapPerMin int32, monthlyIncludedSends int32, msgsThisPeriod int32, status string, priceNairaPerSeat int32, createdAt time.Time, sendUnitsBalance int32, unitBundles []UnitBundle, catalog []PlanCatalogEntry, ) *BillingStatus`

NewBillingStatus instantiates a new BillingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingStatusWithDefaults

`func NewBillingStatusWithDefaults() *BillingStatus`

NewBillingStatusWithDefaults instantiates a new BillingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *BillingStatus) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *BillingStatus) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *BillingStatus) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetDisplayName

`func (o *BillingStatus) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BillingStatus) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BillingStatus) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMode

`func (o *BillingStatus) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BillingStatus) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BillingStatus) SetMode(v string)`

SetMode sets Mode field to given value.


### GetRateCapPerMin

`func (o *BillingStatus) GetRateCapPerMin() int32`

GetRateCapPerMin returns the RateCapPerMin field if non-nil, zero value otherwise.

### GetRateCapPerMinOk

`func (o *BillingStatus) GetRateCapPerMinOk() (*int32, bool)`

GetRateCapPerMinOk returns a tuple with the RateCapPerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCapPerMin

`func (o *BillingStatus) SetRateCapPerMin(v int32)`

SetRateCapPerMin sets RateCapPerMin field to given value.


### GetMonthlyIncludedSends

`func (o *BillingStatus) GetMonthlyIncludedSends() int32`

GetMonthlyIncludedSends returns the MonthlyIncludedSends field if non-nil, zero value otherwise.

### GetMonthlyIncludedSendsOk

`func (o *BillingStatus) GetMonthlyIncludedSendsOk() (*int32, bool)`

GetMonthlyIncludedSendsOk returns a tuple with the MonthlyIncludedSends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyIncludedSends

`func (o *BillingStatus) SetMonthlyIncludedSends(v int32)`

SetMonthlyIncludedSends sets MonthlyIncludedSends field to given value.


### GetMsgsThisPeriod

`func (o *BillingStatus) GetMsgsThisPeriod() int32`

GetMsgsThisPeriod returns the MsgsThisPeriod field if non-nil, zero value otherwise.

### GetMsgsThisPeriodOk

`func (o *BillingStatus) GetMsgsThisPeriodOk() (*int32, bool)`

GetMsgsThisPeriodOk returns a tuple with the MsgsThisPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgsThisPeriod

`func (o *BillingStatus) SetMsgsThisPeriod(v int32)`

SetMsgsThisPeriod sets MsgsThisPeriod field to given value.


### GetStatus

`func (o *BillingStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPriceNairaPerSeat

`func (o *BillingStatus) GetPriceNairaPerSeat() int32`

GetPriceNairaPerSeat returns the PriceNairaPerSeat field if non-nil, zero value otherwise.

### GetPriceNairaPerSeatOk

`func (o *BillingStatus) GetPriceNairaPerSeatOk() (*int32, bool)`

GetPriceNairaPerSeatOk returns a tuple with the PriceNairaPerSeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNairaPerSeat

`func (o *BillingStatus) SetPriceNairaPerSeat(v int32)`

SetPriceNairaPerSeat sets PriceNairaPerSeat field to given value.


### GetSubscribedAt

`func (o *BillingStatus) GetSubscribedAt() time.Time`

GetSubscribedAt returns the SubscribedAt field if non-nil, zero value otherwise.

### GetSubscribedAtOk

`func (o *BillingStatus) GetSubscribedAtOk() (*time.Time, bool)`

GetSubscribedAtOk returns a tuple with the SubscribedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedAt

`func (o *BillingStatus) SetSubscribedAt(v time.Time)`

SetSubscribedAt sets SubscribedAt field to given value.

### HasSubscribedAt

`func (o *BillingStatus) HasSubscribedAt() bool`

HasSubscribedAt returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *BillingStatus) GetCurrentPeriodEnd() time.Time`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *BillingStatus) GetCurrentPeriodEndOk() (*time.Time, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *BillingStatus) SetCurrentPeriodEnd(v time.Time)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *BillingStatus) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BillingStatus) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingStatus) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingStatus) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSendUnitsBalance

`func (o *BillingStatus) GetSendUnitsBalance() int32`

GetSendUnitsBalance returns the SendUnitsBalance field if non-nil, zero value otherwise.

### GetSendUnitsBalanceOk

`func (o *BillingStatus) GetSendUnitsBalanceOk() (*int32, bool)`

GetSendUnitsBalanceOk returns a tuple with the SendUnitsBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendUnitsBalance

`func (o *BillingStatus) SetSendUnitsBalance(v int32)`

SetSendUnitsBalance sets SendUnitsBalance field to given value.


### GetSendUnitsNextExpiry

`func (o *BillingStatus) GetSendUnitsNextExpiry() time.Time`

GetSendUnitsNextExpiry returns the SendUnitsNextExpiry field if non-nil, zero value otherwise.

### GetSendUnitsNextExpiryOk

`func (o *BillingStatus) GetSendUnitsNextExpiryOk() (*time.Time, bool)`

GetSendUnitsNextExpiryOk returns a tuple with the SendUnitsNextExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendUnitsNextExpiry

`func (o *BillingStatus) SetSendUnitsNextExpiry(v time.Time)`

SetSendUnitsNextExpiry sets SendUnitsNextExpiry field to given value.

### HasSendUnitsNextExpiry

`func (o *BillingStatus) HasSendUnitsNextExpiry() bool`

HasSendUnitsNextExpiry returns a boolean if a field has been set.

### GetUnitBundles

`func (o *BillingStatus) GetUnitBundles() []UnitBundle`

GetUnitBundles returns the UnitBundles field if non-nil, zero value otherwise.

### GetUnitBundlesOk

`func (o *BillingStatus) GetUnitBundlesOk() (*[]UnitBundle, bool)`

GetUnitBundlesOk returns a tuple with the UnitBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitBundles

`func (o *BillingStatus) SetUnitBundles(v []UnitBundle)`

SetUnitBundles sets UnitBundles field to given value.


### GetCatalog

`func (o *BillingStatus) GetCatalog() []PlanCatalogEntry`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *BillingStatus) GetCatalogOk() (*[]PlanCatalogEntry, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *BillingStatus) SetCatalog(v []PlanCatalogEntry)`

SetCatalog sets Catalog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


