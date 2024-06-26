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

// checks if the FlowValuesRequestDTOResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowValuesRequestDTOResourceObjectAttributes{}

// FlowValuesRequestDTOResourceObjectAttributes struct for FlowValuesRequestDTOResourceObjectAttributes
type FlowValuesRequestDTOResourceObjectAttributes struct {
	// List of statistics to query for. All rate statistics will be returned in fractional form [0.0, 1.0]
	Statistics []string `json:"statistics"`
	Timeframe CampaignValuesRequestDTOResourceObjectAttributesTimeframe `json:"timeframe"`
	// ID of the metric to be used for conversion statistics
	ConversionMetricId string `json:"conversion_metric_id"`
	// API filter string used to filter the query. Allowed filters are flow_id, send_channel, flow_message_id. Allowed operators are equals, contains-any. Only one filter can be used per attribute, only AND can be used as a combination operator. Max of 100 messages per ANY filter.
	Filter NullableString `json:"filter,omitempty"`
}

type _FlowValuesRequestDTOResourceObjectAttributes FlowValuesRequestDTOResourceObjectAttributes

// NewFlowValuesRequestDTOResourceObjectAttributes instantiates a new FlowValuesRequestDTOResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowValuesRequestDTOResourceObjectAttributes(statistics []string, timeframe CampaignValuesRequestDTOResourceObjectAttributesTimeframe, conversionMetricId string) *FlowValuesRequestDTOResourceObjectAttributes {
	this := FlowValuesRequestDTOResourceObjectAttributes{}
	this.Statistics = statistics
	this.Timeframe = timeframe
	this.ConversionMetricId = conversionMetricId
	return &this
}

// NewFlowValuesRequestDTOResourceObjectAttributesWithDefaults instantiates a new FlowValuesRequestDTOResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowValuesRequestDTOResourceObjectAttributesWithDefaults() *FlowValuesRequestDTOResourceObjectAttributes {
	this := FlowValuesRequestDTOResourceObjectAttributes{}
	return &this
}

// GetStatistics returns the Statistics field value
func (o *FlowValuesRequestDTOResourceObjectAttributes) GetStatistics() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value
// and a boolean to check if the value has been set.
func (o *FlowValuesRequestDTOResourceObjectAttributes) GetStatisticsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Statistics, true
}

// SetStatistics sets field value
func (o *FlowValuesRequestDTOResourceObjectAttributes) SetStatistics(v []string) {
	o.Statistics = v
}

// GetTimeframe returns the Timeframe field value
func (o *FlowValuesRequestDTOResourceObjectAttributes) GetTimeframe() CampaignValuesRequestDTOResourceObjectAttributesTimeframe {
	if o == nil {
		var ret CampaignValuesRequestDTOResourceObjectAttributesTimeframe
		return ret
	}

	return o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value
// and a boolean to check if the value has been set.
func (o *FlowValuesRequestDTOResourceObjectAttributes) GetTimeframeOk() (*CampaignValuesRequestDTOResourceObjectAttributesTimeframe, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeframe, true
}

// SetTimeframe sets field value
func (o *FlowValuesRequestDTOResourceObjectAttributes) SetTimeframe(v CampaignValuesRequestDTOResourceObjectAttributesTimeframe) {
	o.Timeframe = v
}

// GetConversionMetricId returns the ConversionMetricId field value
func (o *FlowValuesRequestDTOResourceObjectAttributes) GetConversionMetricId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConversionMetricId
}

// GetConversionMetricIdOk returns a tuple with the ConversionMetricId field value
// and a boolean to check if the value has been set.
func (o *FlowValuesRequestDTOResourceObjectAttributes) GetConversionMetricIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversionMetricId, true
}

// SetConversionMetricId sets field value
func (o *FlowValuesRequestDTOResourceObjectAttributes) SetConversionMetricId(v string) {
	o.ConversionMetricId = v
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlowValuesRequestDTOResourceObjectAttributes) GetFilter() string {
	if o == nil || IsNil(o.Filter.Get()) {
		var ret string
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlowValuesRequestDTOResourceObjectAttributes) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *FlowValuesRequestDTOResourceObjectAttributes) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableString and assigns it to the Filter field.
func (o *FlowValuesRequestDTOResourceObjectAttributes) SetFilter(v string) {
	o.Filter.Set(&v)
}
// SetFilterNil sets the value for Filter to be an explicit nil
func (o *FlowValuesRequestDTOResourceObjectAttributes) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *FlowValuesRequestDTOResourceObjectAttributes) UnsetFilter() {
	o.Filter.Unset()
}

func (o FlowValuesRequestDTOResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowValuesRequestDTOResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statistics"] = o.Statistics
	toSerialize["timeframe"] = o.Timeframe
	toSerialize["conversion_metric_id"] = o.ConversionMetricId
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	return toSerialize, nil
}

func (o *FlowValuesRequestDTOResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"statistics",
		"timeframe",
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

	varFlowValuesRequestDTOResourceObjectAttributes := _FlowValuesRequestDTOResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlowValuesRequestDTOResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = FlowValuesRequestDTOResourceObjectAttributes(varFlowValuesRequestDTOResourceObjectAttributes)

	return err
}

type NullableFlowValuesRequestDTOResourceObjectAttributes struct {
	value *FlowValuesRequestDTOResourceObjectAttributes
	isSet bool
}

func (v NullableFlowValuesRequestDTOResourceObjectAttributes) Get() *FlowValuesRequestDTOResourceObjectAttributes {
	return v.value
}

func (v *NullableFlowValuesRequestDTOResourceObjectAttributes) Set(val *FlowValuesRequestDTOResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowValuesRequestDTOResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowValuesRequestDTOResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowValuesRequestDTOResourceObjectAttributes(val *FlowValuesRequestDTOResourceObjectAttributes) *NullableFlowValuesRequestDTOResourceObjectAttributes {
	return &NullableFlowValuesRequestDTOResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableFlowValuesRequestDTOResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowValuesRequestDTOResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


