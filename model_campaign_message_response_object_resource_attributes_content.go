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
	"fmt"
)

// CampaignMessageResponseObjectResourceAttributesContent - Additional attributes relating to the content of the message
type CampaignMessageResponseObjectResourceAttributesContent struct {
	EmailContentSubObject *EmailContentSubObject
	SMSContentSubObject *SMSContentSubObject
}

// EmailContentSubObjectAsCampaignMessageResponseObjectResourceAttributesContent is a convenience function that returns EmailContentSubObject wrapped in CampaignMessageResponseObjectResourceAttributesContent
func EmailContentSubObjectAsCampaignMessageResponseObjectResourceAttributesContent(v *EmailContentSubObject) CampaignMessageResponseObjectResourceAttributesContent {
	return CampaignMessageResponseObjectResourceAttributesContent{
		EmailContentSubObject: v,
	}
}

// SMSContentSubObjectAsCampaignMessageResponseObjectResourceAttributesContent is a convenience function that returns SMSContentSubObject wrapped in CampaignMessageResponseObjectResourceAttributesContent
func SMSContentSubObjectAsCampaignMessageResponseObjectResourceAttributesContent(v *SMSContentSubObject) CampaignMessageResponseObjectResourceAttributesContent {
	return CampaignMessageResponseObjectResourceAttributesContent{
		SMSContentSubObject: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CampaignMessageResponseObjectResourceAttributesContent) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EmailContentSubObject
	err = newStrictDecoder(data).Decode(&dst.EmailContentSubObject)
	if err == nil {
		jsonEmailContentSubObject, _ := json.Marshal(dst.EmailContentSubObject)
		if string(jsonEmailContentSubObject) == "{}" { // empty struct
			dst.EmailContentSubObject = nil
		} else {
			match++
		}
	} else {
		dst.EmailContentSubObject = nil
	}

	// try to unmarshal data into SMSContentSubObject
	err = newStrictDecoder(data).Decode(&dst.SMSContentSubObject)
	if err == nil {
		jsonSMSContentSubObject, _ := json.Marshal(dst.SMSContentSubObject)
		if string(jsonSMSContentSubObject) == "{}" { // empty struct
			dst.SMSContentSubObject = nil
		} else {
			match++
		}
	} else {
		dst.SMSContentSubObject = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EmailContentSubObject = nil
		dst.SMSContentSubObject = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CampaignMessageResponseObjectResourceAttributesContent)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CampaignMessageResponseObjectResourceAttributesContent)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CampaignMessageResponseObjectResourceAttributesContent) MarshalJSON() ([]byte, error) {
	if src.EmailContentSubObject != nil {
		return json.Marshal(&src.EmailContentSubObject)
	}

	if src.SMSContentSubObject != nil {
		return json.Marshal(&src.SMSContentSubObject)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CampaignMessageResponseObjectResourceAttributesContent) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EmailContentSubObject != nil {
		return obj.EmailContentSubObject
	}

	if obj.SMSContentSubObject != nil {
		return obj.SMSContentSubObject
	}

	// all schemas are nil
	return nil
}

type NullableCampaignMessageResponseObjectResourceAttributesContent struct {
	value *CampaignMessageResponseObjectResourceAttributesContent
	isSet bool
}

func (v NullableCampaignMessageResponseObjectResourceAttributesContent) Get() *CampaignMessageResponseObjectResourceAttributesContent {
	return v.value
}

func (v *NullableCampaignMessageResponseObjectResourceAttributesContent) Set(val *CampaignMessageResponseObjectResourceAttributesContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignMessageResponseObjectResourceAttributesContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignMessageResponseObjectResourceAttributesContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignMessageResponseObjectResourceAttributesContent(val *CampaignMessageResponseObjectResourceAttributesContent) *NullableCampaignMessageResponseObjectResourceAttributesContent {
	return &NullableCampaignMessageResponseObjectResourceAttributesContent{value: val, isSet: true}
}

func (v NullableCampaignMessageResponseObjectResourceAttributesContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignMessageResponseObjectResourceAttributesContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


