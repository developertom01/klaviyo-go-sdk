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

// checks if the ImageResponseObjectResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageResponseObjectResourceAttributes{}

// ImageResponseObjectResourceAttributes struct for ImageResponseObjectResourceAttributes
type ImageResponseObjectResourceAttributes struct {
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
	Format string `json:"format"`
	Size int32 `json:"size"`
	Hidden bool `json:"hidden"`
	UpdatedAt time.Time `json:"updated_at"`
}

type _ImageResponseObjectResourceAttributes ImageResponseObjectResourceAttributes

// NewImageResponseObjectResourceAttributes instantiates a new ImageResponseObjectResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageResponseObjectResourceAttributes(name string, imageUrl string, format string, size int32, hidden bool, updatedAt time.Time) *ImageResponseObjectResourceAttributes {
	this := ImageResponseObjectResourceAttributes{}
	this.Name = name
	this.ImageUrl = imageUrl
	this.Format = format
	this.Size = size
	this.Hidden = hidden
	this.UpdatedAt = updatedAt
	return &this
}

// NewImageResponseObjectResourceAttributesWithDefaults instantiates a new ImageResponseObjectResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageResponseObjectResourceAttributesWithDefaults() *ImageResponseObjectResourceAttributes {
	this := ImageResponseObjectResourceAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *ImageResponseObjectResourceAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImageResponseObjectResourceAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImageResponseObjectResourceAttributes) SetName(v string) {
	o.Name = v
}

// GetImageUrl returns the ImageUrl field value
func (o *ImageResponseObjectResourceAttributes) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *ImageResponseObjectResourceAttributes) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *ImageResponseObjectResourceAttributes) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetFormat returns the Format field value
func (o *ImageResponseObjectResourceAttributes) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *ImageResponseObjectResourceAttributes) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *ImageResponseObjectResourceAttributes) SetFormat(v string) {
	o.Format = v
}

// GetSize returns the Size field value
func (o *ImageResponseObjectResourceAttributes) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ImageResponseObjectResourceAttributes) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ImageResponseObjectResourceAttributes) SetSize(v int32) {
	o.Size = v
}

// GetHidden returns the Hidden field value
func (o *ImageResponseObjectResourceAttributes) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *ImageResponseObjectResourceAttributes) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *ImageResponseObjectResourceAttributes) SetHidden(v bool) {
	o.Hidden = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ImageResponseObjectResourceAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ImageResponseObjectResourceAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ImageResponseObjectResourceAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ImageResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageResponseObjectResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["image_url"] = o.ImageUrl
	toSerialize["format"] = o.Format
	toSerialize["size"] = o.Size
	toSerialize["hidden"] = o.Hidden
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ImageResponseObjectResourceAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"image_url",
		"format",
		"size",
		"hidden",
		"updated_at",
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

	varImageResponseObjectResourceAttributes := _ImageResponseObjectResourceAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImageResponseObjectResourceAttributes)

	if err != nil {
		return err
	}

	*o = ImageResponseObjectResourceAttributes(varImageResponseObjectResourceAttributes)

	return err
}

type NullableImageResponseObjectResourceAttributes struct {
	value *ImageResponseObjectResourceAttributes
	isSet bool
}

func (v NullableImageResponseObjectResourceAttributes) Get() *ImageResponseObjectResourceAttributes {
	return v.value
}

func (v *NullableImageResponseObjectResourceAttributes) Set(val *ImageResponseObjectResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableImageResponseObjectResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableImageResponseObjectResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageResponseObjectResourceAttributes(val *ImageResponseObjectResourceAttributes) *NullableImageResponseObjectResourceAttributes {
	return &NullableImageResponseObjectResourceAttributes{value: val, isSet: true}
}

func (v NullableImageResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageResponseObjectResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


