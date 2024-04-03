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

// ImageEnum the model 'ImageEnum'
type ImageEnum string

// List of ImageEnum
const (
	IMAGE ImageEnum = "image"
)

// All allowed values of ImageEnum enum
var AllowedImageEnumEnumValues = []ImageEnum{
	"image",
}

func (v *ImageEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImageEnum(value)
	for _, existing := range AllowedImageEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImageEnum", value)
}

// NewImageEnumFromValue returns a pointer to a valid ImageEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageEnumFromValue(v string) (*ImageEnum, error) {
	ev := ImageEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImageEnum: valid values are %v", v, AllowedImageEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageEnum) IsValid() bool {
	for _, existing := range AllowedImageEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageEnum value
func (v ImageEnum) Ptr() *ImageEnum {
	return &v
}

type NullableImageEnum struct {
	value *ImageEnum
	isSet bool
}

func (v NullableImageEnum) Get() *ImageEnum {
	return v.value
}

func (v *NullableImageEnum) Set(val *ImageEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableImageEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableImageEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageEnum(val *ImageEnum) *NullableImageEnum {
	return &NullableImageEnum{value: val, isSet: true}
}

func (v NullableImageEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
