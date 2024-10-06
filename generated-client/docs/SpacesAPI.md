# \SpacesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvailableBlockNetsApiSpacesSpaceBlocksBlockAvailableGet**](SpacesAPI.md#AvailableBlockNetsApiSpacesSpaceBlocksBlockAvailableGet) | **Get** /api/spaces/{space}/blocks/{block}/available | List Available Block Networks
[**CreateBlockApiSpacesSpaceBlocksPost**](SpacesAPI.md#CreateBlockApiSpacesSpaceBlocksPost) | **Post** /api/spaces/{space}/blocks | Create a new Block
[**CreateBlockNetApiSpacesSpaceBlocksBlockNetworksPost**](SpacesAPI.md#CreateBlockNetApiSpacesSpaceBlocksBlockNetworksPost) | **Post** /api/spaces/{space}/blocks/{block}/networks | Add Block Network
[**CreateBlockReservationApiSpacesSpaceBlocksBlockReservationsPost**](SpacesAPI.md#CreateBlockReservationApiSpacesSpaceBlocksBlockReservationsPost) | **Post** /api/spaces/{space}/blocks/{block}/reservations | Create CIDR Reservation
[**CreateExternalNetworkApiSpacesSpaceBlocksBlockExternalsPost**](SpacesAPI.md#CreateExternalNetworkApiSpacesSpaceBlocksBlockExternalsPost) | **Post** /api/spaces/{space}/blocks/{block}/externals | Create External Network
[**CreateExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsPost**](SpacesAPI.md#CreateExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsPost) | **Post** /api/spaces/{space}/blocks/{block}/externals/{external}/subnets | Create External Network Subnet
[**CreateExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPost**](SpacesAPI.md#CreateExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPost) | **Post** /api/spaces/{space}/blocks/{block}/externals/{external}/subnets/{subnet}/endpoints | Add External Network Subnet Endpoint
[**CreateMultiBlockReservationApiSpacesSpaceReservationsPost**](SpacesAPI.md#CreateMultiBlockReservationApiSpacesSpaceReservationsPost) | **Post** /api/spaces/{space}/reservations | Create CIDR Reservation from List of Blocks
[**CreateSpaceApiSpacesPost**](SpacesAPI.md#CreateSpaceApiSpacesPost) | **Post** /api/spaces | Create New Space
[**DeleteBlockApiSpacesSpaceBlocksBlockDelete**](SpacesAPI.md#DeleteBlockApiSpacesSpaceBlocksBlockDelete) | **Delete** /api/spaces/{space}/blocks/{block} | Delete a Block
[**DeleteBlockNetsApiSpacesSpaceBlocksBlockNetworksDelete**](SpacesAPI.md#DeleteBlockNetsApiSpacesSpaceBlocksBlockNetworksDelete) | **Delete** /api/spaces/{space}/blocks/{block}/networks | Remove Block Networks
[**DeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsDelete**](SpacesAPI.md#DeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsDelete) | **Delete** /api/spaces/{space}/blocks/{block}/reservations | Delete CIDR Reservations
[**DeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationDelete**](SpacesAPI.md#DeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationDelete) | **Delete** /api/spaces/{space}/blocks/{block}/reservations/{reservation} | Delete CIDR Reservation
[**DeleteExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalDelete**](SpacesAPI.md#DeleteExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalDelete) | **Delete** /api/spaces/{space}/blocks/{block}/externals/{external} | Remove External Network
[**DeleteExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetDelete**](SpacesAPI.md#DeleteExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetDelete) | **Delete** /api/spaces/{space}/blocks/{block}/externals/{external}/subnets/{subnet} | Remove External Network Subnet
[**DeleteExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointDelete**](SpacesAPI.md#DeleteExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointDelete) | **Delete** /api/spaces/{space}/blocks/{block}/externals/{external}/subnets/{subnet}/endpoints/{endpoint} | Remove External Network Subnet Endpoint
[**DeleteExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsDelete**](SpacesAPI.md#DeleteExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsDelete) | **Delete** /api/spaces/{space}/blocks/{block}/externals/{external}/subnets/{subnet}/endpoints | Remove External Network Subnet Endpoints
[**DeleteSpaceApiSpacesSpaceDelete**](SpacesAPI.md#DeleteSpaceApiSpacesSpaceDelete) | **Delete** /api/spaces/{space} | Delete a Space
[**GetBlockApiSpacesSpaceBlocksBlockGet**](SpacesAPI.md#GetBlockApiSpacesSpaceBlocksBlockGet) | **Get** /api/spaces/{space}/blocks/{block} | Get Block Details
[**GetBlockNetsApiSpacesSpaceBlocksBlockNetworksGet**](SpacesAPI.md#GetBlockNetsApiSpacesSpaceBlocksBlockNetworksGet) | **Get** /api/spaces/{space}/blocks/{block}/networks | List Block Networks
[**GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsGet**](SpacesAPI.md#GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsGet) | **Get** /api/spaces/{space}/blocks/{block}/reservations | Get Block Reservations
[**GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationGet**](SpacesAPI.md#GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationGet) | **Get** /api/spaces/{space}/blocks/{block}/reservations/{reservation} | Get Block Reservation
[**GetBlocksApiSpacesSpaceBlocksGet**](SpacesAPI.md#GetBlocksApiSpacesSpaceBlocksGet) | **Get** /api/spaces/{space}/blocks | Get all Blocks within a Space
[**GetExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalGet**](SpacesAPI.md#GetExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalGet) | **Get** /api/spaces/{space}/blocks/{block}/externals/{external} | Get External Network
[**GetExternalNetworksApiSpacesSpaceBlocksBlockExternalsGet**](SpacesAPI.md#GetExternalNetworksApiSpacesSpaceBlocksBlockExternalsGet) | **Get** /api/spaces/{space}/blocks/{block}/externals | List External Networks
[**GetExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetGet**](SpacesAPI.md#GetExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetGet) | **Get** /api/spaces/{space}/blocks/{block}/externals/{external}/subnets/{subnet} | Get External Network Subnet
[**GetExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointGet**](SpacesAPI.md#GetExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointGet) | **Get** /api/spaces/{space}/blocks/{block}/externals/{external}/subnets/{subnet}/endpoints/{endpoint} | Get External Network Subnet Endpoint
[**GetExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsGet**](SpacesAPI.md#GetExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsGet) | **Get** /api/spaces/{space}/blocks/{block}/externals/{external}/subnets/{subnet}/endpoints | List External Network Subnet Endpoints
[**GetExternalSubnetsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsGet**](SpacesAPI.md#GetExternalSubnetsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsGet) | **Get** /api/spaces/{space}/blocks/{block}/externals/{external}/subnets | List External Network Subnets
[**GetMultiBlockReservationsApiSpacesSpaceReservationsGet**](SpacesAPI.md#GetMultiBlockReservationsApiSpacesSpaceReservationsGet) | **Get** /api/spaces/{space}/reservations | Get Reservations for all Blocks within a Space
[**GetSpaceApiSpacesSpaceGet**](SpacesAPI.md#GetSpaceApiSpacesSpaceGet) | **Get** /api/spaces/{space} | Get Space Details
[**GetSpacesApiSpacesGet**](SpacesAPI.md#GetSpacesApiSpacesGet) | **Get** /api/spaces | Get All Spaces
[**UpdateBlockApiSpacesSpaceBlocksBlockPatch**](SpacesAPI.md#UpdateBlockApiSpacesSpaceBlocksBlockPatch) | **Patch** /api/spaces/{space}/blocks/{block} | Update Block Details
[**UpdateBlockVnetsApiSpacesSpaceBlocksBlockNetworksPut**](SpacesAPI.md#UpdateBlockVnetsApiSpacesSpaceBlocksBlockNetworksPut) | **Put** /api/spaces/{space}/blocks/{block}/networks | Replace Block Networks
[**UpdateExtEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointPatch**](SpacesAPI.md#UpdateExtEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointPatch) | **Patch** /api/spaces/{space}/blocks/{block}/externals/{external}/subnets/{subnet}/endpoints/{endpoint} | Update External Endpoint Details
[**UpdateExtNetworkApiSpacesSpaceBlocksBlockExternalsExternalPatch**](SpacesAPI.md#UpdateExtNetworkApiSpacesSpaceBlocksBlockExternalsExternalPatch) | **Patch** /api/spaces/{space}/blocks/{block}/externals/{external} | Update External Network Details
[**UpdateExtSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetPatch**](SpacesAPI.md#UpdateExtSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetPatch) | **Patch** /api/spaces/{space}/blocks/{block}/externals/{external}/subnets/{subnet} | Update External Subnet Details
[**UpdateExternalSubnetEnpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPut**](SpacesAPI.md#UpdateExternalSubnetEnpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPut) | **Put** /api/spaces/{space}/blocks/{block}/externals/{external}/subnets/{subnet}/endpoints | Replace External Network Subnet Endpoints
[**UpdateSpaceApiSpacesSpacePatch**](SpacesAPI.md#UpdateSpaceApiSpacesSpacePatch) | **Patch** /api/spaces/{space} | Update Space Details



## AvailableBlockNetsApiSpacesSpaceBlocksBlockAvailableGet

> ResponseAvailableBlockNetsApiSpacesSpaceBlocksBlockAvailableGet AvailableBlockNetsApiSpacesSpaceBlocksBlockAvailableGet(ctx, space, block).Expand(expand).Authorization(authorization).Execute()

List Available Block Networks



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	expand := true // bool | Expand network references to full network objects (optional) (default to false)
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.AvailableBlockNetsApiSpacesSpaceBlocksBlockAvailableGet(context.Background(), space, block).Expand(expand).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.AvailableBlockNetsApiSpacesSpaceBlocksBlockAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableBlockNetsApiSpacesSpaceBlocksBlockAvailableGet`: ResponseAvailableBlockNetsApiSpacesSpaceBlocksBlockAvailableGet
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.AvailableBlockNetsApiSpacesSpaceBlocksBlockAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableBlockNetsApiSpacesSpaceBlocksBlockAvailableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **bool** | Expand network references to full network objects | [default to false]
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ResponseAvailableBlockNetsApiSpacesSpaceBlocksBlockAvailableGet**](ResponseAvailableBlockNetsApiSpacesSpaceBlocksBlockAvailableGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBlockApiSpacesSpaceBlocksPost

> Block CreateBlockApiSpacesSpaceBlocksPost(ctx, space).BlockReq(blockReq).Authorization(authorization).Execute()

Create a new Block



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
	space := "space_example" // string | Name of the target Space
	blockReq := *openapiclient.NewBlockReq("Name_example", "Cidr_example") // BlockReq | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.CreateBlockApiSpacesSpaceBlocksPost(context.Background(), space).BlockReq(blockReq).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.CreateBlockApiSpacesSpaceBlocksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlockApiSpacesSpaceBlocksPost`: Block
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.CreateBlockApiSpacesSpaceBlocksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlockApiSpacesSpaceBlocksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockReq** | [**BlockReq**](BlockReq.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBlockNetApiSpacesSpaceBlocksBlockNetworksPost

> BlockBasic CreateBlockNetApiSpacesSpaceBlocksBlockNetworksPost(ctx, space, block).VNet(vNet).Authorization(authorization).Execute()

Add Block Network



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	vNet := *openapiclient.NewVNet("Id_example") // VNet | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.CreateBlockNetApiSpacesSpaceBlocksBlockNetworksPost(context.Background(), space, block).VNet(vNet).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.CreateBlockNetApiSpacesSpaceBlocksBlockNetworksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlockNetApiSpacesSpaceBlocksBlockNetworksPost`: BlockBasic
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.CreateBlockNetApiSpacesSpaceBlocksBlockNetworksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlockNetApiSpacesSpaceBlocksBlockNetworksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vNet** | [**VNet**](VNet.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**BlockBasic**](BlockBasic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBlockReservationApiSpacesSpaceBlocksBlockReservationsPost

> ReservationExpand CreateBlockReservationApiSpacesSpaceBlocksBlockReservationsPost(ctx, space, block).BlockCIDRReq(blockCIDRReq).Authorization(authorization).Execute()

Create CIDR Reservation



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	blockCIDRReq := *openapiclient.NewBlockCIDRReq() // BlockCIDRReq | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.CreateBlockReservationApiSpacesSpaceBlocksBlockReservationsPost(context.Background(), space, block).BlockCIDRReq(blockCIDRReq).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.CreateBlockReservationApiSpacesSpaceBlocksBlockReservationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlockReservationApiSpacesSpaceBlocksBlockReservationsPost`: ReservationExpand
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.CreateBlockReservationApiSpacesSpaceBlocksBlockReservationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlockReservationApiSpacesSpaceBlocksBlockReservationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **blockCIDRReq** | [**BlockCIDRReq**](BlockCIDRReq.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ReservationExpand**](ReservationExpand.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExternalNetworkApiSpacesSpaceBlocksBlockExternalsPost

> ExtNetExpand CreateExternalNetworkApiSpacesSpaceBlocksBlockExternalsPost(ctx, space, block).ExtNetReq(extNetReq).Authorization(authorization).Execute()

Create External Network



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	extNetReq := *openapiclient.NewExtNetReq("Name_example") // ExtNetReq | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.CreateExternalNetworkApiSpacesSpaceBlocksBlockExternalsPost(context.Background(), space, block).ExtNetReq(extNetReq).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.CreateExternalNetworkApiSpacesSpaceBlocksBlockExternalsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalNetworkApiSpacesSpaceBlocksBlockExternalsPost`: ExtNetExpand
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.CreateExternalNetworkApiSpacesSpaceBlocksBlockExternalsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalNetworkApiSpacesSpaceBlocksBlockExternalsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **extNetReq** | [**ExtNetReq**](ExtNetReq.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ExtNetExpand**](ExtNetExpand.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsPost

> ExtSubnetExpand CreateExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsPost(ctx, space, block, external).ExtSubnetReq(extSubnetReq).Authorization(authorization).Execute()

Create External Network Subnet



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target External Network
	extSubnetReq := *openapiclient.NewExtSubnetReq("Name_example") // ExtSubnetReq | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.CreateExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsPost(context.Background(), space, block, external).ExtSubnetReq(extSubnetReq).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.CreateExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsPost`: ExtSubnetExpand
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.CreateExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target External Network | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **extSubnetReq** | [**ExtSubnetReq**](ExtSubnetReq.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ExtSubnetExpand**](ExtSubnetExpand.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPost

> ExtEndpoint CreateExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPost(ctx, space, block, external, subnet).ExtEndpointReq(extEndpointReq).Authorization(authorization).Execute()

Add External Network Subnet Endpoint



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target External Network
	subnet := "subnet_example" // string | Name of the target External Network Subnet
	extEndpointReq := *openapiclient.NewExtEndpointReq("Name_example", "Desc_example", "Ip_example") // ExtEndpointReq | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.CreateExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPost(context.Background(), space, block, external, subnet).ExtEndpointReq(extEndpointReq).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.CreateExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPost`: ExtEndpoint
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.CreateExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target External Network | 
**subnet** | **string** | Name of the target External Network Subnet | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **extEndpointReq** | [**ExtEndpointReq**](ExtEndpointReq.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ExtEndpoint**](ExtEndpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMultiBlockReservationApiSpacesSpaceReservationsPost

> ReservationExpand CreateMultiBlockReservationApiSpacesSpaceReservationsPost(ctx, space).SpaceCIDRReq(spaceCIDRReq).Authorization(authorization).Execute()

Create CIDR Reservation from List of Blocks



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
	space := "space_example" // string | Name of the target Space
	spaceCIDRReq := *openapiclient.NewSpaceCIDRReq([]interface{}{nil}, int32(123)) // SpaceCIDRReq | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.CreateMultiBlockReservationApiSpacesSpaceReservationsPost(context.Background(), space).SpaceCIDRReq(spaceCIDRReq).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.CreateMultiBlockReservationApiSpacesSpaceReservationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMultiBlockReservationApiSpacesSpaceReservationsPost`: ReservationExpand
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.CreateMultiBlockReservationApiSpacesSpaceReservationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMultiBlockReservationApiSpacesSpaceReservationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spaceCIDRReq** | [**SpaceCIDRReq**](SpaceCIDRReq.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ReservationExpand**](ReservationExpand.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpaceApiSpacesPost

> Space CreateSpaceApiSpacesPost(ctx).SpaceReq(spaceReq).Authorization(authorization).Execute()

Create New Space



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
	spaceReq := *openapiclient.NewSpaceReq("Name_example", "Desc_example") // SpaceReq | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.CreateSpaceApiSpacesPost(context.Background()).SpaceReq(spaceReq).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.CreateSpaceApiSpacesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpaceApiSpacesPost`: Space
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.CreateSpaceApiSpacesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpaceApiSpacesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spaceReq** | [**SpaceReq**](SpaceReq.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**Space**](Space.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlockApiSpacesSpaceBlocksBlockDelete

> interface{} DeleteBlockApiSpacesSpaceBlocksBlockDelete(ctx, space, block).Force(force).Authorization(authorization).Execute()

Delete a Block



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	force := true // bool | Forcefully delete a Block with existing networks and/or reservations (optional)
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.DeleteBlockApiSpacesSpaceBlocksBlockDelete(context.Background(), space, block).Force(force).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.DeleteBlockApiSpacesSpaceBlocksBlockDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlockApiSpacesSpaceBlocksBlockDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.DeleteBlockApiSpacesSpaceBlocksBlockDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockApiSpacesSpaceBlocksBlockDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | Forcefully delete a Block with existing networks and/or reservations | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlockNetsApiSpacesSpaceBlocksBlockNetworksDelete

> interface{} DeleteBlockNetsApiSpacesSpaceBlocksBlockNetworksDelete(ctx, space, block).RequestBody(requestBody).Authorization(authorization).Execute()

Remove Block Networks



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	requestBody := []string{"Property_example"} // []string | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.DeleteBlockNetsApiSpacesSpaceBlocksBlockNetworksDelete(context.Background(), space, block).RequestBody(requestBody).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.DeleteBlockNetsApiSpacesSpaceBlocksBlockNetworksDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlockNetsApiSpacesSpaceBlocksBlockNetworksDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.DeleteBlockNetsApiSpacesSpaceBlocksBlockNetworksDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockNetsApiSpacesSpaceBlocksBlockNetworksDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **[]string** |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsDelete

> DeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsDelete(ctx, space, block).RequestBody(requestBody).Authorization(authorization).Execute()

Delete CIDR Reservations



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	requestBody := []string{"Property_example"} // []string | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SpacesAPI.DeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsDelete(context.Background(), space, block).RequestBody(requestBody).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.DeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **[]string** |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationDelete

> DeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationDelete(ctx, space, block, reservation).Authorization(authorization).Execute()

Delete CIDR Reservation



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	reservation := "reservation_example" // string | ID of the target Reservation
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SpacesAPI.DeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationDelete(context.Background(), space, block, reservation).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.DeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**reservation** | **string** | ID of the target Reservation | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Azure Bearer token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalDelete

> interface{} DeleteExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalDelete(ctx, space, block, external).Force(force).Authorization(authorization).Execute()

Remove External Network



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target external network
	force := true // bool | Forcefully delete an External Network with existing Subnets (optional)
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.DeleteExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalDelete(context.Background(), space, block, external).Force(force).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.DeleteExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.DeleteExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target external network | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | Forcefully delete an External Network with existing Subnets | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetDelete

> interface{} DeleteExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetDelete(ctx, space, block, external, subnet).Force(force).Authorization(authorization).Execute()

Remove External Network Subnet



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target external network
	subnet := "subnet_example" // string | Name of the target external subnet
	force := true // bool | Forcefully delete an External Network with existing Subnets (optional)
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.DeleteExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetDelete(context.Background(), space, block, external, subnet).Force(force).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.DeleteExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.DeleteExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target external network | 
**subnet** | **string** | Name of the target external subnet | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **force** | **bool** | Forcefully delete an External Network with existing Subnets | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointDelete

> interface{} DeleteExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointDelete(ctx, space, block, external, subnet, endpoint).Authorization(authorization).Execute()

Remove External Network Subnet Endpoint



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target external network
	subnet := "subnet_example" // string | Name of the target external subnet
	endpoint := "endpoint_example" // string | Name of the target external subnet endpoint
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.DeleteExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointDelete(context.Background(), space, block, external, subnet, endpoint).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.DeleteExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.DeleteExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target external network | 
**subnet** | **string** | Name of the target external subnet | 
**endpoint** | **string** | Name of the target external subnet endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **authorization** | **string** | Azure Bearer token | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsDelete

> interface{} DeleteExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsDelete(ctx, space, block, external, subnet).RequestBody(requestBody).Authorization(authorization).Execute()

Remove External Network Subnet Endpoints



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target External Network
	subnet := "subnet_example" // string | Name of the target External Network Subnet
	requestBody := []string{"Property_example"} // []string | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.DeleteExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsDelete(context.Background(), space, block, external, subnet).RequestBody(requestBody).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.DeleteExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.DeleteExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target External Network | 
**subnet** | **string** | Name of the target External Network Subnet | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **requestBody** | **[]string** |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpaceApiSpacesSpaceDelete

> interface{} DeleteSpaceApiSpacesSpaceDelete(ctx, space).Force(force).Authorization(authorization).Execute()

Delete a Space



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
	space := "space_example" // string | Name of the target Space
	force := true // bool | Forcefully delete a Space with existing Blocks (optional)
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.DeleteSpaceApiSpacesSpaceDelete(context.Background(), space).Force(force).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.DeleteSpaceApiSpacesSpaceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSpaceApiSpacesSpaceDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.DeleteSpaceApiSpacesSpaceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpaceApiSpacesSpaceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Forcefully delete a Space with existing Blocks | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockApiSpacesSpaceBlocksBlockGet

> ResponseGetBlockApiSpacesSpaceBlocksBlockGet GetBlockApiSpacesSpaceBlocksBlockGet(ctx, space, block).Expand(expand).Utilization(utilization).Authorization(authorization).Execute()

Get Block Details



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	expand := true // bool | Expand network references to full network objects (optional) (default to false)
	utilization := true // bool | Append utilization information for each network (optional) (default to false)
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetBlockApiSpacesSpaceBlocksBlockGet(context.Background(), space, block).Expand(expand).Utilization(utilization).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetBlockApiSpacesSpaceBlocksBlockGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockApiSpacesSpaceBlocksBlockGet`: ResponseGetBlockApiSpacesSpaceBlocksBlockGet
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetBlockApiSpacesSpaceBlocksBlockGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockApiSpacesSpaceBlocksBlockGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **bool** | Expand network references to full network objects | [default to false]
 **utilization** | **bool** | Append utilization information for each network | [default to false]
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ResponseGetBlockApiSpacesSpaceBlocksBlockGet**](ResponseGetBlockApiSpacesSpaceBlocksBlockGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockNetsApiSpacesSpaceBlocksBlockNetworksGet

> ResponseGetBlockNetsApiSpacesSpaceBlocksBlockNetworksGet GetBlockNetsApiSpacesSpaceBlocksBlockNetworksGet(ctx, space, block).Expand(expand).Authorization(authorization).Execute()

List Block Networks



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	expand := true // bool | Expand network references to full network objects (optional) (default to false)
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetBlockNetsApiSpacesSpaceBlocksBlockNetworksGet(context.Background(), space, block).Expand(expand).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetBlockNetsApiSpacesSpaceBlocksBlockNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockNetsApiSpacesSpaceBlocksBlockNetworksGet`: ResponseGetBlockNetsApiSpacesSpaceBlocksBlockNetworksGet
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetBlockNetsApiSpacesSpaceBlocksBlockNetworksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockNetsApiSpacesSpaceBlocksBlockNetworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **bool** | Expand network references to full network objects | [default to false]
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ResponseGetBlockNetsApiSpacesSpaceBlocksBlockNetworksGet**](ResponseGetBlockNetsApiSpacesSpaceBlocksBlockNetworksGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsGet

> []ReservationExpand GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsGet(ctx, space, block).Settled(settled).Authorization(authorization).Execute()

Get Block Reservations



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	settled := true // bool | Include settled reservations. (optional) (default to false)
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsGet(context.Background(), space, block).Settled(settled).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsGet`: []ReservationExpand
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockReservationsApiSpacesSpaceBlocksBlockReservationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **settled** | **bool** | Include settled reservations. | [default to false]
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**[]ReservationExpand**](ReservationExpand.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationGet

> ReservationExpand GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationGet(ctx, space, block, reservation).Authorization(authorization).Execute()

Get Block Reservation



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	reservation := "reservation_example" // string | ID of the target Reservation
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationGet(context.Background(), space, block, reservation).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationGet`: ReservationExpand
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**reservation** | **string** | ID of the target Reservation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockReservationsApiSpacesSpaceBlocksBlockReservationsReservationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ReservationExpand**](ReservationExpand.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlocksApiSpacesSpaceBlocksGet

> ResponseGetBlocksApiSpacesSpaceBlocksGet GetBlocksApiSpacesSpaceBlocksGet(ctx, space).Expand(expand).Utilization(utilization).Authorization(authorization).Execute()

Get all Blocks within a Space



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
	space := "space_example" // string | Name of the target Space
	expand := true // bool | Expand network references to full network objects (optional) (default to false)
	utilization := true // bool | Append utilization information for each network (optional) (default to false)
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetBlocksApiSpacesSpaceBlocksGet(context.Background(), space).Expand(expand).Utilization(utilization).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetBlocksApiSpacesSpaceBlocksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlocksApiSpacesSpaceBlocksGet`: ResponseGetBlocksApiSpacesSpaceBlocksGet
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetBlocksApiSpacesSpaceBlocksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlocksApiSpacesSpaceBlocksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **bool** | Expand network references to full network objects | [default to false]
 **utilization** | **bool** | Append utilization information for each network | [default to false]
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ResponseGetBlocksApiSpacesSpaceBlocksGet**](ResponseGetBlocksApiSpacesSpaceBlocksGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalGet

> ExtNet GetExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalGet(ctx, space, block, external).Authorization(authorization).Execute()

Get External Network



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target external network
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalGet(context.Background(), space, block, external).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalGet`: ExtNet
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target external network | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalNetworkApiSpacesSpaceBlocksBlockExternalsExternalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ExtNet**](ExtNet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalNetworksApiSpacesSpaceBlocksBlockExternalsGet

> []ExtNet GetExternalNetworksApiSpacesSpaceBlocksBlockExternalsGet(ctx, space, block).Authorization(authorization).Execute()

List External Networks



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetExternalNetworksApiSpacesSpaceBlocksBlockExternalsGet(context.Background(), space, block).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetExternalNetworksApiSpacesSpaceBlocksBlockExternalsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalNetworksApiSpacesSpaceBlocksBlockExternalsGet`: []ExtNet
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetExternalNetworksApiSpacesSpaceBlocksBlockExternalsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalNetworksApiSpacesSpaceBlocksBlockExternalsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Azure Bearer token | 

### Return type

[**[]ExtNet**](ExtNet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetGet

> ExtSubnet GetExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetGet(ctx, space, block, external, subnet).Authorization(authorization).Execute()

Get External Network Subnet



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target external network
	subnet := "subnet_example" // string | Name of the target external subnet
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetGet(context.Background(), space, block, external, subnet).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetGet`: ExtSubnet
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target external network | 
**subnet** | **string** | Name of the target external subnet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ExtSubnet**](ExtSubnet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointGet

> ExtEndpoint GetExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointGet(ctx, space, block, external, subnet, endpoint).Authorization(authorization).Execute()

Get External Network Subnet Endpoint



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target external network
	subnet := "subnet_example" // string | Name of the target external subnet
	endpoint := "endpoint_example" // string | Name of the target external subnet endpoint
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointGet(context.Background(), space, block, external, subnet, endpoint).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointGet`: ExtEndpoint
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target external network | 
**subnet** | **string** | Name of the target external subnet | 
**endpoint** | **string** | Name of the target external subnet endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalSubnetEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ExtEndpoint**](ExtEndpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsGet

> []ExtEndpoint GetExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsGet(ctx, space, block, external, subnet).Authorization(authorization).Execute()

List External Network Subnet Endpoints



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target External Network
	subnet := "subnet_example" // string | Name of the target External Network Subnet
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsGet(context.Background(), space, block, external, subnet).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsGet`: []ExtEndpoint
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target External Network | 
**subnet** | **string** | Name of the target External Network Subnet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalSubnetEndpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Azure Bearer token | 

### Return type

[**[]ExtEndpoint**](ExtEndpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalSubnetsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsGet

> []ExtSubnet GetExternalSubnetsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsGet(ctx, space, block, external).Authorization(authorization).Execute()

List External Network Subnets



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target External Network
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetExternalSubnetsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsGet(context.Background(), space, block, external).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetExternalSubnetsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalSubnetsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsGet`: []ExtSubnet
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetExternalSubnetsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target External Network | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalSubnetsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Azure Bearer token | 

### Return type

[**[]ExtSubnet**](ExtSubnet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultiBlockReservationsApiSpacesSpaceReservationsGet

> []ReservationExpand GetMultiBlockReservationsApiSpacesSpaceReservationsGet(ctx, space).Settled(settled).Authorization(authorization).Execute()

Get Reservations for all Blocks within a Space



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
	space := "space_example" // string | Name of the target Space
	settled := true // bool | Include settled reservations. (optional) (default to false)
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetMultiBlockReservationsApiSpacesSpaceReservationsGet(context.Background(), space).Settled(settled).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetMultiBlockReservationsApiSpacesSpaceReservationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultiBlockReservationsApiSpacesSpaceReservationsGet`: []ReservationExpand
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetMultiBlockReservationsApiSpacesSpaceReservationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultiBlockReservationsApiSpacesSpaceReservationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settled** | **bool** | Include settled reservations. | [default to false]
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**[]ReservationExpand**](ReservationExpand.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceApiSpacesSpaceGet

> ResponseGetSpaceApiSpacesSpaceGet GetSpaceApiSpacesSpaceGet(ctx, space).Expand(expand).Utilization(utilization).Authorization(authorization).Execute()

Get Space Details



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
	space := "space_example" // string | Name of the target Space
	expand := true // bool | Expand network references to full network objects (optional) (default to false)
	utilization := true // bool | Append utilization information for each network (optional) (default to false)
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetSpaceApiSpacesSpaceGet(context.Background(), space).Expand(expand).Utilization(utilization).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetSpaceApiSpacesSpaceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceApiSpacesSpaceGet`: ResponseGetSpaceApiSpacesSpaceGet
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetSpaceApiSpacesSpaceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceApiSpacesSpaceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **bool** | Expand network references to full network objects | [default to false]
 **utilization** | **bool** | Append utilization information for each network | [default to false]
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ResponseGetSpaceApiSpacesSpaceGet**](ResponseGetSpaceApiSpacesSpaceGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpacesApiSpacesGet

> ResponseGetSpacesApiSpacesGet GetSpacesApiSpacesGet(ctx).Expand(expand).Utilization(utilization).Authorization(authorization).Execute()

Get All Spaces



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
	expand := true // bool | Expand network references to full network objects (optional) (default to false)
	utilization := true // bool | Append utilization information for each network (optional) (default to false)
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetSpacesApiSpacesGet(context.Background()).Expand(expand).Utilization(utilization).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetSpacesApiSpacesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpacesApiSpacesGet`: ResponseGetSpacesApiSpacesGet
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetSpacesApiSpacesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacesApiSpacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **bool** | Expand network references to full network objects | [default to false]
 **utilization** | **bool** | Append utilization information for each network | [default to false]
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ResponseGetSpacesApiSpacesGet**](ResponseGetSpacesApiSpacesGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlockApiSpacesSpaceBlocksBlockPatch

> Block UpdateBlockApiSpacesSpaceBlocksBlockPatch(ctx, space, block).JSONPatch(jSONPatch).Authorization(authorization).Execute()

Update Block Details



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	jSONPatch := []openapiclient.JSONPatch{*openapiclient.NewJSONPatch("Op_example", "Path_example")} // []JSONPatch | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.UpdateBlockApiSpacesSpaceBlocksBlockPatch(context.Background(), space, block).JSONPatch(jSONPatch).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.UpdateBlockApiSpacesSpaceBlocksBlockPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBlockApiSpacesSpaceBlocksBlockPatch`: Block
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.UpdateBlockApiSpacesSpaceBlocksBlockPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlockApiSpacesSpaceBlocksBlockPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jSONPatch** | [**[]JSONPatch**](JSONPatch.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlockVnetsApiSpacesSpaceBlocksBlockNetworksPut

> []Network UpdateBlockVnetsApiSpacesSpaceBlocksBlockNetworksPut(ctx, space, block).RequestBody(requestBody).Authorization(authorization).Execute()

Replace Block Networks



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	requestBody := []string{"Property_example"} // []string | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.UpdateBlockVnetsApiSpacesSpaceBlocksBlockNetworksPut(context.Background(), space, block).RequestBody(requestBody).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.UpdateBlockVnetsApiSpacesSpaceBlocksBlockNetworksPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBlockVnetsApiSpacesSpaceBlocksBlockNetworksPut`: []Network
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.UpdateBlockVnetsApiSpacesSpaceBlocksBlockNetworksPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlockVnetsApiSpacesSpaceBlocksBlockNetworksPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **[]string** |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**[]Network**](Network.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExtEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointPatch

> ExtEndpoint UpdateExtEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointPatch(ctx, space, block, external, subnet, endpoint).JSONPatch(jSONPatch).Authorization(authorization).Execute()

Update External Endpoint Details



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target External Network
	subnet := "subnet_example" // string | Name of the target external subnet
	endpoint := "endpoint_example" // string | Name of the target external subnet endpoint
	jSONPatch := []openapiclient.JSONPatch{*openapiclient.NewJSONPatch("Op_example", "Path_example")} // []JSONPatch | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.UpdateExtEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointPatch(context.Background(), space, block, external, subnet, endpoint).JSONPatch(jSONPatch).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.UpdateExtEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExtEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointPatch`: ExtEndpoint
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.UpdateExtEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target External Network | 
**subnet** | **string** | Name of the target external subnet | 
**endpoint** | **string** | Name of the target external subnet endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtEndpointApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsEndpointPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **jSONPatch** | [**[]JSONPatch**](JSONPatch.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ExtEndpoint**](ExtEndpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExtNetworkApiSpacesSpaceBlocksBlockExternalsExternalPatch

> ExtNet UpdateExtNetworkApiSpacesSpaceBlocksBlockExternalsExternalPatch(ctx, space, block, external).JSONPatch(jSONPatch).Authorization(authorization).Execute()

Update External Network Details



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target External Network
	jSONPatch := []openapiclient.JSONPatch{*openapiclient.NewJSONPatch("Op_example", "Path_example")} // []JSONPatch | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.UpdateExtNetworkApiSpacesSpaceBlocksBlockExternalsExternalPatch(context.Background(), space, block, external).JSONPatch(jSONPatch).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.UpdateExtNetworkApiSpacesSpaceBlocksBlockExternalsExternalPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExtNetworkApiSpacesSpaceBlocksBlockExternalsExternalPatch`: ExtNet
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.UpdateExtNetworkApiSpacesSpaceBlocksBlockExternalsExternalPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target External Network | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtNetworkApiSpacesSpaceBlocksBlockExternalsExternalPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **jSONPatch** | [**[]JSONPatch**](JSONPatch.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ExtNet**](ExtNet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExtSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetPatch

> ExtSubnet UpdateExtSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetPatch(ctx, space, block, external, subnet).JSONPatch(jSONPatch).Authorization(authorization).Execute()

Update External Subnet Details



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target External Network
	subnet := "subnet_example" // string | Name of the target external subnet
	jSONPatch := []openapiclient.JSONPatch{*openapiclient.NewJSONPatch("Op_example", "Path_example")} // []JSONPatch | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.UpdateExtSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetPatch(context.Background(), space, block, external, subnet).JSONPatch(jSONPatch).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.UpdateExtSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExtSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetPatch`: ExtSubnet
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.UpdateExtSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target External Network | 
**subnet** | **string** | Name of the target external subnet | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtSubnetApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **jSONPatch** | [**[]JSONPatch**](JSONPatch.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ExtSubnet**](ExtSubnet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExternalSubnetEnpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPut

> []ExtEndpoint UpdateExternalSubnetEnpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPut(ctx, space, block, external, subnet).ExtEndpointReq(extEndpointReq).Authorization(authorization).Execute()

Replace External Network Subnet Endpoints



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
	space := "space_example" // string | Name of the target Space
	block := "block_example" // string | Name of the target Block
	external := "external_example" // string | Name of the target External Network
	subnet := "subnet_example" // string | Name of the target External Network Subnet
	extEndpointReq := []openapiclient.ExtEndpointReq{*openapiclient.NewExtEndpointReq("Name_example", "Desc_example", "Ip_example")} // []ExtEndpointReq | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.UpdateExternalSubnetEnpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPut(context.Background(), space, block, external, subnet).ExtEndpointReq(extEndpointReq).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.UpdateExternalSubnetEnpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExternalSubnetEnpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPut`: []ExtEndpoint
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.UpdateExternalSubnetEnpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 
**block** | **string** | Name of the target Block | 
**external** | **string** | Name of the target External Network | 
**subnet** | **string** | Name of the target External Network Subnet | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalSubnetEnpointsApiSpacesSpaceBlocksBlockExternalsExternalSubnetsSubnetEndpointsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **extEndpointReq** | [**[]ExtEndpointReq**](ExtEndpointReq.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**[]ExtEndpoint**](ExtEndpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpaceApiSpacesSpacePatch

> interface{} UpdateSpaceApiSpacesSpacePatch(ctx, space).JSONPatch(jSONPatch).Authorization(authorization).Execute()

Update Space Details



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
	space := "space_example" // string | Name of the target Space
	jSONPatch := []openapiclient.JSONPatch{*openapiclient.NewJSONPatch("Op_example", "Path_example")} // []JSONPatch | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.UpdateSpaceApiSpacesSpacePatch(context.Background(), space).JSONPatch(jSONPatch).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.UpdateSpaceApiSpacesSpacePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSpaceApiSpacesSpacePatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.UpdateSpaceApiSpacesSpacePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Name of the target Space | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpaceApiSpacesSpacePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSONPatch** | [**[]JSONPatch**](JSONPatch.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

