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

// ProfileSuppressionBulkCreateJobEnum the model 'ProfileSuppressionBulkCreateJobEnum'
type ProfileSuppressionBulkCreateJobEnum string

// List of ProfileSuppressionBulkCreateJobEnum
const (
	PROFILE_SUPPRESSION_BULK_CREATE_JOB ProfileSuppressionBulkCreateJobEnum = "profile-suppression-bulk-create-job"
)

// All allowed values of ProfileSuppressionBulkCreateJobEnum enum
var AllowedProfileSuppressionBulkCreateJobEnumEnumValues = []ProfileSuppressionBulkCreateJobEnum{
	"profile-suppression-bulk-create-job",
}

func (v *ProfileSuppressionBulkCreateJobEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProfileSuppressionBulkCreateJobEnum(value)
	for _, existing := range AllowedProfileSuppressionBulkCreateJobEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProfileSuppressionBulkCreateJobEnum", value)
}

// NewProfileSuppressionBulkCreateJobEnumFromValue returns a pointer to a valid ProfileSuppressionBulkCreateJobEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProfileSuppressionBulkCreateJobEnumFromValue(v string) (*ProfileSuppressionBulkCreateJobEnum, error) {
	ev := ProfileSuppressionBulkCreateJobEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProfileSuppressionBulkCreateJobEnum: valid values are %v", v, AllowedProfileSuppressionBulkCreateJobEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProfileSuppressionBulkCreateJobEnum) IsValid() bool {
	for _, existing := range AllowedProfileSuppressionBulkCreateJobEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProfileSuppressionBulkCreateJobEnum value
func (v ProfileSuppressionBulkCreateJobEnum) Ptr() *ProfileSuppressionBulkCreateJobEnum {
	return &v
}

type NullableProfileSuppressionBulkCreateJobEnum struct {
	value *ProfileSuppressionBulkCreateJobEnum
	isSet bool
}

func (v NullableProfileSuppressionBulkCreateJobEnum) Get() *ProfileSuppressionBulkCreateJobEnum {
	return v.value
}

func (v *NullableProfileSuppressionBulkCreateJobEnum) Set(val *ProfileSuppressionBulkCreateJobEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileSuppressionBulkCreateJobEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileSuppressionBulkCreateJobEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileSuppressionBulkCreateJobEnum(val *ProfileSuppressionBulkCreateJobEnum) *NullableProfileSuppressionBulkCreateJobEnum {
	return &NullableProfileSuppressionBulkCreateJobEnum{value: val, isSet: true}
}

func (v NullableProfileSuppressionBulkCreateJobEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileSuppressionBulkCreateJobEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

