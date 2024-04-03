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

// checks if the SegmentPartialUpdateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentPartialUpdateQueryResourceObjectAttributes{}

// SegmentPartialUpdateQueryResourceObjectAttributes struct for SegmentPartialUpdateQueryResourceObjectAttributes
type SegmentPartialUpdateQueryResourceObjectAttributes struct {
	Name NullableString `json:"name,omitempty"`
	IsStarred NullableBool `json:"is_starred,omitempty"`
}

// NewSegmentPartialUpdateQueryResourceObjectAttributes instantiates a new SegmentPartialUpdateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentPartialUpdateQueryResourceObjectAttributes() *SegmentPartialUpdateQueryResourceObjectAttributes {
	this := SegmentPartialUpdateQueryResourceObjectAttributes{}
	return &this
}

// NewSegmentPartialUpdateQueryResourceObjectAttributesWithDefaults instantiates a new SegmentPartialUpdateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentPartialUpdateQueryResourceObjectAttributesWithDefaults() *SegmentPartialUpdateQueryResourceObjectAttributes {
	this := SegmentPartialUpdateQueryResourceObjectAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentPartialUpdateQueryResourceObjectAttributes) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentPartialUpdateQueryResourceObjectAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SegmentPartialUpdateQueryResourceObjectAttributes) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SegmentPartialUpdateQueryResourceObjectAttributes) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SegmentPartialUpdateQueryResourceObjectAttributes) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SegmentPartialUpdateQueryResourceObjectAttributes) UnsetName() {
	o.Name.Unset()
}

// GetIsStarred returns the IsStarred field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentPartialUpdateQueryResourceObjectAttributes) GetIsStarred() bool {
	if o == nil || IsNil(o.IsStarred.Get()) {
		var ret bool
		return ret
	}
	return *o.IsStarred.Get()
}

// GetIsStarredOk returns a tuple with the IsStarred field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentPartialUpdateQueryResourceObjectAttributes) GetIsStarredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsStarred.Get(), o.IsStarred.IsSet()
}

// HasIsStarred returns a boolean if a field has been set.
func (o *SegmentPartialUpdateQueryResourceObjectAttributes) HasIsStarred() bool {
	if o != nil && o.IsStarred.IsSet() {
		return true
	}

	return false
}

// SetIsStarred gets a reference to the given NullableBool and assigns it to the IsStarred field.
func (o *SegmentPartialUpdateQueryResourceObjectAttributes) SetIsStarred(v bool) {
	o.IsStarred.Set(&v)
}
// SetIsStarredNil sets the value for IsStarred to be an explicit nil
func (o *SegmentPartialUpdateQueryResourceObjectAttributes) SetIsStarredNil() {
	o.IsStarred.Set(nil)
}

// UnsetIsStarred ensures that no value is present for IsStarred, not even an explicit nil
func (o *SegmentPartialUpdateQueryResourceObjectAttributes) UnsetIsStarred() {
	o.IsStarred.Unset()
}

func (o SegmentPartialUpdateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentPartialUpdateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.IsStarred.IsSet() {
		toSerialize["is_starred"] = o.IsStarred.Get()
	}
	return toSerialize, nil
}

type NullableSegmentPartialUpdateQueryResourceObjectAttributes struct {
	value *SegmentPartialUpdateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableSegmentPartialUpdateQueryResourceObjectAttributes) Get() *SegmentPartialUpdateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableSegmentPartialUpdateQueryResourceObjectAttributes) Set(val *SegmentPartialUpdateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentPartialUpdateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentPartialUpdateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentPartialUpdateQueryResourceObjectAttributes(val *SegmentPartialUpdateQueryResourceObjectAttributes) *NullableSegmentPartialUpdateQueryResourceObjectAttributes {
	return &NullableSegmentPartialUpdateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableSegmentPartialUpdateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentPartialUpdateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

