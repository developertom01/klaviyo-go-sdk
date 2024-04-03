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

// checks if the CatalogItemDeleteJobCreateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogItemDeleteJobCreateQueryResourceObjectAttributes{}

// CatalogItemDeleteJobCreateQueryResourceObjectAttributes struct for CatalogItemDeleteJobCreateQueryResourceObjectAttributes
type CatalogItemDeleteJobCreateQueryResourceObjectAttributes struct {
	Items CatalogItemDeleteJobCreateQueryResourceObjectAttributesItems `json:"items"`
}

type _CatalogItemDeleteJobCreateQueryResourceObjectAttributes CatalogItemDeleteJobCreateQueryResourceObjectAttributes

// NewCatalogItemDeleteJobCreateQueryResourceObjectAttributes instantiates a new CatalogItemDeleteJobCreateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogItemDeleteJobCreateQueryResourceObjectAttributes(items CatalogItemDeleteJobCreateQueryResourceObjectAttributesItems) *CatalogItemDeleteJobCreateQueryResourceObjectAttributes {
	this := CatalogItemDeleteJobCreateQueryResourceObjectAttributes{}
	this.Items = items
	return &this
}

// NewCatalogItemDeleteJobCreateQueryResourceObjectAttributesWithDefaults instantiates a new CatalogItemDeleteJobCreateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogItemDeleteJobCreateQueryResourceObjectAttributesWithDefaults() *CatalogItemDeleteJobCreateQueryResourceObjectAttributes {
	this := CatalogItemDeleteJobCreateQueryResourceObjectAttributes{}
	return &this
}

// GetItems returns the Items field value
func (o *CatalogItemDeleteJobCreateQueryResourceObjectAttributes) GetItems() CatalogItemDeleteJobCreateQueryResourceObjectAttributesItems {
	if o == nil {
		var ret CatalogItemDeleteJobCreateQueryResourceObjectAttributesItems
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CatalogItemDeleteJobCreateQueryResourceObjectAttributes) GetItemsOk() (*CatalogItemDeleteJobCreateQueryResourceObjectAttributesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *CatalogItemDeleteJobCreateQueryResourceObjectAttributes) SetItems(v CatalogItemDeleteJobCreateQueryResourceObjectAttributesItems) {
	o.Items = v
}

func (o CatalogItemDeleteJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogItemDeleteJobCreateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *CatalogItemDeleteJobCreateQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogItemDeleteJobCreateQueryResourceObjectAttributes := _CatalogItemDeleteJobCreateQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogItemDeleteJobCreateQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = CatalogItemDeleteJobCreateQueryResourceObjectAttributes(varCatalogItemDeleteJobCreateQueryResourceObjectAttributes)

	return err
}

type NullableCatalogItemDeleteJobCreateQueryResourceObjectAttributes struct {
	value *CatalogItemDeleteJobCreateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableCatalogItemDeleteJobCreateQueryResourceObjectAttributes) Get() *CatalogItemDeleteJobCreateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableCatalogItemDeleteJobCreateQueryResourceObjectAttributes) Set(val *CatalogItemDeleteJobCreateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogItemDeleteJobCreateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogItemDeleteJobCreateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogItemDeleteJobCreateQueryResourceObjectAttributes(val *CatalogItemDeleteJobCreateQueryResourceObjectAttributes) *NullableCatalogItemDeleteJobCreateQueryResourceObjectAttributes {
	return &NullableCatalogItemDeleteJobCreateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableCatalogItemDeleteJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogItemDeleteJobCreateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


