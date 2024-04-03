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

// checks if the PostProfileImportJobResponseDataRelationshipsProfilesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostProfileImportJobResponseDataRelationshipsProfilesDataInner{}

// PostProfileImportJobResponseDataRelationshipsProfilesDataInner struct for PostProfileImportJobResponseDataRelationshipsProfilesDataInner
type PostProfileImportJobResponseDataRelationshipsProfilesDataInner struct {
	Type ProfileEnum `json:"type"`
	// IDs of the created/updated profiles
	Id string `json:"id"`
}

type _PostProfileImportJobResponseDataRelationshipsProfilesDataInner PostProfileImportJobResponseDataRelationshipsProfilesDataInner

// NewPostProfileImportJobResponseDataRelationshipsProfilesDataInner instantiates a new PostProfileImportJobResponseDataRelationshipsProfilesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostProfileImportJobResponseDataRelationshipsProfilesDataInner(type_ ProfileEnum, id string) *PostProfileImportJobResponseDataRelationshipsProfilesDataInner {
	this := PostProfileImportJobResponseDataRelationshipsProfilesDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewPostProfileImportJobResponseDataRelationshipsProfilesDataInnerWithDefaults instantiates a new PostProfileImportJobResponseDataRelationshipsProfilesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostProfileImportJobResponseDataRelationshipsProfilesDataInnerWithDefaults() *PostProfileImportJobResponseDataRelationshipsProfilesDataInner {
	this := PostProfileImportJobResponseDataRelationshipsProfilesDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *PostProfileImportJobResponseDataRelationshipsProfilesDataInner) GetType() ProfileEnum {
	if o == nil {
		var ret ProfileEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostProfileImportJobResponseDataRelationshipsProfilesDataInner) GetTypeOk() (*ProfileEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostProfileImportJobResponseDataRelationshipsProfilesDataInner) SetType(v ProfileEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PostProfileImportJobResponseDataRelationshipsProfilesDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostProfileImportJobResponseDataRelationshipsProfilesDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostProfileImportJobResponseDataRelationshipsProfilesDataInner) SetId(v string) {
	o.Id = v
}

func (o PostProfileImportJobResponseDataRelationshipsProfilesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostProfileImportJobResponseDataRelationshipsProfilesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *PostProfileImportJobResponseDataRelationshipsProfilesDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varPostProfileImportJobResponseDataRelationshipsProfilesDataInner := _PostProfileImportJobResponseDataRelationshipsProfilesDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostProfileImportJobResponseDataRelationshipsProfilesDataInner)

	if err != nil {
		return err
	}

	*o = PostProfileImportJobResponseDataRelationshipsProfilesDataInner(varPostProfileImportJobResponseDataRelationshipsProfilesDataInner)

	return err
}

type NullablePostProfileImportJobResponseDataRelationshipsProfilesDataInner struct {
	value *PostProfileImportJobResponseDataRelationshipsProfilesDataInner
	isSet bool
}

func (v NullablePostProfileImportJobResponseDataRelationshipsProfilesDataInner) Get() *PostProfileImportJobResponseDataRelationshipsProfilesDataInner {
	return v.value
}

func (v *NullablePostProfileImportJobResponseDataRelationshipsProfilesDataInner) Set(val *PostProfileImportJobResponseDataRelationshipsProfilesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePostProfileImportJobResponseDataRelationshipsProfilesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePostProfileImportJobResponseDataRelationshipsProfilesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostProfileImportJobResponseDataRelationshipsProfilesDataInner(val *PostProfileImportJobResponseDataRelationshipsProfilesDataInner) *NullablePostProfileImportJobResponseDataRelationshipsProfilesDataInner {
	return &NullablePostProfileImportJobResponseDataRelationshipsProfilesDataInner{value: val, isSet: true}
}

func (v NullablePostProfileImportJobResponseDataRelationshipsProfilesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostProfileImportJobResponseDataRelationshipsProfilesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

