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

// checks if the CatalogCategoryCreateJobCreateQueryResourceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogCategoryCreateJobCreateQueryResourceObject{}

// CatalogCategoryCreateJobCreateQueryResourceObject struct for CatalogCategoryCreateJobCreateQueryResourceObject
type CatalogCategoryCreateJobCreateQueryResourceObject struct {
	Type CatalogCategoryBulkCreateJobEnum `json:"type"`
	Attributes CatalogCategoryCreateJobCreateQueryResourceObjectAttributes `json:"attributes"`
}

type _CatalogCategoryCreateJobCreateQueryResourceObject CatalogCategoryCreateJobCreateQueryResourceObject

// NewCatalogCategoryCreateJobCreateQueryResourceObject instantiates a new CatalogCategoryCreateJobCreateQueryResourceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogCategoryCreateJobCreateQueryResourceObject(type_ CatalogCategoryBulkCreateJobEnum, attributes CatalogCategoryCreateJobCreateQueryResourceObjectAttributes) *CatalogCategoryCreateJobCreateQueryResourceObject {
	this := CatalogCategoryCreateJobCreateQueryResourceObject{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCatalogCategoryCreateJobCreateQueryResourceObjectWithDefaults instantiates a new CatalogCategoryCreateJobCreateQueryResourceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogCategoryCreateJobCreateQueryResourceObjectWithDefaults() *CatalogCategoryCreateJobCreateQueryResourceObject {
	this := CatalogCategoryCreateJobCreateQueryResourceObject{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogCategoryCreateJobCreateQueryResourceObject) GetType() CatalogCategoryBulkCreateJobEnum {
	if o == nil {
		var ret CatalogCategoryBulkCreateJobEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogCategoryCreateJobCreateQueryResourceObject) GetTypeOk() (*CatalogCategoryBulkCreateJobEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogCategoryCreateJobCreateQueryResourceObject) SetType(v CatalogCategoryBulkCreateJobEnum) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CatalogCategoryCreateJobCreateQueryResourceObject) GetAttributes() CatalogCategoryCreateJobCreateQueryResourceObjectAttributes {
	if o == nil {
		var ret CatalogCategoryCreateJobCreateQueryResourceObjectAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CatalogCategoryCreateJobCreateQueryResourceObject) GetAttributesOk() (*CatalogCategoryCreateJobCreateQueryResourceObjectAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CatalogCategoryCreateJobCreateQueryResourceObject) SetAttributes(v CatalogCategoryCreateJobCreateQueryResourceObjectAttributes) {
	o.Attributes = v
}

func (o CatalogCategoryCreateJobCreateQueryResourceObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogCategoryCreateJobCreateQueryResourceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *CatalogCategoryCreateJobCreateQueryResourceObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
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

	varCatalogCategoryCreateJobCreateQueryResourceObject := _CatalogCategoryCreateJobCreateQueryResourceObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogCategoryCreateJobCreateQueryResourceObject)

	if err != nil {
		return err
	}

	*o = CatalogCategoryCreateJobCreateQueryResourceObject(varCatalogCategoryCreateJobCreateQueryResourceObject)

	return err
}

type NullableCatalogCategoryCreateJobCreateQueryResourceObject struct {
	value *CatalogCategoryCreateJobCreateQueryResourceObject
	isSet bool
}

func (v NullableCatalogCategoryCreateJobCreateQueryResourceObject) Get() *CatalogCategoryCreateJobCreateQueryResourceObject {
	return v.value
}

func (v *NullableCatalogCategoryCreateJobCreateQueryResourceObject) Set(val *CatalogCategoryCreateJobCreateQueryResourceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogCategoryCreateJobCreateQueryResourceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogCategoryCreateJobCreateQueryResourceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogCategoryCreateJobCreateQueryResourceObject(val *CatalogCategoryCreateJobCreateQueryResourceObject) *NullableCatalogCategoryCreateJobCreateQueryResourceObject {
	return &NullableCatalogCategoryCreateJobCreateQueryResourceObject{value: val, isSet: true}
}

func (v NullableCatalogCategoryCreateJobCreateQueryResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogCategoryCreateJobCreateQueryResourceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


