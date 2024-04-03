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

// checks if the PatchFlowResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchFlowResponse{}

// PatchFlowResponse struct for PatchFlowResponse
type PatchFlowResponse struct {
	Data PatchFlowResponseData `json:"data"`
}

type _PatchFlowResponse PatchFlowResponse

// NewPatchFlowResponse instantiates a new PatchFlowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchFlowResponse(data PatchFlowResponseData) *PatchFlowResponse {
	this := PatchFlowResponse{}
	this.Data = data
	return &this
}

// NewPatchFlowResponseWithDefaults instantiates a new PatchFlowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchFlowResponseWithDefaults() *PatchFlowResponse {
	this := PatchFlowResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PatchFlowResponse) GetData() PatchFlowResponseData {
	if o == nil {
		var ret PatchFlowResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PatchFlowResponse) GetDataOk() (*PatchFlowResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PatchFlowResponse) SetData(v PatchFlowResponseData) {
	o.Data = v
}

func (o PatchFlowResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchFlowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PatchFlowResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPatchFlowResponse := _PatchFlowResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatchFlowResponse)

	if err != nil {
		return err
	}

	*o = PatchFlowResponse(varPatchFlowResponse)

	return err
}

type NullablePatchFlowResponse struct {
	value *PatchFlowResponse
	isSet bool
}

func (v NullablePatchFlowResponse) Get() *PatchFlowResponse {
	return v.value
}

func (v *NullablePatchFlowResponse) Set(val *PatchFlowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchFlowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchFlowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchFlowResponse(val *PatchFlowResponse) *NullablePatchFlowResponse {
	return &NullablePatchFlowResponse{value: val, isSet: true}
}

func (v NullablePatchFlowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchFlowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


