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

// checks if the PostCatalogItemDeleteJobResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCatalogItemDeleteJobResponseDataRelationships{}

// PostCatalogItemDeleteJobResponseDataRelationships struct for PostCatalogItemDeleteJobResponseDataRelationships
type PostCatalogItemDeleteJobResponseDataRelationships struct {
	Items *PostCatalogItemDeleteJobResponseDataRelationshipsItems `json:"items,omitempty"`
}

// NewPostCatalogItemDeleteJobResponseDataRelationships instantiates a new PostCatalogItemDeleteJobResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCatalogItemDeleteJobResponseDataRelationships() *PostCatalogItemDeleteJobResponseDataRelationships {
	this := PostCatalogItemDeleteJobResponseDataRelationships{}
	return &this
}

// NewPostCatalogItemDeleteJobResponseDataRelationshipsWithDefaults instantiates a new PostCatalogItemDeleteJobResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCatalogItemDeleteJobResponseDataRelationshipsWithDefaults() *PostCatalogItemDeleteJobResponseDataRelationships {
	this := PostCatalogItemDeleteJobResponseDataRelationships{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PostCatalogItemDeleteJobResponseDataRelationships) GetItems() PostCatalogItemDeleteJobResponseDataRelationshipsItems {
	if o == nil || IsNil(o.Items) {
		var ret PostCatalogItemDeleteJobResponseDataRelationshipsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCatalogItemDeleteJobResponseDataRelationships) GetItemsOk() (*PostCatalogItemDeleteJobResponseDataRelationshipsItems, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PostCatalogItemDeleteJobResponseDataRelationships) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given PostCatalogItemDeleteJobResponseDataRelationshipsItems and assigns it to the Items field.
func (o *PostCatalogItemDeleteJobResponseDataRelationships) SetItems(v PostCatalogItemDeleteJobResponseDataRelationshipsItems) {
	o.Items = &v
}

func (o PostCatalogItemDeleteJobResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCatalogItemDeleteJobResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullablePostCatalogItemDeleteJobResponseDataRelationships struct {
	value *PostCatalogItemDeleteJobResponseDataRelationships
	isSet bool
}

func (v NullablePostCatalogItemDeleteJobResponseDataRelationships) Get() *PostCatalogItemDeleteJobResponseDataRelationships {
	return v.value
}

func (v *NullablePostCatalogItemDeleteJobResponseDataRelationships) Set(val *PostCatalogItemDeleteJobResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCatalogItemDeleteJobResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCatalogItemDeleteJobResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCatalogItemDeleteJobResponseDataRelationships(val *PostCatalogItemDeleteJobResponseDataRelationships) *NullablePostCatalogItemDeleteJobResponseDataRelationships {
	return &NullablePostCatalogItemDeleteJobResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePostCatalogItemDeleteJobResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCatalogItemDeleteJobResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


