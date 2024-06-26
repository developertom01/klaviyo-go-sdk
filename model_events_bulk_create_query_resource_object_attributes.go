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

// checks if the EventsBulkCreateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsBulkCreateQueryResourceObjectAttributes{}

// EventsBulkCreateQueryResourceObjectAttributes struct for EventsBulkCreateQueryResourceObjectAttributes
type EventsBulkCreateQueryResourceObjectAttributes struct {
	Profile EventCreateQueryV2ResourceObjectAttributesProfile `json:"profile"`
	Events EventsBulkCreateQueryResourceObjectAttributesEvents `json:"events"`
}

type _EventsBulkCreateQueryResourceObjectAttributes EventsBulkCreateQueryResourceObjectAttributes

// NewEventsBulkCreateQueryResourceObjectAttributes instantiates a new EventsBulkCreateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsBulkCreateQueryResourceObjectAttributes(profile EventCreateQueryV2ResourceObjectAttributesProfile, events EventsBulkCreateQueryResourceObjectAttributesEvents) *EventsBulkCreateQueryResourceObjectAttributes {
	this := EventsBulkCreateQueryResourceObjectAttributes{}
	this.Profile = profile
	this.Events = events
	return &this
}

// NewEventsBulkCreateQueryResourceObjectAttributesWithDefaults instantiates a new EventsBulkCreateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsBulkCreateQueryResourceObjectAttributesWithDefaults() *EventsBulkCreateQueryResourceObjectAttributes {
	this := EventsBulkCreateQueryResourceObjectAttributes{}
	return &this
}

// GetProfile returns the Profile field value
func (o *EventsBulkCreateQueryResourceObjectAttributes) GetProfile() EventCreateQueryV2ResourceObjectAttributesProfile {
	if o == nil {
		var ret EventCreateQueryV2ResourceObjectAttributesProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *EventsBulkCreateQueryResourceObjectAttributes) GetProfileOk() (*EventCreateQueryV2ResourceObjectAttributesProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *EventsBulkCreateQueryResourceObjectAttributes) SetProfile(v EventCreateQueryV2ResourceObjectAttributesProfile) {
	o.Profile = v
}

// GetEvents returns the Events field value
func (o *EventsBulkCreateQueryResourceObjectAttributes) GetEvents() EventsBulkCreateQueryResourceObjectAttributesEvents {
	if o == nil {
		var ret EventsBulkCreateQueryResourceObjectAttributesEvents
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *EventsBulkCreateQueryResourceObjectAttributes) GetEventsOk() (*EventsBulkCreateQueryResourceObjectAttributesEvents, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Events, true
}

// SetEvents sets field value
func (o *EventsBulkCreateQueryResourceObjectAttributes) SetEvents(v EventsBulkCreateQueryResourceObjectAttributesEvents) {
	o.Events = v
}

func (o EventsBulkCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsBulkCreateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profile"] = o.Profile
	toSerialize["events"] = o.Events
	return toSerialize, nil
}

func (o *EventsBulkCreateQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profile",
		"events",
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

	varEventsBulkCreateQueryResourceObjectAttributes := _EventsBulkCreateQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventsBulkCreateQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = EventsBulkCreateQueryResourceObjectAttributes(varEventsBulkCreateQueryResourceObjectAttributes)

	return err
}

type NullableEventsBulkCreateQueryResourceObjectAttributes struct {
	value *EventsBulkCreateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableEventsBulkCreateQueryResourceObjectAttributes) Get() *EventsBulkCreateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableEventsBulkCreateQueryResourceObjectAttributes) Set(val *EventsBulkCreateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsBulkCreateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsBulkCreateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsBulkCreateQueryResourceObjectAttributes(val *EventsBulkCreateQueryResourceObjectAttributes) *NullableEventsBulkCreateQueryResourceObjectAttributes {
	return &NullableEventsBulkCreateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableEventsBulkCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsBulkCreateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


