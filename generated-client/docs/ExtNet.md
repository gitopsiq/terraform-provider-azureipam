# ExtNet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Desc** | **string** |  | 
**Cidr** | **string** |  | 
**Subnets** | [**[]ExtSubnet**](ExtSubnet.md) |  | 

## Methods

### NewExtNet

`func NewExtNet(name string, desc string, cidr string, subnets []ExtSubnet, ) *ExtNet`

NewExtNet instantiates a new ExtNet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtNetWithDefaults

`func NewExtNetWithDefaults() *ExtNet`

NewExtNetWithDefaults instantiates a new ExtNet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtNet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtNet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtNet) SetName(v string)`

SetName sets Name field to given value.


### GetDesc

`func (o *ExtNet) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ExtNet) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ExtNet) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetCidr

`func (o *ExtNet) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ExtNet) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ExtNet) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetSubnets

`func (o *ExtNet) GetSubnets() []ExtSubnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *ExtNet) GetSubnetsOk() (*[]ExtSubnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *ExtNet) SetSubnets(v []ExtSubnet)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


