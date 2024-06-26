/*
Klaviyo API

The Klaviyo REST API. Please visit https://developers.klaviyo.com for more details.

API version: 2024-02-15
Contact: developers@klaviyo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package klaviyo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FlowSeriesRequestDTOResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowSeriesRequestDTOResourceObjectAttributes{}

// FlowSeriesRequestDTOResourceObjectAttributes struct for FlowSeriesRequestDTOResourceObjectAttributes
type FlowSeriesRequestDTOResourceObjectAttributes struct {
	// List of statistics to query for. All rate statistics will be returned in fractional form [0.0, 1.0]
	Statistics []string `json:"statistics"`
	Timeframe CampaignValuesRequestDTOResourceObjectAttributesTimeframe `json:"timeframe"`
	// The interval used to aggregate data within the series request. If hourly is used, the timeframe cannot be longer than 7 days. If daily is used, the timeframe cannot be longer than 60 days. If monthly is used, the timeframe cannot be longer than 52 weeks.
	Interval string `json:"interval"`
	// ID of the metric to be used for conversion statistics
	ConversionMetricId string `json:"conversion_metric_id"`
	// API filter string used to filter the query. Allowed filters are flow_id, send_channel, flow_message_id. Allowed operators are equals, contains-any. Only one filter can be used per attribute, only AND can be used as a combination operator. Max of 100 messages per ANY filter.
	Filter NullableString `json:"filter,omitempty"`
}

type _FlowSeriesRequestDTOResourceObjectAttributes FlowSeriesRequestDTOResourceObjectAttributes

// NewFlowSeriesRequestDTOResourceObjectAttributes instantiates a new FlowSeriesRequestDTOResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowSeriesRequestDTOResourceObjectAttributes(statistics []string, timeframe CampaignValuesRequestDTOResourceObjectAttributesTimeframe, interval string, conversionMetricId string) *FlowSeriesRequestDTOResourceObjectAttributes {
	this := FlowSeriesRequestDTOResourceObjectAttributes{}
	this.Statistics = statistics
	this.Timeframe = timeframe
	this.Interval = interval
	this.ConversionMetricId = conversionMetricId
	return &this
}

// NewFlowSeriesRequestDTOResourceObjectAttributesWithDefaults instantiates a new FlowSeriesRequestDTOResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowSeriesRequestDTOResourceObjectAttributesWithDefaults() *FlowSeriesRequestDTOResourceObjectAttributes {
	this := FlowSeriesRequestDTOResourceObjectAttributes{}
	return &this
}

// GetStatistics returns the Statistics field value
func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetStatistics() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value
// and a boolean to check if the value has been set.
func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetStatisticsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Statistics, true
}

// SetStatistics sets field value
func (o *FlowSeriesRequestDTOResourceObjectAttributes) SetStatistics(v []string) {
	o.Statistics = v
}

// GetTimeframe returns the Timeframe field value
func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetTimeframe() CampaignValuesRequestDTOResourceObjectAttributesTimeframe {
	if o == nil {
		var ret CampaignValuesRequestDTOResourceObjectAttributesTimeframe
		return ret
	}

	return o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value
// and a boolean to check if the value has been set.
func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetTimeframeOk() (*CampaignValuesRequestDTOResourceObjectAttributesTimeframe, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeframe, true
}

// SetTimeframe sets field value
func (o *FlowSeriesRequestDTOResourceObjectAttributes) SetTimeframe(v CampaignValuesRequestDTOResourceObjectAttributesTimeframe) {
	o.Timeframe = v
}

// GetInterval returns the Interval field value
func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *FlowSeriesRequestDTOResourceObjectAttributes) SetInterval(v string) {
	o.Interval = v
}

// GetConversionMetricId returns the ConversionMetricId field value
func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetConversionMetricId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConversionMetricId
}

// GetConversionMetricIdOk returns a tuple with the ConversionMetricId field value
// and a boolean to check if the value has been set.
func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetConversionMetricIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversionMetricId, true
}

// SetConversionMetricId sets field value
func (o *FlowSeriesRequestDTOResourceObjectAttributes) SetConversionMetricId(v string) {
	o.ConversionMetricId = v
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetFilter() string {
	if o == nil || IsNil(o.Filter.Get()) {
		var ret string
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlowSeriesRequestDTOResourceObjectAttributes) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *FlowSeriesRequestDTOResourceObjectAttributes) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableString and assigns it to the Filter field.
func (o *FlowSeriesRequestDTOResourceObjectAttributes) SetFilter(v string) {
	o.Filter.Set(&v)
}
// SetFilterNil sets the value for Filter to be an explicit nil
func (o *FlowSeriesRequestDTOResourceObjectAttributes) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *FlowSeriesRequestDTOResourceObjectAttributes) UnsetFilter() {
	o.Filter.Unset()
}

func (o FlowSeriesRequestDTOResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowSeriesRequestDTOResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statistics"] = o.Statistics
	toSerialize["timeframe"] = o.Timeframe
	toSerialize["interval"] = o.Interval
	toSerialize["conversion_metric_id"] = o.ConversionMetricId
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	return toSerialize, nil
}

func (o *FlowSeriesRequestDTOResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"statistics",
		"timeframe",
		"interval",
		"conversion_metric_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFlowSeriesRequestDTOResourceObjectAttributes := _FlowSeriesRequestDTOResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlowSeriesRequestDTOResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = FlowSeriesRequestDTOResourceObjectAttributes(varFlowSeriesRequestDTOResourceObjectAttributes)

	return err
}

type NullableFlowSeriesRequestDTOResourceObjectAttributes struct {
	value *FlowSeriesRequestDTOResourceObjectAttributes
	isSet bool
}

func (v NullableFlowSeriesRequestDTOResourceObjectAttributes) Get() *FlowSeriesRequestDTOResourceObjectAttributes {
	return v.value
}

func (v *NullableFlowSeriesRequestDTOResourceObjectAttributes) Set(val *FlowSeriesRequestDTOResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowSeriesRequestDTOResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowSeriesRequestDTOResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowSeriesRequestDTOResourceObjectAttributes(val *FlowSeriesRequestDTOResourceObjectAttributes) *NullableFlowSeriesRequestDTOResourceObjectAttributes {
	return &NullableFlowSeriesRequestDTOResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableFlowSeriesRequestDTOResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowSeriesRequestDTOResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


