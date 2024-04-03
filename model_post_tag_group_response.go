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

// checks if the PostTagGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostTagGroupResponse{}

// PostTagGroupResponse struct for PostTagGroupResponse
type PostTagGroupResponse struct {
	Data PostTagGroupResponseData `json:"data"`
}

type _PostTagGroupResponse PostTagGroupResponse

// NewPostTagGroupResponse instantiates a new PostTagGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTagGroupResponse(data PostTagGroupResponseData) *PostTagGroupResponse {
	this := PostTagGroupResponse{}
	this.Data = data
	return &this
}

// NewPostTagGroupResponseWithDefaults instantiates a new PostTagGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTagGroupResponseWithDefaults() *PostTagGroupResponse {
	this := PostTagGroupResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PostTagGroupResponse) GetData() PostTagGroupResponseData {
	if o == nil {
		var ret PostTagGroupResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PostTagGroupResponse) GetDataOk() (*PostTagGroupResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PostTagGroupResponse) SetData(v PostTagGroupResponseData) {
	o.Data = v
}

func (o PostTagGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostTagGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PostTagGroupResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPostTagGroupResponse := _PostTagGroupResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostTagGroupResponse)

	if err != nil {
		return err
	}

	*o = PostTagGroupResponse(varPostTagGroupResponse)

	return err
}

type NullablePostTagGroupResponse struct {
	value *PostTagGroupResponse
	isSet bool
}

func (v NullablePostTagGroupResponse) Get() *PostTagGroupResponse {
	return v.value
}

func (v *NullablePostTagGroupResponse) Set(val *PostTagGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTagGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTagGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTagGroupResponse(val *PostTagGroupResponse) *NullablePostTagGroupResponse {
	return &NullablePostTagGroupResponse{value: val, isSet: true}
}

func (v NullablePostTagGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTagGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


