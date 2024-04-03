# \CampaignsAPI

All URIs are relative to *https://a.klaviyo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaign**](CampaignsAPI.md#CreateCampaign) | **Post** /api/campaigns/ | Create Campaign
[**CreateCampaignClone**](CampaignsAPI.md#CreateCampaignClone) | **Post** /api/campaign-clone/ | Create Campaign Clone
[**CreateCampaignMessageAssignTemplate**](CampaignsAPI.md#CreateCampaignMessageAssignTemplate) | **Post** /api/campaign-message-assign-template/ | Assign Campaign Message Template
[**CreateCampaignRecipientEstimationJob**](CampaignsAPI.md#CreateCampaignRecipientEstimationJob) | **Post** /api/campaign-recipient-estimation-jobs/ | Create Campaign Recipient Estimation Job
[**CreateCampaignSendJob**](CampaignsAPI.md#CreateCampaignSendJob) | **Post** /api/campaign-send-jobs/ | Create Campaign Send Job
[**DeleteCampaign**](CampaignsAPI.md#DeleteCampaign) | **Delete** /api/campaigns/{id}/ | Delete Campaign
[**GetCampaign**](CampaignsAPI.md#GetCampaign) | **Get** /api/campaigns/{id}/ | Get Campaign
[**GetCampaignCampaignMessages**](CampaignsAPI.md#GetCampaignCampaignMessages) | **Get** /api/campaigns/{id}/campaign-messages/ | Get Campaign Campaign Messages
[**GetCampaignMessage**](CampaignsAPI.md#GetCampaignMessage) | **Get** /api/campaign-messages/{id}/ | Get Campaign Message
[**GetCampaignMessageCampaign**](CampaignsAPI.md#GetCampaignMessageCampaign) | **Get** /api/campaign-messages/{id}/campaign/ | Get Campaign Message Campaign
[**GetCampaignMessageRelationshipsCampaign**](CampaignsAPI.md#GetCampaignMessageRelationshipsCampaign) | **Get** /api/campaign-messages/{id}/relationships/campaign/ | Get Campaign Message Relationships Campaign
[**GetCampaignMessageRelationshipsTemplate**](CampaignsAPI.md#GetCampaignMessageRelationshipsTemplate) | **Get** /api/campaign-messages/{id}/relationships/template/ | Get Campaign Message Relationships Template
[**GetCampaignMessageTemplate**](CampaignsAPI.md#GetCampaignMessageTemplate) | **Get** /api/campaign-messages/{id}/template/ | Get Campaign Message Template
[**GetCampaignRecipientEstimation**](CampaignsAPI.md#GetCampaignRecipientEstimation) | **Get** /api/campaign-recipient-estimations/{id}/ | Get Campaign Recipient Estimation
[**GetCampaignRecipientEstimationJob**](CampaignsAPI.md#GetCampaignRecipientEstimationJob) | **Get** /api/campaign-recipient-estimation-jobs/{id}/ | Get Campaign Recipient Estimation Job
[**GetCampaignRelationshipsCampaignMessages**](CampaignsAPI.md#GetCampaignRelationshipsCampaignMessages) | **Get** /api/campaigns/{id}/relationships/campaign-messages/ | Get Campaign Relationships Campaign Messages
[**GetCampaignRelationshipsTags**](CampaignsAPI.md#GetCampaignRelationshipsTags) | **Get** /api/campaigns/{id}/relationships/tags/ | Get Campaign Relationships Tags
[**GetCampaignSendJob**](CampaignsAPI.md#GetCampaignSendJob) | **Get** /api/campaign-send-jobs/{id}/ | Get Campaign Send Job
[**GetCampaignTags**](CampaignsAPI.md#GetCampaignTags) | **Get** /api/campaigns/{id}/tags/ | Get Campaign Tags
[**GetCampaigns**](CampaignsAPI.md#GetCampaigns) | **Get** /api/campaigns/ | Get Campaigns
[**UpdateCampaign**](CampaignsAPI.md#UpdateCampaign) | **Patch** /api/campaigns/{id}/ | Update Campaign
[**UpdateCampaignMessage**](CampaignsAPI.md#UpdateCampaignMessage) | **Patch** /api/campaign-messages/{id}/ | Update Campaign Message
[**UpdateCampaignSendJob**](CampaignsAPI.md#UpdateCampaignSendJob) | **Patch** /api/campaign-send-jobs/{id}/ | Update Campaign Send Job



## CreateCampaign

> PostCampaignResponse CreateCampaign(ctx).Revision(revision).CampaignCreateQuery(campaignCreateQuery).Execute()

Create Campaign



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
	campaignCreateQuery := *openapiclient.NewCampaignCreateQuery(*openapiclient.NewCampaignCreateQueryResourceObject(openapiclient.CampaignEnum("campaign"), *openapiclient.NewCampaignCreateQueryResourceObjectAttributes("My new campaign", *openapiclient.NewAudiencesSubObject(), *openapiclient.NewCampaignCreateQueryResourceObjectAttributesCampaignMessages([]openapiclient.CampaignMessageCreateQueryResourceObject{*openapiclient.NewCampaignMessageCreateQueryResourceObject(openapiclient.CampaignMessageEnum("campaign-message"), *openapiclient.NewCampaignMessageCreateQueryResourceObjectAttributes("email"))})))) // CampaignCreateQuery | Creates a campaign from parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.CreateCampaign(context.Background()).Revision(revision).CampaignCreateQuery(campaignCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.CreateCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCampaign`: PostCampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.CreateCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **campaignCreateQuery** | [**CampaignCreateQuery**](CampaignCreateQuery.md) | Creates a campaign from parameters | 

### Return type

[**PostCampaignResponse**](PostCampaignResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaignClone

> PostCampaignResponse CreateCampaignClone(ctx).Revision(revision).CampaignCloneQuery(campaignCloneQuery).Execute()

Create Campaign Clone



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
	campaignCloneQuery := *openapiclient.NewCampaignCloneQuery(*openapiclient.NewCampaignCloneQueryResourceObject(openapiclient.CampaignEnum("campaign"), "Id_example", *openapiclient.NewCampaignCloneQueryResourceObjectAttributes())) // CampaignCloneQuery | Clones a campaign from an existing campaign

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.CreateCampaignClone(context.Background()).Revision(revision).CampaignCloneQuery(campaignCloneQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.CreateCampaignClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCampaignClone`: PostCampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.CreateCampaignClone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **campaignCloneQuery** | [**CampaignCloneQuery**](CampaignCloneQuery.md) | Clones a campaign from an existing campaign | 

### Return type

[**PostCampaignResponse**](PostCampaignResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaignMessageAssignTemplate

> PostCampaignMessageResponse CreateCampaignMessageAssignTemplate(ctx).Revision(revision).CampaignMessageAssignTemplateQuery(campaignMessageAssignTemplateQuery).Execute()

Assign Campaign Message Template



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
	campaignMessageAssignTemplateQuery := *openapiclient.NewCampaignMessageAssignTemplateQuery(*openapiclient.NewCampaignMessageAssignTemplateQueryResourceObject(openapiclient.CampaignMessageEnum("campaign-message"), "Id_example", *openapiclient.NewCampaignMessageAssignTemplateQueryResourceObjectRelationships(*openapiclient.NewCampaignMessageAssignTemplateQueryResourceObjectRelationshipsTemplate(*openapiclient.NewCampaignMessageAssignTemplateQueryResourceObjectRelationshipsTemplateData(openapiclient.TemplateEnum("template"), "RipRmi"))))) // CampaignMessageAssignTemplateQuery | Takes a reusable template, clones it, and assigns the non-reusable clone to the message.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.CreateCampaignMessageAssignTemplate(context.Background()).Revision(revision).CampaignMessageAssignTemplateQuery(campaignMessageAssignTemplateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.CreateCampaignMessageAssignTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCampaignMessageAssignTemplate`: PostCampaignMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.CreateCampaignMessageAssignTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignMessageAssignTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **campaignMessageAssignTemplateQuery** | [**CampaignMessageAssignTemplateQuery**](CampaignMessageAssignTemplateQuery.md) | Takes a reusable template, clones it, and assigns the non-reusable clone to the message. | 

### Return type

[**PostCampaignMessageResponse**](PostCampaignMessageResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaignRecipientEstimationJob

> PostCampaignRecipientEstimationJobResponse CreateCampaignRecipientEstimationJob(ctx).Revision(revision).CampaignRecipientEstimationJobCreateQuery(campaignRecipientEstimationJobCreateQuery).Execute()

Create Campaign Recipient Estimation Job



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
	campaignRecipientEstimationJobCreateQuery := *openapiclient.NewCampaignRecipientEstimationJobCreateQuery(*openapiclient.NewCampaignRecipientEstimationJobCreateQueryResourceObject(openapiclient.CampaignRecipientEstimationJobEnum("campaign-recipient-estimation-job"), "Id_example")) // CampaignRecipientEstimationJobCreateQuery | Trigger an asynchronous job to update the estimated number of recipients for the given campaign ID. Use the `Get Campaign Recipient Estimation Job` endpoint to retrieve the status of this estimation job. Use the `Get Campaign Recipient Estimation` endpoint to retrieve the estimated recipient count for a given campaign.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.CreateCampaignRecipientEstimationJob(context.Background()).Revision(revision).CampaignRecipientEstimationJobCreateQuery(campaignRecipientEstimationJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.CreateCampaignRecipientEstimationJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCampaignRecipientEstimationJob`: PostCampaignRecipientEstimationJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.CreateCampaignRecipientEstimationJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRecipientEstimationJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **campaignRecipientEstimationJobCreateQuery** | [**CampaignRecipientEstimationJobCreateQuery**](CampaignRecipientEstimationJobCreateQuery.md) | Trigger an asynchronous job to update the estimated number of recipients for the given campaign ID. Use the &#x60;Get Campaign Recipient Estimation Job&#x60; endpoint to retrieve the status of this estimation job. Use the &#x60;Get Campaign Recipient Estimation&#x60; endpoint to retrieve the estimated recipient count for a given campaign. | 

### Return type

[**PostCampaignRecipientEstimationJobResponse**](PostCampaignRecipientEstimationJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaignSendJob

> PostCampaignSendJobResponse CreateCampaignSendJob(ctx).Revision(revision).CampaignSendJobCreateQuery(campaignSendJobCreateQuery).Execute()

Create Campaign Send Job



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
	campaignSendJobCreateQuery := *openapiclient.NewCampaignSendJobCreateQuery(*openapiclient.NewCampaignSendJobCreateQueryResourceObject(openapiclient.CampaignSendJobEnum("campaign-send-job"), "Id_example")) // CampaignSendJobCreateQuery | Trigger the campaign to send asynchronously

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.CreateCampaignSendJob(context.Background()).Revision(revision).CampaignSendJobCreateQuery(campaignSendJobCreateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.CreateCampaignSendJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCampaignSendJob`: PostCampaignSendJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.CreateCampaignSendJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignSendJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **campaignSendJobCreateQuery** | [**CampaignSendJobCreateQuery**](CampaignSendJobCreateQuery.md) | Trigger the campaign to send asynchronously | 

### Return type

[**PostCampaignSendJobResponse**](PostCampaignSendJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaign

> DeleteCampaign(ctx, id).Revision(revision).Execute()

Delete Campaign



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
	id := "id_example" // string | The campaign ID to be deleted
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.DeleteCampaign(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.DeleteCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The campaign ID to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignRequest struct via the builder pattern


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


## GetCampaign

> GetCampaignResponseCompoundDocument GetCampaign(ctx, id).Revision(revision).FieldsCampaignMessage(fieldsCampaignMessage).FieldsCampaign(fieldsCampaign).FieldsTag(fieldsTag).Include(include).Execute()

Get Campaign



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
	id := "id_example" // string | The campaign ID to be retrieved
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsCampaignMessage := []string{"FieldsCampaignMessage_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCampaign := []string{"FieldsCampaign_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsTag := []string{"FieldsTag_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaign(context.Background(), id).Revision(revision).FieldsCampaignMessage(fieldsCampaignMessage).FieldsCampaign(fieldsCampaign).FieldsTag(fieldsTag).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaign`: GetCampaignResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The campaign ID to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCampaignMessage** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCampaign** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsTag** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetCampaignResponseCompoundDocument**](GetCampaignResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignCampaignMessages

> GetCampaignMessageResponseCollectionCompoundDocument GetCampaignCampaignMessages(ctx, id).Revision(revision).FieldsCampaignMessage(fieldsCampaignMessage).FieldsCampaign(fieldsCampaign).FieldsTemplate(fieldsTemplate).Include(include).Execute()

Get Campaign Campaign Messages



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
	fieldsCampaignMessage := []string{"FieldsCampaignMessage_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCampaign := []string{"FieldsCampaign_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsTemplate := []string{"FieldsTemplate_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaignCampaignMessages(context.Background(), id).Revision(revision).FieldsCampaignMessage(fieldsCampaignMessage).FieldsCampaign(fieldsCampaign).FieldsTemplate(fieldsTemplate).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignCampaignMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignCampaignMessages`: GetCampaignMessageResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignCampaignMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignCampaignMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCampaignMessage** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCampaign** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsTemplate** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetCampaignMessageResponseCollectionCompoundDocument**](GetCampaignMessageResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignMessage

> GetCampaignMessageResponseCompoundDocument GetCampaignMessage(ctx, id).Revision(revision).FieldsCampaignMessage(fieldsCampaignMessage).FieldsCampaign(fieldsCampaign).FieldsTemplate(fieldsTemplate).Include(include).Execute()

Get Campaign Message



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
	id := "id_example" // string | The message ID to be retrieved
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsCampaignMessage := []string{"FieldsCampaignMessage_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCampaign := []string{"FieldsCampaign_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsTemplate := []string{"FieldsTemplate_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaignMessage(context.Background(), id).Revision(revision).FieldsCampaignMessage(fieldsCampaignMessage).FieldsCampaign(fieldsCampaign).FieldsTemplate(fieldsTemplate).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignMessage`: GetCampaignMessageResponseCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The message ID to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCampaignMessage** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCampaign** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsTemplate** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 

### Return type

[**GetCampaignMessageResponseCompoundDocument**](GetCampaignMessageResponseCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignMessageCampaign

> GetCampaignResponse GetCampaignMessageCampaign(ctx, id).Revision(revision).FieldsCampaign(fieldsCampaign).Execute()

Get Campaign Message Campaign



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
	fieldsCampaign := []string{"FieldsCampaign_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaignMessageCampaign(context.Background(), id).Revision(revision).FieldsCampaign(fieldsCampaign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignMessageCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignMessageCampaign`: GetCampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignMessageCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignMessageCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCampaign** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetCampaignResponse**](GetCampaignResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignMessageRelationshipsCampaign

> GetCampaignMessageCampaignRelationshipListResponse GetCampaignMessageRelationshipsCampaign(ctx, id).Revision(revision).Execute()

Get Campaign Message Relationships Campaign



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
	resp, r, err := apiClient.CampaignsAPI.GetCampaignMessageRelationshipsCampaign(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignMessageRelationshipsCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignMessageRelationshipsCampaign`: GetCampaignMessageCampaignRelationshipListResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignMessageRelationshipsCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignMessageRelationshipsCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetCampaignMessageCampaignRelationshipListResponse**](GetCampaignMessageCampaignRelationshipListResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignMessageRelationshipsTemplate

> GetCampaignMessageTemplateRelationshipListResponse GetCampaignMessageRelationshipsTemplate(ctx, id).Revision(revision).Execute()

Get Campaign Message Relationships Template



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
	resp, r, err := apiClient.CampaignsAPI.GetCampaignMessageRelationshipsTemplate(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignMessageRelationshipsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignMessageRelationshipsTemplate`: GetCampaignMessageTemplateRelationshipListResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignMessageRelationshipsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignMessageRelationshipsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetCampaignMessageTemplateRelationshipListResponse**](GetCampaignMessageTemplateRelationshipListResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignMessageTemplate

> GetTemplateResponse GetCampaignMessageTemplate(ctx, id).Revision(revision).FieldsTemplate(fieldsTemplate).Execute()

Get Campaign Message Template



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
	fieldsTemplate := []string{"FieldsTemplate_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaignMessageTemplate(context.Background(), id).Revision(revision).FieldsTemplate(fieldsTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignMessageTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignMessageTemplate`: GetTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignMessageTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignMessageTemplateRequest struct via the builder pattern


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


## GetCampaignRecipientEstimation

> GetCampaignRecipientEstimationResponse GetCampaignRecipientEstimation(ctx, id).Revision(revision).FieldsCampaignRecipientEstimation(fieldsCampaignRecipientEstimation).Execute()

Get Campaign Recipient Estimation



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
	id := "id_example" // string | The ID of the campaign for which to get the estimated number of recipients
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsCampaignRecipientEstimation := []string{"FieldsCampaignRecipientEstimation_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaignRecipientEstimation(context.Background(), id).Revision(revision).FieldsCampaignRecipientEstimation(fieldsCampaignRecipientEstimation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignRecipientEstimation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignRecipientEstimation`: GetCampaignRecipientEstimationResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignRecipientEstimation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the campaign for which to get the estimated number of recipients | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRecipientEstimationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCampaignRecipientEstimation** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetCampaignRecipientEstimationResponse**](GetCampaignRecipientEstimationResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignRecipientEstimationJob

> GetCampaignRecipientEstimationJobResponse GetCampaignRecipientEstimationJob(ctx, id).Revision(revision).FieldsCampaignRecipientEstimationJob(fieldsCampaignRecipientEstimationJob).Execute()

Get Campaign Recipient Estimation Job



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
	id := "id_example" // string | The ID of the campaign to get recipient estimation status
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsCampaignRecipientEstimationJob := []string{"FieldsCampaignRecipientEstimationJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaignRecipientEstimationJob(context.Background(), id).Revision(revision).FieldsCampaignRecipientEstimationJob(fieldsCampaignRecipientEstimationJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignRecipientEstimationJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignRecipientEstimationJob`: GetCampaignRecipientEstimationJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignRecipientEstimationJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the campaign to get recipient estimation status | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRecipientEstimationJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCampaignRecipientEstimationJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetCampaignRecipientEstimationJobResponse**](GetCampaignRecipientEstimationJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignRelationshipsCampaignMessages

> GetCampaignMessagesRelationshipListResponseCollection GetCampaignRelationshipsCampaignMessages(ctx, id).Revision(revision).Execute()

Get Campaign Relationships Campaign Messages



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
	resp, r, err := apiClient.CampaignsAPI.GetCampaignRelationshipsCampaignMessages(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignRelationshipsCampaignMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignRelationshipsCampaignMessages`: GetCampaignMessagesRelationshipListResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignRelationshipsCampaignMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRelationshipsCampaignMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetCampaignMessagesRelationshipListResponseCollection**](GetCampaignMessagesRelationshipListResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignRelationshipsTags

> GetCampaignTagRelationshipListResponseCollection GetCampaignRelationshipsTags(ctx, id).Revision(revision).Execute()

Get Campaign Relationships Tags



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
	resp, r, err := apiClient.CampaignsAPI.GetCampaignRelationshipsTags(context.Background(), id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignRelationshipsTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignRelationshipsTags`: GetCampaignTagRelationshipListResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignRelationshipsTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRelationshipsTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]

### Return type

[**GetCampaignTagRelationshipListResponseCollection**](GetCampaignTagRelationshipListResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignSendJob

> GetCampaignSendJobResponse GetCampaignSendJob(ctx, id).Revision(revision).FieldsCampaignSendJob(fieldsCampaignSendJob).Execute()

Get Campaign Send Job



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
	id := "id_example" // string | The ID of the campaign to send
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsCampaignSendJob := []string{"FieldsCampaignSendJob_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaignSendJob(context.Background(), id).Revision(revision).FieldsCampaignSendJob(fieldsCampaignSendJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignSendJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignSendJob`: GetCampaignSendJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignSendJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the campaign to send | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignSendJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCampaignSendJob** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetCampaignSendJobResponse**](GetCampaignSendJobResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignTags

> GetTagResponseCollection GetCampaignTags(ctx, id).Revision(revision).FieldsTag(fieldsTag).Execute()

Get Campaign Tags



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
	resp, r, err := apiClient.CampaignsAPI.GetCampaignTags(context.Background(), id).Revision(revision).FieldsTag(fieldsTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignTags`: GetTagResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignTagsRequest struct via the builder pattern


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


## GetCampaigns

> GetCampaignResponseCollectionCompoundDocument GetCampaigns(ctx).Filter(filter).Revision(revision).FieldsCampaignMessage(fieldsCampaignMessage).FieldsCampaign(fieldsCampaign).FieldsTag(fieldsTag).Include(include).PageCursor(pageCursor).Sort(sort).Execute()

Get Campaigns



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
	filter := "equals(messages.channel,'sms')" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering<br>Allowed field(s)/operator(s):<br>`messages.channel`: `equals`<br>`name`: `contains`<br>`status`: `any`, `equals`<br>`archived`: `equals`<br>`created_at`: `greater-or-equal`, `greater-than`, `less-or-equal`, `less-than`<br>`scheduled_at`: `greater-or-equal`, `greater-than`, `less-or-equal`, `less-than`<br>`updated_at`: `greater-or-equal`, `greater-than`, `less-or-equal`, `less-than`
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsCampaignMessage := []string{"FieldsCampaignMessage_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsCampaign := []string{"FieldsCampaign_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	fieldsTag := []string{"FieldsTag_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)
	include := []string{"Include_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships (optional)
	pageCursor := "pageCursor_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination (optional)
	sort := "sort_example" // string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaigns(context.Background()).Filter(filter).Revision(revision).FieldsCampaignMessage(fieldsCampaignMessage).FieldsCampaign(fieldsCampaign).FieldsTag(fieldsTag).Include(include).PageCursor(pageCursor).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaigns`: GetCampaignResponseCollectionCompoundDocument
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;messages.channel&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;name&#x60;: &#x60;contains&#x60;&lt;br&gt;&#x60;status&#x60;: &#x60;any&#x60;, &#x60;equals&#x60;&lt;br&gt;&#x60;archived&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;created_at&#x60;: &#x60;greater-or-equal&#x60;, &#x60;greater-than&#x60;, &#x60;less-or-equal&#x60;, &#x60;less-than&#x60;&lt;br&gt;&#x60;scheduled_at&#x60;: &#x60;greater-or-equal&#x60;, &#x60;greater-than&#x60;, &#x60;less-or-equal&#x60;, &#x60;less-than&#x60;&lt;br&gt;&#x60;updated_at&#x60;: &#x60;greater-or-equal&#x60;, &#x60;greater-than&#x60;, &#x60;less-or-equal&#x60;, &#x60;less-than&#x60; | 
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsCampaignMessage** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsCampaign** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **fieldsTag** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 
 **include** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships | 
 **pageCursor** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination | 
 **sort** | **string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting | 

### Return type

[**GetCampaignResponseCollectionCompoundDocument**](GetCampaignResponseCollectionCompoundDocument.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaign

> PatchCampaignResponse UpdateCampaign(ctx, id).Revision(revision).CampaignPartialUpdateQuery(campaignPartialUpdateQuery).Execute()

Update Campaign



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
	id := "id_example" // string | The campaign ID to be retrieved
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	campaignPartialUpdateQuery := *openapiclient.NewCampaignPartialUpdateQuery(*openapiclient.NewCampaignPartialUpdateQueryResourceObject(openapiclient.CampaignEnum("campaign"), "Id_example", *openapiclient.NewCampaignPartialUpdateQueryResourceObjectAttributes())) // CampaignPartialUpdateQuery | Update a campaign and return it

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.UpdateCampaign(context.Background(), id).Revision(revision).CampaignPartialUpdateQuery(campaignPartialUpdateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.UpdateCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCampaign`: PatchCampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.UpdateCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The campaign ID to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **campaignPartialUpdateQuery** | [**CampaignPartialUpdateQuery**](CampaignPartialUpdateQuery.md) | Update a campaign and return it | 

### Return type

[**PatchCampaignResponse**](PatchCampaignResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaignMessage

> PatchCampaignMessageResponse UpdateCampaignMessage(ctx, id).Revision(revision).CampaignMessagePartialUpdateQuery(campaignMessagePartialUpdateQuery).Execute()

Update Campaign Message



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
	id := "id_example" // string | The message ID to be retrieved
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	campaignMessagePartialUpdateQuery := *openapiclient.NewCampaignMessagePartialUpdateQuery(*openapiclient.NewCampaignMessagePartialUpdateQueryResourceObject(openapiclient.CampaignMessageEnum("campaign-message"), "Id_example", *openapiclient.NewCampaignMessagePartialUpdateQueryResourceObjectAttributes())) // CampaignMessagePartialUpdateQuery | Update a message and return it

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.UpdateCampaignMessage(context.Background(), id).Revision(revision).CampaignMessagePartialUpdateQuery(campaignMessagePartialUpdateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.UpdateCampaignMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCampaignMessage`: PatchCampaignMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.UpdateCampaignMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The message ID to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **campaignMessagePartialUpdateQuery** | [**CampaignMessagePartialUpdateQuery**](CampaignMessagePartialUpdateQuery.md) | Update a message and return it | 

### Return type

[**PatchCampaignMessageResponse**](PatchCampaignMessageResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaignSendJob

> UpdateCampaignSendJob(ctx, id).Revision(revision).CampaignSendJobPartialUpdateQuery(campaignSendJobPartialUpdateQuery).Execute()

Update Campaign Send Job



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
	id := "id_example" // string | The ID of the currently sending campaign to cancel or revert
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	campaignSendJobPartialUpdateQuery := *openapiclient.NewCampaignSendJobPartialUpdateQuery(*openapiclient.NewCampaignSendJobPartialUpdateQueryResourceObject(openapiclient.CampaignSendJobEnum("campaign-send-job"), "Id_example", *openapiclient.NewCampaignSendJobPartialUpdateQueryResourceObjectAttributes("cancel"))) // CampaignSendJobPartialUpdateQuery | Permanently cancel the campaign, setting the status to CANCELED or revert the campaign, setting the status back to DRAFT

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.UpdateCampaignSendJob(context.Background(), id).Revision(revision).CampaignSendJobPartialUpdateQuery(campaignSendJobPartialUpdateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.UpdateCampaignSendJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the currently sending campaign to cancel or revert | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignSendJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **campaignSendJobPartialUpdateQuery** | [**CampaignSendJobPartialUpdateQuery**](CampaignSendJobPartialUpdateQuery.md) | Permanently cancel the campaign, setting the status to CANCELED or revert the campaign, setting the status back to DRAFT | 

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

