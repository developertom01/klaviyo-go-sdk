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

// checks if the TagGroupResponseObjectResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagGroupResponseObjectResourceAttributes{}

// TagGroupResponseObjectResourceAttributes struct for TagGroupResponseObjectResourceAttributes
type TagGroupResponseObjectResourceAttributes struct {
	// The Tag Group name
	Name string `json:"name"`
	// If a tag group is non-exclusive, any given related resource (campaign, flow, etc.) can be linked to multiple tags from that tag group. If a tag group is exclusive, any given related resource can only be linked to one tag from that tag group.
	Exclusive bool `json:"exclusive"`
	// Every company automatically has one Default Tag Group. The Default Tag Group cannot be deleted, and no other Default Tag Groups can be created. This value is true for the Default Tag Group and false for all other Tag Groups.
	Default bool `json:"default"`
}

type _TagGroupResponseObjectResourceAttributes TagGroupResponseObjectResourceAttributes

// NewTagGroupResponseObjectResourceAttributes instantiates a new TagGroupResponseObjectResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagGroupResponseObjectResourceAttributes(name string, exclusive bool, default_ bool) *TagGroupResponseObjectResourceAttributes {
	this := TagGroupResponseObjectResourceAttributes{}
	this.Name = name
	this.Exclusive = exclusive
	this.Default = default_
	return &this
}

// NewTagGroupResponseObjectResourceAttributesWithDefaults instantiates a new TagGroupResponseObjectResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagGroupResponseObjectResourceAttributesWithDefaults() *TagGroupResponseObjectResourceAttributes {
	this := TagGroupResponseObjectResourceAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *TagGroupResponseObjectResourceAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagGroupResponseObjectResourceAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagGroupResponseObjectResourceAttributes) SetName(v string) {
	o.Name = v
}

// GetExclusive returns the Exclusive field value
func (o *TagGroupResponseObjectResourceAttributes) GetExclusive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exclusive
}

// GetExclusiveOk returns a tuple with the Exclusive field value
// and a boolean to check if the value has been set.
func (o *TagGroupResponseObjectResourceAttributes) GetExclusiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exclusive, true
}

// SetExclusive sets field value
func (o *TagGroupResponseObjectResourceAttributes) SetExclusive(v bool) {
	o.Exclusive = v
}

// GetDefault returns the Default field value
func (o *TagGroupResponseObjectResourceAttributes) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *TagGroupResponseObjectResourceAttributes) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *TagGroupResponseObjectResourceAttributes) SetDefault(v bool) {
	o.Default = v
}

func (o TagGroupResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagGroupResponseObjectResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["exclusive"] = o.Exclusive
	toSerialize["default"] = o.Default
	return toSerialize, nil
}

func (o *TagGroupResponseObjectResourceAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"exclusive",
		"default",
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

	varTagGroupResponseObjectResourceAttributes := _TagGroupResponseObjectResourceAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagGroupResponseObjectResourceAttributes)

	if err != nil {
		return err
	}

	*o = TagGroupResponseObjectResourceAttributes(varTagGroupResponseObjectResourceAttributes)

	return err
}

type NullableTagGroupResponseObjectResourceAttributes struct {
	value *TagGroupResponseObjectResourceAttributes
	isSet bool
}

func (v NullableTagGroupResponseObjectResourceAttributes) Get() *TagGroupResponseObjectResourceAttributes {
	return v.value
}

func (v *NullableTagGroupResponseObjectResourceAttributes) Set(val *TagGroupResponseObjectResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableTagGroupResponseObjectResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableTagGroupResponseObjectResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagGroupResponseObjectResourceAttributes(val *TagGroupResponseObjectResourceAttributes) *NullableTagGroupResponseObjectResourceAttributes {
	return &NullableTagGroupResponseObjectResourceAttributes{value: val, isSet: true}
}

func (v NullableTagGroupResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagGroupResponseObjectResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


