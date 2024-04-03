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

// checks if the TagGroupResponseObjectResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagGroupResponseObjectResource{}

// TagGroupResponseObjectResource struct for TagGroupResponseObjectResource
type TagGroupResponseObjectResource struct {
	Type TagGroupEnum `json:"type"`
	// The Tag Group ID
	Id string `json:"id"`
	Attributes TagGroupResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
}

type _TagGroupResponseObjectResource TagGroupResponseObjectResource

// NewTagGroupResponseObjectResource instantiates a new TagGroupResponseObjectResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagGroupResponseObjectResource(type_ TagGroupEnum, id string, attributes TagGroupResponseObjectResourceAttributes, links ObjectLinks) *TagGroupResponseObjectResource {
	this := TagGroupResponseObjectResource{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewTagGroupResponseObjectResourceWithDefaults instantiates a new TagGroupResponseObjectResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagGroupResponseObjectResourceWithDefaults() *TagGroupResponseObjectResource {
	this := TagGroupResponseObjectResource{}
	return &this
}

// GetType returns the Type field value
func (o *TagGroupResponseObjectResource) GetType() TagGroupEnum {
	if o == nil {
		var ret TagGroupEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TagGroupResponseObjectResource) GetTypeOk() (*TagGroupEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TagGroupResponseObjectResource) SetType(v TagGroupEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TagGroupResponseObjectResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagGroupResponseObjectResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TagGroupResponseObjectResource) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *TagGroupResponseObjectResource) GetAttributes() TagGroupResponseObjectResourceAttributes {
	if o == nil {
		var ret TagGroupResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TagGroupResponseObjectResource) GetAttributesOk() (*TagGroupResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TagGroupResponseObjectResource) SetAttributes(v TagGroupResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *TagGroupResponseObjectResource) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *TagGroupResponseObjectResource) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *TagGroupResponseObjectResource) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o TagGroupResponseObjectResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagGroupResponseObjectResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *TagGroupResponseObjectResource) UnmarshalJSON(data []byte) (err error) {
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

	varTagGroupResponseObjectResource := _TagGroupResponseObjectResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagGroupResponseObjectResource)

	if err != nil {
		return err
	}

	*o = TagGroupResponseObjectResource(varTagGroupResponseObjectResource)

	return err
}

type NullableTagGroupResponseObjectResource struct {
	value *TagGroupResponseObjectResource
	isSet bool
}

func (v NullableTagGroupResponseObjectResource) Get() *TagGroupResponseObjectResource {
	return v.value
}

func (v *NullableTagGroupResponseObjectResource) Set(val *TagGroupResponseObjectResource) {
	v.value = val
	v.isSet = true
}

func (v NullableTagGroupResponseObjectResource) IsSet() bool {
	return v.isSet
}

func (v *NullableTagGroupResponseObjectResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagGroupResponseObjectResource(val *TagGroupResponseObjectResource) *NullableTagGroupResponseObjectResource {
	return &NullableTagGroupResponseObjectResource{value: val, isSet: true}
}

func (v NullableTagGroupResponseObjectResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagGroupResponseObjectResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


