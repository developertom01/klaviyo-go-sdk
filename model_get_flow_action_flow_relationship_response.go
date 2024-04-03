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

// checks if the GetFlowActionFlowRelationshipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFlowActionFlowRelationshipResponse{}

// GetFlowActionFlowRelationshipResponse struct for GetFlowActionFlowRelationshipResponse
type GetFlowActionFlowRelationshipResponse struct {
	Data GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlowData `json:"data"`
}

type _GetFlowActionFlowRelationshipResponse GetFlowActionFlowRelationshipResponse

// NewGetFlowActionFlowRelationshipResponse instantiates a new GetFlowActionFlowRelationshipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlowActionFlowRelationshipResponse(data GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlowData) *GetFlowActionFlowRelationshipResponse {
	this := GetFlowActionFlowRelationshipResponse{}
	this.Data = data
	return &this
}

// NewGetFlowActionFlowRelationshipResponseWithDefaults instantiates a new GetFlowActionFlowRelationshipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlowActionFlowRelationshipResponseWithDefaults() *GetFlowActionFlowRelationshipResponse {
	this := GetFlowActionFlowRelationshipResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetFlowActionFlowRelationshipResponse) GetData() GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlowData {
	if o == nil {
		var ret GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlowData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetFlowActionFlowRelationshipResponse) GetDataOk() (*GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlowData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetFlowActionFlowRelationshipResponse) SetData(v GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlowData) {
	o.Data = v
}

func (o GetFlowActionFlowRelationshipResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlowActionFlowRelationshipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetFlowActionFlowRelationshipResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetFlowActionFlowRelationshipResponse := _GetFlowActionFlowRelationshipResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFlowActionFlowRelationshipResponse)

	if err != nil {
		return err
	}

	*o = GetFlowActionFlowRelationshipResponse(varGetFlowActionFlowRelationshipResponse)

	return err
}

type NullableGetFlowActionFlowRelationshipResponse struct {
	value *GetFlowActionFlowRelationshipResponse
	isSet bool
}

func (v NullableGetFlowActionFlowRelationshipResponse) Get() *GetFlowActionFlowRelationshipResponse {
	return v.value
}

func (v *NullableGetFlowActionFlowRelationshipResponse) Set(val *GetFlowActionFlowRelationshipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlowActionFlowRelationshipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlowActionFlowRelationshipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlowActionFlowRelationshipResponse(val *GetFlowActionFlowRelationshipResponse) *NullableGetFlowActionFlowRelationshipResponse {
	return &NullableGetFlowActionFlowRelationshipResponse{value: val, isSet: true}
}

func (v NullableGetFlowActionFlowRelationshipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlowActionFlowRelationshipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


