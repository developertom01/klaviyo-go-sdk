# MetricAggregateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricId** | **string** | The metric ID used in the aggregation. | 
**PageCursor** | Pointer to **string** | Optional pagination cursor to iterate over large result sets | [optional] 
**Measurements** | **[]string** | Measurement key, e.g. &#x60;unique&#x60;, &#x60;sum_value&#x60;, &#x60;count&#x60; | 
**Interval** | Pointer to **NullableString** | Aggregation interval, e.g. \&quot;hour\&quot;, \&quot;day\&quot;, \&quot;week\&quot;, \&quot;month\&quot; | [optional] [default to "day"]
**PageSize** | Pointer to **NullableInt32** | Alter the maximum number of returned rows in a single page of aggregation results | [optional] [default to 500]
**By** | Pointer to **[]string** | Optional attribute(s) used for partitioning by the aggregation function | [optional] 
**ReturnFields** | Pointer to **[]string** | Provide fields to limit the returned data | [optional] 
**Filter** | **[]string** | List of filters, must include time range using ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).             These filters follow a similar format to those in &#x60;GET&#x60; requests, the primary difference is that this endpoint asks for a list.             The time range can be filtered by providing a &#x60;greater-or-equal&#x60; and a &#x60;less-than&#x60; filter on the &#x60;datetime&#x60; field. | 
**Timezone** | Pointer to **NullableString** | The timezone used for processing the query, e.g. &#x60;&#39;America/New_York&#39;&#x60;.             This field is validated against a list of common timezones from the [IANA Time Zone Database](https://www.iana.org/time-zones).             While most are supported, a few notable exceptions are &#x60;Factory&#x60;, &#x60;Europe/Kyiv&#x60; and &#x60;Pacific/Kanton&#x60;. This field is case-sensitive. | [optional] [default to "UTC"]
**Sort** | Pointer to **string** | Provide a sort key (e.g. -$message) | [optional] 

## Methods

### NewMetricAggregateQueryResourceObjectAttributes

`func NewMetricAggregateQueryResourceObjectAttributes(metricId string, measurements []string, filter []string, ) *MetricAggregateQueryResourceObjectAttributes`

NewMetricAggregateQueryResourceObjectAttributes instantiates a new MetricAggregateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricAggregateQueryResourceObjectAttributesWithDefaults

`func NewMetricAggregateQueryResourceObjectAttributesWithDefaults() *MetricAggregateQueryResourceObjectAttributes`

NewMetricAggregateQueryResourceObjectAttributesWithDefaults instantiates a new MetricAggregateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricId

`func (o *MetricAggregateQueryResourceObjectAttributes) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MetricAggregateQueryResourceObjectAttributes) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MetricAggregateQueryResourceObjectAttributes) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.


### GetPageCursor

`func (o *MetricAggregateQueryResourceObjectAttributes) GetPageCursor() string`

GetPageCursor returns the PageCursor field if non-nil, zero value otherwise.

### GetPageCursorOk

`func (o *MetricAggregateQueryResourceObjectAttributes) GetPageCursorOk() (*string, bool)`

GetPageCursorOk returns a tuple with the PageCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCursor

`func (o *MetricAggregateQueryResourceObjectAttributes) SetPageCursor(v string)`

SetPageCursor sets PageCursor field to given value.

### HasPageCursor

`func (o *MetricAggregateQueryResourceObjectAttributes) HasPageCursor() bool`

HasPageCursor returns a boolean if a field has been set.

### GetMeasurements

`func (o *MetricAggregateQueryResourceObjectAttributes) GetMeasurements() []string`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *MetricAggregateQueryResourceObjectAttributes) GetMeasurementsOk() (*[]string, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *MetricAggregateQueryResourceObjectAttributes) SetMeasurements(v []string)`

SetMeasurements sets Measurements field to given value.


### GetInterval

`func (o *MetricAggregateQueryResourceObjectAttributes) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MetricAggregateQueryResourceObjectAttributes) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MetricAggregateQueryResourceObjectAttributes) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *MetricAggregateQueryResourceObjectAttributes) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *MetricAggregateQueryResourceObjectAttributes) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *MetricAggregateQueryResourceObjectAttributes) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetPageSize

`func (o *MetricAggregateQueryResourceObjectAttributes) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *MetricAggregateQueryResourceObjectAttributes) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *MetricAggregateQueryResourceObjectAttributes) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *MetricAggregateQueryResourceObjectAttributes) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *MetricAggregateQueryResourceObjectAttributes) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *MetricAggregateQueryResourceObjectAttributes) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetBy

`func (o *MetricAggregateQueryResourceObjectAttributes) GetBy() []string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *MetricAggregateQueryResourceObjectAttributes) GetByOk() (*[]string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *MetricAggregateQueryResourceObjectAttributes) SetBy(v []string)`

SetBy sets By field to given value.

### HasBy

`func (o *MetricAggregateQueryResourceObjectAttributes) HasBy() bool`

HasBy returns a boolean if a field has been set.

### SetByNil

`func (o *MetricAggregateQueryResourceObjectAttributes) SetByNil(b bool)`

 SetByNil sets the value for By to be an explicit nil

### UnsetBy
`func (o *MetricAggregateQueryResourceObjectAttributes) UnsetBy()`

UnsetBy ensures that no value is present for By, not even an explicit nil
### GetReturnFields

`func (o *MetricAggregateQueryResourceObjectAttributes) GetReturnFields() []string`

GetReturnFields returns the ReturnFields field if non-nil, zero value otherwise.

### GetReturnFieldsOk

`func (o *MetricAggregateQueryResourceObjectAttributes) GetReturnFieldsOk() (*[]string, bool)`

GetReturnFieldsOk returns a tuple with the ReturnFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnFields

`func (o *MetricAggregateQueryResourceObjectAttributes) SetReturnFields(v []string)`

SetReturnFields sets ReturnFields field to given value.

### HasReturnFields

`func (o *MetricAggregateQueryResourceObjectAttributes) HasReturnFields() bool`

HasReturnFields returns a boolean if a field has been set.

### SetReturnFieldsNil

`func (o *MetricAggregateQueryResourceObjectAttributes) SetReturnFieldsNil(b bool)`

 SetReturnFieldsNil sets the value for ReturnFields to be an explicit nil

### UnsetReturnFields
`func (o *MetricAggregateQueryResourceObjectAttributes) UnsetReturnFields()`

UnsetReturnFields ensures that no value is present for ReturnFields, not even an explicit nil
### GetFilter

`func (o *MetricAggregateQueryResourceObjectAttributes) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MetricAggregateQueryResourceObjectAttributes) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MetricAggregateQueryResourceObjectAttributes) SetFilter(v []string)`

SetFilter sets Filter field to given value.


### GetTimezone

`func (o *MetricAggregateQueryResourceObjectAttributes) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MetricAggregateQueryResourceObjectAttributes) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MetricAggregateQueryResourceObjectAttributes) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *MetricAggregateQueryResourceObjectAttributes) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *MetricAggregateQueryResourceObjectAttributes) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *MetricAggregateQueryResourceObjectAttributes) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetSort

`func (o *MetricAggregateQueryResourceObjectAttributes) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *MetricAggregateQueryResourceObjectAttributes) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *MetricAggregateQueryResourceObjectAttributes) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *MetricAggregateQueryResourceObjectAttributes) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


