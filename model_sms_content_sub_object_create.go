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

// checks if the SMSContentSubObjectCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SMSContentSubObjectCreate{}

// SMSContentSubObjectCreate struct for SMSContentSubObjectCreate
type SMSContentSubObjectCreate struct {
	// The message body
	Body NullableString `json:"body,omitempty"`
}

// NewSMSContentSubObjectCreate instantiates a new SMSContentSubObjectCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMSContentSubObjectCreate() *SMSContentSubObjectCreate {
	this := SMSContentSubObjectCreate{}
	return &this
}

// NewSMSContentSubObjectCreateWithDefaults instantiates a new SMSContentSubObjectCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMSContentSubObjectCreateWithDefaults() *SMSContentSubObjectCreate {
	this := SMSContentSubObjectCreate{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SMSContentSubObjectCreate) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SMSContentSubObjectCreate) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *SMSContentSubObjectCreate) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *SMSContentSubObjectCreate) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *SMSContentSubObjectCreate) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *SMSContentSubObjectCreate) UnsetBody() {
	o.Body.Unset()
}

func (o SMSContentSubObjectCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SMSContentSubObjectCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	return toSerialize, nil
}

type NullableSMSContentSubObjectCreate struct {
	value *SMSContentSubObjectCreate
	isSet bool
}

func (v NullableSMSContentSubObjectCreate) Get() *SMSContentSubObjectCreate {
	return v.value
}

func (v *NullableSMSContentSubObjectCreate) Set(val *SMSContentSubObjectCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSMSContentSubObjectCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSMSContentSubObjectCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMSContentSubObjectCreate(val *SMSContentSubObjectCreate) *NullableSMSContentSubObjectCreate {
	return &NullableSMSContentSubObjectCreate{value: val, isSet: true}
}

func (v NullableSMSContentSubObjectCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMSContentSubObjectCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


