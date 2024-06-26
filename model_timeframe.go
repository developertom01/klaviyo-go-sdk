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

// checks if the Timeframe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Timeframe{}

// Timeframe struct for Timeframe
type Timeframe struct {
	// Pre-defined key that represents a set timeframe
	Key string `json:"key"`
}

type _Timeframe Timeframe

// NewTimeframe instantiates a new Timeframe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeframe(key string) *Timeframe {
	this := Timeframe{}
	this.Key = key
	return &this
}

// NewTimeframeWithDefaults instantiates a new Timeframe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeframeWithDefaults() *Timeframe {
	this := Timeframe{}
	return &this
}

// GetKey returns the Key field value
func (o *Timeframe) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Timeframe) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Timeframe) SetKey(v string) {
	o.Key = v
}

func (o Timeframe) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Timeframe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

func (o *Timeframe) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
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

	varTimeframe := _Timeframe{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimeframe)

	if err != nil {
		return err
	}

	*o = Timeframe(varTimeframe)

	return err
}

type NullableTimeframe struct {
	value *Timeframe
	isSet bool
}

func (v NullableTimeframe) Get() *Timeframe {
	return v.value
}

func (v *NullableTimeframe) Set(val *Timeframe) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeframe) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeframe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeframe(val *Timeframe) *NullableTimeframe {
	return &NullableTimeframe{value: val, isSet: true}
}

func (v NullableTimeframe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeframe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


