# \CatalogsAPI

All URIs are relative to *https://a.klaviyo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackInStockSubscription**](CatalogsAPI.md#CreateBackInStockSubscription) | **Post** /api/back-in-stock-subscriptions/ | Create Back In Stock Subscription
[**CreateCatalogCategory**](CatalogsAPI.md#CreateCatalogCategory) | **Post** /api/catalog-categories/ | Create Catalog Category
[**CreateCatalogCategoryRelationshipsItems**](CatalogsAPI.md#CreateCatalogCategoryRelationshipsItems) | **Post** /api/catalog-categories/{id}/relationships/items/ | Create Catalog Category Relationships Items
[**CreateCatalogItem**](CatalogsAPI.md#CreateCatalogItem) | **Post** /api/catalog-items/ | Create Catalog Item
[**CreateCatalogItemRelationshipsCategories**](CatalogsAPI.md#CreateCatalogItemRelationshipsCategories) | **Post** /api/catalog-items/{id}/relationships/categories/ | Create Catalog Item Relationships Categories
[**CreateCatalogVariant**](CatalogsAPI.md#CreateCatalogVariant) | **Post** /api/catalog-variants/ | Create Catalog Variant
[**DeleteCatalogCategory**](CatalogsAPI.md#DeleteCatalogCategory) | **Delete** /api/catalog-categories/{id}/ | Delete Catalog Category
[**DeleteCatalogCategoryRelationshipsItems**](CatalogsAPI.md#DeleteCatalogCategoryRelationshipsItems) | **Delete** /api/catalog-categories/{id}/relationships/items/ | Delete Catalog Category Relationships Items
[**DeleteCatalogItem**](CatalogsAPI.md#DeleteCatalogItem) | **Delete** /api/catalog-items/{id}/ | Delete Catalog Item
[**DeleteCatalogItemRelationshipsCategories**](CatalogsAPI.md#DeleteCatalogItemRelationshipsCategories) | **Delete** /api/catalog-items/{id}/relationships/categories/ | Delete Catalog Item Relationships Categories
[**DeleteCatalogVariant**](CatalogsAPI.md#DeleteCatalogVariant) | **Delete** /api/catalog-variants/{id}/ | Delete Catalog Variant
[**GetCatalogCategories**](CatalogsAPI.md#GetCatalogCategories) | **Get** /api/catalog-categories/ | Get Catalog Categories
[**GetCatalogCategory**](CatalogsAPI.md#GetCatalogCategory) | **Get** /api/catalog-categories/{id}/ | Get Catalog Category
[**GetCatalogCategoryItems**](CatalogsAPI.md#GetCatalogCategoryItems) | **Get** /api/catalog-categories/{id}/items/ | Get Catalog Category Items
[**GetCatalogCategoryRelationshipsItems**](CatalogsAPI.md#GetCatalogCategoryRelationshipsItems) | **Get** /api/catalog-categories/{id}/relationships/items/ | Get Catalog Category Relationships Items
[**GetCatalogItem**](CatalogsAPI.md#GetCatalogItem) | **Get** /api/catalog-items/{id}/ | Get Catalog Item
[**GetCatalogItemCategories**](CatalogsAPI.md#GetCatalogItemCategories) | **Get** /api/catalog-items/{id}/categories/ | Get Catalog Item Categories
[**GetCatalogItemRelationshipsCategories**](CatalogsAPI.md#GetCatalogItemRelationshipsCategories) | **Get** /api/catalog-items/{id}/relationships/categories/ | Get Catalog Item Relationships Categories
[**GetCatalogItemVariants**](CatalogsAPI.md#GetCatalogItemVariants) | **Get** /api/catalog-items/{id}/variants/ | Get Catalog Item Variants
[**GetCatalogItems**](CatalogsAPI.md#GetCatalogItems) | **Get** /api/catalog-items/ | Get Catalog Items
[**GetCatalogVariant**](CatalogsAPI.md#GetCatalogVariant) | **Get** /api/catalog-variants/{id}/ | Get Catalog Variant
[**GetCatalogVariants**](CatalogsAPI.md#GetCatalogVariants) | **Get** /api/catalog-variants/ | Get Catalog Variants
[**GetCreateCategoriesJob**](CatalogsAPI.md#GetCreateCategoriesJob) | **Get** /api/catalog-category-bulk-create-jobs/{job_id}/ | Get Create Categories Job
[**GetCreateCategoriesJobs**](CatalogsAPI.md#GetCreateCategoriesJobs) | **Get** /api/catalog-category-bulk-create-jobs/ | Get Create Categories Jobs
[**GetCreateItemsJob**](CatalogsAPI.md#GetCreateItemsJob) | **Get** /api/catalog-item-bulk-create-jobs/{job_id}/ | Get Create Items Job
[**GetCreateItemsJobs**](CatalogsAPI.md#GetCreateItemsJobs) | **Get** /api/catalog-item-bulk-create-jobs/ | Get Create Items Jobs
[**GetCreateVariantsJob**](CatalogsAPI.md#GetCreateVariantsJob) | **Get** /api/catalog-variant-bulk-create-jobs/{job_id}/ | Get Create Variants Job
[**GetCreateVariantsJobs**](CatalogsAPI.md#GetCreateVariantsJobs) | **Get** /api/catalog-variant-bulk-create-jobs/ | Get Create Variants Jobs
[**GetDeleteCategoriesJob**](CatalogsAPI.md#GetDeleteCategoriesJob) | **Get** /api/catalog-category-bulk-delete-jobs/{job_id}/ | Get Delete Categories Job
[**GetDeleteCategoriesJobs**](CatalogsAPI.md#GetDeleteCategoriesJobs) | **Get** /api/catalog-category-bulk-delete-jobs/ | Get Delete Categories Jobs
[**GetDeleteItemsJob**](CatalogsAPI.md#GetDeleteItemsJob) | **Get** /api/catalog-item-bulk-delete-jobs/{job_id}/ | Get Delete Items Job
[**GetDeleteItemsJobs**](CatalogsAPI.md#GetDeleteItemsJobs) | **Get** /api/catalog-item-bulk-delete-jobs/ | Get Delete Items Jobs
[**GetDeleteVariantsJob**](CatalogsAPI.md#GetDeleteVariantsJob) | **Get** /api/catalog-variant-bulk-delete-jobs/{job_id}/ | Get Delete Variants Job
[**GetDeleteVariantsJobs**](CatalogsAPI.md#GetDeleteVariantsJobs) | **Get** /api/catalog-variant-bulk-delete-jobs/ | Get Delete Variants Jobs
[**GetUpdateCategoriesJob**](CatalogsAPI.md#GetUpdateCategoriesJob) | **Get** /api/catalog-category-bulk-update-jobs/{job_id}/ | Get Update Categories Job
[**GetUpdateCategoriesJobs**](CatalogsAPI.md#GetUpdateCategoriesJobs) | **Get** /api/catalog-category-bulk-update-jobs/ | Get Update Categories Jobs
[**GetUpdateItemsJob**](CatalogsAPI.md#GetUpdateItemsJob) | **Get** /api/catalog-item-bulk-update-jobs/{job_id}/ | Get Update Items Job
[**GetUpdateItemsJobs**](CatalogsAPI.md#GetUpdateItemsJobs) | **Get** /api/catalog-item-bulk-update-jobs/ | Get Update Items Jobs
[**GetUpdateVariantsJob**](CatalogsAPI.md#GetUpdateVariantsJob) | **Get** /api/catalog-variant-bulk-update-jobs/{job_id}/ | Get Update Variants Job
[**GetUpdateVariantsJobs**](CatalogsAPI.md#GetUpdateVariantsJobs) | **Get** /api/catalog-variant-bulk-update-jobs/ | Get Update Variants Jobs
[**SpawnCreateCategoriesJob**](CatalogsAPI.md#SpawnCreateCategoriesJob) | **Post** /api/catalog-category-bulk-create-jobs/ | Spawn Create Categories Job
[**SpawnCreateItemsJob**](CatalogsAPI.md#SpawnCreateItemsJob) | **Post** /api/catalog-item-bulk-create-jobs/ | Spawn Create Items Job
[**SpawnCreateVariantsJob**](CatalogsAPI.md#SpawnCreateVariantsJob) | **Post** /api/catalog-variant-bulk-create-jobs/ | Spawn Create Variants Job
[**SpawnDeleteCategoriesJob**](CatalogsAPI.md#SpawnDeleteCategoriesJob) | **Post** /api/catalog-category-bulk-delete-jobs/ | Spawn Delete Categories Job
[**SpawnDeleteItemsJob**](CatalogsAPI.md#SpawnDeleteItemsJob) | **Post** /api/catalog-item-bulk-delete-jobs/ | Spawn Delete Items Job
[**SpawnDeleteVariantsJob**](CatalogsAPI.md#SpawnDeleteVariantsJob) | **Post** /api/catalog-variant-bulk-delete-jobs/ | Spawn Delete Variants Job
[**SpawnUpdateCategoriesJob**](CatalogsAPI.md#SpawnUpdateCategoriesJob) | **Post** /api/catalog-category-bulk-update-jobs/ | Spawn Update Categories Job
[**SpawnUpdateItemsJob**](CatalogsAPI.md#SpawnUpdateItemsJob) | **Post** /api/catalog-item-bulk-update-jobs/ | Spawn Update Items Job
[**SpawnUpdateVariantsJob**](CatalogsAPI.md#SpawnUpdateVariantsJob) | **Post** /api/catalog-variant-bulk-update-jobs/ | Spawn Update Variants Job
[**UpdateCatalogCategory**](CatalogsAPI.md#UpdateCatalogCategory) | **Patch** /api/catalog-categories/{id}/ | Update Catalog Category
[**UpdateCatalogCategoryRelationshipsItems**](CatalogsAPI.md#UpdateCatalogCategoryRelationshipsItems) | **Patch** /api/catalog-categories/{id}/relationships/items/ | Update Catalog Category Relationships Items
[**UpdateCatalogItem**](CatalogsAPI.md#UpdateCatalogItem) | **Patch** /api/catalog-items/{id}/ | Update Catalog Item
[**UpdateCatalogItemRelationshipsCategories**](CatalogsAPI.md#UpdateCatalogItemRelationshipsCategories) | **Patch** /api/catalog-items/{id}/relationships/categories/ | Update Catalog Item Relationships Categories
[**UpdateCatalogVariant**](CatalogsAPI.md#UpdateCatalogVariant) | **Patch** /api/catalog-variants/{id}/ | Update Catalog Variant



## CreateBackInStockSubscription

> CreateBackInStockSubscription(ctx).Revision(revision).ServerBISSubscriptionCreateQuery(serverBISSubscriptionCreateQuery).Execute()

Create Back In Stock Subscription



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
	serverBISSubscriptionCreateQuery := *openapiclient.NewServerBISSubscriptionCreateQuery(*openapiclient.NewServerBISSubscriptionCreateQueryResourceObject(openapiclient.BackInStockSubscriptionEnum("back-in-stock-subscription"), *openapiclient.NewServerBISSubscriptionCreateQueryResourceObjectAttributes([]string{"Channels_example"}), *openapiclient.NewServerBISSubscriptionCreateQueryResourceObjectRelationships(*openapiclient.NewServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariant(*openapiclient.NewServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData(openapiclient.CatalogVariantEnum("catalog-variant"), "$custom:::$default:::SAMPLE-DATA-ITEM-1-VARIANT-MEDIUM"))))) // ServerBISSubscriptionCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogsAPI.CreateBackInStockSubscription(context.Background()).Revision(revision).ServerBISSubscriptionCreateQuery(serverBISSubscriptionCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.CreateBackInStockSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackInStockSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **serverBISSubscriptionCreateQuery** | [**ServerBISSubscriptionCreateQuery**](ServerBISSubscriptionCreateQuery.md) |  | 

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


## CreateCatalogCategory

> PostCatalogCategoryResponse CreateCatalogCategory(ctx).Revision(revision).CatalogCategoryCreateQuery(catalogCategoryCreateQuery).Execute()

Create Catalog Category



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
	catalogCategoryCreateQuery := *openapiclient.NewCatalogCategoryCreateQuery(*openapiclient.NewCatalogCategoryCreateQueryResourceObject(openapiclient.CatalogCategoryEnum("catalog-category"), *openapiclient.NewCatalogCategoryCreateQueryResourceObjectAttributes("SAMPLE-DATA-CATEGORY-APPAREL", "Sample Data Category Apparel"))) // CatalogCategoryCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.CreateCatalogCategory(context.Background()).Revision(revision).CatalogCategoryCreateQuery(catalogCategoryCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.CreateCatalogCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCatalogCategory`: PostCatalogCategoryResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.CreateCatalogCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCatalogCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogCategoryCreateQuery** | [**CatalogCategoryCreateQuery**](CatalogCategoryCreateQuery.md) |  | 

### Return type

[**PostCatalogCategoryResponse**](PostCatalogCategoryResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCatalogCategoryRelationshipsItems

> CreateCatalogCategoryRelationshipsItems(ctx, id).Revision(revision).CatalogCategoryItemOp(catalogCategoryItemOp).Execute()

Create Catalog Category Relationships Items



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
	catalogCategoryItemOp := *openapiclient.NewCatalogCategoryItemOp([]openapiclient.GetCatalogCategoryItemListResponseCollectionDataInner{*openapiclient.NewGetCatalogCategoryItemListResponseCollectionDataInner(openapiclient.CatalogItemEnum("catalog-item"), "$custom:::$default:::SAMPLE-DATA-ITEM-1")}) // CatalogCategoryItemOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogsAPI.CreateCatalogCategoryRelationshipsItems(context.Background(), id).Revision(revision).CatalogCategoryItemOp(catalogCategoryItemOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.CreateCatalogCategoryRelationshipsItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCatalogCategoryRelationshipsItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogCategoryItemOp** | [**CatalogCategoryItemOp**](CatalogCategoryItemOp.md) |  | 

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


## CreateCatalogItem

> PostCatalogItemResponse CreateCatalogItem(ctx).Revision(revision).CatalogItemCreateQuery(catalogItemCreateQuery).Execute()

Create Catalog Item



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
	catalogItemCreateQuery := *openapiclient.NewCatalogItemCreateQuery(*openapiclient.NewCatalogItemCreateQueryResourceObject(openapiclient.CatalogItemEnum("catalog-item"), *openapiclient.NewCatalogItemCreateQueryResourceObjectAttributes("SAMPLE-DATA-ITEM-1", "Ocean Blue Shirt (Sample)", "Ocean blue cotton shirt with a narrow collar and buttons down the front and long sleeves. Comfortable fit and titled kaleidoscope patterns.", "https://via.placeholder.com/150"))) // CatalogItemCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.CreateCatalogItem(context.Background()).Revision(revision).CatalogItemCreateQuery(catalogItemCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.CreateCatalogItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCatalogItem`: PostCatalogItemResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.CreateCatalogItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCatalogItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogItemCreateQuery** | [**CatalogItemCreateQuery**](CatalogItemCreateQuery.md) |  | 

### Return type

[**PostCatalogItemResponse**](PostCatalogItemResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCatalogItemRelationshipsCategories

> CreateCatalogItemRelationshipsCategories(ctx, id).Revision(revision).CatalogItemCategoryOp(catalogItemCategoryOp).Execute()

Create Catalog Item Relationships Categories



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
	catalogItemCategoryOp := *openapiclient.NewCatalogItemCategoryOp([]openapiclient.GetCatalogItemCategoryListResponseCollectionDataInner{*openapiclient.NewGetCatalogItemCategoryListResponseCollectionDataInner(openapiclient.CatalogCategoryEnum("catalog-category"), "$custom:::$default:::SAMPLE-DATA-CATEGORY-APPAREL")}) // CatalogItemCategoryOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogsAPI.CreateCatalogItemRelationshipsCategories(context.Background(), id).Revision(revision).CatalogItemCategoryOp(catalogItemCategoryOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.CreateCatalogItemRelationshipsCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCatalogItemRelationshipsCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogItemCategoryOp** | [**CatalogItemCategoryOp**](CatalogItemCategoryOp.md) |  | 

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


## CreateCatalogVariant

> PostCatalogVariantResponse CreateCatalogVariant(ctx).Revision(revision).CatalogVariantCreateQuery(catalogVariantCreateQuery).Execute()

Create Catalog Variant



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
	catalogVariantCreateQuery := *openapiclient.NewCatalogVariantCreateQuery(*openapiclient.NewCatalogVariantCreateQueryResourceObject(openapiclient.CatalogVariantEnum("catalog-variant"), *openapiclient.NewCatalogVariantCreateQueryResourceObjectAttributes("SAMPLE-DATA-ITEM-1-VARIANT-MEDIUM", "Ocean Blue Shirt (Sample) Variant Medium", "Ocean blue cotton shirt with a narrow collar and buttons down the front and long sleeves. Comfortable fit and titled kaleidoscope patterns.", "OBS-MD", float32(25), float32(42), "https://via.placeholder.com/150"), *openapiclient.NewCatalogVariantCreateQueryResourceObjectRelationships(*openapiclient.NewCatalogVariantCreateQueryResourceObjectRelationshipsItem(*openapiclient.NewCatalogVariantCreateQueryResourceObjectRelationshipsItemData(openapiclient.CatalogItemEnum("catalog-item"), "$custom:::$default:::SAMPLE-DATA-ITEM-1"))))) // CatalogVariantCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.CreateCatalogVariant(context.Background()).Revision(revision).CatalogVariantCreateQuery(catalogVariantCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.CreateCatalogVariant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCatalogVariant`: PostCatalogVariantResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.CreateCatalogVariant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCatalogVariantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogVariantCreateQuery** | [**CatalogVariantCreateQuery**](CatalogVariantCreateQuery.md) |  | 

### Return type

[**PostCatalogVariantResponse**](PostCatalogVariantResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCatalogCategory

> DeleteCatalogCategory(ctx, id).Revision(revision).Execute()

Delete Catalog Category



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
	id := "$custom:::$default:::SAMPLE-DATA-CATEGORY-APPAREL" // string | The catalog category ID is a compound ID (string), with format: `{integration}:::{catalog}:::{external_id}`. Currently, the only supported integration type is `$custom`, and the only supported catalog is `$default`.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogsAPI.DeleteCatalogCategory(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.DeleteCatalogCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The catalog category ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCatalogCategoryRequest struct via the builder pattern


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


## DeleteCatalogCategoryRelationshipsItems

> DeleteCatalogCategoryRelationshipsItems(ctx, id).Revision(revision).CatalogCategoryItemOp(catalogCategoryItemOp).Execute()

Delete Catalog Category Relationships Items



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
	catalogCategoryItemOp := *openapiclient.NewCatalogCategoryItemOp([]openapiclient.GetCatalogCategoryItemListResponseCollectionDataInner{*openapiclient.NewGetCatalogCategoryItemListResponseCollectionDataInner(openapiclient.CatalogItemEnum("catalog-item"), "$custom:::$default:::SAMPLE-DATA-ITEM-1")}) // CatalogCategoryItemOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogsAPI.DeleteCatalogCategoryRelationshipsItems(context.Background(), id).Revision(revision).CatalogCategoryItemOp(catalogCategoryItemOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.DeleteCatalogCategoryRelationshipsItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCatalogCategoryRelationshipsItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogCategoryItemOp** | [**CatalogCategoryItemOp**](CatalogCategoryItemOp.md) |  | 

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


## DeleteCatalogItem

> DeleteCatalogItem(ctx, id).Revision(revision).Execute()

Delete Catalog Item



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
	id := "$custom:::$default:::SAMPLE-DATA-ITEM-1" // string | The catalog item ID is a compound ID (string), with format: `{integration}:::{catalog}:::{external_id}`. Currently, the only supported integration type is `$custom`, and the only supported catalog is `$default`.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogsAPI.DeleteCatalogItem(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.DeleteCatalogItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The catalog item ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCatalogItemRequest struct via the builder pattern


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


## DeleteCatalogItemRelationshipsCategories

> DeleteCatalogItemRelationshipsCategories(ctx, id).Revision(revision).CatalogItemCategoryOp(catalogItemCategoryOp).Execute()

Delete Catalog Item Relationships Categories



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
	catalogItemCategoryOp := *openapiclient.NewCatalogItemCategoryOp([]openapiclient.GetCatalogItemCategoryListResponseCollectionDataInner{*openapiclient.NewGetCatalogItemCategoryListResponseCollectionDataInner(openapiclient.CatalogCategoryEnum("catalog-category"), "$custom:::$default:::SAMPLE-DATA-CATEGORY-APPAREL")}) // CatalogItemCategoryOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogsAPI.DeleteCatalogItemRelationshipsCategories(context.Background(), id).Revision(revision).CatalogItemCategoryOp(catalogItemCategoryOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.DeleteCatalogItemRelationshipsCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCatalogItemRelationshipsCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogItemCategoryOp** | [**CatalogItemCategoryOp**](CatalogItemCategoryOp.md) |  | 

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


## DeleteCatalogVariant

> DeleteCatalogVariant(ctx, id).Revision(revision).Execute()

Delete Catalog Variant



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
	id := "$custom:::$default:::SAMPLE-DATA-ITEM-1-VARIANT-MEDIUM" // string | The catalog variant ID is a compound ID (string), with format: `{integration}:::{catalog}:::{external_id}`. Currently, the only supported integration type is `$custom`, and the only supported catalog is `$default`.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogsAPI.DeleteCatalogVariant(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.DeleteCatalogVariant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The catalog variant ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCatalogVariantRequest struct via the builder pattern


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


## GetCatalogCategories

> GetCatalogCategoryResponseCollection GetCatalogCategories(ctx).Revision(revision).FieldsCatalogCategory(fieldsCatalogCategory).Filter(filter).PageCursor(pageCursor).Sort(sort).Execute()

Get Catalog Categories



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
	fieldsCatalogCategory := []string{"FieldsCatalogCategory_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`ids`: `any`<br>`item.id`: `equals`<br>`name`: `contains` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCatalogCategories(context.Background()).Revision(revision).FieldsCatalogCategory(fieldsCatalogCategory).Filter(filter).PageCursor(pageCursor).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCatalogCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogCategories`: GetCatalogCategoryResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCatalogCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogCategory** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;ids&#x60;: &#x60;any&#x60;&lt;br&gt;&#x60;item.id&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;name&#x60;: &#x60;contains&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetCatalogCategoryResponseCollection**](GetCatalogCategoryResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogCategory

> GetCatalogCategoryResponse GetCatalogCategory(ctx, id).Revision(revision).FieldsCatalogCategory(fieldsCatalogCategory).Execute()

Get Catalog Category



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
	id := "$custom:::$default:::SAMPLE-DATA-CATEGORY-APPAREL" // string | The catalog category ID is a compound ID (string), with format: `{integration}:::{catalog}:::{external_id}`. Currently, the only supported integration type is `$custom`, and the only supported catalog is `$default`.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsCatalogCategory := []string{"FieldsCatalogCategory_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCatalogCategory(context.Background(), id).Revision(revision).FieldsCatalogCategory(fieldsCatalogCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCatalogCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogCategory`: GetCatalogCategoryResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCatalogCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The catalog category ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogCategory** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetCatalogCategoryResponse**](GetCatalogCategoryResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogCategoryItems

> GetCatalogItemResponseCollectionCompoundDocument GetCatalogCategoryItems(ctx, id).Revision(revision).FieldsCatalogItem(fieldsCatalogItem).FieldsCatalogVariant(fieldsCatalogVariant).Filter(filter).Include(include).PageCursor(pageCursor).Sort(sort).Execute()

Get Catalog Category Items



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
	fieldsCatalogItem := []string{"FieldsCatalogItem_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCatalogVariant := []string{"FieldsCatalogVariant_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`ids`: `any`<br>`category.id`: `equals`<br>`title`: `contains`<br>`published`: `equals` (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCatalogCategoryItems(context.Background(), id).Revision(revision).FieldsCatalogItem(fieldsCatalogItem).FieldsCatalogVariant(fieldsCatalogVariant).Filter(filter).Include(include).PageCursor(pageCursor).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCatalogCategoryItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogCategoryItems`: GetCatalogItemResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCatalogCategoryItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogCategoryItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogItem** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCatalogVariant** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;ids&#x60;: &#x60;any&#x60;&lt;br&gt;&#x60;category.id&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;title&#x60;: &#x60;contains&#x60;&lt;br&gt;&#x60;published&#x60;: &#x60;equals&#x60; | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetCatalogItemResponseCollectionCompoundDocument**](GetCatalogItemResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogCategoryRelationshipsItems

> GetCatalogCategoryItemListResponseCollection GetCatalogCategoryRelationshipsItems(ctx, id).Revision(revision).PageCursor(pageCursor).Execute()

Get Catalog Category Relationships Items



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
	resp, r, err := apiClient.CatalogsAPI.GetCatalogCategoryRelationshipsItems(context.Background(), id).Revision(revision).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCatalogCategoryRelationshipsItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogCategoryRelationshipsItems`: GetCatalogCategoryItemListResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCatalogCategoryRelationshipsItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogCategoryRelationshipsItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCatalogCategoryItemListResponseCollection**](GetCatalogCategoryItemListResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogItem

> GetCatalogItemResponseCompoundDocument GetCatalogItem(ctx, id).Revision(revision).FieldsCatalogItem(fieldsCatalogItem).FieldsCatalogVariant(fieldsCatalogVariant).Include(include).Execute()

Get Catalog Item



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
	id := "$custom:::$default:::SAMPLE-DATA-ITEM-1" // string | The catalog item ID is a compound ID (string), with format: `{integration}:::{catalog}:::{external_id}`. Currently, the only supported integration type is `$custom`, and the only supported catalog is `$default`.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsCatalogItem := []string{"FieldsCatalogItem_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCatalogVariant := []string{"FieldsCatalogVariant_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCatalogItem(context.Background(), id).Revision(revision).FieldsCatalogItem(fieldsCatalogItem).FieldsCatalogVariant(fieldsCatalogVariant).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCatalogItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogItem`: GetCatalogItemResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCatalogItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The catalog item ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogItem** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCatalogVariant** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetCatalogItemResponseCompoundDocument**](GetCatalogItemResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogItemCategories

> GetCatalogCategoryResponseCollection GetCatalogItemCategories(ctx, id).Revision(revision).FieldsCatalogCategory(fieldsCatalogCategory).Filter(filter).PageCursor(pageCursor).Sort(sort).Execute()

Get Catalog Item Categories



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
	fieldsCatalogCategory := []string{"FieldsCatalogCategory_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`ids`: `any`<br>`item.id`: `equals`<br>`name`: `contains` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCatalogItemCategories(context.Background(), id).Revision(revision).FieldsCatalogCategory(fieldsCatalogCategory).Filter(filter).PageCursor(pageCursor).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCatalogItemCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogItemCategories`: GetCatalogCategoryResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCatalogItemCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogItemCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogCategory** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;ids&#x60;: &#x60;any&#x60;&lt;br&gt;&#x60;item.id&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;name&#x60;: &#x60;contains&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetCatalogCategoryResponseCollection**](GetCatalogCategoryResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogItemRelationshipsCategories

> GetCatalogItemCategoryListResponseCollection GetCatalogItemRelationshipsCategories(ctx, id).Revision(revision).PageCursor(pageCursor).Execute()

Get Catalog Item Relationships Categories



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
	resp, r, err := apiClient.CatalogsAPI.GetCatalogItemRelationshipsCategories(context.Background(), id).Revision(revision).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCatalogItemRelationshipsCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogItemRelationshipsCategories`: GetCatalogItemCategoryListResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCatalogItemRelationshipsCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogItemRelationshipsCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCatalogItemCategoryListResponseCollection**](GetCatalogItemCategoryListResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogItemVariants

> GetCatalogVariantResponseCollection GetCatalogItemVariants(ctx, id).Revision(revision).FieldsCatalogVariant(fieldsCatalogVariant).Filter(filter).PageCursor(pageCursor).Sort(sort).Execute()

Get Catalog Item Variants



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
	fieldsCatalogVariant := []string{"FieldsCatalogVariant_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`ids`: `any`<br>`item.id`: `equals`<br>`sku`: `equals`<br>`title`: `contains`<br>`published`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCatalogItemVariants(context.Background(), id).Revision(revision).FieldsCatalogVariant(fieldsCatalogVariant).Filter(filter).PageCursor(pageCursor).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCatalogItemVariants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogItemVariants`: GetCatalogVariantResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCatalogItemVariants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogItemVariantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogVariant** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;ids&#x60;: &#x60;any&#x60;&lt;br&gt;&#x60;item.id&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;sku&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;title&#x60;: &#x60;contains&#x60;&lt;br&gt;&#x60;published&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetCatalogVariantResponseCollection**](GetCatalogVariantResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogItems

> GetCatalogItemResponseCollectionCompoundDocument GetCatalogItems(ctx).Revision(revision).FieldsCatalogItem(fieldsCatalogItem).FieldsCatalogVariant(fieldsCatalogVariant).Filter(filter).Include(include).PageCursor(pageCursor).Sort(sort).Execute()

Get Catalog Items



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
	fieldsCatalogItem := []string{"FieldsCatalogItem_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCatalogVariant := []string{"FieldsCatalogVariant_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`ids`: `any`<br>`category.id`: `equals`<br>`title`: `contains`<br>`published`: `equals` (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCatalogItems(context.Background()).Revision(revision).FieldsCatalogItem(fieldsCatalogItem).FieldsCatalogVariant(fieldsCatalogVariant).Filter(filter).Include(include).PageCursor(pageCursor).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCatalogItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogItems`: GetCatalogItemResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCatalogItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogItem** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCatalogVariant** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;ids&#x60;: &#x60;any&#x60;&lt;br&gt;&#x60;category.id&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;title&#x60;: &#x60;contains&#x60;&lt;br&gt;&#x60;published&#x60;: &#x60;equals&#x60; | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetCatalogItemResponseCollectionCompoundDocument**](GetCatalogItemResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogVariant

> GetCatalogVariantResponse GetCatalogVariant(ctx, id).Revision(revision).FieldsCatalogVariant(fieldsCatalogVariant).Execute()

Get Catalog Variant



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
	id := "$custom:::$default:::SAMPLE-DATA-ITEM-1-VARIANT-MEDIUM" // string | The catalog variant ID is a compound ID (string), with format: `{integration}:::{catalog}:::{external_id}`. Currently, the only supported integration type is `$custom`, and the only supported catalog is `$default`.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsCatalogVariant := []string{"FieldsCatalogVariant_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCatalogVariant(context.Background(), id).Revision(revision).FieldsCatalogVariant(fieldsCatalogVariant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCatalogVariant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogVariant`: GetCatalogVariantResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCatalogVariant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The catalog variant ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogVariantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogVariant** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetCatalogVariantResponse**](GetCatalogVariantResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogVariants

> GetCatalogVariantResponseCollection GetCatalogVariants(ctx).Revision(revision).FieldsCatalogVariant(fieldsCatalogVariant).Filter(filter).PageCursor(pageCursor).Sort(sort).Execute()

Get Catalog Variants



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
	fieldsCatalogVariant := []string{"FieldsCatalogVariant_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`ids`: `any`<br>`item.id`: `equals`<br>`sku`: `equals`<br>`title`: `contains`<br>`published`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCatalogVariants(context.Background()).Revision(revision).FieldsCatalogVariant(fieldsCatalogVariant).Filter(filter).PageCursor(pageCursor).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCatalogVariants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogVariants`: GetCatalogVariantResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCatalogVariants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogVariantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogVariant** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;ids&#x60;: &#x60;any&#x60;&lt;br&gt;&#x60;item.id&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;sku&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;title&#x60;: &#x60;contains&#x60;&lt;br&gt;&#x60;published&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetCatalogVariantResponseCollection**](GetCatalogVariantResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreateCategoriesJob

> GetCatalogCategoryCreateJobResponseCompoundDocument GetCreateCategoriesJob(ctx, jobId).Revision(revision).FieldsCatalogCategoryBulkCreateJob(fieldsCatalogCategoryBulkCreateJob).FieldsCatalogCategory(fieldsCatalogCategory).Include(include).Execute()

Get Create Categories Job



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
	fieldsCatalogCategoryBulkCreateJob := []string{"FieldsCatalogCategoryBulkCreateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCatalogCategory := []string{"FieldsCatalogCategory_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCreateCategoriesJob(context.Background(), jobId).Revision(revision).FieldsCatalogCategoryBulkCreateJob(fieldsCatalogCategoryBulkCreateJob).FieldsCatalogCategory(fieldsCatalogCategory).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCreateCategoriesJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreateCategoriesJob`: GetCatalogCategoryCreateJobResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCreateCategoriesJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | ID of the job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreateCategoriesJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogCategoryBulkCreateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCatalogCategory** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetCatalogCategoryCreateJobResponseCompoundDocument**](GetCatalogCategoryCreateJobResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreateCategoriesJobs

> GetCatalogCategoryCreateJobResponseCollectionCompoundDocument GetCreateCategoriesJobs(ctx).Revision(revision).FieldsCatalogCategoryBulkCreateJob(fieldsCatalogCategoryBulkCreateJob).Filter(filter).PageCursor(pageCursor).Execute()

Get Create Categories Jobs



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
	fieldsCatalogCategoryBulkCreateJob := []string{"FieldsCatalogCategoryBulkCreateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`status`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCreateCategoriesJobs(context.Background()).Revision(revision).FieldsCatalogCategoryBulkCreateJob(fieldsCatalogCategoryBulkCreateJob).Filter(filter).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCreateCategoriesJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreateCategoriesJobs`: GetCatalogCategoryCreateJobResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCreateCategoriesJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreateCategoriesJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogCategoryBulkCreateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;status&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCatalogCategoryCreateJobResponseCollectionCompoundDocument**](GetCatalogCategoryCreateJobResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreateItemsJob

> GetCatalogItemCreateJobResponseCompoundDocument GetCreateItemsJob(ctx, jobId).Revision(revision).FieldsCatalogItemBulkCreateJob(fieldsCatalogItemBulkCreateJob).FieldsCatalogItem(fieldsCatalogItem).Include(include).Execute()

Get Create Items Job



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
	fieldsCatalogItemBulkCreateJob := []string{"FieldsCatalogItemBulkCreateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCatalogItem := []string{"FieldsCatalogItem_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCreateItemsJob(context.Background(), jobId).Revision(revision).FieldsCatalogItemBulkCreateJob(fieldsCatalogItemBulkCreateJob).FieldsCatalogItem(fieldsCatalogItem).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCreateItemsJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreateItemsJob`: GetCatalogItemCreateJobResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCreateItemsJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | ID of the job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreateItemsJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogItemBulkCreateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCatalogItem** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetCatalogItemCreateJobResponseCompoundDocument**](GetCatalogItemCreateJobResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreateItemsJobs

> GetCatalogItemCreateJobResponseCollectionCompoundDocument GetCreateItemsJobs(ctx).Revision(revision).FieldsCatalogItemBulkCreateJob(fieldsCatalogItemBulkCreateJob).Filter(filter).PageCursor(pageCursor).Execute()

Get Create Items Jobs



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
	fieldsCatalogItemBulkCreateJob := []string{"FieldsCatalogItemBulkCreateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`status`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCreateItemsJobs(context.Background()).Revision(revision).FieldsCatalogItemBulkCreateJob(fieldsCatalogItemBulkCreateJob).Filter(filter).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCreateItemsJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreateItemsJobs`: GetCatalogItemCreateJobResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCreateItemsJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreateItemsJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogItemBulkCreateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;status&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCatalogItemCreateJobResponseCollectionCompoundDocument**](GetCatalogItemCreateJobResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreateVariantsJob

> GetCatalogVariantCreateJobResponseCompoundDocument GetCreateVariantsJob(ctx, jobId).Revision(revision).FieldsCatalogVariantBulkCreateJob(fieldsCatalogVariantBulkCreateJob).FieldsCatalogVariant(fieldsCatalogVariant).Include(include).Execute()

Get Create Variants Job



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
	fieldsCatalogVariantBulkCreateJob := []string{"FieldsCatalogVariantBulkCreateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCatalogVariant := []string{"FieldsCatalogVariant_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCreateVariantsJob(context.Background(), jobId).Revision(revision).FieldsCatalogVariantBulkCreateJob(fieldsCatalogVariantBulkCreateJob).FieldsCatalogVariant(fieldsCatalogVariant).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCreateVariantsJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreateVariantsJob`: GetCatalogVariantCreateJobResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCreateVariantsJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | ID of the job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreateVariantsJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogVariantBulkCreateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCatalogVariant** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetCatalogVariantCreateJobResponseCompoundDocument**](GetCatalogVariantCreateJobResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreateVariantsJobs

> GetCatalogVariantCreateJobResponseCollectionCompoundDocument GetCreateVariantsJobs(ctx).Revision(revision).FieldsCatalogVariantBulkCreateJob(fieldsCatalogVariantBulkCreateJob).Filter(filter).PageCursor(pageCursor).Execute()

Get Create Variants Jobs



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
	fieldsCatalogVariantBulkCreateJob := []string{"FieldsCatalogVariantBulkCreateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`status`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCreateVariantsJobs(context.Background()).Revision(revision).FieldsCatalogVariantBulkCreateJob(fieldsCatalogVariantBulkCreateJob).Filter(filter).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCreateVariantsJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreateVariantsJobs`: GetCatalogVariantCreateJobResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCreateVariantsJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreateVariantsJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogVariantBulkCreateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;status&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCatalogVariantCreateJobResponseCollectionCompoundDocument**](GetCatalogVariantCreateJobResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeleteCategoriesJob

> GetCatalogCategoryDeleteJobResponse GetDeleteCategoriesJob(ctx, jobId).Revision(revision).FieldsCatalogCategoryBulkDeleteJob(fieldsCatalogCategoryBulkDeleteJob).Execute()

Get Delete Categories Job



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
	fieldsCatalogCategoryBulkDeleteJob := []string{"FieldsCatalogCategoryBulkDeleteJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetDeleteCategoriesJob(context.Background(), jobId).Revision(revision).FieldsCatalogCategoryBulkDeleteJob(fieldsCatalogCategoryBulkDeleteJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetDeleteCategoriesJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeleteCategoriesJob`: GetCatalogCategoryDeleteJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetDeleteCategoriesJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | ID of the job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeleteCategoriesJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogCategoryBulkDeleteJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetCatalogCategoryDeleteJobResponse**](GetCatalogCategoryDeleteJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeleteCategoriesJobs

> GetCatalogCategoryDeleteJobResponseCollection GetDeleteCategoriesJobs(ctx).Revision(revision).FieldsCatalogCategoryBulkDeleteJob(fieldsCatalogCategoryBulkDeleteJob).Filter(filter).PageCursor(pageCursor).Execute()

Get Delete Categories Jobs



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
	fieldsCatalogCategoryBulkDeleteJob := []string{"FieldsCatalogCategoryBulkDeleteJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`status`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetDeleteCategoriesJobs(context.Background()).Revision(revision).FieldsCatalogCategoryBulkDeleteJob(fieldsCatalogCategoryBulkDeleteJob).Filter(filter).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetDeleteCategoriesJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeleteCategoriesJobs`: GetCatalogCategoryDeleteJobResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetDeleteCategoriesJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeleteCategoriesJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogCategoryBulkDeleteJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;status&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCatalogCategoryDeleteJobResponseCollection**](GetCatalogCategoryDeleteJobResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeleteItemsJob

> GetCatalogItemDeleteJobResponse GetDeleteItemsJob(ctx, jobId).Revision(revision).FieldsCatalogItemBulkDeleteJob(fieldsCatalogItemBulkDeleteJob).Execute()

Get Delete Items Job



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
	fieldsCatalogItemBulkDeleteJob := []string{"FieldsCatalogItemBulkDeleteJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetDeleteItemsJob(context.Background(), jobId).Revision(revision).FieldsCatalogItemBulkDeleteJob(fieldsCatalogItemBulkDeleteJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetDeleteItemsJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeleteItemsJob`: GetCatalogItemDeleteJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetDeleteItemsJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | ID of the job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeleteItemsJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogItemBulkDeleteJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetCatalogItemDeleteJobResponse**](GetCatalogItemDeleteJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeleteItemsJobs

> GetCatalogItemDeleteJobResponseCollection GetDeleteItemsJobs(ctx).Revision(revision).FieldsCatalogItemBulkDeleteJob(fieldsCatalogItemBulkDeleteJob).Filter(filter).PageCursor(pageCursor).Execute()

Get Delete Items Jobs



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
	fieldsCatalogItemBulkDeleteJob := []string{"FieldsCatalogItemBulkDeleteJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`status`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetDeleteItemsJobs(context.Background()).Revision(revision).FieldsCatalogItemBulkDeleteJob(fieldsCatalogItemBulkDeleteJob).Filter(filter).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetDeleteItemsJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeleteItemsJobs`: GetCatalogItemDeleteJobResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetDeleteItemsJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeleteItemsJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogItemBulkDeleteJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;status&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCatalogItemDeleteJobResponseCollection**](GetCatalogItemDeleteJobResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeleteVariantsJob

> GetCatalogVariantDeleteJobResponse GetDeleteVariantsJob(ctx, jobId).Revision(revision).FieldsCatalogVariantBulkDeleteJob(fieldsCatalogVariantBulkDeleteJob).Execute()

Get Delete Variants Job



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
	fieldsCatalogVariantBulkDeleteJob := []string{"FieldsCatalogVariantBulkDeleteJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetDeleteVariantsJob(context.Background(), jobId).Revision(revision).FieldsCatalogVariantBulkDeleteJob(fieldsCatalogVariantBulkDeleteJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetDeleteVariantsJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeleteVariantsJob`: GetCatalogVariantDeleteJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetDeleteVariantsJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | ID of the job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeleteVariantsJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogVariantBulkDeleteJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetCatalogVariantDeleteJobResponse**](GetCatalogVariantDeleteJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeleteVariantsJobs

> GetCatalogVariantDeleteJobResponseCollection GetDeleteVariantsJobs(ctx).Revision(revision).FieldsCatalogVariantBulkDeleteJob(fieldsCatalogVariantBulkDeleteJob).Filter(filter).PageCursor(pageCursor).Execute()

Get Delete Variants Jobs



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
	fieldsCatalogVariantBulkDeleteJob := []string{"FieldsCatalogVariantBulkDeleteJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`status`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetDeleteVariantsJobs(context.Background()).Revision(revision).FieldsCatalogVariantBulkDeleteJob(fieldsCatalogVariantBulkDeleteJob).Filter(filter).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetDeleteVariantsJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeleteVariantsJobs`: GetCatalogVariantDeleteJobResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetDeleteVariantsJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeleteVariantsJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogVariantBulkDeleteJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;status&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCatalogVariantDeleteJobResponseCollection**](GetCatalogVariantDeleteJobResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateCategoriesJob

> GetCatalogCategoryUpdateJobResponseCompoundDocument GetUpdateCategoriesJob(ctx, jobId).Revision(revision).FieldsCatalogCategoryBulkUpdateJob(fieldsCatalogCategoryBulkUpdateJob).FieldsCatalogCategory(fieldsCatalogCategory).Include(include).Execute()

Get Update Categories Job



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
	fieldsCatalogCategoryBulkUpdateJob := []string{"FieldsCatalogCategoryBulkUpdateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCatalogCategory := []string{"FieldsCatalogCategory_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetUpdateCategoriesJob(context.Background(), jobId).Revision(revision).FieldsCatalogCategoryBulkUpdateJob(fieldsCatalogCategoryBulkUpdateJob).FieldsCatalogCategory(fieldsCatalogCategory).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetUpdateCategoriesJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpdateCategoriesJob`: GetCatalogCategoryUpdateJobResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetUpdateCategoriesJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | ID of the job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpdateCategoriesJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogCategoryBulkUpdateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCatalogCategory** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetCatalogCategoryUpdateJobResponseCompoundDocument**](GetCatalogCategoryUpdateJobResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateCategoriesJobs

> GetCatalogCategoryUpdateJobResponseCollectionCompoundDocument GetUpdateCategoriesJobs(ctx).Revision(revision).FieldsCatalogCategoryBulkUpdateJob(fieldsCatalogCategoryBulkUpdateJob).Filter(filter).PageCursor(pageCursor).Execute()

Get Update Categories Jobs



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
	fieldsCatalogCategoryBulkUpdateJob := []string{"FieldsCatalogCategoryBulkUpdateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`status`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetUpdateCategoriesJobs(context.Background()).Revision(revision).FieldsCatalogCategoryBulkUpdateJob(fieldsCatalogCategoryBulkUpdateJob).Filter(filter).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetUpdateCategoriesJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpdateCategoriesJobs`: GetCatalogCategoryUpdateJobResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetUpdateCategoriesJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUpdateCategoriesJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogCategoryBulkUpdateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;status&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCatalogCategoryUpdateJobResponseCollectionCompoundDocument**](GetCatalogCategoryUpdateJobResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateItemsJob

> GetCatalogItemUpdateJobResponseCompoundDocument GetUpdateItemsJob(ctx, jobId).Revision(revision).FieldsCatalogItemBulkUpdateJob(fieldsCatalogItemBulkUpdateJob).FieldsCatalogItem(fieldsCatalogItem).Include(include).Execute()

Get Update Items Job



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
	fieldsCatalogItemBulkUpdateJob := []string{"FieldsCatalogItemBulkUpdateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCatalogItem := []string{"FieldsCatalogItem_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetUpdateItemsJob(context.Background(), jobId).Revision(revision).FieldsCatalogItemBulkUpdateJob(fieldsCatalogItemBulkUpdateJob).FieldsCatalogItem(fieldsCatalogItem).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetUpdateItemsJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpdateItemsJob`: GetCatalogItemUpdateJobResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetUpdateItemsJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | ID of the job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpdateItemsJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogItemBulkUpdateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCatalogItem** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetCatalogItemUpdateJobResponseCompoundDocument**](GetCatalogItemUpdateJobResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateItemsJobs

> GetCatalogItemUpdateJobResponseCollectionCompoundDocument GetUpdateItemsJobs(ctx).Revision(revision).FieldsCatalogItemBulkUpdateJob(fieldsCatalogItemBulkUpdateJob).Filter(filter).PageCursor(pageCursor).Execute()

Get Update Items Jobs



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
	fieldsCatalogItemBulkUpdateJob := []string{"FieldsCatalogItemBulkUpdateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`status`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetUpdateItemsJobs(context.Background()).Revision(revision).FieldsCatalogItemBulkUpdateJob(fieldsCatalogItemBulkUpdateJob).Filter(filter).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetUpdateItemsJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpdateItemsJobs`: GetCatalogItemUpdateJobResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetUpdateItemsJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUpdateItemsJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogItemBulkUpdateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;status&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCatalogItemUpdateJobResponseCollectionCompoundDocument**](GetCatalogItemUpdateJobResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateVariantsJob

> GetCatalogVariantUpdateJobResponseCompoundDocument GetUpdateVariantsJob(ctx, jobId).Revision(revision).FieldsCatalogVariantBulkUpdateJob(fieldsCatalogVariantBulkUpdateJob).FieldsCatalogVariant(fieldsCatalogVariant).Include(include).Execute()

Get Update Variants Job



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
	fieldsCatalogVariantBulkUpdateJob := []string{"FieldsCatalogVariantBulkUpdateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCatalogVariant := []string{"FieldsCatalogVariant_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetUpdateVariantsJob(context.Background(), jobId).Revision(revision).FieldsCatalogVariantBulkUpdateJob(fieldsCatalogVariantBulkUpdateJob).FieldsCatalogVariant(fieldsCatalogVariant).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetUpdateVariantsJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpdateVariantsJob`: GetCatalogVariantUpdateJobResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetUpdateVariantsJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | ID of the job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpdateVariantsJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogVariantBulkUpdateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCatalogVariant** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetCatalogVariantUpdateJobResponseCompoundDocument**](GetCatalogVariantUpdateJobResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateVariantsJobs

> GetCatalogVariantUpdateJobResponseCollectionCompoundDocument GetUpdateVariantsJobs(ctx).Revision(revision).FieldsCatalogVariantBulkUpdateJob(fieldsCatalogVariantBulkUpdateJob).Filter(filter).PageCursor(pageCursor).Execute()

Get Update Variants Jobs



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
	fieldsCatalogVariantBulkUpdateJob := []string{"FieldsCatalogVariantBulkUpdateJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`status`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetUpdateVariantsJobs(context.Background()).Revision(revision).FieldsCatalogVariantBulkUpdateJob(fieldsCatalogVariantBulkUpdateJob).Filter(filter).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetUpdateVariantsJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpdateVariantsJobs`: GetCatalogVariantUpdateJobResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetUpdateVariantsJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUpdateVariantsJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCatalogVariantBulkUpdateJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;status&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetCatalogVariantUpdateJobResponseCollectionCompoundDocument**](GetCatalogVariantUpdateJobResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpawnCreateCategoriesJob

> PostCatalogCategoryCreateJobResponse SpawnCreateCategoriesJob(ctx).Revision(revision).CatalogCategoryCreateJobCreateQuery(catalogCategoryCreateJobCreateQuery).Execute()

Spawn Create Categories Job



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
	catalogCategoryCreateJobCreateQuery := *openapiclient.NewCatalogCategoryCreateJobCreateQuery(*openapiclient.NewCatalogCategoryCreateJobCreateQueryResourceObject(openapiclient.CatalogCategoryBulkCreateJobEnum("catalog-category-bulk-create-job"), *openapiclient.NewCatalogCategoryCreateJobCreateQueryResourceObjectAttributes(*openapiclient.NewCatalogCategoryCreateJobCreateQueryResourceObjectAttributesCategories([]openapiclient.CatalogCategoryCreateQueryResourceObject{*openapiclient.NewCatalogCategoryCreateQueryResourceObject(openapiclient.CatalogCategoryEnum("catalog-category"), *openapiclient.NewCatalogCategoryCreateQueryResourceObjectAttributes("SAMPLE-DATA-CATEGORY-APPAREL", "Sample Data Category Apparel"))})))) // CatalogCategoryCreateJobCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.SpawnCreateCategoriesJob(context.Background()).Revision(revision).CatalogCategoryCreateJobCreateQuery(catalogCategoryCreateJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.SpawnCreateCategoriesJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpawnCreateCategoriesJob`: PostCatalogCategoryCreateJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.SpawnCreateCategoriesJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpawnCreateCategoriesJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogCategoryCreateJobCreateQuery** | [**CatalogCategoryCreateJobCreateQuery**](CatalogCategoryCreateJobCreateQuery.md) |  | 

### Return type

[**PostCatalogCategoryCreateJobResponse**](PostCatalogCategoryCreateJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpawnCreateItemsJob

> PostCatalogItemCreateJobResponse SpawnCreateItemsJob(ctx).Revision(revision).CatalogItemCreateJobCreateQuery(catalogItemCreateJobCreateQuery).Execute()

Spawn Create Items Job



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
	catalogItemCreateJobCreateQuery := *openapiclient.NewCatalogItemCreateJobCreateQuery(*openapiclient.NewCatalogItemCreateJobCreateQueryResourceObject(openapiclient.CatalogItemBulkCreateJobEnum("catalog-item-bulk-create-job"), *openapiclient.NewCatalogItemCreateJobCreateQueryResourceObjectAttributes(*openapiclient.NewCatalogItemCreateJobCreateQueryResourceObjectAttributesItems([]openapiclient.CatalogItemCreateQueryResourceObject{*openapiclient.NewCatalogItemCreateQueryResourceObject(openapiclient.CatalogItemEnum("catalog-item"), *openapiclient.NewCatalogItemCreateQueryResourceObjectAttributes("SAMPLE-DATA-ITEM-1", "Ocean Blue Shirt (Sample)", "Ocean blue cotton shirt with a narrow collar and buttons down the front and long sleeves. Comfortable fit and titled kaleidoscope patterns.", "https://via.placeholder.com/150"))})))) // CatalogItemCreateJobCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.SpawnCreateItemsJob(context.Background()).Revision(revision).CatalogItemCreateJobCreateQuery(catalogItemCreateJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.SpawnCreateItemsJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpawnCreateItemsJob`: PostCatalogItemCreateJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.SpawnCreateItemsJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpawnCreateItemsJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogItemCreateJobCreateQuery** | [**CatalogItemCreateJobCreateQuery**](CatalogItemCreateJobCreateQuery.md) |  | 

### Return type

[**PostCatalogItemCreateJobResponse**](PostCatalogItemCreateJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpawnCreateVariantsJob

> PostCatalogVariantCreateJobResponse SpawnCreateVariantsJob(ctx).Revision(revision).CatalogVariantCreateJobCreateQuery(catalogVariantCreateJobCreateQuery).Execute()

Spawn Create Variants Job



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
	catalogVariantCreateJobCreateQuery := *openapiclient.NewCatalogVariantCreateJobCreateQuery(*openapiclient.NewCatalogVariantCreateJobCreateQueryResourceObject(openapiclient.CatalogVariantBulkCreateJobEnum("catalog-variant-bulk-create-job"), *openapiclient.NewCatalogVariantCreateJobCreateQueryResourceObjectAttributes(*openapiclient.NewCatalogVariantCreateJobCreateQueryResourceObjectAttributesVariants([]openapiclient.CatalogVariantCreateQueryResourceObject{*openapiclient.NewCatalogVariantCreateQueryResourceObject(openapiclient.CatalogVariantEnum("catalog-variant"), *openapiclient.NewCatalogVariantCreateQueryResourceObjectAttributes("SAMPLE-DATA-ITEM-1-VARIANT-MEDIUM", "Ocean Blue Shirt (Sample) Variant Medium", "Ocean blue cotton shirt with a narrow collar and buttons down the front and long sleeves. Comfortable fit and titled kaleidoscope patterns.", "OBS-MD", float32(25), float32(42), "https://via.placeholder.com/150"), *openapiclient.NewCatalogVariantCreateQueryResourceObjectRelationships(*openapiclient.NewCatalogVariantCreateQueryResourceObjectRelationshipsItem(*openapiclient.NewCatalogVariantCreateQueryResourceObjectRelationshipsItemData(openapiclient.CatalogItemEnum("catalog-item"), "$custom:::$default:::SAMPLE-DATA-ITEM-1"))))})))) // CatalogVariantCreateJobCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.SpawnCreateVariantsJob(context.Background()).Revision(revision).CatalogVariantCreateJobCreateQuery(catalogVariantCreateJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.SpawnCreateVariantsJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpawnCreateVariantsJob`: PostCatalogVariantCreateJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.SpawnCreateVariantsJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpawnCreateVariantsJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogVariantCreateJobCreateQuery** | [**CatalogVariantCreateJobCreateQuery**](CatalogVariantCreateJobCreateQuery.md) |  | 

### Return type

[**PostCatalogVariantCreateJobResponse**](PostCatalogVariantCreateJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpawnDeleteCategoriesJob

> PostCatalogCategoryDeleteJobResponse SpawnDeleteCategoriesJob(ctx).Revision(revision).CatalogCategoryDeleteJobCreateQuery(catalogCategoryDeleteJobCreateQuery).Execute()

Spawn Delete Categories Job



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
	catalogCategoryDeleteJobCreateQuery := *openapiclient.NewCatalogCategoryDeleteJobCreateQuery(*openapiclient.NewCatalogCategoryDeleteJobCreateQueryResourceObject(openapiclient.CatalogCategoryBulkDeleteJobEnum("catalog-category-bulk-delete-job"), *openapiclient.NewCatalogCategoryDeleteJobCreateQueryResourceObjectAttributes(*openapiclient.NewCatalogCategoryDeleteJobCreateQueryResourceObjectAttributesCategories([]openapiclient.CatalogCategoryDeleteQueryResourceObject{*openapiclient.NewCatalogCategoryDeleteQueryResourceObject(openapiclient.CatalogCategoryEnum("catalog-category"), "$custom:::$default:::SAMPLE-DATA-CATEGORY-APPAREL")})))) // CatalogCategoryDeleteJobCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.SpawnDeleteCategoriesJob(context.Background()).Revision(revision).CatalogCategoryDeleteJobCreateQuery(catalogCategoryDeleteJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.SpawnDeleteCategoriesJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpawnDeleteCategoriesJob`: PostCatalogCategoryDeleteJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.SpawnDeleteCategoriesJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpawnDeleteCategoriesJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogCategoryDeleteJobCreateQuery** | [**CatalogCategoryDeleteJobCreateQuery**](CatalogCategoryDeleteJobCreateQuery.md) |  | 

### Return type

[**PostCatalogCategoryDeleteJobResponse**](PostCatalogCategoryDeleteJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpawnDeleteItemsJob

> PostCatalogItemDeleteJobResponse SpawnDeleteItemsJob(ctx).Revision(revision).CatalogItemDeleteJobCreateQuery(catalogItemDeleteJobCreateQuery).Execute()

Spawn Delete Items Job



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
	catalogItemDeleteJobCreateQuery := *openapiclient.NewCatalogItemDeleteJobCreateQuery(*openapiclient.NewCatalogItemDeleteJobCreateQueryResourceObject(openapiclient.CatalogItemBulkDeleteJobEnum("catalog-item-bulk-delete-job"), *openapiclient.NewCatalogItemDeleteJobCreateQueryResourceObjectAttributes(*openapiclient.NewCatalogItemDeleteJobCreateQueryResourceObjectAttributesItems([]openapiclient.CatalogItemDeleteQueryResourceObject{*openapiclient.NewCatalogItemDeleteQueryResourceObject(openapiclient.CatalogItemEnum("catalog-item"), "$custom:::$default:::SAMPLE-DATA-ITEM-1")})))) // CatalogItemDeleteJobCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.SpawnDeleteItemsJob(context.Background()).Revision(revision).CatalogItemDeleteJobCreateQuery(catalogItemDeleteJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.SpawnDeleteItemsJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpawnDeleteItemsJob`: PostCatalogItemDeleteJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.SpawnDeleteItemsJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpawnDeleteItemsJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogItemDeleteJobCreateQuery** | [**CatalogItemDeleteJobCreateQuery**](CatalogItemDeleteJobCreateQuery.md) |  | 

### Return type

[**PostCatalogItemDeleteJobResponse**](PostCatalogItemDeleteJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpawnDeleteVariantsJob

> PostCatalogVariantDeleteJobResponse SpawnDeleteVariantsJob(ctx).Revision(revision).CatalogVariantDeleteJobCreateQuery(catalogVariantDeleteJobCreateQuery).Execute()

Spawn Delete Variants Job



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
	catalogVariantDeleteJobCreateQuery := *openapiclient.NewCatalogVariantDeleteJobCreateQuery(*openapiclient.NewCatalogVariantDeleteJobCreateQueryResourceObject(openapiclient.CatalogVariantBulkDeleteJobEnum("catalog-variant-bulk-delete-job"), *openapiclient.NewCatalogVariantDeleteJobCreateQueryResourceObjectAttributes(*openapiclient.NewCatalogVariantDeleteJobCreateQueryResourceObjectAttributesVariants([]openapiclient.CatalogVariantDeleteQueryResourceObject{*openapiclient.NewCatalogVariantDeleteQueryResourceObject(openapiclient.CatalogVariantEnum("catalog-variant"), "$custom:::$default:::SAMPLE-DATA-ITEM-1-VARIANT-MEDIUM")})))) // CatalogVariantDeleteJobCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.SpawnDeleteVariantsJob(context.Background()).Revision(revision).CatalogVariantDeleteJobCreateQuery(catalogVariantDeleteJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.SpawnDeleteVariantsJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpawnDeleteVariantsJob`: PostCatalogVariantDeleteJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.SpawnDeleteVariantsJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpawnDeleteVariantsJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogVariantDeleteJobCreateQuery** | [**CatalogVariantDeleteJobCreateQuery**](CatalogVariantDeleteJobCreateQuery.md) |  | 

### Return type

[**PostCatalogVariantDeleteJobResponse**](PostCatalogVariantDeleteJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpawnUpdateCategoriesJob

> PostCatalogCategoryUpdateJobResponse SpawnUpdateCategoriesJob(ctx).Revision(revision).CatalogCategoryUpdateJobCreateQuery(catalogCategoryUpdateJobCreateQuery).Execute()

Spawn Update Categories Job



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
	catalogCategoryUpdateJobCreateQuery := *openapiclient.NewCatalogCategoryUpdateJobCreateQuery(*openapiclient.NewCatalogCategoryUpdateJobCreateQueryResourceObject(openapiclient.CatalogCategoryBulkUpdateJobEnum("catalog-category-bulk-update-job"), *openapiclient.NewCatalogCategoryUpdateJobCreateQueryResourceObjectAttributes(*openapiclient.NewCatalogCategoryUpdateJobCreateQueryResourceObjectAttributesCategories([]openapiclient.CatalogCategoryUpdateQueryResourceObject{*openapiclient.NewCatalogCategoryUpdateQueryResourceObject(openapiclient.CatalogCategoryEnum("catalog-category"), "$custom:::$default:::SAMPLE-DATA-CATEGORY-APPAREL", *openapiclient.NewCatalogCategoryUpdateQueryResourceObjectAttributes())})))) // CatalogCategoryUpdateJobCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.SpawnUpdateCategoriesJob(context.Background()).Revision(revision).CatalogCategoryUpdateJobCreateQuery(catalogCategoryUpdateJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.SpawnUpdateCategoriesJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpawnUpdateCategoriesJob`: PostCatalogCategoryUpdateJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.SpawnUpdateCategoriesJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpawnUpdateCategoriesJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogCategoryUpdateJobCreateQuery** | [**CatalogCategoryUpdateJobCreateQuery**](CatalogCategoryUpdateJobCreateQuery.md) |  | 

### Return type

[**PostCatalogCategoryUpdateJobResponse**](PostCatalogCategoryUpdateJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpawnUpdateItemsJob

> PostCatalogItemUpdateJobResponse SpawnUpdateItemsJob(ctx).Revision(revision).CatalogItemUpdateJobCreateQuery(catalogItemUpdateJobCreateQuery).Execute()

Spawn Update Items Job



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
	catalogItemUpdateJobCreateQuery := *openapiclient.NewCatalogItemUpdateJobCreateQuery(*openapiclient.NewCatalogItemUpdateJobCreateQueryResourceObject(openapiclient.CatalogItemBulkUpdateJobEnum("catalog-item-bulk-update-job"), *openapiclient.NewCatalogItemUpdateJobCreateQueryResourceObjectAttributes(*openapiclient.NewCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems([]openapiclient.CatalogItemUpdateQueryResourceObject{*openapiclient.NewCatalogItemUpdateQueryResourceObject(openapiclient.CatalogItemEnum("catalog-item"), "$custom:::$default:::SAMPLE-DATA-ITEM-1", *openapiclient.NewCatalogItemUpdateQueryResourceObjectAttributes())})))) // CatalogItemUpdateJobCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.SpawnUpdateItemsJob(context.Background()).Revision(revision).CatalogItemUpdateJobCreateQuery(catalogItemUpdateJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.SpawnUpdateItemsJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpawnUpdateItemsJob`: PostCatalogItemUpdateJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.SpawnUpdateItemsJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpawnUpdateItemsJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogItemUpdateJobCreateQuery** | [**CatalogItemUpdateJobCreateQuery**](CatalogItemUpdateJobCreateQuery.md) |  | 

### Return type

[**PostCatalogItemUpdateJobResponse**](PostCatalogItemUpdateJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpawnUpdateVariantsJob

> PostCatalogVariantUpdateJobResponse SpawnUpdateVariantsJob(ctx).Revision(revision).CatalogVariantUpdateJobCreateQuery(catalogVariantUpdateJobCreateQuery).Execute()

Spawn Update Variants Job



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
	catalogVariantUpdateJobCreateQuery := *openapiclient.NewCatalogVariantUpdateJobCreateQuery(*openapiclient.NewCatalogVariantUpdateJobCreateQueryResourceObject(openapiclient.CatalogVariantBulkUpdateJobEnum("catalog-variant-bulk-update-job"), *openapiclient.NewCatalogVariantUpdateJobCreateQueryResourceObjectAttributes(*openapiclient.NewCatalogVariantUpdateJobCreateQueryResourceObjectAttributesVariants([]openapiclient.CatalogVariantUpdateQueryResourceObject{*openapiclient.NewCatalogVariantUpdateQueryResourceObject(openapiclient.CatalogVariantEnum("catalog-variant"), "$custom:::$default:::SAMPLE-DATA-ITEM-1-VARIANT-MEDIUM", *openapiclient.NewCatalogVariantUpdateQueryResourceObjectAttributes())})))) // CatalogVariantUpdateJobCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.SpawnUpdateVariantsJob(context.Background()).Revision(revision).CatalogVariantUpdateJobCreateQuery(catalogVariantUpdateJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.SpawnUpdateVariantsJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpawnUpdateVariantsJob`: PostCatalogVariantUpdateJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.SpawnUpdateVariantsJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpawnUpdateVariantsJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogVariantUpdateJobCreateQuery** | [**CatalogVariantUpdateJobCreateQuery**](CatalogVariantUpdateJobCreateQuery.md) |  | 

### Return type

[**PostCatalogVariantUpdateJobResponse**](PostCatalogVariantUpdateJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogCategory

> PatchCatalogCategoryResponse UpdateCatalogCategory(ctx, id).Revision(revision).CatalogCategoryUpdateQuery(catalogCategoryUpdateQuery).Execute()

Update Catalog Category



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
	id := "$custom:::$default:::SAMPLE-DATA-CATEGORY-APPAREL" // string | The catalog category ID is a compound ID (string), with format: `{integration}:::{catalog}:::{external_id}`. Currently, the only supported integration type is `$custom`, and the only supported catalog is `$default`.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	catalogCategoryUpdateQuery := *openapiclient.NewCatalogCategoryUpdateQuery(*openapiclient.NewCatalogCategoryUpdateQueryResourceObject(openapiclient.CatalogCategoryEnum("catalog-category"), "$custom:::$default:::SAMPLE-DATA-CATEGORY-APPAREL", *openapiclient.NewCatalogCategoryUpdateQueryResourceObjectAttributes())) // CatalogCategoryUpdateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.UpdateCatalogCategory(context.Background(), id).Revision(revision).CatalogCategoryUpdateQuery(catalogCategoryUpdateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.UpdateCatalogCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCatalogCategory`: PatchCatalogCategoryResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.UpdateCatalogCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The catalog category ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogCategoryUpdateQuery** | [**CatalogCategoryUpdateQuery**](CatalogCategoryUpdateQuery.md) |  | 

### Return type

[**PatchCatalogCategoryResponse**](PatchCatalogCategoryResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogCategoryRelationshipsItems

> UpdateCatalogCategoryRelationshipsItems(ctx, id).Revision(revision).CatalogCategoryItemOp(catalogCategoryItemOp).Execute()

Update Catalog Category Relationships Items



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
	catalogCategoryItemOp := *openapiclient.NewCatalogCategoryItemOp([]openapiclient.GetCatalogCategoryItemListResponseCollectionDataInner{*openapiclient.NewGetCatalogCategoryItemListResponseCollectionDataInner(openapiclient.CatalogItemEnum("catalog-item"), "$custom:::$default:::SAMPLE-DATA-ITEM-1")}) // CatalogCategoryItemOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogsAPI.UpdateCatalogCategoryRelationshipsItems(context.Background(), id).Revision(revision).CatalogCategoryItemOp(catalogCategoryItemOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.UpdateCatalogCategoryRelationshipsItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogCategoryRelationshipsItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogCategoryItemOp** | [**CatalogCategoryItemOp**](CatalogCategoryItemOp.md) |  | 

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


## UpdateCatalogItem

> PatchCatalogItemResponse UpdateCatalogItem(ctx, id).Revision(revision).CatalogItemUpdateQuery(catalogItemUpdateQuery).Execute()

Update Catalog Item



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
	id := "$custom:::$default:::SAMPLE-DATA-ITEM-1" // string | The catalog item ID is a compound ID (string), with format: `{integration}:::{catalog}:::{external_id}`. Currently, the only supported integration type is `$custom`, and the only supported catalog is `$default`.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	catalogItemUpdateQuery := *openapiclient.NewCatalogItemUpdateQuery(*openapiclient.NewCatalogItemUpdateQueryResourceObject(openapiclient.CatalogItemEnum("catalog-item"), "$custom:::$default:::SAMPLE-DATA-ITEM-1", *openapiclient.NewCatalogItemUpdateQueryResourceObjectAttributes())) // CatalogItemUpdateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.UpdateCatalogItem(context.Background(), id).Revision(revision).CatalogItemUpdateQuery(catalogItemUpdateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.UpdateCatalogItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCatalogItem`: PatchCatalogItemResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.UpdateCatalogItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The catalog item ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogItemUpdateQuery** | [**CatalogItemUpdateQuery**](CatalogItemUpdateQuery.md) |  | 

### Return type

[**PatchCatalogItemResponse**](PatchCatalogItemResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogItemRelationshipsCategories

> UpdateCatalogItemRelationshipsCategories(ctx, id).Revision(revision).CatalogItemCategoryOp(catalogItemCategoryOp).Execute()

Update Catalog Item Relationships Categories



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
	catalogItemCategoryOp := *openapiclient.NewCatalogItemCategoryOp([]openapiclient.GetCatalogItemCategoryListResponseCollectionDataInner{*openapiclient.NewGetCatalogItemCategoryListResponseCollectionDataInner(openapiclient.CatalogCategoryEnum("catalog-category"), "$custom:::$default:::SAMPLE-DATA-CATEGORY-APPAREL")}) // CatalogItemCategoryOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogsAPI.UpdateCatalogItemRelationshipsCategories(context.Background(), id).Revision(revision).CatalogItemCategoryOp(catalogItemCategoryOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.UpdateCatalogItemRelationshipsCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogItemRelationshipsCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogItemCategoryOp** | [**CatalogItemCategoryOp**](CatalogItemCategoryOp.md) |  | 

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


## UpdateCatalogVariant

> PatchCatalogVariantResponse UpdateCatalogVariant(ctx, id).Revision(revision).CatalogVariantUpdateQuery(catalogVariantUpdateQuery).Execute()

Update Catalog Variant



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
	id := "$custom:::$default:::SAMPLE-DATA-ITEM-1-VARIANT-MEDIUM" // string | The catalog variant ID is a compound ID (string), with format: `{integration}:::{catalog}:::{external_id}`. Currently, the only supported integration type is `$custom`, and the only supported catalog is `$default`.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	catalogVariantUpdateQuery := *openapiclient.NewCatalogVariantUpdateQuery(*openapiclient.NewCatalogVariantUpdateQueryResourceObject(openapiclient.CatalogVariantEnum("catalog-variant"), "$custom:::$default:::SAMPLE-DATA-ITEM-1-VARIANT-MEDIUM", *openapiclient.NewCatalogVariantUpdateQueryResourceObjectAttributes())) // CatalogVariantUpdateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.UpdateCatalogVariant(context.Background(), id).Revision(revision).CatalogVariantUpdateQuery(catalogVariantUpdateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.UpdateCatalogVariant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCatalogVariant`: PatchCatalogVariantResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.UpdateCatalogVariant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The catalog variant ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogVariantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **catalogVariantUpdateQuery** | [**CatalogVariantUpdateQuery**](CatalogVariantUpdateQuery.md) |  | 

### Return type

[**PatchCatalogVariantResponse**](PatchCatalogVariantResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

