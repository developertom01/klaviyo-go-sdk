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

// AccountEnum the model 'AccountEnum'
type AccountEnum string

// List of AccountEnum
const (
	ACCOUNT AccountEnum = "account"
)

// All allowed values of AccountEnum enum
var AllowedAccountEnumEnumValues = []AccountEnum{
	"account",
}

func (v *AccountEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountEnum(value)
	for _, existing := range AllowedAccountEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountEnum", value)
}

// NewAccountEnumFromValue returns a pointer to a valid AccountEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountEnumFromValue(v string) (*AccountEnum, error) {
	ev := AccountEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountEnum: valid values are %v", v, AllowedAccountEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountEnum) IsValid() bool {
	for _, existing := range AllowedAccountEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountEnum value
func (v AccountEnum) Ptr() *AccountEnum {
	return &v
}

type NullableAccountEnum struct {
	value *AccountEnum
	isSet bool
}

func (v NullableAccountEnum) Get() *AccountEnum {
	return v.value
}

func (v *NullableAccountEnum) Set(val *AccountEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountEnum(val *AccountEnum) *NullableAccountEnum {
	return &NullableAccountEnum{value: val, isSet: true}
}

func (v NullableAccountEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

