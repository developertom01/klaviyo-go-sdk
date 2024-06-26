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

// checks if the SendStrategySubObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendStrategySubObject{}

// SendStrategySubObject struct for SendStrategySubObject
type SendStrategySubObject struct {
	// Describes the shape of the options object. Allowed values: ['static', 'throttled', 'immediate', 'smart_send_time']
	Method string `json:"method"`
	OptionsStatic *StaticScheduleOptions `json:"options_static,omitempty"`
	OptionsThrottled *ThrottledScheduleOptions `json:"options_throttled,omitempty"`
	OptionsSto *STOScheduleOptions `json:"options_sto,omitempty"`
}

type _SendStrategySubObject SendStrategySubObject

// NewSendStrategySubObject instantiates a new SendStrategySubObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendStrategySubObject(method string) *SendStrategySubObject {
	this := SendStrategySubObject{}
	this.Method = method
	return &this
}

// NewSendStrategySubObjectWithDefaults instantiates a new SendStrategySubObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendStrategySubObjectWithDefaults() *SendStrategySubObject {
	this := SendStrategySubObject{}
	return &this
}

// GetMethod returns the Method field value
func (o *SendStrategySubObject) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SendStrategySubObject) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SendStrategySubObject) SetMethod(v string) {
	o.Method = v
}

// GetOptionsStatic returns the OptionsStatic field value if set, zero value otherwise.
func (o *SendStrategySubObject) GetOptionsStatic() StaticScheduleOptions {
	if o == nil || IsNil(o.OptionsStatic) {
		var ret StaticScheduleOptions
		return ret
	}
	return *o.OptionsStatic
}

// GetOptionsStaticOk returns a tuple with the OptionsStatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendStrategySubObject) GetOptionsStaticOk() (*StaticScheduleOptions, bool) {
	if o == nil || IsNil(o.OptionsStatic) {
		return nil, false
	}
	return o.OptionsStatic, true
}

// HasOptionsStatic returns a boolean if a field has been set.
func (o *SendStrategySubObject) HasOptionsStatic() bool {
	if o != nil && !IsNil(o.OptionsStatic) {
		return true
	}

	return false
}

// SetOptionsStatic gets a reference to the given StaticScheduleOptions and assigns it to the OptionsStatic field.
func (o *SendStrategySubObject) SetOptionsStatic(v StaticScheduleOptions) {
	o.OptionsStatic = &v
}

// GetOptionsThrottled returns the OptionsThrottled field value if set, zero value otherwise.
func (o *SendStrategySubObject) GetOptionsThrottled() ThrottledScheduleOptions {
	if o == nil || IsNil(o.OptionsThrottled) {
		var ret ThrottledScheduleOptions
		return ret
	}
	return *o.OptionsThrottled
}

// GetOptionsThrottledOk returns a tuple with the OptionsThrottled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendStrategySubObject) GetOptionsThrottledOk() (*ThrottledScheduleOptions, bool) {
	if o == nil || IsNil(o.OptionsThrottled) {
		return nil, false
	}
	return o.OptionsThrottled, true
}

// HasOptionsThrottled returns a boolean if a field has been set.
func (o *SendStrategySubObject) HasOptionsThrottled() bool {
	if o != nil && !IsNil(o.OptionsThrottled) {
		return true
	}

	return false
}

// SetOptionsThrottled gets a reference to the given ThrottledScheduleOptions and assigns it to the OptionsThrottled field.
func (o *SendStrategySubObject) SetOptionsThrottled(v ThrottledScheduleOptions) {
	o.OptionsThrottled = &v
}

// GetOptionsSto returns the OptionsSto field value if set, zero value otherwise.
func (o *SendStrategySubObject) GetOptionsSto() STOScheduleOptions {
	if o == nil || IsNil(o.OptionsSto) {
		var ret STOScheduleOptions
		return ret
	}
	return *o.OptionsSto
}

// GetOptionsStoOk returns a tuple with the OptionsSto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendStrategySubObject) GetOptionsStoOk() (*STOScheduleOptions, bool) {
	if o == nil || IsNil(o.OptionsSto) {
		return nil, false
	}
	return o.OptionsSto, true
}

// HasOptionsSto returns a boolean if a field has been set.
func (o *SendStrategySubObject) HasOptionsSto() bool {
	if o != nil && !IsNil(o.OptionsSto) {
		return true
	}

	return false
}

// SetOptionsSto gets a reference to the given STOScheduleOptions and assigns it to the OptionsSto field.
func (o *SendStrategySubObject) SetOptionsSto(v STOScheduleOptions) {
	o.OptionsSto = &v
}

func (o SendStrategySubObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendStrategySubObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["method"] = o.Method
	if !IsNil(o.OptionsStatic) {
		toSerialize["options_static"] = o.OptionsStatic
	}
	if !IsNil(o.OptionsThrottled) {
		toSerialize["options_throttled"] = o.OptionsThrottled
	}
	if !IsNil(o.OptionsSto) {
		toSerialize["options_sto"] = o.OptionsSto
	}
	return toSerialize, nil
}

func (o *SendStrategySubObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
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

	varSendStrategySubObject := _SendStrategySubObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSendStrategySubObject)

	if err != nil {
		return err
	}

	*o = SendStrategySubObject(varSendStrategySubObject)

	return err
}

type NullableSendStrategySubObject struct {
	value *SendStrategySubObject
	isSet bool
}

func (v NullableSendStrategySubObject) Get() *SendStrategySubObject {
	return v.value
}

func (v *NullableSendStrategySubObject) Set(val *SendStrategySubObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSendStrategySubObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSendStrategySubObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendStrategySubObject(val *SendStrategySubObject) *NullableSendStrategySubObject {
	return &NullableSendStrategySubObject{value: val, isSet: true}
}

func (v NullableSendStrategySubObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendStrategySubObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


