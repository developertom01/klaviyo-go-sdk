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
)

// checks if the GetFlowMessageResponseCompoundDocumentDataAllOfRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFlowMessageResponseCompoundDocumentDataAllOfRelationships{}

// GetFlowMessageResponseCompoundDocumentDataAllOfRelationships struct for GetFlowMessageResponseCompoundDocumentDataAllOfRelationships
type GetFlowMessageResponseCompoundDocumentDataAllOfRelationships struct {
	FlowAction *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction `json:"flow-action,omitempty"`
	Template *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate `json:"template,omitempty"`
}

// NewGetFlowMessageResponseCompoundDocumentDataAllOfRelationships instantiates a new GetFlowMessageResponseCompoundDocumentDataAllOfRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlowMessageResponseCompoundDocumentDataAllOfRelationships() *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships {
	this := GetFlowMessageResponseCompoundDocumentDataAllOfRelationships{}
	return &this
}

// NewGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsWithDefaults instantiates a new GetFlowMessageResponseCompoundDocumentDataAllOfRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsWithDefaults() *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships {
	this := GetFlowMessageResponseCompoundDocumentDataAllOfRelationships{}
	return &this
}

// GetFlowAction returns the FlowAction field value if set, zero value otherwise.
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships) GetFlowAction() GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction {
	if o == nil || IsNil(o.FlowAction) {
		var ret GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction
		return ret
	}
	return *o.FlowAction
}

// GetFlowActionOk returns a tuple with the FlowAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships) GetFlowActionOk() (*GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction, bool) {
	if o == nil || IsNil(o.FlowAction) {
		return nil, false
	}
	return o.FlowAction, true
}

// HasFlowAction returns a boolean if a field has been set.
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships) HasFlowAction() bool {
	if o != nil && !IsNil(o.FlowAction) {
		return true
	}

	return false
}

// SetFlowAction gets a reference to the given GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction and assigns it to the FlowAction field.
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships) SetFlowAction(v GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) {
	o.FlowAction = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships) GetTemplate() GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate {
	if o == nil || IsNil(o.Template) {
		var ret GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships) GetTemplateOk() (*GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate and assigns it to the Template field.
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships) SetTemplate(v GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate) {
	o.Template = &v
}

func (o GetFlowMessageResponseCompoundDocumentDataAllOfRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlowMessageResponseCompoundDocumentDataAllOfRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlowAction) {
		toSerialize["flow-action"] = o.FlowAction
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return toSerialize, nil
}

type NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationships struct {
	value *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships
	isSet bool
}

func (v NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationships) Get() *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships {
	return v.value
}

func (v *NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationships) Set(val *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationships(val *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships) *NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationships {
	return &NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationships{value: val, isSet: true}
}

func (v NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

