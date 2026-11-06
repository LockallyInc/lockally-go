# TrackingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackOpens** | **bool** |  | 
**TrackClicks** | **bool** |  | 
**UpdatedAt** | Pointer to **time.Time** | Absent until a policy has been set. | [optional] 
**Configured** | **bool** | Whether this deployment holds a tracking signing secret. False means tracking cannot work whatever the flags say. | 

## Methods

### NewTrackingPolicy

`func NewTrackingPolicy(trackOpens bool, trackClicks bool, configured bool, ) *TrackingPolicy`

NewTrackingPolicy instantiates a new TrackingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingPolicyWithDefaults

`func NewTrackingPolicyWithDefaults() *TrackingPolicy`

NewTrackingPolicyWithDefaults instantiates a new TrackingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackOpens

`func (o *TrackingPolicy) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *TrackingPolicy) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *TrackingPolicy) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.


### GetTrackClicks

`func (o *TrackingPolicy) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *TrackingPolicy) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *TrackingPolicy) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.


### GetUpdatedAt

`func (o *TrackingPolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TrackingPolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TrackingPolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TrackingPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetConfigured

`func (o *TrackingPolicy) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *TrackingPolicy) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *TrackingPolicy) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


