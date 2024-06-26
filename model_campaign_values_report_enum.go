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

// CampaignValuesReportEnum the model 'CampaignValuesReportEnum'
type CampaignValuesReportEnum string

// List of CampaignValuesReportEnum
const (
	CAMPAIGN_VALUES_REPORT CampaignValuesReportEnum = "campaign-values-report"
)

// All allowed values of CampaignValuesReportEnum enum
var AllowedCampaignValuesReportEnumEnumValues = []CampaignValuesReportEnum{
	"campaign-values-report",
}

func (v *CampaignValuesReportEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CampaignValuesReportEnum(value)
	for _, existing := range AllowedCampaignValuesReportEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CampaignValuesReportEnum", value)
}

// NewCampaignValuesReportEnumFromValue returns a pointer to a valid CampaignValuesReportEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignValuesReportEnumFromValue(v string) (*CampaignValuesReportEnum, error) {
	ev := CampaignValuesReportEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignValuesReportEnum: valid values are %v", v, AllowedCampaignValuesReportEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignValuesReportEnum) IsValid() bool {
	for _, existing := range AllowedCampaignValuesReportEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CampaignValuesReportEnum value
func (v CampaignValuesReportEnum) Ptr() *CampaignValuesReportEnum {
	return &v
}

type NullableCampaignValuesReportEnum struct {
	value *CampaignValuesReportEnum
	isSet bool
}

func (v NullableCampaignValuesReportEnum) Get() *CampaignValuesReportEnum {
	return v.value
}

func (v *NullableCampaignValuesReportEnum) Set(val *CampaignValuesReportEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignValuesReportEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignValuesReportEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignValuesReportEnum(val *CampaignValuesReportEnum) *NullableCampaignValuesReportEnum {
	return &NullableCampaignValuesReportEnum{value: val, isSet: true}
}

func (v NullableCampaignValuesReportEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignValuesReportEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

