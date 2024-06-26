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
	"strings"
)


// MetricsAPIService MetricsAPI service
type MetricsAPIService service

type ApiGetMetricRequest struct {
	ctx context.Context
	ApiService *MetricsAPIService
	id string
	revision *string
	fieldsMetric *[]string
}

// API endpoint revision (format: YYYY-MM-DD[.suffix])
func (r ApiGetMetricRequest) Revision(revision string) ApiGetMetricRequest {
	r.revision = &revision
	return r
}

// For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets
func (r ApiGetMetricRequest) FieldsMetric(fieldsMetric []string) ApiGetMetricRequest {
	r.fieldsMetric = &fieldsMetric
	return r
}

func (r ApiGetMetricRequest) Execute() (*GetMetricResponse, *http.Response, error) {
	return r.ApiService.GetMetricExecute(r)
}

/*
GetMetric Get Metric

Get a metric with the given metric ID.<br><br>*Rate limits*:<br>Burst: `10/s`<br>Steady: `150/m`

**Scopes:**
`metrics:read`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Metric ID
 @return ApiGetMetricRequest
*/
func (a *MetricsAPIService) GetMetric(ctx context.Context, id string) ApiGetMetricRequest {
	return ApiGetMetricRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetMetricResponse
func (a *MetricsAPIService) GetMetricExecute(r ApiGetMetricRequest) (*GetMetricResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMetricResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsAPIService.GetMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/metrics/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.revision == nil {
		return localVarReturnValue, nil, reportError("revision is required and must be specified")
	}

	if r.fieldsMetric != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[metric]", r.fieldsMetric, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGetMetricsRequest struct {
	ctx context.Context
	ApiService *MetricsAPIService
	revision *string
	fieldsMetric *[]string
	filter *string
	pageCursor *string
}

// API endpoint revision (format: YYYY-MM-DD[.suffix])
func (r ApiGetMetricsRequest) Revision(revision string) ApiGetMetricsRequest {
	r.revision = &revision
	return r
}

// For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets
func (r ApiGetMetricsRequest) FieldsMetric(fieldsMetric []string) ApiGetMetricsRequest {
	r.fieldsMetric = &fieldsMetric
	return r
}

// For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering&lt;br&gt;Allowed field(s)/operator(s):&lt;br&gt;&#x60;integration.name&#x60;: &#x60;equals&#x60;&lt;br&gt;&#x60;integration.category&#x60;: &#x60;equals&#x60;
func (r ApiGetMetricsRequest) Filter(filter string) ApiGetMetricsRequest {
	r.filter = &filter
	return r
}

// For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination
func (r ApiGetMetricsRequest) PageCursor(pageCursor string) ApiGetMetricsRequest {
	r.pageCursor = &pageCursor
	return r
}

func (r ApiGetMetricsRequest) Execute() (*GetMetricResponseCollection, *http.Response, error) {
	return r.ApiService.GetMetricsExecute(r)
}

/*
GetMetrics Get Metrics

Get all metrics in an account.

Requests can be filtered by the following fields:
integration `name`, integration `category`

Returns a maximum of 200 results per page.<br><br>*Rate limits*:<br>Burst: `10/s`<br>Steady: `150/m`

**Scopes:**
`metrics:read`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMetricsRequest
*/
func (a *MetricsAPIService) GetMetrics(ctx context.Context) ApiGetMetricsRequest {
	return ApiGetMetricsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMetricResponseCollection
func (a *MetricsAPIService) GetMetricsExecute(r ApiGetMetricsRequest) (*GetMetricResponseCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMetricResponseCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsAPIService.GetMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/metrics/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.revision == nil {
		return localVarReturnValue, nil, reportError("revision is required and must be specified")
	}

	if r.fieldsMetric != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[metric]", r.fieldsMetric, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	if r.pageCursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[cursor]", r.pageCursor, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiQueryMetricAggregatesRequest struct {
	ctx context.Context
	ApiService *MetricsAPIService
	revision *string
	metricAggregateQuery *MetricAggregateQuery
}

// API endpoint revision (format: YYYY-MM-DD[.suffix])
func (r ApiQueryMetricAggregatesRequest) Revision(revision string) ApiQueryMetricAggregatesRequest {
	r.revision = &revision
	return r
}

// Retrieve Metric Aggregations
func (r ApiQueryMetricAggregatesRequest) MetricAggregateQuery(metricAggregateQuery MetricAggregateQuery) ApiQueryMetricAggregatesRequest {
	r.metricAggregateQuery = &metricAggregateQuery
	return r
}

func (r ApiQueryMetricAggregatesRequest) Execute() (*PostMetricAggregateResponse, *http.Response, error) {
	return r.ApiService.QueryMetricAggregatesExecute(r)
}

/*
QueryMetricAggregates Query Metric Aggregates

Query and aggregate event data associated with a metric, including native Klaviyo metrics, integration-specific metrics, and custom events. Queries must be passed in the JSON body of your `POST` request.

Results can be filtered and grouped by time, event, or profile dimensions.

To learn more about how to use this endpoint, check out our new [Using the Query Metric Aggregates Endpoint guide](https://developers.klaviyo.com/en/docs/using-the-query-metric-aggregates-endpoint).

**Request body parameters** (nested under `attributes`):

* `return_fields`: request specific fields using [sparse fieldsets](https://developers.klaviyo.com/en/reference/api_overview#sparse-fieldsets)
* `sort`: sort results by a specified field, such as `"-timestamp"`
* `page_cursor`: results can be paginated with [cursor-based pagination](https://developers.klaviyo.com/en/reference/api_overview#pagination)
* `page_size`: limit the number of returned results per page
* `by`: optional attributes used to group by the aggregation function
    * When using `by` attributes, an empty `dimensions` response is expected when the counts for the events do not have the associated dimension requested by the set `by` attribute. For example, a query including `"by": ["$flow"]` will return an empty dimensions response for counts of metrics not associated with a `$flow`
* `measurement`: the measurement key supports the following values:
    * `"sum_value"`: perform a summation of the `_Event Value_`, optionally partitioned over any dimension provided in the `by` field
    * `"count"`: counts the number of events associated to a metric, optionally partitioned over any dimension provided in the `by` field
    * `"unique"` counts the number of unique customers associated to a metric, optionally partitioned over any dimension provided in the `by` field
* `interval`: aggregation interval, such as `"hour"`,`"day"`,`"week"`, and `"month"`
* `metric_id`: the metric ID used in the aggregation
* `filter`: list of filters for specific fields, must include time range using ISO 8601 format (`"YYYY-MM-DDTHH:MM:SS.mmmmmm"`)
    * The time range can be filtered by providing a `greater-or-equal` filter on the datetime field, such as `"greater-or-equal(datetime,2021-07-01T00:00:00)"` and a `less-than` filter on the same datetime field, such as `"less-than(datetime,2022-07-01T00:00:00)"`
    * The time range may span a maximum of one year. Time range dates may be set to a maximum of 5 years prior to the current date
    * Filter the list of supported aggregate dimensions using the common filter syntax, such as `"equals(URL,\"https://www.klaviyo.com/\")"`
* `timezone`: the timezone used when processing the query. Case sensitive. This field is validated against a list of common timezones from the [IANA Time Zone Database](https://www.iana.org/time-zones)
    * While the payload accepts a timezone, the response datetimes returned will be in UTC.

For a comprehensive list of native Klaviyo metrics and their associated attributes for grouping and filtering, please refer to the [metrics attributes guide](https://developers.klaviyo.com/en/docs/supported_metrics_and_attributes).<br><br>*Rate limits*:<br>Burst: `3/s`<br>Steady: `60/m`

**Scopes:**
`metrics:read`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryMetricAggregatesRequest
*/
func (a *MetricsAPIService) QueryMetricAggregates(ctx context.Context) ApiQueryMetricAggregatesRequest {
	return ApiQueryMetricAggregatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostMetricAggregateResponse
func (a *MetricsAPIService) QueryMetricAggregatesExecute(r ApiQueryMetricAggregatesRequest) (*PostMetricAggregateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostMetricAggregateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsAPIService.QueryMetricAggregates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/metric-aggregates/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.revision == nil {
		return localVarReturnValue, nil, reportError("revision is required and must be specified")
	}
	if r.metricAggregateQuery == nil {
		return localVarReturnValue, nil, reportError("metricAggregateQuery is required and must be specified")
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
	localVarPostBody = r.metricAggregateQuery
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
