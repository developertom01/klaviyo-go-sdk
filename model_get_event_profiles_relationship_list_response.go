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

// checks if the GetEventProfilesRelationshipListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEventProfilesRelationshipListResponse{}

// GetEventProfilesRelationshipListResponse struct for GetEventProfilesRelationshipListResponse
type GetEventProfilesRelationshipListResponse struct {
	Data GetEventProfilesRelationshipListResponseData `json:"data"`
}

type _GetEventProfilesRelationshipListResponse GetEventProfilesRelationshipListResponse

// NewGetEventProfilesRelationshipListResponse instantiates a new GetEventProfilesRelationshipListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventProfilesRelationshipListResponse(data GetEventProfilesRelationshipListResponseData) *GetEventProfilesRelationshipListResponse {
	this := GetEventProfilesRelationshipListResponse{}
	this.Data = data
	return &this
}

// NewGetEventProfilesRelationshipListResponseWithDefaults instantiates a new GetEventProfilesRelationshipListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventProfilesRelationshipListResponseWithDefaults() *GetEventProfilesRelationshipListResponse {
	this := GetEventProfilesRelationshipListResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetEventProfilesRelationshipListResponse) GetData() GetEventProfilesRelationshipListResponseData {
	if o == nil {
		var ret GetEventProfilesRelationshipListResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetEventProfilesRelationshipListResponse) GetDataOk() (*GetEventProfilesRelationshipListResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetEventProfilesRelationshipListResponse) SetData(v GetEventProfilesRelationshipListResponseData) {
	o.Data = v
}

func (o GetEventProfilesRelationshipListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEventProfilesRelationshipListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetEventProfilesRelationshipListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetEventProfilesRelationshipListResponse := _GetEventProfilesRelationshipListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEventProfilesRelationshipListResponse)

	if err != nil {
		return err
	}

	*o = GetEventProfilesRelationshipListResponse(varGetEventProfilesRelationshipListResponse)

	return err
}

type NullableGetEventProfilesRelationshipListResponse struct {
	value *GetEventProfilesRelationshipListResponse
	isSet bool
}

func (v NullableGetEventProfilesRelationshipListResponse) Get() *GetEventProfilesRelationshipListResponse {
	return v.value
}

func (v *NullableGetEventProfilesRelationshipListResponse) Set(val *GetEventProfilesRelationshipListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventProfilesRelationshipListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventProfilesRelationshipListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventProfilesRelationshipListResponse(val *GetEventProfilesRelationshipListResponse) *NullableGetEventProfilesRelationshipListResponse {
	return &NullableGetEventProfilesRelationshipListResponse{value: val, isSet: true}
}

func (v NullableGetEventProfilesRelationshipListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventProfilesRelationshipListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

