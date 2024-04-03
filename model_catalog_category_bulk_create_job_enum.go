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

// CatalogCategoryBulkCreateJobEnum the model 'CatalogCategoryBulkCreateJobEnum'
type CatalogCategoryBulkCreateJobEnum string

// List of CatalogCategoryBulkCreateJobEnum
const (
	CATALOG_CATEGORY_BULK_CREATE_JOB CatalogCategoryBulkCreateJobEnum = "catalog-category-bulk-create-job"
)

// All allowed values of CatalogCategoryBulkCreateJobEnum enum
var AllowedCatalogCategoryBulkCreateJobEnumEnumValues = []CatalogCategoryBulkCreateJobEnum{
	"catalog-category-bulk-create-job",
}

func (v *CatalogCategoryBulkCreateJobEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CatalogCategoryBulkCreateJobEnum(value)
	for _, existing := range AllowedCatalogCategoryBulkCreateJobEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CatalogCategoryBulkCreateJobEnum", value)
}

// NewCatalogCategoryBulkCreateJobEnumFromValue returns a pointer to a valid CatalogCategoryBulkCreateJobEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCatalogCategoryBulkCreateJobEnumFromValue(v string) (*CatalogCategoryBulkCreateJobEnum, error) {
	ev := CatalogCategoryBulkCreateJobEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CatalogCategoryBulkCreateJobEnum: valid values are %v", v, AllowedCatalogCategoryBulkCreateJobEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CatalogCategoryBulkCreateJobEnum) IsValid() bool {
	for _, existing := range AllowedCatalogCategoryBulkCreateJobEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CatalogCategoryBulkCreateJobEnum value
func (v CatalogCategoryBulkCreateJobEnum) Ptr() *CatalogCategoryBulkCreateJobEnum {
	return &v
}

type NullableCatalogCategoryBulkCreateJobEnum struct {
	value *CatalogCategoryBulkCreateJobEnum
	isSet bool
}

func (v NullableCatalogCategoryBulkCreateJobEnum) Get() *CatalogCategoryBulkCreateJobEnum {
	return v.value
}

func (v *NullableCatalogCategoryBulkCreateJobEnum) Set(val *CatalogCategoryBulkCreateJobEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogCategoryBulkCreateJobEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogCategoryBulkCreateJobEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogCategoryBulkCreateJobEnum(val *CatalogCategoryBulkCreateJobEnum) *NullableCatalogCategoryBulkCreateJobEnum {
	return &NullableCatalogCategoryBulkCreateJobEnum{value: val, isSet: true}
}

func (v NullableCatalogCategoryBulkCreateJobEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogCategoryBulkCreateJobEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
