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
)

// checks if the ImagePartialUpdateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImagePartialUpdateQueryResourceObjectAttributes{}

// ImagePartialUpdateQueryResourceObjectAttributes struct for ImagePartialUpdateQueryResourceObjectAttributes
type ImagePartialUpdateQueryResourceObjectAttributes struct {
	// A name for the image.
	Name NullableString `json:"name,omitempty"`
	// If true, this image is not shown in the asset library.
	Hidden NullableBool `json:"hidden,omitempty"`
}

// NewImagePartialUpdateQueryResourceObjectAttributes instantiates a new ImagePartialUpdateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImagePartialUpdateQueryResourceObjectAttributes() *ImagePartialUpdateQueryResourceObjectAttributes {
	this := ImagePartialUpdateQueryResourceObjectAttributes{}
	return &this
}

// NewImagePartialUpdateQueryResourceObjectAttributesWithDefaults instantiates a new ImagePartialUpdateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImagePartialUpdateQueryResourceObjectAttributesWithDefaults() *ImagePartialUpdateQueryResourceObjectAttributes {
	this := ImagePartialUpdateQueryResourceObjectAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImagePartialUpdateQueryResourceObjectAttributes) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImagePartialUpdateQueryResourceObjectAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ImagePartialUpdateQueryResourceObjectAttributes) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ImagePartialUpdateQueryResourceObjectAttributes) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ImagePartialUpdateQueryResourceObjectAttributes) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ImagePartialUpdateQueryResourceObjectAttributes) UnsetName() {
	o.Name.Unset()
}

// GetHidden returns the Hidden field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImagePartialUpdateQueryResourceObjectAttributes) GetHidden() bool {
	if o == nil || IsNil(o.Hidden.Get()) {
		var ret bool
		return ret
	}
	return *o.Hidden.Get()
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImagePartialUpdateQueryResourceObjectAttributes) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hidden.Get(), o.Hidden.IsSet()
}

// HasHidden returns a boolean if a field has been set.
func (o *ImagePartialUpdateQueryResourceObjectAttributes) HasHidden() bool {
	if o != nil && o.Hidden.IsSet() {
		return true
	}

	return false
}

// SetHidden gets a reference to the given NullableBool and assigns it to the Hidden field.
func (o *ImagePartialUpdateQueryResourceObjectAttributes) SetHidden(v bool) {
	o.Hidden.Set(&v)
}
// SetHiddenNil sets the value for Hidden to be an explicit nil
func (o *ImagePartialUpdateQueryResourceObjectAttributes) SetHiddenNil() {
	o.Hidden.Set(nil)
}

// UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
func (o *ImagePartialUpdateQueryResourceObjectAttributes) UnsetHidden() {
	o.Hidden.Unset()
}

func (o ImagePartialUpdateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImagePartialUpdateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Hidden.IsSet() {
		toSerialize["hidden"] = o.Hidden.Get()
	}
	return toSerialize, nil
}

type NullableImagePartialUpdateQueryResourceObjectAttributes struct {
	value *ImagePartialUpdateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableImagePartialUpdateQueryResourceObjectAttributes) Get() *ImagePartialUpdateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableImagePartialUpdateQueryResourceObjectAttributes) Set(val *ImagePartialUpdateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableImagePartialUpdateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableImagePartialUpdateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImagePartialUpdateQueryResourceObjectAttributes(val *ImagePartialUpdateQueryResourceObjectAttributes) *NullableImagePartialUpdateQueryResourceObjectAttributes {
	return &NullableImagePartialUpdateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableImagePartialUpdateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImagePartialUpdateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


