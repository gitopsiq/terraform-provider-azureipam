# \ToolsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CidrCheckApiToolsCidrCheckPost**](ToolsAPI.md#CidrCheckApiToolsCidrCheckPost) | **Post** /api/tools/cidrCheck | Find Virtual Networks that Overlap a Given CIDR Range
[**NextAvailableSubnetApiToolsNextAvailableSubnetPost**](ToolsAPI.md#NextAvailableSubnetApiToolsNextAvailableSubnetPost) | **Post** /api/tools/nextAvailableSubnet | Get Next Available Subnet in a Virtual Network
[**NextAvailableVnetApiToolsNextAvailableVNetPost**](ToolsAPI.md#NextAvailableVnetApiToolsNextAvailableVNetPost) | **Post** /api/tools/nextAvailableVNet | Get Next Available Virtual Network from a List of Blocks



## CidrCheckApiToolsCidrCheckPost

> []CIDRCheckRes CidrCheckApiToolsCidrCheckPost(ctx).CIDRCheckReq(cIDRCheckReq).Authorization(authorization).Execute()

Find Virtual Networks that Overlap a Given CIDR Range



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cIDRCheckReq := *openapiclient.NewCIDRCheckReq("Cidr_example") // CIDRCheckReq | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.CidrCheckApiToolsCidrCheckPost(context.Background()).CIDRCheckReq(cIDRCheckReq).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.CidrCheckApiToolsCidrCheckPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CidrCheckApiToolsCidrCheckPost`: []CIDRCheckRes
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.CidrCheckApiToolsCidrCheckPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCidrCheckApiToolsCidrCheckPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cIDRCheckReq** | [**CIDRCheckReq**](CIDRCheckReq.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**[]CIDRCheckRes**](CIDRCheckRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NextAvailableSubnetApiToolsNextAvailableSubnetPost

> NewSubnetCIDR NextAvailableSubnetApiToolsNextAvailableSubnetPost(ctx).SubnetCIDRReq(subnetCIDRReq).Authorization(authorization).Execute()

Get Next Available Subnet in a Virtual Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	subnetCIDRReq := *openapiclient.NewSubnetCIDRReq("VnetId_example", int32(123)) // SubnetCIDRReq | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.NextAvailableSubnetApiToolsNextAvailableSubnetPost(context.Background()).SubnetCIDRReq(subnetCIDRReq).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.NextAvailableSubnetApiToolsNextAvailableSubnetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NextAvailableSubnetApiToolsNextAvailableSubnetPost`: NewSubnetCIDR
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.NextAvailableSubnetApiToolsNextAvailableSubnetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNextAvailableSubnetApiToolsNextAvailableSubnetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subnetCIDRReq** | [**SubnetCIDRReq**](SubnetCIDRReq.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**NewSubnetCIDR**](NewSubnetCIDR.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NextAvailableVnetApiToolsNextAvailableVNetPost

> NewVNetCIDR NextAvailableVnetApiToolsNextAvailableVNetPost(ctx).VNetCIDRReq(vNetCIDRReq).Authorization(authorization).Execute()

Get Next Available Virtual Network from a List of Blocks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	vNetCIDRReq := *openapiclient.NewVNetCIDRReq("Space_example", []string{"Blocks_example"}, int32(123)) // VNetCIDRReq | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.NextAvailableVnetApiToolsNextAvailableVNetPost(context.Background()).VNetCIDRReq(vNetCIDRReq).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.NextAvailableVnetApiToolsNextAvailableVNetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NextAvailableVnetApiToolsNextAvailableVNetPost`: NewVNetCIDR
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.NextAvailableVnetApiToolsNextAvailableVNetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNextAvailableVnetApiToolsNextAvailableVNetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vNetCIDRReq** | [**VNetCIDRReq**](VNetCIDRReq.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**NewVNetCIDR**](NewVNetCIDR.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

