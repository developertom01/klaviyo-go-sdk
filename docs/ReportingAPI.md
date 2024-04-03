# \ReportingAPI

All URIs are relative to *https://a.klaviyo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryCampaignValues**](ReportingAPI.md#QueryCampaignValues) | **Post** /api/campaign-values-reports/ | Query Campaign Values
[**QueryFlowSeries**](ReportingAPI.md#QueryFlowSeries) | **Post** /api/flow-series-reports/ | Query Flow Series
[**QueryFlowValues**](ReportingAPI.md#QueryFlowValues) | **Post** /api/flow-values-reports/ | Query Flow Values



## QueryCampaignValues

> PostCampaignValuesResponseDTO QueryCampaignValues(ctx).Revision(revision).CampaignValuesRequestDTO(campaignValuesRequestDTO).PageCursor(pageCursor).Execute()

Query Campaign Values



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	campaignValuesRequestDTO := *openapiclient.NewCampaignValuesRequestDTO(*openapiclient.NewCampaignValuesRequestDTOResourceObject(openapiclient.CampaignValuesReportEnum("campaign-values-report"), *openapiclient.NewCampaignValuesRequestDTOResourceObjectAttributes([]string{"Statistics_example"}, openapiclient.CampaignValuesRequestDTOResourceObject_attributes_timeframe{CustomTimeframe: openapiclient.NewCustomTimeframe(time.Now(), time.Now())}, "RESQ6t"))) // CampaignValuesRequestDTO | 
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingAPI.QueryCampaignValues(context.Background()).Revision(revision).CampaignValuesRequestDTO(campaignValuesRequestDTO).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.QueryCampaignValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryCampaignValues`: PostCampaignValuesResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.QueryCampaignValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryCampaignValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **campaignValuesRequestDTO** | [**CampaignValuesRequestDTO**](CampaignValuesRequestDTO.md) |  | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**PostCampaignValuesResponseDTO**](PostCampaignValuesResponseDTO.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryFlowSeries

> PostFlowSeriesResponseDTO QueryFlowSeries(ctx).Revision(revision).FlowSeriesRequestDTO(flowSeriesRequestDTO).PageCursor(pageCursor).Execute()

Query Flow Series



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	flowSeriesRequestDTO := *openapiclient.NewFlowSeriesRequestDTO(*openapiclient.NewFlowSeriesRequestDTOResourceObject(openapiclient.FlowSeriesReportEnum("flow-series-report"), *openapiclient.NewFlowSeriesRequestDTOResourceObjectAttributes([]string{"Statistics_example"}, openapiclient.CampaignValuesRequestDTOResourceObject_attributes_timeframe{CustomTimeframe: openapiclient.NewCustomTimeframe(time.Now(), time.Now())}, "weekly", "RESQ6t"))) // FlowSeriesRequestDTO | 
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingAPI.QueryFlowSeries(context.Background()).Revision(revision).FlowSeriesRequestDTO(flowSeriesRequestDTO).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.QueryFlowSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryFlowSeries`: PostFlowSeriesResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.QueryFlowSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryFlowSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **flowSeriesRequestDTO** | [**FlowSeriesRequestDTO**](FlowSeriesRequestDTO.md) |  | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**PostFlowSeriesResponseDTO**](PostFlowSeriesResponseDTO.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryFlowValues

> PostFlowValuesResponseDTO QueryFlowValues(ctx).Revision(revision).FlowValuesRequestDTO(flowValuesRequestDTO).PageCursor(pageCursor).Execute()

Query Flow Values



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func main() {
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	flowValuesRequestDTO := *openapiclient.NewFlowValuesRequestDTO(*openapiclient.NewFlowValuesRequestDTOResourceObject(openapiclient.FlowValuesReportEnum("flow-values-report"), *openapiclient.NewFlowValuesRequestDTOResourceObjectAttributes([]string{"Statistics_example"}, openapiclient.CampaignValuesRequestDTOResourceObject_attributes_timeframe{CustomTimeframe: openapiclient.NewCustomTimeframe(time.Now(), time.Now())}, "RESQ6t"))) // FlowValuesRequestDTO | 
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingAPI.QueryFlowValues(context.Background()).Revision(revision).FlowValuesRequestDTO(flowValuesRequestDTO).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.QueryFlowValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryFlowValues`: PostFlowValuesResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.QueryFlowValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryFlowValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **flowValuesRequestDTO** | [**FlowValuesRequestDTO**](FlowValuesRequestDTO.md) |  | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 

### Return type

[**PostFlowValuesResponseDTO**](PostFlowValuesResponseDTO.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

