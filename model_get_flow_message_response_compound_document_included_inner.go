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

// GetFlowMessageResponseCompoundDocumentIncludedInner - struct for GetFlowMessageResponseCompoundDocumentIncludedInner
type GetFlowMessageResponseCompoundDocumentIncludedInner struct {
	FlowActionResponseObjectResource *FlowActionResponseObjectResource
	TemplateResponseObjectResource *TemplateResponseObjectResource
}

// FlowActionResponseObjectResourceAsGetFlowMessageResponseCompoundDocumentIncludedInner is a convenience function that returns FlowActionResponseObjectResource wrapped in GetFlowMessageResponseCompoundDocumentIncludedInner
func FlowActionResponseObjectResourceAsGetFlowMessageResponseCompoundDocumentIncludedInner(v *FlowActionResponseObjectResource) GetFlowMessageResponseCompoundDocumentIncludedInner {
	return GetFlowMessageResponseCompoundDocumentIncludedInner{
		FlowActionResponseObjectResource: v,
	}
}

// TemplateResponseObjectResourceAsGetFlowMessageResponseCompoundDocumentIncludedInner is a convenience function that returns TemplateResponseObjectResource wrapped in GetFlowMessageResponseCompoundDocumentIncludedInner
func TemplateResponseObjectResourceAsGetFlowMessageResponseCompoundDocumentIncludedInner(v *TemplateResponseObjectResource) GetFlowMessageResponseCompoundDocumentIncludedInner {
	return GetFlowMessageResponseCompoundDocumentIncludedInner{
		TemplateResponseObjectResource: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetFlowMessageResponseCompoundDocumentIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FlowActionResponseObjectResource
	err = newStrictDecoder(data).Decode(&dst.FlowActionResponseObjectResource)
	if err == nil {
		jsonFlowActionResponseObjectResource, _ := json.Marshal(dst.FlowActionResponseObjectResource)
		if string(jsonFlowActionResponseObjectResource) == "{}" { // empty struct
			dst.FlowActionResponseObjectResource = nil
		} else {
			match++
		}
	} else {
		dst.FlowActionResponseObjectResource = nil
	}

	// try to unmarshal data into TemplateResponseObjectResource
	err = newStrictDecoder(data).Decode(&dst.TemplateResponseObjectResource)
	if err == nil {
		jsonTemplateResponseObjectResource, _ := json.Marshal(dst.TemplateResponseObjectResource)
		if string(jsonTemplateResponseObjectResource) == "{}" { // empty struct
			dst.TemplateResponseObjectResource = nil
		} else {
			match++
		}
	} else {
		dst.TemplateResponseObjectResource = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.FlowActionResponseObjectResource = nil
		dst.TemplateResponseObjectResource = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetFlowMessageResponseCompoundDocumentIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetFlowMessageResponseCompoundDocumentIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetFlowMessageResponseCompoundDocumentIncludedInner) MarshalJSON() ([]byte, error) {
	if src.FlowActionResponseObjectResource != nil {
		return json.Marshal(&src.FlowActionResponseObjectResource)
	}

	if src.TemplateResponseObjectResource != nil {
		return json.Marshal(&src.TemplateResponseObjectResource)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetFlowMessageResponseCompoundDocumentIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.FlowActionResponseObjectResource != nil {
		return obj.FlowActionResponseObjectResource
	}

	if obj.TemplateResponseObjectResource != nil {
		return obj.TemplateResponseObjectResource
	}

	// all schemas are nil
	return nil
}

type NullableGetFlowMessageResponseCompoundDocumentIncludedInner struct {
	value *GetFlowMessageResponseCompoundDocumentIncludedInner
	isSet bool
}

func (v NullableGetFlowMessageResponseCompoundDocumentIncludedInner) Get() *GetFlowMessageResponseCompoundDocumentIncludedInner {
	return v.value
}

func (v *NullableGetFlowMessageResponseCompoundDocumentIncludedInner) Set(val *GetFlowMessageResponseCompoundDocumentIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlowMessageResponseCompoundDocumentIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlowMessageResponseCompoundDocumentIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlowMessageResponseCompoundDocumentIncludedInner(val *GetFlowMessageResponseCompoundDocumentIncludedInner) *NullableGetFlowMessageResponseCompoundDocumentIncludedInner {
	return &NullableGetFlowMessageResponseCompoundDocumentIncludedInner{value: val, isSet: true}
}

func (v NullableGetFlowMessageResponseCompoundDocumentIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlowMessageResponseCompoundDocumentIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


