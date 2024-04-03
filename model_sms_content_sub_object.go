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

// checks if the SMSContentSubObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SMSContentSubObject{}

// SMSContentSubObject struct for SMSContentSubObject
type SMSContentSubObject struct {
	// The message body
	Body NullableString `json:"body,omitempty"`
	// URL for included media
	MediaUrl NullableString `json:"media_url,omitempty"`
}

// NewSMSContentSubObject instantiates a new SMSContentSubObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMSContentSubObject() *SMSContentSubObject {
	this := SMSContentSubObject{}
	return &this
}

// NewSMSContentSubObjectWithDefaults instantiates a new SMSContentSubObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMSContentSubObjectWithDefaults() *SMSContentSubObject {
	this := SMSContentSubObject{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SMSContentSubObject) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SMSContentSubObject) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *SMSContentSubObject) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *SMSContentSubObject) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *SMSContentSubObject) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *SMSContentSubObject) UnsetBody() {
	o.Body.Unset()
}

// GetMediaUrl returns the MediaUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SMSContentSubObject) GetMediaUrl() string {
	if o == nil || IsNil(o.MediaUrl.Get()) {
		var ret string
		return ret
	}
	return *o.MediaUrl.Get()
}

// GetMediaUrlOk returns a tuple with the MediaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SMSContentSubObject) GetMediaUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaUrl.Get(), o.MediaUrl.IsSet()
}

// HasMediaUrl returns a boolean if a field has been set.
func (o *SMSContentSubObject) HasMediaUrl() bool {
	if o != nil && o.MediaUrl.IsSet() {
		return true
	}

	return false
}

// SetMediaUrl gets a reference to the given NullableString and assigns it to the MediaUrl field.
func (o *SMSContentSubObject) SetMediaUrl(v string) {
	o.MediaUrl.Set(&v)
}
// SetMediaUrlNil sets the value for MediaUrl to be an explicit nil
func (o *SMSContentSubObject) SetMediaUrlNil() {
	o.MediaUrl.Set(nil)
}

// UnsetMediaUrl ensures that no value is present for MediaUrl, not even an explicit nil
func (o *SMSContentSubObject) UnsetMediaUrl() {
	o.MediaUrl.Unset()
}

func (o SMSContentSubObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SMSContentSubObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	if o.MediaUrl.IsSet() {
		toSerialize["media_url"] = o.MediaUrl.Get()
	}
	return toSerialize, nil
}

type NullableSMSContentSubObject struct {
	value *SMSContentSubObject
	isSet bool
}

func (v NullableSMSContentSubObject) Get() *SMSContentSubObject {
	return v.value
}

func (v *NullableSMSContentSubObject) Set(val *SMSContentSubObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSMSContentSubObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSMSContentSubObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMSContentSubObject(val *SMSContentSubObject) *NullableSMSContentSubObject {
	return &NullableSMSContentSubObject{value: val, isSet: true}
}

func (v NullableSMSContentSubObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMSContentSubObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

