# FlowSeriesRequestDTOResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | **[]string** | List of statistics to query for. All rate statistics will be returned in fractional form [0.0, 1.0] | 
**Timeframe** | [**CampaignValuesRequestDTOResourceObjectAttributesTimeframe**](CampaignValuesRequestDTOResourceObjectAttributesTimeframe.md) |  | 
**Interval** | **string** | The interval used to aggregate data within the series request. If hourly is used, the timeframe cannot be longer than 7 days. If daily is used, the timeframe cannot be longer than 60 days. If monthly is used, the timeframe cannot be longer than 52 weeks. | 
**ConversionMetricId** | **string** | ID of the metric to be used for conversion statistics | 
**Filter** | Pointer to **NullableString** | API filter string used to filter the query. Allowed filters are flow_id, send_channel, flow_message_id. Allowed operators are equals, contains-any. Only one filter can be used per attribute, only AND can be used as a combination operator. Max of 100 messages per ANY filter. | [optional] 

## Methods

### NewFlowSeriesRequestDTOResourceObjectAttributes

`func NewFlowSeriesRequestDTOResourceObjectAttributes(statistics []string, timeframe CampaignValuesRequestDTOResourceObjectAttributesTimeframe, interval string, conversionMetricId string, ) *FlowSeriesRequestDTOResourceObjectAttributes`

NewFlowSeriesRequestDTOResourceObjectAttributes instantiates a new FlowSeriesRequestDTOResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowSeriesRequestDTOResourceObjectAttributesWithDefaults

`func NewFlowSeriesRequestDTOResourceObjectAttributesWithDefaults() *FlowSeriesRequestDTOResourceObjectAttributes`

NewFlowSeriesRequestDTOResourceObjectAttributesWithDefaults instantiates a new FlowSeriesRequestDTOResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetStatistics() []string`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetStatisticsOk() (*[]string, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) SetStatistics(v []string)`

SetStatistics sets Statistics field to given value.


### GetTimeframe

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetTimeframe() CampaignValuesRequestDTOResourceObjectAttributesTimeframe`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetTimeframeOk() (*CampaignValuesRequestDTOResourceObjectAttributesTimeframe, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) SetTimeframe(v CampaignValuesRequestDTOResourceObjectAttributesTimeframe)`

SetTimeframe sets Timeframe field to given value.


### GetInterval

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetConversionMetricId

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetConversionMetricId() string`

GetConversionMetricId returns the ConversionMetricId field if non-nil, zero value otherwise.

### GetConversionMetricIdOk

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetConversionMetricIdOk() (*string, bool)`

GetConversionMetricIdOk returns a tuple with the ConversionMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionMetricId

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) SetConversionMetricId(v string)`

SetConversionMetricId sets ConversionMetricId field to given value.


### GetFilter

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *FlowSeriesRequestDTOResourceObjectAttributes) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *FlowSeriesRequestDTOResourceObjectAttributes) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


