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

// checks if the SegmentResponseObjectResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentResponseObjectResource{}

// SegmentResponseObjectResource struct for SegmentResponseObjectResource
type SegmentResponseObjectResource struct {
	Type SegmentEnum `json:"type"`
	Id string `json:"id"`
	Attributes SegmentListResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
}

type _SegmentResponseObjectResource SegmentResponseObjectResource

// NewSegmentResponseObjectResource instantiates a new SegmentResponseObjectResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentResponseObjectResource(type_ SegmentEnum, id string, attributes SegmentListResponseObjectResourceAttributes, links ObjectLinks) *SegmentResponseObjectResource {
	this := SegmentResponseObjectResource{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewSegmentResponseObjectResourceWithDefaults instantiates a new SegmentResponseObjectResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentResponseObjectResourceWithDefaults() *SegmentResponseObjectResource {
	this := SegmentResponseObjectResource{}
	return &this
}

// GetType returns the Type field value
func (o *SegmentResponseObjectResource) GetType() SegmentEnum {
	if o == nil {
		var ret SegmentEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SegmentResponseObjectResource) GetTypeOk() (*SegmentEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SegmentResponseObjectResource) SetType(v SegmentEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SegmentResponseObjectResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SegmentResponseObjectResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SegmentResponseObjectResource) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *SegmentResponseObjectResource) GetAttributes() SegmentListResponseObjectResourceAttributes {
	if o == nil {
		var ret SegmentListResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SegmentResponseObjectResource) GetAttributesOk() (*SegmentListResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SegmentResponseObjectResource) SetAttributes(v SegmentListResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *SegmentResponseObjectResource) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SegmentResponseObjectResource) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SegmentResponseObjectResource) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o SegmentResponseObjectResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentResponseObjectResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *SegmentResponseObjectResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varSegmentResponseObjectResource := _SegmentResponseObjectResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSegmentResponseObjectResource)

	if err != nil {
		return err
	}

	*o = SegmentResponseObjectResource(varSegmentResponseObjectResource)

	return err
}

type NullableSegmentResponseObjectResource struct {
	value *SegmentResponseObjectResource
	isSet bool
}

func (v NullableSegmentResponseObjectResource) Get() *SegmentResponseObjectResource {
	return v.value
}

func (v *NullableSegmentResponseObjectResource) Set(val *SegmentResponseObjectResource) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentResponseObjectResource) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentResponseObjectResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentResponseObjectResource(val *SegmentResponseObjectResource) *NullableSegmentResponseObjectResource {
	return &NullableSegmentResponseObjectResource{value: val, isSet: true}
}

func (v NullableSegmentResponseObjectResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentResponseObjectResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

