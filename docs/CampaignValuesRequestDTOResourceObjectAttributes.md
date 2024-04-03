# CampaignValuesRequestDTOResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | **[]string** | List of statistics to query for. All rate statistics will be returned in fractional form [0.0, 1.0] | 
**Timeframe** | [**CampaignValuesRequestDTOResourceObjectAttributesTimeframe**](CampaignValuesRequestDTOResourceObjectAttributesTimeframe.md) |  | 
**ConversionMetricId** | **string** | ID of the metric to be used for conversion statistics | 
**Filter** | Pointer to **NullableString** | API filter string used to filter the query. Allowed filters are send_channel, campaign_id. Allowed operators are equals, contains-any. Only one filter can be used per attribute, only AND can be used as a combination operator. Max of 100 messages per ANY filter. | [optional] 

## Methods

### NewCampaignValuesRequestDTOResourceObjectAttributes

`func NewCampaignValuesRequestDTOResourceObjectAttributes(statistics []string, timeframe CampaignValuesRequestDTOResourceObjectAttributesTimeframe, conversionMetricId string, ) *CampaignValuesRequestDTOResourceObjectAttributes`

NewCampaignValuesRequestDTOResourceObjectAttributes instantiates a new CampaignValuesRequestDTOResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignValuesRequestDTOResourceObjectAttributesWithDefaults

`func NewCampaignValuesRequestDTOResourceObjectAttributesWithDefaults() *CampaignValuesRequestDTOResourceObjectAttributes`

NewCampaignValuesRequestDTOResourceObjectAttributesWithDefaults instantiates a new CampaignValuesRequestDTOResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) GetStatistics() []string`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) GetStatisticsOk() (*[]string, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) SetStatistics(v []string)`

SetStatistics sets Statistics field to given value.


### GetTimeframe

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) GetTimeframe() CampaignValuesRequestDTOResourceObjectAttributesTimeframe`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) GetTimeframeOk() (*CampaignValuesRequestDTOResourceObjectAttributesTimeframe, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) SetTimeframe(v CampaignValuesRequestDTOResourceObjectAttributesTimeframe)`

SetTimeframe sets Timeframe field to given value.


### GetConversionMetricId

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) GetConversionMetricId() string`

GetConversionMetricId returns the ConversionMetricId field if non-nil, zero value otherwise.

### GetConversionMetricIdOk

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) GetConversionMetricIdOk() (*string, bool)`

GetConversionMetricIdOk returns a tuple with the ConversionMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionMetricId

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) SetConversionMetricId(v string)`

SetConversionMetricId sets ConversionMetricId field to given value.


### GetFilter

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *CampaignValuesRequestDTOResourceObjectAttributes) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *CampaignValuesRequestDTOResourceObjectAttributes) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


