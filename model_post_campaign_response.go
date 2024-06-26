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

// checks if the PostCampaignResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCampaignResponse{}

// PostCampaignResponse struct for PostCampaignResponse
type PostCampaignResponse struct {
	Data PostCampaignResponseData `json:"data"`
}

type _PostCampaignResponse PostCampaignResponse

// NewPostCampaignResponse instantiates a new PostCampaignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCampaignResponse(data PostCampaignResponseData) *PostCampaignResponse {
	this := PostCampaignResponse{}
	this.Data = data
	return &this
}

// NewPostCampaignResponseWithDefaults instantiates a new PostCampaignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCampaignResponseWithDefaults() *PostCampaignResponse {
	this := PostCampaignResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PostCampaignResponse) GetData() PostCampaignResponseData {
	if o == nil {
		var ret PostCampaignResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PostCampaignResponse) GetDataOk() (*PostCampaignResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PostCampaignResponse) SetData(v PostCampaignResponseData) {
	o.Data = v
}

func (o PostCampaignResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCampaignResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PostCampaignResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPostCampaignResponse := _PostCampaignResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCampaignResponse)

	if err != nil {
		return err
	}

	*o = PostCampaignResponse(varPostCampaignResponse)

	return err
}

type NullablePostCampaignResponse struct {
	value *PostCampaignResponse
	isSet bool
}

func (v NullablePostCampaignResponse) Get() *PostCampaignResponse {
	return v.value
}

func (v *NullablePostCampaignResponse) Set(val *PostCampaignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCampaignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCampaignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCampaignResponse(val *PostCampaignResponse) *NullablePostCampaignResponse {
	return &NullablePostCampaignResponse{value: val, isSet: true}
}

func (v NullablePostCampaignResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCampaignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


