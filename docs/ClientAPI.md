# \ClientAPI

All URIs are relative to *https://a.klaviyo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateClientEvents**](ClientAPI.md#BulkCreateClientEvents) | **Post** /client/event-bulk-create/ | Bulk Create Client Events
[**CreateClientBackInStockSubscription**](ClientAPI.md#CreateClientBackInStockSubscription) | **Post** /client/back-in-stock-subscriptions/ | Create Client Back In Stock Subscription
[**CreateClientEvent**](ClientAPI.md#CreateClientEvent) | **Post** /client/events/ | Create Client Event
[**CreateClientProfile**](ClientAPI.md#CreateClientProfile) | **Post** /client/profiles/ | Create or Update Client Profile
[**CreateClientPushToken**](ClientAPI.md#CreateClientPushToken) | **Post** /client/push-tokens/ | Create or Update Client Push Token
[**CreateClientSubscription**](ClientAPI.md#CreateClientSubscription) | **Post** /client/subscriptions/ | Create Client Subscription
[**UnregisterClientPushToken**](ClientAPI.md#UnregisterClientPushToken) | **Post** /client/push-token-unregister/ | Unregister Client Push Token



## BulkCreateClientEvents

> BulkCreateClientEvents(ctx).CompanyId(companyId).Revision(revision).EventsBulkCreateQuery(eventsBulkCreateQuery).Execute()

Bulk Create Client Events



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
	companyId := "PUBLIC_API_KEY" // string | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	eventsBulkCreateQuery := *openapiclient.NewEventsBulkCreateQuery(*openapiclient.NewEventsBulkCreateQueryResourceObject(openapiclient.EventBulkCreateEnum("event-bulk-create"), *openapiclient.NewEventsBulkCreateQueryResourceObjectAttributes(*openapiclient.NewEventCreateQueryV2ResourceObjectAttributesProfile(*openapiclient.NewOnsiteProfileCreateQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewOnsiteProfileCreateQueryResourceObjectAttributes())), *openapiclient.NewEventsBulkCreateQueryResourceObjectAttributesEvents([]openapiclient.BaseEventCreateQueryResourceObject{*openapiclient.NewBaseEventCreateQueryResourceObject(openapiclient.EventEnum("event"), *openapiclient.NewBaseEventCreateQueryResourceObjectAttributes(map[string]interface{}({"Brand":"Kids Book","Categories":["Fiction","Children"],"ProductID":1111,"ProductName":"Winnie the Pooh","$extra":{"URL":"http://www.example.com/path/to/product","ImageURL":"http://www.example.com/path/to/product/image.png"}}), *openapiclient.NewEventCreateQueryV2ResourceObjectAttributesMetric(*openapiclient.NewMetricCreateQueryResourceObject(openapiclient.MetricEnum("metric"), *openapiclient.NewMetricCreateQueryResourceObjectAttributes("Viewed Product")))))})))) // EventsBulkCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientAPI.BulkCreateClientEvents(context.Background()).CompanyId(companyId).Revision(revision).EventsBulkCreateQuery(eventsBulkCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.BulkCreateClientEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateClientEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details. | 
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **eventsBulkCreateQuery** | [**EventsBulkCreateQuery**](EventsBulkCreateQuery.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClientBackInStockSubscription

> CreateClientBackInStockSubscription(ctx).CompanyId(companyId).Revision(revision).ClientBISSubscriptionCreateQuery(clientBISSubscriptionCreateQuery).Execute()

Create Client Back In Stock Subscription



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
	companyId := "PUBLIC_API_KEY" // string | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	clientBISSubscriptionCreateQuery := *openapiclient.NewClientBISSubscriptionCreateQuery(*openapiclient.NewClientBISSubscriptionCreateQueryResourceObject(openapiclient.BackInStockSubscriptionEnum("back-in-stock-subscription"), *openapiclient.NewClientBISSubscriptionCreateQueryResourceObjectAttributes([]string{"Channels_example"}, *openapiclient.NewClientBISSubscriptionCreateQueryResourceObjectAttributesProfile(*openapiclient.NewProfileIdentifierDTOResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewProfileIdentifierDTOResourceObjectAttributes()))), *openapiclient.NewServerBISSubscriptionCreateQueryResourceObjectRelationships(*openapiclient.NewServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariant(*openapiclient.NewServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData(openapiclient.CatalogVariantEnum("catalog-variant"), "$custom:::$default:::SAMPLE-DATA-ITEM-1-VARIANT-MEDIUM"))))) // ClientBISSubscriptionCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientAPI.CreateClientBackInStockSubscription(context.Background()).CompanyId(companyId).Revision(revision).ClientBISSubscriptionCreateQuery(clientBISSubscriptionCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateClientBackInStockSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientBackInStockSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details. | 
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **clientBISSubscriptionCreateQuery** | [**ClientBISSubscriptionCreateQuery**](ClientBISSubscriptionCreateQuery.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClientEvent

> CreateClientEvent(ctx).CompanyId(companyId).Revision(revision).EventCreateQueryV2(eventCreateQueryV2).Execute()

Create Client Event



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
	companyId := "PUBLIC_API_KEY" // string | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	eventCreateQueryV2 := *openapiclient.NewEventCreateQueryV2(*openapiclient.NewEventCreateQueryV2ResourceObject(openapiclient.EventEnum("event"), *openapiclient.NewEventCreateQueryV2ResourceObjectAttributes(map[string]interface{}({"Brand":"Kids Book","Categories":["Fiction","Children"],"ProductID":1111,"ProductName":"Winnie the Pooh","$extra":{"URL":"http://www.example.com/path/to/product","ImageURL":"http://www.example.com/path/to/product/image.png"}}), *openapiclient.NewEventCreateQueryV2ResourceObjectAttributesMetric(*openapiclient.NewMetricCreateQueryResourceObject(openapiclient.MetricEnum("metric"), *openapiclient.NewMetricCreateQueryResourceObjectAttributes("Viewed Product"))), *openapiclient.NewEventCreateQueryV2ResourceObjectAttributesProfile(*openapiclient.NewOnsiteProfileCreateQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewOnsiteProfileCreateQueryResourceObjectAttributes()))))) // EventCreateQueryV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientAPI.CreateClientEvent(context.Background()).CompanyId(companyId).Revision(revision).EventCreateQueryV2(eventCreateQueryV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateClientEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details. | 
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **eventCreateQueryV2** | [**EventCreateQueryV2**](EventCreateQueryV2.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClientProfile

> CreateClientProfile(ctx).CompanyId(companyId).Revision(revision).OnsiteProfileCreateQuery(onsiteProfileCreateQuery).Execute()

Create or Update Client Profile



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
	companyId := "PUBLIC_API_KEY" // string | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	onsiteProfileCreateQuery := *openapiclient.NewOnsiteProfileCreateQuery(*openapiclient.NewOnsiteProfileCreateQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewOnsiteProfileCreateQueryResourceObjectAttributes())) // OnsiteProfileCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientAPI.CreateClientProfile(context.Background()).CompanyId(companyId).Revision(revision).OnsiteProfileCreateQuery(onsiteProfileCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateClientProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details. | 
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **onsiteProfileCreateQuery** | [**OnsiteProfileCreateQuery**](OnsiteProfileCreateQuery.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClientPushToken

> CreateClientPushToken(ctx).CompanyId(companyId).Revision(revision).PushTokenCreateQuery(pushTokenCreateQuery).Execute()

Create or Update Client Push Token



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
	companyId := "PUBLIC_API_KEY" // string | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	pushTokenCreateQuery := *openapiclient.NewPushTokenCreateQuery(*openapiclient.NewPushTokenCreateQueryResourceObject(openapiclient.PushTokenEnum("push-token"), *openapiclient.NewPushTokenCreateQueryResourceObjectAttributes("1234567890", "Platform_example", "APNs", *openapiclient.NewPushTokenCreateQueryResourceObjectAttributesProfile(*openapiclient.NewProfileUpsertQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewProfileUpsertQueryResourceObjectAttributes()))))) // PushTokenCreateQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientAPI.CreateClientPushToken(context.Background()).CompanyId(companyId).Revision(revision).PushTokenCreateQuery(pushTokenCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateClientPushToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientPushTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details. | 
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **pushTokenCreateQuery** | [**PushTokenCreateQuery**](PushTokenCreateQuery.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClientSubscription

> CreateClientSubscription(ctx).CompanyId(companyId).Revision(revision).OnsiteSubscriptionCreateQuery(onsiteSubscriptionCreateQuery).Execute()

Create Client Subscription



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
	companyId := "PUBLIC_API_KEY" // string | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	onsiteSubscriptionCreateQuery := *openapiclient.NewOnsiteSubscriptionCreateQuery(*openapiclient.NewOnsiteSubscriptionCreateQueryResourceObject(openapiclient.SubscriptionEnum("subscription"), *openapiclient.NewOnsiteSubscriptionCreateQueryResourceObjectAttributes(*openapiclient.NewOnsiteSubscriptionCreateQueryResourceObjectAttributesProfile(*openapiclient.NewProfileUpsertQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewProfileUpsertQueryResourceObjectAttributes()))), *openapiclient.NewOnsiteSubscriptionCreateQueryResourceObjectRelationships(*openapiclient.NewOnsiteSubscriptionCreateQueryResourceObjectRelationshipsList(*openapiclient.NewOnsiteSubscriptionCreateQueryResourceObjectRelationshipsListData(openapiclient.ListEnum("list"), "Y6nRLr"))))) // OnsiteSubscriptionCreateQuery | Creates a subscription and consent records for Email and or SMS channels based on the provided email and phone_number attributes respectively. One of either email or phone_number must be provided. To create a subscription and consent record for only one channel but still include the other channel as a profile property the consent channel can be provided as a top level attribute and the other channel can be included in the properties object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientAPI.CreateClientSubscription(context.Background()).CompanyId(companyId).Revision(revision).OnsiteSubscriptionCreateQuery(onsiteSubscriptionCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateClientSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details. | 
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **onsiteSubscriptionCreateQuery** | [**OnsiteSubscriptionCreateQuery**](OnsiteSubscriptionCreateQuery.md) | Creates a subscription and consent records for Email and or SMS channels based on the provided email and phone_number attributes respectively. One of either email or phone_number must be provided. To create a subscription and consent record for only one channel but still include the other channel as a profile property the consent channel can be provided as a top level attribute and the other channel can be included in the properties object. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterClientPushToken

> UnregisterClientPushToken(ctx).CompanyId(companyId).Revision(revision).PushTokenUnregisterQuery(pushTokenUnregisterQuery).Execute()

Unregister Client Push Token



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
	companyId := "PUBLIC_API_KEY" // string | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details.
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	pushTokenUnregisterQuery := *openapiclient.NewPushTokenUnregisterQuery(*openapiclient.NewPushTokenUnregisterQueryResourceObject(openapiclient.PushTokenUnregisterEnum("push-token-unregister"), *openapiclient.NewPushTokenUnregisterQueryResourceObjectAttributes("1234567890", "Platform_example", *openapiclient.NewPushTokenCreateQueryResourceObjectAttributesProfile(*openapiclient.NewProfileUpsertQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewProfileUpsertQueryResourceObjectAttributes()))))) // PushTokenUnregisterQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientAPI.UnregisterClientPushToken(context.Background()).CompanyId(companyId).Revision(revision).PushTokenUnregisterQuery(pushTokenUnregisterQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.UnregisterClientPushToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterClientPushTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Your Public API Key / Site ID. See [this article](https://help.klaviyo.com/hc/en-us/articles/115005062267) for more details. | 
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **pushTokenUnregisterQuery** | [**PushTokenUnregisterQuery**](PushTokenUnregisterQuery.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

