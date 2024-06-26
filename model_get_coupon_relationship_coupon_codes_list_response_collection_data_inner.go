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

// checks if the GetCouponRelationshipCouponCodesListResponseCollectionDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCouponRelationshipCouponCodesListResponseCollectionDataInner{}

// GetCouponRelationshipCouponCodesListResponseCollectionDataInner struct for GetCouponRelationshipCouponCodesListResponseCollectionDataInner
type GetCouponRelationshipCouponCodesListResponseCollectionDataInner struct {
	Type CouponCodeEnum `json:"type"`
	// A list of coupon code IDs that are in the given coupon.
	Id string `json:"id"`
}

type _GetCouponRelationshipCouponCodesListResponseCollectionDataInner GetCouponRelationshipCouponCodesListResponseCollectionDataInner

// NewGetCouponRelationshipCouponCodesListResponseCollectionDataInner instantiates a new GetCouponRelationshipCouponCodesListResponseCollectionDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCouponRelationshipCouponCodesListResponseCollectionDataInner(type_ CouponCodeEnum, id string) *GetCouponRelationshipCouponCodesListResponseCollectionDataInner {
	this := GetCouponRelationshipCouponCodesListResponseCollectionDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGetCouponRelationshipCouponCodesListResponseCollectionDataInnerWithDefaults instantiates a new GetCouponRelationshipCouponCodesListResponseCollectionDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCouponRelationshipCouponCodesListResponseCollectionDataInnerWithDefaults() *GetCouponRelationshipCouponCodesListResponseCollectionDataInner {
	this := GetCouponRelationshipCouponCodesListResponseCollectionDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GetCouponRelationshipCouponCodesListResponseCollectionDataInner) GetType() CouponCodeEnum {
	if o == nil {
		var ret CouponCodeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetCouponRelationshipCouponCodesListResponseCollectionDataInner) GetTypeOk() (*CouponCodeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetCouponRelationshipCouponCodesListResponseCollectionDataInner) SetType(v CouponCodeEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetCouponRelationshipCouponCodesListResponseCollectionDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetCouponRelationshipCouponCodesListResponseCollectionDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetCouponRelationshipCouponCodesListResponseCollectionDataInner) SetId(v string) {
	o.Id = v
}

func (o GetCouponRelationshipCouponCodesListResponseCollectionDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCouponRelationshipCouponCodesListResponseCollectionDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *GetCouponRelationshipCouponCodesListResponseCollectionDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varGetCouponRelationshipCouponCodesListResponseCollectionDataInner := _GetCouponRelationshipCouponCodesListResponseCollectionDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCouponRelationshipCouponCodesListResponseCollectionDataInner)

	if err != nil {
		return err
	}

	*o = GetCouponRelationshipCouponCodesListResponseCollectionDataInner(varGetCouponRelationshipCouponCodesListResponseCollectionDataInner)

	return err
}

type NullableGetCouponRelationshipCouponCodesListResponseCollectionDataInner struct {
	value *GetCouponRelationshipCouponCodesListResponseCollectionDataInner
	isSet bool
}

func (v NullableGetCouponRelationshipCouponCodesListResponseCollectionDataInner) Get() *GetCouponRelationshipCouponCodesListResponseCollectionDataInner {
	return v.value
}

func (v *NullableGetCouponRelationshipCouponCodesListResponseCollectionDataInner) Set(val *GetCouponRelationshipCouponCodesListResponseCollectionDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCouponRelationshipCouponCodesListResponseCollectionDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCouponRelationshipCouponCodesListResponseCollectionDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCouponRelationshipCouponCodesListResponseCollectionDataInner(val *GetCouponRelationshipCouponCodesListResponseCollectionDataInner) *NullableGetCouponRelationshipCouponCodesListResponseCollectionDataInner {
	return &NullableGetCouponRelationshipCouponCodesListResponseCollectionDataInner{value: val, isSet: true}
}

func (v NullableGetCouponRelationshipCouponCodesListResponseCollectionDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCouponRelationshipCouponCodesListResponseCollectionDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


