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
	"bytes"
	"fmt"
)

// checks if the EventCreateQueryV2ResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventCreateQueryV2ResourceObjectAttributes{}

// EventCreateQueryV2ResourceObjectAttributes struct for EventCreateQueryV2ResourceObjectAttributes
type EventCreateQueryV2ResourceObjectAttributes struct {
	// Properties of this event. Any top level property (that are not objects) can be used to create segments. The $extra property is a special property. This records any non-segmentable values that can be referenced later. For example, HTML templates are useful on a segment but are not used to create a segment. There are limits placed onto the size of the data present. This must not exceed 5 MB. This must not exceed 300 event properties. A single string cannot be larger than 100 KB. Each array must not exceed 4000 elements. The properties cannot contain more than 10 nested levels.
	Properties map[string]interface{} `json:"properties"`
	// When this event occurred. By default, the time the request was received will be used. The time is truncated to the second. The time must be after the year 2000 and can only be up to 1 year in the future.
	Time NullableTime `json:"time,omitempty"`
	// A numeric, monetary value to associate with this event. For example, the dollar amount of a purchase.
	Value NullableFloat32 `json:"value,omitempty"`
	// The ISO 4217 currency code of the value associated with the event.
	ValueCurrency NullableString `json:"value_currency,omitempty"`
	// A unique identifier for an event. If the unique_id is repeated for the same profile and metric, only the first processed event will be recorded. If this is not present, this will use the time to the second. Using the default, this limits only one event per profile per second.
	UniqueId NullableString `json:"unique_id,omitempty"`
	Metric EventCreateQueryV2ResourceObjectAttributesMetric `json:"metric"`
	Profile EventCreateQueryV2ResourceObjectAttributesProfile `json:"profile"`
}

type _EventCreateQueryV2ResourceObjectAttributes EventCreateQueryV2ResourceObjectAttributes

// NewEventCreateQueryV2ResourceObjectAttributes instantiates a new EventCreateQueryV2ResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventCreateQueryV2ResourceObjectAttributes(properties map[string]interface{}, metric EventCreateQueryV2ResourceObjectAttributesMetric, profile EventCreateQueryV2ResourceObjectAttributesProfile) *EventCreateQueryV2ResourceObjectAttributes {
	this := EventCreateQueryV2ResourceObjectAttributes{}
	this.Properties = properties
	this.Metric = metric
	this.Profile = profile
	return &this
}

// NewEventCreateQueryV2ResourceObjectAttributesWithDefaults instantiates a new EventCreateQueryV2ResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventCreateQueryV2ResourceObjectAttributesWithDefaults() *EventCreateQueryV2ResourceObjectAttributes {
	this := EventCreateQueryV2ResourceObjectAttributes{}
	return &this
}

// GetProperties returns the Properties field value
func (o *EventCreateQueryV2ResourceObjectAttributes) GetProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *EventCreateQueryV2ResourceObjectAttributes) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *EventCreateQueryV2ResourceObjectAttributes) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetTime returns the Time field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventCreateQueryV2ResourceObjectAttributes) GetTime() time.Time {
	if o == nil || IsNil(o.Time.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Time.Get()
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventCreateQueryV2ResourceObjectAttributes) GetTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Time.Get(), o.Time.IsSet()
}

// HasTime returns a boolean if a field has been set.
func (o *EventCreateQueryV2ResourceObjectAttributes) HasTime() bool {
	if o != nil && o.Time.IsSet() {
		return true
	}

	return false
}

// SetTime gets a reference to the given NullableTime and assigns it to the Time field.
func (o *EventCreateQueryV2ResourceObjectAttributes) SetTime(v time.Time) {
	o.Time.Set(&v)
}
// SetTimeNil sets the value for Time to be an explicit nil
func (o *EventCreateQueryV2ResourceObjectAttributes) SetTimeNil() {
	o.Time.Set(nil)
}

// UnsetTime ensures that no value is present for Time, not even an explicit nil
func (o *EventCreateQueryV2ResourceObjectAttributes) UnsetTime() {
	o.Time.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventCreateQueryV2ResourceObjectAttributes) GetValue() float32 {
	if o == nil || IsNil(o.Value.Get()) {
		var ret float32
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventCreateQueryV2ResourceObjectAttributes) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *EventCreateQueryV2ResourceObjectAttributes) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableFloat32 and assigns it to the Value field.
func (o *EventCreateQueryV2ResourceObjectAttributes) SetValue(v float32) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *EventCreateQueryV2ResourceObjectAttributes) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *EventCreateQueryV2ResourceObjectAttributes) UnsetValue() {
	o.Value.Unset()
}

// GetValueCurrency returns the ValueCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventCreateQueryV2ResourceObjectAttributes) GetValueCurrency() string {
	if o == nil || IsNil(o.ValueCurrency.Get()) {
		var ret string
		return ret
	}
	return *o.ValueCurrency.Get()
}

// GetValueCurrencyOk returns a tuple with the ValueCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventCreateQueryV2ResourceObjectAttributes) GetValueCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueCurrency.Get(), o.ValueCurrency.IsSet()
}

// HasValueCurrency returns a boolean if a field has been set.
func (o *EventCreateQueryV2ResourceObjectAttributes) HasValueCurrency() bool {
	if o != nil && o.ValueCurrency.IsSet() {
		return true
	}

	return false
}

// SetValueCurrency gets a reference to the given NullableString and assigns it to the ValueCurrency field.
func (o *EventCreateQueryV2ResourceObjectAttributes) SetValueCurrency(v string) {
	o.ValueCurrency.Set(&v)
}
// SetValueCurrencyNil sets the value for ValueCurrency to be an explicit nil
func (o *EventCreateQueryV2ResourceObjectAttributes) SetValueCurrencyNil() {
	o.ValueCurrency.Set(nil)
}

// UnsetValueCurrency ensures that no value is present for ValueCurrency, not even an explicit nil
func (o *EventCreateQueryV2ResourceObjectAttributes) UnsetValueCurrency() {
	o.ValueCurrency.Unset()
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventCreateQueryV2ResourceObjectAttributes) GetUniqueId() string {
	if o == nil || IsNil(o.UniqueId.Get()) {
		var ret string
		return ret
	}
	return *o.UniqueId.Get()
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventCreateQueryV2ResourceObjectAttributes) GetUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UniqueId.Get(), o.UniqueId.IsSet()
}

// HasUniqueId returns a boolean if a field has been set.
func (o *EventCreateQueryV2ResourceObjectAttributes) HasUniqueId() bool {
	if o != nil && o.UniqueId.IsSet() {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given NullableString and assigns it to the UniqueId field.
func (o *EventCreateQueryV2ResourceObjectAttributes) SetUniqueId(v string) {
	o.UniqueId.Set(&v)
}
// SetUniqueIdNil sets the value for UniqueId to be an explicit nil
func (o *EventCreateQueryV2ResourceObjectAttributes) SetUniqueIdNil() {
	o.UniqueId.Set(nil)
}

// UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
func (o *EventCreateQueryV2ResourceObjectAttributes) UnsetUniqueId() {
	o.UniqueId.Unset()
}

// GetMetric returns the Metric field value
func (o *EventCreateQueryV2ResourceObjectAttributes) GetMetric() EventCreateQueryV2ResourceObjectAttributesMetric {
	if o == nil {
		var ret EventCreateQueryV2ResourceObjectAttributesMetric
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *EventCreateQueryV2ResourceObjectAttributes) GetMetricOk() (*EventCreateQueryV2ResourceObjectAttributesMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *EventCreateQueryV2ResourceObjectAttributes) SetMetric(v EventCreateQueryV2ResourceObjectAttributesMetric) {
	o.Metric = v
}

// GetProfile returns the Profile field value
func (o *EventCreateQueryV2ResourceObjectAttributes) GetProfile() EventCreateQueryV2ResourceObjectAttributesProfile {
	if o == nil {
		var ret EventCreateQueryV2ResourceObjectAttributesProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *EventCreateQueryV2ResourceObjectAttributes) GetProfileOk() (*EventCreateQueryV2ResourceObjectAttributesProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *EventCreateQueryV2ResourceObjectAttributes) SetProfile(v EventCreateQueryV2ResourceObjectAttributesProfile) {
	o.Profile = v
}

func (o EventCreateQueryV2ResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventCreateQueryV2ResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["properties"] = o.Properties
	if o.Time.IsSet() {
		toSerialize["time"] = o.Time.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.ValueCurrency.IsSet() {
		toSerialize["value_currency"] = o.ValueCurrency.Get()
	}
	if o.UniqueId.IsSet() {
		toSerialize["unique_id"] = o.UniqueId.Get()
	}
	toSerialize["metric"] = o.Metric
	toSerialize["profile"] = o.Profile
	return toSerialize, nil
}

func (o *EventCreateQueryV2ResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"properties",
		"metric",
		"profile",
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

	varEventCreateQueryV2ResourceObjectAttributes := _EventCreateQueryV2ResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventCreateQueryV2ResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = EventCreateQueryV2ResourceObjectAttributes(varEventCreateQueryV2ResourceObjectAttributes)

	return err
}

type NullableEventCreateQueryV2ResourceObjectAttributes struct {
	value *EventCreateQueryV2ResourceObjectAttributes
	isSet bool
}

func (v NullableEventCreateQueryV2ResourceObjectAttributes) Get() *EventCreateQueryV2ResourceObjectAttributes {
	return v.value
}

func (v *NullableEventCreateQueryV2ResourceObjectAttributes) Set(val *EventCreateQueryV2ResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCreateQueryV2ResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCreateQueryV2ResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCreateQueryV2ResourceObjectAttributes(val *EventCreateQueryV2ResourceObjectAttributes) *NullableEventCreateQueryV2ResourceObjectAttributes {
	return &NullableEventCreateQueryV2ResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableEventCreateQueryV2ResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCreateQueryV2ResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

