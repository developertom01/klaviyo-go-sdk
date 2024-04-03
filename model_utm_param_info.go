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

// checks if the UtmParamInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtmParamInfo{}

// UtmParamInfo struct for UtmParamInfo
type UtmParamInfo struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type _UtmParamInfo UtmParamInfo

// NewUtmParamInfo instantiates a new UtmParamInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtmParamInfo(name string, value string) *UtmParamInfo {
	this := UtmParamInfo{}
	this.Name = name
	this.Value = value
	return &this
}

// NewUtmParamInfoWithDefaults instantiates a new UtmParamInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtmParamInfoWithDefaults() *UtmParamInfo {
	this := UtmParamInfo{}
	return &this
}

// GetName returns the Name field value
func (o *UtmParamInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UtmParamInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UtmParamInfo) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *UtmParamInfo) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UtmParamInfo) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UtmParamInfo) SetValue(v string) {
	o.Value = v
}

func (o UtmParamInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtmParamInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *UtmParamInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varUtmParamInfo := _UtmParamInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUtmParamInfo)

	if err != nil {
		return err
	}

	*o = UtmParamInfo(varUtmParamInfo)

	return err
}

type NullableUtmParamInfo struct {
	value *UtmParamInfo
	isSet bool
}

func (v NullableUtmParamInfo) Get() *UtmParamInfo {
	return v.value
}

func (v *NullableUtmParamInfo) Set(val *UtmParamInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUtmParamInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUtmParamInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtmParamInfo(val *UtmParamInfo) *NullableUtmParamInfo {
	return &NullableUtmParamInfo{value: val, isSet: true}
}

func (v NullableUtmParamInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtmParamInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

