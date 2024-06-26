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

// checks if the ListCreateQueryResourceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCreateQueryResourceObject{}

// ListCreateQueryResourceObject struct for ListCreateQueryResourceObject
type ListCreateQueryResourceObject struct {
	Type ListEnum `json:"type"`
	Attributes ListCreateQueryResourceObjectAttributes `json:"attributes"`
}

type _ListCreateQueryResourceObject ListCreateQueryResourceObject

// NewListCreateQueryResourceObject instantiates a new ListCreateQueryResourceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCreateQueryResourceObject(type_ ListEnum, attributes ListCreateQueryResourceObjectAttributes) *ListCreateQueryResourceObject {
	this := ListCreateQueryResourceObject{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewListCreateQueryResourceObjectWithDefaults instantiates a new ListCreateQueryResourceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCreateQueryResourceObjectWithDefaults() *ListCreateQueryResourceObject {
	this := ListCreateQueryResourceObject{}
	return &this
}

// GetType returns the Type field value
func (o *ListCreateQueryResourceObject) GetType() ListEnum {
	if o == nil {
		var ret ListEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListCreateQueryResourceObject) GetTypeOk() (*ListEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListCreateQueryResourceObject) SetType(v ListEnum) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ListCreateQueryResourceObject) GetAttributes() ListCreateQueryResourceObjectAttributes {
	if o == nil {
		var ret ListCreateQueryResourceObjectAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ListCreateQueryResourceObject) GetAttributesOk() (*ListCreateQueryResourceObjectAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ListCreateQueryResourceObject) SetAttributes(v ListCreateQueryResourceObjectAttributes) {
	o.Attributes = v
}

func (o ListCreateQueryResourceObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCreateQueryResourceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *ListCreateQueryResourceObject) UnmarshalJSON(data []byte) (err error) {
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

	varListCreateQueryResourceObject := _ListCreateQueryResourceObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListCreateQueryResourceObject)

	if err != nil {
		return err
	}

	*o = ListCreateQueryResourceObject(varListCreateQueryResourceObject)

	return err
}

type NullableListCreateQueryResourceObject struct {
	value *ListCreateQueryResourceObject
	isSet bool
}

func (v NullableListCreateQueryResourceObject) Get() *ListCreateQueryResourceObject {
	return v.value
}

func (v *NullableListCreateQueryResourceObject) Set(val *ListCreateQueryResourceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListCreateQueryResourceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListCreateQueryResourceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCreateQueryResourceObject(val *ListCreateQueryResourceObject) *NullableListCreateQueryResourceObject {
	return &NullableListCreateQueryResourceObject{value: val, isSet: true}
}

func (v NullableListCreateQueryResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCreateQueryResourceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


