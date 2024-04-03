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

// checks if the CampaignRecipientEstimationResponseObjectResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignRecipientEstimationResponseObjectResourceAttributes{}

// CampaignRecipientEstimationResponseObjectResourceAttributes struct for CampaignRecipientEstimationResponseObjectResourceAttributes
type CampaignRecipientEstimationResponseObjectResourceAttributes struct {
	// The estimated number of unique recipients the campaign will send to
	EstimatedRecipientCount int32 `json:"estimated_recipient_count"`
}

type _CampaignRecipientEstimationResponseObjectResourceAttributes CampaignRecipientEstimationResponseObjectResourceAttributes

// NewCampaignRecipientEstimationResponseObjectResourceAttributes instantiates a new CampaignRecipientEstimationResponseObjectResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignRecipientEstimationResponseObjectResourceAttributes(estimatedRecipientCount int32) *CampaignRecipientEstimationResponseObjectResourceAttributes {
	this := CampaignRecipientEstimationResponseObjectResourceAttributes{}
	this.EstimatedRecipientCount = estimatedRecipientCount
	return &this
}

// NewCampaignRecipientEstimationResponseObjectResourceAttributesWithDefaults instantiates a new CampaignRecipientEstimationResponseObjectResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignRecipientEstimationResponseObjectResourceAttributesWithDefaults() *CampaignRecipientEstimationResponseObjectResourceAttributes {
	this := CampaignRecipientEstimationResponseObjectResourceAttributes{}
	return &this
}

// GetEstimatedRecipientCount returns the EstimatedRecipientCount field value
func (o *CampaignRecipientEstimationResponseObjectResourceAttributes) GetEstimatedRecipientCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EstimatedRecipientCount
}

// GetEstimatedRecipientCountOk returns a tuple with the EstimatedRecipientCount field value
// and a boolean to check if the value has been set.
func (o *CampaignRecipientEstimationResponseObjectResourceAttributes) GetEstimatedRecipientCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedRecipientCount, true
}

// SetEstimatedRecipientCount sets field value
func (o *CampaignRecipientEstimationResponseObjectResourceAttributes) SetEstimatedRecipientCount(v int32) {
	o.EstimatedRecipientCount = v
}

func (o CampaignRecipientEstimationResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignRecipientEstimationResponseObjectResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["estimated_recipient_count"] = o.EstimatedRecipientCount
	return toSerialize, nil
}

func (o *CampaignRecipientEstimationResponseObjectResourceAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"estimated_recipient_count",
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

	varCampaignRecipientEstimationResponseObjectResourceAttributes := _CampaignRecipientEstimationResponseObjectResourceAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaignRecipientEstimationResponseObjectResourceAttributes)

	if err != nil {
		return err
	}

	*o = CampaignRecipientEstimationResponseObjectResourceAttributes(varCampaignRecipientEstimationResponseObjectResourceAttributes)

	return err
}

type NullableCampaignRecipientEstimationResponseObjectResourceAttributes struct {
	value *CampaignRecipientEstimationResponseObjectResourceAttributes
	isSet bool
}

func (v NullableCampaignRecipientEstimationResponseObjectResourceAttributes) Get() *CampaignRecipientEstimationResponseObjectResourceAttributes {
	return v.value
}

func (v *NullableCampaignRecipientEstimationResponseObjectResourceAttributes) Set(val *CampaignRecipientEstimationResponseObjectResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignRecipientEstimationResponseObjectResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignRecipientEstimationResponseObjectResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignRecipientEstimationResponseObjectResourceAttributes(val *CampaignRecipientEstimationResponseObjectResourceAttributes) *NullableCampaignRecipientEstimationResponseObjectResourceAttributes {
	return &NullableCampaignRecipientEstimationResponseObjectResourceAttributes{value: val, isSet: true}
}

func (v NullableCampaignRecipientEstimationResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignRecipientEstimationResponseObjectResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


