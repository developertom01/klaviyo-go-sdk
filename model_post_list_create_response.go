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

// checks if the PostListCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostListCreateResponse{}

// PostListCreateResponse struct for PostListCreateResponse
type PostListCreateResponse struct {
	Data PostListCreateResponseData `json:"data"`
}

type _PostListCreateResponse PostListCreateResponse

// NewPostListCreateResponse instantiates a new PostListCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostListCreateResponse(data PostListCreateResponseData) *PostListCreateResponse {
	this := PostListCreateResponse{}
	this.Data = data
	return &this
}

// NewPostListCreateResponseWithDefaults instantiates a new PostListCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostListCreateResponseWithDefaults() *PostListCreateResponse {
	this := PostListCreateResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PostListCreateResponse) GetData() PostListCreateResponseData {
	if o == nil {
		var ret PostListCreateResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PostListCreateResponse) GetDataOk() (*PostListCreateResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PostListCreateResponse) SetData(v PostListCreateResponseData) {
	o.Data = v
}

func (o PostListCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostListCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PostListCreateResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPostListCreateResponse := _PostListCreateResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostListCreateResponse)

	if err != nil {
		return err
	}

	*o = PostListCreateResponse(varPostListCreateResponse)

	return err
}

type NullablePostListCreateResponse struct {
	value *PostListCreateResponse
	isSet bool
}

func (v NullablePostListCreateResponse) Get() *PostListCreateResponse {
	return v.value
}

func (v *NullablePostListCreateResponse) Set(val *PostListCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostListCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostListCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostListCreateResponse(val *PostListCreateResponse) *NullablePostListCreateResponse {
	return &NullablePostListCreateResponse{value: val, isSet: true}
}

func (v NullablePostListCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostListCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

