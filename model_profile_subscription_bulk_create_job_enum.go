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

// ProfileSubscriptionBulkCreateJobEnum the model 'ProfileSubscriptionBulkCreateJobEnum'
type ProfileSubscriptionBulkCreateJobEnum string

// List of ProfileSubscriptionBulkCreateJobEnum
const (
	PROFILE_SUBSCRIPTION_BULK_CREATE_JOB ProfileSubscriptionBulkCreateJobEnum = "profile-subscription-bulk-create-job"
)

// All allowed values of ProfileSubscriptionBulkCreateJobEnum enum
var AllowedProfileSubscriptionBulkCreateJobEnumEnumValues = []ProfileSubscriptionBulkCreateJobEnum{
	"profile-subscription-bulk-create-job",
}

func (v *ProfileSubscriptionBulkCreateJobEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProfileSubscriptionBulkCreateJobEnum(value)
	for _, existing := range AllowedProfileSubscriptionBulkCreateJobEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProfileSubscriptionBulkCreateJobEnum", value)
}

// NewProfileSubscriptionBulkCreateJobEnumFromValue returns a pointer to a valid ProfileSubscriptionBulkCreateJobEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProfileSubscriptionBulkCreateJobEnumFromValue(v string) (*ProfileSubscriptionBulkCreateJobEnum, error) {
	ev := ProfileSubscriptionBulkCreateJobEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProfileSubscriptionBulkCreateJobEnum: valid values are %v", v, AllowedProfileSubscriptionBulkCreateJobEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProfileSubscriptionBulkCreateJobEnum) IsValid() bool {
	for _, existing := range AllowedProfileSubscriptionBulkCreateJobEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProfileSubscriptionBulkCreateJobEnum value
func (v ProfileSubscriptionBulkCreateJobEnum) Ptr() *ProfileSubscriptionBulkCreateJobEnum {
	return &v
}

type NullableProfileSubscriptionBulkCreateJobEnum struct {
	value *ProfileSubscriptionBulkCreateJobEnum
	isSet bool
}

func (v NullableProfileSubscriptionBulkCreateJobEnum) Get() *ProfileSubscriptionBulkCreateJobEnum {
	return v.value
}

func (v *NullableProfileSubscriptionBulkCreateJobEnum) Set(val *ProfileSubscriptionBulkCreateJobEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileSubscriptionBulkCreateJobEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileSubscriptionBulkCreateJobEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileSubscriptionBulkCreateJobEnum(val *ProfileSubscriptionBulkCreateJobEnum) *NullableProfileSubscriptionBulkCreateJobEnum {
	return &NullableProfileSubscriptionBulkCreateJobEnum{value: val, isSet: true}
}

func (v NullableProfileSubscriptionBulkCreateJobEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileSubscriptionBulkCreateJobEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

