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

// checks if the GetProfileResponseDataAllOfRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProfileResponseDataAllOfRelationships{}

// GetProfileResponseDataAllOfRelationships struct for GetProfileResponseDataAllOfRelationships
type GetProfileResponseDataAllOfRelationships struct {
	Lists *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile `json:"lists,omitempty"`
	Segments *GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile `json:"segments,omitempty"`
}

// NewGetProfileResponseDataAllOfRelationships instantiates a new GetProfileResponseDataAllOfRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProfileResponseDataAllOfRelationships() *GetProfileResponseDataAllOfRelationships {
	this := GetProfileResponseDataAllOfRelationships{}
	return &this
}

// NewGetProfileResponseDataAllOfRelationshipsWithDefaults instantiates a new GetProfileResponseDataAllOfRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProfileResponseDataAllOfRelationshipsWithDefaults() *GetProfileResponseDataAllOfRelationships {
	this := GetProfileResponseDataAllOfRelationships{}
	return &this
}

// GetLists returns the Lists field value if set, zero value otherwise.
func (o *GetProfileResponseDataAllOfRelationships) GetLists() GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile {
	if o == nil || IsNil(o.Lists) {
		var ret GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile
		return ret
	}
	return *o.Lists
}

// GetListsOk returns a tuple with the Lists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfileResponseDataAllOfRelationships) GetListsOk() (*GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile, bool) {
	if o == nil || IsNil(o.Lists) {
		return nil, false
	}
	return o.Lists, true
}

// HasLists returns a boolean if a field has been set.
func (o *GetProfileResponseDataAllOfRelationships) HasLists() bool {
	if o != nil && !IsNil(o.Lists) {
		return true
	}

	return false
}

// SetLists gets a reference to the given GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile and assigns it to the Lists field.
func (o *GetProfileResponseDataAllOfRelationships) SetLists(v GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) {
	o.Lists = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *GetProfileResponseDataAllOfRelationships) GetSegments() GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile {
	if o == nil || IsNil(o.Segments) {
		var ret GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile
		return ret
	}
	return *o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfileResponseDataAllOfRelationships) GetSegmentsOk() (*GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *GetProfileResponseDataAllOfRelationships) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile and assigns it to the Segments field.
func (o *GetProfileResponseDataAllOfRelationships) SetSegments(v GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) {
	o.Segments = &v
}

func (o GetProfileResponseDataAllOfRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProfileResponseDataAllOfRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lists) {
		toSerialize["lists"] = o.Lists
	}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	return toSerialize, nil
}

type NullableGetProfileResponseDataAllOfRelationships struct {
	value *GetProfileResponseDataAllOfRelationships
	isSet bool
}

func (v NullableGetProfileResponseDataAllOfRelationships) Get() *GetProfileResponseDataAllOfRelationships {
	return v.value
}

func (v *NullableGetProfileResponseDataAllOfRelationships) Set(val *GetProfileResponseDataAllOfRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProfileResponseDataAllOfRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProfileResponseDataAllOfRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProfileResponseDataAllOfRelationships(val *GetProfileResponseDataAllOfRelationships) *NullableGetProfileResponseDataAllOfRelationships {
	return &NullableGetProfileResponseDataAllOfRelationships{value: val, isSet: true}
}

func (v NullableGetProfileResponseDataAllOfRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProfileResponseDataAllOfRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

