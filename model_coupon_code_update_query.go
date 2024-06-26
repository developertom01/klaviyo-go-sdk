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

// checks if the CouponCodeUpdateQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponCodeUpdateQuery{}

// CouponCodeUpdateQuery struct for CouponCodeUpdateQuery
type CouponCodeUpdateQuery struct {
	Data CouponCodeUpdateQueryResourceObject `json:"data"`
}

type _CouponCodeUpdateQuery CouponCodeUpdateQuery

// NewCouponCodeUpdateQuery instantiates a new CouponCodeUpdateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodeUpdateQuery(data CouponCodeUpdateQueryResourceObject) *CouponCodeUpdateQuery {
	this := CouponCodeUpdateQuery{}
	this.Data = data
	return &this
}

// NewCouponCodeUpdateQueryWithDefaults instantiates a new CouponCodeUpdateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodeUpdateQueryWithDefaults() *CouponCodeUpdateQuery {
	this := CouponCodeUpdateQuery{}
	return &this
}

// GetData returns the Data field value
func (o *CouponCodeUpdateQuery) GetData() CouponCodeUpdateQueryResourceObject {
	if o == nil {
		var ret CouponCodeUpdateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CouponCodeUpdateQuery) GetDataOk() (*CouponCodeUpdateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CouponCodeUpdateQuery) SetData(v CouponCodeUpdateQueryResourceObject) {
	o.Data = v
}

func (o CouponCodeUpdateQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponCodeUpdateQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CouponCodeUpdateQuery) UnmarshalJSON(data []byte) (err error) {
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

	varCouponCodeUpdateQuery := _CouponCodeUpdateQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCouponCodeUpdateQuery)

	if err != nil {
		return err
	}

	*o = CouponCodeUpdateQuery(varCouponCodeUpdateQuery)

	return err
}

type NullableCouponCodeUpdateQuery struct {
	value *CouponCodeUpdateQuery
	isSet bool
}

func (v NullableCouponCodeUpdateQuery) Get() *CouponCodeUpdateQuery {
	return v.value
}

func (v *NullableCouponCodeUpdateQuery) Set(val *CouponCodeUpdateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodeUpdateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodeUpdateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodeUpdateQuery(val *CouponCodeUpdateQuery) *NullableCouponCodeUpdateQuery {
	return &NullableCouponCodeUpdateQuery{value: val, isSet: true}
}

func (v NullableCouponCodeUpdateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodeUpdateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


