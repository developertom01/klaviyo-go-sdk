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

// checks if the PostCatalogVariantDeleteJobResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCatalogVariantDeleteJobResponseDataRelationships{}

// PostCatalogVariantDeleteJobResponseDataRelationships struct for PostCatalogVariantDeleteJobResponseDataRelationships
type PostCatalogVariantDeleteJobResponseDataRelationships struct {
	Variants *PostCatalogVariantDeleteJobResponseDataRelationshipsVariants `json:"variants,omitempty"`
}

// NewPostCatalogVariantDeleteJobResponseDataRelationships instantiates a new PostCatalogVariantDeleteJobResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCatalogVariantDeleteJobResponseDataRelationships() *PostCatalogVariantDeleteJobResponseDataRelationships {
	this := PostCatalogVariantDeleteJobResponseDataRelationships{}
	return &this
}

// NewPostCatalogVariantDeleteJobResponseDataRelationshipsWithDefaults instantiates a new PostCatalogVariantDeleteJobResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCatalogVariantDeleteJobResponseDataRelationshipsWithDefaults() *PostCatalogVariantDeleteJobResponseDataRelationships {
	this := PostCatalogVariantDeleteJobResponseDataRelationships{}
	return &this
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *PostCatalogVariantDeleteJobResponseDataRelationships) GetVariants() PostCatalogVariantDeleteJobResponseDataRelationshipsVariants {
	if o == nil || IsNil(o.Variants) {
		var ret PostCatalogVariantDeleteJobResponseDataRelationshipsVariants
		return ret
	}
	return *o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCatalogVariantDeleteJobResponseDataRelationships) GetVariantsOk() (*PostCatalogVariantDeleteJobResponseDataRelationshipsVariants, bool) {
	if o == nil || IsNil(o.Variants) {
		return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *PostCatalogVariantDeleteJobResponseDataRelationships) HasVariants() bool {
	if o != nil && !IsNil(o.Variants) {
		return true
	}

	return false
}

// SetVariants gets a reference to the given PostCatalogVariantDeleteJobResponseDataRelationshipsVariants and assigns it to the Variants field.
func (o *PostCatalogVariantDeleteJobResponseDataRelationships) SetVariants(v PostCatalogVariantDeleteJobResponseDataRelationshipsVariants) {
	o.Variants = &v
}

func (o PostCatalogVariantDeleteJobResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCatalogVariantDeleteJobResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Variants) {
		toSerialize["variants"] = o.Variants
	}
	return toSerialize, nil
}

type NullablePostCatalogVariantDeleteJobResponseDataRelationships struct {
	value *PostCatalogVariantDeleteJobResponseDataRelationships
	isSet bool
}

func (v NullablePostCatalogVariantDeleteJobResponseDataRelationships) Get() *PostCatalogVariantDeleteJobResponseDataRelationships {
	return v.value
}

func (v *NullablePostCatalogVariantDeleteJobResponseDataRelationships) Set(val *PostCatalogVariantDeleteJobResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCatalogVariantDeleteJobResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCatalogVariantDeleteJobResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCatalogVariantDeleteJobResponseDataRelationships(val *PostCatalogVariantDeleteJobResponseDataRelationships) *NullablePostCatalogVariantDeleteJobResponseDataRelationships {
	return &NullablePostCatalogVariantDeleteJobResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePostCatalogVariantDeleteJobResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCatalogVariantDeleteJobResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


