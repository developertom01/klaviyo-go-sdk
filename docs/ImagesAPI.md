# \ImagesAPI

All URIs are relative to *https://a.klaviyo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImage**](ImagesAPI.md#GetImage) | **Get** /api/images/{id}/ | Get Image
[**GetImages**](ImagesAPI.md#GetImages) | **Get** /api/images/ | Get Images
[**UpdateImage**](ImagesAPI.md#UpdateImage) | **Patch** /api/images/{id}/ | Update Image
[**UploadImageFromFile**](ImagesAPI.md#UploadImageFromFile) | **Post** /api/image-upload/ | Upload Image From File
[**UploadImageFromUrl**](ImagesAPI.md#UploadImageFromUrl) | **Post** /api/images/ | Upload Image From URL



## GetImage

> GetImageResponse GetImage(ctx, id).Revision(revision).FieldsImage(fieldsImage).Execute()

Get Image



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
	id := "7" // string | The ID of the image
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsImage := []string{"FieldsImage_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImage(context.Background(), id).Revision(revision).FieldsImage(fieldsImage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImage`: GetImageResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the image | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsImage** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetImageResponse**](GetImageResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImages

> GetImageResponseCollection GetImages(ctx).Revision(revision).FieldsImage(fieldsImage).Filter(filter).PageCursor(pageCursor).PageSize(pageSize).Sort(sort).Execute()

Get Images



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
	fieldsImage := []string{"FieldsImage_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`id`: `any`, `equals`<br>`updated_at`: `greater-or-equal`, `greater-than`, `less-or-equal`, `less-than`<br>`format`: `any`, `equals`<br>`name`: `any`, `contains`, `ends-with`, `equals`, `starts-with`<br>`size`: `greater-or-equal`, `greater-than`, `less-or-equal`, `less-than`<br>`hidden`: `any`, `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	pageSize := int32(56) // int32 | Default: 20. Min: 1. Max: 100. (optional) (default to 20)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImages(context.Background()).Revision(revision).FieldsImage(fieldsImage).Filter(filter).PageCursor(pageCursor).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImages`: GetImageResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsImage** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;id&#x60;: &#x60;any&#x60;, &#x60;equals&#x60;&lt;br&gt;&#x60;updated_at&#x60;: &#x60;greater-or-equal&#x60;, &#x60;greater-than&#x60;, &#x60;less-or-equal&#x60;, &#x60;less-than&#x60;&lt;br&gt;&#x60;format&#x60;: &#x60;any&#x60;, &#x60;equals&#x60;&lt;br&gt;&#x60;name&#x60;: &#x60;any&#x60;, &#x60;contains&#x60;, &#x60;ends-with&#x60;, &#x60;equals&#x60;, &#x60;starts-with&#x60;&lt;br&gt;&#x60;size&#x60;: &#x60;greater-or-equal&#x60;, &#x60;greater-than&#x60;, &#x60;less-or-equal&#x60;, &#x60;less-than&#x60;&lt;br&gt;&#x60;hidden&#x60;: &#x60;any&#x60;, &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **pageSize** | **int32** | Default: 20. Min: 1. Max: 100. | [default to 20]
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetImageResponseCollection**](GetImageResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateImage

> PatchImageResponse UpdateImage(ctx, id).Revision(revision).ImagePartialUpdateQuery(imagePartialUpdateQuery).Execute()

Update Image



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
	id := "7" // string | The ID of the image
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	imagePartialUpdateQuery := *openapiclient.NewImagePartialUpdateQuery(*openapiclient.NewImagePartialUpdateQueryResourceObject(openapiclient.ImageEnum("image"), "7", *openapiclient.NewImagePartialUpdateQueryResourceObjectAttributes())) // ImagePartialUpdateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.UpdateImage(context.Background(), id).Revision(revision).ImagePartialUpdateQuery(imagePartialUpdateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.UpdateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateImage`: PatchImageResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.UpdateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the image | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **imagePartialUpdateQuery** | [**ImagePartialUpdateQuery**](ImagePartialUpdateQuery.md) |  | 

### Return type

[**PatchImageResponse**](PatchImageResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadImageFromFile

> PostImageResponse UploadImageFromFile(ctx).Revision(revision).File(file).Name(name).Hidden(hidden).Execute()

Upload Image From File



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
	file := os.NewFile(1234, "some_file") // *os.File | The image file to upload. Supported image formats: jpeg,png,gif. Maximum image size: 5MB.
	name := "name_example" // string | A name for the image.  Defaults to the filename if not provided.  If the name matches an existing image, a suffix will be added. (optional)
	hidden := true // bool | If true, this image is not shown in the asset library. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.UploadImageFromFile(context.Background()).Revision(revision).File(file).Name(name).Hidden(hidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.UploadImageFromFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadImageFromFile`: PostImageResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.UploadImageFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadImageFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **file** | ***os.File** | The image file to upload. Supported image formats: jpeg,png,gif. Maximum image size: 5MB. | 
 **name** | **string** | A name for the image.  Defaults to the filename if not provided.  If the name matches an existing image, a suffix will be added. | 
 **hidden** | **bool** | If true, this image is not shown in the asset library. | [default to false]

### Return type

[**PostImageResponse**](PostImageResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadImageFromUrl

> PostImageResponse UploadImageFromUrl(ctx).Revision(revision).ImageCreateQuery(imageCreateQuery).Execute()

Upload Image From URL



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
	imageCreateQuery := *openapiclient.NewImageCreateQuery(*openapiclient.NewImageCreateQueryResourceObject(openapiclient.ImageEnum("image"), *openapiclient.NewImageCreateQueryResourceObjectAttributes("https://www.example.com/example.jpg"))) // ImageCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.UploadImageFromUrl(context.Background()).Revision(revision).ImageCreateQuery(imageCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.UploadImageFromUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadImageFromUrl`: PostImageResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.UploadImageFromUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadImageFromUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **imageCreateQuery** | [**ImageCreateQuery**](ImageCreateQuery.md) |  | 

### Return type

[**PostImageResponse**](PostImageResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

