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

// checks if the GetSegmentListResponseCollectionCompoundDocumentDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSegmentListResponseCollectionCompoundDocumentDataInner{}

// GetSegmentListResponseCollectionCompoundDocumentDataInner struct for GetSegmentListResponseCollectionCompoundDocumentDataInner
type GetSegmentListResponseCollectionCompoundDocumentDataInner struct {
	Type SegmentEnum `json:"type"`
	Id string `json:"id"`
	Attributes SegmentListResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
	Relationships *GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationships `json:"relationships,omitempty"`
}

type _GetSegmentListResponseCollectionCompoundDocumentDataInner GetSegmentListResponseCollectionCompoundDocumentDataInner

// NewGetSegmentListResponseCollectionCompoundDocumentDataInner instantiates a new GetSegmentListResponseCollectionCompoundDocumentDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSegmentListResponseCollectionCompoundDocumentDataInner(type_ SegmentEnum, id string, attributes SegmentListResponseObjectResourceAttributes, links ObjectLinks) *GetSegmentListResponseCollectionCompoundDocumentDataInner {
	this := GetSegmentListResponseCollectionCompoundDocumentDataInner{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewGetSegmentListResponseCollectionCompoundDocumentDataInnerWithDefaults instantiates a new GetSegmentListResponseCollectionCompoundDocumentDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSegmentListResponseCollectionCompoundDocumentDataInnerWithDefaults() *GetSegmentListResponseCollectionCompoundDocumentDataInner {
	this := GetSegmentListResponseCollectionCompoundDocumentDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) GetType() SegmentEnum {
	if o == nil {
		var ret SegmentEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) GetTypeOk() (*SegmentEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) SetType(v SegmentEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) GetAttributes() SegmentListResponseObjectResourceAttributes {
	if o == nil {
		var ret SegmentListResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) GetAttributesOk() (*SegmentListResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) SetAttributes(v SegmentListResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) SetLinks(v ObjectLinks) {
	o.Links = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) GetRelationships() GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) GetRelationshipsOk() (*GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationships and assigns it to the Relationships field.
func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) SetRelationships(v GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationships) {
	o.Relationships = &v
}

func (o GetSegmentListResponseCollectionCompoundDocumentDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSegmentListResponseCollectionCompoundDocumentDataInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetSegmentListResponseCollectionCompoundDocumentDataInner) UnmarshalJSON(data []byte) (err error) {
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

	varGetSegmentListResponseCollectionCompoundDocumentDataInner := _GetSegmentListResponseCollectionCompoundDocumentDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSegmentListResponseCollectionCompoundDocumentDataInner)

	if err != nil {
		return err
	}

	*o = GetSegmentListResponseCollectionCompoundDocumentDataInner(varGetSegmentListResponseCollectionCompoundDocumentDataInner)

	return err
}

type NullableGetSegmentListResponseCollectionCompoundDocumentDataInner struct {
	value *GetSegmentListResponseCollectionCompoundDocumentDataInner
	isSet bool
}

func (v NullableGetSegmentListResponseCollectionCompoundDocumentDataInner) Get() *GetSegmentListResponseCollectionCompoundDocumentDataInner {
	return v.value
}

func (v *NullableGetSegmentListResponseCollectionCompoundDocumentDataInner) Set(val *GetSegmentListResponseCollectionCompoundDocumentDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSegmentListResponseCollectionCompoundDocumentDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSegmentListResponseCollectionCompoundDocumentDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSegmentListResponseCollectionCompoundDocumentDataInner(val *GetSegmentListResponseCollectionCompoundDocumentDataInner) *NullableGetSegmentListResponseCollectionCompoundDocumentDataInner {
	return &NullableGetSegmentListResponseCollectionCompoundDocumentDataInner{value: val, isSet: true}
}

func (v NullableGetSegmentListResponseCollectionCompoundDocumentDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSegmentListResponseCollectionCompoundDocumentDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

