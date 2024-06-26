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

// checks if the TemplateCreateQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateCreateQuery{}

// TemplateCreateQuery struct for TemplateCreateQuery
type TemplateCreateQuery struct {
	Data TemplateCreateQueryResourceObject `json:"data"`
}

type _TemplateCreateQuery TemplateCreateQuery

// NewTemplateCreateQuery instantiates a new TemplateCreateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateCreateQuery(data TemplateCreateQueryResourceObject) *TemplateCreateQuery {
	this := TemplateCreateQuery{}
	this.Data = data
	return &this
}

// NewTemplateCreateQueryWithDefaults instantiates a new TemplateCreateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateCreateQueryWithDefaults() *TemplateCreateQuery {
	this := TemplateCreateQuery{}
	return &this
}

// GetData returns the Data field value
func (o *TemplateCreateQuery) GetData() TemplateCreateQueryResourceObject {
	if o == nil {
		var ret TemplateCreateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TemplateCreateQuery) GetDataOk() (*TemplateCreateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TemplateCreateQuery) SetData(v TemplateCreateQueryResourceObject) {
	o.Data = v
}

func (o TemplateCreateQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateCreateQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *TemplateCreateQuery) UnmarshalJSON(data []byte) (err error) {
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

	varTemplateCreateQuery := _TemplateCreateQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTemplateCreateQuery)

	if err != nil {
		return err
	}

	*o = TemplateCreateQuery(varTemplateCreateQuery)

	return err
}

type NullableTemplateCreateQuery struct {
	value *TemplateCreateQuery
	isSet bool
}

func (v NullableTemplateCreateQuery) Get() *TemplateCreateQuery {
	return v.value
}

func (v *NullableTemplateCreateQuery) Set(val *TemplateCreateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateCreateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateCreateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateCreateQuery(val *TemplateCreateQuery) *NullableTemplateCreateQuery {
	return &NullableTemplateCreateQuery{value: val, isSet: true}
}

func (v NullableTemplateCreateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateCreateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


