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

// checks if the CampaignSendJobPartialUpdateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignSendJobPartialUpdateQueryResourceObjectAttributes{}

// CampaignSendJobPartialUpdateQueryResourceObjectAttributes struct for CampaignSendJobPartialUpdateQueryResourceObjectAttributes
type CampaignSendJobPartialUpdateQueryResourceObjectAttributes struct {
	// The action you would like to take with this send job from among 'cancel' and 'revert'
	Action string `json:"action"`
}

type _CampaignSendJobPartialUpdateQueryResourceObjectAttributes CampaignSendJobPartialUpdateQueryResourceObjectAttributes

// NewCampaignSendJobPartialUpdateQueryResourceObjectAttributes instantiates a new CampaignSendJobPartialUpdateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignSendJobPartialUpdateQueryResourceObjectAttributes(action string) *CampaignSendJobPartialUpdateQueryResourceObjectAttributes {
	this := CampaignSendJobPartialUpdateQueryResourceObjectAttributes{}
	this.Action = action
	return &this
}

// NewCampaignSendJobPartialUpdateQueryResourceObjectAttributesWithDefaults instantiates a new CampaignSendJobPartialUpdateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignSendJobPartialUpdateQueryResourceObjectAttributesWithDefaults() *CampaignSendJobPartialUpdateQueryResourceObjectAttributes {
	this := CampaignSendJobPartialUpdateQueryResourceObjectAttributes{}
	return &this
}

// GetAction returns the Action field value
func (o *CampaignSendJobPartialUpdateQueryResourceObjectAttributes) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *CampaignSendJobPartialUpdateQueryResourceObjectAttributes) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *CampaignSendJobPartialUpdateQueryResourceObjectAttributes) SetAction(v string) {
	o.Action = v
}

func (o CampaignSendJobPartialUpdateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignSendJobPartialUpdateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

func (o *CampaignSendJobPartialUpdateQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varCampaignSendJobPartialUpdateQueryResourceObjectAttributes := _CampaignSendJobPartialUpdateQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaignSendJobPartialUpdateQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = CampaignSendJobPartialUpdateQueryResourceObjectAttributes(varCampaignSendJobPartialUpdateQueryResourceObjectAttributes)

	return err
}

type NullableCampaignSendJobPartialUpdateQueryResourceObjectAttributes struct {
	value *CampaignSendJobPartialUpdateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableCampaignSendJobPartialUpdateQueryResourceObjectAttributes) Get() *CampaignSendJobPartialUpdateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableCampaignSendJobPartialUpdateQueryResourceObjectAttributes) Set(val *CampaignSendJobPartialUpdateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignSendJobPartialUpdateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignSendJobPartialUpdateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignSendJobPartialUpdateQueryResourceObjectAttributes(val *CampaignSendJobPartialUpdateQueryResourceObjectAttributes) *NullableCampaignSendJobPartialUpdateQueryResourceObjectAttributes {
	return &NullableCampaignSendJobPartialUpdateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableCampaignSendJobPartialUpdateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignSendJobPartialUpdateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


