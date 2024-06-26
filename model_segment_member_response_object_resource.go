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

// checks if the SegmentMemberResponseObjectResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentMemberResponseObjectResource{}

// SegmentMemberResponseObjectResource struct for SegmentMemberResponseObjectResource
type SegmentMemberResponseObjectResource struct {
	Type ProfileEnum `json:"type"`
	// Primary key that uniquely identifies this profile. Generated by Klaviyo.
	Id NullableString `json:"id,omitempty"`
	Attributes SegmentMemberResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
}

type _SegmentMemberResponseObjectResource SegmentMemberResponseObjectResource

// NewSegmentMemberResponseObjectResource instantiates a new SegmentMemberResponseObjectResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentMemberResponseObjectResource(type_ ProfileEnum, attributes SegmentMemberResponseObjectResourceAttributes, links ObjectLinks) *SegmentMemberResponseObjectResource {
	this := SegmentMemberResponseObjectResource{}
	this.Type = type_
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewSegmentMemberResponseObjectResourceWithDefaults instantiates a new SegmentMemberResponseObjectResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentMemberResponseObjectResourceWithDefaults() *SegmentMemberResponseObjectResource {
	this := SegmentMemberResponseObjectResource{}
	return &this
}

// GetType returns the Type field value
func (o *SegmentMemberResponseObjectResource) GetType() ProfileEnum {
	if o == nil {
		var ret ProfileEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SegmentMemberResponseObjectResource) GetTypeOk() (*ProfileEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SegmentMemberResponseObjectResource) SetType(v ProfileEnum) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentMemberResponseObjectResource) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentMemberResponseObjectResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *SegmentMemberResponseObjectResource) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *SegmentMemberResponseObjectResource) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *SegmentMemberResponseObjectResource) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *SegmentMemberResponseObjectResource) UnsetId() {
	o.Id.Unset()
}

// GetAttributes returns the Attributes field value
func (o *SegmentMemberResponseObjectResource) GetAttributes() SegmentMemberResponseObjectResourceAttributes {
	if o == nil {
		var ret SegmentMemberResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SegmentMemberResponseObjectResource) GetAttributesOk() (*SegmentMemberResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SegmentMemberResponseObjectResource) SetAttributes(v SegmentMemberResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *SegmentMemberResponseObjectResource) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SegmentMemberResponseObjectResource) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SegmentMemberResponseObjectResource) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o SegmentMemberResponseObjectResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentMemberResponseObjectResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *SegmentMemberResponseObjectResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
		"links",
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

	varSegmentMemberResponseObjectResource := _SegmentMemberResponseObjectResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSegmentMemberResponseObjectResource)

	if err != nil {
		return err
	}

	*o = SegmentMemberResponseObjectResource(varSegmentMemberResponseObjectResource)

	return err
}

type NullableSegmentMemberResponseObjectResource struct {
	value *SegmentMemberResponseObjectResource
	isSet bool
}

func (v NullableSegmentMemberResponseObjectResource) Get() *SegmentMemberResponseObjectResource {
	return v.value
}

func (v *NullableSegmentMemberResponseObjectResource) Set(val *SegmentMemberResponseObjectResource) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentMemberResponseObjectResource) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentMemberResponseObjectResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentMemberResponseObjectResource(val *SegmentMemberResponseObjectResource) *NullableSegmentMemberResponseObjectResource {
	return &NullableSegmentMemberResponseObjectResource{value: val, isSet: true}
}

func (v NullableSegmentMemberResponseObjectResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentMemberResponseObjectResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


