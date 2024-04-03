# \DataPrivacyAPI

All URIs are relative to *https://a.klaviyo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestProfileDeletion**](DataPrivacyAPI.md#RequestProfileDeletion) | **Post** /api/data-privacy-deletion-jobs/ | Request Profile Deletion



## RequestProfileDeletion

> RequestProfileDeletion(ctx).Revision(revision).DataPrivacyCreateDeletionJobQuery(dataPrivacyCreateDeletionJobQuery).Execute()

Request Profile Deletion



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
	dataPrivacyCreateDeletionJobQuery := *openapiclient.NewDataPrivacyCreateDeletionJobQuery(*openapiclient.NewDataPrivacyCreateDeletionJobQueryResourceObject(openapiclient.DataPrivacyDeletionJobEnum("data-privacy-deletion-job"), *openapiclient.NewDataPrivacyCreateDeletionJobQueryResourceObjectAttributes(*openapiclient.NewDataPrivacyCreateDeletionJobQueryResourceObjectAttributesProfile(*openapiclient.NewDataPrivacyProfileQueryResourceObject(openapiclient.ProfileEnum("profile"), *openapiclient.NewDataPrivacyProfileQueryResourceObjectAttributes()))))) // DataPrivacyCreateDeletionJobQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataPrivacyAPI.RequestProfileDeletion(context.Background()).Revision(revision).DataPrivacyCreateDeletionJobQuery(dataPrivacyCreateDeletionJobQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataPrivacyAPI.RequestProfileDeletion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestProfileDeletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **dataPrivacyCreateDeletionJobQuery** | [**DataPrivacyCreateDeletionJobQuery**](DataPrivacyCreateDeletionJobQuery.md) |  | 

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

