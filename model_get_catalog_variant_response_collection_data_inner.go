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

// checks if the GetCatalogVariantResponseCollectionDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCatalogVariantResponseCollectionDataInner{}

// GetCatalogVariantResponseCollectionDataInner struct for GetCatalogVariantResponseCollectionDataInner
type GetCatalogVariantResponseCollectionDataInner struct {
	Type CatalogVariantEnum `json:"type"`
	// The catalog variant ID is a compound ID (string), with format: `{integration}:::{catalog}:::{external_id}`. Currently, the only supported integration type is `$custom`, and the only supported catalog is `$default`.
	Id string `json:"id"`
	Attributes CatalogVariantResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
	Relationships *GetCatalogVariantResponseCollectionDataInnerAllOfRelationships `json:"relationships,omitempty"`
}

type _GetCatalogVariantResponseCollectionDataInner GetCatalogVariantResponseCollectionDataInner

// NewGetCatalogVariantResponseCollectionDataInner instantiates a new GetCatalogVariantResponseCollectionDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogVariantResponseCollectionDataInner(type_ CatalogVariantEnum, id string, attributes CatalogVariantResponseObjectResourceAttributes, links ObjectLinks) *GetCatalogVariantResponseCollectionDataInner {
	this := GetCatalogVariantResponseCollectionDataInner{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewGetCatalogVariantResponseCollectionDataInnerWithDefaults instantiates a new GetCatalogVariantResponseCollectionDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogVariantResponseCollectionDataInnerWithDefaults() *GetCatalogVariantResponseCollectionDataInner {
	this := GetCatalogVariantResponseCollectionDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GetCatalogVariantResponseCollectionDataInner) GetType() CatalogVariantEnum {
	if o == nil {
		var ret CatalogVariantEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetCatalogVariantResponseCollectionDataInner) GetTypeOk() (*CatalogVariantEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetCatalogVariantResponseCollectionDataInner) SetType(v CatalogVariantEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetCatalogVariantResponseCollectionDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetCatalogVariantResponseCollectionDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetCatalogVariantResponseCollectionDataInner) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *GetCatalogVariantResponseCollectionDataInner) GetAttributes() CatalogVariantResponseObjectResourceAttributes {
	if o == nil {
		var ret CatalogVariantResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetCatalogVariantResponseCollectionDataInner) GetAttributesOk() (*CatalogVariantResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GetCatalogVariantResponseCollectionDataInner) SetAttributes(v CatalogVariantResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *GetCatalogVariantResponseCollectionDataInner) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetCatalogVariantResponseCollectionDataInner) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetCatalogVariantResponseCollectionDataInner) SetLinks(v ObjectLinks) {
	o.Links = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GetCatalogVariantResponseCollectionDataInner) GetRelationships() GetCatalogVariantResponseCollectionDataInnerAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetCatalogVariantResponseCollectionDataInnerAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCatalogVariantResponseCollectionDataInner) GetRelationshipsOk() (*GetCatalogVariantResponseCollectionDataInnerAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GetCatalogVariantResponseCollectionDataInner) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetCatalogVariantResponseCollectionDataInnerAllOfRelationships and assigns it to the Relationships field.
func (o *GetCatalogVariantResponseCollectionDataInner) SetRelationships(v GetCatalogVariantResponseCollectionDataInnerAllOfRelationships) {
	o.Relationships = &v
}

func (o GetCatalogVariantResponseCollectionDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCatalogVariantResponseCollectionDataInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetCatalogVariantResponseCollectionDataInner) UnmarshalJSON(data []byte) (err error) {
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

	varGetCatalogVariantResponseCollectionDataInner := _GetCatalogVariantResponseCollectionDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCatalogVariantResponseCollectionDataInner)

	if err != nil {
		return err
	}

	*o = GetCatalogVariantResponseCollectionDataInner(varGetCatalogVariantResponseCollectionDataInner)

	return err
}

type NullableGetCatalogVariantResponseCollectionDataInner struct {
	value *GetCatalogVariantResponseCollectionDataInner
	isSet bool
}

func (v NullableGetCatalogVariantResponseCollectionDataInner) Get() *GetCatalogVariantResponseCollectionDataInner {
	return v.value
}

func (v *NullableGetCatalogVariantResponseCollectionDataInner) Set(val *GetCatalogVariantResponseCollectionDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogVariantResponseCollectionDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogVariantResponseCollectionDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogVariantResponseCollectionDataInner(val *GetCatalogVariantResponseCollectionDataInner) *NullableGetCatalogVariantResponseCollectionDataInner {
	return &NullableGetCatalogVariantResponseCollectionDataInner{value: val, isSet: true}
}

func (v NullableGetCatalogVariantResponseCollectionDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogVariantResponseCollectionDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

