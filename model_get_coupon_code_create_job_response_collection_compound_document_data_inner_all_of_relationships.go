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

// checks if the GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships{}

// GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships struct for GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships
type GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships struct {
	CouponCodes *GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponCodes `json:"coupon-codes,omitempty"`
}

// NewGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships instantiates a new GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships() *GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	this := GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships{}
	return &this
}

// NewGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsWithDefaults instantiates a new GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsWithDefaults() *GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	this := GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships{}
	return &this
}

// GetCouponCodes returns the CouponCodes field value if set, zero value otherwise.
func (o *GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetCouponCodes() GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponCodes {
	if o == nil || IsNil(o.CouponCodes) {
		var ret GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponCodes
		return ret
	}
	return *o.CouponCodes
}

// GetCouponCodesOk returns a tuple with the CouponCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetCouponCodesOk() (*GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponCodes, bool) {
	if o == nil || IsNil(o.CouponCodes) {
		return nil, false
	}
	return o.CouponCodes, true
}

// HasCouponCodes returns a boolean if a field has been set.
func (o *GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) HasCouponCodes() bool {
	if o != nil && !IsNil(o.CouponCodes) {
		return true
	}

	return false
}

// SetCouponCodes gets a reference to the given GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponCodes and assigns it to the CouponCodes field.
func (o *GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) SetCouponCodes(v GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponCodes) {
	o.CouponCodes = &v
}

func (o GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CouponCodes) {
		toSerialize["coupon-codes"] = o.CouponCodes
	}
	return toSerialize, nil
}

type NullableGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships struct {
	value *GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships
	isSet bool
}

func (v NullableGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) Get() *GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	return v.value
}

func (v *NullableGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) Set(val *GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships(val *GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) *NullableGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	return &NullableGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships{value: val, isSet: true}
}

func (v NullableGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


