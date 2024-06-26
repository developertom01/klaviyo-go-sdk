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

// BackInStockSubscriptionEnum the model 'BackInStockSubscriptionEnum'
type BackInStockSubscriptionEnum string

// List of BackInStockSubscriptionEnum
const (
	BACK_IN_STOCK_SUBSCRIPTION BackInStockSubscriptionEnum = "back-in-stock-subscription"
)

// All allowed values of BackInStockSubscriptionEnum enum
var AllowedBackInStockSubscriptionEnumEnumValues = []BackInStockSubscriptionEnum{
	"back-in-stock-subscription",
}

func (v *BackInStockSubscriptionEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackInStockSubscriptionEnum(value)
	for _, existing := range AllowedBackInStockSubscriptionEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackInStockSubscriptionEnum", value)
}

// NewBackInStockSubscriptionEnumFromValue returns a pointer to a valid BackInStockSubscriptionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackInStockSubscriptionEnumFromValue(v string) (*BackInStockSubscriptionEnum, error) {
	ev := BackInStockSubscriptionEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackInStockSubscriptionEnum: valid values are %v", v, AllowedBackInStockSubscriptionEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackInStockSubscriptionEnum) IsValid() bool {
	for _, existing := range AllowedBackInStockSubscriptionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackInStockSubscriptionEnum value
func (v BackInStockSubscriptionEnum) Ptr() *BackInStockSubscriptionEnum {
	return &v
}

type NullableBackInStockSubscriptionEnum struct {
	value *BackInStockSubscriptionEnum
	isSet bool
}

func (v NullableBackInStockSubscriptionEnum) Get() *BackInStockSubscriptionEnum {
	return v.value
}

func (v *NullableBackInStockSubscriptionEnum) Set(val *BackInStockSubscriptionEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableBackInStockSubscriptionEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableBackInStockSubscriptionEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackInStockSubscriptionEnum(val *BackInStockSubscriptionEnum) *NullableBackInStockSubscriptionEnum {
	return &NullableBackInStockSubscriptionEnum{value: val, isSet: true}
}

func (v NullableBackInStockSubscriptionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackInStockSubscriptionEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

