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

// checks if the GetCouponCodeRelationshipCouponResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCouponCodeRelationshipCouponResponse{}

// GetCouponCodeRelationshipCouponResponse struct for GetCouponCodeRelationshipCouponResponse
type GetCouponCodeRelationshipCouponResponse struct {
	Data GetCouponCodeRelationshipCouponResponseData `json:"data"`
}

type _GetCouponCodeRelationshipCouponResponse GetCouponCodeRelationshipCouponResponse

// NewGetCouponCodeRelationshipCouponResponse instantiates a new GetCouponCodeRelationshipCouponResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCouponCodeRelationshipCouponResponse(data GetCouponCodeRelationshipCouponResponseData) *GetCouponCodeRelationshipCouponResponse {
	this := GetCouponCodeRelationshipCouponResponse{}
	this.Data = data
	return &this
}

// NewGetCouponCodeRelationshipCouponResponseWithDefaults instantiates a new GetCouponCodeRelationshipCouponResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCouponCodeRelationshipCouponResponseWithDefaults() *GetCouponCodeRelationshipCouponResponse {
	this := GetCouponCodeRelationshipCouponResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetCouponCodeRelationshipCouponResponse) GetData() GetCouponCodeRelationshipCouponResponseData {
	if o == nil {
		var ret GetCouponCodeRelationshipCouponResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCouponCodeRelationshipCouponResponse) GetDataOk() (*GetCouponCodeRelationshipCouponResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetCouponCodeRelationshipCouponResponse) SetData(v GetCouponCodeRelationshipCouponResponseData) {
	o.Data = v
}

func (o GetCouponCodeRelationshipCouponResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCouponCodeRelationshipCouponResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetCouponCodeRelationshipCouponResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetCouponCodeRelationshipCouponResponse := _GetCouponCodeRelationshipCouponResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCouponCodeRelationshipCouponResponse)

	if err != nil {
		return err
	}

	*o = GetCouponCodeRelationshipCouponResponse(varGetCouponCodeRelationshipCouponResponse)

	return err
}

type NullableGetCouponCodeRelationshipCouponResponse struct {
	value *GetCouponCodeRelationshipCouponResponse
	isSet bool
}

func (v NullableGetCouponCodeRelationshipCouponResponse) Get() *GetCouponCodeRelationshipCouponResponse {
	return v.value
}

func (v *NullableGetCouponCodeRelationshipCouponResponse) Set(val *GetCouponCodeRelationshipCouponResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCouponCodeRelationshipCouponResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCouponCodeRelationshipCouponResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCouponCodeRelationshipCouponResponse(val *GetCouponCodeRelationshipCouponResponse) *NullableGetCouponCodeRelationshipCouponResponse {
	return &NullableGetCouponCodeRelationshipCouponResponse{value: val, isSet: true}
}

func (v NullableGetCouponCodeRelationshipCouponResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCouponCodeRelationshipCouponResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

