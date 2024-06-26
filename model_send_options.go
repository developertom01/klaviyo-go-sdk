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

// checks if the SendOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendOptions{}

// SendOptions struct for SendOptions
type SendOptions struct {
	UseSmartSending bool `json:"use_smart_sending"`
	IsTransactional bool `json:"is_transactional"`
}

type _SendOptions SendOptions

// NewSendOptions instantiates a new SendOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendOptions(useSmartSending bool, isTransactional bool) *SendOptions {
	this := SendOptions{}
	this.UseSmartSending = useSmartSending
	this.IsTransactional = isTransactional
	return &this
}

// NewSendOptionsWithDefaults instantiates a new SendOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendOptionsWithDefaults() *SendOptions {
	this := SendOptions{}
	return &this
}

// GetUseSmartSending returns the UseSmartSending field value
func (o *SendOptions) GetUseSmartSending() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseSmartSending
}

// GetUseSmartSendingOk returns a tuple with the UseSmartSending field value
// and a boolean to check if the value has been set.
func (o *SendOptions) GetUseSmartSendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseSmartSending, true
}

// SetUseSmartSending sets field value
func (o *SendOptions) SetUseSmartSending(v bool) {
	o.UseSmartSending = v
}

// GetIsTransactional returns the IsTransactional field value
func (o *SendOptions) GetIsTransactional() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTransactional
}

// GetIsTransactionalOk returns a tuple with the IsTransactional field value
// and a boolean to check if the value has been set.
func (o *SendOptions) GetIsTransactionalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTransactional, true
}

// SetIsTransactional sets field value
func (o *SendOptions) SetIsTransactional(v bool) {
	o.IsTransactional = v
}

func (o SendOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["use_smart_sending"] = o.UseSmartSending
	toSerialize["is_transactional"] = o.IsTransactional
	return toSerialize, nil
}

func (o *SendOptions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"use_smart_sending",
		"is_transactional",
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

	varSendOptions := _SendOptions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSendOptions)

	if err != nil {
		return err
	}

	*o = SendOptions(varSendOptions)

	return err
}

type NullableSendOptions struct {
	value *SendOptions
	isSet bool
}

func (v NullableSendOptions) Get() *SendOptions {
	return v.value
}

func (v *NullableSendOptions) Set(val *SendOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSendOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSendOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendOptions(val *SendOptions) *NullableSendOptions {
	return &NullableSendOptions{value: val, isSet: true}
}

func (v NullableSendOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


