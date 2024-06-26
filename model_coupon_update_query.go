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

// checks if the CouponUpdateQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponUpdateQuery{}

// CouponUpdateQuery struct for CouponUpdateQuery
type CouponUpdateQuery struct {
	Data CouponUpdateQueryResourceObject `json:"data"`
}

type _CouponUpdateQuery CouponUpdateQuery

// NewCouponUpdateQuery instantiates a new CouponUpdateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponUpdateQuery(data CouponUpdateQueryResourceObject) *CouponUpdateQuery {
	this := CouponUpdateQuery{}
	this.Data = data
	return &this
}

// NewCouponUpdateQueryWithDefaults instantiates a new CouponUpdateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponUpdateQueryWithDefaults() *CouponUpdateQuery {
	this := CouponUpdateQuery{}
	return &this
}

// GetData returns the Data field value
func (o *CouponUpdateQuery) GetData() CouponUpdateQueryResourceObject {
	if o == nil {
		var ret CouponUpdateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CouponUpdateQuery) GetDataOk() (*CouponUpdateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CouponUpdateQuery) SetData(v CouponUpdateQueryResourceObject) {
	o.Data = v
}

func (o CouponUpdateQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponUpdateQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CouponUpdateQuery) UnmarshalJSON(data []byte) (err error) {
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

	varCouponUpdateQuery := _CouponUpdateQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCouponUpdateQuery)

	if err != nil {
		return err
	}

	*o = CouponUpdateQuery(varCouponUpdateQuery)

	return err
}

type NullableCouponUpdateQuery struct {
	value *CouponUpdateQuery
	isSet bool
}

func (v NullableCouponUpdateQuery) Get() *CouponUpdateQuery {
	return v.value
}

func (v *NullableCouponUpdateQuery) Set(val *CouponUpdateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponUpdateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponUpdateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponUpdateQuery(val *CouponUpdateQuery) *NullableCouponUpdateQuery {
	return &NullableCouponUpdateQuery{value: val, isSet: true}
}

func (v NullableCouponUpdateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponUpdateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


