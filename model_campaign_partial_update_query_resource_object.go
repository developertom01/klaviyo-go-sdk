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

// checks if the CampaignPartialUpdateQueryResourceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignPartialUpdateQueryResourceObject{}

// CampaignPartialUpdateQueryResourceObject struct for CampaignPartialUpdateQueryResourceObject
type CampaignPartialUpdateQueryResourceObject struct {
	Type CampaignEnum `json:"type"`
	// The campaign ID to be retrieved
	Id string `json:"id"`
	Attributes CampaignPartialUpdateQueryResourceObjectAttributes `json:"attributes"`
}

type _CampaignPartialUpdateQueryResourceObject CampaignPartialUpdateQueryResourceObject

// NewCampaignPartialUpdateQueryResourceObject instantiates a new CampaignPartialUpdateQueryResourceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignPartialUpdateQueryResourceObject(type_ CampaignEnum, id string, attributes CampaignPartialUpdateQueryResourceObjectAttributes) *CampaignPartialUpdateQueryResourceObject {
	this := CampaignPartialUpdateQueryResourceObject{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCampaignPartialUpdateQueryResourceObjectWithDefaults instantiates a new CampaignPartialUpdateQueryResourceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignPartialUpdateQueryResourceObjectWithDefaults() *CampaignPartialUpdateQueryResourceObject {
	this := CampaignPartialUpdateQueryResourceObject{}
	return &this
}

// GetType returns the Type field value
func (o *CampaignPartialUpdateQueryResourceObject) GetType() CampaignEnum {
	if o == nil {
		var ret CampaignEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CampaignPartialUpdateQueryResourceObject) GetTypeOk() (*CampaignEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CampaignPartialUpdateQueryResourceObject) SetType(v CampaignEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CampaignPartialUpdateQueryResourceObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CampaignPartialUpdateQueryResourceObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CampaignPartialUpdateQueryResourceObject) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CampaignPartialUpdateQueryResourceObject) GetAttributes() CampaignPartialUpdateQueryResourceObjectAttributes {
	if o == nil {
		var ret CampaignPartialUpdateQueryResourceObjectAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CampaignPartialUpdateQueryResourceObject) GetAttributesOk() (*CampaignPartialUpdateQueryResourceObjectAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CampaignPartialUpdateQueryResourceObject) SetAttributes(v CampaignPartialUpdateQueryResourceObjectAttributes) {
	o.Attributes = v
}

func (o CampaignPartialUpdateQueryResourceObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignPartialUpdateQueryResourceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *CampaignPartialUpdateQueryResourceObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varCampaignPartialUpdateQueryResourceObject := _CampaignPartialUpdateQueryResourceObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaignPartialUpdateQueryResourceObject)

	if err != nil {
		return err
	}

	*o = CampaignPartialUpdateQueryResourceObject(varCampaignPartialUpdateQueryResourceObject)

	return err
}

type NullableCampaignPartialUpdateQueryResourceObject struct {
	value *CampaignPartialUpdateQueryResourceObject
	isSet bool
}

func (v NullableCampaignPartialUpdateQueryResourceObject) Get() *CampaignPartialUpdateQueryResourceObject {
	return v.value
}

func (v *NullableCampaignPartialUpdateQueryResourceObject) Set(val *CampaignPartialUpdateQueryResourceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignPartialUpdateQueryResourceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignPartialUpdateQueryResourceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignPartialUpdateQueryResourceObject(val *CampaignPartialUpdateQueryResourceObject) *NullableCampaignPartialUpdateQueryResourceObject {
	return &NullableCampaignPartialUpdateQueryResourceObject{value: val, isSet: true}
}

func (v NullableCampaignPartialUpdateQueryResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignPartialUpdateQueryResourceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

