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

// checks if the PostCatalogItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCatalogItemResponse{}

// PostCatalogItemResponse struct for PostCatalogItemResponse
type PostCatalogItemResponse struct {
	Data PostCatalogItemResponseData `json:"data"`
}

type _PostCatalogItemResponse PostCatalogItemResponse

// NewPostCatalogItemResponse instantiates a new PostCatalogItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCatalogItemResponse(data PostCatalogItemResponseData) *PostCatalogItemResponse {
	this := PostCatalogItemResponse{}
	this.Data = data
	return &this
}

// NewPostCatalogItemResponseWithDefaults instantiates a new PostCatalogItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCatalogItemResponseWithDefaults() *PostCatalogItemResponse {
	this := PostCatalogItemResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PostCatalogItemResponse) GetData() PostCatalogItemResponseData {
	if o == nil {
		var ret PostCatalogItemResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PostCatalogItemResponse) GetDataOk() (*PostCatalogItemResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PostCatalogItemResponse) SetData(v PostCatalogItemResponseData) {
	o.Data = v
}

func (o PostCatalogItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCatalogItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PostCatalogItemResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPostCatalogItemResponse := _PostCatalogItemResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCatalogItemResponse)

	if err != nil {
		return err
	}

	*o = PostCatalogItemResponse(varPostCatalogItemResponse)

	return err
}

type NullablePostCatalogItemResponse struct {
	value *PostCatalogItemResponse
	isSet bool
}

func (v NullablePostCatalogItemResponse) Get() *PostCatalogItemResponse {
	return v.value
}

func (v *NullablePostCatalogItemResponse) Set(val *PostCatalogItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCatalogItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCatalogItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCatalogItemResponse(val *PostCatalogItemResponse) *NullablePostCatalogItemResponse {
	return &NullablePostCatalogItemResponse{value: val, isSet: true}
}

func (v NullablePostCatalogItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCatalogItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

