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

// checks if the EventsBulkCreateQueryResourceObjectAttributesEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsBulkCreateQueryResourceObjectAttributesEvents{}

// EventsBulkCreateQueryResourceObjectAttributesEvents struct for EventsBulkCreateQueryResourceObjectAttributesEvents
type EventsBulkCreateQueryResourceObjectAttributesEvents struct {
	Data []BaseEventCreateQueryResourceObject `json:"data"`
}

type _EventsBulkCreateQueryResourceObjectAttributesEvents EventsBulkCreateQueryResourceObjectAttributesEvents

// NewEventsBulkCreateQueryResourceObjectAttributesEvents instantiates a new EventsBulkCreateQueryResourceObjectAttributesEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsBulkCreateQueryResourceObjectAttributesEvents(data []BaseEventCreateQueryResourceObject) *EventsBulkCreateQueryResourceObjectAttributesEvents {
	this := EventsBulkCreateQueryResourceObjectAttributesEvents{}
	this.Data = data
	return &this
}

// NewEventsBulkCreateQueryResourceObjectAttributesEventsWithDefaults instantiates a new EventsBulkCreateQueryResourceObjectAttributesEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsBulkCreateQueryResourceObjectAttributesEventsWithDefaults() *EventsBulkCreateQueryResourceObjectAttributesEvents {
	this := EventsBulkCreateQueryResourceObjectAttributesEvents{}
	return &this
}

// GetData returns the Data field value
func (o *EventsBulkCreateQueryResourceObjectAttributesEvents) GetData() []BaseEventCreateQueryResourceObject {
	if o == nil {
		var ret []BaseEventCreateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EventsBulkCreateQueryResourceObjectAttributesEvents) GetDataOk() ([]BaseEventCreateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *EventsBulkCreateQueryResourceObjectAttributesEvents) SetData(v []BaseEventCreateQueryResourceObject) {
	o.Data = v
}

func (o EventsBulkCreateQueryResourceObjectAttributesEvents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsBulkCreateQueryResourceObjectAttributesEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *EventsBulkCreateQueryResourceObjectAttributesEvents) UnmarshalJSON(data []byte) (err error) {
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

	varEventsBulkCreateQueryResourceObjectAttributesEvents := _EventsBulkCreateQueryResourceObjectAttributesEvents{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventsBulkCreateQueryResourceObjectAttributesEvents)

	if err != nil {
		return err
	}

	*o = EventsBulkCreateQueryResourceObjectAttributesEvents(varEventsBulkCreateQueryResourceObjectAttributesEvents)

	return err
}

type NullableEventsBulkCreateQueryResourceObjectAttributesEvents struct {
	value *EventsBulkCreateQueryResourceObjectAttributesEvents
	isSet bool
}

func (v NullableEventsBulkCreateQueryResourceObjectAttributesEvents) Get() *EventsBulkCreateQueryResourceObjectAttributesEvents {
	return v.value
}

func (v *NullableEventsBulkCreateQueryResourceObjectAttributesEvents) Set(val *EventsBulkCreateQueryResourceObjectAttributesEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsBulkCreateQueryResourceObjectAttributesEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsBulkCreateQueryResourceObjectAttributesEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsBulkCreateQueryResourceObjectAttributesEvents(val *EventsBulkCreateQueryResourceObjectAttributesEvents) *NullableEventsBulkCreateQueryResourceObjectAttributesEvents {
	return &NullableEventsBulkCreateQueryResourceObjectAttributesEvents{value: val, isSet: true}
}

func (v NullableEventsBulkCreateQueryResourceObjectAttributesEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsBulkCreateQueryResourceObjectAttributesEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


