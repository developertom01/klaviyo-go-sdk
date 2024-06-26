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

// checks if the PostCatalogCategoryResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCatalogCategoryResponseDataRelationships{}

// PostCatalogCategoryResponseDataRelationships struct for PostCatalogCategoryResponseDataRelationships
type PostCatalogCategoryResponseDataRelationships struct {
	Items *PostCatalogCategoryResponseDataRelationshipsItems `json:"items,omitempty"`
}

// NewPostCatalogCategoryResponseDataRelationships instantiates a new PostCatalogCategoryResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCatalogCategoryResponseDataRelationships() *PostCatalogCategoryResponseDataRelationships {
	this := PostCatalogCategoryResponseDataRelationships{}
	return &this
}

// NewPostCatalogCategoryResponseDataRelationshipsWithDefaults instantiates a new PostCatalogCategoryResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCatalogCategoryResponseDataRelationshipsWithDefaults() *PostCatalogCategoryResponseDataRelationships {
	this := PostCatalogCategoryResponseDataRelationships{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PostCatalogCategoryResponseDataRelationships) GetItems() PostCatalogCategoryResponseDataRelationshipsItems {
	if o == nil || IsNil(o.Items) {
		var ret PostCatalogCategoryResponseDataRelationshipsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCatalogCategoryResponseDataRelationships) GetItemsOk() (*PostCatalogCategoryResponseDataRelationshipsItems, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PostCatalogCategoryResponseDataRelationships) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given PostCatalogCategoryResponseDataRelationshipsItems and assigns it to the Items field.
func (o *PostCatalogCategoryResponseDataRelationships) SetItems(v PostCatalogCategoryResponseDataRelationshipsItems) {
	o.Items = &v
}

func (o PostCatalogCategoryResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCatalogCategoryResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullablePostCatalogCategoryResponseDataRelationships struct {
	value *PostCatalogCategoryResponseDataRelationships
	isSet bool
}

func (v NullablePostCatalogCategoryResponseDataRelationships) Get() *PostCatalogCategoryResponseDataRelationships {
	return v.value
}

func (v *NullablePostCatalogCategoryResponseDataRelationships) Set(val *PostCatalogCategoryResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCatalogCategoryResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCatalogCategoryResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCatalogCategoryResponseDataRelationships(val *PostCatalogCategoryResponseDataRelationships) *NullablePostCatalogCategoryResponseDataRelationships {
	return &NullablePostCatalogCategoryResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePostCatalogCategoryResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCatalogCategoryResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


