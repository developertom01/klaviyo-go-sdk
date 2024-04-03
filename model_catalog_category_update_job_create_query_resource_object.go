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

// checks if the CatalogCategoryUpdateJobCreateQueryResourceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogCategoryUpdateJobCreateQueryResourceObject{}

// CatalogCategoryUpdateJobCreateQueryResourceObject struct for CatalogCategoryUpdateJobCreateQueryResourceObject
type CatalogCategoryUpdateJobCreateQueryResourceObject struct {
	Type CatalogCategoryBulkUpdateJobEnum `json:"type"`
	Attributes CatalogCategoryUpdateJobCreateQueryResourceObjectAttributes `json:"attributes"`
}

type _CatalogCategoryUpdateJobCreateQueryResourceObject CatalogCategoryUpdateJobCreateQueryResourceObject

// NewCatalogCategoryUpdateJobCreateQueryResourceObject instantiates a new CatalogCategoryUpdateJobCreateQueryResourceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogCategoryUpdateJobCreateQueryResourceObject(type_ CatalogCategoryBulkUpdateJobEnum, attributes CatalogCategoryUpdateJobCreateQueryResourceObjectAttributes) *CatalogCategoryUpdateJobCreateQueryResourceObject {
	this := CatalogCategoryUpdateJobCreateQueryResourceObject{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCatalogCategoryUpdateJobCreateQueryResourceObjectWithDefaults instantiates a new CatalogCategoryUpdateJobCreateQueryResourceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogCategoryUpdateJobCreateQueryResourceObjectWithDefaults() *CatalogCategoryUpdateJobCreateQueryResourceObject {
	this := CatalogCategoryUpdateJobCreateQueryResourceObject{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogCategoryUpdateJobCreateQueryResourceObject) GetType() CatalogCategoryBulkUpdateJobEnum {
	if o == nil {
		var ret CatalogCategoryBulkUpdateJobEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogCategoryUpdateJobCreateQueryResourceObject) GetTypeOk() (*CatalogCategoryBulkUpdateJobEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogCategoryUpdateJobCreateQueryResourceObject) SetType(v CatalogCategoryBulkUpdateJobEnum) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CatalogCategoryUpdateJobCreateQueryResourceObject) GetAttributes() CatalogCategoryUpdateJobCreateQueryResourceObjectAttributes {
	if o == nil {
		var ret CatalogCategoryUpdateJobCreateQueryResourceObjectAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CatalogCategoryUpdateJobCreateQueryResourceObject) GetAttributesOk() (*CatalogCategoryUpdateJobCreateQueryResourceObjectAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CatalogCategoryUpdateJobCreateQueryResourceObject) SetAttributes(v CatalogCategoryUpdateJobCreateQueryResourceObjectAttributes) {
	o.Attributes = v
}

func (o CatalogCategoryUpdateJobCreateQueryResourceObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogCategoryUpdateJobCreateQueryResourceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *CatalogCategoryUpdateJobCreateQueryResourceObject) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogCategoryUpdateJobCreateQueryResourceObject := _CatalogCategoryUpdateJobCreateQueryResourceObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogCategoryUpdateJobCreateQueryResourceObject)

	if err != nil {
		return err
	}

	*o = CatalogCategoryUpdateJobCreateQueryResourceObject(varCatalogCategoryUpdateJobCreateQueryResourceObject)

	return err
}

type NullableCatalogCategoryUpdateJobCreateQueryResourceObject struct {
	value *CatalogCategoryUpdateJobCreateQueryResourceObject
	isSet bool
}

func (v NullableCatalogCategoryUpdateJobCreateQueryResourceObject) Get() *CatalogCategoryUpdateJobCreateQueryResourceObject {
	return v.value
}

func (v *NullableCatalogCategoryUpdateJobCreateQueryResourceObject) Set(val *CatalogCategoryUpdateJobCreateQueryResourceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogCategoryUpdateJobCreateQueryResourceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogCategoryUpdateJobCreateQueryResourceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogCategoryUpdateJobCreateQueryResourceObject(val *CatalogCategoryUpdateJobCreateQueryResourceObject) *NullableCatalogCategoryUpdateJobCreateQueryResourceObject {
	return &NullableCatalogCategoryUpdateJobCreateQueryResourceObject{value: val, isSet: true}
}

func (v NullableCatalogCategoryUpdateJobCreateQueryResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogCategoryUpdateJobCreateQueryResourceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


