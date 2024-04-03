# \CouponsAPI

All URIs are relative to *https://a.klaviyo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCoupon**](CouponsAPI.md#CreateCoupon) | **Post** /api/coupons/ | Create Coupon
[**CreateCouponCode**](CouponsAPI.md#CreateCouponCode) | **Post** /api/coupon-codes/ | Create Coupon Code
[**DeleteCoupon**](CouponsAPI.md#DeleteCoupon) | **Delete** /api/coupons/{id}/ | Delete Coupon
[**DeleteCouponCode**](CouponsAPI.md#DeleteCouponCode) | **Delete** /api/coupon-codes/{id}/ | Delete Coupon Code
[**GetCoupon**](CouponsAPI.md#GetCoupon) | **Get** /api/coupons/{id}/ | Get Coupon
[**GetCouponCode**](CouponsAPI.md#GetCouponCode) | **Get** /api/coupon-codes/{id}/ | Get Coupon Code
[**GetCouponCodeBulkCreateJob**](CouponsAPI.md#GetCouponCodeBulkCreateJob) | **Get** /api/coupon-code-bulk-create-jobs/{job_id}/ | Get Coupon Code Bulk Create Job
[**GetCouponCodeBulkCreateJobs**](CouponsAPI.md#GetCouponCodeBulkCreateJobs) | **Get** /api/coupon-code-bulk-create-jobs/ | Get Coupon Code Bulk Create Jobs
[**GetCouponCodeRelationshipsCoupon**](CouponsAPI.md#GetCouponCodeRelationshipsCoupon) | **Get** /api/coupons/{id}/relationships/coupon-codes/ | Get Coupon Code Relationships Coupon
[**GetCouponCodes**](CouponsAPI.md#GetCouponCodes) | **Get** /api/coupon-codes/ | Get Coupon Codes
[**GetCouponCodesForCoupon**](CouponsAPI.md#GetCouponCodesForCoupon) | **Get** /api/coupons/{id}/coupon-codes/ | Get Coupon Codes For Coupon
[**GetCouponForCouponCode**](CouponsAPI.md#GetCouponForCouponCode) | **Get** /api/coupon-codes/{id}/coupon/ | Get Coupon For Coupon Code
[**GetCouponRelationshipsCouponCodes**](CouponsAPI.md#GetCouponRelationshipsCouponCodes) | **Get** /api/coupon-codes/{id}/relationships/coupon/ | Get Coupon Relationships Coupon Codes
[**GetCoupons**](CouponsAPI.md#GetCoupons) | **Get** /api/coupons/ | Get Coupons
[**SpawnCouponCodeBulkCreateJob**](CouponsAPI.md#SpawnCouponCodeBulkCreateJob) | **Post** /api/coupon-code-bulk-create-jobs/ | Spawn Coupon Code Bulk Create Job
[**UpdateCoupon**](CouponsAPI.md#UpdateCoupon) | **Patch** /api/coupons/{id}/ | Update Coupon
[**UpdateCouponCode**](CouponsAPI.md#UpdateCouponCode) | **Patch** /api/coupon-codes/{id}/ | Update Coupon Code



## CreateCoupon

> PostCouponResponse CreateCoupon(ctx).Revision(revision).CouponCreateQuery(couponCreateQuery).Execute()

Create Coupon



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
	couponCreateQuery := *openapiclient.NewCouponCreateQuery(*openapiclient.NewCouponCreateQueryResourceObject(openapiclient.CouponEnum("coupon"), *openapiclient.NewCouponResponseObjectResourceAttributes("10OFF"))) // CouponCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.CreateCoupon(context.Background()).Revision(revision).CouponCreateQuery(couponCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.CreateCoupon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCoupon`: PostCouponResponse
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.CreateCoupon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **couponCreateQuery** | [**CouponCreateQuery**](CouponCreateQuery.md) |  | 

### Return type

[**PostCouponResponse**](PostCouponResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponCode

> PostCouponCodeResponse CreateCouponCode(ctx).Revision(revision).CouponCodeCreateQuery(couponCodeCreateQuery).Execute()

Create Coupon Code



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
	couponCodeCreateQuery := *openapiclient.NewCouponCodeCreateQuery(*openapiclient.NewCouponCodeCreateQueryResourceObject(openapiclient.CouponCodeEnum("coupon-code"), *openapiclient.NewCouponCodeCreateQueryResourceObjectAttributes("ASD325FHK324UJDOI2M3JNES99"), *openapiclient.NewCouponCodeCreateQueryResourceObjectRelationships(*openapiclient.NewCouponCodeCreateQueryResourceObjectRelationshipsCoupon(*openapiclient.NewGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponData(openapiclient.CouponEnum("coupon"), "10OFF"))))) // CouponCodeCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.CreateCouponCode(context.Background()).Revision(revision).CouponCodeCreateQuery(couponCodeCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.CreateCouponCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCouponCode`: PostCouponCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.CreateCouponCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **couponCodeCreateQuery** | [**CouponCodeCreateQuery**](CouponCodeCreateQuery.md) |  | 

### Return type

[**PostCouponCodeResponse**](PostCouponCodeResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCoupon

> DeleteCoupon(ctx, id).Revision(revision).Execute()

Delete Coupon



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
	id := "10OFF" // string | The internal id of a Coupon is equivalent to its external id stored within an integration.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CouponsAPI.DeleteCoupon(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.DeleteCoupon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The internal id of a Coupon is equivalent to its external id stored within an integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

 (empty response body)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCouponCode

> DeleteCouponCode(ctx, id).Revision(revision).Execute()

Delete Coupon Code



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
	id := "10OFF-ASD325FHK324UJDOI2M3JNES99" // string | The id of a coupon code is a combination of its unique code and the id of the coupon it is associated with.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CouponsAPI.DeleteCouponCode(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.DeleteCouponCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of a coupon code is a combination of its unique code and the id of the coupon it is associated with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

 (empty response body)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCoupon

> GetCouponResponse GetCoupon(ctx, id).Revision(revision).FieldsCoupon(fieldsCoupon).Execute()

Get Coupon



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
	id := "10OFF" // string | The internal id of a Coupon is equivalent to its external id stored within an integration.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsCoupon := []string{"FieldsCoupon_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.GetCoupon(context.Background(), id).Revision(revision).FieldsCoupon(fieldsCoupon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.GetCoupon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCoupon`: GetCouponResponse
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.GetCoupon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The internal id of a Coupon is equivalent to its external id stored within an integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCoupon** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetCouponResponse**](GetCouponResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponCode

> GetCouponCodeResponseCompoundDocument GetCouponCode(ctx, id).Revision(revision).FieldsCouponCode(fieldsCouponCode).FieldsCoupon(fieldsCoupon).Include(include).Execute()

Get Coupon Code



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
	id := "10OFF-ASD325FHK324UJDOI2M3JNES99" // string | The id of a coupon code is a combination of its unique code and the id of the coupon it is associated with.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsCouponCode := []string{"FieldsCouponCode_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCoupon := []string{"FieldsCoupon_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.GetCouponCode(context.Background(), id).Revision(revision).FieldsCouponCode(fieldsCouponCode).FieldsCoupon(fieldsCoupon).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.GetCouponCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCouponCode`: GetCouponCodeResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.GetCouponCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of a coupon code is a combination of its unique code and the id of the coupon it is associated with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCouponCode** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCoupon** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetCouponCodeResponseCompoundDocument**](GetCouponCodeResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponCodeBulkCreateJob

> GetCouponCodeCreateJobResponseCompoundDocument GetCouponCodeBulkCreateJob(ctx, jobId).Revision(revision).FieldsCouponCodeBulkCreateJob(fieldsCouponCodeBulkCreateJob).FieldsCouponCode(fieldsCouponCode).Include(include).Execute()

Get Coupon Code Bulk Create Job



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
	fieldsCouponCodeBulkCreateJob := []string{"FieldsCouponCodeBulkCreateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCouponCode := []string{"FieldsCouponCode_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.GetCouponCodeBulkCreateJob(context.Background(), jobId).Revision(revision).FieldsCouponCodeBulkCreateJob(fieldsCouponCodeBulkCreateJob).FieldsCouponCode(fieldsCouponCode).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.GetCouponCodeBulkCreateJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCouponCodeBulkCreateJob`: GetCouponCodeCreateJobResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.GetCouponCodeBulkCreateJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | ID of the job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponCodeBulkCreateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCouponCodeBulkCreateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCouponCode** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetCouponCodeCreateJobResponseCompoundDocument**](GetCouponCodeCreateJobResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponCodeBulkCreateJobs

> GetCouponCodeCreateJobResponseCollectionCompoundDocument GetCouponCodeBulkCreateJobs(ctx).Revision(revision).FieldsCouponCodeBulkCreateJob(fieldsCouponCodeBulkCreateJob).Filter(filter).PageCursor(pageCursor).Execute()

Get Coupon Code Bulk Create Jobs



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
	fieldsCouponCodeBulkCreateJob := []string{"FieldsCouponCodeBulkCreateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`status`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.GetCouponCodeBulkCreateJobs(context.Background()).Revision(revision).FieldsCouponCodeBulkCreateJob(fieldsCouponCodeBulkCreateJob).Filter(filter).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.GetCouponCodeBulkCreateJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCouponCodeBulkCreateJobs`: GetCouponCodeCreateJobResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.GetCouponCodeBulkCreateJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponCodeBulkCreateJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCouponCodeBulkCreateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;status&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCouponCodeCreateJobResponseCollectionCompoundDocument**](GetCouponCodeCreateJobResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponCodeRelationshipsCoupon

> GetCouponRelationshipCouponCodesListResponseCollection GetCouponCodeRelationshipsCoupon(ctx, id).Revision(revision).PageCursor(pageCursor).Execute()

Get Coupon Code Relationships Coupon



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.GetCouponCodeRelationshipsCoupon(context.Background(), id).Revision(revision).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.GetCouponCodeRelationshipsCoupon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCouponCodeRelationshipsCoupon`: GetCouponRelationshipCouponCodesListResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.GetCouponCodeRelationshipsCoupon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponCodeRelationshipsCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCouponRelationshipCouponCodesListResponseCollection**](GetCouponRelationshipCouponCodesListResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponCodes

> GetCouponCodeResponseCollectionCompoundDocument GetCouponCodes(ctx).Revision(revision).FieldsCouponCode(fieldsCouponCode).FieldsCoupon(fieldsCoupon).Filter(filter).Include(include).PageCursor(pageCursor).Execute()

Get Coupon Codes



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
	fieldsCouponCode := []string{"FieldsCouponCode_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCoupon := []string{"FieldsCoupon_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`expires_at`: `greater-or-equal`, `greater-than`, `less-or-equal`, `less-than`<br>`status`: `equals`<br>`coupon.id`: `any`, `equals`<br>`profile.id`: `any`, `equals` (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.GetCouponCodes(context.Background()).Revision(revision).FieldsCouponCode(fieldsCouponCode).FieldsCoupon(fieldsCoupon).Filter(filter).Include(include).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.GetCouponCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCouponCodes`: GetCouponCodeResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.GetCouponCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCouponCode** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCoupon** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;expires_at&#x60;: &#x60;greater-or-equal&#x60;, &#x60;greater-than&#x60;, &#x60;less-or-equal&#x60;, &#x60;less-than&#x60;&lt;br&gt;&#x60;status&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;coupon.id&#x60;: &#x60;any&#x60;, &#x60;equals&#x60;&lt;br&gt;&#x60;profile.id&#x60;: &#x60;any&#x60;, &#x60;equals&#x60; | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCouponCodeResponseCollectionCompoundDocument**](GetCouponCodeResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponCodesForCoupon

> GetCouponCodeResponseCollection GetCouponCodesForCoupon(ctx, id).Revision(revision).FieldsCouponCode(fieldsCouponCode).Filter(filter).PageCursor(pageCursor).Execute()

Get Coupon Codes For Coupon



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
	fieldsCouponCode := []string{"FieldsCouponCode_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`expires_at`: `greater-or-equal`, `greater-than`, `less-or-equal`, `less-than`<br>`status`: `equals`<br>`coupon.id`: `any`, `equals`<br>`profile.id`: `any`, `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.GetCouponCodesForCoupon(context.Background(), id).Revision(revision).FieldsCouponCode(fieldsCouponCode).Filter(filter).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.GetCouponCodesForCoupon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCouponCodesForCoupon`: GetCouponCodeResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.GetCouponCodesForCoupon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponCodesForCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCouponCode** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;expires_at&#x60;: &#x60;greater-or-equal&#x60;, &#x60;greater-than&#x60;, &#x60;less-or-equal&#x60;, &#x60;less-than&#x60;&lt;br&gt;&#x60;status&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;coupon.id&#x60;: &#x60;any&#x60;, &#x60;equals&#x60;&lt;br&gt;&#x60;profile.id&#x60;: &#x60;any&#x60;, &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCouponCodeResponseCollection**](GetCouponCodeResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponForCouponCode

> GetCouponResponseCollection GetCouponForCouponCode(ctx, id).Revision(revision).FieldsCoupon(fieldsCoupon).Execute()

Get Coupon For Coupon Code



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
	fieldsCoupon := []string{"FieldsCoupon_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.GetCouponForCouponCode(context.Background(), id).Revision(revision).FieldsCoupon(fieldsCoupon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.GetCouponForCouponCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCouponForCouponCode`: GetCouponResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.GetCouponForCouponCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponForCouponCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCoupon** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetCouponResponseCollection**](GetCouponResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponRelationshipsCouponCodes

> GetCouponCodeRelationshipCouponResponse GetCouponRelationshipsCouponCodes(ctx, id).Revision(revision).Execute()

Get Coupon Relationships Coupon Codes



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
	resp, r, err := apiClient.CouponsAPI.GetCouponRelationshipsCouponCodes(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.GetCouponRelationshipsCouponCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCouponRelationshipsCouponCodes`: GetCouponCodeRelationshipCouponResponse
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.GetCouponRelationshipsCouponCodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponRelationshipsCouponCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetCouponCodeRelationshipCouponResponse**](GetCouponCodeRelationshipCouponResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCoupons

> GetCouponResponseCollection GetCoupons(ctx).Revision(revision).FieldsCoupon(fieldsCoupon).PageCursor(pageCursor).Execute()

Get Coupons



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
	fieldsCoupon := []string{"FieldsCoupon_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.GetCoupons(context.Background()).Revision(revision).FieldsCoupon(fieldsCoupon).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.GetCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCoupons`: GetCouponResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.GetCoupons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCoupon** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCouponResponseCollection**](GetCouponResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpawnCouponCodeBulkCreateJob

> PostCouponCodeCreateJobResponse SpawnCouponCodeBulkCreateJob(ctx).Revision(revision).CouponCodeCreateJobCreateQuery(couponCodeCreateJobCreateQuery).Execute()

Spawn Coupon Code Bulk Create Job



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
	couponCodeCreateJobCreateQuery := *openapiclient.NewCouponCodeCreateJobCreateQuery(*openapiclient.NewCouponCodeCreateJobCreateQueryResourceObject(openapiclient.CouponCodeBulkCreateJobEnum("coupon-code-bulk-create-job"), *openapiclient.NewCouponCodeCreateJobCreateQueryResourceObjectAttributes(*openapiclient.NewCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes([]openapiclient.CouponCodeCreateQueryResourceObject{*openapiclient.NewCouponCodeCreateQueryResourceObject(openapiclient.CouponCodeEnum("coupon-code"), *openapiclient.NewCouponCodeCreateQueryResourceObjectAttributes("ASD325FHK324UJDOI2M3JNES99"), *openapiclient.NewCouponCodeCreateQueryResourceObjectRelationships(*openapiclient.NewCouponCodeCreateQueryResourceObjectRelationshipsCoupon(*openapiclient.NewGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponData(openapiclient.CouponEnum("coupon"), "10OFF"))))})))) // CouponCodeCreateJobCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.SpawnCouponCodeBulkCreateJob(context.Background()).Revision(revision).CouponCodeCreateJobCreateQuery(couponCodeCreateJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.SpawnCouponCodeBulkCreateJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpawnCouponCodeBulkCreateJob`: PostCouponCodeCreateJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.SpawnCouponCodeBulkCreateJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpawnCouponCodeBulkCreateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **couponCodeCreateJobCreateQuery** | [**CouponCodeCreateJobCreateQuery**](CouponCodeCreateJobCreateQuery.md) |  | 

### Return type

[**PostCouponCodeCreateJobResponse**](PostCouponCodeCreateJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCoupon

> PatchCouponResponse UpdateCoupon(ctx, id).Revision(revision).CouponUpdateQuery(couponUpdateQuery).Execute()

Update Coupon



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
	id := "10OFF" // string | The internal id of a Coupon is equivalent to its external id stored within an integration.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	couponUpdateQuery := *openapiclient.NewCouponUpdateQuery(*openapiclient.NewCouponUpdateQueryResourceObject(openapiclient.CouponEnum("coupon"), "10OFF", *openapiclient.NewCouponUpdateQueryResourceObjectAttributes())) // CouponUpdateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.UpdateCoupon(context.Background(), id).Revision(revision).CouponUpdateQuery(couponUpdateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.UpdateCoupon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCoupon`: PatchCouponResponse
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.UpdateCoupon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The internal id of a Coupon is equivalent to its external id stored within an integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **couponUpdateQuery** | [**CouponUpdateQuery**](CouponUpdateQuery.md) |  | 

### Return type

[**PatchCouponResponse**](PatchCouponResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCouponCode

> PatchCouponCodeResponse UpdateCouponCode(ctx, id).Revision(revision).CouponCodeUpdateQuery(couponCodeUpdateQuery).Execute()

Update Coupon Code



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
	id := "10OFF-ASD325FHK324UJDOI2M3JNES99" // string | The id of a coupon code is a combination of its unique code and the id of the coupon it is associated with.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	couponCodeUpdateQuery := *openapiclient.NewCouponCodeUpdateQuery(*openapiclient.NewCouponCodeUpdateQueryResourceObject(openapiclient.CouponCodeEnum("coupon-code"), "10OFF-ASD325FHK324UJDOI2M3JNES99", *openapiclient.NewCouponCodeUpdateQueryResourceObjectAttributes())) // CouponCodeUpdateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.UpdateCouponCode(context.Background(), id).Revision(revision).CouponCodeUpdateQuery(couponCodeUpdateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.UpdateCouponCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCouponCode`: PatchCouponCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.UpdateCouponCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of a coupon code is a combination of its unique code and the id of the coupon it is associated with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCouponCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **couponCodeUpdateQuery** | [**CouponCodeUpdateQuery**](CouponCodeUpdateQuery.md) |  | 

### Return type

[**PatchCouponCodeResponse**](PatchCouponCodeResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

