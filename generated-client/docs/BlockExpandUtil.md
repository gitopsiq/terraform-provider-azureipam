# BlockExpandUtil

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Cidr** | **string** |  | 
**Vnets** | [**[]VNetExpandUtil**](VNetExpandUtil.md) |  | 
**Externals** | [**[]ExtNet**](ExtNet.md) |  | 
**Resv** | [**[]Reservation**](Reservation.md) |  | 
**Size** | **int32** |  | 
**Used** | **int32** |  | 

## Methods

### NewBlockExpandUtil

`func NewBlockExpandUtil(name string, cidr string, vnets []VNetExpandUtil, externals []ExtNet, resv []Reservation, size int32, used int32, ) *BlockExpandUtil`

NewBlockExpandUtil instantiates a new BlockExpandUtil object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockExpandUtilWithDefaults

`func NewBlockExpandUtilWithDefaults() *BlockExpandUtil`

NewBlockExpandUtilWithDefaults instantiates a new BlockExpandUtil object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlockExpandUtil) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlockExpandUtil) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlockExpandUtil) SetName(v string)`

SetName sets Name field to given value.


### GetCidr

`func (o *BlockExpandUtil) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *BlockExpandUtil) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *BlockExpandUtil) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetVnets

`func (o *BlockExpandUtil) GetVnets() []VNetExpandUtil`

GetVnets returns the Vnets field if non-nil, zero value otherwise.

### GetVnetsOk

`func (o *BlockExpandUtil) GetVnetsOk() (*[]VNetExpandUtil, bool)`

GetVnetsOk returns a tuple with the Vnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnets

`func (o *BlockExpandUtil) SetVnets(v []VNetExpandUtil)`

SetVnets sets Vnets field to given value.


### GetExternals

`func (o *BlockExpandUtil) GetExternals() []ExtNet`

GetExternals returns the Externals field if non-nil, zero value otherwise.

### GetExternalsOk

`func (o *BlockExpandUtil) GetExternalsOk() (*[]ExtNet, bool)`

GetExternalsOk returns a tuple with the Externals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternals

`func (o *BlockExpandUtil) SetExternals(v []ExtNet)`

SetExternals sets Externals field to given value.


### GetResv

`func (o *BlockExpandUtil) GetResv() []Reservation`

GetResv returns the Resv field if non-nil, zero value otherwise.

### GetResvOk

`func (o *BlockExpandUtil) GetResvOk() (*[]Reservation, bool)`

GetResvOk returns a tuple with the Resv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResv

`func (o *BlockExpandUtil) SetResv(v []Reservation)`

SetResv sets Resv field to given value.


### GetSize

`func (o *BlockExpandUtil) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BlockExpandUtil) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BlockExpandUtil) SetSize(v int32)`

SetSize sets Size field to given value.


### GetUsed

`func (o *BlockExpandUtil) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *BlockExpandUtil) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *BlockExpandUtil) SetUsed(v int32)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


