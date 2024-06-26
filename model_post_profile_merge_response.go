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

// checks if the PostProfileMergeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostProfileMergeResponse{}

// PostProfileMergeResponse struct for PostProfileMergeResponse
type PostProfileMergeResponse struct {
	Data PostProfileMergeResponseData `json:"data"`
}

type _PostProfileMergeResponse PostProfileMergeResponse

// NewPostProfileMergeResponse instantiates a new PostProfileMergeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostProfileMergeResponse(data PostProfileMergeResponseData) *PostProfileMergeResponse {
	this := PostProfileMergeResponse{}
	this.Data = data
	return &this
}

// NewPostProfileMergeResponseWithDefaults instantiates a new PostProfileMergeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostProfileMergeResponseWithDefaults() *PostProfileMergeResponse {
	this := PostProfileMergeResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PostProfileMergeResponse) GetData() PostProfileMergeResponseData {
	if o == nil {
		var ret PostProfileMergeResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PostProfileMergeResponse) GetDataOk() (*PostProfileMergeResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PostProfileMergeResponse) SetData(v PostProfileMergeResponseData) {
	o.Data = v
}

func (o PostProfileMergeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostProfileMergeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PostProfileMergeResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPostProfileMergeResponse := _PostProfileMergeResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostProfileMergeResponse)

	if err != nil {
		return err
	}

	*o = PostProfileMergeResponse(varPostProfileMergeResponse)

	return err
}

type NullablePostProfileMergeResponse struct {
	value *PostProfileMergeResponse
	isSet bool
}

func (v NullablePostProfileMergeResponse) Get() *PostProfileMergeResponse {
	return v.value
}

func (v *NullablePostProfileMergeResponse) Set(val *PostProfileMergeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostProfileMergeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostProfileMergeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostProfileMergeResponse(val *PostProfileMergeResponse) *NullablePostProfileMergeResponse {
	return &NullablePostProfileMergeResponse{value: val, isSet: true}
}

func (v NullablePostProfileMergeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostProfileMergeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


