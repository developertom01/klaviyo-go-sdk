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
)

// checks if the GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships{}

// GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships struct for GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships
type GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships struct {
	Categories *GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategories `json:"categories,omitempty"`
}

// NewGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships instantiates a new GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships() *GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	this := GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships{}
	return &this
}

// NewGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsWithDefaults instantiates a new GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsWithDefaults() *GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	this := GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetCategories() GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategories {
	if o == nil || IsNil(o.Categories) {
		var ret GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategories
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetCategoriesOk() (*GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategories, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategories and assigns it to the Categories field.
func (o *GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) SetCategories(v GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategories) {
	o.Categories = &v
}

func (o GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

type NullableGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships struct {
	value *GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships
	isSet bool
}

func (v NullableGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) Get() *GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	return v.value
}

func (v *NullableGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) Set(val *GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships(val *GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) *NullableGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	return &NullableGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships{value: val, isSet: true}
}

func (v NullableGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


