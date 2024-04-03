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

// checks if the GetCampaignMessageCampaignRelationshipListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaignMessageCampaignRelationshipListResponse{}

// GetCampaignMessageCampaignRelationshipListResponse struct for GetCampaignMessageCampaignRelationshipListResponse
type GetCampaignMessageCampaignRelationshipListResponse struct {
	Data GetCampaignMessageCampaignRelationshipListResponseData `json:"data"`
}

type _GetCampaignMessageCampaignRelationshipListResponse GetCampaignMessageCampaignRelationshipListResponse

// NewGetCampaignMessageCampaignRelationshipListResponse instantiates a new GetCampaignMessageCampaignRelationshipListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaignMessageCampaignRelationshipListResponse(data GetCampaignMessageCampaignRelationshipListResponseData) *GetCampaignMessageCampaignRelationshipListResponse {
	this := GetCampaignMessageCampaignRelationshipListResponse{}
	this.Data = data
	return &this
}

// NewGetCampaignMessageCampaignRelationshipListResponseWithDefaults instantiates a new GetCampaignMessageCampaignRelationshipListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaignMessageCampaignRelationshipListResponseWithDefaults() *GetCampaignMessageCampaignRelationshipListResponse {
	this := GetCampaignMessageCampaignRelationshipListResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetCampaignMessageCampaignRelationshipListResponse) GetData() GetCampaignMessageCampaignRelationshipListResponseData {
	if o == nil {
		var ret GetCampaignMessageCampaignRelationshipListResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCampaignMessageCampaignRelationshipListResponse) GetDataOk() (*GetCampaignMessageCampaignRelationshipListResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetCampaignMessageCampaignRelationshipListResponse) SetData(v GetCampaignMessageCampaignRelationshipListResponseData) {
	o.Data = v
}

func (o GetCampaignMessageCampaignRelationshipListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCampaignMessageCampaignRelationshipListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetCampaignMessageCampaignRelationshipListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetCampaignMessageCampaignRelationshipListResponse := _GetCampaignMessageCampaignRelationshipListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCampaignMessageCampaignRelationshipListResponse)

	if err != nil {
		return err
	}

	*o = GetCampaignMessageCampaignRelationshipListResponse(varGetCampaignMessageCampaignRelationshipListResponse)

	return err
}

type NullableGetCampaignMessageCampaignRelationshipListResponse struct {
	value *GetCampaignMessageCampaignRelationshipListResponse
	isSet bool
}

func (v NullableGetCampaignMessageCampaignRelationshipListResponse) Get() *GetCampaignMessageCampaignRelationshipListResponse {
	return v.value
}

func (v *NullableGetCampaignMessageCampaignRelationshipListResponse) Set(val *GetCampaignMessageCampaignRelationshipListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaignMessageCampaignRelationshipListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaignMessageCampaignRelationshipListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaignMessageCampaignRelationshipListResponse(val *GetCampaignMessageCampaignRelationshipListResponse) *NullableGetCampaignMessageCampaignRelationshipListResponse {
	return &NullableGetCampaignMessageCampaignRelationshipListResponse{value: val, isSet: true}
}

func (v NullableGetCampaignMessageCampaignRelationshipListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCampaignMessageCampaignRelationshipListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


