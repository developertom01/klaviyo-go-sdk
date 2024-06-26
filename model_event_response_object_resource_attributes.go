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
	"time"
)

// checks if the EventResponseObjectResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventResponseObjectResourceAttributes{}

// EventResponseObjectResourceAttributes struct for EventResponseObjectResourceAttributes
type EventResponseObjectResourceAttributes struct {
	// Event timestamp in seconds
	Timestamp NullableInt32 `json:"timestamp,omitempty"`
	// Event properties, can include identifiers and extra properties
	EventProperties map[string]interface{} `json:"event_properties,omitempty"`
	// Event timestamp in ISO8601 format (YYYY-MM-DDTHH:MM:SS+hh:mm)
	Datetime NullableTime `json:"datetime,omitempty"`
	// A unique identifier for the event, this can be used as a cursor in pagination
	Uuid NullableString `json:"uuid,omitempty"`
}

// NewEventResponseObjectResourceAttributes instantiates a new EventResponseObjectResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResponseObjectResourceAttributes() *EventResponseObjectResourceAttributes {
	this := EventResponseObjectResourceAttributes{}
	return &this
}

// NewEventResponseObjectResourceAttributesWithDefaults instantiates a new EventResponseObjectResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResponseObjectResourceAttributesWithDefaults() *EventResponseObjectResourceAttributes {
	this := EventResponseObjectResourceAttributes{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseObjectResourceAttributes) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret int32
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseObjectResourceAttributes) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *EventResponseObjectResourceAttributes) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableInt32 and assigns it to the Timestamp field.
func (o *EventResponseObjectResourceAttributes) SetTimestamp(v int32) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *EventResponseObjectResourceAttributes) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *EventResponseObjectResourceAttributes) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetEventProperties returns the EventProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseObjectResourceAttributes) GetEventProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.EventProperties
}

// GetEventPropertiesOk returns a tuple with the EventProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseObjectResourceAttributes) GetEventPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.EventProperties) {
		return map[string]interface{}{}, false
	}
	return o.EventProperties, true
}

// HasEventProperties returns a boolean if a field has been set.
func (o *EventResponseObjectResourceAttributes) HasEventProperties() bool {
	if o != nil && !IsNil(o.EventProperties) {
		return true
	}

	return false
}

// SetEventProperties gets a reference to the given map[string]interface{} and assigns it to the EventProperties field.
func (o *EventResponseObjectResourceAttributes) SetEventProperties(v map[string]interface{}) {
	o.EventProperties = v
}

// GetDatetime returns the Datetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseObjectResourceAttributes) GetDatetime() time.Time {
	if o == nil || IsNil(o.Datetime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Datetime.Get()
}

// GetDatetimeOk returns a tuple with the Datetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseObjectResourceAttributes) GetDatetimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Datetime.Get(), o.Datetime.IsSet()
}

// HasDatetime returns a boolean if a field has been set.
func (o *EventResponseObjectResourceAttributes) HasDatetime() bool {
	if o != nil && o.Datetime.IsSet() {
		return true
	}

	return false
}

// SetDatetime gets a reference to the given NullableTime and assigns it to the Datetime field.
func (o *EventResponseObjectResourceAttributes) SetDatetime(v time.Time) {
	o.Datetime.Set(&v)
}
// SetDatetimeNil sets the value for Datetime to be an explicit nil
func (o *EventResponseObjectResourceAttributes) SetDatetimeNil() {
	o.Datetime.Set(nil)
}

// UnsetDatetime ensures that no value is present for Datetime, not even an explicit nil
func (o *EventResponseObjectResourceAttributes) UnsetDatetime() {
	o.Datetime.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseObjectResourceAttributes) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseObjectResourceAttributes) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *EventResponseObjectResourceAttributes) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *EventResponseObjectResourceAttributes) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *EventResponseObjectResourceAttributes) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *EventResponseObjectResourceAttributes) UnsetUuid() {
	o.Uuid.Unset()
}

func (o EventResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventResponseObjectResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.EventProperties != nil {
		toSerialize["event_properties"] = o.EventProperties
	}
	if o.Datetime.IsSet() {
		toSerialize["datetime"] = o.Datetime.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	return toSerialize, nil
}

type NullableEventResponseObjectResourceAttributes struct {
	value *EventResponseObjectResourceAttributes
	isSet bool
}

func (v NullableEventResponseObjectResourceAttributes) Get() *EventResponseObjectResourceAttributes {
	return v.value
}

func (v *NullableEventResponseObjectResourceAttributes) Set(val *EventResponseObjectResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResponseObjectResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResponseObjectResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResponseObjectResourceAttributes(val *EventResponseObjectResourceAttributes) *NullableEventResponseObjectResourceAttributes {
	return &NullableEventResponseObjectResourceAttributes{value: val, isSet: true}
}

func (v NullableEventResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResponseObjectResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


