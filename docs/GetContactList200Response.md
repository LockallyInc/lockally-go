# GetContactList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**ContactList**](ContactList.md) |  | [optional] 
**Members** | Pointer to [**[]GetContactList200ResponseMembersInner**](GetContactList200ResponseMembersInner.md) |  | [optional] 

## Methods

### NewGetContactList200Response

`func NewGetContactList200Response() *GetContactList200Response`

NewGetContactList200Response instantiates a new GetContactList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContactList200ResponseWithDefaults

`func NewGetContactList200ResponseWithDefaults() *GetContactList200Response`

NewGetContactList200ResponseWithDefaults instantiates a new GetContactList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *GetContactList200Response) GetList() ContactList`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetContactList200Response) GetListOk() (*ContactList, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetContactList200Response) SetList(v ContactList)`

SetList sets List field to given value.

### HasList

`func (o *GetContactList200Response) HasList() bool`

HasList returns a boolean if a field has been set.

### GetMembers

`func (o *GetContactList200Response) GetMembers() []GetContactList200ResponseMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetContactList200Response) GetMembersOk() (*[]GetContactList200ResponseMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetContactList200Response) SetMembers(v []GetContactList200ResponseMembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetContactList200Response) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


