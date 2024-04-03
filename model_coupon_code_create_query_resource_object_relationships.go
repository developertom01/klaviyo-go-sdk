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

// checks if the CouponCodeCreateQueryResourceObjectRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponCodeCreateQueryResourceObjectRelationships{}

// CouponCodeCreateQueryResourceObjectRelationships struct for CouponCodeCreateQueryResourceObjectRelationships
type CouponCodeCreateQueryResourceObjectRelationships struct {
	Coupon CouponCodeCreateQueryResourceObjectRelationshipsCoupon `json:"coupon"`
}

type _CouponCodeCreateQueryResourceObjectRelationships CouponCodeCreateQueryResourceObjectRelationships

// NewCouponCodeCreateQueryResourceObjectRelationships instantiates a new CouponCodeCreateQueryResourceObjectRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodeCreateQueryResourceObjectRelationships(coupon CouponCodeCreateQueryResourceObjectRelationshipsCoupon) *CouponCodeCreateQueryResourceObjectRelationships {
	this := CouponCodeCreateQueryResourceObjectRelationships{}
	this.Coupon = coupon
	return &this
}

// NewCouponCodeCreateQueryResourceObjectRelationshipsWithDefaults instantiates a new CouponCodeCreateQueryResourceObjectRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodeCreateQueryResourceObjectRelationshipsWithDefaults() *CouponCodeCreateQueryResourceObjectRelationships {
	this := CouponCodeCreateQueryResourceObjectRelationships{}
	return &this
}

// GetCoupon returns the Coupon field value
func (o *CouponCodeCreateQueryResourceObjectRelationships) GetCoupon() CouponCodeCreateQueryResourceObjectRelationshipsCoupon {
	if o == nil {
		var ret CouponCodeCreateQueryResourceObjectRelationshipsCoupon
		return ret
	}

	return o.Coupon
}

// GetCouponOk returns a tuple with the Coupon field value
// and a boolean to check if the value has been set.
func (o *CouponCodeCreateQueryResourceObjectRelationships) GetCouponOk() (*CouponCodeCreateQueryResourceObjectRelationshipsCoupon, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coupon, true
}

// SetCoupon sets field value
func (o *CouponCodeCreateQueryResourceObjectRelationships) SetCoupon(v CouponCodeCreateQueryResourceObjectRelationshipsCoupon) {
	o.Coupon = v
}

func (o CouponCodeCreateQueryResourceObjectRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponCodeCreateQueryResourceObjectRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["coupon"] = o.Coupon
	return toSerialize, nil
}

func (o *CouponCodeCreateQueryResourceObjectRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"coupon",
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

	varCouponCodeCreateQueryResourceObjectRelationships := _CouponCodeCreateQueryResourceObjectRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCouponCodeCreateQueryResourceObjectRelationships)

	if err != nil {
		return err
	}

	*o = CouponCodeCreateQueryResourceObjectRelationships(varCouponCodeCreateQueryResourceObjectRelationships)

	return err
}

type NullableCouponCodeCreateQueryResourceObjectRelationships struct {
	value *CouponCodeCreateQueryResourceObjectRelationships
	isSet bool
}

func (v NullableCouponCodeCreateQueryResourceObjectRelationships) Get() *CouponCodeCreateQueryResourceObjectRelationships {
	return v.value
}

func (v *NullableCouponCodeCreateQueryResourceObjectRelationships) Set(val *CouponCodeCreateQueryResourceObjectRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodeCreateQueryResourceObjectRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodeCreateQueryResourceObjectRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodeCreateQueryResourceObjectRelationships(val *CouponCodeCreateQueryResourceObjectRelationships) *NullableCouponCodeCreateQueryResourceObjectRelationships {
	return &NullableCouponCodeCreateQueryResourceObjectRelationships{value: val, isSet: true}
}

func (v NullableCouponCodeCreateQueryResourceObjectRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodeCreateQueryResourceObjectRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

