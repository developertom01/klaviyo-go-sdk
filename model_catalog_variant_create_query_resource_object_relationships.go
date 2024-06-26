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

// checks if the CatalogVariantCreateQueryResourceObjectRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogVariantCreateQueryResourceObjectRelationships{}

// CatalogVariantCreateQueryResourceObjectRelationships struct for CatalogVariantCreateQueryResourceObjectRelationships
type CatalogVariantCreateQueryResourceObjectRelationships struct {
	Item CatalogVariantCreateQueryResourceObjectRelationshipsItem `json:"item"`
}

type _CatalogVariantCreateQueryResourceObjectRelationships CatalogVariantCreateQueryResourceObjectRelationships

// NewCatalogVariantCreateQueryResourceObjectRelationships instantiates a new CatalogVariantCreateQueryResourceObjectRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogVariantCreateQueryResourceObjectRelationships(item CatalogVariantCreateQueryResourceObjectRelationshipsItem) *CatalogVariantCreateQueryResourceObjectRelationships {
	this := CatalogVariantCreateQueryResourceObjectRelationships{}
	this.Item = item
	return &this
}

// NewCatalogVariantCreateQueryResourceObjectRelationshipsWithDefaults instantiates a new CatalogVariantCreateQueryResourceObjectRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogVariantCreateQueryResourceObjectRelationshipsWithDefaults() *CatalogVariantCreateQueryResourceObjectRelationships {
	this := CatalogVariantCreateQueryResourceObjectRelationships{}
	return &this
}

// GetItem returns the Item field value
func (o *CatalogVariantCreateQueryResourceObjectRelationships) GetItem() CatalogVariantCreateQueryResourceObjectRelationshipsItem {
	if o == nil {
		var ret CatalogVariantCreateQueryResourceObjectRelationshipsItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantCreateQueryResourceObjectRelationships) GetItemOk() (*CatalogVariantCreateQueryResourceObjectRelationshipsItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *CatalogVariantCreateQueryResourceObjectRelationships) SetItem(v CatalogVariantCreateQueryResourceObjectRelationshipsItem) {
	o.Item = v
}

func (o CatalogVariantCreateQueryResourceObjectRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogVariantCreateQueryResourceObjectRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["item"] = o.Item
	return toSerialize, nil
}

func (o *CatalogVariantCreateQueryResourceObjectRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"item",
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

	varCatalogVariantCreateQueryResourceObjectRelationships := _CatalogVariantCreateQueryResourceObjectRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogVariantCreateQueryResourceObjectRelationships)

	if err != nil {
		return err
	}

	*o = CatalogVariantCreateQueryResourceObjectRelationships(varCatalogVariantCreateQueryResourceObjectRelationships)

	return err
}

type NullableCatalogVariantCreateQueryResourceObjectRelationships struct {
	value *CatalogVariantCreateQueryResourceObjectRelationships
	isSet bool
}

func (v NullableCatalogVariantCreateQueryResourceObjectRelationships) Get() *CatalogVariantCreateQueryResourceObjectRelationships {
	return v.value
}

func (v *NullableCatalogVariantCreateQueryResourceObjectRelationships) Set(val *CatalogVariantCreateQueryResourceObjectRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogVariantCreateQueryResourceObjectRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogVariantCreateQueryResourceObjectRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogVariantCreateQueryResourceObjectRelationships(val *CatalogVariantCreateQueryResourceObjectRelationships) *NullableCatalogVariantCreateQueryResourceObjectRelationships {
	return &NullableCatalogVariantCreateQueryResourceObjectRelationships{value: val, isSet: true}
}

func (v NullableCatalogVariantCreateQueryResourceObjectRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogVariantCreateQueryResourceObjectRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


