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

// checks if the PostProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostProfileResponse{}

// PostProfileResponse struct for PostProfileResponse
type PostProfileResponse struct {
	Data PostProfileResponseData `json:"data"`
}

type _PostProfileResponse PostProfileResponse

// NewPostProfileResponse instantiates a new PostProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostProfileResponse(data PostProfileResponseData) *PostProfileResponse {
	this := PostProfileResponse{}
	this.Data = data
	return &this
}

// NewPostProfileResponseWithDefaults instantiates a new PostProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostProfileResponseWithDefaults() *PostProfileResponse {
	this := PostProfileResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PostProfileResponse) GetData() PostProfileResponseData {
	if o == nil {
		var ret PostProfileResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PostProfileResponse) GetDataOk() (*PostProfileResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PostProfileResponse) SetData(v PostProfileResponseData) {
	o.Data = v
}

func (o PostProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PostProfileResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPostProfileResponse := _PostProfileResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostProfileResponse)

	if err != nil {
		return err
	}

	*o = PostProfileResponse(varPostProfileResponse)

	return err
}

type NullablePostProfileResponse struct {
	value *PostProfileResponse
	isSet bool
}

func (v NullablePostProfileResponse) Get() *PostProfileResponse {
	return v.value
}

func (v *NullablePostProfileResponse) Set(val *PostProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostProfileResponse(val *PostProfileResponse) *NullablePostProfileResponse {
	return &NullablePostProfileResponse{value: val, isSet: true}
}

func (v NullablePostProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


