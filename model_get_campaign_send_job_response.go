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

// checks if the GetCampaignSendJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaignSendJobResponse{}

// GetCampaignSendJobResponse struct for GetCampaignSendJobResponse
type GetCampaignSendJobResponse struct {
	Data CampaignSendJobResponseObjectResource `json:"data"`
}

type _GetCampaignSendJobResponse GetCampaignSendJobResponse

// NewGetCampaignSendJobResponse instantiates a new GetCampaignSendJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaignSendJobResponse(data CampaignSendJobResponseObjectResource) *GetCampaignSendJobResponse {
	this := GetCampaignSendJobResponse{}
	this.Data = data
	return &this
}

// NewGetCampaignSendJobResponseWithDefaults instantiates a new GetCampaignSendJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaignSendJobResponseWithDefaults() *GetCampaignSendJobResponse {
	this := GetCampaignSendJobResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetCampaignSendJobResponse) GetData() CampaignSendJobResponseObjectResource {
	if o == nil {
		var ret CampaignSendJobResponseObjectResource
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCampaignSendJobResponse) GetDataOk() (*CampaignSendJobResponseObjectResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetCampaignSendJobResponse) SetData(v CampaignSendJobResponseObjectResource) {
	o.Data = v
}

func (o GetCampaignSendJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCampaignSendJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetCampaignSendJobResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetCampaignSendJobResponse := _GetCampaignSendJobResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCampaignSendJobResponse)

	if err != nil {
		return err
	}

	*o = GetCampaignSendJobResponse(varGetCampaignSendJobResponse)

	return err
}

type NullableGetCampaignSendJobResponse struct {
	value *GetCampaignSendJobResponse
	isSet bool
}

func (v NullableGetCampaignSendJobResponse) Get() *GetCampaignSendJobResponse {
	return v.value
}

func (v *NullableGetCampaignSendJobResponse) Set(val *GetCampaignSendJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaignSendJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaignSendJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaignSendJobResponse(val *GetCampaignSendJobResponse) *NullableGetCampaignSendJobResponse {
	return &NullableGetCampaignSendJobResponse{value: val, isSet: true}
}

func (v NullableGetCampaignSendJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCampaignSendJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


