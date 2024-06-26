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

// checks if the CatalogCategoryUpdateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogCategoryUpdateQueryResourceObjectAttributes{}

// CatalogCategoryUpdateQueryResourceObjectAttributes struct for CatalogCategoryUpdateQueryResourceObjectAttributes
type CatalogCategoryUpdateQueryResourceObjectAttributes struct {
	// The name of the catalog category.
	Name NullableString `json:"name,omitempty"`
}

// NewCatalogCategoryUpdateQueryResourceObjectAttributes instantiates a new CatalogCategoryUpdateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogCategoryUpdateQueryResourceObjectAttributes() *CatalogCategoryUpdateQueryResourceObjectAttributes {
	this := CatalogCategoryUpdateQueryResourceObjectAttributes{}
	return &this
}

// NewCatalogCategoryUpdateQueryResourceObjectAttributesWithDefaults instantiates a new CatalogCategoryUpdateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogCategoryUpdateQueryResourceObjectAttributesWithDefaults() *CatalogCategoryUpdateQueryResourceObjectAttributes {
	this := CatalogCategoryUpdateQueryResourceObjectAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogCategoryUpdateQueryResourceObjectAttributes) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogCategoryUpdateQueryResourceObjectAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CatalogCategoryUpdateQueryResourceObjectAttributes) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CatalogCategoryUpdateQueryResourceObjectAttributes) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CatalogCategoryUpdateQueryResourceObjectAttributes) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CatalogCategoryUpdateQueryResourceObjectAttributes) UnsetName() {
	o.Name.Unset()
}

func (o CatalogCategoryUpdateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogCategoryUpdateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableCatalogCategoryUpdateQueryResourceObjectAttributes struct {
	value *CatalogCategoryUpdateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableCatalogCategoryUpdateQueryResourceObjectAttributes) Get() *CatalogCategoryUpdateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableCatalogCategoryUpdateQueryResourceObjectAttributes) Set(val *CatalogCategoryUpdateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogCategoryUpdateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogCategoryUpdateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogCategoryUpdateQueryResourceObjectAttributes(val *CatalogCategoryUpdateQueryResourceObjectAttributes) *NullableCatalogCategoryUpdateQueryResourceObjectAttributes {
	return &NullableCatalogCategoryUpdateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableCatalogCategoryUpdateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogCategoryUpdateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


