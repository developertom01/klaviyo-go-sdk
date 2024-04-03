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

// ListEnum the model 'ListEnum'
type ListEnum string

// List of ListEnum
const (
	LIST ListEnum = "list"
)

// All allowed values of ListEnum enum
var AllowedListEnumEnumValues = []ListEnum{
	"list",
}

func (v *ListEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListEnum(value)
	for _, existing := range AllowedListEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListEnum", value)
}

// NewListEnumFromValue returns a pointer to a valid ListEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListEnumFromValue(v string) (*ListEnum, error) {
	ev := ListEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListEnum: valid values are %v", v, AllowedListEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListEnum) IsValid() bool {
	for _, existing := range AllowedListEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListEnum value
func (v ListEnum) Ptr() *ListEnum {
	return &v
}

type NullableListEnum struct {
	value *ListEnum
	isSet bool
}

func (v NullableListEnum) Get() *ListEnum {
	return v.value
}

func (v *NullableListEnum) Set(val *ListEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableListEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableListEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEnum(val *ListEnum) *NullableListEnum {
	return &NullableListEnum{value: val, isSet: true}
}

func (v NullableListEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

