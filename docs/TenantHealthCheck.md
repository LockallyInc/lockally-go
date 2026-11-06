# TenantHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Stable identifier, today one of &#x60;domain_verification&#x60;, &#x60;smtp&#x60;, &#x60;storage&#x60;, &#x60;security&#x60;. Treat unknown keys as new checks, not errors. | 
**Label** | **string** | Human-readable name for display. | 
**Status** | **string** | &#x60;healthy&#x60;, &#x60;warning&#x60; or &#x60;critical&#x60;. | 
**Detail** | **string** |  | 

## Methods

### NewTenantHealthCheck

`func NewTenantHealthCheck(key string, label string, status string, detail string, ) *TenantHealthCheck`

NewTenantHealthCheck instantiates a new TenantHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantHealthCheckWithDefaults

`func NewTenantHealthCheckWithDefaults() *TenantHealthCheck`

NewTenantHealthCheckWithDefaults instantiates a new TenantHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TenantHealthCheck) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TenantHealthCheck) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TenantHealthCheck) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *TenantHealthCheck) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TenantHealthCheck) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TenantHealthCheck) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStatus

`func (o *TenantHealthCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TenantHealthCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TenantHealthCheck) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *TenantHealthCheck) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *TenantHealthCheck) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *TenantHealthCheck) SetDetail(v string)`

SetDetail sets Detail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


