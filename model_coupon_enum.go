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

// CouponEnum the model 'CouponEnum'
type CouponEnum string

// List of CouponEnum
const (
	COUPON CouponEnum = "coupon"
)

// All allowed values of CouponEnum enum
var AllowedCouponEnumEnumValues = []CouponEnum{
	"coupon",
}

func (v *CouponEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CouponEnum(value)
	for _, existing := range AllowedCouponEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CouponEnum", value)
}

// NewCouponEnumFromValue returns a pointer to a valid CouponEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCouponEnumFromValue(v string) (*CouponEnum, error) {
	ev := CouponEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CouponEnum: valid values are %v", v, AllowedCouponEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CouponEnum) IsValid() bool {
	for _, existing := range AllowedCouponEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CouponEnum value
func (v CouponEnum) Ptr() *CouponEnum {
	return &v
}

type NullableCouponEnum struct {
	value *CouponEnum
	isSet bool
}

func (v NullableCouponEnum) Get() *CouponEnum {
	return v.value
}

func (v *NullableCouponEnum) Set(val *CouponEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponEnum(val *CouponEnum) *NullableCouponEnum {
	return &NullableCouponEnum{value: val, isSet: true}
}

func (v NullableCouponEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

