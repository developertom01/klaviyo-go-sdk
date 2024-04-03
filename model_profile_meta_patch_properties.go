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
)

// checks if the ProfileMetaPatchProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileMetaPatchProperties{}

// ProfileMetaPatchProperties struct for ProfileMetaPatchProperties
type ProfileMetaPatchProperties struct {
	// Append a simple value or values to this property array
	Append map[string]interface{} `json:"append,omitempty"`
	// Remove a simple value or values from this property array
	Unappend map[string]interface{} `json:"unappend,omitempty"`
	Unset NullableProfileMetaPatchPropertiesUnset `json:"unset,omitempty"`
}

// NewProfileMetaPatchProperties instantiates a new ProfileMetaPatchProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileMetaPatchProperties() *ProfileMetaPatchProperties {
	this := ProfileMetaPatchProperties{}
	return &this
}

// NewProfileMetaPatchPropertiesWithDefaults instantiates a new ProfileMetaPatchProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileMetaPatchPropertiesWithDefaults() *ProfileMetaPatchProperties {
	this := ProfileMetaPatchProperties{}
	return &this
}

// GetAppend returns the Append field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileMetaPatchProperties) GetAppend() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Append
}

// GetAppendOk returns a tuple with the Append field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileMetaPatchProperties) GetAppendOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Append) {
		return map[string]interface{}{}, false
	}
	return o.Append, true
}

// HasAppend returns a boolean if a field has been set.
func (o *ProfileMetaPatchProperties) HasAppend() bool {
	if o != nil && !IsNil(o.Append) {
		return true
	}

	return false
}

// SetAppend gets a reference to the given map[string]interface{} and assigns it to the Append field.
func (o *ProfileMetaPatchProperties) SetAppend(v map[string]interface{}) {
	o.Append = v
}

// GetUnappend returns the Unappend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileMetaPatchProperties) GetUnappend() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Unappend
}

// GetUnappendOk returns a tuple with the Unappend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileMetaPatchProperties) GetUnappendOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Unappend) {
		return map[string]interface{}{}, false
	}
	return o.Unappend, true
}

// HasUnappend returns a boolean if a field has been set.
func (o *ProfileMetaPatchProperties) HasUnappend() bool {
	if o != nil && !IsNil(o.Unappend) {
		return true
	}

	return false
}

// SetUnappend gets a reference to the given map[string]interface{} and assigns it to the Unappend field.
func (o *ProfileMetaPatchProperties) SetUnappend(v map[string]interface{}) {
	o.Unappend = v
}

// GetUnset returns the Unset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileMetaPatchProperties) GetUnset() ProfileMetaPatchPropertiesUnset {
	if o == nil || IsNil(o.Unset.Get()) {
		var ret ProfileMetaPatchPropertiesUnset
		return ret
	}
	return *o.Unset.Get()
}

// GetUnsetOk returns a tuple with the Unset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileMetaPatchProperties) GetUnsetOk() (*ProfileMetaPatchPropertiesUnset, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unset.Get(), o.Unset.IsSet()
}

// HasUnset returns a boolean if a field has been set.
func (o *ProfileMetaPatchProperties) HasUnset() bool {
	if o != nil && o.Unset.IsSet() {
		return true
	}

	return false
}

// SetUnset gets a reference to the given NullableProfileMetaPatchPropertiesUnset and assigns it to the Unset field.
func (o *ProfileMetaPatchProperties) SetUnset(v ProfileMetaPatchPropertiesUnset) {
	o.Unset.Set(&v)
}
// SetUnsetNil sets the value for Unset to be an explicit nil
func (o *ProfileMetaPatchProperties) SetUnsetNil() {
	o.Unset.Set(nil)
}

// UnsetUnset ensures that no value is present for Unset, not even an explicit nil
func (o *ProfileMetaPatchProperties) UnsetUnset() {
	o.Unset.Unset()
}

func (o ProfileMetaPatchProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileMetaPatchProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Append != nil {
		toSerialize["append"] = o.Append
	}
	if o.Unappend != nil {
		toSerialize["unappend"] = o.Unappend
	}
	if o.Unset.IsSet() {
		toSerialize["unset"] = o.Unset.Get()
	}
	return toSerialize, nil
}

type NullableProfileMetaPatchProperties struct {
	value *ProfileMetaPatchProperties
	isSet bool
}

func (v NullableProfileMetaPatchProperties) Get() *ProfileMetaPatchProperties {
	return v.value
}

func (v *NullableProfileMetaPatchProperties) Set(val *ProfileMetaPatchProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileMetaPatchProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileMetaPatchProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileMetaPatchProperties(val *ProfileMetaPatchProperties) *NullableProfileMetaPatchProperties {
	return &NullableProfileMetaPatchProperties{value: val, isSet: true}
}

func (v NullableProfileMetaPatchProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileMetaPatchProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


