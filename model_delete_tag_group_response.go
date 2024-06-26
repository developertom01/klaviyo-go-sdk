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

// checks if the DeleteTagGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteTagGroupResponse{}

// DeleteTagGroupResponse struct for DeleteTagGroupResponse
type DeleteTagGroupResponse struct {
	Data PostTagGroupResponseData `json:"data"`
}

type _DeleteTagGroupResponse DeleteTagGroupResponse

// NewDeleteTagGroupResponse instantiates a new DeleteTagGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTagGroupResponse(data PostTagGroupResponseData) *DeleteTagGroupResponse {
	this := DeleteTagGroupResponse{}
	this.Data = data
	return &this
}

// NewDeleteTagGroupResponseWithDefaults instantiates a new DeleteTagGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTagGroupResponseWithDefaults() *DeleteTagGroupResponse {
	this := DeleteTagGroupResponse{}
	return &this
}

// GetData returns the Data field value
func (o *DeleteTagGroupResponse) GetData() PostTagGroupResponseData {
	if o == nil {
		var ret PostTagGroupResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeleteTagGroupResponse) GetDataOk() (*PostTagGroupResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeleteTagGroupResponse) SetData(v PostTagGroupResponseData) {
	o.Data = v
}

func (o DeleteTagGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTagGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *DeleteTagGroupResponse) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteTagGroupResponse := _DeleteTagGroupResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteTagGroupResponse)

	if err != nil {
		return err
	}

	*o = DeleteTagGroupResponse(varDeleteTagGroupResponse)

	return err
}

type NullableDeleteTagGroupResponse struct {
	value *DeleteTagGroupResponse
	isSet bool
}

func (v NullableDeleteTagGroupResponse) Get() *DeleteTagGroupResponse {
	return v.value
}

func (v *NullableDeleteTagGroupResponse) Set(val *DeleteTagGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTagGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTagGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTagGroupResponse(val *DeleteTagGroupResponse) *NullableDeleteTagGroupResponse {
	return &NullableDeleteTagGroupResponse{value: val, isSet: true}
}

func (v NullableDeleteTagGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTagGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


