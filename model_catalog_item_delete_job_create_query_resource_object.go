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

// checks if the CatalogItemDeleteJobCreateQueryResourceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogItemDeleteJobCreateQueryResourceObject{}

// CatalogItemDeleteJobCreateQueryResourceObject struct for CatalogItemDeleteJobCreateQueryResourceObject
type CatalogItemDeleteJobCreateQueryResourceObject struct {
	Type CatalogItemBulkDeleteJobEnum `json:"type"`
	Attributes CatalogItemDeleteJobCreateQueryResourceObjectAttributes `json:"attributes"`
}

type _CatalogItemDeleteJobCreateQueryResourceObject CatalogItemDeleteJobCreateQueryResourceObject

// NewCatalogItemDeleteJobCreateQueryResourceObject instantiates a new CatalogItemDeleteJobCreateQueryResourceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogItemDeleteJobCreateQueryResourceObject(type_ CatalogItemBulkDeleteJobEnum, attributes CatalogItemDeleteJobCreateQueryResourceObjectAttributes) *CatalogItemDeleteJobCreateQueryResourceObject {
	this := CatalogItemDeleteJobCreateQueryResourceObject{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCatalogItemDeleteJobCreateQueryResourceObjectWithDefaults instantiates a new CatalogItemDeleteJobCreateQueryResourceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogItemDeleteJobCreateQueryResourceObjectWithDefaults() *CatalogItemDeleteJobCreateQueryResourceObject {
	this := CatalogItemDeleteJobCreateQueryResourceObject{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogItemDeleteJobCreateQueryResourceObject) GetType() CatalogItemBulkDeleteJobEnum {
	if o == nil {
		var ret CatalogItemBulkDeleteJobEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogItemDeleteJobCreateQueryResourceObject) GetTypeOk() (*CatalogItemBulkDeleteJobEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogItemDeleteJobCreateQueryResourceObject) SetType(v CatalogItemBulkDeleteJobEnum) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CatalogItemDeleteJobCreateQueryResourceObject) GetAttributes() CatalogItemDeleteJobCreateQueryResourceObjectAttributes {
	if o == nil {
		var ret CatalogItemDeleteJobCreateQueryResourceObjectAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CatalogItemDeleteJobCreateQueryResourceObject) GetAttributesOk() (*CatalogItemDeleteJobCreateQueryResourceObjectAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CatalogItemDeleteJobCreateQueryResourceObject) SetAttributes(v CatalogItemDeleteJobCreateQueryResourceObjectAttributes) {
	o.Attributes = v
}

func (o CatalogItemDeleteJobCreateQueryResourceObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogItemDeleteJobCreateQueryResourceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *CatalogItemDeleteJobCreateQueryResourceObject) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogItemDeleteJobCreateQueryResourceObject := _CatalogItemDeleteJobCreateQueryResourceObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogItemDeleteJobCreateQueryResourceObject)

	if err != nil {
		return err
	}

	*o = CatalogItemDeleteJobCreateQueryResourceObject(varCatalogItemDeleteJobCreateQueryResourceObject)

	return err
}

type NullableCatalogItemDeleteJobCreateQueryResourceObject struct {
	value *CatalogItemDeleteJobCreateQueryResourceObject
	isSet bool
}

func (v NullableCatalogItemDeleteJobCreateQueryResourceObject) Get() *CatalogItemDeleteJobCreateQueryResourceObject {
	return v.value
}

func (v *NullableCatalogItemDeleteJobCreateQueryResourceObject) Set(val *CatalogItemDeleteJobCreateQueryResourceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogItemDeleteJobCreateQueryResourceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogItemDeleteJobCreateQueryResourceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogItemDeleteJobCreateQueryResourceObject(val *CatalogItemDeleteJobCreateQueryResourceObject) *NullableCatalogItemDeleteJobCreateQueryResourceObject {
	return &NullableCatalogItemDeleteJobCreateQueryResourceObject{value: val, isSet: true}
}

func (v NullableCatalogItemDeleteJobCreateQueryResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogItemDeleteJobCreateQueryResourceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


