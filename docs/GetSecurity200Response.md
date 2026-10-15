# GetSecurity200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverallStatus** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**[]GetSecurity200ResponseStatsInner**](GetSecurity200ResponseStatsInner.md) |  | [optional] 
**Alerts** | Pointer to [**[]GetSecurity200ResponseAlertsInner**](GetSecurity200ResponseAlertsInner.md) |  | [optional] 

## Methods

### NewGetSecurity200Response

`func NewGetSecurity200Response() *GetSecurity200Response`

NewGetSecurity200Response instantiates a new GetSecurity200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecurity200ResponseWithDefaults

`func NewGetSecurity200ResponseWithDefaults() *GetSecurity200Response`

NewGetSecurity200ResponseWithDefaults instantiates a new GetSecurity200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverallStatus

`func (o *GetSecurity200Response) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *GetSecurity200Response) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *GetSecurity200Response) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *GetSecurity200Response) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### GetStats

`func (o *GetSecurity200Response) GetStats() []GetSecurity200ResponseStatsInner`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetSecurity200Response) GetStatsOk() (*[]GetSecurity200ResponseStatsInner, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetSecurity200Response) SetStats(v []GetSecurity200ResponseStatsInner)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetSecurity200Response) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetAlerts

`func (o *GetSecurity200Response) GetAlerts() []GetSecurity200ResponseAlertsInner`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *GetSecurity200Response) GetAlertsOk() (*[]GetSecurity200ResponseAlertsInner, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *GetSecurity200Response) SetAlerts(v []GetSecurity200ResponseAlertsInner)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *GetSecurity200Response) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


