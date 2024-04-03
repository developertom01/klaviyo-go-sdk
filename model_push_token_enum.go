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

// PushTokenEnum the model 'PushTokenEnum'
type PushTokenEnum string

// List of PushTokenEnum
const (
	PUSH_TOKEN PushTokenEnum = "push-token"
)

// All allowed values of PushTokenEnum enum
var AllowedPushTokenEnumEnumValues = []PushTokenEnum{
	"push-token",
}

func (v *PushTokenEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PushTokenEnum(value)
	for _, existing := range AllowedPushTokenEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PushTokenEnum", value)
}

// NewPushTokenEnumFromValue returns a pointer to a valid PushTokenEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPushTokenEnumFromValue(v string) (*PushTokenEnum, error) {
	ev := PushTokenEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PushTokenEnum: valid values are %v", v, AllowedPushTokenEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PushTokenEnum) IsValid() bool {
	for _, existing := range AllowedPushTokenEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PushTokenEnum value
func (v PushTokenEnum) Ptr() *PushTokenEnum {
	return &v
}

type NullablePushTokenEnum struct {
	value *PushTokenEnum
	isSet bool
}

func (v NullablePushTokenEnum) Get() *PushTokenEnum {
	return v.value
}

func (v *NullablePushTokenEnum) Set(val *PushTokenEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePushTokenEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePushTokenEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushTokenEnum(val *PushTokenEnum) *NullablePushTokenEnum {
	return &NullablePushTokenEnum{value: val, isSet: true}
}

func (v NullablePushTokenEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushTokenEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

