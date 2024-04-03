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

// checks if the ProfileIdentifierDTOResourceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileIdentifierDTOResourceObject{}

// ProfileIdentifierDTOResourceObject struct for ProfileIdentifierDTOResourceObject
type ProfileIdentifierDTOResourceObject struct {
	Type ProfileEnum `json:"type"`
	// Primary key that uniquely identifies this profile. Generated by Klaviyo.
	Id NullableString `json:"id,omitempty"`
	Attributes ProfileIdentifierDTOResourceObjectAttributes `json:"attributes"`
}

type _ProfileIdentifierDTOResourceObject ProfileIdentifierDTOResourceObject

// NewProfileIdentifierDTOResourceObject instantiates a new ProfileIdentifierDTOResourceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileIdentifierDTOResourceObject(type_ ProfileEnum, attributes ProfileIdentifierDTOResourceObjectAttributes) *ProfileIdentifierDTOResourceObject {
	this := ProfileIdentifierDTOResourceObject{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewProfileIdentifierDTOResourceObjectWithDefaults instantiates a new ProfileIdentifierDTOResourceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileIdentifierDTOResourceObjectWithDefaults() *ProfileIdentifierDTOResourceObject {
	this := ProfileIdentifierDTOResourceObject{}
	return &this
}

// GetType returns the Type field value
func (o *ProfileIdentifierDTOResourceObject) GetType() ProfileEnum {
	if o == nil {
		var ret ProfileEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProfileIdentifierDTOResourceObject) GetTypeOk() (*ProfileEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProfileIdentifierDTOResourceObject) SetType(v ProfileEnum) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileIdentifierDTOResourceObject) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileIdentifierDTOResourceObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ProfileIdentifierDTOResourceObject) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ProfileIdentifierDTOResourceObject) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ProfileIdentifierDTOResourceObject) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ProfileIdentifierDTOResourceObject) UnsetId() {
	o.Id.Unset()
}

// GetAttributes returns the Attributes field value
func (o *ProfileIdentifierDTOResourceObject) GetAttributes() ProfileIdentifierDTOResourceObjectAttributes {
	if o == nil {
		var ret ProfileIdentifierDTOResourceObjectAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ProfileIdentifierDTOResourceObject) GetAttributesOk() (*ProfileIdentifierDTOResourceObjectAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ProfileIdentifierDTOResourceObject) SetAttributes(v ProfileIdentifierDTOResourceObjectAttributes) {
	o.Attributes = v
}

func (o ProfileIdentifierDTOResourceObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileIdentifierDTOResourceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *ProfileIdentifierDTOResourceObject) UnmarshalJSON(data []byte) (err error) {
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

	varProfileIdentifierDTOResourceObject := _ProfileIdentifierDTOResourceObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProfileIdentifierDTOResourceObject)

	if err != nil {
		return err
	}

	*o = ProfileIdentifierDTOResourceObject(varProfileIdentifierDTOResourceObject)

	return err
}

type NullableProfileIdentifierDTOResourceObject struct {
	value *ProfileIdentifierDTOResourceObject
	isSet bool
}

func (v NullableProfileIdentifierDTOResourceObject) Get() *ProfileIdentifierDTOResourceObject {
	return v.value
}

func (v *NullableProfileIdentifierDTOResourceObject) Set(val *ProfileIdentifierDTOResourceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileIdentifierDTOResourceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileIdentifierDTOResourceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileIdentifierDTOResourceObject(val *ProfileIdentifierDTOResourceObject) *NullableProfileIdentifierDTOResourceObject {
	return &NullableProfileIdentifierDTOResourceObject{value: val, isSet: true}
}

func (v NullableProfileIdentifierDTOResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileIdentifierDTOResourceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

