# ResponseGetUserApiUsersMeGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DarkMode** | **bool** |  | 
**ApiRefresh** | **int32** |  | 
**IsAdmin** | **bool** |  | 
**Views** | [**map[string]ViewSettings**](ViewSettings.md) |  | 

## Methods

### NewResponseGetUserApiUsersMeGet

`func NewResponseGetUserApiUsersMeGet(id string, darkMode bool, apiRefresh int32, isAdmin bool, views map[string]ViewSettings, ) *ResponseGetUserApiUsersMeGet`

NewResponseGetUserApiUsersMeGet instantiates a new ResponseGetUserApiUsersMeGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseGetUserApiUsersMeGetWithDefaults

`func NewResponseGetUserApiUsersMeGetWithDefaults() *ResponseGetUserApiUsersMeGet`

NewResponseGetUserApiUsersMeGetWithDefaults instantiates a new ResponseGetUserApiUsersMeGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseGetUserApiUsersMeGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseGetUserApiUsersMeGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseGetUserApiUsersMeGet) SetId(v string)`

SetId sets Id field to given value.


### GetDarkMode

`func (o *ResponseGetUserApiUsersMeGet) GetDarkMode() bool`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *ResponseGetUserApiUsersMeGet) GetDarkModeOk() (*bool, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *ResponseGetUserApiUsersMeGet) SetDarkMode(v bool)`

SetDarkMode sets DarkMode field to given value.


### GetApiRefresh

`func (o *ResponseGetUserApiUsersMeGet) GetApiRefresh() int32`

GetApiRefresh returns the ApiRefresh field if non-nil, zero value otherwise.

### GetApiRefreshOk

`func (o *ResponseGetUserApiUsersMeGet) GetApiRefreshOk() (*int32, bool)`

GetApiRefreshOk returns a tuple with the ApiRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRefresh

`func (o *ResponseGetUserApiUsersMeGet) SetApiRefresh(v int32)`

SetApiRefresh sets ApiRefresh field to given value.


### GetIsAdmin

`func (o *ResponseGetUserApiUsersMeGet) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *ResponseGetUserApiUsersMeGet) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *ResponseGetUserApiUsersMeGet) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.


### GetViews

`func (o *ResponseGetUserApiUsersMeGet) GetViews() map[string]ViewSettings`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *ResponseGetUserApiUsersMeGet) GetViewsOk() (*map[string]ViewSettings, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *ResponseGetUserApiUsersMeGet) SetViews(v map[string]ViewSettings)`

SetViews sets Views field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


