# \MetricsAPI

All URIs are relative to *https://a.klaviyo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetric**](MetricsAPI.md#GetMetric) | **Get** /api/metrics/{id}/ | Get Metric
[**GetMetrics**](MetricsAPI.md#GetMetrics) | **Get** /api/metrics/ | Get Metrics
[**QueryMetricAggregates**](MetricsAPI.md#QueryMetricAggregates) | **Post** /api/metric-aggregates/ | Query Metric Aggregates



## GetMetric

> GetMetricResponse GetMetric(ctx, id).Revision(revision).FieldsMetric(fieldsMetric).Execute()

Get Metric



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
	id := "id_example" // string | Metric ID
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsMetric := []string{"FieldsMetric_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetMetric(context.Background(), id).Revision(revision).FieldsMetric(fieldsMetric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetric`: GetMetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Metric ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricRequest struct via the builder pattern


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


## GetMetrics

> GetMetricResponseCollection GetMetrics(ctx).Revision(revision).FieldsMetric(fieldsMetric).Filter(filter).PageCursor(pageCursor).Execute()

Get Metrics



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
	fieldsMetric := []string{"FieldsMetric_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	filter := "filter_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`integration.name`: `equals`<br>`integration.category`: `equals` (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetMetrics(context.Background()).Revision(revision).FieldsMetric(fieldsMetric).Filter(filter).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetrics`: GetMetricResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsMetric** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;integration.name&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;integration.category&#x60;: &#x60;equals&#x60; | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**GetMetricResponseCollection**](GetMetricResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryMetricAggregates

> PostMetricAggregateResponse QueryMetricAggregates(ctx).Revision(revision).MetricAggregateQuery(metricAggregateQuery).Execute()

Query Metric Aggregates



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
	metricAggregateQuery := *openapiclient.NewMetricAggregateQuery(*openapiclient.NewMetricAggregateQueryResourceObject(openapiclient.MetricAggregateEnum("metric-aggregate"), *openapiclient.NewMetricAggregateQueryResourceObjectAttributes("0rG4eQ", []string{"Measurements_example"}, []string{"Filter_example"}))) // MetricAggregateQuery | Retrieve Metric Aggregations

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.QueryMetricAggregates(context.Background()).Revision(revision).MetricAggregateQuery(metricAggregateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.QueryMetricAggregates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryMetricAggregates`: PostMetricAggregateResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.QueryMetricAggregates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryMetricAggregatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **metricAggregateQuery** | [**MetricAggregateQuery**](MetricAggregateQuery.md) | Retrieve Metric Aggregations | 

### Return type

[**PostMetricAggregateResponse**](PostMetricAggregateResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

