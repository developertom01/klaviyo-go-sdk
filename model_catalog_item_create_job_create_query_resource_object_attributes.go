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

// checks if the CatalogItemCreateJobCreateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogItemCreateJobCreateQueryResourceObjectAttributes{}

// CatalogItemCreateJobCreateQueryResourceObjectAttributes struct for CatalogItemCreateJobCreateQueryResourceObjectAttributes
type CatalogItemCreateJobCreateQueryResourceObjectAttributes struct {
	Items CatalogItemCreateJobCreateQueryResourceObjectAttributesItems `json:"items"`
}

type _CatalogItemCreateJobCreateQueryResourceObjectAttributes CatalogItemCreateJobCreateQueryResourceObjectAttributes

// NewCatalogItemCreateJobCreateQueryResourceObjectAttributes instantiates a new CatalogItemCreateJobCreateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogItemCreateJobCreateQueryResourceObjectAttributes(items CatalogItemCreateJobCreateQueryResourceObjectAttributesItems) *CatalogItemCreateJobCreateQueryResourceObjectAttributes {
	this := CatalogItemCreateJobCreateQueryResourceObjectAttributes{}
	this.Items = items
	return &this
}

// NewCatalogItemCreateJobCreateQueryResourceObjectAttributesWithDefaults instantiates a new CatalogItemCreateJobCreateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogItemCreateJobCreateQueryResourceObjectAttributesWithDefaults() *CatalogItemCreateJobCreateQueryResourceObjectAttributes {
	this := CatalogItemCreateJobCreateQueryResourceObjectAttributes{}
	return &this
}

// GetItems returns the Items field value
func (o *CatalogItemCreateJobCreateQueryResourceObjectAttributes) GetItems() CatalogItemCreateJobCreateQueryResourceObjectAttributesItems {
	if o == nil {
		var ret CatalogItemCreateJobCreateQueryResourceObjectAttributesItems
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CatalogItemCreateJobCreateQueryResourceObjectAttributes) GetItemsOk() (*CatalogItemCreateJobCreateQueryResourceObjectAttributesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *CatalogItemCreateJobCreateQueryResourceObjectAttributes) SetItems(v CatalogItemCreateJobCreateQueryResourceObjectAttributesItems) {
	o.Items = v
}

func (o CatalogItemCreateJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogItemCreateJobCreateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *CatalogItemCreateJobCreateQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogItemCreateJobCreateQueryResourceObjectAttributes := _CatalogItemCreateJobCreateQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogItemCreateJobCreateQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = CatalogItemCreateJobCreateQueryResourceObjectAttributes(varCatalogItemCreateJobCreateQueryResourceObjectAttributes)

	return err
}

type NullableCatalogItemCreateJobCreateQueryResourceObjectAttributes struct {
	value *CatalogItemCreateJobCreateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableCatalogItemCreateJobCreateQueryResourceObjectAttributes) Get() *CatalogItemCreateJobCreateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableCatalogItemCreateJobCreateQueryResourceObjectAttributes) Set(val *CatalogItemCreateJobCreateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogItemCreateJobCreateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogItemCreateJobCreateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogItemCreateJobCreateQueryResourceObjectAttributes(val *CatalogItemCreateJobCreateQueryResourceObjectAttributes) *NullableCatalogItemCreateJobCreateQueryResourceObjectAttributes {
	return &NullableCatalogItemCreateJobCreateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableCatalogItemCreateJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogItemCreateJobCreateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


