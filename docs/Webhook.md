# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Url** | **string** |  | 
**Events** | **[]string** |  | 
**Paused** | **bool** |  | 
**PausedAt** | Pointer to **time.Time** |  | [optional] 
**LastSuccessAt** | Pointer to **time.Time** |  | [optional] 
**LastFailureAt** | Pointer to **time.Time** |  | [optional] 
**ConsecutiveFailures** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**SigningSecret** | Pointer to **string** | Hex-encoded HMAC-SHA256 key. Present ONLY on POST response. | [optional] 

## Methods

### NewWebhook

`func NewWebhook(id string, tenantId string, url string, events []string, paused bool, consecutiveFailures int32, createdAt time.Time, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Webhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *Webhook) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Webhook) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Webhook) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEvents

`func (o *Webhook) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Webhook) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Webhook) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetPaused

`func (o *Webhook) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *Webhook) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *Webhook) SetPaused(v bool)`

SetPaused sets Paused field to given value.


### GetPausedAt

`func (o *Webhook) GetPausedAt() time.Time`

GetPausedAt returns the PausedAt field if non-nil, zero value otherwise.

### GetPausedAtOk

`func (o *Webhook) GetPausedAtOk() (*time.Time, bool)`

GetPausedAtOk returns a tuple with the PausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedAt

`func (o *Webhook) SetPausedAt(v time.Time)`

SetPausedAt sets PausedAt field to given value.

### HasPausedAt

`func (o *Webhook) HasPausedAt() bool`

HasPausedAt returns a boolean if a field has been set.

### GetLastSuccessAt

`func (o *Webhook) GetLastSuccessAt() time.Time`

GetLastSuccessAt returns the LastSuccessAt field if non-nil, zero value otherwise.

### GetLastSuccessAtOk

`func (o *Webhook) GetLastSuccessAtOk() (*time.Time, bool)`

GetLastSuccessAtOk returns a tuple with the LastSuccessAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessAt

`func (o *Webhook) SetLastSuccessAt(v time.Time)`

SetLastSuccessAt sets LastSuccessAt field to given value.

### HasLastSuccessAt

`func (o *Webhook) HasLastSuccessAt() bool`

HasLastSuccessAt returns a boolean if a field has been set.

### GetLastFailureAt

`func (o *Webhook) GetLastFailureAt() time.Time`

GetLastFailureAt returns the LastFailureAt field if non-nil, zero value otherwise.

### GetLastFailureAtOk

`func (o *Webhook) GetLastFailureAtOk() (*time.Time, bool)`

GetLastFailureAtOk returns a tuple with the LastFailureAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureAt

`func (o *Webhook) SetLastFailureAt(v time.Time)`

SetLastFailureAt sets LastFailureAt field to given value.

### HasLastFailureAt

`func (o *Webhook) HasLastFailureAt() bool`

HasLastFailureAt returns a boolean if a field has been set.

### GetConsecutiveFailures

`func (o *Webhook) GetConsecutiveFailures() int32`

GetConsecutiveFailures returns the ConsecutiveFailures field if non-nil, zero value otherwise.

### GetConsecutiveFailuresOk

`func (o *Webhook) GetConsecutiveFailuresOk() (*int32, bool)`

GetConsecutiveFailuresOk returns a tuple with the ConsecutiveFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveFailures

`func (o *Webhook) SetConsecutiveFailures(v int32)`

SetConsecutiveFailures sets ConsecutiveFailures field to given value.


### GetCreatedAt

`func (o *Webhook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Webhook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Webhook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSigningSecret

`func (o *Webhook) GetSigningSecret() string`

GetSigningSecret returns the SigningSecret field if non-nil, zero value otherwise.

### GetSigningSecretOk

`func (o *Webhook) GetSigningSecretOk() (*string, bool)`

GetSigningSecretOk returns a tuple with the SigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecret

`func (o *Webhook) SetSigningSecret(v string)`

SetSigningSecret sets SigningSecret field to given value.

### HasSigningSecret

`func (o *Webhook) HasSigningSecret() bool`

HasSigningSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


