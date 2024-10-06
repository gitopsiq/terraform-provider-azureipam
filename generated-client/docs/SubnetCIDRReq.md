# SubnetCIDRReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VnetId** | **string** |  | 
**Size** | **int32** |  | 
**ReverseSearch** | Pointer to **NullableBool** |  | [optional] 
**SmallestCidr** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSubnetCIDRReq

`func NewSubnetCIDRReq(vnetId string, size int32, ) *SubnetCIDRReq`

NewSubnetCIDRReq instantiates a new SubnetCIDRReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetCIDRReqWithDefaults

`func NewSubnetCIDRReqWithDefaults() *SubnetCIDRReq`

NewSubnetCIDRReqWithDefaults instantiates a new SubnetCIDRReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVnetId

`func (o *SubnetCIDRReq) GetVnetId() string`

GetVnetId returns the VnetId field if non-nil, zero value otherwise.

### GetVnetIdOk

`func (o *SubnetCIDRReq) GetVnetIdOk() (*string, bool)`

GetVnetIdOk returns a tuple with the VnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetId

`func (o *SubnetCIDRReq) SetVnetId(v string)`

SetVnetId sets VnetId field to given value.


### GetSize

`func (o *SubnetCIDRReq) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SubnetCIDRReq) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SubnetCIDRReq) SetSize(v int32)`

SetSize sets Size field to given value.


### GetReverseSearch

`func (o *SubnetCIDRReq) GetReverseSearch() bool`

GetReverseSearch returns the ReverseSearch field if non-nil, zero value otherwise.

### GetReverseSearchOk

`func (o *SubnetCIDRReq) GetReverseSearchOk() (*bool, bool)`

GetReverseSearchOk returns a tuple with the ReverseSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseSearch

`func (o *SubnetCIDRReq) SetReverseSearch(v bool)`

SetReverseSearch sets ReverseSearch field to given value.

### HasReverseSearch

`func (o *SubnetCIDRReq) HasReverseSearch() bool`

HasReverseSearch returns a boolean if a field has been set.

### SetReverseSearchNil

`func (o *SubnetCIDRReq) SetReverseSearchNil(b bool)`

 SetReverseSearchNil sets the value for ReverseSearch to be an explicit nil

### UnsetReverseSearch
`func (o *SubnetCIDRReq) UnsetReverseSearch()`

UnsetReverseSearch ensures that no value is present for ReverseSearch, not even an explicit nil
### GetSmallestCidr

`func (o *SubnetCIDRReq) GetSmallestCidr() bool`

GetSmallestCidr returns the SmallestCidr field if non-nil, zero value otherwise.

### GetSmallestCidrOk

`func (o *SubnetCIDRReq) GetSmallestCidrOk() (*bool, bool)`

GetSmallestCidrOk returns a tuple with the SmallestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallestCidr

`func (o *SubnetCIDRReq) SetSmallestCidr(v bool)`

SetSmallestCidr sets SmallestCidr field to given value.

### HasSmallestCidr

`func (o *SubnetCIDRReq) HasSmallestCidr() bool`

HasSmallestCidr returns a boolean if a field has been set.

### SetSmallestCidrNil

`func (o *SubnetCIDRReq) SetSmallestCidrNil(b bool)`

 SetSmallestCidrNil sets the value for SmallestCidr to be an explicit nil

### UnsetSmallestCidr
`func (o *SubnetCIDRReq) UnsetSmallestCidr()`

UnsetSmallestCidr ensures that no value is present for SmallestCidr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


