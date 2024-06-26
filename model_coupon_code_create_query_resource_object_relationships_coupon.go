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

// checks if the CouponCodeCreateQueryResourceObjectRelationshipsCoupon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponCodeCreateQueryResourceObjectRelationshipsCoupon{}

// CouponCodeCreateQueryResourceObjectRelationshipsCoupon struct for CouponCodeCreateQueryResourceObjectRelationshipsCoupon
type CouponCodeCreateQueryResourceObjectRelationshipsCoupon struct {
	Data GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponData `json:"data"`
}

type _CouponCodeCreateQueryResourceObjectRelationshipsCoupon CouponCodeCreateQueryResourceObjectRelationshipsCoupon

// NewCouponCodeCreateQueryResourceObjectRelationshipsCoupon instantiates a new CouponCodeCreateQueryResourceObjectRelationshipsCoupon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodeCreateQueryResourceObjectRelationshipsCoupon(data GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponData) *CouponCodeCreateQueryResourceObjectRelationshipsCoupon {
	this := CouponCodeCreateQueryResourceObjectRelationshipsCoupon{}
	this.Data = data
	return &this
}

// NewCouponCodeCreateQueryResourceObjectRelationshipsCouponWithDefaults instantiates a new CouponCodeCreateQueryResourceObjectRelationshipsCoupon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodeCreateQueryResourceObjectRelationshipsCouponWithDefaults() *CouponCodeCreateQueryResourceObjectRelationshipsCoupon {
	this := CouponCodeCreateQueryResourceObjectRelationshipsCoupon{}
	return &this
}

// GetData returns the Data field value
func (o *CouponCodeCreateQueryResourceObjectRelationshipsCoupon) GetData() GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponData {
	if o == nil {
		var ret GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CouponCodeCreateQueryResourceObjectRelationshipsCoupon) GetDataOk() (*GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CouponCodeCreateQueryResourceObjectRelationshipsCoupon) SetData(v GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponData) {
	o.Data = v
}

func (o CouponCodeCreateQueryResourceObjectRelationshipsCoupon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponCodeCreateQueryResourceObjectRelationshipsCoupon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CouponCodeCreateQueryResourceObjectRelationshipsCoupon) UnmarshalJSON(data []byte) (err error) {
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

	varCouponCodeCreateQueryResourceObjectRelationshipsCoupon := _CouponCodeCreateQueryResourceObjectRelationshipsCoupon{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCouponCodeCreateQueryResourceObjectRelationshipsCoupon)

	if err != nil {
		return err
	}

	*o = CouponCodeCreateQueryResourceObjectRelationshipsCoupon(varCouponCodeCreateQueryResourceObjectRelationshipsCoupon)

	return err
}

type NullableCouponCodeCreateQueryResourceObjectRelationshipsCoupon struct {
	value *CouponCodeCreateQueryResourceObjectRelationshipsCoupon
	isSet bool
}

func (v NullableCouponCodeCreateQueryResourceObjectRelationshipsCoupon) Get() *CouponCodeCreateQueryResourceObjectRelationshipsCoupon {
	return v.value
}

func (v *NullableCouponCodeCreateQueryResourceObjectRelationshipsCoupon) Set(val *CouponCodeCreateQueryResourceObjectRelationshipsCoupon) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodeCreateQueryResourceObjectRelationshipsCoupon) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodeCreateQueryResourceObjectRelationshipsCoupon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodeCreateQueryResourceObjectRelationshipsCoupon(val *CouponCodeCreateQueryResourceObjectRelationshipsCoupon) *NullableCouponCodeCreateQueryResourceObjectRelationshipsCoupon {
	return &NullableCouponCodeCreateQueryResourceObjectRelationshipsCoupon{value: val, isSet: true}
}

func (v NullableCouponCodeCreateQueryResourceObjectRelationshipsCoupon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodeCreateQueryResourceObjectRelationshipsCoupon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


