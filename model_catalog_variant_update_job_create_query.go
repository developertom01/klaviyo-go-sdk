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

// checks if the CatalogVariantUpdateJobCreateQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogVariantUpdateJobCreateQuery{}

// CatalogVariantUpdateJobCreateQuery struct for CatalogVariantUpdateJobCreateQuery
type CatalogVariantUpdateJobCreateQuery struct {
	Data CatalogVariantUpdateJobCreateQueryResourceObject `json:"data"`
}

type _CatalogVariantUpdateJobCreateQuery CatalogVariantUpdateJobCreateQuery

// NewCatalogVariantUpdateJobCreateQuery instantiates a new CatalogVariantUpdateJobCreateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogVariantUpdateJobCreateQuery(data CatalogVariantUpdateJobCreateQueryResourceObject) *CatalogVariantUpdateJobCreateQuery {
	this := CatalogVariantUpdateJobCreateQuery{}
	this.Data = data
	return &this
}

// NewCatalogVariantUpdateJobCreateQueryWithDefaults instantiates a new CatalogVariantUpdateJobCreateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogVariantUpdateJobCreateQueryWithDefaults() *CatalogVariantUpdateJobCreateQuery {
	this := CatalogVariantUpdateJobCreateQuery{}
	return &this
}

// GetData returns the Data field value
func (o *CatalogVariantUpdateJobCreateQuery) GetData() CatalogVariantUpdateJobCreateQueryResourceObject {
	if o == nil {
		var ret CatalogVariantUpdateJobCreateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantUpdateJobCreateQuery) GetDataOk() (*CatalogVariantUpdateJobCreateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CatalogVariantUpdateJobCreateQuery) SetData(v CatalogVariantUpdateJobCreateQueryResourceObject) {
	o.Data = v
}

func (o CatalogVariantUpdateJobCreateQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogVariantUpdateJobCreateQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CatalogVariantUpdateJobCreateQuery) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogVariantUpdateJobCreateQuery := _CatalogVariantUpdateJobCreateQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogVariantUpdateJobCreateQuery)

	if err != nil {
		return err
	}

	*o = CatalogVariantUpdateJobCreateQuery(varCatalogVariantUpdateJobCreateQuery)

	return err
}

type NullableCatalogVariantUpdateJobCreateQuery struct {
	value *CatalogVariantUpdateJobCreateQuery
	isSet bool
}

func (v NullableCatalogVariantUpdateJobCreateQuery) Get() *CatalogVariantUpdateJobCreateQuery {
	return v.value
}

func (v *NullableCatalogVariantUpdateJobCreateQuery) Set(val *CatalogVariantUpdateJobCreateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogVariantUpdateJobCreateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogVariantUpdateJobCreateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogVariantUpdateJobCreateQuery(val *CatalogVariantUpdateJobCreateQuery) *NullableCatalogVariantUpdateJobCreateQuery {
	return &NullableCatalogVariantUpdateJobCreateQuery{value: val, isSet: true}
}

func (v NullableCatalogVariantUpdateJobCreateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogVariantUpdateJobCreateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


