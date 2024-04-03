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

// checks if the PostCouponCodeCreateJobResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCouponCodeCreateJobResponseData{}

// PostCouponCodeCreateJobResponseData struct for PostCouponCodeCreateJobResponseData
type PostCouponCodeCreateJobResponseData struct {
	Type CouponCodeBulkCreateJobEnum `json:"type"`
	// Unique identifier for retrieving the job. Generated by Klaviyo.
	Id string `json:"id"`
	Attributes CouponCodeCreateJobResponseObjectResourceAttributes `json:"attributes"`
	Relationships *GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships `json:"relationships,omitempty"`
	Links ObjectLinks `json:"links"`
}

type _PostCouponCodeCreateJobResponseData PostCouponCodeCreateJobResponseData

// NewPostCouponCodeCreateJobResponseData instantiates a new PostCouponCodeCreateJobResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCouponCodeCreateJobResponseData(type_ CouponCodeBulkCreateJobEnum, id string, attributes CouponCodeCreateJobResponseObjectResourceAttributes, links ObjectLinks) *PostCouponCodeCreateJobResponseData {
	this := PostCouponCodeCreateJobResponseData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewPostCouponCodeCreateJobResponseDataWithDefaults instantiates a new PostCouponCodeCreateJobResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCouponCodeCreateJobResponseDataWithDefaults() *PostCouponCodeCreateJobResponseData {
	this := PostCouponCodeCreateJobResponseData{}
	return &this
}

// GetType returns the Type field value
func (o *PostCouponCodeCreateJobResponseData) GetType() CouponCodeBulkCreateJobEnum {
	if o == nil {
		var ret CouponCodeBulkCreateJobEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostCouponCodeCreateJobResponseData) GetTypeOk() (*CouponCodeBulkCreateJobEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostCouponCodeCreateJobResponseData) SetType(v CouponCodeBulkCreateJobEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PostCouponCodeCreateJobResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostCouponCodeCreateJobResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostCouponCodeCreateJobResponseData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PostCouponCodeCreateJobResponseData) GetAttributes() CouponCodeCreateJobResponseObjectResourceAttributes {
	if o == nil {
		var ret CouponCodeCreateJobResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PostCouponCodeCreateJobResponseData) GetAttributesOk() (*CouponCodeCreateJobResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PostCouponCodeCreateJobResponseData) SetAttributes(v CouponCodeCreateJobResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PostCouponCodeCreateJobResponseData) GetRelationships() GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCouponCodeCreateJobResponseData) GetRelationshipsOk() (*GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PostCouponCodeCreateJobResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships and assigns it to the Relationships field.
func (o *PostCouponCodeCreateJobResponseData) SetRelationships(v GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *PostCouponCodeCreateJobResponseData) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PostCouponCodeCreateJobResponseData) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PostCouponCodeCreateJobResponseData) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o PostCouponCodeCreateJobResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCouponCodeCreateJobResponseData) ToMap() (map[string]interface{}, error) {
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

func (o *PostCouponCodeCreateJobResponseData) UnmarshalJSON(data []byte) (err error) {
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

	varPostCouponCodeCreateJobResponseData := _PostCouponCodeCreateJobResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCouponCodeCreateJobResponseData)

	if err != nil {
		return err
	}

	*o = PostCouponCodeCreateJobResponseData(varPostCouponCodeCreateJobResponseData)

	return err
}

type NullablePostCouponCodeCreateJobResponseData struct {
	value *PostCouponCodeCreateJobResponseData
	isSet bool
}

func (v NullablePostCouponCodeCreateJobResponseData) Get() *PostCouponCodeCreateJobResponseData {
	return v.value
}

func (v *NullablePostCouponCodeCreateJobResponseData) Set(val *PostCouponCodeCreateJobResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCouponCodeCreateJobResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCouponCodeCreateJobResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCouponCodeCreateJobResponseData(val *PostCouponCodeCreateJobResponseData) *NullablePostCouponCodeCreateJobResponseData {
	return &NullablePostCouponCodeCreateJobResponseData{value: val, isSet: true}
}

func (v NullablePostCouponCodeCreateJobResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCouponCodeCreateJobResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


