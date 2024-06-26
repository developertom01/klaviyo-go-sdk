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

// CatalogVariantEnum the model 'CatalogVariantEnum'
type CatalogVariantEnum string

// List of CatalogVariantEnum
const (
	CATALOG_VARIANT CatalogVariantEnum = "catalog-variant"
)

// All allowed values of CatalogVariantEnum enum
var AllowedCatalogVariantEnumEnumValues = []CatalogVariantEnum{
	"catalog-variant",
}

func (v *CatalogVariantEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CatalogVariantEnum(value)
	for _, existing := range AllowedCatalogVariantEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CatalogVariantEnum", value)
}

// NewCatalogVariantEnumFromValue returns a pointer to a valid CatalogVariantEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCatalogVariantEnumFromValue(v string) (*CatalogVariantEnum, error) {
	ev := CatalogVariantEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CatalogVariantEnum: valid values are %v", v, AllowedCatalogVariantEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CatalogVariantEnum) IsValid() bool {
	for _, existing := range AllowedCatalogVariantEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CatalogVariantEnum value
func (v CatalogVariantEnum) Ptr() *CatalogVariantEnum {
	return &v
}

type NullableCatalogVariantEnum struct {
	value *CatalogVariantEnum
	isSet bool
}

func (v NullableCatalogVariantEnum) Get() *CatalogVariantEnum {
	return v.value
}

func (v *NullableCatalogVariantEnum) Set(val *CatalogVariantEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogVariantEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogVariantEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogVariantEnum(val *CatalogVariantEnum) *NullableCatalogVariantEnum {
	return &NullableCatalogVariantEnum{value: val, isSet: true}
}

func (v NullableCatalogVariantEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogVariantEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

