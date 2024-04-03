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

// checks if the CatalogItemUpdateJobCreateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogItemUpdateJobCreateQueryResourceObjectAttributes{}

// CatalogItemUpdateJobCreateQueryResourceObjectAttributes struct for CatalogItemUpdateJobCreateQueryResourceObjectAttributes
type CatalogItemUpdateJobCreateQueryResourceObjectAttributes struct {
	Items CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems `json:"items"`
}

type _CatalogItemUpdateJobCreateQueryResourceObjectAttributes CatalogItemUpdateJobCreateQueryResourceObjectAttributes

// NewCatalogItemUpdateJobCreateQueryResourceObjectAttributes instantiates a new CatalogItemUpdateJobCreateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogItemUpdateJobCreateQueryResourceObjectAttributes(items CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) *CatalogItemUpdateJobCreateQueryResourceObjectAttributes {
	this := CatalogItemUpdateJobCreateQueryResourceObjectAttributes{}
	this.Items = items
	return &this
}

// NewCatalogItemUpdateJobCreateQueryResourceObjectAttributesWithDefaults instantiates a new CatalogItemUpdateJobCreateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogItemUpdateJobCreateQueryResourceObjectAttributesWithDefaults() *CatalogItemUpdateJobCreateQueryResourceObjectAttributes {
	this := CatalogItemUpdateJobCreateQueryResourceObjectAttributes{}
	return &this
}

// GetItems returns the Items field value
func (o *CatalogItemUpdateJobCreateQueryResourceObjectAttributes) GetItems() CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems {
	if o == nil {
		var ret CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CatalogItemUpdateJobCreateQueryResourceObjectAttributes) GetItemsOk() (*CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *CatalogItemUpdateJobCreateQueryResourceObjectAttributes) SetItems(v CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) {
	o.Items = v
}

func (o CatalogItemUpdateJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogItemUpdateJobCreateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *CatalogItemUpdateJobCreateQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
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

	varCatalogItemUpdateJobCreateQueryResourceObjectAttributes := _CatalogItemUpdateJobCreateQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogItemUpdateJobCreateQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = CatalogItemUpdateJobCreateQueryResourceObjectAttributes(varCatalogItemUpdateJobCreateQueryResourceObjectAttributes)

	return err
}

type NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributes struct {
	value *CatalogItemUpdateJobCreateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributes) Get() *CatalogItemUpdateJobCreateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributes) Set(val *CatalogItemUpdateJobCreateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogItemUpdateJobCreateQueryResourceObjectAttributes(val *CatalogItemUpdateJobCreateQueryResourceObjectAttributes) *NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributes {
	return &NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

