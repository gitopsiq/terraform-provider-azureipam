# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DarkMode** | **bool** |  | 
**ApiRefresh** | **int32** |  | 
**IsAdmin** | **bool** |  | 

## Methods

### NewUser

`func NewUser(id string, darkMode bool, apiRefresh int32, isAdmin bool, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetDarkMode

`func (o *User) GetDarkMode() bool`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *User) GetDarkModeOk() (*bool, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *User) SetDarkMode(v bool)`

SetDarkMode sets DarkMode field to given value.


### GetApiRefresh

`func (o *User) GetApiRefresh() int32`

GetApiRefresh returns the ApiRefresh field if non-nil, zero value otherwise.

### GetApiRefreshOk

`func (o *User) GetApiRefreshOk() (*int32, bool)`

GetApiRefreshOk returns a tuple with the ApiRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRefresh

`func (o *User) SetApiRefresh(v int32)`

SetApiRefresh sets ApiRefresh field to given value.


### GetIsAdmin

`func (o *User) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *User) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *User) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


