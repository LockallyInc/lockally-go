# GetIntegrationsSummary200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiRequestsToday** | Pointer to **int32** |  | [optional] 
**ActiveApiKeys** | Pointer to **int32** |  | [optional] 
**ApiKeys** | Pointer to [**[]GetIntegrationsSummary200ResponseApiKeysInner**](GetIntegrationsSummary200ResponseApiKeysInner.md) |  | [optional] 
**WebhookFailures** | Pointer to **int32** |  | [optional] 
**WebhooksTotal** | Pointer to **int32** |  | [optional] 
**Webhooks** | Pointer to [**[]GetIntegrationsSummary200ResponseWebhooksInner**](GetIntegrationsSummary200ResponseWebhooksInner.md) |  | [optional] 
**GeneratedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetIntegrationsSummary200Response

`func NewGetIntegrationsSummary200Response() *GetIntegrationsSummary200Response`

NewGetIntegrationsSummary200Response instantiates a new GetIntegrationsSummary200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIntegrationsSummary200ResponseWithDefaults

`func NewGetIntegrationsSummary200ResponseWithDefaults() *GetIntegrationsSummary200Response`

NewGetIntegrationsSummary200ResponseWithDefaults instantiates a new GetIntegrationsSummary200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiRequestsToday

`func (o *GetIntegrationsSummary200Response) GetApiRequestsToday() int32`

GetApiRequestsToday returns the ApiRequestsToday field if non-nil, zero value otherwise.

### GetApiRequestsTodayOk

`func (o *GetIntegrationsSummary200Response) GetApiRequestsTodayOk() (*int32, bool)`

GetApiRequestsTodayOk returns a tuple with the ApiRequestsToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRequestsToday

`func (o *GetIntegrationsSummary200Response) SetApiRequestsToday(v int32)`

SetApiRequestsToday sets ApiRequestsToday field to given value.

### HasApiRequestsToday

`func (o *GetIntegrationsSummary200Response) HasApiRequestsToday() bool`

HasApiRequestsToday returns a boolean if a field has been set.

### GetActiveApiKeys

`func (o *GetIntegrationsSummary200Response) GetActiveApiKeys() int32`

GetActiveApiKeys returns the ActiveApiKeys field if non-nil, zero value otherwise.

### GetActiveApiKeysOk

`func (o *GetIntegrationsSummary200Response) GetActiveApiKeysOk() (*int32, bool)`

GetActiveApiKeysOk returns a tuple with the ActiveApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveApiKeys

`func (o *GetIntegrationsSummary200Response) SetActiveApiKeys(v int32)`

SetActiveApiKeys sets ActiveApiKeys field to given value.

### HasActiveApiKeys

`func (o *GetIntegrationsSummary200Response) HasActiveApiKeys() bool`

HasActiveApiKeys returns a boolean if a field has been set.

### GetApiKeys

`func (o *GetIntegrationsSummary200Response) GetApiKeys() []GetIntegrationsSummary200ResponseApiKeysInner`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *GetIntegrationsSummary200Response) GetApiKeysOk() (*[]GetIntegrationsSummary200ResponseApiKeysInner, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *GetIntegrationsSummary200Response) SetApiKeys(v []GetIntegrationsSummary200ResponseApiKeysInner)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *GetIntegrationsSummary200Response) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### GetWebhookFailures

`func (o *GetIntegrationsSummary200Response) GetWebhookFailures() int32`

GetWebhookFailures returns the WebhookFailures field if non-nil, zero value otherwise.

### GetWebhookFailuresOk

`func (o *GetIntegrationsSummary200Response) GetWebhookFailuresOk() (*int32, bool)`

GetWebhookFailuresOk returns a tuple with the WebhookFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailures

`func (o *GetIntegrationsSummary200Response) SetWebhookFailures(v int32)`

SetWebhookFailures sets WebhookFailures field to given value.

### HasWebhookFailures

`func (o *GetIntegrationsSummary200Response) HasWebhookFailures() bool`

HasWebhookFailures returns a boolean if a field has been set.

### GetWebhooksTotal

`func (o *GetIntegrationsSummary200Response) GetWebhooksTotal() int32`

GetWebhooksTotal returns the WebhooksTotal field if non-nil, zero value otherwise.

### GetWebhooksTotalOk

`func (o *GetIntegrationsSummary200Response) GetWebhooksTotalOk() (*int32, bool)`

GetWebhooksTotalOk returns a tuple with the WebhooksTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksTotal

`func (o *GetIntegrationsSummary200Response) SetWebhooksTotal(v int32)`

SetWebhooksTotal sets WebhooksTotal field to given value.

### HasWebhooksTotal

`func (o *GetIntegrationsSummary200Response) HasWebhooksTotal() bool`

HasWebhooksTotal returns a boolean if a field has been set.

### GetWebhooks

`func (o *GetIntegrationsSummary200Response) GetWebhooks() []GetIntegrationsSummary200ResponseWebhooksInner`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *GetIntegrationsSummary200Response) GetWebhooksOk() (*[]GetIntegrationsSummary200ResponseWebhooksInner, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *GetIntegrationsSummary200Response) SetWebhooks(v []GetIntegrationsSummary200ResponseWebhooksInner)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *GetIntegrationsSummary200Response) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *GetIntegrationsSummary200Response) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *GetIntegrationsSummary200Response) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *GetIntegrationsSummary200Response) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *GetIntegrationsSummary200Response) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


