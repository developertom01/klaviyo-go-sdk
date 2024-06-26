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
	"fmt"
)

// EventBulkCreateEnum the model 'EventBulkCreateEnum'
type EventBulkCreateEnum string

// List of EventBulkCreateEnum
const (
	EVENT_BULK_CREATE EventBulkCreateEnum = "event-bulk-create"
)

// All allowed values of EventBulkCreateEnum enum
var AllowedEventBulkCreateEnumEnumValues = []EventBulkCreateEnum{
	"event-bulk-create",
}

func (v *EventBulkCreateEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventBulkCreateEnum(value)
	for _, existing := range AllowedEventBulkCreateEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventBulkCreateEnum", value)
}

// NewEventBulkCreateEnumFromValue returns a pointer to a valid EventBulkCreateEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventBulkCreateEnumFromValue(v string) (*EventBulkCreateEnum, error) {
	ev := EventBulkCreateEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventBulkCreateEnum: valid values are %v", v, AllowedEventBulkCreateEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventBulkCreateEnum) IsValid() bool {
	for _, existing := range AllowedEventBulkCreateEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventBulkCreateEnum value
func (v EventBulkCreateEnum) Ptr() *EventBulkCreateEnum {
	return &v
}

type NullableEventBulkCreateEnum struct {
	value *EventBulkCreateEnum
	isSet bool
}

func (v NullableEventBulkCreateEnum) Get() *EventBulkCreateEnum {
	return v.value
}

func (v *NullableEventBulkCreateEnum) Set(val *EventBulkCreateEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableEventBulkCreateEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableEventBulkCreateEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventBulkCreateEnum(val *EventBulkCreateEnum) *NullableEventBulkCreateEnum {
	return &NullableEventBulkCreateEnum{value: val, isSet: true}
}

func (v NullableEventBulkCreateEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventBulkCreateEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

