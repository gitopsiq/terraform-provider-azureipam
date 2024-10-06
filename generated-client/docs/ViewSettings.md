# ViewSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | **map[string]map[string]interface{}** |  | 
**Order** | **[]string** |  | 
**Sort** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewViewSettings

`func NewViewSettings(values map[string]map[string]interface{}, order []string, ) *ViewSettings`

NewViewSettings instantiates a new ViewSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewSettingsWithDefaults

`func NewViewSettingsWithDefaults() *ViewSettings`

NewViewSettingsWithDefaults instantiates a new ViewSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *ViewSettings) GetValues() map[string]map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ViewSettings) GetValuesOk() (*map[string]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ViewSettings) SetValues(v map[string]map[string]interface{})`

SetValues sets Values field to given value.


### GetOrder

`func (o *ViewSettings) GetOrder() []string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ViewSettings) GetOrderOk() (*[]string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ViewSettings) SetOrder(v []string)`

SetOrder sets Order field to given value.


### GetSort

`func (o *ViewSettings) GetSort() map[string]interface{}`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ViewSettings) GetSortOk() (*map[string]interface{}, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ViewSettings) SetSort(v map[string]interface{})`

SetSort sets Sort field to given value.

### HasSort

`func (o *ViewSettings) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *ViewSettings) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *ViewSettings) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


