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

// checks if the GetTagGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTagGroupResponse{}

// GetTagGroupResponse struct for GetTagGroupResponse
type GetTagGroupResponse struct {
	Data GetTagGroupResponseCollectionDataInner `json:"data"`
}

type _GetTagGroupResponse GetTagGroupResponse

// NewGetTagGroupResponse instantiates a new GetTagGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTagGroupResponse(data GetTagGroupResponseCollectionDataInner) *GetTagGroupResponse {
	this := GetTagGroupResponse{}
	this.Data = data
	return &this
}

// NewGetTagGroupResponseWithDefaults instantiates a new GetTagGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTagGroupResponseWithDefaults() *GetTagGroupResponse {
	this := GetTagGroupResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetTagGroupResponse) GetData() GetTagGroupResponseCollectionDataInner {
	if o == nil {
		var ret GetTagGroupResponseCollectionDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetTagGroupResponse) GetDataOk() (*GetTagGroupResponseCollectionDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetTagGroupResponse) SetData(v GetTagGroupResponseCollectionDataInner) {
	o.Data = v
}

func (o GetTagGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTagGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetTagGroupResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetTagGroupResponse := _GetTagGroupResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTagGroupResponse)

	if err != nil {
		return err
	}

	*o = GetTagGroupResponse(varGetTagGroupResponse)

	return err
}

type NullableGetTagGroupResponse struct {
	value *GetTagGroupResponse
	isSet bool
}

func (v NullableGetTagGroupResponse) Get() *GetTagGroupResponse {
	return v.value
}

func (v *NullableGetTagGroupResponse) Set(val *GetTagGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTagGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTagGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTagGroupResponse(val *GetTagGroupResponse) *NullableGetTagGroupResponse {
	return &NullableGetTagGroupResponse{value: val, isSet: true}
}

func (v NullableGetTagGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTagGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


