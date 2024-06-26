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

// checks if the AttributionResponseObjectResourceRelationshipsEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributionResponseObjectResourceRelationshipsEvent{}

// AttributionResponseObjectResourceRelationshipsEvent struct for AttributionResponseObjectResourceRelationshipsEvent
type AttributionResponseObjectResourceRelationshipsEvent struct {
	Data AttributionResponseObjectResourceRelationshipsEventData `json:"data"`
}

type _AttributionResponseObjectResourceRelationshipsEvent AttributionResponseObjectResourceRelationshipsEvent

// NewAttributionResponseObjectResourceRelationshipsEvent instantiates a new AttributionResponseObjectResourceRelationshipsEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributionResponseObjectResourceRelationshipsEvent(data AttributionResponseObjectResourceRelationshipsEventData) *AttributionResponseObjectResourceRelationshipsEvent {
	this := AttributionResponseObjectResourceRelationshipsEvent{}
	this.Data = data
	return &this
}

// NewAttributionResponseObjectResourceRelationshipsEventWithDefaults instantiates a new AttributionResponseObjectResourceRelationshipsEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributionResponseObjectResourceRelationshipsEventWithDefaults() *AttributionResponseObjectResourceRelationshipsEvent {
	this := AttributionResponseObjectResourceRelationshipsEvent{}
	return &this
}

// GetData returns the Data field value
func (o *AttributionResponseObjectResourceRelationshipsEvent) GetData() AttributionResponseObjectResourceRelationshipsEventData {
	if o == nil {
		var ret AttributionResponseObjectResourceRelationshipsEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AttributionResponseObjectResourceRelationshipsEvent) GetDataOk() (*AttributionResponseObjectResourceRelationshipsEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AttributionResponseObjectResourceRelationshipsEvent) SetData(v AttributionResponseObjectResourceRelationshipsEventData) {
	o.Data = v
}

func (o AttributionResponseObjectResourceRelationshipsEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributionResponseObjectResourceRelationshipsEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AttributionResponseObjectResourceRelationshipsEvent) UnmarshalJSON(data []byte) (err error) {
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

	varAttributionResponseObjectResourceRelationshipsEvent := _AttributionResponseObjectResourceRelationshipsEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAttributionResponseObjectResourceRelationshipsEvent)

	if err != nil {
		return err
	}

	*o = AttributionResponseObjectResourceRelationshipsEvent(varAttributionResponseObjectResourceRelationshipsEvent)

	return err
}

type NullableAttributionResponseObjectResourceRelationshipsEvent struct {
	value *AttributionResponseObjectResourceRelationshipsEvent
	isSet bool
}

func (v NullableAttributionResponseObjectResourceRelationshipsEvent) Get() *AttributionResponseObjectResourceRelationshipsEvent {
	return v.value
}

func (v *NullableAttributionResponseObjectResourceRelationshipsEvent) Set(val *AttributionResponseObjectResourceRelationshipsEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributionResponseObjectResourceRelationshipsEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributionResponseObjectResourceRelationshipsEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributionResponseObjectResourceRelationshipsEvent(val *AttributionResponseObjectResourceRelationshipsEvent) *NullableAttributionResponseObjectResourceRelationshipsEvent {
	return &NullableAttributionResponseObjectResourceRelationshipsEvent{value: val, isSet: true}
}

func (v NullableAttributionResponseObjectResourceRelationshipsEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributionResponseObjectResourceRelationshipsEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


