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

// checks if the PostCatalogCategoryUpdateJobResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCatalogCategoryUpdateJobResponseData{}

// PostCatalogCategoryUpdateJobResponseData struct for PostCatalogCategoryUpdateJobResponseData
type PostCatalogCategoryUpdateJobResponseData struct {
	Type CatalogCategoryBulkUpdateJobEnum `json:"type"`
	// Unique identifier for retrieving the job. Generated by Klaviyo.
	Id string `json:"id"`
	Attributes CouponCodeCreateJobResponseObjectResourceAttributes `json:"attributes"`
	Relationships *GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships `json:"relationships,omitempty"`
	Links ObjectLinks `json:"links"`
}

type _PostCatalogCategoryUpdateJobResponseData PostCatalogCategoryUpdateJobResponseData

// NewPostCatalogCategoryUpdateJobResponseData instantiates a new PostCatalogCategoryUpdateJobResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCatalogCategoryUpdateJobResponseData(type_ CatalogCategoryBulkUpdateJobEnum, id string, attributes CouponCodeCreateJobResponseObjectResourceAttributes, links ObjectLinks) *PostCatalogCategoryUpdateJobResponseData {
	this := PostCatalogCategoryUpdateJobResponseData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewPostCatalogCategoryUpdateJobResponseDataWithDefaults instantiates a new PostCatalogCategoryUpdateJobResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCatalogCategoryUpdateJobResponseDataWithDefaults() *PostCatalogCategoryUpdateJobResponseData {
	this := PostCatalogCategoryUpdateJobResponseData{}
	return &this
}

// GetType returns the Type field value
func (o *PostCatalogCategoryUpdateJobResponseData) GetType() CatalogCategoryBulkUpdateJobEnum {
	if o == nil {
		var ret CatalogCategoryBulkUpdateJobEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostCatalogCategoryUpdateJobResponseData) GetTypeOk() (*CatalogCategoryBulkUpdateJobEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostCatalogCategoryUpdateJobResponseData) SetType(v CatalogCategoryBulkUpdateJobEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PostCatalogCategoryUpdateJobResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostCatalogCategoryUpdateJobResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostCatalogCategoryUpdateJobResponseData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PostCatalogCategoryUpdateJobResponseData) GetAttributes() CouponCodeCreateJobResponseObjectResourceAttributes {
	if o == nil {
		var ret CouponCodeCreateJobResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PostCatalogCategoryUpdateJobResponseData) GetAttributesOk() (*CouponCodeCreateJobResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PostCatalogCategoryUpdateJobResponseData) SetAttributes(v CouponCodeCreateJobResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PostCatalogCategoryUpdateJobResponseData) GetRelationships() GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCatalogCategoryUpdateJobResponseData) GetRelationshipsOk() (*GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PostCatalogCategoryUpdateJobResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships and assigns it to the Relationships field.
func (o *PostCatalogCategoryUpdateJobResponseData) SetRelationships(v GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *PostCatalogCategoryUpdateJobResponseData) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PostCatalogCategoryUpdateJobResponseData) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PostCatalogCategoryUpdateJobResponseData) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o PostCatalogCategoryUpdateJobResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCatalogCategoryUpdateJobResponseData) ToMap() (map[string]interface{}, error) {
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

func (o *PostCatalogCategoryUpdateJobResponseData) UnmarshalJSON(data []byte) (err error) {
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

	varPostCatalogCategoryUpdateJobResponseData := _PostCatalogCategoryUpdateJobResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCatalogCategoryUpdateJobResponseData)

	if err != nil {
		return err
	}

	*o = PostCatalogCategoryUpdateJobResponseData(varPostCatalogCategoryUpdateJobResponseData)

	return err
}

type NullablePostCatalogCategoryUpdateJobResponseData struct {
	value *PostCatalogCategoryUpdateJobResponseData
	isSet bool
}

func (v NullablePostCatalogCategoryUpdateJobResponseData) Get() *PostCatalogCategoryUpdateJobResponseData {
	return v.value
}

func (v *NullablePostCatalogCategoryUpdateJobResponseData) Set(val *PostCatalogCategoryUpdateJobResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCatalogCategoryUpdateJobResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCatalogCategoryUpdateJobResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCatalogCategoryUpdateJobResponseData(val *PostCatalogCategoryUpdateJobResponseData) *NullablePostCatalogCategoryUpdateJobResponseData {
	return &NullablePostCatalogCategoryUpdateJobResponseData{value: val, isSet: true}
}

func (v NullablePostCatalogCategoryUpdateJobResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCatalogCategoryUpdateJobResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

