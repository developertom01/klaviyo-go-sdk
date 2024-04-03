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

// CampaignRecipientEstimationJobEnum the model 'CampaignRecipientEstimationJobEnum'
type CampaignRecipientEstimationJobEnum string

// List of CampaignRecipientEstimationJobEnum
const (
	CAMPAIGN_RECIPIENT_ESTIMATION_JOB CampaignRecipientEstimationJobEnum = "campaign-recipient-estimation-job"
)

// All allowed values of CampaignRecipientEstimationJobEnum enum
var AllowedCampaignRecipientEstimationJobEnumEnumValues = []CampaignRecipientEstimationJobEnum{
	"campaign-recipient-estimation-job",
}

func (v *CampaignRecipientEstimationJobEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CampaignRecipientEstimationJobEnum(value)
	for _, existing := range AllowedCampaignRecipientEstimationJobEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CampaignRecipientEstimationJobEnum", value)
}

// NewCampaignRecipientEstimationJobEnumFromValue returns a pointer to a valid CampaignRecipientEstimationJobEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignRecipientEstimationJobEnumFromValue(v string) (*CampaignRecipientEstimationJobEnum, error) {
	ev := CampaignRecipientEstimationJobEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignRecipientEstimationJobEnum: valid values are %v", v, AllowedCampaignRecipientEstimationJobEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignRecipientEstimationJobEnum) IsValid() bool {
	for _, existing := range AllowedCampaignRecipientEstimationJobEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CampaignRecipientEstimationJobEnum value
func (v CampaignRecipientEstimationJobEnum) Ptr() *CampaignRecipientEstimationJobEnum {
	return &v
}

type NullableCampaignRecipientEstimationJobEnum struct {
	value *CampaignRecipientEstimationJobEnum
	isSet bool
}

func (v NullableCampaignRecipientEstimationJobEnum) Get() *CampaignRecipientEstimationJobEnum {
	return v.value
}

func (v *NullableCampaignRecipientEstimationJobEnum) Set(val *CampaignRecipientEstimationJobEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignRecipientEstimationJobEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignRecipientEstimationJobEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignRecipientEstimationJobEnum(val *CampaignRecipientEstimationJobEnum) *NullableCampaignRecipientEstimationJobEnum {
	return &NullableCampaignRecipientEstimationJobEnum{value: val, isSet: true}
}

func (v NullableCampaignRecipientEstimationJobEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignRecipientEstimationJobEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
