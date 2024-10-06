# SpaceBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Desc** | **string** |  | 
**Blocks** | [**[]BlockBasic**](BlockBasic.md) |  | 

## Methods

### NewSpaceBasic

`func NewSpaceBasic(name string, desc string, blocks []BlockBasic, ) *SpaceBasic`

NewSpaceBasic instantiates a new SpaceBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceBasicWithDefaults

`func NewSpaceBasicWithDefaults() *SpaceBasic`

NewSpaceBasicWithDefaults instantiates a new SpaceBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SpaceBasic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpaceBasic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpaceBasic) SetName(v string)`

SetName sets Name field to given value.


### GetDesc

`func (o *SpaceBasic) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *SpaceBasic) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *SpaceBasic) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetBlocks

`func (o *SpaceBasic) GetBlocks() []BlockBasic`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *SpaceBasic) GetBlocksOk() (*[]BlockBasic, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *SpaceBasic) SetBlocks(v []BlockBasic)`

SetBlocks sets Blocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


