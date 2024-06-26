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

// checks if the CampaignRecipientEstimationJobResponseObjectResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignRecipientEstimationJobResponseObjectResourceAttributes{}

// CampaignRecipientEstimationJobResponseObjectResourceAttributes struct for CampaignRecipientEstimationJobResponseObjectResourceAttributes
type CampaignRecipientEstimationJobResponseObjectResourceAttributes struct {
	// The status of the recipient estimation job
	Status string `json:"status"`
}

type _CampaignRecipientEstimationJobResponseObjectResourceAttributes CampaignRecipientEstimationJobResponseObjectResourceAttributes

// NewCampaignRecipientEstimationJobResponseObjectResourceAttributes instantiates a new CampaignRecipientEstimationJobResponseObjectResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignRecipientEstimationJobResponseObjectResourceAttributes(status string) *CampaignRecipientEstimationJobResponseObjectResourceAttributes {
	this := CampaignRecipientEstimationJobResponseObjectResourceAttributes{}
	this.Status = status
	return &this
}

// NewCampaignRecipientEstimationJobResponseObjectResourceAttributesWithDefaults instantiates a new CampaignRecipientEstimationJobResponseObjectResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignRecipientEstimationJobResponseObjectResourceAttributesWithDefaults() *CampaignRecipientEstimationJobResponseObjectResourceAttributes {
	this := CampaignRecipientEstimationJobResponseObjectResourceAttributes{}
	return &this
}

// GetStatus returns the Status field value
func (o *CampaignRecipientEstimationJobResponseObjectResourceAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CampaignRecipientEstimationJobResponseObjectResourceAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CampaignRecipientEstimationJobResponseObjectResourceAttributes) SetStatus(v string) {
	o.Status = v
}

func (o CampaignRecipientEstimationJobResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignRecipientEstimationJobResponseObjectResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *CampaignRecipientEstimationJobResponseObjectResourceAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varCampaignRecipientEstimationJobResponseObjectResourceAttributes := _CampaignRecipientEstimationJobResponseObjectResourceAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaignRecipientEstimationJobResponseObjectResourceAttributes)

	if err != nil {
		return err
	}

	*o = CampaignRecipientEstimationJobResponseObjectResourceAttributes(varCampaignRecipientEstimationJobResponseObjectResourceAttributes)

	return err
}

type NullableCampaignRecipientEstimationJobResponseObjectResourceAttributes struct {
	value *CampaignRecipientEstimationJobResponseObjectResourceAttributes
	isSet bool
}

func (v NullableCampaignRecipientEstimationJobResponseObjectResourceAttributes) Get() *CampaignRecipientEstimationJobResponseObjectResourceAttributes {
	return v.value
}

func (v *NullableCampaignRecipientEstimationJobResponseObjectResourceAttributes) Set(val *CampaignRecipientEstimationJobResponseObjectResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignRecipientEstimationJobResponseObjectResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignRecipientEstimationJobResponseObjectResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignRecipientEstimationJobResponseObjectResourceAttributes(val *CampaignRecipientEstimationJobResponseObjectResourceAttributes) *NullableCampaignRecipientEstimationJobResponseObjectResourceAttributes {
	return &NullableCampaignRecipientEstimationJobResponseObjectResourceAttributes{value: val, isSet: true}
}

func (v NullableCampaignRecipientEstimationJobResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignRecipientEstimationJobResponseObjectResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


