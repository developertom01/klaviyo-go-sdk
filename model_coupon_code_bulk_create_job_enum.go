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

// CouponCodeBulkCreateJobEnum the model 'CouponCodeBulkCreateJobEnum'
type CouponCodeBulkCreateJobEnum string

// List of CouponCodeBulkCreateJobEnum
const (
	COUPON_CODE_BULK_CREATE_JOB CouponCodeBulkCreateJobEnum = "coupon-code-bulk-create-job"
)

// All allowed values of CouponCodeBulkCreateJobEnum enum
var AllowedCouponCodeBulkCreateJobEnumEnumValues = []CouponCodeBulkCreateJobEnum{
	"coupon-code-bulk-create-job",
}

func (v *CouponCodeBulkCreateJobEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CouponCodeBulkCreateJobEnum(value)
	for _, existing := range AllowedCouponCodeBulkCreateJobEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CouponCodeBulkCreateJobEnum", value)
}

// NewCouponCodeBulkCreateJobEnumFromValue returns a pointer to a valid CouponCodeBulkCreateJobEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCouponCodeBulkCreateJobEnumFromValue(v string) (*CouponCodeBulkCreateJobEnum, error) {
	ev := CouponCodeBulkCreateJobEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CouponCodeBulkCreateJobEnum: valid values are %v", v, AllowedCouponCodeBulkCreateJobEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CouponCodeBulkCreateJobEnum) IsValid() bool {
	for _, existing := range AllowedCouponCodeBulkCreateJobEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CouponCodeBulkCreateJobEnum value
func (v CouponCodeBulkCreateJobEnum) Ptr() *CouponCodeBulkCreateJobEnum {
	return &v
}

type NullableCouponCodeBulkCreateJobEnum struct {
	value *CouponCodeBulkCreateJobEnum
	isSet bool
}

func (v NullableCouponCodeBulkCreateJobEnum) Get() *CouponCodeBulkCreateJobEnum {
	return v.value
}

func (v *NullableCouponCodeBulkCreateJobEnum) Set(val *CouponCodeBulkCreateJobEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodeBulkCreateJobEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodeBulkCreateJobEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodeBulkCreateJobEnum(val *CouponCodeBulkCreateJobEnum) *NullableCouponCodeBulkCreateJobEnum {
	return &NullableCouponCodeBulkCreateJobEnum{value: val, isSet: true}
}

func (v NullableCouponCodeBulkCreateJobEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodeBulkCreateJobEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

