# V1UsageGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailboxesActive** | **int32** | Mailboxes that are neither disabled nor soft-deleted. | 
**MailboxesTotal** | Pointer to **int32** | All mailboxes for this tenant, including disabled/soft-deleted. | [optional] 
**DomainsVerified** | Pointer to **int32** | Domains that have passed DNS verification. | [optional] 
**DomainsTotal** | Pointer to **int32** |  | [optional] 
**MessagesSentLast60s** | Pointer to **int32** | Sends in the 60-second window ending now. Used by the rate-cap check. | [optional] 
**MessagesSentTodayUtc** | Pointer to **int32** | Sends since 00:00 UTC. Compared against &#x60;daily_msg_quota&#x60;. | [optional] 
**MessagesSentLast30d** | Pointer to **int32** | Rolling 30-day send count (not calendar month). | [optional] 
**BytesStored** | Pointer to **int64** | Lifetime sum of &#x60;messages.size_bytes&#x60; for this tenant. | [optional] 
**RateCapPerMin** | Pointer to **int32** | Per-tenant outbound rate cap (sends per minute). | [optional] 
**DailyMsgQuota** | Pointer to **int32** | Per-tenant daily send quota (UTC day boundary). | [optional] 
**WebhooksTotal** | Pointer to **int32** |  | [optional] 
**WebhooksPaused** | Pointer to **int32** | Webhook subscriptions auto-paused after 50 consecutive failures (LT2). | [optional] 
**GeneratedAt** | **time.Time** | When this snapshot was generated, RFC 3339 UTC. | 

## Methods

### NewV1UsageGet200Response

`func NewV1UsageGet200Response(mailboxesActive int32, generatedAt time.Time, ) *V1UsageGet200Response`

NewV1UsageGet200Response instantiates a new V1UsageGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UsageGet200ResponseWithDefaults

`func NewV1UsageGet200ResponseWithDefaults() *V1UsageGet200Response`

NewV1UsageGet200ResponseWithDefaults instantiates a new V1UsageGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxesActive

`func (o *V1UsageGet200Response) GetMailboxesActive() int32`

GetMailboxesActive returns the MailboxesActive field if non-nil, zero value otherwise.

### GetMailboxesActiveOk

`func (o *V1UsageGet200Response) GetMailboxesActiveOk() (*int32, bool)`

GetMailboxesActiveOk returns a tuple with the MailboxesActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxesActive

`func (o *V1UsageGet200Response) SetMailboxesActive(v int32)`

SetMailboxesActive sets MailboxesActive field to given value.


### GetMailboxesTotal

`func (o *V1UsageGet200Response) GetMailboxesTotal() int32`

GetMailboxesTotal returns the MailboxesTotal field if non-nil, zero value otherwise.

### GetMailboxesTotalOk

`func (o *V1UsageGet200Response) GetMailboxesTotalOk() (*int32, bool)`

GetMailboxesTotalOk returns a tuple with the MailboxesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxesTotal

`func (o *V1UsageGet200Response) SetMailboxesTotal(v int32)`

SetMailboxesTotal sets MailboxesTotal field to given value.

### HasMailboxesTotal

`func (o *V1UsageGet200Response) HasMailboxesTotal() bool`

HasMailboxesTotal returns a boolean if a field has been set.

### GetDomainsVerified

`func (o *V1UsageGet200Response) GetDomainsVerified() int32`

GetDomainsVerified returns the DomainsVerified field if non-nil, zero value otherwise.

### GetDomainsVerifiedOk

`func (o *V1UsageGet200Response) GetDomainsVerifiedOk() (*int32, bool)`

GetDomainsVerifiedOk returns a tuple with the DomainsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainsVerified

`func (o *V1UsageGet200Response) SetDomainsVerified(v int32)`

SetDomainsVerified sets DomainsVerified field to given value.

### HasDomainsVerified

`func (o *V1UsageGet200Response) HasDomainsVerified() bool`

HasDomainsVerified returns a boolean if a field has been set.

### GetDomainsTotal

`func (o *V1UsageGet200Response) GetDomainsTotal() int32`

GetDomainsTotal returns the DomainsTotal field if non-nil, zero value otherwise.

### GetDomainsTotalOk

`func (o *V1UsageGet200Response) GetDomainsTotalOk() (*int32, bool)`

GetDomainsTotalOk returns a tuple with the DomainsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainsTotal

`func (o *V1UsageGet200Response) SetDomainsTotal(v int32)`

SetDomainsTotal sets DomainsTotal field to given value.

### HasDomainsTotal

`func (o *V1UsageGet200Response) HasDomainsTotal() bool`

HasDomainsTotal returns a boolean if a field has been set.

### GetMessagesSentLast60s

`func (o *V1UsageGet200Response) GetMessagesSentLast60s() int32`

GetMessagesSentLast60s returns the MessagesSentLast60s field if non-nil, zero value otherwise.

### GetMessagesSentLast60sOk

`func (o *V1UsageGet200Response) GetMessagesSentLast60sOk() (*int32, bool)`

GetMessagesSentLast60sOk returns a tuple with the MessagesSentLast60s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSentLast60s

`func (o *V1UsageGet200Response) SetMessagesSentLast60s(v int32)`

SetMessagesSentLast60s sets MessagesSentLast60s field to given value.

### HasMessagesSentLast60s

`func (o *V1UsageGet200Response) HasMessagesSentLast60s() bool`

HasMessagesSentLast60s returns a boolean if a field has been set.

### GetMessagesSentTodayUtc

`func (o *V1UsageGet200Response) GetMessagesSentTodayUtc() int32`

GetMessagesSentTodayUtc returns the MessagesSentTodayUtc field if non-nil, zero value otherwise.

### GetMessagesSentTodayUtcOk

`func (o *V1UsageGet200Response) GetMessagesSentTodayUtcOk() (*int32, bool)`

GetMessagesSentTodayUtcOk returns a tuple with the MessagesSentTodayUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSentTodayUtc

`func (o *V1UsageGet200Response) SetMessagesSentTodayUtc(v int32)`

SetMessagesSentTodayUtc sets MessagesSentTodayUtc field to given value.

### HasMessagesSentTodayUtc

`func (o *V1UsageGet200Response) HasMessagesSentTodayUtc() bool`

HasMessagesSentTodayUtc returns a boolean if a field has been set.

### GetMessagesSentLast30d

`func (o *V1UsageGet200Response) GetMessagesSentLast30d() int32`

GetMessagesSentLast30d returns the MessagesSentLast30d field if non-nil, zero value otherwise.

### GetMessagesSentLast30dOk

`func (o *V1UsageGet200Response) GetMessagesSentLast30dOk() (*int32, bool)`

GetMessagesSentLast30dOk returns a tuple with the MessagesSentLast30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSentLast30d

`func (o *V1UsageGet200Response) SetMessagesSentLast30d(v int32)`

SetMessagesSentLast30d sets MessagesSentLast30d field to given value.

### HasMessagesSentLast30d

`func (o *V1UsageGet200Response) HasMessagesSentLast30d() bool`

HasMessagesSentLast30d returns a boolean if a field has been set.

### GetBytesStored

`func (o *V1UsageGet200Response) GetBytesStored() int64`

GetBytesStored returns the BytesStored field if non-nil, zero value otherwise.

### GetBytesStoredOk

`func (o *V1UsageGet200Response) GetBytesStoredOk() (*int64, bool)`

GetBytesStoredOk returns a tuple with the BytesStored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesStored

`func (o *V1UsageGet200Response) SetBytesStored(v int64)`

SetBytesStored sets BytesStored field to given value.

### HasBytesStored

`func (o *V1UsageGet200Response) HasBytesStored() bool`

HasBytesStored returns a boolean if a field has been set.

### GetRateCapPerMin

`func (o *V1UsageGet200Response) GetRateCapPerMin() int32`

GetRateCapPerMin returns the RateCapPerMin field if non-nil, zero value otherwise.

### GetRateCapPerMinOk

`func (o *V1UsageGet200Response) GetRateCapPerMinOk() (*int32, bool)`

GetRateCapPerMinOk returns a tuple with the RateCapPerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCapPerMin

`func (o *V1UsageGet200Response) SetRateCapPerMin(v int32)`

SetRateCapPerMin sets RateCapPerMin field to given value.

### HasRateCapPerMin

`func (o *V1UsageGet200Response) HasRateCapPerMin() bool`

HasRateCapPerMin returns a boolean if a field has been set.

### GetDailyMsgQuota

`func (o *V1UsageGet200Response) GetDailyMsgQuota() int32`

GetDailyMsgQuota returns the DailyMsgQuota field if non-nil, zero value otherwise.

### GetDailyMsgQuotaOk

`func (o *V1UsageGet200Response) GetDailyMsgQuotaOk() (*int32, bool)`

GetDailyMsgQuotaOk returns a tuple with the DailyMsgQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyMsgQuota

`func (o *V1UsageGet200Response) SetDailyMsgQuota(v int32)`

SetDailyMsgQuota sets DailyMsgQuota field to given value.

### HasDailyMsgQuota

`func (o *V1UsageGet200Response) HasDailyMsgQuota() bool`

HasDailyMsgQuota returns a boolean if a field has been set.

### GetWebhooksTotal

`func (o *V1UsageGet200Response) GetWebhooksTotal() int32`

GetWebhooksTotal returns the WebhooksTotal field if non-nil, zero value otherwise.

### GetWebhooksTotalOk

`func (o *V1UsageGet200Response) GetWebhooksTotalOk() (*int32, bool)`

GetWebhooksTotalOk returns a tuple with the WebhooksTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksTotal

`func (o *V1UsageGet200Response) SetWebhooksTotal(v int32)`

SetWebhooksTotal sets WebhooksTotal field to given value.

### HasWebhooksTotal

`func (o *V1UsageGet200Response) HasWebhooksTotal() bool`

HasWebhooksTotal returns a boolean if a field has been set.

### GetWebhooksPaused

`func (o *V1UsageGet200Response) GetWebhooksPaused() int32`

GetWebhooksPaused returns the WebhooksPaused field if non-nil, zero value otherwise.

### GetWebhooksPausedOk

`func (o *V1UsageGet200Response) GetWebhooksPausedOk() (*int32, bool)`

GetWebhooksPausedOk returns a tuple with the WebhooksPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksPaused

`func (o *V1UsageGet200Response) SetWebhooksPaused(v int32)`

SetWebhooksPaused sets WebhooksPaused field to given value.

### HasWebhooksPaused

`func (o *V1UsageGet200Response) HasWebhooksPaused() bool`

HasWebhooksPaused returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *V1UsageGet200Response) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *V1UsageGet200Response) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *V1UsageGet200Response) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


