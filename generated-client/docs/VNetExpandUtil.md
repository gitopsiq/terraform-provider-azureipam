# VNetExpandUtil

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Id** | **string** |  | 
**Prefixes** | **[]string** |  | 
**Subnets** | [**[]SubnetUtil**](SubnetUtil.md) |  | 
**ResourceGroup** | **string** |  | 
**SubscriptionId** | **string** |  | 
**TenantId** | **string** |  | 
**Size** | **int32** |  | 
**Used** | **int32** |  | 

## Methods

### NewVNetExpandUtil

`func NewVNetExpandUtil(name string, id string, prefixes []string, subnets []SubnetUtil, resourceGroup string, subscriptionId string, tenantId string, size int32, used int32, ) *VNetExpandUtil`

NewVNetExpandUtil instantiates a new VNetExpandUtil object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVNetExpandUtilWithDefaults

`func NewVNetExpandUtilWithDefaults() *VNetExpandUtil`

NewVNetExpandUtilWithDefaults instantiates a new VNetExpandUtil object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VNetExpandUtil) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VNetExpandUtil) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VNetExpandUtil) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *VNetExpandUtil) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VNetExpandUtil) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VNetExpandUtil) SetId(v string)`

SetId sets Id field to given value.


### GetPrefixes

`func (o *VNetExpandUtil) GetPrefixes() []string`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *VNetExpandUtil) GetPrefixesOk() (*[]string, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *VNetExpandUtil) SetPrefixes(v []string)`

SetPrefixes sets Prefixes field to given value.


### GetSubnets

`func (o *VNetExpandUtil) GetSubnets() []SubnetUtil`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *VNetExpandUtil) GetSubnetsOk() (*[]SubnetUtil, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *VNetExpandUtil) SetSubnets(v []SubnetUtil)`

SetSubnets sets Subnets field to given value.


### GetResourceGroup

`func (o *VNetExpandUtil) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *VNetExpandUtil) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *VNetExpandUtil) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetSubscriptionId

`func (o *VNetExpandUtil) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *VNetExpandUtil) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *VNetExpandUtil) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *VNetExpandUtil) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *VNetExpandUtil) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *VNetExpandUtil) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSize

`func (o *VNetExpandUtil) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VNetExpandUtil) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VNetExpandUtil) SetSize(v int32)`

SetSize sets Size field to given value.


### GetUsed

`func (o *VNetExpandUtil) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *VNetExpandUtil) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *VNetExpandUtil) SetUsed(v int32)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


