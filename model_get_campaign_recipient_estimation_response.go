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

// checks if the GetCampaignRecipientEstimationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaignRecipientEstimationResponse{}

// GetCampaignRecipientEstimationResponse struct for GetCampaignRecipientEstimationResponse
type GetCampaignRecipientEstimationResponse struct {
	Data CampaignRecipientEstimationResponseObjectResource `json:"data"`
}

type _GetCampaignRecipientEstimationResponse GetCampaignRecipientEstimationResponse

// NewGetCampaignRecipientEstimationResponse instantiates a new GetCampaignRecipientEstimationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaignRecipientEstimationResponse(data CampaignRecipientEstimationResponseObjectResource) *GetCampaignRecipientEstimationResponse {
	this := GetCampaignRecipientEstimationResponse{}
	this.Data = data
	return &this
}

// NewGetCampaignRecipientEstimationResponseWithDefaults instantiates a new GetCampaignRecipientEstimationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaignRecipientEstimationResponseWithDefaults() *GetCampaignRecipientEstimationResponse {
	this := GetCampaignRecipientEstimationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetCampaignRecipientEstimationResponse) GetData() CampaignRecipientEstimationResponseObjectResource {
	if o == nil {
		var ret CampaignRecipientEstimationResponseObjectResource
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCampaignRecipientEstimationResponse) GetDataOk() (*CampaignRecipientEstimationResponseObjectResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetCampaignRecipientEstimationResponse) SetData(v CampaignRecipientEstimationResponseObjectResource) {
	o.Data = v
}

func (o GetCampaignRecipientEstimationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCampaignRecipientEstimationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetCampaignRecipientEstimationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetCampaignRecipientEstimationResponse := _GetCampaignRecipientEstimationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCampaignRecipientEstimationResponse)

	if err != nil {
		return err
	}

	*o = GetCampaignRecipientEstimationResponse(varGetCampaignRecipientEstimationResponse)

	return err
}

type NullableGetCampaignRecipientEstimationResponse struct {
	value *GetCampaignRecipientEstimationResponse
	isSet bool
}

func (v NullableGetCampaignRecipientEstimationResponse) Get() *GetCampaignRecipientEstimationResponse {
	return v.value
}

func (v *NullableGetCampaignRecipientEstimationResponse) Set(val *GetCampaignRecipientEstimationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaignRecipientEstimationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaignRecipientEstimationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaignRecipientEstimationResponse(val *GetCampaignRecipientEstimationResponse) *NullableGetCampaignRecipientEstimationResponse {
	return &NullableGetCampaignRecipientEstimationResponse{value: val, isSet: true}
}

func (v NullableGetCampaignRecipientEstimationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCampaignRecipientEstimationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


