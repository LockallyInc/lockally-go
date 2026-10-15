# GetAuditSummary200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminActionsToday** | Pointer to **int32** |  | [optional] 
**RecentLogins** | Pointer to [**[]AuditEvent**](AuditEvent.md) |  | [optional] 
**RecentExports** | Pointer to [**[]AuditEvent**](AuditEvent.md) |  | [optional] 
**DeletedMailboxes** | Pointer to [**[]AuditEvent**](AuditEvent.md) |  | [optional] 
**GeneratedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetAuditSummary200Response

`func NewGetAuditSummary200Response() *GetAuditSummary200Response`

NewGetAuditSummary200Response instantiates a new GetAuditSummary200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuditSummary200ResponseWithDefaults

`func NewGetAuditSummary200ResponseWithDefaults() *GetAuditSummary200Response`

NewGetAuditSummary200ResponseWithDefaults instantiates a new GetAuditSummary200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminActionsToday

`func (o *GetAuditSummary200Response) GetAdminActionsToday() int32`

GetAdminActionsToday returns the AdminActionsToday field if non-nil, zero value otherwise.

### GetAdminActionsTodayOk

`func (o *GetAuditSummary200Response) GetAdminActionsTodayOk() (*int32, bool)`

GetAdminActionsTodayOk returns a tuple with the AdminActionsToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminActionsToday

`func (o *GetAuditSummary200Response) SetAdminActionsToday(v int32)`

SetAdminActionsToday sets AdminActionsToday field to given value.

### HasAdminActionsToday

`func (o *GetAuditSummary200Response) HasAdminActionsToday() bool`

HasAdminActionsToday returns a boolean if a field has been set.

### GetRecentLogins

`func (o *GetAuditSummary200Response) GetRecentLogins() []AuditEvent`

GetRecentLogins returns the RecentLogins field if non-nil, zero value otherwise.

### GetRecentLoginsOk

`func (o *GetAuditSummary200Response) GetRecentLoginsOk() (*[]AuditEvent, bool)`

GetRecentLoginsOk returns a tuple with the RecentLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentLogins

`func (o *GetAuditSummary200Response) SetRecentLogins(v []AuditEvent)`

SetRecentLogins sets RecentLogins field to given value.

### HasRecentLogins

`func (o *GetAuditSummary200Response) HasRecentLogins() bool`

HasRecentLogins returns a boolean if a field has been set.

### GetRecentExports

`func (o *GetAuditSummary200Response) GetRecentExports() []AuditEvent`

GetRecentExports returns the RecentExports field if non-nil, zero value otherwise.

### GetRecentExportsOk

`func (o *GetAuditSummary200Response) GetRecentExportsOk() (*[]AuditEvent, bool)`

GetRecentExportsOk returns a tuple with the RecentExports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentExports

`func (o *GetAuditSummary200Response) SetRecentExports(v []AuditEvent)`

SetRecentExports sets RecentExports field to given value.

### HasRecentExports

`func (o *GetAuditSummary200Response) HasRecentExports() bool`

HasRecentExports returns a boolean if a field has been set.

### GetDeletedMailboxes

`func (o *GetAuditSummary200Response) GetDeletedMailboxes() []AuditEvent`

GetDeletedMailboxes returns the DeletedMailboxes field if non-nil, zero value otherwise.

### GetDeletedMailboxesOk

`func (o *GetAuditSummary200Response) GetDeletedMailboxesOk() (*[]AuditEvent, bool)`

GetDeletedMailboxesOk returns a tuple with the DeletedMailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedMailboxes

`func (o *GetAuditSummary200Response) SetDeletedMailboxes(v []AuditEvent)`

SetDeletedMailboxes sets DeletedMailboxes field to given value.

### HasDeletedMailboxes

`func (o *GetAuditSummary200Response) HasDeletedMailboxes() bool`

HasDeletedMailboxes returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *GetAuditSummary200Response) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *GetAuditSummary200Response) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *GetAuditSummary200Response) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *GetAuditSummary200Response) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


