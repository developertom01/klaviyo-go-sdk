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

// checks if the PostCatalogItemUpdateJobResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCatalogItemUpdateJobResponseData{}

// PostCatalogItemUpdateJobResponseData struct for PostCatalogItemUpdateJobResponseData
type PostCatalogItemUpdateJobResponseData struct {
	Type CatalogItemBulkUpdateJobEnum `json:"type"`
	// Unique identifier for retrieving the job. Generated by Klaviyo.
	Id string `json:"id"`
	Attributes CouponCodeCreateJobResponseObjectResourceAttributes `json:"attributes"`
	Relationships *GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships `json:"relationships,omitempty"`
	Links ObjectLinks `json:"links"`
}

type _PostCatalogItemUpdateJobResponseData PostCatalogItemUpdateJobResponseData

// NewPostCatalogItemUpdateJobResponseData instantiates a new PostCatalogItemUpdateJobResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCatalogItemUpdateJobResponseData(type_ CatalogItemBulkUpdateJobEnum, id string, attributes CouponCodeCreateJobResponseObjectResourceAttributes, links ObjectLinks) *PostCatalogItemUpdateJobResponseData {
	this := PostCatalogItemUpdateJobResponseData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewPostCatalogItemUpdateJobResponseDataWithDefaults instantiates a new PostCatalogItemUpdateJobResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCatalogItemUpdateJobResponseDataWithDefaults() *PostCatalogItemUpdateJobResponseData {
	this := PostCatalogItemUpdateJobResponseData{}
	return &this
}

// GetType returns the Type field value
func (o *PostCatalogItemUpdateJobResponseData) GetType() CatalogItemBulkUpdateJobEnum {
	if o == nil {
		var ret CatalogItemBulkUpdateJobEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostCatalogItemUpdateJobResponseData) GetTypeOk() (*CatalogItemBulkUpdateJobEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostCatalogItemUpdateJobResponseData) SetType(v CatalogItemBulkUpdateJobEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PostCatalogItemUpdateJobResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostCatalogItemUpdateJobResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostCatalogItemUpdateJobResponseData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PostCatalogItemUpdateJobResponseData) GetAttributes() CouponCodeCreateJobResponseObjectResourceAttributes {
	if o == nil {
		var ret CouponCodeCreateJobResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PostCatalogItemUpdateJobResponseData) GetAttributesOk() (*CouponCodeCreateJobResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PostCatalogItemUpdateJobResponseData) SetAttributes(v CouponCodeCreateJobResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PostCatalogItemUpdateJobResponseData) GetRelationships() GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCatalogItemUpdateJobResponseData) GetRelationshipsOk() (*GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PostCatalogItemUpdateJobResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships and assigns it to the Relationships field.
func (o *PostCatalogItemUpdateJobResponseData) SetRelationships(v GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *PostCatalogItemUpdateJobResponseData) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PostCatalogItemUpdateJobResponseData) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PostCatalogItemUpdateJobResponseData) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o PostCatalogItemUpdateJobResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCatalogItemUpdateJobResponseData) ToMap() (map[string]interface{}, error) {
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

func (o *PostCatalogItemUpdateJobResponseData) UnmarshalJSON(data []byte) (err error) {
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

	varPostCatalogItemUpdateJobResponseData := _PostCatalogItemUpdateJobResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCatalogItemUpdateJobResponseData)

	if err != nil {
		return err
	}

	*o = PostCatalogItemUpdateJobResponseData(varPostCatalogItemUpdateJobResponseData)

	return err
}

type NullablePostCatalogItemUpdateJobResponseData struct {
	value *PostCatalogItemUpdateJobResponseData
	isSet bool
}

func (v NullablePostCatalogItemUpdateJobResponseData) Get() *PostCatalogItemUpdateJobResponseData {
	return v.value
}

func (v *NullablePostCatalogItemUpdateJobResponseData) Set(val *PostCatalogItemUpdateJobResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCatalogItemUpdateJobResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCatalogItemUpdateJobResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCatalogItemUpdateJobResponseData(val *PostCatalogItemUpdateJobResponseData) *NullablePostCatalogItemUpdateJobResponseData {
	return &NullablePostCatalogItemUpdateJobResponseData{value: val, isSet: true}
}

func (v NullablePostCatalogItemUpdateJobResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCatalogItemUpdateJobResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


