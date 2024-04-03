# FlowValuesRequestDTOResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | **[]string** | List of statistics to query for. All rate statistics will be returned in fractional form [0.0, 1.0] | 
**Timeframe** | [**CampaignValuesRequestDTOResourceObjectAttributesTimeframe**](CampaignValuesRequestDTOResourceObjectAttributesTimeframe.md) |  | 
**ConversionMetricId** | **string** | ID of the metric to be used for conversion statistics | 
**Filter** | Pointer to **NullableString** | API filter string used to filter the query. Allowed filters are flow_id, send_channel, flow_message_id. Allowed operators are equals, contains-any. Only one filter can be used per attribute, only AND can be used as a combination operator. Max of 100 messages per ANY filter. | [optional] 

## Methods

### NewFlowValuesRequestDTOResourceObjectAttributes

`func NewFlowValuesRequestDTOResourceObjectAttributes(statistics []string, timeframe CampaignValuesRequestDTOResourceObjectAttributesTimeframe, conversionMetricId string, ) *FlowValuesRequestDTOResourceObjectAttributes`

NewFlowValuesRequestDTOResourceObjectAttributes instantiates a new FlowValuesRequestDTOResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowValuesRequestDTOResourceObjectAttributesWithDefaults

`func NewFlowValuesRequestDTOResourceObjectAttributesWithDefaults() *FlowValuesRequestDTOResourceObjectAttributes`

NewFlowValuesRequestDTOResourceObjectAttributesWithDefaults instantiates a new FlowValuesRequestDTOResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *FlowValuesRequestDTOResourceObjectAttributes) GetStatistics() []string`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *FlowValuesRequestDTOResourceObjectAttributes) GetStatisticsOk() (*[]string, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *FlowValuesRequestDTOResourceObjectAttributes) SetStatistics(v []string)`

SetStatistics sets Statistics field to given value.


### GetTimeframe

`func (o *FlowValuesRequestDTOResourceObjectAttributes) GetTimeframe() CampaignValuesRequestDTOResourceObjectAttributesTimeframe`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *FlowValuesRequestDTOResourceObjectAttributes) GetTimeframeOk() (*CampaignValuesRequestDTOResourceObjectAttributesTimeframe, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *FlowValuesRequestDTOResourceObjectAttributes) SetTimeframe(v CampaignValuesRequestDTOResourceObjectAttributesTimeframe)`

SetTimeframe sets Timeframe field to given value.


### GetConversionMetricId

`func (o *FlowValuesRequestDTOResourceObjectAttributes) GetConversionMetricId() string`

GetConversionMetricId returns the ConversionMetricId field if non-nil, zero value otherwise.

### GetConversionMetricIdOk

`func (o *FlowValuesRequestDTOResourceObjectAttributes) GetConversionMetricIdOk() (*string, bool)`

GetConversionMetricIdOk returns a tuple with the ConversionMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionMetricId

`func (o *FlowValuesRequestDTOResourceObjectAttributes) SetConversionMetricId(v string)`

SetConversionMetricId sets ConversionMetricId field to given value.


### GetFilter

`func (o *FlowValuesRequestDTOResourceObjectAttributes) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *FlowValuesRequestDTOResourceObjectAttributes) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *FlowValuesRequestDTOResourceObjectAttributes) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *FlowValuesRequestDTOResourceObjectAttributes) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *FlowValuesRequestDTOResourceObjectAttributes) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *FlowValuesRequestDTOResourceObjectAttributes) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


