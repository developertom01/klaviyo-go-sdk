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

// checks if the GetSegmentResponseCollectionDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSegmentResponseCollectionDataInner{}

// GetSegmentResponseCollectionDataInner struct for GetSegmentResponseCollectionDataInner
type GetSegmentResponseCollectionDataInner struct {
	Type SegmentEnum `json:"type"`
	Id string `json:"id"`
	Attributes SegmentListResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
	Relationships *GetListResponseCollectionDataInnerAllOfRelationships `json:"relationships,omitempty"`
}

type _GetSegmentResponseCollectionDataInner GetSegmentResponseCollectionDataInner

// NewGetSegmentResponseCollectionDataInner instantiates a new GetSegmentResponseCollectionDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSegmentResponseCollectionDataInner(type_ SegmentEnum, id string, attributes SegmentListResponseObjectResourceAttributes, links ObjectLinks) *GetSegmentResponseCollectionDataInner {
	this := GetSegmentResponseCollectionDataInner{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewGetSegmentResponseCollectionDataInnerWithDefaults instantiates a new GetSegmentResponseCollectionDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSegmentResponseCollectionDataInnerWithDefaults() *GetSegmentResponseCollectionDataInner {
	this := GetSegmentResponseCollectionDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GetSegmentResponseCollectionDataInner) GetType() SegmentEnum {
	if o == nil {
		var ret SegmentEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetSegmentResponseCollectionDataInner) GetTypeOk() (*SegmentEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetSegmentResponseCollectionDataInner) SetType(v SegmentEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetSegmentResponseCollectionDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetSegmentResponseCollectionDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetSegmentResponseCollectionDataInner) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *GetSegmentResponseCollectionDataInner) GetAttributes() SegmentListResponseObjectResourceAttributes {
	if o == nil {
		var ret SegmentListResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetSegmentResponseCollectionDataInner) GetAttributesOk() (*SegmentListResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GetSegmentResponseCollectionDataInner) SetAttributes(v SegmentListResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *GetSegmentResponseCollectionDataInner) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetSegmentResponseCollectionDataInner) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetSegmentResponseCollectionDataInner) SetLinks(v ObjectLinks) {
	o.Links = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GetSegmentResponseCollectionDataInner) GetRelationships() GetListResponseCollectionDataInnerAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetListResponseCollectionDataInnerAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSegmentResponseCollectionDataInner) GetRelationshipsOk() (*GetListResponseCollectionDataInnerAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GetSegmentResponseCollectionDataInner) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetListResponseCollectionDataInnerAllOfRelationships and assigns it to the Relationships field.
func (o *GetSegmentResponseCollectionDataInner) SetRelationships(v GetListResponseCollectionDataInnerAllOfRelationships) {
	o.Relationships = &v
}

func (o GetSegmentResponseCollectionDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSegmentResponseCollectionDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

func (o *GetSegmentResponseCollectionDataInner) UnmarshalJSON(data []byte) (err error) {
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

	varGetSegmentResponseCollectionDataInner := _GetSegmentResponseCollectionDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSegmentResponseCollectionDataInner)

	if err != nil {
		return err
	}

	*o = GetSegmentResponseCollectionDataInner(varGetSegmentResponseCollectionDataInner)

	return err
}

type NullableGetSegmentResponseCollectionDataInner struct {
	value *GetSegmentResponseCollectionDataInner
	isSet bool
}

func (v NullableGetSegmentResponseCollectionDataInner) Get() *GetSegmentResponseCollectionDataInner {
	return v.value
}

func (v *NullableGetSegmentResponseCollectionDataInner) Set(val *GetSegmentResponseCollectionDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSegmentResponseCollectionDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSegmentResponseCollectionDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSegmentResponseCollectionDataInner(val *GetSegmentResponseCollectionDataInner) *NullableGetSegmentResponseCollectionDataInner {
	return &NullableGetSegmentResponseCollectionDataInner{value: val, isSet: true}
}

func (v NullableGetSegmentResponseCollectionDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSegmentResponseCollectionDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

