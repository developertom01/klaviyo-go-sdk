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

// checks if the CatalogCategoryCreateJobCreateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogCategoryCreateJobCreateQueryResourceObjectAttributes{}

// CatalogCategoryCreateJobCreateQueryResourceObjectAttributes struct for CatalogCategoryCreateJobCreateQueryResourceObjectAttributes
type CatalogCategoryCreateJobCreateQueryResourceObjectAttributes struct {
	Categories CatalogCategoryCreateJobCreateQueryResourceObjectAttributesCategories `json:"categories"`
}

type _CatalogCategoryCreateJobCreateQueryResourceObjectAttributes CatalogCategoryCreateJobCreateQueryResourceObjectAttributes

// NewCatalogCategoryCreateJobCreateQueryResourceObjectAttributes instantiates a new CatalogCategoryCreateJobCreateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogCategoryCreateJobCreateQueryResourceObjectAttributes(categories CatalogCategoryCreateJobCreateQueryResourceObjectAttributesCategories) *CatalogCategoryCreateJobCreateQueryResourceObjectAttributes {
	this := CatalogCategoryCreateJobCreateQueryResourceObjectAttributes{}
	this.Categories = categories
	return &this
}

// NewCatalogCategoryCreateJobCreateQueryResourceObjectAttributesWithDefaults instantiates a new CatalogCategoryCreateJobCreateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogCategoryCreateJobCreateQueryResourceObjectAttributesWithDefaults() *CatalogCategoryCreateJobCreateQueryResourceObjectAttributes {
	this := CatalogCategoryCreateJobCreateQueryResourceObjectAttributes{}
	return &this
}

// GetCategories returns the Categories field value
func (o *CatalogCategoryCreateJobCreateQueryResourceObjectAttributes) GetCategories() CatalogCategoryCreateJobCreateQueryResourceObjectAttributesCategories {
	if o == nil {
		var ret CatalogCategoryCreateJobCreateQueryResourceObjectAttributesCategories
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *CatalogCategoryCreateJobCreateQueryResourceObjectAttributes) GetCategoriesOk() (*CatalogCategoryCreateJobCreateQueryResourceObjectAttributesCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Categories, true
}

// SetCategories sets field value
func (o *CatalogCategoryCreateJobCreateQueryResourceObjectAttributes) SetCategories(v CatalogCategoryCreateJobCreateQueryResourceObjectAttributesCategories) {
	o.Categories = v
}

func (o CatalogCategoryCreateJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogCategoryCreateJobCreateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["categories"] = o.Categories
	return toSerialize, nil
}

func (o *CatalogCategoryCreateJobCreateQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"categories",
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

	varCatalogCategoryCreateJobCreateQueryResourceObjectAttributes := _CatalogCategoryCreateJobCreateQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogCategoryCreateJobCreateQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = CatalogCategoryCreateJobCreateQueryResourceObjectAttributes(varCatalogCategoryCreateJobCreateQueryResourceObjectAttributes)

	return err
}

type NullableCatalogCategoryCreateJobCreateQueryResourceObjectAttributes struct {
	value *CatalogCategoryCreateJobCreateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableCatalogCategoryCreateJobCreateQueryResourceObjectAttributes) Get() *CatalogCategoryCreateJobCreateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableCatalogCategoryCreateJobCreateQueryResourceObjectAttributes) Set(val *CatalogCategoryCreateJobCreateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogCategoryCreateJobCreateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogCategoryCreateJobCreateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogCategoryCreateJobCreateQueryResourceObjectAttributes(val *CatalogCategoryCreateJobCreateQueryResourceObjectAttributes) *NullableCatalogCategoryCreateJobCreateQueryResourceObjectAttributes {
	return &NullableCatalogCategoryCreateJobCreateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableCatalogCategoryCreateJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogCategoryCreateJobCreateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

