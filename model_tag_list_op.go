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

// checks if the TagListOp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagListOp{}

// TagListOp struct for TagListOp
type TagListOp struct {
	Data []TagListOpDataInner `json:"data"`
}

type _TagListOp TagListOp

// NewTagListOp instantiates a new TagListOp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagListOp(data []TagListOpDataInner) *TagListOp {
	this := TagListOp{}
	this.Data = data
	return &this
}

// NewTagListOpWithDefaults instantiates a new TagListOp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagListOpWithDefaults() *TagListOp {
	this := TagListOp{}
	return &this
}

// GetData returns the Data field value
func (o *TagListOp) GetData() []TagListOpDataInner {
	if o == nil {
		var ret []TagListOpDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TagListOp) GetDataOk() ([]TagListOpDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *TagListOp) SetData(v []TagListOpDataInner) {
	o.Data = v
}

func (o TagListOp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagListOp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *TagListOp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varTagListOp := _TagListOp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagListOp)

	if err != nil {
		return err
	}

	*o = TagListOp(varTagListOp)

	return err
}

type NullableTagListOp struct {
	value *TagListOp
	isSet bool
}

func (v NullableTagListOp) Get() *TagListOp {
	return v.value
}

func (v *NullableTagListOp) Set(val *TagListOp) {
	v.value = val
	v.isSet = true
}

func (v NullableTagListOp) IsSet() bool {
	return v.isSet
}

func (v *NullableTagListOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagListOp(val *TagListOp) *NullableTagListOp {
	return &NullableTagListOp{value: val, isSet: true}
}

func (v NullableTagListOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagListOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


