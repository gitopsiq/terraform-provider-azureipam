# ExtSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Desc** | **string** |  | 
**Cidr** | **string** |  | 
**Endpoints** | [**[]ExtEndpoint**](ExtEndpoint.md) |  | 

## Methods

### NewExtSubnet

`func NewExtSubnet(name string, desc string, cidr string, endpoints []ExtEndpoint, ) *ExtSubnet`

NewExtSubnet instantiates a new ExtSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtSubnetWithDefaults

`func NewExtSubnetWithDefaults() *ExtSubnet`

NewExtSubnetWithDefaults instantiates a new ExtSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtSubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtSubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtSubnet) SetName(v string)`

SetName sets Name field to given value.


### GetDesc

`func (o *ExtSubnet) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ExtSubnet) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ExtSubnet) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetCidr

`func (o *ExtSubnet) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ExtSubnet) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ExtSubnet) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetEndpoints

`func (o *ExtSubnet) GetEndpoints() []ExtEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ExtSubnet) GetEndpointsOk() (*[]ExtEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ExtSubnet) SetEndpoints(v []ExtEndpoint)`

SetEndpoints sets Endpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


