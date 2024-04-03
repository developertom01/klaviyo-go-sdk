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

// checks if the PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories{}

// PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories struct for PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories
type PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories struct {
	Data []PostCatalogCategoryDeleteJobResponseDataRelationshipsCategoriesDataInner `json:"data"`
	Links *RelationshipLinks `json:"links,omitempty"`
}

type _PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories

// NewPostCatalogCategoryDeleteJobResponseDataRelationshipsCategories instantiates a new PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCatalogCategoryDeleteJobResponseDataRelationshipsCategories(data []PostCatalogCategoryDeleteJobResponseDataRelationshipsCategoriesDataInner) *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories {
	this := PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories{}
	this.Data = data
	return &this
}

// NewPostCatalogCategoryDeleteJobResponseDataRelationshipsCategoriesWithDefaults instantiates a new PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCatalogCategoryDeleteJobResponseDataRelationshipsCategoriesWithDefaults() *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories {
	this := PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories{}
	return &this
}

// GetData returns the Data field value
func (o *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) GetData() []PostCatalogCategoryDeleteJobResponseDataRelationshipsCategoriesDataInner {
	if o == nil {
		var ret []PostCatalogCategoryDeleteJobResponseDataRelationshipsCategoriesDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) GetDataOk() ([]PostCatalogCategoryDeleteJobResponseDataRelationshipsCategoriesDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) SetData(v []PostCatalogCategoryDeleteJobResponseDataRelationshipsCategoriesDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

func (o PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varPostCatalogCategoryDeleteJobResponseDataRelationshipsCategories := _PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCatalogCategoryDeleteJobResponseDataRelationshipsCategories)

	if err != nil {
		return err
	}

	*o = PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories(varPostCatalogCategoryDeleteJobResponseDataRelationshipsCategories)

	return err
}

type NullablePostCatalogCategoryDeleteJobResponseDataRelationshipsCategories struct {
	value *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories
	isSet bool
}

func (v NullablePostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) Get() *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories {
	return v.value
}

func (v *NullablePostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) Set(val *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCatalogCategoryDeleteJobResponseDataRelationshipsCategories(val *PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) *NullablePostCatalogCategoryDeleteJobResponseDataRelationshipsCategories {
	return &NullablePostCatalogCategoryDeleteJobResponseDataRelationshipsCategories{value: val, isSet: true}
}

func (v NullablePostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCatalogCategoryDeleteJobResponseDataRelationshipsCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


