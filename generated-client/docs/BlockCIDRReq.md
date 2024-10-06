# BlockCIDRReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **NullableInt32** |  | [optional] 
**Cidr** | Pointer to **NullableString** |  | [optional] 
**Desc** | Pointer to **NullableString** |  | [optional] 
**ReverseSearch** | Pointer to **NullableBool** |  | [optional] 
**SmallestCidr** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewBlockCIDRReq

`func NewBlockCIDRReq() *BlockCIDRReq`

NewBlockCIDRReq instantiates a new BlockCIDRReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockCIDRReqWithDefaults

`func NewBlockCIDRReqWithDefaults() *BlockCIDRReq`

NewBlockCIDRReqWithDefaults instantiates a new BlockCIDRReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *BlockCIDRReq) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BlockCIDRReq) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BlockCIDRReq) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *BlockCIDRReq) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *BlockCIDRReq) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *BlockCIDRReq) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetCidr

`func (o *BlockCIDRReq) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *BlockCIDRReq) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *BlockCIDRReq) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *BlockCIDRReq) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### SetCidrNil

`func (o *BlockCIDRReq) SetCidrNil(b bool)`

 SetCidrNil sets the value for Cidr to be an explicit nil

### UnsetCidr
`func (o *BlockCIDRReq) UnsetCidr()`

UnsetCidr ensures that no value is present for Cidr, not even an explicit nil
### GetDesc

`func (o *BlockCIDRReq) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *BlockCIDRReq) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *BlockCIDRReq) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *BlockCIDRReq) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### SetDescNil

`func (o *BlockCIDRReq) SetDescNil(b bool)`

 SetDescNil sets the value for Desc to be an explicit nil

### UnsetDesc
`func (o *BlockCIDRReq) UnsetDesc()`

UnsetDesc ensures that no value is present for Desc, not even an explicit nil
### GetReverseSearch

`func (o *BlockCIDRReq) GetReverseSearch() bool`

GetReverseSearch returns the ReverseSearch field if non-nil, zero value otherwise.

### GetReverseSearchOk

`func (o *BlockCIDRReq) GetReverseSearchOk() (*bool, bool)`

GetReverseSearchOk returns a tuple with the ReverseSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseSearch

`func (o *BlockCIDRReq) SetReverseSearch(v bool)`

SetReverseSearch sets ReverseSearch field to given value.

### HasReverseSearch

`func (o *BlockCIDRReq) HasReverseSearch() bool`

HasReverseSearch returns a boolean if a field has been set.

### SetReverseSearchNil

`func (o *BlockCIDRReq) SetReverseSearchNil(b bool)`

 SetReverseSearchNil sets the value for ReverseSearch to be an explicit nil

### UnsetReverseSearch
`func (o *BlockCIDRReq) UnsetReverseSearch()`

UnsetReverseSearch ensures that no value is present for ReverseSearch, not even an explicit nil
### GetSmallestCidr

`func (o *BlockCIDRReq) GetSmallestCidr() bool`

GetSmallestCidr returns the SmallestCidr field if non-nil, zero value otherwise.

### GetSmallestCidrOk

`func (o *BlockCIDRReq) GetSmallestCidrOk() (*bool, bool)`

GetSmallestCidrOk returns a tuple with the SmallestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallestCidr

`func (o *BlockCIDRReq) SetSmallestCidr(v bool)`

SetSmallestCidr sets SmallestCidr field to given value.

### HasSmallestCidr

`func (o *BlockCIDRReq) HasSmallestCidr() bool`

HasSmallestCidr returns a boolean if a field has been set.

### SetSmallestCidrNil

`func (o *BlockCIDRReq) SetSmallestCidrNil(b bool)`

 SetSmallestCidrNil sets the value for SmallestCidr to be an explicit nil

### UnsetSmallestCidr
`func (o *BlockCIDRReq) UnsetSmallestCidr()`

UnsetSmallestCidr ensures that no value is present for SmallestCidr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


