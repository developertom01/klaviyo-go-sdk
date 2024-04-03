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

// checks if the GetListRetrieveResponseCompoundDocumentDataAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetListRetrieveResponseCompoundDocumentDataAllOfAttributes{}

// GetListRetrieveResponseCompoundDocumentDataAllOfAttributes struct for GetListRetrieveResponseCompoundDocumentDataAllOfAttributes
type GetListRetrieveResponseCompoundDocumentDataAllOfAttributes struct {
	ProfileCount NullableInt32 `json:"profile_count,omitempty"`
}

// NewGetListRetrieveResponseCompoundDocumentDataAllOfAttributes instantiates a new GetListRetrieveResponseCompoundDocumentDataAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetListRetrieveResponseCompoundDocumentDataAllOfAttributes() *GetListRetrieveResponseCompoundDocumentDataAllOfAttributes {
	this := GetListRetrieveResponseCompoundDocumentDataAllOfAttributes{}
	return &this
}

// NewGetListRetrieveResponseCompoundDocumentDataAllOfAttributesWithDefaults instantiates a new GetListRetrieveResponseCompoundDocumentDataAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetListRetrieveResponseCompoundDocumentDataAllOfAttributesWithDefaults() *GetListRetrieveResponseCompoundDocumentDataAllOfAttributes {
	this := GetListRetrieveResponseCompoundDocumentDataAllOfAttributes{}
	return &this
}

// GetProfileCount returns the ProfileCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetListRetrieveResponseCompoundDocumentDataAllOfAttributes) GetProfileCount() int32 {
	if o == nil || IsNil(o.ProfileCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ProfileCount.Get()
}

// GetProfileCountOk returns a tuple with the ProfileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetListRetrieveResponseCompoundDocumentDataAllOfAttributes) GetProfileCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileCount.Get(), o.ProfileCount.IsSet()
}

// HasProfileCount returns a boolean if a field has been set.
func (o *GetListRetrieveResponseCompoundDocumentDataAllOfAttributes) HasProfileCount() bool {
	if o != nil && o.ProfileCount.IsSet() {
		return true
	}

	return false
}

// SetProfileCount gets a reference to the given NullableInt32 and assigns it to the ProfileCount field.
func (o *GetListRetrieveResponseCompoundDocumentDataAllOfAttributes) SetProfileCount(v int32) {
	o.ProfileCount.Set(&v)
}
// SetProfileCountNil sets the value for ProfileCount to be an explicit nil
func (o *GetListRetrieveResponseCompoundDocumentDataAllOfAttributes) SetProfileCountNil() {
	o.ProfileCount.Set(nil)
}

// UnsetProfileCount ensures that no value is present for ProfileCount, not even an explicit nil
func (o *GetListRetrieveResponseCompoundDocumentDataAllOfAttributes) UnsetProfileCount() {
	o.ProfileCount.Unset()
}

func (o GetListRetrieveResponseCompoundDocumentDataAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetListRetrieveResponseCompoundDocumentDataAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProfileCount.IsSet() {
		toSerialize["profile_count"] = o.ProfileCount.Get()
	}
	return toSerialize, nil
}

type NullableGetListRetrieveResponseCompoundDocumentDataAllOfAttributes struct {
	value *GetListRetrieveResponseCompoundDocumentDataAllOfAttributes
	isSet bool
}

func (v NullableGetListRetrieveResponseCompoundDocumentDataAllOfAttributes) Get() *GetListRetrieveResponseCompoundDocumentDataAllOfAttributes {
	return v.value
}

func (v *NullableGetListRetrieveResponseCompoundDocumentDataAllOfAttributes) Set(val *GetListRetrieveResponseCompoundDocumentDataAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetListRetrieveResponseCompoundDocumentDataAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetListRetrieveResponseCompoundDocumentDataAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetListRetrieveResponseCompoundDocumentDataAllOfAttributes(val *GetListRetrieveResponseCompoundDocumentDataAllOfAttributes) *NullableGetListRetrieveResponseCompoundDocumentDataAllOfAttributes {
	return &NullableGetListRetrieveResponseCompoundDocumentDataAllOfAttributes{value: val, isSet: true}
}

func (v NullableGetListRetrieveResponseCompoundDocumentDataAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetListRetrieveResponseCompoundDocumentDataAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


