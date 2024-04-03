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

// ProfileMetaPatchPropertiesUnset - Remove a key or keys (and their values) completely from properties
type ProfileMetaPatchPropertiesUnset struct {
	ArrayOfString *[]string
	String *string
}

// []stringAsProfileMetaPatchPropertiesUnset is a convenience function that returns []string wrapped in ProfileMetaPatchPropertiesUnset
func ArrayOfStringAsProfileMetaPatchPropertiesUnset(v *[]string) ProfileMetaPatchPropertiesUnset {
	return ProfileMetaPatchPropertiesUnset{
		ArrayOfString: v,
	}
}

// stringAsProfileMetaPatchPropertiesUnset is a convenience function that returns string wrapped in ProfileMetaPatchPropertiesUnset
func StringAsProfileMetaPatchPropertiesUnset(v *string) ProfileMetaPatchPropertiesUnset {
	return ProfileMetaPatchPropertiesUnset{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ProfileMetaPatchPropertiesUnset) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfString = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ProfileMetaPatchPropertiesUnset)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ProfileMetaPatchPropertiesUnset)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ProfileMetaPatchPropertiesUnset) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ProfileMetaPatchPropertiesUnset) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableProfileMetaPatchPropertiesUnset struct {
	value *ProfileMetaPatchPropertiesUnset
	isSet bool
}

func (v NullableProfileMetaPatchPropertiesUnset) Get() *ProfileMetaPatchPropertiesUnset {
	return v.value
}

func (v *NullableProfileMetaPatchPropertiesUnset) Set(val *ProfileMetaPatchPropertiesUnset) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileMetaPatchPropertiesUnset) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileMetaPatchPropertiesUnset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileMetaPatchPropertiesUnset(val *ProfileMetaPatchPropertiesUnset) *NullableProfileMetaPatchPropertiesUnset {
	return &NullableProfileMetaPatchPropertiesUnset{value: val, isSet: true}
}

func (v NullableProfileMetaPatchPropertiesUnset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileMetaPatchPropertiesUnset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

