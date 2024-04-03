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

// checks if the PushTokenUnregisterQueryResourceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushTokenUnregisterQueryResourceObject{}

// PushTokenUnregisterQueryResourceObject struct for PushTokenUnregisterQueryResourceObject
type PushTokenUnregisterQueryResourceObject struct {
	Type PushTokenUnregisterEnum `json:"type"`
	Attributes PushTokenUnregisterQueryResourceObjectAttributes `json:"attributes"`
}

type _PushTokenUnregisterQueryResourceObject PushTokenUnregisterQueryResourceObject

// NewPushTokenUnregisterQueryResourceObject instantiates a new PushTokenUnregisterQueryResourceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushTokenUnregisterQueryResourceObject(type_ PushTokenUnregisterEnum, attributes PushTokenUnregisterQueryResourceObjectAttributes) *PushTokenUnregisterQueryResourceObject {
	this := PushTokenUnregisterQueryResourceObject{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPushTokenUnregisterQueryResourceObjectWithDefaults instantiates a new PushTokenUnregisterQueryResourceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushTokenUnregisterQueryResourceObjectWithDefaults() *PushTokenUnregisterQueryResourceObject {
	this := PushTokenUnregisterQueryResourceObject{}
	return &this
}

// GetType returns the Type field value
func (o *PushTokenUnregisterQueryResourceObject) GetType() PushTokenUnregisterEnum {
	if o == nil {
		var ret PushTokenUnregisterEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PushTokenUnregisterQueryResourceObject) GetTypeOk() (*PushTokenUnregisterEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PushTokenUnregisterQueryResourceObject) SetType(v PushTokenUnregisterEnum) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PushTokenUnregisterQueryResourceObject) GetAttributes() PushTokenUnregisterQueryResourceObjectAttributes {
	if o == nil {
		var ret PushTokenUnregisterQueryResourceObjectAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PushTokenUnregisterQueryResourceObject) GetAttributesOk() (*PushTokenUnregisterQueryResourceObjectAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PushTokenUnregisterQueryResourceObject) SetAttributes(v PushTokenUnregisterQueryResourceObjectAttributes) {
	o.Attributes = v
}

func (o PushTokenUnregisterQueryResourceObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushTokenUnregisterQueryResourceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *PushTokenUnregisterQueryResourceObject) UnmarshalJSON(data []byte) (err error) {
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

	varPushTokenUnregisterQueryResourceObject := _PushTokenUnregisterQueryResourceObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPushTokenUnregisterQueryResourceObject)

	if err != nil {
		return err
	}

	*o = PushTokenUnregisterQueryResourceObject(varPushTokenUnregisterQueryResourceObject)

	return err
}

type NullablePushTokenUnregisterQueryResourceObject struct {
	value *PushTokenUnregisterQueryResourceObject
	isSet bool
}

func (v NullablePushTokenUnregisterQueryResourceObject) Get() *PushTokenUnregisterQueryResourceObject {
	return v.value
}

func (v *NullablePushTokenUnregisterQueryResourceObject) Set(val *PushTokenUnregisterQueryResourceObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePushTokenUnregisterQueryResourceObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePushTokenUnregisterQueryResourceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushTokenUnregisterQueryResourceObject(val *PushTokenUnregisterQueryResourceObject) *NullablePushTokenUnregisterQueryResourceObject {
	return &NullablePushTokenUnregisterQueryResourceObject{value: val, isSet: true}
}

func (v NullablePushTokenUnregisterQueryResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushTokenUnregisterQueryResourceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


