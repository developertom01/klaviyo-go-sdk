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

// checks if the CatalogItemUpdateQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogItemUpdateQuery{}

// CatalogItemUpdateQuery struct for CatalogItemUpdateQuery
type CatalogItemUpdateQuery struct {
	Data CatalogItemUpdateQueryResourceObject `json:"data"`
}

type _CatalogItemUpdateQuery CatalogItemUpdateQuery

// NewCatalogItemUpdateQuery instantiates a new CatalogItemUpdateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogItemUpdateQuery(data CatalogItemUpdateQueryResourceObject) *CatalogItemUpdateQuery {
	this := CatalogItemUpdateQuery{}
	this.Data = data
	return &this
}

// NewCatalogItemUpdateQueryWithDefaults instantiates a new CatalogItemUpdateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogItemUpdateQueryWithDefaults() *CatalogItemUpdateQuery {
	this := CatalogItemUpdateQuery{}
	return &this
}

// GetData returns the Data field value
func (o *CatalogItemUpdateQuery) GetData() CatalogItemUpdateQueryResourceObject {
	if o == nil {
		var ret CatalogItemUpdateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CatalogItemUpdateQuery) GetDataOk() (*CatalogItemUpdateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CatalogItemUpdateQuery) SetData(v CatalogItemUpdateQueryResourceObject) {
	o.Data = v
}

func (o CatalogItemUpdateQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogItemUpdateQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CatalogItemUpdateQuery) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogItemUpdateQuery := _CatalogItemUpdateQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogItemUpdateQuery)

	if err != nil {
		return err
	}

	*o = CatalogItemUpdateQuery(varCatalogItemUpdateQuery)

	return err
}

type NullableCatalogItemUpdateQuery struct {
	value *CatalogItemUpdateQuery
	isSet bool
}

func (v NullableCatalogItemUpdateQuery) Get() *CatalogItemUpdateQuery {
	return v.value
}

func (v *NullableCatalogItemUpdateQuery) Set(val *CatalogItemUpdateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogItemUpdateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogItemUpdateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogItemUpdateQuery(val *CatalogItemUpdateQuery) *NullableCatalogItemUpdateQuery {
	return &NullableCatalogItemUpdateQuery{value: val, isSet: true}
}

func (v NullableCatalogItemUpdateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogItemUpdateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

