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

// ImportErrorEnum the model 'ImportErrorEnum'
type ImportErrorEnum string

// List of ImportErrorEnum
const (
	IMPORT_ERROR ImportErrorEnum = "import-error"
)

// All allowed values of ImportErrorEnum enum
var AllowedImportErrorEnumEnumValues = []ImportErrorEnum{
	"import-error",
}

func (v *ImportErrorEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImportErrorEnum(value)
	for _, existing := range AllowedImportErrorEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImportErrorEnum", value)
}

// NewImportErrorEnumFromValue returns a pointer to a valid ImportErrorEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImportErrorEnumFromValue(v string) (*ImportErrorEnum, error) {
	ev := ImportErrorEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImportErrorEnum: valid values are %v", v, AllowedImportErrorEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImportErrorEnum) IsValid() bool {
	for _, existing := range AllowedImportErrorEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImportErrorEnum value
func (v ImportErrorEnum) Ptr() *ImportErrorEnum {
	return &v
}

type NullableImportErrorEnum struct {
	value *ImportErrorEnum
	isSet bool
}

func (v NullableImportErrorEnum) Get() *ImportErrorEnum {
	return v.value
}

func (v *NullableImportErrorEnum) Set(val *ImportErrorEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableImportErrorEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableImportErrorEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportErrorEnum(val *ImportErrorEnum) *NullableImportErrorEnum {
	return &NullableImportErrorEnum{value: val, isSet: true}
}

func (v NullableImportErrorEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportErrorEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

