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

// checks if the AttributionResponseObjectResourceRelationshipsAttributedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributionResponseObjectResourceRelationshipsAttributedEvent{}

// AttributionResponseObjectResourceRelationshipsAttributedEvent struct for AttributionResponseObjectResourceRelationshipsAttributedEvent
type AttributionResponseObjectResourceRelationshipsAttributedEvent struct {
	Data AttributionResponseObjectResourceRelationshipsAttributedEventData `json:"data"`
}

type _AttributionResponseObjectResourceRelationshipsAttributedEvent AttributionResponseObjectResourceRelationshipsAttributedEvent

// NewAttributionResponseObjectResourceRelationshipsAttributedEvent instantiates a new AttributionResponseObjectResourceRelationshipsAttributedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributionResponseObjectResourceRelationshipsAttributedEvent(data AttributionResponseObjectResourceRelationshipsAttributedEventData) *AttributionResponseObjectResourceRelationshipsAttributedEvent {
	this := AttributionResponseObjectResourceRelationshipsAttributedEvent{}
	this.Data = data
	return &this
}

// NewAttributionResponseObjectResourceRelationshipsAttributedEventWithDefaults instantiates a new AttributionResponseObjectResourceRelationshipsAttributedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributionResponseObjectResourceRelationshipsAttributedEventWithDefaults() *AttributionResponseObjectResourceRelationshipsAttributedEvent {
	this := AttributionResponseObjectResourceRelationshipsAttributedEvent{}
	return &this
}

// GetData returns the Data field value
func (o *AttributionResponseObjectResourceRelationshipsAttributedEvent) GetData() AttributionResponseObjectResourceRelationshipsAttributedEventData {
	if o == nil {
		var ret AttributionResponseObjectResourceRelationshipsAttributedEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AttributionResponseObjectResourceRelationshipsAttributedEvent) GetDataOk() (*AttributionResponseObjectResourceRelationshipsAttributedEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AttributionResponseObjectResourceRelationshipsAttributedEvent) SetData(v AttributionResponseObjectResourceRelationshipsAttributedEventData) {
	o.Data = v
}

func (o AttributionResponseObjectResourceRelationshipsAttributedEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributionResponseObjectResourceRelationshipsAttributedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AttributionResponseObjectResourceRelationshipsAttributedEvent) UnmarshalJSON(data []byte) (err error) {
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

	varAttributionResponseObjectResourceRelationshipsAttributedEvent := _AttributionResponseObjectResourceRelationshipsAttributedEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAttributionResponseObjectResourceRelationshipsAttributedEvent)

	if err != nil {
		return err
	}

	*o = AttributionResponseObjectResourceRelationshipsAttributedEvent(varAttributionResponseObjectResourceRelationshipsAttributedEvent)

	return err
}

type NullableAttributionResponseObjectResourceRelationshipsAttributedEvent struct {
	value *AttributionResponseObjectResourceRelationshipsAttributedEvent
	isSet bool
}

func (v NullableAttributionResponseObjectResourceRelationshipsAttributedEvent) Get() *AttributionResponseObjectResourceRelationshipsAttributedEvent {
	return v.value
}

func (v *NullableAttributionResponseObjectResourceRelationshipsAttributedEvent) Set(val *AttributionResponseObjectResourceRelationshipsAttributedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributionResponseObjectResourceRelationshipsAttributedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributionResponseObjectResourceRelationshipsAttributedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributionResponseObjectResourceRelationshipsAttributedEvent(val *AttributionResponseObjectResourceRelationshipsAttributedEvent) *NullableAttributionResponseObjectResourceRelationshipsAttributedEvent {
	return &NullableAttributionResponseObjectResourceRelationshipsAttributedEvent{value: val, isSet: true}
}

func (v NullableAttributionResponseObjectResourceRelationshipsAttributedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributionResponseObjectResourceRelationshipsAttributedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


