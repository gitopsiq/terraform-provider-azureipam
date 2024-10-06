# \AdminAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddExclusionsApiAdminExclusionsPost**](AdminAPI.md#AddExclusionsApiAdminExclusionsPost) | **Post** /api/admin/exclusions | Add Excluded Subscription(s)
[**CreateAdminApiAdminAdminsPost**](AdminAPI.md#CreateAdminApiAdminAdminsPost) | **Post** /api/admin/admins | Create IPAM Admin
[**DeleteAdminApiAdminAdminsObjectIdDelete**](AdminAPI.md#DeleteAdminApiAdminAdminsObjectIdDelete) | **Delete** /api/admin/admins/{objectId} | Delete IPAM Admin
[**GetAdminsApiAdminAdminsGet**](AdminAPI.md#GetAdminsApiAdminAdminsGet) | **Get** /api/admin/admins | Get All Admins
[**GetAdminsApiAdminAdminsObjectIdGet**](AdminAPI.md#GetAdminsApiAdminAdminsObjectIdGet) | **Get** /api/admin/admins/{objectId} | Get IPAM Admin
[**GetExclusionsApiAdminExclusionsGet**](AdminAPI.md#GetExclusionsApiAdminExclusionsGet) | **Get** /api/admin/exclusions | Get Excluded Subscriptions
[**RemoveExclusionApiAdminExclusionsSubscriptionIdDelete**](AdminAPI.md#RemoveExclusionApiAdminExclusionsSubscriptionIdDelete) | **Delete** /api/admin/exclusions/{subscriptionId} | Remove Excluded Subscription
[**UpdateAdminsApiAdminAdminsPut**](AdminAPI.md#UpdateAdminsApiAdminAdminsPut) | **Put** /api/admin/admins | Replace IPAM Admins
[**UpdateExclusionsApiAdminExclusionsPut**](AdminAPI.md#UpdateExclusionsApiAdminExclusionsPut) | **Put** /api/admin/exclusions | Replace Excluded Subscriptions



## AddExclusionsApiAdminExclusionsPost

> interface{} AddExclusionsApiAdminExclusionsPost(ctx).RequestBody(requestBody).Authorization(authorization).Execute()

Add Excluded Subscription(s)



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
	requestBody := []string{"Property_example"} // []string | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AddExclusionsApiAdminExclusionsPost(context.Background()).RequestBody(requestBody).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AddExclusionsApiAdminExclusionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddExclusionsApiAdminExclusionsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AddExclusionsApiAdminExclusionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddExclusionsApiAdminExclusionsPostRequest struct via the builder pattern


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


## CreateAdminApiAdminAdminsPost

> interface{} CreateAdminApiAdminAdminsPost(ctx).Admin(admin).Authorization(authorization).Execute()

Create IPAM Admin



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
	admin := *openapiclient.NewAdmin("Type_example", "Name_example", "Id_example") // Admin | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CreateAdminApiAdminAdminsPost(context.Background()).Admin(admin).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CreateAdminApiAdminAdminsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAdminApiAdminAdminsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CreateAdminApiAdminAdminsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminApiAdminAdminsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **admin** | [**Admin**](Admin.md) |  | 
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


## DeleteAdminApiAdminAdminsObjectIdDelete

> interface{} DeleteAdminApiAdminAdminsObjectIdDelete(ctx, objectId).Authorization(authorization).Execute()

Delete IPAM Admin



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
	objectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Azure AD ObjectID for the target user
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.DeleteAdminApiAdminAdminsObjectIdDelete(context.Background(), objectId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.DeleteAdminApiAdminAdminsObjectIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAdminApiAdminAdminsObjectIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.DeleteAdminApiAdminAdminsObjectIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | Azure AD ObjectID for the target user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdminApiAdminAdminsObjectIdDeleteRequest struct via the builder pattern


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


## GetAdminsApiAdminAdminsGet

> []Admin GetAdminsApiAdminAdminsGet(ctx).Authorization(authorization).Execute()

Get All Admins



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
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetAdminsApiAdminAdminsGet(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdminsApiAdminAdminsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminsApiAdminAdminsGet`: []Admin
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdminsApiAdminAdminsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminsApiAdminAdminsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**[]Admin**](Admin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminsApiAdminAdminsObjectIdGet

> Admin GetAdminsApiAdminAdminsObjectIdGet(ctx, objectId).Authorization(authorization).Execute()

Get IPAM Admin



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
	objectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Azure AD ObjectID for the target user
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetAdminsApiAdminAdminsObjectIdGet(context.Background(), objectId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdminsApiAdminAdminsObjectIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminsApiAdminAdminsObjectIdGet`: Admin
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdminsApiAdminAdminsObjectIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | Azure AD ObjectID for the target user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminsApiAdminAdminsObjectIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Azure Bearer token | 

### Return type

[**Admin**](Admin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExclusionsApiAdminExclusionsGet

> []string GetExclusionsApiAdminExclusionsGet(ctx).Authorization(authorization).Execute()

Get Excluded Subscriptions



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
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetExclusionsApiAdminExclusionsGet(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetExclusionsApiAdminExclusionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExclusionsApiAdminExclusionsGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetExclusionsApiAdminExclusionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExclusionsApiAdminExclusionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Azure Bearer token | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveExclusionApiAdminExclusionsSubscriptionIdDelete

> interface{} RemoveExclusionApiAdminExclusionsSubscriptionIdDelete(ctx, subscriptionId).Authorization(authorization).Execute()

Remove Excluded Subscription



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
	subscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Azure Subscription ID
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.RemoveExclusionApiAdminExclusionsSubscriptionIdDelete(context.Background(), subscriptionId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.RemoveExclusionApiAdminExclusionsSubscriptionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveExclusionApiAdminExclusionsSubscriptionIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.RemoveExclusionApiAdminExclusionsSubscriptionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Azure Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveExclusionApiAdminExclusionsSubscriptionIdDeleteRequest struct via the builder pattern


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


## UpdateAdminsApiAdminAdminsPut

> interface{} UpdateAdminsApiAdminAdminsPut(ctx).Admin(admin).Authorization(authorization).Execute()

Replace IPAM Admins



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
	admin := []openapiclient.Admin{*openapiclient.NewAdmin("Type_example", "Name_example", "Id_example")} // []Admin | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.UpdateAdminsApiAdminAdminsPut(context.Background()).Admin(admin).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.UpdateAdminsApiAdminAdminsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdminsApiAdminAdminsPut`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.UpdateAdminsApiAdminAdminsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdminsApiAdminAdminsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **admin** | [**[]Admin**](Admin.md) |  | 
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


## UpdateExclusionsApiAdminExclusionsPut

> interface{} UpdateExclusionsApiAdminExclusionsPut(ctx).RequestBody(requestBody).Authorization(authorization).Execute()

Replace Excluded Subscriptions



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
	requestBody := []string{"Property_example"} // []string | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.UpdateExclusionsApiAdminExclusionsPut(context.Background()).RequestBody(requestBody).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.UpdateExclusionsApiAdminExclusionsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExclusionsApiAdminExclusionsPut`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.UpdateExclusionsApiAdminExclusionsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExclusionsApiAdminExclusionsPutRequest struct via the builder pattern


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

