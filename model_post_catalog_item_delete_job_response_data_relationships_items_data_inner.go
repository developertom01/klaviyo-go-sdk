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

// checks if the PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner{}

// PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner struct for PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner
type PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner struct {
	Type CatalogItemEnum `json:"type"`
	// IDs of the deleted catalog items.
	Id string `json:"id"`
}

type _PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner

// NewPostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner instantiates a new PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner(type_ CatalogItemEnum, id string) *PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner {
	this := PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewPostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInnerWithDefaults instantiates a new PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInnerWithDefaults() *PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner {
	this := PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) GetType() CatalogItemEnum {
	if o == nil {
		var ret CatalogItemEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) GetTypeOk() (*CatalogItemEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) SetType(v CatalogItemEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) SetId(v string) {
	o.Id = v
}

func (o PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) UnmarshalJSON(data []byte) (err error) {
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

	varPostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner := _PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner)

	if err != nil {
		return err
	}

	*o = PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner(varPostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner)

	return err
}

type NullablePostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner struct {
	value *PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner
	isSet bool
}

func (v NullablePostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) Get() *PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner {
	return v.value
}

func (v *NullablePostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) Set(val *PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner(val *PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) *NullablePostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner {
	return &NullablePostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner{value: val, isSet: true}
}

func (v NullablePostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


