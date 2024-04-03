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

// checks if the GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData{}

// GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData struct for GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData
type GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData struct {
	Type CampaignEnum `json:"type"`
	// The parent campaign id
	Id string `json:"id"`
}

type _GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData

// NewGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData instantiates a new GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData(type_ CampaignEnum, id string) *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData {
	this := GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignDataWithDefaults instantiates a new GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignDataWithDefaults() *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData {
	this := GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData{}
	return &this
}

// GetType returns the Type field value
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) GetType() CampaignEnum {
	if o == nil {
		var ret CampaignEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) GetTypeOk() (*CampaignEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) SetType(v CampaignEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) SetId(v string) {
	o.Id = v
}

func (o GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData := _GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData)

	if err != nil {
		return err
	}

	*o = GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData(varGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData)

	return err
}

type NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData struct {
	value *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData
	isSet bool
}

func (v NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) Get() *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData {
	return v.value
}

func (v *NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) Set(val *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData(val *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) *NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData {
	return &NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData{value: val, isSet: true}
}

func (v NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

