# BlockBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Cidr** | **string** |  | 
**Vnets** | [**[]VNet**](VNet.md) |  | 
**Externals** | [**[]ExtNet**](ExtNet.md) |  | 
**Resv** | [**[]Reservation**](Reservation.md) |  | 

## Methods

### NewBlockBasic

`func NewBlockBasic(name string, cidr string, vnets []VNet, externals []ExtNet, resv []Reservation, ) *BlockBasic`

NewBlockBasic instantiates a new BlockBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockBasicWithDefaults

`func NewBlockBasicWithDefaults() *BlockBasic`

NewBlockBasicWithDefaults instantiates a new BlockBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlockBasic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlockBasic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlockBasic) SetName(v string)`

SetName sets Name field to given value.


### GetCidr

`func (o *BlockBasic) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *BlockBasic) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *BlockBasic) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetVnets

`func (o *BlockBasic) GetVnets() []VNet`

GetVnets returns the Vnets field if non-nil, zero value otherwise.

### GetVnetsOk

`func (o *BlockBasic) GetVnetsOk() (*[]VNet, bool)`

GetVnetsOk returns a tuple with the Vnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnets

`func (o *BlockBasic) SetVnets(v []VNet)`

SetVnets sets Vnets field to given value.


### GetExternals

`func (o *BlockBasic) GetExternals() []ExtNet`

GetExternals returns the Externals field if non-nil, zero value otherwise.

### GetExternalsOk

`func (o *BlockBasic) GetExternalsOk() (*[]ExtNet, bool)`

GetExternalsOk returns a tuple with the Externals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternals

`func (o *BlockBasic) SetExternals(v []ExtNet)`

SetExternals sets Externals field to given value.


### GetResv

`func (o *BlockBasic) GetResv() []Reservation`

GetResv returns the Resv field if non-nil, zero value otherwise.

### GetResvOk

`func (o *BlockBasic) GetResvOk() (*[]Reservation, bool)`

GetResvOk returns a tuple with the Resv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResv

`func (o *BlockBasic) SetResv(v []Reservation)`

SetResv sets Resv field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


