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

// checks if the GetTagResponseCollectionCompoundDocumentDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTagResponseCollectionCompoundDocumentDataInner{}

// GetTagResponseCollectionCompoundDocumentDataInner struct for GetTagResponseCollectionCompoundDocumentDataInner
type GetTagResponseCollectionCompoundDocumentDataInner struct {
	Type TagEnum `json:"type"`
	// The Tag ID
	Id string `json:"id"`
	Attributes TagResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
	Relationships *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationships `json:"relationships,omitempty"`
}

type _GetTagResponseCollectionCompoundDocumentDataInner GetTagResponseCollectionCompoundDocumentDataInner

// NewGetTagResponseCollectionCompoundDocumentDataInner instantiates a new GetTagResponseCollectionCompoundDocumentDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTagResponseCollectionCompoundDocumentDataInner(type_ TagEnum, id string, attributes TagResponseObjectResourceAttributes, links ObjectLinks) *GetTagResponseCollectionCompoundDocumentDataInner {
	this := GetTagResponseCollectionCompoundDocumentDataInner{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewGetTagResponseCollectionCompoundDocumentDataInnerWithDefaults instantiates a new GetTagResponseCollectionCompoundDocumentDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTagResponseCollectionCompoundDocumentDataInnerWithDefaults() *GetTagResponseCollectionCompoundDocumentDataInner {
	this := GetTagResponseCollectionCompoundDocumentDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GetTagResponseCollectionCompoundDocumentDataInner) GetType() TagEnum {
	if o == nil {
		var ret TagEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetTagResponseCollectionCompoundDocumentDataInner) GetTypeOk() (*TagEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetTagResponseCollectionCompoundDocumentDataInner) SetType(v TagEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetTagResponseCollectionCompoundDocumentDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetTagResponseCollectionCompoundDocumentDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetTagResponseCollectionCompoundDocumentDataInner) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *GetTagResponseCollectionCompoundDocumentDataInner) GetAttributes() TagResponseObjectResourceAttributes {
	if o == nil {
		var ret TagResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetTagResponseCollectionCompoundDocumentDataInner) GetAttributesOk() (*TagResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GetTagResponseCollectionCompoundDocumentDataInner) SetAttributes(v TagResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *GetTagResponseCollectionCompoundDocumentDataInner) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetTagResponseCollectionCompoundDocumentDataInner) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetTagResponseCollectionCompoundDocumentDataInner) SetLinks(v ObjectLinks) {
	o.Links = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GetTagResponseCollectionCompoundDocumentDataInner) GetRelationships() GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTagResponseCollectionCompoundDocumentDataInner) GetRelationshipsOk() (*GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GetTagResponseCollectionCompoundDocumentDataInner) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationships and assigns it to the Relationships field.
func (o *GetTagResponseCollectionCompoundDocumentDataInner) SetRelationships(v GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationships) {
	o.Relationships = &v
}

func (o GetTagResponseCollectionCompoundDocumentDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTagResponseCollectionCompoundDocumentDataInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetTagResponseCollectionCompoundDocumentDataInner) UnmarshalJSON(data []byte) (err error) {
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

	varGetTagResponseCollectionCompoundDocumentDataInner := _GetTagResponseCollectionCompoundDocumentDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTagResponseCollectionCompoundDocumentDataInner)

	if err != nil {
		return err
	}

	*o = GetTagResponseCollectionCompoundDocumentDataInner(varGetTagResponseCollectionCompoundDocumentDataInner)

	return err
}

type NullableGetTagResponseCollectionCompoundDocumentDataInner struct {
	value *GetTagResponseCollectionCompoundDocumentDataInner
	isSet bool
}

func (v NullableGetTagResponseCollectionCompoundDocumentDataInner) Get() *GetTagResponseCollectionCompoundDocumentDataInner {
	return v.value
}

func (v *NullableGetTagResponseCollectionCompoundDocumentDataInner) Set(val *GetTagResponseCollectionCompoundDocumentDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTagResponseCollectionCompoundDocumentDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTagResponseCollectionCompoundDocumentDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTagResponseCollectionCompoundDocumentDataInner(val *GetTagResponseCollectionCompoundDocumentDataInner) *NullableGetTagResponseCollectionCompoundDocumentDataInner {
	return &NullableGetTagResponseCollectionCompoundDocumentDataInner{value: val, isSet: true}
}

func (v NullableGetTagResponseCollectionCompoundDocumentDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTagResponseCollectionCompoundDocumentDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


