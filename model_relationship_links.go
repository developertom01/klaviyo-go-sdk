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

// checks if the RelationshipLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelationshipLinks{}

// RelationshipLinks struct for RelationshipLinks
type RelationshipLinks struct {
	Self string `json:"self"`
	Related string `json:"related"`
}

type _RelationshipLinks RelationshipLinks

// NewRelationshipLinks instantiates a new RelationshipLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipLinks(self string, related string) *RelationshipLinks {
	this := RelationshipLinks{}
	this.Self = self
	this.Related = related
	return &this
}

// NewRelationshipLinksWithDefaults instantiates a new RelationshipLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipLinksWithDefaults() *RelationshipLinks {
	this := RelationshipLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *RelationshipLinks) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *RelationshipLinks) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *RelationshipLinks) SetSelf(v string) {
	o.Self = v
}

// GetRelated returns the Related field value
func (o *RelationshipLinks) GetRelated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
func (o *RelationshipLinks) GetRelatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Related, true
}

// SetRelated sets field value
func (o *RelationshipLinks) SetRelated(v string) {
	o.Related = v
}

func (o RelationshipLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelationshipLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	toSerialize["related"] = o.Related
	return toSerialize, nil
}

func (o *RelationshipLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
		"related",
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

	varRelationshipLinks := _RelationshipLinks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRelationshipLinks)

	if err != nil {
		return err
	}

	*o = RelationshipLinks(varRelationshipLinks)

	return err
}

type NullableRelationshipLinks struct {
	value *RelationshipLinks
	isSet bool
}

func (v NullableRelationshipLinks) Get() *RelationshipLinks {
	return v.value
}

func (v *NullableRelationshipLinks) Set(val *RelationshipLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipLinks(val *RelationshipLinks) *NullableRelationshipLinks {
	return &NullableRelationshipLinks{value: val, isSet: true}
}

func (v NullableRelationshipLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


