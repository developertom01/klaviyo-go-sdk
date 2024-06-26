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

// checks if the CatalogCategoryItemOp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogCategoryItemOp{}

// CatalogCategoryItemOp struct for CatalogCategoryItemOp
type CatalogCategoryItemOp struct {
	Data []GetCatalogCategoryItemListResponseCollectionDataInner `json:"data"`
}

type _CatalogCategoryItemOp CatalogCategoryItemOp

// NewCatalogCategoryItemOp instantiates a new CatalogCategoryItemOp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogCategoryItemOp(data []GetCatalogCategoryItemListResponseCollectionDataInner) *CatalogCategoryItemOp {
	this := CatalogCategoryItemOp{}
	this.Data = data
	return &this
}

// NewCatalogCategoryItemOpWithDefaults instantiates a new CatalogCategoryItemOp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogCategoryItemOpWithDefaults() *CatalogCategoryItemOp {
	this := CatalogCategoryItemOp{}
	return &this
}

// GetData returns the Data field value
func (o *CatalogCategoryItemOp) GetData() []GetCatalogCategoryItemListResponseCollectionDataInner {
	if o == nil {
		var ret []GetCatalogCategoryItemListResponseCollectionDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CatalogCategoryItemOp) GetDataOk() ([]GetCatalogCategoryItemListResponseCollectionDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CatalogCategoryItemOp) SetData(v []GetCatalogCategoryItemListResponseCollectionDataInner) {
	o.Data = v
}

func (o CatalogCategoryItemOp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogCategoryItemOp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CatalogCategoryItemOp) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogCategoryItemOp := _CatalogCategoryItemOp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogCategoryItemOp)

	if err != nil {
		return err
	}

	*o = CatalogCategoryItemOp(varCatalogCategoryItemOp)

	return err
}

type NullableCatalogCategoryItemOp struct {
	value *CatalogCategoryItemOp
	isSet bool
}

func (v NullableCatalogCategoryItemOp) Get() *CatalogCategoryItemOp {
	return v.value
}

func (v *NullableCatalogCategoryItemOp) Set(val *CatalogCategoryItemOp) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogCategoryItemOp) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogCategoryItemOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogCategoryItemOp(val *CatalogCategoryItemOp) *NullableCatalogCategoryItemOp {
	return &NullableCatalogCategoryItemOp{value: val, isSet: true}
}

func (v NullableCatalogCategoryItemOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogCategoryItemOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


