# SpaceUtil

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Desc** | **string** |  | 
**Blocks** | [**[]BlockUtil**](BlockUtil.md) |  | 
**Size** | **int32** |  | 
**Used** | **int32** |  | 

## Methods

### NewSpaceUtil

`func NewSpaceUtil(name string, desc string, blocks []BlockUtil, size int32, used int32, ) *SpaceUtil`

NewSpaceUtil instantiates a new SpaceUtil object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceUtilWithDefaults

`func NewSpaceUtilWithDefaults() *SpaceUtil`

NewSpaceUtilWithDefaults instantiates a new SpaceUtil object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SpaceUtil) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpaceUtil) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpaceUtil) SetName(v string)`

SetName sets Name field to given value.


### GetDesc

`func (o *SpaceUtil) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *SpaceUtil) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *SpaceUtil) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetBlocks

`func (o *SpaceUtil) GetBlocks() []BlockUtil`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *SpaceUtil) GetBlocksOk() (*[]BlockUtil, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *SpaceUtil) SetBlocks(v []BlockUtil)`

SetBlocks sets Blocks field to given value.


### GetSize

`func (o *SpaceUtil) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SpaceUtil) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SpaceUtil) SetSize(v int32)`

SetSize sets Size field to given value.


### GetUsed

`func (o *SpaceUtil) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *SpaceUtil) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *SpaceUtil) SetUsed(v int32)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


