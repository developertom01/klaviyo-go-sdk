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

// checks if the CampaignCreateQueryResourceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignCreateQueryResourceObject{}

// CampaignCreateQueryResourceObject struct for CampaignCreateQueryResourceObject
type CampaignCreateQueryResourceObject struct {
	Type CampaignEnum `json:"type"`
	Attributes CampaignCreateQueryResourceObjectAttributes `json:"attributes"`
}

type _CampaignCreateQueryResourceObject CampaignCreateQueryResourceObject

// NewCampaignCreateQueryResourceObject instantiates a new CampaignCreateQueryResourceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignCreateQueryResourceObject(type_ CampaignEnum, attributes CampaignCreateQueryResourceObjectAttributes) *CampaignCreateQueryResourceObject {
	this := CampaignCreateQueryResourceObject{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCampaignCreateQueryResourceObjectWithDefaults instantiates a new CampaignCreateQueryResourceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignCreateQueryResourceObjectWithDefaults() *CampaignCreateQueryResourceObject {
	this := CampaignCreateQueryResourceObject{}
	return &this
}

// GetType returns the Type field value
func (o *CampaignCreateQueryResourceObject) GetType() CampaignEnum {
	if o == nil {
		var ret CampaignEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CampaignCreateQueryResourceObject) GetTypeOk() (*CampaignEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CampaignCreateQueryResourceObject) SetType(v CampaignEnum) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CampaignCreateQueryResourceObject) GetAttributes() CampaignCreateQueryResourceObjectAttributes {
	if o == nil {
		var ret CampaignCreateQueryResourceObjectAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CampaignCreateQueryResourceObject) GetAttributesOk() (*CampaignCreateQueryResourceObjectAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CampaignCreateQueryResourceObject) SetAttributes(v CampaignCreateQueryResourceObjectAttributes) {
	o.Attributes = v
}

func (o CampaignCreateQueryResourceObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignCreateQueryResourceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *CampaignCreateQueryResourceObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
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

	varCampaignCreateQueryResourceObject := _CampaignCreateQueryResourceObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaignCreateQueryResourceObject)

	if err != nil {
		return err
	}

	*o = CampaignCreateQueryResourceObject(varCampaignCreateQueryResourceObject)

	return err
}

type NullableCampaignCreateQueryResourceObject struct {
	value *CampaignCreateQueryResourceObject
	isSet bool
}

func (v NullableCampaignCreateQueryResourceObject) Get() *CampaignCreateQueryResourceObject {
	return v.value
}

func (v *NullableCampaignCreateQueryResourceObject) Set(val *CampaignCreateQueryResourceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignCreateQueryResourceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignCreateQueryResourceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignCreateQueryResourceObject(val *CampaignCreateQueryResourceObject) *NullableCampaignCreateQueryResourceObject {
	return &NullableCampaignCreateQueryResourceObject{value: val, isSet: true}
}

func (v NullableCampaignCreateQueryResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignCreateQueryResourceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

