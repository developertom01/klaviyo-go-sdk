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

// checks if the GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner{}

// GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner struct for GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner
type GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner struct {
	Type CatalogItemBulkCreateJobEnum `json:"type"`
	// Unique identifier for retrieving the job. Generated by Klaviyo.
	Id string `json:"id"`
	Attributes CouponCodeCreateJobResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
	Relationships *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships `json:"relationships,omitempty"`
}

type _GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner

// NewGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner instantiates a new GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner(type_ CatalogItemBulkCreateJobEnum, id string, attributes CouponCodeCreateJobResponseObjectResourceAttributes, links ObjectLinks) *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner {
	this := GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerWithDefaults instantiates a new GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerWithDefaults() *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner {
	this := GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) GetType() CatalogItemBulkCreateJobEnum {
	if o == nil {
		var ret CatalogItemBulkCreateJobEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) GetTypeOk() (*CatalogItemBulkCreateJobEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) SetType(v CatalogItemBulkCreateJobEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) GetAttributes() CouponCodeCreateJobResponseObjectResourceAttributes {
	if o == nil {
		var ret CouponCodeCreateJobResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) GetAttributesOk() (*CouponCodeCreateJobResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) SetAttributes(v CouponCodeCreateJobResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) SetLinks(v ObjectLinks) {
	o.Links = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) GetRelationships() GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) GetRelationshipsOk() (*GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships and assigns it to the Relationships field.
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) SetRelationships(v GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) {
	o.Relationships = &v
}

func (o GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) UnmarshalJSON(data []byte) (err error) {
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

	varGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner := _GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner)

	if err != nil {
		return err
	}

	*o = GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner(varGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner)

	return err
}

type NullableGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner struct {
	value *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner
	isSet bool
}

func (v NullableGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) Get() *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner {
	return v.value
}

func (v *NullableGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) Set(val *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner(val *GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) *NullableGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner {
	return &NullableGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner{value: val, isSet: true}
}

func (v NullableGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


