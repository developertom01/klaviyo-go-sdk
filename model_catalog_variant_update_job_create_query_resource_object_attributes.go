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

// checks if the CatalogVariantUpdateJobCreateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogVariantUpdateJobCreateQueryResourceObjectAttributes{}

// CatalogVariantUpdateJobCreateQueryResourceObjectAttributes struct for CatalogVariantUpdateJobCreateQueryResourceObjectAttributes
type CatalogVariantUpdateJobCreateQueryResourceObjectAttributes struct {
	Variants CatalogVariantUpdateJobCreateQueryResourceObjectAttributesVariants `json:"variants"`
}

type _CatalogVariantUpdateJobCreateQueryResourceObjectAttributes CatalogVariantUpdateJobCreateQueryResourceObjectAttributes

// NewCatalogVariantUpdateJobCreateQueryResourceObjectAttributes instantiates a new CatalogVariantUpdateJobCreateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogVariantUpdateJobCreateQueryResourceObjectAttributes(variants CatalogVariantUpdateJobCreateQueryResourceObjectAttributesVariants) *CatalogVariantUpdateJobCreateQueryResourceObjectAttributes {
	this := CatalogVariantUpdateJobCreateQueryResourceObjectAttributes{}
	this.Variants = variants
	return &this
}

// NewCatalogVariantUpdateJobCreateQueryResourceObjectAttributesWithDefaults instantiates a new CatalogVariantUpdateJobCreateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogVariantUpdateJobCreateQueryResourceObjectAttributesWithDefaults() *CatalogVariantUpdateJobCreateQueryResourceObjectAttributes {
	this := CatalogVariantUpdateJobCreateQueryResourceObjectAttributes{}
	return &this
}

// GetVariants returns the Variants field value
func (o *CatalogVariantUpdateJobCreateQueryResourceObjectAttributes) GetVariants() CatalogVariantUpdateJobCreateQueryResourceObjectAttributesVariants {
	if o == nil {
		var ret CatalogVariantUpdateJobCreateQueryResourceObjectAttributesVariants
		return ret
	}

	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantUpdateJobCreateQueryResourceObjectAttributes) GetVariantsOk() (*CatalogVariantUpdateJobCreateQueryResourceObjectAttributesVariants, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variants, true
}

// SetVariants sets field value
func (o *CatalogVariantUpdateJobCreateQueryResourceObjectAttributes) SetVariants(v CatalogVariantUpdateJobCreateQueryResourceObjectAttributesVariants) {
	o.Variants = v
}

func (o CatalogVariantUpdateJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogVariantUpdateJobCreateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["variants"] = o.Variants
	return toSerialize, nil
}

func (o *CatalogVariantUpdateJobCreateQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"variants",
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

	varCatalogVariantUpdateJobCreateQueryResourceObjectAttributes := _CatalogVariantUpdateJobCreateQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogVariantUpdateJobCreateQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = CatalogVariantUpdateJobCreateQueryResourceObjectAttributes(varCatalogVariantUpdateJobCreateQueryResourceObjectAttributes)

	return err
}

type NullableCatalogVariantUpdateJobCreateQueryResourceObjectAttributes struct {
	value *CatalogVariantUpdateJobCreateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableCatalogVariantUpdateJobCreateQueryResourceObjectAttributes) Get() *CatalogVariantUpdateJobCreateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableCatalogVariantUpdateJobCreateQueryResourceObjectAttributes) Set(val *CatalogVariantUpdateJobCreateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogVariantUpdateJobCreateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogVariantUpdateJobCreateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogVariantUpdateJobCreateQueryResourceObjectAttributes(val *CatalogVariantUpdateJobCreateQueryResourceObjectAttributes) *NullableCatalogVariantUpdateJobCreateQueryResourceObjectAttributes {
	return &NullableCatalogVariantUpdateJobCreateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableCatalogVariantUpdateJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogVariantUpdateJobCreateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


