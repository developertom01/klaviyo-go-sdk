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
	"fmt"
)

// MetricAggregateEnum the model 'MetricAggregateEnum'
type MetricAggregateEnum string

// List of MetricAggregateEnum
const (
	METRIC_AGGREGATE MetricAggregateEnum = "metric-aggregate"
)

// All allowed values of MetricAggregateEnum enum
var AllowedMetricAggregateEnumEnumValues = []MetricAggregateEnum{
	"metric-aggregate",
}

func (v *MetricAggregateEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetricAggregateEnum(value)
	for _, existing := range AllowedMetricAggregateEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetricAggregateEnum", value)
}

// NewMetricAggregateEnumFromValue returns a pointer to a valid MetricAggregateEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetricAggregateEnumFromValue(v string) (*MetricAggregateEnum, error) {
	ev := MetricAggregateEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetricAggregateEnum: valid values are %v", v, AllowedMetricAggregateEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetricAggregateEnum) IsValid() bool {
	for _, existing := range AllowedMetricAggregateEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricAggregateEnum value
func (v MetricAggregateEnum) Ptr() *MetricAggregateEnum {
	return &v
}

type NullableMetricAggregateEnum struct {
	value *MetricAggregateEnum
	isSet bool
}

func (v NullableMetricAggregateEnum) Get() *MetricAggregateEnum {
	return v.value
}

func (v *NullableMetricAggregateEnum) Set(val *MetricAggregateEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricAggregateEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricAggregateEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricAggregateEnum(val *MetricAggregateEnum) *NullableMetricAggregateEnum {
	return &NullableMetricAggregateEnum{value: val, isSet: true}
}

func (v NullableMetricAggregateEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricAggregateEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

