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

// checks if the ImageCreateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageCreateQueryResourceObjectAttributes{}

// ImageCreateQueryResourceObjectAttributes struct for ImageCreateQueryResourceObjectAttributes
type ImageCreateQueryResourceObjectAttributes struct {
	// A name for the image.  Defaults to the filename if not provided.  If the name matches an existing image, a suffix will be added.
	Name NullableString `json:"name,omitempty"`
	// An existing image url to import the image from. Alternatively, you may specify a base-64 encoded data-uri (`data:image/...`). Supported image formats: jpeg,png,gif. Maximum image size: 5MB.
	ImportFromUrl string `json:"import_from_url"`
	// If true, this image is not shown in the asset library.
	Hidden NullableBool `json:"hidden,omitempty"`
}

type _ImageCreateQueryResourceObjectAttributes ImageCreateQueryResourceObjectAttributes

// NewImageCreateQueryResourceObjectAttributes instantiates a new ImageCreateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageCreateQueryResourceObjectAttributes(importFromUrl string) *ImageCreateQueryResourceObjectAttributes {
	this := ImageCreateQueryResourceObjectAttributes{}
	this.ImportFromUrl = importFromUrl
	var hidden bool = false
	this.Hidden = *NewNullableBool(&hidden)
	return &this
}

// NewImageCreateQueryResourceObjectAttributesWithDefaults instantiates a new ImageCreateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageCreateQueryResourceObjectAttributesWithDefaults() *ImageCreateQueryResourceObjectAttributes {
	this := ImageCreateQueryResourceObjectAttributes{}
	var hidden bool = false
	this.Hidden = *NewNullableBool(&hidden)
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageCreateQueryResourceObjectAttributes) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageCreateQueryResourceObjectAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ImageCreateQueryResourceObjectAttributes) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ImageCreateQueryResourceObjectAttributes) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ImageCreateQueryResourceObjectAttributes) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ImageCreateQueryResourceObjectAttributes) UnsetName() {
	o.Name.Unset()
}

// GetImportFromUrl returns the ImportFromUrl field value
func (o *ImageCreateQueryResourceObjectAttributes) GetImportFromUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImportFromUrl
}

// GetImportFromUrlOk returns a tuple with the ImportFromUrl field value
// and a boolean to check if the value has been set.
func (o *ImageCreateQueryResourceObjectAttributes) GetImportFromUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportFromUrl, true
}

// SetImportFromUrl sets field value
func (o *ImageCreateQueryResourceObjectAttributes) SetImportFromUrl(v string) {
	o.ImportFromUrl = v
}

// GetHidden returns the Hidden field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageCreateQueryResourceObjectAttributes) GetHidden() bool {
	if o == nil || IsNil(o.Hidden.Get()) {
		var ret bool
		return ret
	}
	return *o.Hidden.Get()
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageCreateQueryResourceObjectAttributes) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hidden.Get(), o.Hidden.IsSet()
}

// HasHidden returns a boolean if a field has been set.
func (o *ImageCreateQueryResourceObjectAttributes) HasHidden() bool {
	if o != nil && o.Hidden.IsSet() {
		return true
	}

	return false
}

// SetHidden gets a reference to the given NullableBool and assigns it to the Hidden field.
func (o *ImageCreateQueryResourceObjectAttributes) SetHidden(v bool) {
	o.Hidden.Set(&v)
}
// SetHiddenNil sets the value for Hidden to be an explicit nil
func (o *ImageCreateQueryResourceObjectAttributes) SetHiddenNil() {
	o.Hidden.Set(nil)
}

// UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
func (o *ImageCreateQueryResourceObjectAttributes) UnsetHidden() {
	o.Hidden.Unset()
}

func (o ImageCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageCreateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["import_from_url"] = o.ImportFromUrl
	if o.Hidden.IsSet() {
		toSerialize["hidden"] = o.Hidden.Get()
	}
	return toSerialize, nil
}

func (o *ImageCreateQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"import_from_url",
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

	varImageCreateQueryResourceObjectAttributes := _ImageCreateQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImageCreateQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = ImageCreateQueryResourceObjectAttributes(varImageCreateQueryResourceObjectAttributes)

	return err
}

type NullableImageCreateQueryResourceObjectAttributes struct {
	value *ImageCreateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableImageCreateQueryResourceObjectAttributes) Get() *ImageCreateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableImageCreateQueryResourceObjectAttributes) Set(val *ImageCreateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableImageCreateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableImageCreateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageCreateQueryResourceObjectAttributes(val *ImageCreateQueryResourceObjectAttributes) *NullableImageCreateQueryResourceObjectAttributes {
	return &NullableImageCreateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableImageCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageCreateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


