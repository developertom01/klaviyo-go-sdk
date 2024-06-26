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

// checks if the GetImageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetImageResponse{}

// GetImageResponse struct for GetImageResponse
type GetImageResponse struct {
	Data ImageResponseObjectResource `json:"data"`
}

type _GetImageResponse GetImageResponse

// NewGetImageResponse instantiates a new GetImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetImageResponse(data ImageResponseObjectResource) *GetImageResponse {
	this := GetImageResponse{}
	this.Data = data
	return &this
}

// NewGetImageResponseWithDefaults instantiates a new GetImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetImageResponseWithDefaults() *GetImageResponse {
	this := GetImageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetImageResponse) GetData() ImageResponseObjectResource {
	if o == nil {
		var ret ImageResponseObjectResource
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetImageResponse) GetDataOk() (*ImageResponseObjectResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetImageResponse) SetData(v ImageResponseObjectResource) {
	o.Data = v
}

func (o GetImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetImageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetImageResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetImageResponse := _GetImageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetImageResponse)

	if err != nil {
		return err
	}

	*o = GetImageResponse(varGetImageResponse)

	return err
}

type NullableGetImageResponse struct {
	value *GetImageResponse
	isSet bool
}

func (v NullableGetImageResponse) Get() *GetImageResponse {
	return v.value
}

func (v *NullableGetImageResponse) Set(val *GetImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetImageResponse(val *GetImageResponse) *NullableGetImageResponse {
	return &NullableGetImageResponse{value: val, isSet: true}
}

func (v NullableGetImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


