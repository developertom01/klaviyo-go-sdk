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

// checks if the GetCatalogCategoryResponseCollectionDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCatalogCategoryResponseCollectionDataInner{}

// GetCatalogCategoryResponseCollectionDataInner struct for GetCatalogCategoryResponseCollectionDataInner
type GetCatalogCategoryResponseCollectionDataInner struct {
	Type CatalogCategoryEnum `json:"type"`
	// The catalog category ID is a compound ID (string), with format: `{integration}:::{catalog}:::{external_id}`. Currently, the only supported integration type is `$custom`, and the only supported catalog is `$default`.
	Id string `json:"id"`
	Attributes CatalogCategoryResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
	Relationships *GetCatalogCategoryResponseCollectionDataInnerAllOfRelationships `json:"relationships,omitempty"`
}

type _GetCatalogCategoryResponseCollectionDataInner GetCatalogCategoryResponseCollectionDataInner

// NewGetCatalogCategoryResponseCollectionDataInner instantiates a new GetCatalogCategoryResponseCollectionDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogCategoryResponseCollectionDataInner(type_ CatalogCategoryEnum, id string, attributes CatalogCategoryResponseObjectResourceAttributes, links ObjectLinks) *GetCatalogCategoryResponseCollectionDataInner {
	this := GetCatalogCategoryResponseCollectionDataInner{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewGetCatalogCategoryResponseCollectionDataInnerWithDefaults instantiates a new GetCatalogCategoryResponseCollectionDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogCategoryResponseCollectionDataInnerWithDefaults() *GetCatalogCategoryResponseCollectionDataInner {
	this := GetCatalogCategoryResponseCollectionDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GetCatalogCategoryResponseCollectionDataInner) GetType() CatalogCategoryEnum {
	if o == nil {
		var ret CatalogCategoryEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetCatalogCategoryResponseCollectionDataInner) GetTypeOk() (*CatalogCategoryEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetCatalogCategoryResponseCollectionDataInner) SetType(v CatalogCategoryEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetCatalogCategoryResponseCollectionDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetCatalogCategoryResponseCollectionDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetCatalogCategoryResponseCollectionDataInner) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *GetCatalogCategoryResponseCollectionDataInner) GetAttributes() CatalogCategoryResponseObjectResourceAttributes {
	if o == nil {
		var ret CatalogCategoryResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetCatalogCategoryResponseCollectionDataInner) GetAttributesOk() (*CatalogCategoryResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GetCatalogCategoryResponseCollectionDataInner) SetAttributes(v CatalogCategoryResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *GetCatalogCategoryResponseCollectionDataInner) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetCatalogCategoryResponseCollectionDataInner) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetCatalogCategoryResponseCollectionDataInner) SetLinks(v ObjectLinks) {
	o.Links = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GetCatalogCategoryResponseCollectionDataInner) GetRelationships() GetCatalogCategoryResponseCollectionDataInnerAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetCatalogCategoryResponseCollectionDataInnerAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCatalogCategoryResponseCollectionDataInner) GetRelationshipsOk() (*GetCatalogCategoryResponseCollectionDataInnerAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GetCatalogCategoryResponseCollectionDataInner) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetCatalogCategoryResponseCollectionDataInnerAllOfRelationships and assigns it to the Relationships field.
func (o *GetCatalogCategoryResponseCollectionDataInner) SetRelationships(v GetCatalogCategoryResponseCollectionDataInnerAllOfRelationships) {
	o.Relationships = &v
}

func (o GetCatalogCategoryResponseCollectionDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCatalogCategoryResponseCollectionDataInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetCatalogCategoryResponseCollectionDataInner) UnmarshalJSON(data []byte) (err error) {
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

	varGetCatalogCategoryResponseCollectionDataInner := _GetCatalogCategoryResponseCollectionDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCatalogCategoryResponseCollectionDataInner)

	if err != nil {
		return err
	}

	*o = GetCatalogCategoryResponseCollectionDataInner(varGetCatalogCategoryResponseCollectionDataInner)

	return err
}

type NullableGetCatalogCategoryResponseCollectionDataInner struct {
	value *GetCatalogCategoryResponseCollectionDataInner
	isSet bool
}

func (v NullableGetCatalogCategoryResponseCollectionDataInner) Get() *GetCatalogCategoryResponseCollectionDataInner {
	return v.value
}

func (v *NullableGetCatalogCategoryResponseCollectionDataInner) Set(val *GetCatalogCategoryResponseCollectionDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogCategoryResponseCollectionDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogCategoryResponseCollectionDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogCategoryResponseCollectionDataInner(val *GetCatalogCategoryResponseCollectionDataInner) *NullableGetCatalogCategoryResponseCollectionDataInner {
	return &NullableGetCatalogCategoryResponseCollectionDataInner{value: val, isSet: true}
}

func (v NullableGetCatalogCategoryResponseCollectionDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogCategoryResponseCollectionDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


