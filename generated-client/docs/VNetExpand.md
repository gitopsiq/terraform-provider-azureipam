# VNetExpand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Id** | **string** |  | 
**Prefixes** | **[]string** |  | 
**Subnets** | [**[]Subnet**](Subnet.md) |  | 
**ResourceGroup** | **string** |  | 
**SubscriptionId** | **string** |  | 
**TenantId** | **string** |  | 

## Methods

### NewVNetExpand

`func NewVNetExpand(name string, id string, prefixes []string, subnets []Subnet, resourceGroup string, subscriptionId string, tenantId string, ) *VNetExpand`

NewVNetExpand instantiates a new VNetExpand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVNetExpandWithDefaults

`func NewVNetExpandWithDefaults() *VNetExpand`

NewVNetExpandWithDefaults instantiates a new VNetExpand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VNetExpand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VNetExpand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VNetExpand) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *VNetExpand) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VNetExpand) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VNetExpand) SetId(v string)`

SetId sets Id field to given value.


### GetPrefixes

`func (o *VNetExpand) GetPrefixes() []string`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *VNetExpand) GetPrefixesOk() (*[]string, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *VNetExpand) SetPrefixes(v []string)`

SetPrefixes sets Prefixes field to given value.


### GetSubnets

`func (o *VNetExpand) GetSubnets() []Subnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *VNetExpand) GetSubnetsOk() (*[]Subnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *VNetExpand) SetSubnets(v []Subnet)`

SetSubnets sets Subnets field to given value.


### GetResourceGroup

`func (o *VNetExpand) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *VNetExpand) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *VNetExpand) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetSubscriptionId

`func (o *VNetExpand) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *VNetExpand) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *VNetExpand) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *VNetExpand) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *VNetExpand) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *VNetExpand) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


