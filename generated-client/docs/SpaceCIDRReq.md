# SpaceCIDRReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocks** | **[]interface{}** |  | 
**Size** | **int32** |  | 
**Desc** | Pointer to **NullableString** |  | [optional] 
**ReverseSearch** | Pointer to **NullableBool** |  | [optional] 
**SmallestCidr** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSpaceCIDRReq

`func NewSpaceCIDRReq(blocks []interface{}, size int32, ) *SpaceCIDRReq`

NewSpaceCIDRReq instantiates a new SpaceCIDRReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceCIDRReqWithDefaults

`func NewSpaceCIDRReqWithDefaults() *SpaceCIDRReq`

NewSpaceCIDRReqWithDefaults instantiates a new SpaceCIDRReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocks

`func (o *SpaceCIDRReq) GetBlocks() []interface{}`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *SpaceCIDRReq) GetBlocksOk() (*[]interface{}, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *SpaceCIDRReq) SetBlocks(v []interface{})`

SetBlocks sets Blocks field to given value.


### GetSize

`func (o *SpaceCIDRReq) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SpaceCIDRReq) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SpaceCIDRReq) SetSize(v int32)`

SetSize sets Size field to given value.


### GetDesc

`func (o *SpaceCIDRReq) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *SpaceCIDRReq) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *SpaceCIDRReq) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *SpaceCIDRReq) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### SetDescNil

`func (o *SpaceCIDRReq) SetDescNil(b bool)`

 SetDescNil sets the value for Desc to be an explicit nil

### UnsetDesc
`func (o *SpaceCIDRReq) UnsetDesc()`

UnsetDesc ensures that no value is present for Desc, not even an explicit nil
### GetReverseSearch

`func (o *SpaceCIDRReq) GetReverseSearch() bool`

GetReverseSearch returns the ReverseSearch field if non-nil, zero value otherwise.

### GetReverseSearchOk

`func (o *SpaceCIDRReq) GetReverseSearchOk() (*bool, bool)`

GetReverseSearchOk returns a tuple with the ReverseSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseSearch

`func (o *SpaceCIDRReq) SetReverseSearch(v bool)`

SetReverseSearch sets ReverseSearch field to given value.

### HasReverseSearch

`func (o *SpaceCIDRReq) HasReverseSearch() bool`

HasReverseSearch returns a boolean if a field has been set.

### SetReverseSearchNil

`func (o *SpaceCIDRReq) SetReverseSearchNil(b bool)`

 SetReverseSearchNil sets the value for ReverseSearch to be an explicit nil

### UnsetReverseSearch
`func (o *SpaceCIDRReq) UnsetReverseSearch()`

UnsetReverseSearch ensures that no value is present for ReverseSearch, not even an explicit nil
### GetSmallestCidr

`func (o *SpaceCIDRReq) GetSmallestCidr() bool`

GetSmallestCidr returns the SmallestCidr field if non-nil, zero value otherwise.

### GetSmallestCidrOk

`func (o *SpaceCIDRReq) GetSmallestCidrOk() (*bool, bool)`

GetSmallestCidrOk returns a tuple with the SmallestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallestCidr

`func (o *SpaceCIDRReq) SetSmallestCidr(v bool)`

SetSmallestCidr sets SmallestCidr field to given value.

### HasSmallestCidr

`func (o *SpaceCIDRReq) HasSmallestCidr() bool`

HasSmallestCidr returns a boolean if a field has been set.

### SetSmallestCidrNil

`func (o *SpaceCIDRReq) SetSmallestCidrNil(b bool)`

 SetSmallestCidrNil sets the value for SmallestCidr to be an explicit nil

### UnsetSmallestCidr
`func (o *SpaceCIDRReq) UnsetSmallestCidr()`

UnsetSmallestCidr ensures that no value is present for SmallestCidr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


