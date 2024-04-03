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

// CampaignMessageEnum the model 'CampaignMessageEnum'
type CampaignMessageEnum string

// List of CampaignMessageEnum
const (
	CAMPAIGN_MESSAGE CampaignMessageEnum = "campaign-message"
)

// All allowed values of CampaignMessageEnum enum
var AllowedCampaignMessageEnumEnumValues = []CampaignMessageEnum{
	"campaign-message",
}

func (v *CampaignMessageEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CampaignMessageEnum(value)
	for _, existing := range AllowedCampaignMessageEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CampaignMessageEnum", value)
}

// NewCampaignMessageEnumFromValue returns a pointer to a valid CampaignMessageEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignMessageEnumFromValue(v string) (*CampaignMessageEnum, error) {
	ev := CampaignMessageEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignMessageEnum: valid values are %v", v, AllowedCampaignMessageEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignMessageEnum) IsValid() bool {
	for _, existing := range AllowedCampaignMessageEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CampaignMessageEnum value
func (v CampaignMessageEnum) Ptr() *CampaignMessageEnum {
	return &v
}

type NullableCampaignMessageEnum struct {
	value *CampaignMessageEnum
	isSet bool
}

func (v NullableCampaignMessageEnum) Get() *CampaignMessageEnum {
	return v.value
}

func (v *NullableCampaignMessageEnum) Set(val *CampaignMessageEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignMessageEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignMessageEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignMessageEnum(val *CampaignMessageEnum) *NullableCampaignMessageEnum {
	return &NullableCampaignMessageEnum{value: val, isSet: true}
}

func (v NullableCampaignMessageEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignMessageEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

