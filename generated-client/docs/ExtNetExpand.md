# ExtNetExpand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Desc** | **string** |  | 
**Space** | **string** |  | 
**Block** | **string** |  | 
**Cidr** | **string** |  | 
**Subnets** | [**[]ExtSubnet**](ExtSubnet.md) |  | 

## Methods

### NewExtNetExpand

`func NewExtNetExpand(name string, desc string, space string, block string, cidr string, subnets []ExtSubnet, ) *ExtNetExpand`

NewExtNetExpand instantiates a new ExtNetExpand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtNetExpandWithDefaults

`func NewExtNetExpandWithDefaults() *ExtNetExpand`

NewExtNetExpandWithDefaults instantiates a new ExtNetExpand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtNetExpand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtNetExpand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtNetExpand) SetName(v string)`

SetName sets Name field to given value.


### GetDesc

`func (o *ExtNetExpand) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ExtNetExpand) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ExtNetExpand) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetSpace

`func (o *ExtNetExpand) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *ExtNetExpand) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *ExtNetExpand) SetSpace(v string)`

SetSpace sets Space field to given value.


### GetBlock

`func (o *ExtNetExpand) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *ExtNetExpand) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *ExtNetExpand) SetBlock(v string)`

SetBlock sets Block field to given value.


### GetCidr

`func (o *ExtNetExpand) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ExtNetExpand) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ExtNetExpand) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetSubnets

`func (o *ExtNetExpand) GetSubnets() []ExtSubnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *ExtNetExpand) GetSubnetsOk() (*[]ExtSubnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *ExtNetExpand) SetSubnets(v []ExtSubnet)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


