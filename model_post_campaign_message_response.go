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

// checks if the PostCampaignMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCampaignMessageResponse{}

// PostCampaignMessageResponse struct for PostCampaignMessageResponse
type PostCampaignMessageResponse struct {
	Data PostCampaignMessageResponseData `json:"data"`
}

type _PostCampaignMessageResponse PostCampaignMessageResponse

// NewPostCampaignMessageResponse instantiates a new PostCampaignMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCampaignMessageResponse(data PostCampaignMessageResponseData) *PostCampaignMessageResponse {
	this := PostCampaignMessageResponse{}
	this.Data = data
	return &this
}

// NewPostCampaignMessageResponseWithDefaults instantiates a new PostCampaignMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCampaignMessageResponseWithDefaults() *PostCampaignMessageResponse {
	this := PostCampaignMessageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PostCampaignMessageResponse) GetData() PostCampaignMessageResponseData {
	if o == nil {
		var ret PostCampaignMessageResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PostCampaignMessageResponse) GetDataOk() (*PostCampaignMessageResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PostCampaignMessageResponse) SetData(v PostCampaignMessageResponseData) {
	o.Data = v
}

func (o PostCampaignMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCampaignMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PostCampaignMessageResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPostCampaignMessageResponse := _PostCampaignMessageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCampaignMessageResponse)

	if err != nil {
		return err
	}

	*o = PostCampaignMessageResponse(varPostCampaignMessageResponse)

	return err
}

type NullablePostCampaignMessageResponse struct {
	value *PostCampaignMessageResponse
	isSet bool
}

func (v NullablePostCampaignMessageResponse) Get() *PostCampaignMessageResponse {
	return v.value
}

func (v *NullablePostCampaignMessageResponse) Set(val *PostCampaignMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCampaignMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCampaignMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCampaignMessageResponse(val *PostCampaignMessageResponse) *NullablePostCampaignMessageResponse {
	return &NullablePostCampaignMessageResponse{value: val, isSet: true}
}

func (v NullablePostCampaignMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCampaignMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


