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

// checks if the CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems{}

// CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems Array of catalog items to update.
type CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems struct {
	Data []CatalogItemUpdateQueryResourceObject `json:"data"`
}

type _CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems

// NewCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems instantiates a new CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems(data []CatalogItemUpdateQueryResourceObject) *CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems {
	this := CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems{}
	this.Data = data
	return &this
}

// NewCatalogItemUpdateJobCreateQueryResourceObjectAttributesItemsWithDefaults instantiates a new CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogItemUpdateJobCreateQueryResourceObjectAttributesItemsWithDefaults() *CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems {
	this := CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems{}
	return &this
}

// GetData returns the Data field value
func (o *CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) GetData() []CatalogItemUpdateQueryResourceObject {
	if o == nil {
		var ret []CatalogItemUpdateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) GetDataOk() ([]CatalogItemUpdateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) SetData(v []CatalogItemUpdateQueryResourceObject) {
	o.Data = v
}

func (o CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems := _CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems)

	if err != nil {
		return err
	}

	*o = CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems(varCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems)

	return err
}

type NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems struct {
	value *CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems
	isSet bool
}

func (v NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) Get() *CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems {
	return v.value
}

func (v *NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) Set(val *CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems(val *CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) *NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems {
	return &NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems{value: val, isSet: true}
}

func (v NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogItemUpdateJobCreateQueryResourceObjectAttributesItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


