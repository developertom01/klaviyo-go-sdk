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

// CatalogItemBulkDeleteJobEnum the model 'CatalogItemBulkDeleteJobEnum'
type CatalogItemBulkDeleteJobEnum string

// List of CatalogItemBulkDeleteJobEnum
const (
	CATALOG_ITEM_BULK_DELETE_JOB CatalogItemBulkDeleteJobEnum = "catalog-item-bulk-delete-job"
)

// All allowed values of CatalogItemBulkDeleteJobEnum enum
var AllowedCatalogItemBulkDeleteJobEnumEnumValues = []CatalogItemBulkDeleteJobEnum{
	"catalog-item-bulk-delete-job",
}

func (v *CatalogItemBulkDeleteJobEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CatalogItemBulkDeleteJobEnum(value)
	for _, existing := range AllowedCatalogItemBulkDeleteJobEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CatalogItemBulkDeleteJobEnum", value)
}

// NewCatalogItemBulkDeleteJobEnumFromValue returns a pointer to a valid CatalogItemBulkDeleteJobEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCatalogItemBulkDeleteJobEnumFromValue(v string) (*CatalogItemBulkDeleteJobEnum, error) {
	ev := CatalogItemBulkDeleteJobEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CatalogItemBulkDeleteJobEnum: valid values are %v", v, AllowedCatalogItemBulkDeleteJobEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CatalogItemBulkDeleteJobEnum) IsValid() bool {
	for _, existing := range AllowedCatalogItemBulkDeleteJobEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CatalogItemBulkDeleteJobEnum value
func (v CatalogItemBulkDeleteJobEnum) Ptr() *CatalogItemBulkDeleteJobEnum {
	return &v
}

type NullableCatalogItemBulkDeleteJobEnum struct {
	value *CatalogItemBulkDeleteJobEnum
	isSet bool
}

func (v NullableCatalogItemBulkDeleteJobEnum) Get() *CatalogItemBulkDeleteJobEnum {
	return v.value
}

func (v *NullableCatalogItemBulkDeleteJobEnum) Set(val *CatalogItemBulkDeleteJobEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogItemBulkDeleteJobEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogItemBulkDeleteJobEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogItemBulkDeleteJobEnum(val *CatalogItemBulkDeleteJobEnum) *NullableCatalogItemBulkDeleteJobEnum {
	return &NullableCatalogItemBulkDeleteJobEnum{value: val, isSet: true}
}

func (v NullableCatalogItemBulkDeleteJobEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogItemBulkDeleteJobEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

