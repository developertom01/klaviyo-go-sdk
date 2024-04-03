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

// FlowSeriesReportEnum the model 'FlowSeriesReportEnum'
type FlowSeriesReportEnum string

// List of FlowSeriesReportEnum
const (
	FLOW_SERIES_REPORT FlowSeriesReportEnum = "flow-series-report"
)

// All allowed values of FlowSeriesReportEnum enum
var AllowedFlowSeriesReportEnumEnumValues = []FlowSeriesReportEnum{
	"flow-series-report",
}

func (v *FlowSeriesReportEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FlowSeriesReportEnum(value)
	for _, existing := range AllowedFlowSeriesReportEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FlowSeriesReportEnum", value)
}

// NewFlowSeriesReportEnumFromValue returns a pointer to a valid FlowSeriesReportEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlowSeriesReportEnumFromValue(v string) (*FlowSeriesReportEnum, error) {
	ev := FlowSeriesReportEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FlowSeriesReportEnum: valid values are %v", v, AllowedFlowSeriesReportEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FlowSeriesReportEnum) IsValid() bool {
	for _, existing := range AllowedFlowSeriesReportEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlowSeriesReportEnum value
func (v FlowSeriesReportEnum) Ptr() *FlowSeriesReportEnum {
	return &v
}

type NullableFlowSeriesReportEnum struct {
	value *FlowSeriesReportEnum
	isSet bool
}

func (v NullableFlowSeriesReportEnum) Get() *FlowSeriesReportEnum {
	return v.value
}

func (v *NullableFlowSeriesReportEnum) Set(val *FlowSeriesReportEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowSeriesReportEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowSeriesReportEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowSeriesReportEnum(val *FlowSeriesReportEnum) *NullableFlowSeriesReportEnum {
	return &NullableFlowSeriesReportEnum{value: val, isSet: true}
}

func (v NullableFlowSeriesReportEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowSeriesReportEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
