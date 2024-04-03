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

// checks if the ImageCreateQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageCreateQuery{}

// ImageCreateQuery struct for ImageCreateQuery
type ImageCreateQuery struct {
	Data ImageCreateQueryResourceObject `json:"data"`
}

type _ImageCreateQuery ImageCreateQuery

// NewImageCreateQuery instantiates a new ImageCreateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageCreateQuery(data ImageCreateQueryResourceObject) *ImageCreateQuery {
	this := ImageCreateQuery{}
	this.Data = data
	return &this
}

// NewImageCreateQueryWithDefaults instantiates a new ImageCreateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageCreateQueryWithDefaults() *ImageCreateQuery {
	this := ImageCreateQuery{}
	return &this
}

// GetData returns the Data field value
func (o *ImageCreateQuery) GetData() ImageCreateQueryResourceObject {
	if o == nil {
		var ret ImageCreateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ImageCreateQuery) GetDataOk() (*ImageCreateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ImageCreateQuery) SetData(v ImageCreateQueryResourceObject) {
	o.Data = v
}

func (o ImageCreateQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageCreateQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ImageCreateQuery) UnmarshalJSON(data []byte) (err error) {
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

	varImageCreateQuery := _ImageCreateQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImageCreateQuery)

	if err != nil {
		return err
	}

	*o = ImageCreateQuery(varImageCreateQuery)

	return err
}

type NullableImageCreateQuery struct {
	value *ImageCreateQuery
	isSet bool
}

func (v NullableImageCreateQuery) Get() *ImageCreateQuery {
	return v.value
}

func (v *NullableImageCreateQuery) Set(val *ImageCreateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableImageCreateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableImageCreateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageCreateQuery(val *ImageCreateQuery) *NullableImageCreateQuery {
	return &NullableImageCreateQuery{value: val, isSet: true}
}

func (v NullableImageCreateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageCreateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


