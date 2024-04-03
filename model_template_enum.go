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

// TemplateEnum the model 'TemplateEnum'
type TemplateEnum string

// List of TemplateEnum
const (
	TEMPLATE TemplateEnum = "template"
)

// All allowed values of TemplateEnum enum
var AllowedTemplateEnumEnumValues = []TemplateEnum{
	"template",
}

func (v *TemplateEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TemplateEnum(value)
	for _, existing := range AllowedTemplateEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TemplateEnum", value)
}

// NewTemplateEnumFromValue returns a pointer to a valid TemplateEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTemplateEnumFromValue(v string) (*TemplateEnum, error) {
	ev := TemplateEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TemplateEnum: valid values are %v", v, AllowedTemplateEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TemplateEnum) IsValid() bool {
	for _, existing := range AllowedTemplateEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TemplateEnum value
func (v TemplateEnum) Ptr() *TemplateEnum {
	return &v
}

type NullableTemplateEnum struct {
	value *TemplateEnum
	isSet bool
}

func (v NullableTemplateEnum) Get() *TemplateEnum {
	return v.value
}

func (v *NullableTemplateEnum) Set(val *TemplateEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateEnum(val *TemplateEnum) *NullableTemplateEnum {
	return &NullableTemplateEnum{value: val, isSet: true}
}

func (v NullableTemplateEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

