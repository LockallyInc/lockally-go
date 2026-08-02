# HealthzGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Time** | **time.Time** |  | 

## Methods

### NewHealthzGet200Response

`func NewHealthzGet200Response(status string, time time.Time, ) *HealthzGet200Response`

NewHealthzGet200Response instantiates a new HealthzGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthzGet200ResponseWithDefaults

`func NewHealthzGet200ResponseWithDefaults() *HealthzGet200Response`

NewHealthzGet200ResponseWithDefaults instantiates a new HealthzGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HealthzGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthzGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthzGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTime

`func (o *HealthzGet200Response) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *HealthzGet200Response) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *HealthzGet200Response) SetTime(v time.Time)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


