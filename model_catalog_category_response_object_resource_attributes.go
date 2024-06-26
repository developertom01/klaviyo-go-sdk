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
	"time"
)

// checks if the CatalogCategoryResponseObjectResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogCategoryResponseObjectResourceAttributes{}

// CatalogCategoryResponseObjectResourceAttributes struct for CatalogCategoryResponseObjectResourceAttributes
type CatalogCategoryResponseObjectResourceAttributes struct {
	// The ID of the catalog category in an external system.
	ExternalId NullableString `json:"external_id,omitempty"`
	// The name of the catalog category.
	Name NullableString `json:"name,omitempty"`
	// Date and time when the catalog category was last updated, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
	Updated NullableTime `json:"updated,omitempty"`
}

// NewCatalogCategoryResponseObjectResourceAttributes instantiates a new CatalogCategoryResponseObjectResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogCategoryResponseObjectResourceAttributes() *CatalogCategoryResponseObjectResourceAttributes {
	this := CatalogCategoryResponseObjectResourceAttributes{}
	return &this
}

// NewCatalogCategoryResponseObjectResourceAttributesWithDefaults instantiates a new CatalogCategoryResponseObjectResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogCategoryResponseObjectResourceAttributesWithDefaults() *CatalogCategoryResponseObjectResourceAttributes {
	this := CatalogCategoryResponseObjectResourceAttributes{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogCategoryResponseObjectResourceAttributes) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogCategoryResponseObjectResourceAttributes) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *CatalogCategoryResponseObjectResourceAttributes) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *CatalogCategoryResponseObjectResourceAttributes) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *CatalogCategoryResponseObjectResourceAttributes) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *CatalogCategoryResponseObjectResourceAttributes) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogCategoryResponseObjectResourceAttributes) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogCategoryResponseObjectResourceAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CatalogCategoryResponseObjectResourceAttributes) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CatalogCategoryResponseObjectResourceAttributes) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CatalogCategoryResponseObjectResourceAttributes) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CatalogCategoryResponseObjectResourceAttributes) UnsetName() {
	o.Name.Unset()
}

// GetUpdated returns the Updated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogCategoryResponseObjectResourceAttributes) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Updated.Get()
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogCategoryResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Updated.Get(), o.Updated.IsSet()
}

// HasUpdated returns a boolean if a field has been set.
func (o *CatalogCategoryResponseObjectResourceAttributes) HasUpdated() bool {
	if o != nil && o.Updated.IsSet() {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given NullableTime and assigns it to the Updated field.
func (o *CatalogCategoryResponseObjectResourceAttributes) SetUpdated(v time.Time) {
	o.Updated.Set(&v)
}
// SetUpdatedNil sets the value for Updated to be an explicit nil
func (o *CatalogCategoryResponseObjectResourceAttributes) SetUpdatedNil() {
	o.Updated.Set(nil)
}

// UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
func (o *CatalogCategoryResponseObjectResourceAttributes) UnsetUpdated() {
	o.Updated.Unset()
}

func (o CatalogCategoryResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogCategoryResponseObjectResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalId.IsSet() {
		toSerialize["external_id"] = o.ExternalId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Updated.IsSet() {
		toSerialize["updated"] = o.Updated.Get()
	}
	return toSerialize, nil
}

type NullableCatalogCategoryResponseObjectResourceAttributes struct {
	value *CatalogCategoryResponseObjectResourceAttributes
	isSet bool
}

func (v NullableCatalogCategoryResponseObjectResourceAttributes) Get() *CatalogCategoryResponseObjectResourceAttributes {
	return v.value
}

func (v *NullableCatalogCategoryResponseObjectResourceAttributes) Set(val *CatalogCategoryResponseObjectResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogCategoryResponseObjectResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogCategoryResponseObjectResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogCategoryResponseObjectResourceAttributes(val *CatalogCategoryResponseObjectResourceAttributes) *NullableCatalogCategoryResponseObjectResourceAttributes {
	return &NullableCatalogCategoryResponseObjectResourceAttributes{value: val, isSet: true}
}

func (v NullableCatalogCategoryResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogCategoryResponseObjectResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


