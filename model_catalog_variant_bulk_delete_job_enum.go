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

// CatalogVariantBulkDeleteJobEnum the model 'CatalogVariantBulkDeleteJobEnum'
type CatalogVariantBulkDeleteJobEnum string

// List of CatalogVariantBulkDeleteJobEnum
const (
	CATALOG_VARIANT_BULK_DELETE_JOB CatalogVariantBulkDeleteJobEnum = "catalog-variant-bulk-delete-job"
)

// All allowed values of CatalogVariantBulkDeleteJobEnum enum
var AllowedCatalogVariantBulkDeleteJobEnumEnumValues = []CatalogVariantBulkDeleteJobEnum{
	"catalog-variant-bulk-delete-job",
}

func (v *CatalogVariantBulkDeleteJobEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CatalogVariantBulkDeleteJobEnum(value)
	for _, existing := range AllowedCatalogVariantBulkDeleteJobEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CatalogVariantBulkDeleteJobEnum", value)
}

// NewCatalogVariantBulkDeleteJobEnumFromValue returns a pointer to a valid CatalogVariantBulkDeleteJobEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCatalogVariantBulkDeleteJobEnumFromValue(v string) (*CatalogVariantBulkDeleteJobEnum, error) {
	ev := CatalogVariantBulkDeleteJobEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CatalogVariantBulkDeleteJobEnum: valid values are %v", v, AllowedCatalogVariantBulkDeleteJobEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CatalogVariantBulkDeleteJobEnum) IsValid() bool {
	for _, existing := range AllowedCatalogVariantBulkDeleteJobEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CatalogVariantBulkDeleteJobEnum value
func (v CatalogVariantBulkDeleteJobEnum) Ptr() *CatalogVariantBulkDeleteJobEnum {
	return &v
}

type NullableCatalogVariantBulkDeleteJobEnum struct {
	value *CatalogVariantBulkDeleteJobEnum
	isSet bool
}

func (v NullableCatalogVariantBulkDeleteJobEnum) Get() *CatalogVariantBulkDeleteJobEnum {
	return v.value
}

func (v *NullableCatalogVariantBulkDeleteJobEnum) Set(val *CatalogVariantBulkDeleteJobEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogVariantBulkDeleteJobEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogVariantBulkDeleteJobEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogVariantBulkDeleteJobEnum(val *CatalogVariantBulkDeleteJobEnum) *NullableCatalogVariantBulkDeleteJobEnum {
	return &NullableCatalogVariantBulkDeleteJobEnum{value: val, isSet: true}
}

func (v NullableCatalogVariantBulkDeleteJobEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogVariantBulkDeleteJobEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
