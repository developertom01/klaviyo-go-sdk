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

// checks if the GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData{}

// GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData struct for GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData
type GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData struct {
	Type TemplateEnum `json:"type"`
	// The associated template id
	Id string `json:"id"`
}

type _GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData

// NewGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData instantiates a new GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData(type_ TemplateEnum, id string) *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData {
	this := GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateDataWithDefaults instantiates a new GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateDataWithDefaults() *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData {
	this := GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData{}
	return &this
}

// GetType returns the Type field value
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) GetType() TemplateEnum {
	if o == nil {
		var ret TemplateEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) GetTypeOk() (*TemplateEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) SetType(v TemplateEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) SetId(v string) {
	o.Id = v
}

func (o GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) UnmarshalJSON(data []byte) (err error) {
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

	varGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData := _GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData)

	if err != nil {
		return err
	}

	*o = GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData(varGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData)

	return err
}

type NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData struct {
	value *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData
	isSet bool
}

func (v NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) Get() *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData {
	return v.value
}

func (v *NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) Set(val *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData(val *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) *NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData {
	return &NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData{value: val, isSet: true}
}

func (v NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

