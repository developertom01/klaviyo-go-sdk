# \EventsAPI

All URIs are relative to *https://a.klaviyo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEvent**](EventsAPI.md#CreateEvent) | **Post** /api/events/ | Create Event
[**GetEvent**](EventsAPI.md#GetEvent) | **Get** /api/events/{id}/ | Get Event
[**GetEventMetric**](EventsAPI.md#GetEventMetric) | **Get** /api/events/{id}/metric/ | Get Event Metric
[**GetEventProfile**](EventsAPI.md#GetEventProfile) | **Get** /api/events/{id}/profile/ | Get Event Profile
[**GetEventRelationshipsMetric**](EventsAPI.md#GetEventRelationshipsMetric) | **Get** /api/events/{id}/relationships/metric/ | Get Event Relationships Metric
[**GetEventRelationshipsProfile**](EventsAPI.md#GetEventRelationshipsProfile) | **Get** /api/events/{id}/relationships/profile/ | Get Event Relationships Profile
[**GetEvents**](EventsAPI.md#GetEvents) | **Get** /api/events/ | Get Events



## CreateEvent

> CreateEvent(ctx).Revision(revision).EventCreateQueryV2(eventCreateQueryV2).Execute()

Create Event



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
	eventCreateQueryV2 := *openapiclient.NewEventCreateQueryV2(*openapiclient.NewEventCreateQueryV2ResourceObject(openapiclient.EventEnum("event"), *openapiclient.NewEventCreateQueryV2ResourceObjectAttributes(map[string]interface{}({"Brand":"Kids Book","Categories":["Fiction","Children"],"ProductID":1111,"ProductName":"Winnie the Pooh","$extra":{"URL":"http://www.example.com/path/to/product","ImageURL":"http://www.example.com/path/to/product/image.png"}}), *openapiclient.NewEventCreateQueryV2ResourceObjectAttributesMetric(*openapiclient.NewMetricCreateQueryResourceObject(openapiclient.MetricEnum("metric"), *openapiclient.NewMetricCreateQueryResourceObjectAttributes("Viewed Product"))), *openapiclient.NewEventCreateQueryV2ResourceObjectAttributesProfile(*openapiclient.NewOnsiteProfileCreateQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewOnsiteProfileCreateQueryResourceObjectAttributes()))))) // EventCreateQueryV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.CreateEvent(context.Background()).Revision(revision).EventCreateQueryV2(eventCreateQueryV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.CreateEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **eventCreateQueryV2** | [**EventCreateQueryV2**](EventCreateQueryV2.md) |  | 

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


## GetEvent

> GetEventResponseCompoundDocument GetEvent(ctx, id).Revision(revision).FieldsEvent(fieldsEvent).FieldsMetric(fieldsMetric).FieldsProfile(fieldsProfile).Include(include).Execute()

Get Event



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
	id := "id_example" // string | ID of the event
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsEvent := []string{"FieldsEvent_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsMetric := []string{"FieldsMetric_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsProfile := []string{"FieldsProfile_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetEvent(context.Background(), id).Revision(revision).FieldsEvent(fieldsEvent).FieldsMetric(fieldsMetric).FieldsProfile(fieldsProfile).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvent`: GetEventResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the event | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsEvent** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsMetric** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsProfile** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetEventResponseCompoundDocument**](GetEventResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventMetric

> GetMetricResponse GetEventMetric(ctx, id).Revision(revision).FieldsMetric(fieldsMetric).Execute()

Get Event Metric



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
	fieldsMetric := []string{"FieldsMetric_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetEventMetric(context.Background(), id).Revision(revision).FieldsMetric(fieldsMetric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventMetric`: GetMetricResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsMetric** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetMetricResponse**](GetMetricResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventProfile

> GetProfileResponse GetEventProfile(ctx, id).Revision(revision).AdditionalFieldsProfile(additionalFieldsProfile).FieldsProfile(fieldsProfile).Execute()

Get Event Profile



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetEventProfile(context.Background(), id).Revision(revision).AdditionalFieldsProfile(additionalFieldsProfile).FieldsProfile(fieldsProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventProfile`: GetProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **additionalFieldsProfile** | **[]string** | Request additional fields not included by default in the response. Supported values: &#39;subscriptions&#39;, &#39;predictive_analytics&#39; | 
 **fieldsProfile** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetProfileResponse**](GetProfileResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventRelationshipsMetric

> GetEventMetricsRelationshipListResponse GetEventRelationshipsMetric(ctx, id).Revision(revision).Execute()

Get Event Relationships Metric



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
	resp, r, err := apiClient.EventsAPI.GetEventRelationshipsMetric(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventRelationshipsMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventRelationshipsMetric`: GetEventMetricsRelationshipListResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventRelationshipsMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRelationshipsMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetEventMetricsRelationshipListResponse**](GetEventMetricsRelationshipListResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventRelationshipsProfile

> GetEventProfilesRelationshipListResponse GetEventRelationshipsProfile(ctx, id).Revision(revision).Execute()

Get Event Relationships Profile



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
	resp, r, err := apiClient.EventsAPI.GetEventRelationshipsProfile(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventRelationshipsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventRelationshipsProfile`: GetEventProfilesRelationshipListResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventRelationshipsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRelationshipsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetEventProfilesRelationshipListResponse**](GetEventProfilesRelationshipListResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> GetEventResponseCollectionCompoundDocument GetEvents(ctx).Revision(revision).FieldsEvent(fieldsEvent).FieldsMetric(fieldsMetric).FieldsProfile(fieldsProfile).Filter(filter).Include(include).PageCursor(pageCursor).Sort(sort).Execute()

Get Events



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
	fieldsEvent := []string{"FieldsEvent_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsMetric := []string{"FieldsMetric_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsProfile := []string{"FieldsProfile_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`metric_id`: `equals`<br>`profile_id`: `equals`<br>`profile`: `has`<br>`datetime`: `greater-or-equal`, `greater-than`, `less-or-equal`, `less-than`<br>`timestamp`: `greater-or-equal`, `greater-than`, `less-or-equal`, `less-than` (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetEvents(context.Background()).Revision(revision).FieldsEvent(fieldsEvent).FieldsMetric(fieldsMetric).FieldsProfile(fieldsProfile).Filter(filter).Include(include).PageCursor(pageCursor).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvents`: GetEventResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsEvent** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsMetric** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsProfile** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;metric_id&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;profile_id&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;profile&#x60;: &#x60;has&#x60;&lt;br&gt;&#x60;datetime&#x60;: &#x60;greater-or-equal&#x60;, &#x60;greater-than&#x60;, &#x60;less-or-equal&#x60;, &#x60;less-than&#x60;&lt;br&gt;&#x60;timestamp&#x60;: &#x60;greater-or-equal&#x60;, &#x60;greater-than&#x60;, &#x60;less-or-equal&#x60;, &#x60;less-than&#x60; | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetEventResponseCollectionCompoundDocument**](GetEventResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

