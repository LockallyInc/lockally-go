# TenantHealthReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverallScore** | **int32** |  | 
**Checks** | [**[]TenantHealthCheck**](TenantHealthCheck.md) |  | 
**GeneratedAt** | **time.Time** |  | 

## Methods

### NewTenantHealthReport

`func NewTenantHealthReport(overallScore int32, checks []TenantHealthCheck, generatedAt time.Time, ) *TenantHealthReport`

NewTenantHealthReport instantiates a new TenantHealthReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantHealthReportWithDefaults

`func NewTenantHealthReportWithDefaults() *TenantHealthReport`

NewTenantHealthReportWithDefaults instantiates a new TenantHealthReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverallScore

`func (o *TenantHealthReport) GetOverallScore() int32`

GetOverallScore returns the OverallScore field if non-nil, zero value otherwise.

### GetOverallScoreOk

`func (o *TenantHealthReport) GetOverallScoreOk() (*int32, bool)`

GetOverallScoreOk returns a tuple with the OverallScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallScore

`func (o *TenantHealthReport) SetOverallScore(v int32)`

SetOverallScore sets OverallScore field to given value.


### GetChecks

`func (o *TenantHealthReport) GetChecks() []TenantHealthCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *TenantHealthReport) GetChecksOk() (*[]TenantHealthCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *TenantHealthReport) SetChecks(v []TenantHealthCheck)`

SetChecks sets Checks field to given value.


### GetGeneratedAt

`func (o *TenantHealthReport) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *TenantHealthReport) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *TenantHealthReport) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


