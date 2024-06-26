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

// CampaignValuesRequestDTOResourceObjectAttributesTimeframe - The timeframe to query for data within. The max length a timeframe can be is 1 year
type CampaignValuesRequestDTOResourceObjectAttributesTimeframe struct {
	CustomTimeframe *CustomTimeframe
	Timeframe *Timeframe
}

// CustomTimeframeAsCampaignValuesRequestDTOResourceObjectAttributesTimeframe is a convenience function that returns CustomTimeframe wrapped in CampaignValuesRequestDTOResourceObjectAttributesTimeframe
func CustomTimeframeAsCampaignValuesRequestDTOResourceObjectAttributesTimeframe(v *CustomTimeframe) CampaignValuesRequestDTOResourceObjectAttributesTimeframe {
	return CampaignValuesRequestDTOResourceObjectAttributesTimeframe{
		CustomTimeframe: v,
	}
}

// TimeframeAsCampaignValuesRequestDTOResourceObjectAttributesTimeframe is a convenience function that returns Timeframe wrapped in CampaignValuesRequestDTOResourceObjectAttributesTimeframe
func TimeframeAsCampaignValuesRequestDTOResourceObjectAttributesTimeframe(v *Timeframe) CampaignValuesRequestDTOResourceObjectAttributesTimeframe {
	return CampaignValuesRequestDTOResourceObjectAttributesTimeframe{
		Timeframe: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CampaignValuesRequestDTOResourceObjectAttributesTimeframe) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomTimeframe
	err = newStrictDecoder(data).Decode(&dst.CustomTimeframe)
	if err == nil {
		jsonCustomTimeframe, _ := json.Marshal(dst.CustomTimeframe)
		if string(jsonCustomTimeframe) == "{}" { // empty struct
			dst.CustomTimeframe = nil
		} else {
			match++
		}
	} else {
		dst.CustomTimeframe = nil
	}

	// try to unmarshal data into Timeframe
	err = newStrictDecoder(data).Decode(&dst.Timeframe)
	if err == nil {
		jsonTimeframe, _ := json.Marshal(dst.Timeframe)
		if string(jsonTimeframe) == "{}" { // empty struct
			dst.Timeframe = nil
		} else {
			match++
		}
	} else {
		dst.Timeframe = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CustomTimeframe = nil
		dst.Timeframe = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CampaignValuesRequestDTOResourceObjectAttributesTimeframe)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CampaignValuesRequestDTOResourceObjectAttributesTimeframe)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CampaignValuesRequestDTOResourceObjectAttributesTimeframe) MarshalJSON() ([]byte, error) {
	if src.CustomTimeframe != nil {
		return json.Marshal(&src.CustomTimeframe)
	}

	if src.Timeframe != nil {
		return json.Marshal(&src.Timeframe)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CampaignValuesRequestDTOResourceObjectAttributesTimeframe) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CustomTimeframe != nil {
		return obj.CustomTimeframe
	}

	if obj.Timeframe != nil {
		return obj.Timeframe
	}

	// all schemas are nil
	return nil
}

type NullableCampaignValuesRequestDTOResourceObjectAttributesTimeframe struct {
	value *CampaignValuesRequestDTOResourceObjectAttributesTimeframe
	isSet bool
}

func (v NullableCampaignValuesRequestDTOResourceObjectAttributesTimeframe) Get() *CampaignValuesRequestDTOResourceObjectAttributesTimeframe {
	return v.value
}

func (v *NullableCampaignValuesRequestDTOResourceObjectAttributesTimeframe) Set(val *CampaignValuesRequestDTOResourceObjectAttributesTimeframe) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignValuesRequestDTOResourceObjectAttributesTimeframe) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignValuesRequestDTOResourceObjectAttributesTimeframe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignValuesRequestDTOResourceObjectAttributesTimeframe(val *CampaignValuesRequestDTOResourceObjectAttributesTimeframe) *NullableCampaignValuesRequestDTOResourceObjectAttributesTimeframe {
	return &NullableCampaignValuesRequestDTOResourceObjectAttributesTimeframe{value: val, isSet: true}
}

func (v NullableCampaignValuesRequestDTOResourceObjectAttributesTimeframe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignValuesRequestDTOResourceObjectAttributesTimeframe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


