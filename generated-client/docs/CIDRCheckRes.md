# CIDRCheckRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Id** | **string** |  | 
**ResourceGroup** | **string** |  | 
**SubscriptionId** | **string** |  | 
**TenantId** | **string** |  | 
**Prefixes** | **[]string** |  | 
**Containers** | [**[]CIDRContainer**](CIDRContainer.md) |  | 

## Methods

### NewCIDRCheckRes

`func NewCIDRCheckRes(name string, id string, resourceGroup string, subscriptionId string, tenantId string, prefixes []string, containers []CIDRContainer, ) *CIDRCheckRes`

NewCIDRCheckRes instantiates a new CIDRCheckRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCIDRCheckResWithDefaults

`func NewCIDRCheckResWithDefaults() *CIDRCheckRes`

NewCIDRCheckResWithDefaults instantiates a new CIDRCheckRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CIDRCheckRes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CIDRCheckRes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CIDRCheckRes) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *CIDRCheckRes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CIDRCheckRes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CIDRCheckRes) SetId(v string)`

SetId sets Id field to given value.


### GetResourceGroup

`func (o *CIDRCheckRes) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *CIDRCheckRes) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *CIDRCheckRes) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetSubscriptionId

`func (o *CIDRCheckRes) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CIDRCheckRes) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CIDRCheckRes) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *CIDRCheckRes) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CIDRCheckRes) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CIDRCheckRes) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPrefixes

`func (o *CIDRCheckRes) GetPrefixes() []string`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *CIDRCheckRes) GetPrefixesOk() (*[]string, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *CIDRCheckRes) SetPrefixes(v []string)`

SetPrefixes sets Prefixes field to given value.


### GetContainers

`func (o *CIDRCheckRes) GetContainers() []CIDRContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *CIDRCheckRes) GetContainersOk() (*[]CIDRContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *CIDRCheckRes) SetContainers(v []CIDRContainer)`

SetContainers sets Containers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


