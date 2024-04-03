# \TemplatesAPI

All URIs are relative to *https://a.klaviyo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplate**](TemplatesAPI.md#CreateTemplate) | **Post** /api/templates/ | Create Template
[**CreateTemplateClone**](TemplatesAPI.md#CreateTemplateClone) | **Post** /api/template-clone/ | Create Template Clone
[**CreateTemplateRender**](TemplatesAPI.md#CreateTemplateRender) | **Post** /api/template-render/ | Create Template Render
[**DeleteTemplate**](TemplatesAPI.md#DeleteTemplate) | **Delete** /api/templates/{id}/ | Delete Template
[**GetTemplate**](TemplatesAPI.md#GetTemplate) | **Get** /api/templates/{id}/ | Get Template
[**GetTemplates**](TemplatesAPI.md#GetTemplates) | **Get** /api/templates/ | Get Templates
[**UpdateTemplate**](TemplatesAPI.md#UpdateTemplate) | **Patch** /api/templates/{id}/ | Update Template



## CreateTemplate

> PostTemplateResponse CreateTemplate(ctx).Revision(revision).TemplateCreateQuery(templateCreateQuery).Execute()

Create Template



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
	templateCreateQuery := *openapiclient.NewTemplateCreateQuery(*openapiclient.NewTemplateCreateQueryResourceObject(openapiclient.TemplateEnum("template"), *openapiclient.NewTemplateCreateQueryResourceObjectAttributes("Monthly Newsletter Template", "EditorType_example"))) // TemplateCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.CreateTemplate(context.Background()).Revision(revision).TemplateCreateQuery(templateCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CreateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplate`: PostTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.CreateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **templateCreateQuery** | [**TemplateCreateQuery**](TemplateCreateQuery.md) |  | 

### Return type

[**PostTemplateResponse**](PostTemplateResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplateClone

> PostTemplateResponse CreateTemplateClone(ctx).Revision(revision).TemplateCloneQuery(templateCloneQuery).Execute()

Create Template Clone



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
	templateCloneQuery := *openapiclient.NewTemplateCloneQuery(*openapiclient.NewTemplateCloneQueryResourceObject(openapiclient.TemplateEnum("template"), "Id_example", *openapiclient.NewTemplateCloneQueryResourceObjectAttributes())) // TemplateCloneQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.CreateTemplateClone(context.Background()).Revision(revision).TemplateCloneQuery(templateCloneQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CreateTemplateClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplateClone`: PostTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.CreateTemplateClone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **templateCloneQuery** | [**TemplateCloneQuery**](TemplateCloneQuery.md) |  | 

### Return type

[**PostTemplateResponse**](PostTemplateResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplateRender

> PostTemplateResponse CreateTemplateRender(ctx).Revision(revision).TemplateRenderQuery(templateRenderQuery).Execute()

Create Template Render



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
	templateRenderQuery := *openapiclient.NewTemplateRenderQuery(*openapiclient.NewTemplateRenderQueryResourceObject(openapiclient.TemplateEnum("template"), "Id_example", *openapiclient.NewTemplateRenderQueryResourceObjectAttributes(map[string]interface{}({"first_name":"Jane","last_name":"Smith"})))) // TemplateRenderQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.CreateTemplateRender(context.Background()).Revision(revision).TemplateRenderQuery(templateRenderQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CreateTemplateRender``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplateRender`: PostTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.CreateTemplateRender`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateRenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **templateRenderQuery** | [**TemplateRenderQuery**](TemplateRenderQuery.md) |  | 

### Return type

[**PostTemplateResponse**](PostTemplateResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplate

> DeleteTemplate(ctx, id).Revision(revision).Execute()

Delete Template



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
	id := "id_example" // string | The ID of template
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TemplatesAPI.DeleteTemplate(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.DeleteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of template | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateRequest struct via the builder pattern


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


## GetTemplate

> GetTemplateResponse GetTemplate(ctx, id).Revision(revision).FieldsTemplate(fieldsTemplate).Execute()

Get Template



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
	id := "id_example" // string | The ID of template
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsTemplate := []string{"FieldsTemplate_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetTemplate(context.Background(), id).Revision(revision).FieldsTemplate(fieldsTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplate`: GetTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsTemplate** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetTemplateResponse**](GetTemplateResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplates

> GetTemplateResponseCollection GetTemplates(ctx).Revision(revision).FieldsTemplate(fieldsTemplate).Filter(filter).PageCursor(pageCursor).Sort(sort).Execute()

Get Templates



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
	fieldsTemplate := []string{"FieldsTemplate_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`id`: `any`, `equals`<br>`name`: `any`, `equals`<br>`created`: `equals`, `greater-or-equal`, `greater-than`, `less-or-equal`, `less-than`<br>`updated`: `equals`, `greater-or-equal`, `greater-than`, `less-or-equal`, `less-than` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetTemplates(context.Background()).Revision(revision).FieldsTemplate(fieldsTemplate).Filter(filter).PageCursor(pageCursor).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplates`: GetTemplateResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsTemplate** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;id&#x60;: &#x60;any&#x60;, &#x60;equals&#x60;&lt;br&gt;&#x60;name&#x60;: &#x60;any&#x60;, &#x60;equals&#x60;&lt;br&gt;&#x60;created&#x60;: &#x60;equals&#x60;, &#x60;greater-or-equal&#x60;, &#x60;greater-than&#x60;, &#x60;less-or-equal&#x60;, &#x60;less-than&#x60;&lt;br&gt;&#x60;updated&#x60;: &#x60;equals&#x60;, &#x60;greater-or-equal&#x60;, &#x60;greater-than&#x60;, &#x60;less-or-equal&#x60;, &#x60;less-than&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetTemplateResponseCollection**](GetTemplateResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplate

> PatchTemplateResponse UpdateTemplate(ctx, id).Revision(revision).TemplateUpdateQuery(templateUpdateQuery).Execute()

Update Template



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
	id := "id_example" // string | The ID of template
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	templateUpdateQuery := *openapiclient.NewTemplateUpdateQuery(*openapiclient.NewTemplateUpdateQueryResourceObject(openapiclient.TemplateEnum("template"), "Id_example", *openapiclient.NewTemplateUpdateQueryResourceObjectAttributes())) // TemplateUpdateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.UpdateTemplate(context.Background(), id).Revision(revision).TemplateUpdateQuery(templateUpdateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.UpdateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTemplate`: PatchTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.UpdateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **templateUpdateQuery** | [**TemplateUpdateQuery**](TemplateUpdateQuery.md) |  | 

### Return type

[**PatchTemplateResponse**](PatchTemplateResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

