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

// checks if the PostProfileImportJobResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostProfileImportJobResponseData{}

// PostProfileImportJobResponseData struct for PostProfileImportJobResponseData
type PostProfileImportJobResponseData struct {
	Type ProfileBulkImportJobEnum `json:"type"`
	// Unique identifier for retrieving the job. Generated by Klaviyo.
	Id string `json:"id"`
	Attributes ProfileImportJobResponseObjectResourceAttributes `json:"attributes"`
	Relationships *PostProfileImportJobResponseDataRelationships `json:"relationships,omitempty"`
	Links ObjectLinks `json:"links"`
}

type _PostProfileImportJobResponseData PostProfileImportJobResponseData

// NewPostProfileImportJobResponseData instantiates a new PostProfileImportJobResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostProfileImportJobResponseData(type_ ProfileBulkImportJobEnum, id string, attributes ProfileImportJobResponseObjectResourceAttributes, links ObjectLinks) *PostProfileImportJobResponseData {
	this := PostProfileImportJobResponseData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewPostProfileImportJobResponseDataWithDefaults instantiates a new PostProfileImportJobResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostProfileImportJobResponseDataWithDefaults() *PostProfileImportJobResponseData {
	this := PostProfileImportJobResponseData{}
	return &this
}

// GetType returns the Type field value
func (o *PostProfileImportJobResponseData) GetType() ProfileBulkImportJobEnum {
	if o == nil {
		var ret ProfileBulkImportJobEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostProfileImportJobResponseData) GetTypeOk() (*ProfileBulkImportJobEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostProfileImportJobResponseData) SetType(v ProfileBulkImportJobEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PostProfileImportJobResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostProfileImportJobResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostProfileImportJobResponseData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PostProfileImportJobResponseData) GetAttributes() ProfileImportJobResponseObjectResourceAttributes {
	if o == nil {
		var ret ProfileImportJobResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PostProfileImportJobResponseData) GetAttributesOk() (*ProfileImportJobResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PostProfileImportJobResponseData) SetAttributes(v ProfileImportJobResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PostProfileImportJobResponseData) GetRelationships() PostProfileImportJobResponseDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PostProfileImportJobResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostProfileImportJobResponseData) GetRelationshipsOk() (*PostProfileImportJobResponseDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PostProfileImportJobResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PostProfileImportJobResponseDataRelationships and assigns it to the Relationships field.
func (o *PostProfileImportJobResponseData) SetRelationships(v PostProfileImportJobResponseDataRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *PostProfileImportJobResponseData) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PostProfileImportJobResponseData) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PostProfileImportJobResponseData) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o PostProfileImportJobResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostProfileImportJobResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *PostProfileImportJobResponseData) UnmarshalJSON(data []byte) (err error) {
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

	varPostProfileImportJobResponseData := _PostProfileImportJobResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostProfileImportJobResponseData)

	if err != nil {
		return err
	}

	*o = PostProfileImportJobResponseData(varPostProfileImportJobResponseData)

	return err
}

type NullablePostProfileImportJobResponseData struct {
	value *PostProfileImportJobResponseData
	isSet bool
}

func (v NullablePostProfileImportJobResponseData) Get() *PostProfileImportJobResponseData {
	return v.value
}

func (v *NullablePostProfileImportJobResponseData) Set(val *PostProfileImportJobResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePostProfileImportJobResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePostProfileImportJobResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostProfileImportJobResponseData(val *PostProfileImportJobResponseData) *NullablePostProfileImportJobResponseData {
	return &NullablePostProfileImportJobResponseData{value: val, isSet: true}
}

func (v NullablePostProfileImportJobResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostProfileImportJobResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

