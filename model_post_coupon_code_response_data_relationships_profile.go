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

// checks if the PostCouponCodeResponseDataRelationshipsProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCouponCodeResponseDataRelationshipsProfile{}

// PostCouponCodeResponseDataRelationshipsProfile struct for PostCouponCodeResponseDataRelationshipsProfile
type PostCouponCodeResponseDataRelationshipsProfile struct {
	Data PostCouponCodeResponseDataRelationshipsProfileData `json:"data"`
	Links *RelationshipLinks `json:"links,omitempty"`
}

type _PostCouponCodeResponseDataRelationshipsProfile PostCouponCodeResponseDataRelationshipsProfile

// NewPostCouponCodeResponseDataRelationshipsProfile instantiates a new PostCouponCodeResponseDataRelationshipsProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCouponCodeResponseDataRelationshipsProfile(data PostCouponCodeResponseDataRelationshipsProfileData) *PostCouponCodeResponseDataRelationshipsProfile {
	this := PostCouponCodeResponseDataRelationshipsProfile{}
	this.Data = data
	return &this
}

// NewPostCouponCodeResponseDataRelationshipsProfileWithDefaults instantiates a new PostCouponCodeResponseDataRelationshipsProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCouponCodeResponseDataRelationshipsProfileWithDefaults() *PostCouponCodeResponseDataRelationshipsProfile {
	this := PostCouponCodeResponseDataRelationshipsProfile{}
	return &this
}

// GetData returns the Data field value
func (o *PostCouponCodeResponseDataRelationshipsProfile) GetData() PostCouponCodeResponseDataRelationshipsProfileData {
	if o == nil {
		var ret PostCouponCodeResponseDataRelationshipsProfileData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PostCouponCodeResponseDataRelationshipsProfile) GetDataOk() (*PostCouponCodeResponseDataRelationshipsProfileData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PostCouponCodeResponseDataRelationshipsProfile) SetData(v PostCouponCodeResponseDataRelationshipsProfileData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostCouponCodeResponseDataRelationshipsProfile) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCouponCodeResponseDataRelationshipsProfile) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostCouponCodeResponseDataRelationshipsProfile) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *PostCouponCodeResponseDataRelationshipsProfile) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

func (o PostCouponCodeResponseDataRelationshipsProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCouponCodeResponseDataRelationshipsProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *PostCouponCodeResponseDataRelationshipsProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varPostCouponCodeResponseDataRelationshipsProfile := _PostCouponCodeResponseDataRelationshipsProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCouponCodeResponseDataRelationshipsProfile)

	if err != nil {
		return err
	}

	*o = PostCouponCodeResponseDataRelationshipsProfile(varPostCouponCodeResponseDataRelationshipsProfile)

	return err
}

type NullablePostCouponCodeResponseDataRelationshipsProfile struct {
	value *PostCouponCodeResponseDataRelationshipsProfile
	isSet bool
}

func (v NullablePostCouponCodeResponseDataRelationshipsProfile) Get() *PostCouponCodeResponseDataRelationshipsProfile {
	return v.value
}

func (v *NullablePostCouponCodeResponseDataRelationshipsProfile) Set(val *PostCouponCodeResponseDataRelationshipsProfile) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCouponCodeResponseDataRelationshipsProfile) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCouponCodeResponseDataRelationshipsProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCouponCodeResponseDataRelationshipsProfile(val *PostCouponCodeResponseDataRelationshipsProfile) *NullablePostCouponCodeResponseDataRelationshipsProfile {
	return &NullablePostCouponCodeResponseDataRelationshipsProfile{value: val, isSet: true}
}

func (v NullablePostCouponCodeResponseDataRelationshipsProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCouponCodeResponseDataRelationshipsProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


