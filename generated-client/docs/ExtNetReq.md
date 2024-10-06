# ExtNetReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Desc** | Pointer to **NullableString** |  | [optional] 
**Cidr** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewExtNetReq

`func NewExtNetReq(name string, ) *ExtNetReq`

NewExtNetReq instantiates a new ExtNetReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtNetReqWithDefaults

`func NewExtNetReqWithDefaults() *ExtNetReq`

NewExtNetReqWithDefaults instantiates a new ExtNetReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtNetReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtNetReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtNetReq) SetName(v string)`

SetName sets Name field to given value.


### GetDesc

`func (o *ExtNetReq) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ExtNetReq) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ExtNetReq) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *ExtNetReq) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### SetDescNil

`func (o *ExtNetReq) SetDescNil(b bool)`

 SetDescNil sets the value for Desc to be an explicit nil

### UnsetDesc
`func (o *ExtNetReq) UnsetDesc()`

UnsetDesc ensures that no value is present for Desc, not even an explicit nil
### GetCidr

`func (o *ExtNetReq) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ExtNetReq) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ExtNetReq) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *ExtNetReq) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### SetCidrNil

`func (o *ExtNetReq) SetCidrNil(b bool)`

 SetCidrNil sets the value for Cidr to be an explicit nil

### UnsetCidr
`func (o *ExtNetReq) UnsetCidr()`

UnsetCidr ensures that no value is present for Cidr, not even an explicit nil
### GetSize

`func (o *ExtNetReq) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ExtNetReq) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ExtNetReq) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ExtNetReq) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ExtNetReq) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ExtNetReq) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


