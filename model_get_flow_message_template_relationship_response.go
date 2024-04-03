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

// checks if the GetFlowMessageTemplateRelationshipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFlowMessageTemplateRelationshipResponse{}

// GetFlowMessageTemplateRelationshipResponse struct for GetFlowMessageTemplateRelationshipResponse
type GetFlowMessageTemplateRelationshipResponse struct {
	Data GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData `json:"data"`
}

type _GetFlowMessageTemplateRelationshipResponse GetFlowMessageTemplateRelationshipResponse

// NewGetFlowMessageTemplateRelationshipResponse instantiates a new GetFlowMessageTemplateRelationshipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlowMessageTemplateRelationshipResponse(data GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) *GetFlowMessageTemplateRelationshipResponse {
	this := GetFlowMessageTemplateRelationshipResponse{}
	this.Data = data
	return &this
}

// NewGetFlowMessageTemplateRelationshipResponseWithDefaults instantiates a new GetFlowMessageTemplateRelationshipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlowMessageTemplateRelationshipResponseWithDefaults() *GetFlowMessageTemplateRelationshipResponse {
	this := GetFlowMessageTemplateRelationshipResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetFlowMessageTemplateRelationshipResponse) GetData() GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData {
	if o == nil {
		var ret GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetFlowMessageTemplateRelationshipResponse) GetDataOk() (*GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetFlowMessageTemplateRelationshipResponse) SetData(v GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) {
	o.Data = v
}

func (o GetFlowMessageTemplateRelationshipResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlowMessageTemplateRelationshipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetFlowMessageTemplateRelationshipResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetFlowMessageTemplateRelationshipResponse := _GetFlowMessageTemplateRelationshipResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFlowMessageTemplateRelationshipResponse)

	if err != nil {
		return err
	}

	*o = GetFlowMessageTemplateRelationshipResponse(varGetFlowMessageTemplateRelationshipResponse)

	return err
}

type NullableGetFlowMessageTemplateRelationshipResponse struct {
	value *GetFlowMessageTemplateRelationshipResponse
	isSet bool
}

func (v NullableGetFlowMessageTemplateRelationshipResponse) Get() *GetFlowMessageTemplateRelationshipResponse {
	return v.value
}

func (v *NullableGetFlowMessageTemplateRelationshipResponse) Set(val *GetFlowMessageTemplateRelationshipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlowMessageTemplateRelationshipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlowMessageTemplateRelationshipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlowMessageTemplateRelationshipResponse(val *GetFlowMessageTemplateRelationshipResponse) *NullableGetFlowMessageTemplateRelationshipResponse {
	return &NullableGetFlowMessageTemplateRelationshipResponse{value: val, isSet: true}
}

func (v NullableGetFlowMessageTemplateRelationshipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlowMessageTemplateRelationshipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

