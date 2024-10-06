# VNetCIDRReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Space** | **string** |  | 
**Blocks** | **[]string** |  | 
**Size** | **int32** |  | 
**ReverseSearch** | Pointer to **NullableBool** |  | [optional] 
**SmallestCidr** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewVNetCIDRReq

`func NewVNetCIDRReq(space string, blocks []string, size int32, ) *VNetCIDRReq`

NewVNetCIDRReq instantiates a new VNetCIDRReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVNetCIDRReqWithDefaults

`func NewVNetCIDRReqWithDefaults() *VNetCIDRReq`

NewVNetCIDRReqWithDefaults instantiates a new VNetCIDRReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpace

`func (o *VNetCIDRReq) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *VNetCIDRReq) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *VNetCIDRReq) SetSpace(v string)`

SetSpace sets Space field to given value.


### GetBlocks

`func (o *VNetCIDRReq) GetBlocks() []string`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *VNetCIDRReq) GetBlocksOk() (*[]string, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *VNetCIDRReq) SetBlocks(v []string)`

SetBlocks sets Blocks field to given value.


### GetSize

`func (o *VNetCIDRReq) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VNetCIDRReq) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VNetCIDRReq) SetSize(v int32)`

SetSize sets Size field to given value.


### GetReverseSearch

`func (o *VNetCIDRReq) GetReverseSearch() bool`

GetReverseSearch returns the ReverseSearch field if non-nil, zero value otherwise.

### GetReverseSearchOk

`func (o *VNetCIDRReq) GetReverseSearchOk() (*bool, bool)`

GetReverseSearchOk returns a tuple with the ReverseSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseSearch

`func (o *VNetCIDRReq) SetReverseSearch(v bool)`

SetReverseSearch sets ReverseSearch field to given value.

### HasReverseSearch

`func (o *VNetCIDRReq) HasReverseSearch() bool`

HasReverseSearch returns a boolean if a field has been set.

### SetReverseSearchNil

`func (o *VNetCIDRReq) SetReverseSearchNil(b bool)`

 SetReverseSearchNil sets the value for ReverseSearch to be an explicit nil

### UnsetReverseSearch
`func (o *VNetCIDRReq) UnsetReverseSearch()`

UnsetReverseSearch ensures that no value is present for ReverseSearch, not even an explicit nil
### GetSmallestCidr

`func (o *VNetCIDRReq) GetSmallestCidr() bool`

GetSmallestCidr returns the SmallestCidr field if non-nil, zero value otherwise.

### GetSmallestCidrOk

`func (o *VNetCIDRReq) GetSmallestCidrOk() (*bool, bool)`

GetSmallestCidrOk returns a tuple with the SmallestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallestCidr

`func (o *VNetCIDRReq) SetSmallestCidr(v bool)`

SetSmallestCidr sets SmallestCidr field to given value.

### HasSmallestCidr

`func (o *VNetCIDRReq) HasSmallestCidr() bool`

HasSmallestCidr returns a boolean if a field has been set.

### SetSmallestCidrNil

`func (o *VNetCIDRReq) SetSmallestCidrNil(b bool)`

 SetSmallestCidrNil sets the value for SmallestCidr to be an explicit nil

### UnsetSmallestCidr
`func (o *VNetCIDRReq) UnsetSmallestCidr()`

UnsetSmallestCidr ensures that no value is present for SmallestCidr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


