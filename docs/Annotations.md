# Annotations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentTypes** | Pointer to **[]string** |  | [optional] 
**MeetingRequest** | Pointer to **bool** |  | [optional] 
**InjectionRisk** | Pointer to [**InjectionRisk**](InjectionRisk.md) |  | [optional] 

## Methods

### NewAnnotations

`func NewAnnotations() *Annotations`

NewAnnotations instantiates a new Annotations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationsWithDefaults

`func NewAnnotationsWithDefaults() *Annotations`

NewAnnotationsWithDefaults instantiates a new Annotations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentTypes

`func (o *Annotations) GetAttachmentTypes() []string`

GetAttachmentTypes returns the AttachmentTypes field if non-nil, zero value otherwise.

### GetAttachmentTypesOk

`func (o *Annotations) GetAttachmentTypesOk() (*[]string, bool)`

GetAttachmentTypesOk returns a tuple with the AttachmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentTypes

`func (o *Annotations) SetAttachmentTypes(v []string)`

SetAttachmentTypes sets AttachmentTypes field to given value.

### HasAttachmentTypes

`func (o *Annotations) HasAttachmentTypes() bool`

HasAttachmentTypes returns a boolean if a field has been set.

### GetMeetingRequest

`func (o *Annotations) GetMeetingRequest() bool`

GetMeetingRequest returns the MeetingRequest field if non-nil, zero value otherwise.

### GetMeetingRequestOk

`func (o *Annotations) GetMeetingRequestOk() (*bool, bool)`

GetMeetingRequestOk returns a tuple with the MeetingRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingRequest

`func (o *Annotations) SetMeetingRequest(v bool)`

SetMeetingRequest sets MeetingRequest field to given value.

### HasMeetingRequest

`func (o *Annotations) HasMeetingRequest() bool`

HasMeetingRequest returns a boolean if a field has been set.

### GetInjectionRisk

`func (o *Annotations) GetInjectionRisk() InjectionRisk`

GetInjectionRisk returns the InjectionRisk field if non-nil, zero value otherwise.

### GetInjectionRiskOk

`func (o *Annotations) GetInjectionRiskOk() (*InjectionRisk, bool)`

GetInjectionRiskOk returns a tuple with the InjectionRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectionRisk

`func (o *Annotations) SetInjectionRisk(v InjectionRisk)`

SetInjectionRisk sets InjectionRisk field to given value.

### HasInjectionRisk

`func (o *Annotations) HasInjectionRisk() bool`

HasInjectionRisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


