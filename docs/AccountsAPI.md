# \AccountsAPI

All URIs are relative to *https://a.klaviyo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](AccountsAPI.md#GetAccount) | **Get** /api/accounts/{id}/ | Get Account
[**GetAccounts**](AccountsAPI.md#GetAccounts) | **Get** /api/accounts/ | Get Accounts



## GetAccount

> GetAccountResponse GetAccount(ctx, id).Revision(revision).FieldsAccount(fieldsAccount).Execute()

Get Account



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
	id := "AbC123" // string | The ID of the account
	revision := "revision_example" // string | API endpoint revision (format: YYYY-MM-DD[.suffix]) (default to "2024-02-15")
	fieldsAccount := []string{"FieldsAccount_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccount(context.Background(), id).Revision(revision).FieldsAccount(fieldsAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: GetAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsAccount** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetAccountResponse**](GetAccountResponse.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccounts

> GetAccountResponseCollection GetAccounts(ctx).Revision(revision).FieldsAccount(fieldsAccount).Execute()

Get Accounts



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
	fieldsAccount := []string{"FieldsAccount_example"} // []string | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccounts(context.Background()).Revision(revision).FieldsAccount(fieldsAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccounts`: GetAccountResponseCollection
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **string** | API endpoint revision (format: YYYY-MM-DD[.suffix]) | [default to &quot;2024-02-15&quot;]
 **fieldsAccount** | **[]string** | For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets | 

### Return type

[**GetAccountResponseCollection**](GetAccountResponseCollection.md)

### Authorization

[Klaviyo-API-Key](../README.md#Klaviyo-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

