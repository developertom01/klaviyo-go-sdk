# \TagsAPI

All URIs are relative to *https://a.klaviyo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagsAPI.md#CreateTag) | **Post** /api/tags/ | Create Tag
[**CreateTagGroup**](TagsAPI.md#CreateTagGroup) | **Post** /api/tag-groups/ | Create Tag Group
[**CreateTagRelationshipsCampaigns**](TagsAPI.md#CreateTagRelationshipsCampaigns) | **Post** /api/tags/{id}/relationships/campaigns/ | Create Tag Relationships Campaigns
[**CreateTagRelationshipsFlows**](TagsAPI.md#CreateTagRelationshipsFlows) | **Post** /api/tags/{id}/relationships/flows/ | Create Tag Relationships Flows
[**CreateTagRelationshipsLists**](TagsAPI.md#CreateTagRelationshipsLists) | **Post** /api/tags/{id}/relationships/lists/ | Create Tag Relationships Lists
[**CreateTagRelationshipsSegments**](TagsAPI.md#CreateTagRelationshipsSegments) | **Post** /api/tags/{id}/relationships/segments/ | Create Tag Relationships Segments
[**DeleteTag**](TagsAPI.md#DeleteTag) | **Delete** /api/tags/{id}/ | Delete Tag
[**DeleteTagGroup**](TagsAPI.md#DeleteTagGroup) | **Delete** /api/tag-groups/{id}/ | Delete Tag Group
[**DeleteTagRelationshipsCampaigns**](TagsAPI.md#DeleteTagRelationshipsCampaigns) | **Delete** /api/tags/{id}/relationships/campaigns/ | Delete Tag Relationships Campaigns
[**DeleteTagRelationshipsFlows**](TagsAPI.md#DeleteTagRelationshipsFlows) | **Delete** /api/tags/{id}/relationships/flows/ | Delete Tag Relationships Flows
[**DeleteTagRelationshipsLists**](TagsAPI.md#DeleteTagRelationshipsLists) | **Delete** /api/tags/{id}/relationships/lists/ | Delete Tag Relationships Lists
[**DeleteTagRelationshipsSegments**](TagsAPI.md#DeleteTagRelationshipsSegments) | **Delete** /api/tags/{id}/relationships/segments/ | Delete Tag Relationships Segments
[**GetTag**](TagsAPI.md#GetTag) | **Get** /api/tags/{id}/ | Get Tag
[**GetTagGroup**](TagsAPI.md#GetTagGroup) | **Get** /api/tag-groups/{id}/ | Get Tag Group
[**GetTagGroupRelationshipsTags**](TagsAPI.md#GetTagGroupRelationshipsTags) | **Get** /api/tag-groups/{id}/relationships/tags/ | Get Tag Group Relationships Tags
[**GetTagGroupTags**](TagsAPI.md#GetTagGroupTags) | **Get** /api/tag-groups/{id}/tags/ | Get Tag Group Tags
[**GetTagGroups**](TagsAPI.md#GetTagGroups) | **Get** /api/tag-groups/ | Get Tag Groups
[**GetTagRelationshipsCampaigns**](TagsAPI.md#GetTagRelationshipsCampaigns) | **Get** /api/tags/{id}/relationships/campaigns/ | Get Tag Relationships Campaigns
[**GetTagRelationshipsFlows**](TagsAPI.md#GetTagRelationshipsFlows) | **Get** /api/tags/{id}/relationships/flows/ | Get Tag Relationships Flows
[**GetTagRelationshipsLists**](TagsAPI.md#GetTagRelationshipsLists) | **Get** /api/tags/{id}/relationships/lists/ | Get Tag Relationships Lists
[**GetTagRelationshipsSegments**](TagsAPI.md#GetTagRelationshipsSegments) | **Get** /api/tags/{id}/relationships/segments/ | Get Tag Relationships Segments
[**GetTagRelationshipsTagGroup**](TagsAPI.md#GetTagRelationshipsTagGroup) | **Get** /api/tags/{id}/relationships/tag-group/ | Get Tag Relationships Tag Group
[**GetTagTagGroup**](TagsAPI.md#GetTagTagGroup) | **Get** /api/tags/{id}/tag-group/ | Get Tag Tag Group
[**GetTags**](TagsAPI.md#GetTags) | **Get** /api/tags/ | Get Tags
[**UpdateTag**](TagsAPI.md#UpdateTag) | **Patch** /api/tags/{id}/ | Update Tag
[**UpdateTagGroup**](TagsAPI.md#UpdateTagGroup) | **Patch** /api/tag-groups/{id}/ | Update Tag Group



## CreateTag

> PostTagResponse CreateTag(ctx).Revision(revision).TagCreateQuery(tagCreateQuery).Execute()

Create Tag



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
	tagCreateQuery := *openapiclient.NewTagCreateQuery(*openapiclient.NewTagCreateQueryResourceObject(openapiclient.TagEnum("tag"), *openapiclient.NewTagResponseObjectResourceAttributes("My Tag"))) // TagCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.CreateTag(context.Background()).Revision(revision).TagCreateQuery(tagCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.CreateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTag`: PostTagResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **tagCreateQuery** | [**TagCreateQuery**](TagCreateQuery.md) |  | 

### Return type

[**PostTagResponse**](PostTagResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTagGroup

> PostTagGroupResponse CreateTagGroup(ctx).Revision(revision).TagGroupCreateQuery(tagGroupCreateQuery).Execute()

Create Tag Group



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
	tagGroupCreateQuery := *openapiclient.NewTagGroupCreateQuery(*openapiclient.NewTagGroupCreateQueryResourceObject(openapiclient.TagGroupEnum("tag-group"), *openapiclient.NewTagGroupCreateQueryResourceObjectAttributes("My Tag Group"))) // TagGroupCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.CreateTagGroup(context.Background()).Revision(revision).TagGroupCreateQuery(tagGroupCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.CreateTagGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTagGroup`: PostTagGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.CreateTagGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **tagGroupCreateQuery** | [**TagGroupCreateQuery**](TagGroupCreateQuery.md) |  | 

### Return type

[**PostTagGroupResponse**](PostTagGroupResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTagRelationshipsCampaigns

> CreateTagRelationshipsCampaigns(ctx, id).Revision(revision).TagCampaignOp(tagCampaignOp).Execute()

Create Tag Relationships Campaigns



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
	tagCampaignOp := *openapiclient.NewTagCampaignOp([]openapiclient.TagCampaignOpDataInner{*openapiclient.NewTagCampaignOpDataInner(openapiclient.CampaignEnum("campaign"), "Id_example")}) // TagCampaignOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.CreateTagRelationshipsCampaigns(context.Background(), id).Revision(revision).TagCampaignOp(tagCampaignOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.CreateTagRelationshipsCampaigns``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateTagRelationshipsCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **tagCampaignOp** | [**TagCampaignOp**](TagCampaignOp.md) |  | 

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


## CreateTagRelationshipsFlows

> CreateTagRelationshipsFlows(ctx, id).Revision(revision).TagFlowOp(tagFlowOp).Execute()

Create Tag Relationships Flows



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
	tagFlowOp := *openapiclient.NewTagFlowOp([]openapiclient.TagFlowOpDataInner{*openapiclient.NewTagFlowOpDataInner(openapiclient.FlowEnum("flow"), "Id_example")}) // TagFlowOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.CreateTagRelationshipsFlows(context.Background(), id).Revision(revision).TagFlowOp(tagFlowOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.CreateTagRelationshipsFlows``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateTagRelationshipsFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **tagFlowOp** | [**TagFlowOp**](TagFlowOp.md) |  | 

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


## CreateTagRelationshipsLists

> CreateTagRelationshipsLists(ctx, id).Revision(revision).TagListOp(tagListOp).Execute()

Create Tag Relationships Lists



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
	tagListOp := *openapiclient.NewTagListOp([]openapiclient.TagListOpDataInner{*openapiclient.NewTagListOpDataInner(openapiclient.ListEnum("list"), "Id_example")}) // TagListOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.CreateTagRelationshipsLists(context.Background(), id).Revision(revision).TagListOp(tagListOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.CreateTagRelationshipsLists``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateTagRelationshipsListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **tagListOp** | [**TagListOp**](TagListOp.md) |  | 

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


## CreateTagRelationshipsSegments

> CreateTagRelationshipsSegments(ctx, id).Revision(revision).TagSegmentOp(tagSegmentOp).Execute()

Create Tag Relationships Segments



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
	tagSegmentOp := *openapiclient.NewTagSegmentOp([]openapiclient.TagSegmentOpDataInner{*openapiclient.NewTagSegmentOpDataInner(openapiclient.SegmentEnum("segment"), "Id_example")}) // TagSegmentOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.CreateTagRelationshipsSegments(context.Background(), id).Revision(revision).TagSegmentOp(tagSegmentOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.CreateTagRelationshipsSegments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateTagRelationshipsSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **tagSegmentOp** | [**TagSegmentOp**](TagSegmentOp.md) |  | 

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


## DeleteTag

> DeleteTag(ctx, id).Revision(revision).Execute()

Delete Tag



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
	id := "abcd1234-ef56-gh78-ij90-abcdef123456" // string | The Tag ID
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.DeleteTag(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.DeleteTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Tag ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


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


## DeleteTagGroup

> DeleteTagGroupResponse DeleteTagGroup(ctx, id).Revision(revision).Execute()

Delete Tag Group



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
	id := "zyxw9876-vu54-ts32-rq10-zyxwvu654321" // string | The Tag Group ID
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.DeleteTagGroup(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.DeleteTagGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTagGroup`: DeleteTagGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.DeleteTagGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Tag Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**DeleteTagGroupResponse**](DeleteTagGroupResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTagRelationshipsCampaigns

> DeleteTagRelationshipsCampaigns(ctx, id).Revision(revision).TagCampaignOp(tagCampaignOp).Execute()

Delete Tag Relationships Campaigns



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
	tagCampaignOp := *openapiclient.NewTagCampaignOp([]openapiclient.TagCampaignOpDataInner{*openapiclient.NewTagCampaignOpDataInner(openapiclient.CampaignEnum("campaign"), "Id_example")}) // TagCampaignOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.DeleteTagRelationshipsCampaigns(context.Background(), id).Revision(revision).TagCampaignOp(tagCampaignOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.DeleteTagRelationshipsCampaigns``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTagRelationshipsCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **tagCampaignOp** | [**TagCampaignOp**](TagCampaignOp.md) |  | 

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


## DeleteTagRelationshipsFlows

> DeleteTagRelationshipsFlows(ctx, id).Revision(revision).TagFlowOp(tagFlowOp).Execute()

Delete Tag Relationships Flows



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
	tagFlowOp := *openapiclient.NewTagFlowOp([]openapiclient.TagFlowOpDataInner{*openapiclient.NewTagFlowOpDataInner(openapiclient.FlowEnum("flow"), "Id_example")}) // TagFlowOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.DeleteTagRelationshipsFlows(context.Background(), id).Revision(revision).TagFlowOp(tagFlowOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.DeleteTagRelationshipsFlows``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTagRelationshipsFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **tagFlowOp** | [**TagFlowOp**](TagFlowOp.md) |  | 

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


## DeleteTagRelationshipsLists

> DeleteTagRelationshipsLists(ctx, id).Revision(revision).TagListOp(tagListOp).Execute()

Delete Tag Relationships Lists



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
	tagListOp := *openapiclient.NewTagListOp([]openapiclient.TagListOpDataInner{*openapiclient.NewTagListOpDataInner(openapiclient.ListEnum("list"), "Id_example")}) // TagListOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.DeleteTagRelationshipsLists(context.Background(), id).Revision(revision).TagListOp(tagListOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.DeleteTagRelationshipsLists``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTagRelationshipsListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **tagListOp** | [**TagListOp**](TagListOp.md) |  | 

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


## DeleteTagRelationshipsSegments

> DeleteTagRelationshipsSegments(ctx, id).Revision(revision).TagSegmentOp(tagSegmentOp).Execute()

Delete Tag Relationships Segments



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
	tagSegmentOp := *openapiclient.NewTagSegmentOp([]openapiclient.TagSegmentOpDataInner{*openapiclient.NewTagSegmentOpDataInner(openapiclient.SegmentEnum("segment"), "Id_example")}) // TagSegmentOp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.DeleteTagRelationshipsSegments(context.Background(), id).Revision(revision).TagSegmentOp(tagSegmentOp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.DeleteTagRelationshipsSegments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTagRelationshipsSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **tagSegmentOp** | [**TagSegmentOp**](TagSegmentOp.md) |  | 

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


## GetTag

> GetTagResponseCompoundDocument GetTag(ctx, id).Revision(revision).FieldsTagGroup(fieldsTagGroup).FieldsTag(fieldsTag).Include(include).Execute()

Get Tag



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
	id := "abcd1234-ef56-gh78-ij90-abcdef123456" // string | The Tag ID
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsTagGroup := []string{"FieldsTagGroup_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsTag := []string{"FieldsTag_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.GetTag(context.Background(), id).Revision(revision).FieldsTagGroup(fieldsTagGroup).FieldsTag(fieldsTag).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTag`: GetTagResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Tag ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsTagGroup** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsTag** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetTagResponseCompoundDocument**](GetTagResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagGroup

> GetTagGroupResponse GetTagGroup(ctx, id).Revision(revision).FieldsTagGroup(fieldsTagGroup).Execute()

Get Tag Group



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
	id := "zyxw9876-vu54-ts32-rq10-zyxwvu654321" // string | The Tag Group ID
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsTagGroup := []string{"FieldsTagGroup_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.GetTagGroup(context.Background(), id).Revision(revision).FieldsTagGroup(fieldsTagGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagGroup`: GetTagGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Tag Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsTagGroup** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetTagGroupResponse**](GetTagGroupResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagGroupRelationshipsTags

> GetTagGroupTagRelationshipsResponseCollection GetTagGroupRelationshipsTags(ctx, id).Revision(revision).Execute()

Get Tag Group Relationships Tags



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
	resp, r, err := apiClient.TagsAPI.GetTagGroupRelationshipsTags(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagGroupRelationshipsTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagGroupRelationshipsTags`: GetTagGroupTagRelationshipsResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagGroupRelationshipsTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagGroupRelationshipsTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetTagGroupTagRelationshipsResponseCollection**](GetTagGroupTagRelationshipsResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagGroupTags

> GetTagResponseCollection GetTagGroupTags(ctx, id).Revision(revision).FieldsTag(fieldsTag).Execute()

Get Tag Group Tags



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
	fieldsTag := []string{"FieldsTag_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.GetTagGroupTags(context.Background(), id).Revision(revision).FieldsTag(fieldsTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagGroupTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagGroupTags`: GetTagResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagGroupTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsTag** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetTagResponseCollection**](GetTagResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagGroups

> GetTagGroupResponseCollection GetTagGroups(ctx).Revision(revision).FieldsTagGroup(fieldsTagGroup).Filter(filter).PageCursor(pageCursor).Sort(sort).Execute()

Get Tag Groups



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
	fieldsTagGroup := []string{"FieldsTagGroup_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`name`: `contains`, `ends-with`, `equals`, `starts-with`<br>`exclusive`: `equals`<br>`default`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.GetTagGroups(context.Background()).Revision(revision).FieldsTagGroup(fieldsTagGroup).Filter(filter).PageCursor(pageCursor).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagGroups`: GetTagGroupResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsTagGroup** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;name&#x60;: &#x60;contains&#x60;, &#x60;ends-with&#x60;, &#x60;equals&#x60;, &#x60;starts-with&#x60;&lt;br&gt;&#x60;exclusive&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;default&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetTagGroupResponseCollection**](GetTagGroupResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagRelationshipsCampaigns

> GetTagCampaignRelationshipsResponseCollection GetTagRelationshipsCampaigns(ctx, id).Revision(revision).Execute()

Get Tag Relationships Campaigns



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
	resp, r, err := apiClient.TagsAPI.GetTagRelationshipsCampaigns(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagRelationshipsCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagRelationshipsCampaigns`: GetTagCampaignRelationshipsResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagRelationshipsCampaigns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRelationshipsCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetTagCampaignRelationshipsResponseCollection**](GetTagCampaignRelationshipsResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagRelationshipsFlows

> GetTagFlowRelationshipsResponseCollection GetTagRelationshipsFlows(ctx, id).Revision(revision).Execute()

Get Tag Relationships Flows



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
	resp, r, err := apiClient.TagsAPI.GetTagRelationshipsFlows(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagRelationshipsFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagRelationshipsFlows`: GetTagFlowRelationshipsResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagRelationshipsFlows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRelationshipsFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetTagFlowRelationshipsResponseCollection**](GetTagFlowRelationshipsResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagRelationshipsLists

> GetTagListRelationshipsResponseCollection GetTagRelationshipsLists(ctx, id).Revision(revision).Execute()

Get Tag Relationships Lists



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
	resp, r, err := apiClient.TagsAPI.GetTagRelationshipsLists(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagRelationshipsLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagRelationshipsLists`: GetTagListRelationshipsResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagRelationshipsLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRelationshipsListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetTagListRelationshipsResponseCollection**](GetTagListRelationshipsResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagRelationshipsSegments

> GetTagSegmentRelationshipsResponseCollection GetTagRelationshipsSegments(ctx, id).Revision(revision).Execute()

Get Tag Relationships Segments



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
	resp, r, err := apiClient.TagsAPI.GetTagRelationshipsSegments(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagRelationshipsSegments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagRelationshipsSegments`: GetTagSegmentRelationshipsResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagRelationshipsSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRelationshipsSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetTagSegmentRelationshipsResponseCollection**](GetTagSegmentRelationshipsResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagRelationshipsTagGroup

> GetTagTagGroupRelationshipsResponse GetTagRelationshipsTagGroup(ctx, id).Revision(revision).Execute()

Get Tag Relationships Tag Group



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
	resp, r, err := apiClient.TagsAPI.GetTagRelationshipsTagGroup(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagRelationshipsTagGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagRelationshipsTagGroup`: GetTagTagGroupRelationshipsResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagRelationshipsTagGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRelationshipsTagGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetTagTagGroupRelationshipsResponse**](GetTagTagGroupRelationshipsResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagTagGroup

> GetTagGroupResponse GetTagTagGroup(ctx, id).Revision(revision).FieldsTagGroup(fieldsTagGroup).Execute()

Get Tag Tag Group



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
	fieldsTagGroup := []string{"FieldsTagGroup_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.GetTagTagGroup(context.Background(), id).Revision(revision).FieldsTagGroup(fieldsTagGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagTagGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagTagGroup`: GetTagGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagTagGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagTagGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsTagGroup** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetTagGroupResponse**](GetTagGroupResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> GetTagResponseCollectionCompoundDocument GetTags(ctx).Revision(revision).FieldsTagGroup(fieldsTagGroup).FieldsTag(fieldsTag).Filter(filter).Include(include).PageCursor(pageCursor).Sort(sort).Execute()

Get Tags



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
	fieldsTagGroup := []string{"FieldsTagGroup_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsTag := []string{"FieldsTag_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`name`: `contains`, `ends-with`, `equals`, `starts-with` (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.GetTags(context.Background()).Revision(revision).FieldsTagGroup(fieldsTagGroup).FieldsTag(fieldsTag).Filter(filter).Include(include).PageCursor(pageCursor).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags`: GetTagResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsTagGroup** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsTag** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;name&#x60;: &#x60;contains&#x60;, &#x60;ends-with&#x60;, &#x60;equals&#x60;, &#x60;starts-with&#x60; | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetTagResponseCollectionCompoundDocument**](GetTagResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTag

> UpdateTag(ctx, id).Revision(revision).TagUpdateQuery(tagUpdateQuery).Execute()

Update Tag



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
	id := "abcd1234-ef56-gh78-ij90-abcdef123456" // string | The Tag ID
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	tagUpdateQuery := *openapiclient.NewTagUpdateQuery(*openapiclient.NewTagUpdateQueryResourceObject(openapiclient.TagEnum("tag"), "abcd1234-ef56-gh78-ij90-abcdef123456", *openapiclient.NewTagResponseObjectResourceAttributes("My Tag"))) // TagUpdateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.UpdateTag(context.Background(), id).Revision(revision).TagUpdateQuery(tagUpdateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.UpdateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Tag ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **tagUpdateQuery** | [**TagUpdateQuery**](TagUpdateQuery.md) |  | 

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


## UpdateTagGroup

> PatchTagGroupResponse UpdateTagGroup(ctx, id).Revision(revision).TagGroupUpdateQuery(tagGroupUpdateQuery).Execute()

Update Tag Group



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
	id := "zyxw9876-vu54-ts32-rq10-zyxwvu654321" // string | The Tag Group ID
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	tagGroupUpdateQuery := *openapiclient.NewTagGroupUpdateQuery(*openapiclient.NewTagGroupUpdateQueryResourceObject(openapiclient.TagGroupEnum("tag-group"), "zyxw9876-vu54-ts32-rq10-zyxwvu654321", *openapiclient.NewTagGroupUpdateQueryResourceObjectAttributes("My Tag Group"))) // TagGroupUpdateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.UpdateTagGroup(context.Background(), id).Revision(revision).TagGroupUpdateQuery(tagGroupUpdateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.UpdateTagGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTagGroup`: PatchTagGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.UpdateTagGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Tag Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **tagGroupUpdateQuery** | [**TagGroupUpdateQuery**](TagGroupUpdateQuery.md) |  | 

### Return type

[**PatchTagGroupResponse**](PatchTagGroupResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

