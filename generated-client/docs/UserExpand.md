# UserExpand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DarkMode** | **bool** |  | 
**ApiRefresh** | **int32** |  | 
**IsAdmin** | **bool** |  | 
**Views** | [**map[string]ViewSettings**](ViewSettings.md) |  | 

## Methods

### NewUserExpand

`func NewUserExpand(id string, darkMode bool, apiRefresh int32, isAdmin bool, views map[string]ViewSettings, ) *UserExpand`

NewUserExpand instantiates a new UserExpand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserExpandWithDefaults

`func NewUserExpandWithDefaults() *UserExpand`

NewUserExpandWithDefaults instantiates a new UserExpand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserExpand) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserExpand) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserExpand) SetId(v string)`

SetId sets Id field to given value.


### GetDarkMode

`func (o *UserExpand) GetDarkMode() bool`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *UserExpand) GetDarkModeOk() (*bool, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *UserExpand) SetDarkMode(v bool)`

SetDarkMode sets DarkMode field to given value.


### GetApiRefresh

`func (o *UserExpand) GetApiRefresh() int32`

GetApiRefresh returns the ApiRefresh field if non-nil, zero value otherwise.

### GetApiRefreshOk

`func (o *UserExpand) GetApiRefreshOk() (*int32, bool)`

GetApiRefreshOk returns a tuple with the ApiRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRefresh

`func (o *UserExpand) SetApiRefresh(v int32)`

SetApiRefresh sets ApiRefresh field to given value.


### GetIsAdmin

`func (o *UserExpand) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UserExpand) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *UserExpand) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.


### GetViews

`func (o *UserExpand) GetViews() map[string]ViewSettings`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *UserExpand) GetViewsOk() (*map[string]ViewSettings, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *UserExpand) SetViews(v map[string]ViewSettings)`

SetViews sets Views field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


