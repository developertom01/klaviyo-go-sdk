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

// checks if the PostCouponResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCouponResponse{}

// PostCouponResponse struct for PostCouponResponse
type PostCouponResponse struct {
	Data PostCouponResponseData `json:"data"`
}

type _PostCouponResponse PostCouponResponse

// NewPostCouponResponse instantiates a new PostCouponResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCouponResponse(data PostCouponResponseData) *PostCouponResponse {
	this := PostCouponResponse{}
	this.Data = data
	return &this
}

// NewPostCouponResponseWithDefaults instantiates a new PostCouponResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCouponResponseWithDefaults() *PostCouponResponse {
	this := PostCouponResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PostCouponResponse) GetData() PostCouponResponseData {
	if o == nil {
		var ret PostCouponResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PostCouponResponse) GetDataOk() (*PostCouponResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PostCouponResponse) SetData(v PostCouponResponseData) {
	o.Data = v
}

func (o PostCouponResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCouponResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PostCouponResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPostCouponResponse := _PostCouponResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCouponResponse)

	if err != nil {
		return err
	}

	*o = PostCouponResponse(varPostCouponResponse)

	return err
}

type NullablePostCouponResponse struct {
	value *PostCouponResponse
	isSet bool
}

func (v NullablePostCouponResponse) Get() *PostCouponResponse {
	return v.value
}

func (v *NullablePostCouponResponse) Set(val *PostCouponResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCouponResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCouponResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCouponResponse(val *PostCouponResponse) *NullablePostCouponResponse {
	return &NullablePostCouponResponse{value: val, isSet: true}
}

func (v NullablePostCouponResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCouponResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


