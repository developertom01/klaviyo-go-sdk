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
	"bytes"
	"fmt"
)

// checks if the CatalogItemCreateQueryResourceObjectRelationshipsCategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogItemCreateQueryResourceObjectRelationshipsCategories{}

// CatalogItemCreateQueryResourceObjectRelationshipsCategories struct for CatalogItemCreateQueryResourceObjectRelationshipsCategories
type CatalogItemCreateQueryResourceObjectRelationshipsCategories struct {
	Data []GetCatalogItemCategoryListResponseCollectionDataInner `json:"data"`
}

type _CatalogItemCreateQueryResourceObjectRelationshipsCategories CatalogItemCreateQueryResourceObjectRelationshipsCategories

// NewCatalogItemCreateQueryResourceObjectRelationshipsCategories instantiates a new CatalogItemCreateQueryResourceObjectRelationshipsCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogItemCreateQueryResourceObjectRelationshipsCategories(data []GetCatalogItemCategoryListResponseCollectionDataInner) *CatalogItemCreateQueryResourceObjectRelationshipsCategories {
	this := CatalogItemCreateQueryResourceObjectRelationshipsCategories{}
	this.Data = data
	return &this
}

// NewCatalogItemCreateQueryResourceObjectRelationshipsCategoriesWithDefaults instantiates a new CatalogItemCreateQueryResourceObjectRelationshipsCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogItemCreateQueryResourceObjectRelationshipsCategoriesWithDefaults() *CatalogItemCreateQueryResourceObjectRelationshipsCategories {
	this := CatalogItemCreateQueryResourceObjectRelationshipsCategories{}
	return &this
}

// GetData returns the Data field value
func (o *CatalogItemCreateQueryResourceObjectRelationshipsCategories) GetData() []GetCatalogItemCategoryListResponseCollectionDataInner {
	if o == nil {
		var ret []GetCatalogItemCategoryListResponseCollectionDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CatalogItemCreateQueryResourceObjectRelationshipsCategories) GetDataOk() ([]GetCatalogItemCategoryListResponseCollectionDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CatalogItemCreateQueryResourceObjectRelationshipsCategories) SetData(v []GetCatalogItemCategoryListResponseCollectionDataInner) {
	o.Data = v
}

func (o CatalogItemCreateQueryResourceObjectRelationshipsCategories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogItemCreateQueryResourceObjectRelationshipsCategories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CatalogItemCreateQueryResourceObjectRelationshipsCategories) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCatalogItemCreateQueryResourceObjectRelationshipsCategories := _CatalogItemCreateQueryResourceObjectRelationshipsCategories{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogItemCreateQueryResourceObjectRelationshipsCategories)

	if err != nil {
		return err
	}

	*o = CatalogItemCreateQueryResourceObjectRelationshipsCategories(varCatalogItemCreateQueryResourceObjectRelationshipsCategories)

	return err
}

type NullableCatalogItemCreateQueryResourceObjectRelationshipsCategories struct {
	value *CatalogItemCreateQueryResourceObjectRelationshipsCategories
	isSet bool
}

func (v NullableCatalogItemCreateQueryResourceObjectRelationshipsCategories) Get() *CatalogItemCreateQueryResourceObjectRelationshipsCategories {
	return v.value
}

func (v *NullableCatalogItemCreateQueryResourceObjectRelationshipsCategories) Set(val *CatalogItemCreateQueryResourceObjectRelationshipsCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogItemCreateQueryResourceObjectRelationshipsCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogItemCreateQueryResourceObjectRelationshipsCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogItemCreateQueryResourceObjectRelationshipsCategories(val *CatalogItemCreateQueryResourceObjectRelationshipsCategories) *NullableCatalogItemCreateQueryResourceObjectRelationshipsCategories {
	return &NullableCatalogItemCreateQueryResourceObjectRelationshipsCategories{value: val, isSet: true}
}

func (v NullableCatalogItemCreateQueryResourceObjectRelationshipsCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogItemCreateQueryResourceObjectRelationshipsCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


