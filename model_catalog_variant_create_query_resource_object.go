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

// checks if the CatalogVariantCreateQueryResourceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogVariantCreateQueryResourceObject{}

// CatalogVariantCreateQueryResourceObject struct for CatalogVariantCreateQueryResourceObject
type CatalogVariantCreateQueryResourceObject struct {
	Type CatalogVariantEnum `json:"type"`
	Attributes CatalogVariantCreateQueryResourceObjectAttributes `json:"attributes"`
	Relationships CatalogVariantCreateQueryResourceObjectRelationships `json:"relationships"`
}

type _CatalogVariantCreateQueryResourceObject CatalogVariantCreateQueryResourceObject

// NewCatalogVariantCreateQueryResourceObject instantiates a new CatalogVariantCreateQueryResourceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogVariantCreateQueryResourceObject(type_ CatalogVariantEnum, attributes CatalogVariantCreateQueryResourceObjectAttributes, relationships CatalogVariantCreateQueryResourceObjectRelationships) *CatalogVariantCreateQueryResourceObject {
	this := CatalogVariantCreateQueryResourceObject{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewCatalogVariantCreateQueryResourceObjectWithDefaults instantiates a new CatalogVariantCreateQueryResourceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogVariantCreateQueryResourceObjectWithDefaults() *CatalogVariantCreateQueryResourceObject {
	this := CatalogVariantCreateQueryResourceObject{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogVariantCreateQueryResourceObject) GetType() CatalogVariantEnum {
	if o == nil {
		var ret CatalogVariantEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantCreateQueryResourceObject) GetTypeOk() (*CatalogVariantEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogVariantCreateQueryResourceObject) SetType(v CatalogVariantEnum) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CatalogVariantCreateQueryResourceObject) GetAttributes() CatalogVariantCreateQueryResourceObjectAttributes {
	if o == nil {
		var ret CatalogVariantCreateQueryResourceObjectAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantCreateQueryResourceObject) GetAttributesOk() (*CatalogVariantCreateQueryResourceObjectAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CatalogVariantCreateQueryResourceObject) SetAttributes(v CatalogVariantCreateQueryResourceObjectAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *CatalogVariantCreateQueryResourceObject) GetRelationships() CatalogVariantCreateQueryResourceObjectRelationships {
	if o == nil {
		var ret CatalogVariantCreateQueryResourceObjectRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantCreateQueryResourceObject) GetRelationshipsOk() (*CatalogVariantCreateQueryResourceObjectRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *CatalogVariantCreateQueryResourceObject) SetRelationships(v CatalogVariantCreateQueryResourceObjectRelationships) {
	o.Relationships = v
}

func (o CatalogVariantCreateQueryResourceObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogVariantCreateQueryResourceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *CatalogVariantCreateQueryResourceObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
		"relationships",
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

	varCatalogVariantCreateQueryResourceObject := _CatalogVariantCreateQueryResourceObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogVariantCreateQueryResourceObject)

	if err != nil {
		return err
	}

	*o = CatalogVariantCreateQueryResourceObject(varCatalogVariantCreateQueryResourceObject)

	return err
}

type NullableCatalogVariantCreateQueryResourceObject struct {
	value *CatalogVariantCreateQueryResourceObject
	isSet bool
}

func (v NullableCatalogVariantCreateQueryResourceObject) Get() *CatalogVariantCreateQueryResourceObject {
	return v.value
}

func (v *NullableCatalogVariantCreateQueryResourceObject) Set(val *CatalogVariantCreateQueryResourceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogVariantCreateQueryResourceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogVariantCreateQueryResourceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogVariantCreateQueryResourceObject(val *CatalogVariantCreateQueryResourceObject) *NullableCatalogVariantCreateQueryResourceObject {
	return &NullableCatalogVariantCreateQueryResourceObject{value: val, isSet: true}
}

func (v NullableCatalogVariantCreateQueryResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogVariantCreateQueryResourceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


