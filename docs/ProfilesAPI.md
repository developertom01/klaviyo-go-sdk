# \ProfilesAPI

All URIs are relative to *https://a.klaviyo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateProfile**](ProfilesAPI.md#CreateOrUpdateProfile) | **Post** /api/profile-import/ | Create or Update Profile
[**CreateProfile**](ProfilesAPI.md#CreateProfile) | **Post** /api/profiles/ | Create Profile
[**CreatePushToken**](ProfilesAPI.md#CreatePushToken) | **Post** /api/push-tokens/ | Create or Update Push Token
[**GetBulkProfileImportJob**](ProfilesAPI.md#GetBulkProfileImportJob) | **Get** /api/profile-bulk-import-jobs/{job_id}/ | Get Bulk Profile Import Job
[**GetBulkProfileImportJobImportErrors**](ProfilesAPI.md#GetBulkProfileImportJobImportErrors) | **Get** /api/profile-bulk-import-jobs/{id}/import-errors/ | Get Bulk Profile Import Job Errors
[**GetBulkProfileImportJobLists**](ProfilesAPI.md#GetBulkProfileImportJobLists) | **Get** /api/profile-bulk-import-jobs/{id}/lists/ | Get Bulk Profile Import Job Lists
[**GetBulkProfileImportJobProfiles**](ProfilesAPI.md#GetBulkProfileImportJobProfiles) | **Get** /api/profile-bulk-import-jobs/{id}/profiles/ | Get Bulk Profile Import Job Profiles
[**GetBulkProfileImportJobRelationshipsLists**](ProfilesAPI.md#GetBulkProfileImportJobRelationshipsLists) | **Get** /api/profile-bulk-import-jobs/{id}/relationships/lists/ | Get Bulk Profile Import Job Relationships Lists
[**GetBulkProfileImportJobRelationshipsProfiles**](ProfilesAPI.md#GetBulkProfileImportJobRelationshipsProfiles) | **Get** /api/profile-bulk-import-jobs/{id}/relationships/profiles/ | Get Bulk Profile Import Job Relationships Profiles
[**GetBulkProfileImportJobs**](ProfilesAPI.md#GetBulkProfileImportJobs) | **Get** /api/profile-bulk-import-jobs/ | Get Bulk Profile Import Jobs
[**GetProfile**](ProfilesAPI.md#GetProfile) | **Get** /api/profiles/{id}/ | Get Profile
[**GetProfileLists**](ProfilesAPI.md#GetProfileLists) | **Get** /api/profiles/{id}/lists/ | Get Profile Lists
[**GetProfileRelationshipsLists**](ProfilesAPI.md#GetProfileRelationshipsLists) | **Get** /api/profiles/{id}/relationships/lists/ | Get Profile Relationships Lists
[**GetProfileRelationshipsSegments**](ProfilesAPI.md#GetProfileRelationshipsSegments) | **Get** /api/profiles/{id}/relationships/segments/ | Get Profile Relationships Segments
[**GetProfileSegments**](ProfilesAPI.md#GetProfileSegments) | **Get** /api/profiles/{id}/segments/ | Get Profile Segments
[**GetProfiles**](ProfilesAPI.md#GetProfiles) | **Get** /api/profiles/ | Get Profiles
[**MergeProfiles**](ProfilesAPI.md#MergeProfiles) | **Post** /api/profile-merge/ | Merge Profiles
[**SpawnBulkProfileImportJob**](ProfilesAPI.md#SpawnBulkProfileImportJob) | **Post** /api/profile-bulk-import-jobs/ | Spawn Bulk Profile Import Job
[**SubscribeProfiles**](ProfilesAPI.md#SubscribeProfiles) | **Post** /api/profile-subscription-bulk-create-jobs/ | Subscribe Profiles
[**SuppressProfiles**](ProfilesAPI.md#SuppressProfiles) | **Post** /api/profile-suppression-bulk-create-jobs/ | Suppress Profiles
[**UnsubscribeProfiles**](ProfilesAPI.md#UnsubscribeProfiles) | **Post** /api/profile-subscription-bulk-delete-jobs/ | Unsubscribe Profiles
[**UnsuppressProfiles**](ProfilesAPI.md#UnsuppressProfiles) | **Post** /api/profile-suppression-bulk-delete-jobs/ | Unsuppress Profiles
[**UpdateProfile**](ProfilesAPI.md#UpdateProfile) | **Patch** /api/profiles/{id}/ | Update Profile



## CreateOrUpdateProfile

> PostProfileResponse CreateOrUpdateProfile(ctx).Revision(revision).ProfileUpsertQuery(profileUpsertQuery).Execute()

Create or Update Profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	profileUpsertQuery := *openapiclient.NewProfileUpsertQuery(*openapiclient.NewProfileUpsertQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewProfileUpsertQueryResourceObjectAttributes())) // ProfileUpsertQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateOrUpdateProfile(context.Background()).Revision(revision).ProfileUpsertQuery(profileUpsertQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateOrUpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrUpdateProfile`: PostProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateOrUpdateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **profileUpsertQuery** | [**ProfileUpsertQuery**](ProfileUpsertQuery.md) |  | 

### Return type

[**PostProfileResponse**](PostProfileResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProfile

> PostProfileResponse CreateProfile(ctx).Revision(revision).ProfileCreateQuery(profileCreateQuery).Execute()

Create Profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	profileCreateQuery := *openapiclient.NewProfileCreateQuery(*openapiclient.NewProfileCreateQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewProfileCreateQueryResourceObjectAttributes())) // ProfileCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateProfile(context.Background()).Revision(revision).ProfileCreateQuery(profileCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProfile`: PostProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **profileCreateQuery** | [**ProfileCreateQuery**](ProfileCreateQuery.md) |  | 

### Return type

[**PostProfileResponse**](PostProfileResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePushToken

> CreatePushToken(ctx).Revision(revision).PushTokenCreateQuery(pushTokenCreateQuery).Execute()

Create or Update Push Token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	pushTokenCreateQuery := *openapiclient.NewPushTokenCreateQuery(*openapiclient.NewPushTokenCreateQueryResourceObject(openapiclient.PushTokenEnum("push-token"), *openapiclient.NewPushTokenCreateQueryResourceObjectAttributes("1234567890", "Platform_example", "APNs", *openapiclient.NewPushTokenCreateQueryResourceObjectAttributesProfile(*openapiclient.NewProfileUpsertQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewProfileUpsertQueryResourceObjectAttributes()))))) // PushTokenCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProfilesAPI.CreatePushToken(context.Background()).Revision(revision).PushTokenCreateQuery(pushTokenCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreatePushToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePushTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **pushTokenCreateQuery** | [**PushTokenCreateQuery**](PushTokenCreateQuery.md) |  | 

### Return type

 (empty response body)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkProfileImportJob

> GetProfileImportJobResponseCompoundDocument GetBulkProfileImportJob(ctx, jobId).Revision(revision).FieldsList(fieldsList).FieldsProfileBulkImportJob(fieldsProfileBulkImportJob).Include(include).Execute()

Get Bulk Profile Import Job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	jobId := "01GSQPBF74KQ5YTDEPP41T1BZH" // string | ID of the job to retrieve.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsList := []string{"FieldsList_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsProfileBulkImportJob := []string{"FieldsProfileBulkImportJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetBulkProfileImportJob(context.Background(), jobId).Revision(revision).FieldsList(fieldsList).FieldsProfileBulkImportJob(fieldsProfileBulkImportJob).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetBulkProfileImportJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkProfileImportJob`: GetProfileImportJobResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetBulkProfileImportJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | ID of the job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkProfileImportJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsList** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsProfileBulkImportJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetProfileImportJobResponseCompoundDocument**](GetProfileImportJobResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkProfileImportJobImportErrors

> GetImportErrorResponseCollection GetBulkProfileImportJobImportErrors(ctx, id).Revision(revision).FieldsImportError(fieldsImportError).PageCursor(pageCursor).PageSize(pageSize).Execute()

Get Bulk Profile Import Job Errors



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	id := "id_example" // string | 
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsImportError := []string{"FieldsImportError_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	pageSize := int32(56) // int32 | Default: 20. Min: 1. Max: 100. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetBulkProfileImportJobImportErrors(context.Background(), id).Revision(revision).FieldsImportError(fieldsImportError).PageCursor(pageCursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetBulkProfileImportJobImportErrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkProfileImportJobImportErrors`: GetImportErrorResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetBulkProfileImportJobImportErrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkProfileImportJobImportErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsImportError** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **pageSize** | **int32** | Default: 20. Min: 1. Max: 100. | [default to 20]

### Return type

[**GetImportErrorResponseCollection**](GetImportErrorResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkProfileImportJobLists

> GetListResponseCollection GetBulkProfileImportJobLists(ctx, id).Revision(revision).FieldsList(fieldsList).Execute()

Get Bulk Profile Import Job Lists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	id := "id_example" // string | 
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsList := []string{"FieldsList_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetBulkProfileImportJobLists(context.Background(), id).Revision(revision).FieldsList(fieldsList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetBulkProfileImportJobLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkProfileImportJobLists`: GetListResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetBulkProfileImportJobLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkProfileImportJobListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsList** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetListResponseCollection**](GetListResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkProfileImportJobProfiles

> GetProfileResponseCollection GetBulkProfileImportJobProfiles(ctx, id).Revision(revision).AdditionalFieldsProfile(additionalFieldsProfile).FieldsProfile(fieldsProfile).PageCursor(pageCursor).PageSize(pageSize).Execute()

Get Bulk Profile Import Job Profiles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	id := "id_example" // string | 
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	additionalFieldsProfile := []string{"AdditionalFieldsProfile_example"} // []string | Request additional fields not included by default in the response. Supported values: 'subscriptions', 'predictive_analytics' (optional)
	fieldsProfile := []string{"FieldsProfile_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	pageSize := int32(56) // int32 | Default: 20. Min: 1. Max: 100. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetBulkProfileImportJobProfiles(context.Background(), id).Revision(revision).AdditionalFieldsProfile(additionalFieldsProfile).FieldsProfile(fieldsProfile).PageCursor(pageCursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetBulkProfileImportJobProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkProfileImportJobProfiles`: GetProfileResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetBulkProfileImportJobProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkProfileImportJobProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **additionalFieldsProfile** | **[]string** | Request additional fields not included by default in the response. Supported values: &#39;subscriptions&#39;, &#39;predictive_analytics&#39; | 
 **fieldsProfile** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **pageSize** | **int32** | Default: 20. Min: 1. Max: 100. | [default to 20]

### Return type

[**GetProfileResponseCollection**](GetProfileResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkProfileImportJobRelationshipsLists

> GetProfileImportJobListRelationshipsResponseCollection GetBulkProfileImportJobRelationshipsLists(ctx, id).Revision(revision).Execute()

Get Bulk Profile Import Job Relationships Lists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	id := "id_example" // string | 
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetBulkProfileImportJobRelationshipsLists(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetBulkProfileImportJobRelationshipsLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkProfileImportJobRelationshipsLists`: GetProfileImportJobListRelationshipsResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetBulkProfileImportJobRelationshipsLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkProfileImportJobRelationshipsListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetProfileImportJobListRelationshipsResponseCollection**](GetProfileImportJobListRelationshipsResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkProfileImportJobRelationshipsProfiles

> GetProfileImportJobProfileRelationshipsResponseCollection GetBulkProfileImportJobRelationshipsProfiles(ctx, id).Revision(revision).PageCursor(pageCursor).PageSize(pageSize).Execute()

Get Bulk Profile Import Job Relationships Profiles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	id := "id_example" // string | 
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	pageSize := int32(56) // int32 | Default: 20. Min: 1. Max: 100. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetBulkProfileImportJobRelationshipsProfiles(context.Background(), id).Revision(revision).PageCursor(pageCursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetBulkProfileImportJobRelationshipsProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkProfileImportJobRelationshipsProfiles`: GetProfileImportJobProfileRelationshipsResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetBulkProfileImportJobRelationshipsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkProfileImportJobRelationshipsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **pageSize** | **int32** | Default: 20. Min: 1. Max: 100. | [default to 20]

### Return type

[**GetProfileImportJobProfileRelationshipsResponseCollection**](GetProfileImportJobProfileRelationshipsResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkProfileImportJobs

> GetProfileImportJobResponseCollectionCompoundDocument GetBulkProfileImportJobs(ctx).Revision(revision).FieldsProfileBulkImportJob(fieldsProfileBulkImportJob).Filter(filter).PageCursor(pageCursor).PageSize(pageSize).Sort(sort).Execute()

Get Bulk Profile Import Jobs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsProfileBulkImportJob := []string{"FieldsProfileBulkImportJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`status`: `any`, `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	pageSize := int32(56) // int32 | Default: 20. Min: 1. Max: 100. (optional) (default to 20)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetBulkProfileImportJobs(context.Background()).Revision(revision).FieldsProfileBulkImportJob(fieldsProfileBulkImportJob).Filter(filter).PageCursor(pageCursor).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetBulkProfileImportJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkProfileImportJobs`: GetProfileImportJobResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetBulkProfileImportJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkProfileImportJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsProfileBulkImportJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;status&#x60;: &#x60;any&#x60;, &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **pageSize** | **int32** | Default: 20. Min: 1. Max: 100. | [default to 20]
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetProfileImportJobResponseCollectionCompoundDocument**](GetProfileImportJobResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfile

> GetProfileResponseCompoundDocument GetProfile(ctx, id).Revision(revision).AdditionalFieldsProfile(additionalFieldsProfile).FieldsList(fieldsList).FieldsProfile(fieldsProfile).FieldsSegment(fieldsSegment).Include(include).Execute()

Get Profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	id := "id_example" // string | 
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	additionalFieldsProfile := []string{"AdditionalFieldsProfile_example"} // []string | Request additional fields not included by default in the response. Supported values: 'subscriptions', 'predictive_analytics' (optional)
	fieldsList := []string{"FieldsList_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsProfile := []string{"FieldsProfile_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsSegment := []string{"FieldsSegment_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetProfile(context.Background(), id).Revision(revision).AdditionalFieldsProfile(additionalFieldsProfile).FieldsList(fieldsList).FieldsProfile(fieldsProfile).FieldsSegment(fieldsSegment).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfile`: GetProfileResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **additionalFieldsProfile** | **[]string** | Request additional fields not included by default in the response. Supported values: &#39;subscriptions&#39;, &#39;predictive_analytics&#39; | 
 **fieldsList** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsProfile** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsSegment** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetProfileResponseCompoundDocument**](GetProfileResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileLists

> GetListResponseCollection GetProfileLists(ctx, id).Revision(revision).FieldsList(fieldsList).Execute()

Get Profile Lists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	id := "id_example" // string | 
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsList := []string{"FieldsList_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetProfileLists(context.Background(), id).Revision(revision).FieldsList(fieldsList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetProfileLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfileLists`: GetListResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetProfileLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsList** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetListResponseCollection**](GetListResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileRelationshipsLists

> GetProfileListRelationshipsResponseCollection GetProfileRelationshipsLists(ctx, id).Revision(revision).Execute()

Get Profile Relationships Lists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	id := "id_example" // string | 
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetProfileRelationshipsLists(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetProfileRelationshipsLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfileRelationshipsLists`: GetProfileListRelationshipsResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetProfileRelationshipsLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRelationshipsListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetProfileListRelationshipsResponseCollection**](GetProfileListRelationshipsResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileRelationshipsSegments

> GetProfileSegmentRelationshipsResponseCollection GetProfileRelationshipsSegments(ctx, id).Revision(revision).Execute()

Get Profile Relationships Segments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	id := "id_example" // string | 
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetProfileRelationshipsSegments(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetProfileRelationshipsSegments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfileRelationshipsSegments`: GetProfileSegmentRelationshipsResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetProfileRelationshipsSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRelationshipsSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetProfileSegmentRelationshipsResponseCollection**](GetProfileSegmentRelationshipsResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileSegments

> GetSegmentResponseCollection GetProfileSegments(ctx, id).Revision(revision).FieldsSegment(fieldsSegment).Execute()

Get Profile Segments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	id := "id_example" // string | 
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsSegment := []string{"FieldsSegment_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetProfileSegments(context.Background(), id).Revision(revision).FieldsSegment(fieldsSegment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetProfileSegments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfileSegments`: GetSegmentResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetProfileSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsSegment** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetSegmentResponseCollection**](GetSegmentResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfiles

> GetProfileResponseCollectionCompoundDocument GetProfiles(ctx).Revision(revision).AdditionalFieldsProfile(additionalFieldsProfile).FieldsProfile(fieldsProfile).Filter(filter).PageCursor(pageCursor).PageSize(pageSize).Sort(sort).Execute()

Get Profiles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	additionalFieldsProfile := []string{"AdditionalFieldsProfile_example"} // []string | Request additional fields not included by default in the response. Supported values: 'subscriptions', 'predictive_analytics' (optional)
	fieldsProfile := []string{"FieldsProfile_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`id`: `any`, `equals`<br>`email`: `any`, `equals`<br>`phone_number`: `any`, `equals`<br>`external_id`: `any`, `equals`<br>`_kx`: `equals`<br>`created`: `greater-than`, `less-than`<br>`updated`: `greater-than`, `less-than`<br>`subscriptions.email.marketing.suppression.timestamp` : `greater-than`, `greater-or-equal`, `less-than`, `less-or-equal`<br>`subscriptions.email.marketing.suppression.reason` : `equals`<br>`subscriptions.email.marketing.list_suppressions.list_id` : `equals`<br>`subscriptions.email.marketing.list_suppressions.reason` : `equals`<br>`subscriptions.email.marketing.list_suppressions.timestamp` : `greater-than`, `greater-or-equal`, `less-than`, `less-or-equal` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	pageSize := int32(56) // int32 | Default: 20. Min: 1. Max: 100. (optional) (default to 20)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetProfiles(context.Background()).Revision(revision).AdditionalFieldsProfile(additionalFieldsProfile).FieldsProfile(fieldsProfile).Filter(filter).PageCursor(pageCursor).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfiles`: GetProfileResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **additionalFieldsProfile** | **[]string** | Request additional fields not included by default in the response. Supported values: &#39;subscriptions&#39;, &#39;predictive_analytics&#39; | 
 **fieldsProfile** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;id&#x60;: &#x60;any&#x60;, &#x60;equals&#x60;&lt;br&gt;&#x60;email&#x60;: &#x60;any&#x60;, &#x60;equals&#x60;&lt;br&gt;&#x60;phone_number&#x60;: &#x60;any&#x60;, &#x60;equals&#x60;&lt;br&gt;&#x60;external_id&#x60;: &#x60;any&#x60;, &#x60;equals&#x60;&lt;br&gt;&#x60;_kx&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;created&#x60;: &#x60;greater-than&#x60;, &#x60;less-than&#x60;&lt;br&gt;&#x60;updated&#x60;: &#x60;greater-than&#x60;, &#x60;less-than&#x60;&lt;br&gt;&#x60;subscriptions.email.marketing.suppression.timestamp&#x60; : &#x60;greater-than&#x60;, &#x60;greater-or-equal&#x60;, &#x60;less-than&#x60;, &#x60;less-or-equal&#x60;&lt;br&gt;&#x60;subscriptions.email.marketing.suppression.reason&#x60; : &#x60;equals&#x60;&lt;br&gt;&#x60;subscriptions.email.marketing.list_suppressions.list_id&#x60; : &#x60;equals&#x60;&lt;br&gt;&#x60;subscriptions.email.marketing.list_suppressions.reason&#x60; : &#x60;equals&#x60;&lt;br&gt;&#x60;subscriptions.email.marketing.list_suppressions.timestamp&#x60; : &#x60;greater-than&#x60;, &#x60;greater-or-equal&#x60;, &#x60;less-than&#x60;, &#x60;less-or-equal&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **pageSize** | **int32** | Default: 20. Min: 1. Max: 100. | [default to 20]
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetProfileResponseCollectionCompoundDocument**](GetProfileResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeProfiles

> PostProfileMergeResponse MergeProfiles(ctx).Revision(revision).ProfileMergeQuery(profileMergeQuery).Execute()

Merge Profiles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	profileMergeQuery := *openapiclient.NewProfileMergeQuery(*openapiclient.NewProfileMergeQueryResourceObject(openapiclient.ProfileMergeEnum("profile-merge"), "01GDDKASAP8TKDDA2GRZDSVP4H", *openapiclient.NewProfileMergeQueryResourceObjectRelationships(*openapiclient.NewProfileMergeQueryResourceObjectRelationshipsProfiles([]openapiclient.ProfileMergeQueryResourceObjectRelationshipsProfilesDataInner{*openapiclient.NewProfileMergeQueryResourceObjectRelationshipsProfilesDataInner(openapiclient.ProfileEnum("profile"), "01GDDKASAP8TKDDA2GRZDSVP4I")})))) // ProfileMergeQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.MergeProfiles(context.Background()).Revision(revision).ProfileMergeQuery(profileMergeQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.MergeProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MergeProfiles`: PostProfileMergeResponse
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.MergeProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMergeProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **profileMergeQuery** | [**ProfileMergeQuery**](ProfileMergeQuery.md) |  | 

### Return type

[**PostProfileMergeResponse**](PostProfileMergeResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpawnBulkProfileImportJob

> PostProfileImportJobResponse SpawnBulkProfileImportJob(ctx).Revision(revision).ProfileImportJobCreateQuery(profileImportJobCreateQuery).Execute()

Spawn Bulk Profile Import Job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	profileImportJobCreateQuery := *openapiclient.NewProfileImportJobCreateQuery(*openapiclient.NewProfileImportJobCreateQueryResourceObject(openapiclient.ProfileBulkImportJobEnum("profile-bulk-import-job"), *openapiclient.NewProfileImportJobCreateQueryResourceObjectAttributes(*openapiclient.NewProfileImportJobCreateQueryResourceObjectAttributesProfiles([]openapiclient.ProfileUpsertQueryResourceObject{*openapiclient.NewProfileUpsertQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewProfileUpsertQueryResourceObjectAttributes())})))) // ProfileImportJobCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.SpawnBulkProfileImportJob(context.Background()).Revision(revision).ProfileImportJobCreateQuery(profileImportJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.SpawnBulkProfileImportJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpawnBulkProfileImportJob`: PostProfileImportJobResponse
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.SpawnBulkProfileImportJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpawnBulkProfileImportJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **profileImportJobCreateQuery** | [**ProfileImportJobCreateQuery**](ProfileImportJobCreateQuery.md) |  | 

### Return type

[**PostProfileImportJobResponse**](PostProfileImportJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribeProfiles

> SubscribeProfiles(ctx).Revision(revision).SubscriptionCreateJobCreateQuery(subscriptionCreateJobCreateQuery).Execute()

Subscribe Profiles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	subscriptionCreateJobCreateQuery := *openapiclient.NewSubscriptionCreateJobCreateQuery(*openapiclient.NewSubscriptionCreateJobCreateQueryResourceObject(openapiclient.ProfileSubscriptionBulkCreateJobEnum("profile-subscription-bulk-create-job"), *openapiclient.NewSubscriptionCreateJobCreateQueryResourceObjectAttributes(*openapiclient.NewSubscriptionCreateJobCreateQueryResourceObjectAttributesProfiles([]openapiclient.ProfileSubscriptionCreateQueryResourceObject{*openapiclient.NewProfileSubscriptionCreateQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewProfileSubscriptionCreateQueryResourceObjectAttributes())})))) // SubscriptionCreateJobCreateQuery | Subscribes one or more profiles to marketing. Currently, supports email and SMS only. All profiles will be added to the provided list. Either email or phone number is required. Both may be specified to subscribe to both channels. If a profile cannot be found matching the given identifier(s), a new profile will be created and then subscribed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProfilesAPI.SubscribeProfiles(context.Background()).Revision(revision).SubscriptionCreateJobCreateQuery(subscriptionCreateJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.SubscribeProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **subscriptionCreateJobCreateQuery** | [**SubscriptionCreateJobCreateQuery**](SubscriptionCreateJobCreateQuery.md) | Subscribes one or more profiles to marketing. Currently, supports email and SMS only. All profiles will be added to the provided list. Either email or phone number is required. Both may be specified to subscribe to both channels. If a profile cannot be found matching the given identifier(s), a new profile will be created and then subscribed. | 

### Return type

 (empty response body)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppressProfiles

> SuppressProfiles(ctx).Revision(revision).SuppressionCreateJobCreateQuery(suppressionCreateJobCreateQuery).Execute()

Suppress Profiles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	suppressionCreateJobCreateQuery := *openapiclient.NewSuppressionCreateJobCreateQuery(*openapiclient.NewSuppressionCreateJobCreateQueryResourceObject(openapiclient.ProfileSuppressionBulkCreateJobEnum("profile-suppression-bulk-create-job"), *openapiclient.NewSuppressionCreateJobCreateQueryResourceObjectAttributes(*openapiclient.NewSuppressionCreateJobCreateQueryResourceObjectAttributesProfiles([]openapiclient.ProfileSuppressionCreateQueryResourceObject{*openapiclient.NewProfileSuppressionCreateQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewProfileSuppressionCreateQueryResourceObjectAttributes())})))) // SuppressionCreateJobCreateQuery | Suppresses one or more profiles from receiving marketing. Currently, supports email only. If a profile is not found with the given email, one will be created and immediately suppressed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProfilesAPI.SuppressProfiles(context.Background()).Revision(revision).SuppressionCreateJobCreateQuery(suppressionCreateJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.SuppressProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuppressProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **suppressionCreateJobCreateQuery** | [**SuppressionCreateJobCreateQuery**](SuppressionCreateJobCreateQuery.md) | Suppresses one or more profiles from receiving marketing. Currently, supports email only. If a profile is not found with the given email, one will be created and immediately suppressed. | 

### Return type

 (empty response body)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsubscribeProfiles

> UnsubscribeProfiles(ctx).Revision(revision).SubscriptionDeleteJobCreateQuery(subscriptionDeleteJobCreateQuery).Execute()

Unsubscribe Profiles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	subscriptionDeleteJobCreateQuery := *openapiclient.NewSubscriptionDeleteJobCreateQuery(*openapiclient.NewSubscriptionDeleteJobCreateQueryResourceObject(openapiclient.ProfileSubscriptionBulkDeleteJobEnum("profile-subscription-bulk-delete-job"), *openapiclient.NewSubscriptionDeleteJobCreateQueryResourceObjectAttributes(*openapiclient.NewSubscriptionDeleteJobCreateQueryResourceObjectAttributesProfiles([]openapiclient.ProfileSubscriptionDeleteQueryResourceObject{*openapiclient.NewProfileSubscriptionDeleteQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewProfileSubscriptionDeleteQueryResourceObjectAttributes())})))) // SubscriptionDeleteJobCreateQuery | Unsubscribes one or more profiles from marketing. Currently, supports email and SMS only. All profiles will be removed from the provided list. Either email or phone number is required. If a profile cannot be found matching the given identifier(s), a new profile will be created and then unsubscribed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProfilesAPI.UnsubscribeProfiles(context.Background()).Revision(revision).SubscriptionDeleteJobCreateQuery(subscriptionDeleteJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.UnsubscribeProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **subscriptionDeleteJobCreateQuery** | [**SubscriptionDeleteJobCreateQuery**](SubscriptionDeleteJobCreateQuery.md) | Unsubscribes one or more profiles from marketing. Currently, supports email and SMS only. All profiles will be removed from the provided list. Either email or phone number is required. If a profile cannot be found matching the given identifier(s), a new profile will be created and then unsubscribed. | 

### Return type

 (empty response body)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsuppressProfiles

> UnsuppressProfiles(ctx).Revision(revision).SuppressionDeleteJobCreateQuery(suppressionDeleteJobCreateQuery).Execute()

Unsuppress Profiles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	suppressionDeleteJobCreateQuery := *openapiclient.NewSuppressionDeleteJobCreateQuery(*openapiclient.NewSuppressionDeleteJobCreateQueryResourceObject(openapiclient.ProfileSuppressionBulkDeleteJobEnum("profile-suppression-bulk-delete-job"), *openapiclient.NewSuppressionDeleteJobCreateQueryResourceObjectAttributes(*openapiclient.NewSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles([]openapiclient.ProfileSuppressionDeleteQueryResourceObject{*openapiclient.NewProfileSuppressionDeleteQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewProfileSuppressionDeleteQueryResourceObjectAttributes())})))) // SuppressionDeleteJobCreateQuery | Unsuppresses one or more profiles from receiving marketing. Currently, supports email only. If a profile is not found with the given email, no action will be taken.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProfilesAPI.UnsuppressProfiles(context.Background()).Revision(revision).SuppressionDeleteJobCreateQuery(suppressionDeleteJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.UnsuppressProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnsuppressProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **suppressionDeleteJobCreateQuery** | [**SuppressionDeleteJobCreateQuery**](SuppressionDeleteJobCreateQuery.md) | Unsuppresses one or more profiles from receiving marketing. Currently, supports email only. If a profile is not found with the given email, no action will be taken. | 

### Return type

 (empty response body)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> PatchProfileResponse UpdateProfile(ctx, id).Revision(revision).ProfilePartialUpdateQuery(profilePartialUpdateQuery).Execute()

Update Profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	id := "01GDDKASAP8TKDDA2GRZDSVP4H" // string | Primary key that uniquely identifies this profile. Generated by Klaviyo.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	profilePartialUpdateQuery := *openapiclient.NewProfilePartialUpdateQuery(*openapiclient.NewProfilePartialUpdateQueryResourceObject(openapiclient.ProfileEnum("profile"), "01GDDKASAP8TKDDA2GRZDSVP4H", *openapiclient.NewProfilePartialUpdateQueryResourceObjectAttributes())) // ProfilePartialUpdateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.UpdateProfile(context.Background(), id).Revision(revision).ProfilePartialUpdateQuery(profilePartialUpdateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.UpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProfile`: PatchProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.UpdateProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Primary key that uniquely identifies this profile. Generated by Klaviyo. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **profilePartialUpdateQuery** | [**ProfilePartialUpdateQuery**](ProfilePartialUpdateQuery.md) |  | 

### Return type

[**PatchProfileResponse**](PatchProfileResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

