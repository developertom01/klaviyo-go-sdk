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
)

// checks if the GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships{}

// GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships struct for GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships
type GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships struct {
	Coupon *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCoupon `json:"coupon,omitempty"`
	Profile *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile `json:"profile,omitempty"`
}

// NewGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships instantiates a new GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships() *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	this := GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships{}
	return &this
}

// NewGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsWithDefaults instantiates a new GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsWithDefaults() *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	this := GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships{}
	return &this
}

// GetCoupon returns the Coupon field value if set, zero value otherwise.
func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetCoupon() GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCoupon {
	if o == nil || IsNil(o.Coupon) {
		var ret GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCoupon
		return ret
	}
	return *o.Coupon
}

// GetCouponOk returns a tuple with the Coupon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetCouponOk() (*GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCoupon, bool) {
	if o == nil || IsNil(o.Coupon) {
		return nil, false
	}
	return o.Coupon, true
}

// HasCoupon returns a boolean if a field has been set.
func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) HasCoupon() bool {
	if o != nil && !IsNil(o.Coupon) {
		return true
	}

	return false
}

// SetCoupon gets a reference to the given GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCoupon and assigns it to the Coupon field.
func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) SetCoupon(v GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCoupon) {
	o.Coupon = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetProfile() GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile {
	if o == nil || IsNil(o.Profile) {
		var ret GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetProfileOk() (*GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile and assigns it to the Profile field.
func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) SetProfile(v GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) {
	o.Profile = &v
}

func (o GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Coupon) {
		toSerialize["coupon"] = o.Coupon
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	return toSerialize, nil
}

type NullableGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships struct {
	value *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships
	isSet bool
}

func (v NullableGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) Get() *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	return v.value
}

func (v *NullableGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) Set(val *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships(val *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) *NullableGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	return &NullableGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships{value: val, isSet: true}
}

func (v NullableGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


