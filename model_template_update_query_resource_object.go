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

// checks if the TemplateUpdateQueryResourceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateUpdateQueryResourceObject{}

// TemplateUpdateQueryResourceObject struct for TemplateUpdateQueryResourceObject
type TemplateUpdateQueryResourceObject struct {
	Type TemplateEnum `json:"type"`
	// The ID of template
	Id string `json:"id"`
	Attributes TemplateUpdateQueryResourceObjectAttributes `json:"attributes"`
}

type _TemplateUpdateQueryResourceObject TemplateUpdateQueryResourceObject

// NewTemplateUpdateQueryResourceObject instantiates a new TemplateUpdateQueryResourceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateUpdateQueryResourceObject(type_ TemplateEnum, id string, attributes TemplateUpdateQueryResourceObjectAttributes) *TemplateUpdateQueryResourceObject {
	this := TemplateUpdateQueryResourceObject{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewTemplateUpdateQueryResourceObjectWithDefaults instantiates a new TemplateUpdateQueryResourceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateUpdateQueryResourceObjectWithDefaults() *TemplateUpdateQueryResourceObject {
	this := TemplateUpdateQueryResourceObject{}
	return &this
}

// GetType returns the Type field value
func (o *TemplateUpdateQueryResourceObject) GetType() TemplateEnum {
	if o == nil {
		var ret TemplateEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TemplateUpdateQueryResourceObject) GetTypeOk() (*TemplateEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TemplateUpdateQueryResourceObject) SetType(v TemplateEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TemplateUpdateQueryResourceObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemplateUpdateQueryResourceObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemplateUpdateQueryResourceObject) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *TemplateUpdateQueryResourceObject) GetAttributes() TemplateUpdateQueryResourceObjectAttributes {
	if o == nil {
		var ret TemplateUpdateQueryResourceObjectAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TemplateUpdateQueryResourceObject) GetAttributesOk() (*TemplateUpdateQueryResourceObjectAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TemplateUpdateQueryResourceObject) SetAttributes(v TemplateUpdateQueryResourceObjectAttributes) {
	o.Attributes = v
}

func (o TemplateUpdateQueryResourceObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateUpdateQueryResourceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *TemplateUpdateQueryResourceObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varTemplateUpdateQueryResourceObject := _TemplateUpdateQueryResourceObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTemplateUpdateQueryResourceObject)

	if err != nil {
		return err
	}

	*o = TemplateUpdateQueryResourceObject(varTemplateUpdateQueryResourceObject)

	return err
}

type NullableTemplateUpdateQueryResourceObject struct {
	value *TemplateUpdateQueryResourceObject
	isSet bool
}

func (v NullableTemplateUpdateQueryResourceObject) Get() *TemplateUpdateQueryResourceObject {
	return v.value
}

func (v *NullableTemplateUpdateQueryResourceObject) Set(val *TemplateUpdateQueryResourceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateUpdateQueryResourceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateUpdateQueryResourceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateUpdateQueryResourceObject(val *TemplateUpdateQueryResourceObject) *NullableTemplateUpdateQueryResourceObject {
	return &NullableTemplateUpdateQueryResourceObject{value: val, isSet: true}
}

func (v NullableTemplateUpdateQueryResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateUpdateQueryResourceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


