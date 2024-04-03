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

// checks if the CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes{}

// CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes Array of coupon codes to create.
type CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes struct {
	Data []CouponCodeCreateQueryResourceObject `json:"data"`
}

type _CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes

// NewCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes instantiates a new CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes(data []CouponCodeCreateQueryResourceObject) *CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes {
	this := CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes{}
	this.Data = data
	return &this
}

// NewCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodesWithDefaults instantiates a new CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodesWithDefaults() *CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes {
	this := CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes{}
	return &this
}

// GetData returns the Data field value
func (o *CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) GetData() []CouponCodeCreateQueryResourceObject {
	if o == nil {
		var ret []CouponCodeCreateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) GetDataOk() ([]CouponCodeCreateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) SetData(v []CouponCodeCreateQueryResourceObject) {
	o.Data = v
}

func (o CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) UnmarshalJSON(data []byte) (err error) {
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

	varCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes := _CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes)

	if err != nil {
		return err
	}

	*o = CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes(varCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes)

	return err
}

type NullableCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes struct {
	value *CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes
	isSet bool
}

func (v NullableCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) Get() *CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes {
	return v.value
}

func (v *NullableCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) Set(val *CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes(val *CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) *NullableCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes {
	return &NullableCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes{value: val, isSet: true}
}

func (v NullableCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

