/*
Klaviyo API

The Klaviyo REST API. Please visit https://developers.klaviyo.com for more details.

API version: 2024-02-15
Contact: developers@klaviyo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package klaviyo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ReportingAPIService ReportingAPI service
type ReportingAPIService service

type ApiQueryCampaignValuesRequest struct {
	ctx context.Context
	ApiService *ReportingAPIService
	revision *string
	campaignValuesRequestDTO *CampaignValuesRequestDTO
	pageCursor *string
}

// API endpoint revision (format: YYYY-MM-DD[.suffix])
func (r ApiQueryCampaignValuesRequest) Revision(revision string) ApiQueryCampaignValuesRequest {
	r.revision = &revision
	return r
}

func (r ApiQueryCampaignValuesRequest) CampaignValuesRequestDTO(campaignValuesRequestDTO CampaignValuesRequestDTO) ApiQueryCampaignValuesRequest {
	r.campaignValuesRequestDTO = &campaignValuesRequestDTO
	return r
}

// For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination
func (r ApiQueryCampaignValuesRequest) PageCursor(pageCursor string) ApiQueryCampaignValuesRequest {
	r.pageCursor = &pageCursor
	return r
}

func (r ApiQueryCampaignValuesRequest) Execute() (*PostCampaignValuesResponseDTO, *http.Response, error) {
	return r.ApiService.QueryCampaignValuesExecute(r)
}

/*
QueryCampaignValues Query Campaign Values

Returns the requested campaign analytics values data<br><br>*Rate limits*:<br>Burst: `1/s`<br>Steady: `2/m`<br>Daily: `225/d`

**Scopes:**
`campaigns:read`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryCampaignValuesRequest
*/
func (a *ReportingAPIService) QueryCampaignValues(ctx context.Context) ApiQueryCampaignValuesRequest {
	return ApiQueryCampaignValuesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostCampaignValuesResponseDTO
func (a *ReportingAPIService) QueryCampaignValuesExecute(r ApiQueryCampaignValuesRequest) (*PostCampaignValuesResponseDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostCampaignValuesResponseDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportingAPIService.QueryCampaignValues")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/campaign-values-reports/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.revision == nil {
		return localVarReturnValue, nil, reportError("revision is required and must be specified")
	}
	if r.campaignValuesRequestDTO == nil {
		return localVarReturnValue, nil, reportError("campaignValuesRequestDTO is required and must be specified")
	}

	if r.pageCursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_cursor", r.pageCursor, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "revision", r.revision, "")
	// body params
	localVarPostBody = r.campaignValuesRequestDTO
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Klaviyo-API-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v GetAccounts4XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v GetAccounts4XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQueryFlowSeriesRequest struct {
	ctx context.Context
	ApiService *ReportingAPIService
	revision *string
	flowSeriesRequestDTO *FlowSeriesRequestDTO
	pageCursor *string
}

// API endpoint revision (format: YYYY-MM-DD[.suffix])
func (r ApiQueryFlowSeriesRequest) Revision(revision string) ApiQueryFlowSeriesRequest {
	r.revision = &revision
	return r
}

func (r ApiQueryFlowSeriesRequest) FlowSeriesRequestDTO(flowSeriesRequestDTO FlowSeriesRequestDTO) ApiQueryFlowSeriesRequest {
	r.flowSeriesRequestDTO = &flowSeriesRequestDTO
	return r
}

// For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination
func (r ApiQueryFlowSeriesRequest) PageCursor(pageCursor string) ApiQueryFlowSeriesRequest {
	r.pageCursor = &pageCursor
	return r
}

func (r ApiQueryFlowSeriesRequest) Execute() (*PostFlowSeriesResponseDTO, *http.Response, error) {
	return r.ApiService.QueryFlowSeriesExecute(r)
}

/*
QueryFlowSeries Query Flow Series

Returns the requested flow analytics series data<br><br>*Rate limits*:<br>Burst: `1/s`<br>Steady: `2/m`<br>Daily: `225/d`

**Scopes:**
`flows:read`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryFlowSeriesRequest
*/
func (a *ReportingAPIService) QueryFlowSeries(ctx context.Context) ApiQueryFlowSeriesRequest {
	return ApiQueryFlowSeriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostFlowSeriesResponseDTO
func (a *ReportingAPIService) QueryFlowSeriesExecute(r ApiQueryFlowSeriesRequest) (*PostFlowSeriesResponseDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostFlowSeriesResponseDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportingAPIService.QueryFlowSeries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/flow-series-reports/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.revision == nil {
		return localVarReturnValue, nil, reportError("revision is required and must be specified")
	}
	if r.flowSeriesRequestDTO == nil {
		return localVarReturnValue, nil, reportError("flowSeriesRequestDTO is required and must be specified")
	}

	if r.pageCursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_cursor", r.pageCursor, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "revision", r.revision, "")
	// body params
	localVarPostBody = r.flowSeriesRequestDTO
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Klaviyo-API-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v GetAccounts4XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v GetAccounts4XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQueryFlowValuesRequest struct {
	ctx context.Context
	ApiService *ReportingAPIService
	revision *string
	flowValuesRequestDTO *FlowValuesRequestDTO
	pageCursor *string
}

// API endpoint revision (format: YYYY-MM-DD[.suffix])
func (r ApiQueryFlowValuesRequest) Revision(revision string) ApiQueryFlowValuesRequest {
	r.revision = &revision
	return r
}

func (r ApiQueryFlowValuesRequest) FlowValuesRequestDTO(flowValuesRequestDTO FlowValuesRequestDTO) ApiQueryFlowValuesRequest {
	r.flowValuesRequestDTO = &flowValuesRequestDTO
	return r
}

// For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination
func (r ApiQueryFlowValuesRequest) PageCursor(pageCursor string) ApiQueryFlowValuesRequest {
	r.pageCursor = &pageCursor
	return r
}

func (r ApiQueryFlowValuesRequest) Execute() (*PostFlowValuesResponseDTO, *http.Response, error) {
	return r.ApiService.QueryFlowValuesExecute(r)
}

/*
QueryFlowValues Query Flow Values

Returns the requested flow analytics values data<br><br>*Rate limits*:<br>Burst: `1/s`<br>Steady: `2/m`<br>Daily: `225/d`

**Scopes:**
`flows:read`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryFlowValuesRequest
*/
func (a *ReportingAPIService) QueryFlowValues(ctx context.Context) ApiQueryFlowValuesRequest {
	return ApiQueryFlowValuesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostFlowValuesResponseDTO
func (a *ReportingAPIService) QueryFlowValuesExecute(r ApiQueryFlowValuesRequest) (*PostFlowValuesResponseDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostFlowValuesResponseDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportingAPIService.QueryFlowValues")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/flow-values-reports/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.revision == nil {
		return localVarReturnValue, nil, reportError("revision is required and must be specified")
	}
	if r.flowValuesRequestDTO == nil {
		return localVarReturnValue, nil, reportError("flowValuesRequestDTO is required and must be specified")
	}

	if r.pageCursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_cursor", r.pageCursor, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "revision", r.revision, "")
	// body params
	localVarPostBody = r.flowValuesRequestDTO
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Klaviyo-API-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v GetAccounts4XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v GetAccounts4XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}