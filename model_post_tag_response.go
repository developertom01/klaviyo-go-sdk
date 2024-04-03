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

// checks if the PostTagResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostTagResponse{}

// PostTagResponse struct for PostTagResponse
type PostTagResponse struct {
	Data PostTagResponseData `json:"data"`
}

type _PostTagResponse PostTagResponse

// NewPostTagResponse instantiates a new PostTagResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTagResponse(data PostTagResponseData) *PostTagResponse {
	this := PostTagResponse{}
	this.Data = data
	return &this
}

// NewPostTagResponseWithDefaults instantiates a new PostTagResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTagResponseWithDefaults() *PostTagResponse {
	this := PostTagResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PostTagResponse) GetData() PostTagResponseData {
	if o == nil {
		var ret PostTagResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PostTagResponse) GetDataOk() (*PostTagResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PostTagResponse) SetData(v PostTagResponseData) {
	o.Data = v
}

func (o PostTagResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostTagResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PostTagResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPostTagResponse := _PostTagResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostTagResponse)

	if err != nil {
		return err
	}

	*o = PostTagResponse(varPostTagResponse)

	return err
}

type NullablePostTagResponse struct {
	value *PostTagResponse
	isSet bool
}

func (v NullablePostTagResponse) Get() *PostTagResponse {
	return v.value
}

func (v *NullablePostTagResponse) Set(val *PostTagResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTagResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTagResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTagResponse(val *PostTagResponse) *NullablePostTagResponse {
	return &NullablePostTagResponse{value: val, isSet: true}
}

func (v NullablePostTagResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTagResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

