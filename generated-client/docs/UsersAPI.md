# \UsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserApiUsersMeGet**](UsersAPI.md#GetUserApiUsersMeGet) | **Get** /api/users/me | Get My User Details
[**GetUsersApiUsersGet**](UsersAPI.md#GetUsersApiUsersGet) | **Get** /api/users | Get All Users
[**UpdateUserApiUsersMePatch**](UsersAPI.md#UpdateUserApiUsersMePatch) | **Patch** /api/users/me | Update User Details



## GetUserApiUsersMeGet

> ResponseGetUserApiUsersMeGet GetUserApiUsersMeGet(ctx).Expand(expand).Authorization(authorization).Execute()

Get My User Details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gitopsiq/terraform-provider-azureipam/generated-client"
)

func main() {
	expand := true // bool | Show expanded user details (optional) (default to false)
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserApiUsersMeGet(context.Background()).Expand(expand).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserApiUsersMeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserApiUsersMeGet`: ResponseGetUserApiUsersMeGet
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserApiUsersMeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserApiUsersMeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **bool** | Show expanded user details | [default to false]
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**ResponseGetUserApiUsersMeGet**](ResponseGetUserApiUsersMeGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersApiUsersGet

> []User GetUsersApiUsersGet(ctx).Authorization(authorization).Execute()

Get All Users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gitopsiq/terraform-provider-azureipam/generated-client"
)

func main() {
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUsersApiUsersGet(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersApiUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersApiUsersGet`: []User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersApiUsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersApiUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserApiUsersMePatch

> User UpdateUserApiUsersMePatch(ctx).JSONPatch(jSONPatch).Authorization(authorization).Execute()

Update User Details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gitopsiq/terraform-provider-azureipam/generated-client"
)

func main() {
	jSONPatch := []openapiclient.JSONPatch{*openapiclient.NewJSONPatch("Op_example", "Path_example")} // []JSONPatch | 
	authorization := "authorization_example" // string | Azure Bearer token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserApiUsersMePatch(context.Background()).JSONPatch(jSONPatch).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserApiUsersMePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserApiUsersMePatch`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserApiUsersMePatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserApiUsersMePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jSONPatch** | [**[]JSONPatch**](JSONPatch.md) |  | 
 **authorization** | **string** | Azure Bearer token | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

