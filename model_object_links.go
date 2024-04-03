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

// checks if the ObjectLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectLinks{}

// ObjectLinks struct for ObjectLinks
type ObjectLinks struct {
	Self string `json:"self"`
}

type _ObjectLinks ObjectLinks

// NewObjectLinks instantiates a new ObjectLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectLinks(self string) *ObjectLinks {
	this := ObjectLinks{}
	this.Self = self
	return &this
}

// NewObjectLinksWithDefaults instantiates a new ObjectLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectLinksWithDefaults() *ObjectLinks {
	this := ObjectLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *ObjectLinks) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ObjectLinks) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *ObjectLinks) SetSelf(v string) {
	o.Self = v
}

func (o ObjectLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	return toSerialize, nil
}

func (o *ObjectLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
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

	varObjectLinks := _ObjectLinks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectLinks)

	if err != nil {
		return err
	}

	*o = ObjectLinks(varObjectLinks)

	return err
}

type NullableObjectLinks struct {
	value *ObjectLinks
	isSet bool
}

func (v NullableObjectLinks) Get() *ObjectLinks {
	return v.value
}

func (v *NullableObjectLinks) Set(val *ObjectLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectLinks(val *ObjectLinks) *NullableObjectLinks {
	return &NullableObjectLinks{value: val, isSet: true}
}

func (v NullableObjectLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


