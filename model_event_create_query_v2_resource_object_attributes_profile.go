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

// checks if the EventCreateQueryV2ResourceObjectAttributesProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventCreateQueryV2ResourceObjectAttributesProfile{}

// EventCreateQueryV2ResourceObjectAttributesProfile struct for EventCreateQueryV2ResourceObjectAttributesProfile
type EventCreateQueryV2ResourceObjectAttributesProfile struct {
	Data OnsiteProfileCreateQueryResourceObject `json:"data"`
}

type _EventCreateQueryV2ResourceObjectAttributesProfile EventCreateQueryV2ResourceObjectAttributesProfile

// NewEventCreateQueryV2ResourceObjectAttributesProfile instantiates a new EventCreateQueryV2ResourceObjectAttributesProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventCreateQueryV2ResourceObjectAttributesProfile(data OnsiteProfileCreateQueryResourceObject) *EventCreateQueryV2ResourceObjectAttributesProfile {
	this := EventCreateQueryV2ResourceObjectAttributesProfile{}
	this.Data = data
	return &this
}

// NewEventCreateQueryV2ResourceObjectAttributesProfileWithDefaults instantiates a new EventCreateQueryV2ResourceObjectAttributesProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventCreateQueryV2ResourceObjectAttributesProfileWithDefaults() *EventCreateQueryV2ResourceObjectAttributesProfile {
	this := EventCreateQueryV2ResourceObjectAttributesProfile{}
	return &this
}

// GetData returns the Data field value
func (o *EventCreateQueryV2ResourceObjectAttributesProfile) GetData() OnsiteProfileCreateQueryResourceObject {
	if o == nil {
		var ret OnsiteProfileCreateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EventCreateQueryV2ResourceObjectAttributesProfile) GetDataOk() (*OnsiteProfileCreateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *EventCreateQueryV2ResourceObjectAttributesProfile) SetData(v OnsiteProfileCreateQueryResourceObject) {
	o.Data = v
}

func (o EventCreateQueryV2ResourceObjectAttributesProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventCreateQueryV2ResourceObjectAttributesProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *EventCreateQueryV2ResourceObjectAttributesProfile) UnmarshalJSON(data []byte) (err error) {
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

	varEventCreateQueryV2ResourceObjectAttributesProfile := _EventCreateQueryV2ResourceObjectAttributesProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventCreateQueryV2ResourceObjectAttributesProfile)

	if err != nil {
		return err
	}

	*o = EventCreateQueryV2ResourceObjectAttributesProfile(varEventCreateQueryV2ResourceObjectAttributesProfile)

	return err
}

type NullableEventCreateQueryV2ResourceObjectAttributesProfile struct {
	value *EventCreateQueryV2ResourceObjectAttributesProfile
	isSet bool
}

func (v NullableEventCreateQueryV2ResourceObjectAttributesProfile) Get() *EventCreateQueryV2ResourceObjectAttributesProfile {
	return v.value
}

func (v *NullableEventCreateQueryV2ResourceObjectAttributesProfile) Set(val *EventCreateQueryV2ResourceObjectAttributesProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCreateQueryV2ResourceObjectAttributesProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCreateQueryV2ResourceObjectAttributesProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCreateQueryV2ResourceObjectAttributesProfile(val *EventCreateQueryV2ResourceObjectAttributesProfile) *NullableEventCreateQueryV2ResourceObjectAttributesProfile {
	return &NullableEventCreateQueryV2ResourceObjectAttributesProfile{value: val, isSet: true}
}

func (v NullableEventCreateQueryV2ResourceObjectAttributesProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCreateQueryV2ResourceObjectAttributesProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

