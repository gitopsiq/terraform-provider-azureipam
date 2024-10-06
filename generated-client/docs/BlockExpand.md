# BlockExpand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Cidr** | **string** |  | 
**Vnets** | [**[]VNetExpand**](VNetExpand.md) |  | 
**Externals** | [**[]ExtNet**](ExtNet.md) |  | 
**Resv** | [**[]Reservation**](Reservation.md) |  | 

## Methods

### NewBlockExpand

`func NewBlockExpand(name string, cidr string, vnets []VNetExpand, externals []ExtNet, resv []Reservation, ) *BlockExpand`

NewBlockExpand instantiates a new BlockExpand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockExpandWithDefaults

`func NewBlockExpandWithDefaults() *BlockExpand`

NewBlockExpandWithDefaults instantiates a new BlockExpand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlockExpand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlockExpand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlockExpand) SetName(v string)`

SetName sets Name field to given value.


### GetCidr

`func (o *BlockExpand) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *BlockExpand) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *BlockExpand) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetVnets

`func (o *BlockExpand) GetVnets() []VNetExpand`

GetVnets returns the Vnets field if non-nil, zero value otherwise.

### GetVnetsOk

`func (o *BlockExpand) GetVnetsOk() (*[]VNetExpand, bool)`

GetVnetsOk returns a tuple with the Vnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnets

`func (o *BlockExpand) SetVnets(v []VNetExpand)`

SetVnets sets Vnets field to given value.


### GetExternals

`func (o *BlockExpand) GetExternals() []ExtNet`

GetExternals returns the Externals field if non-nil, zero value otherwise.

### GetExternalsOk

`func (o *BlockExpand) GetExternalsOk() (*[]ExtNet, bool)`

GetExternalsOk returns a tuple with the Externals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternals

`func (o *BlockExpand) SetExternals(v []ExtNet)`

SetExternals sets Externals field to given value.


### GetResv

`func (o *BlockExpand) GetResv() []Reservation`

GetResv returns the Resv field if non-nil, zero value otherwise.

### GetResvOk

`func (o *BlockExpand) GetResvOk() (*[]Reservation, bool)`

GetResvOk returns a tuple with the Resv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResv

`func (o *BlockExpand) SetResv(v []Reservation)`

SetResv sets Resv field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


