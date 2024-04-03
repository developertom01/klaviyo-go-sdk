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

// checks if the GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction{}

// GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction struct for GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction
type GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction struct {
	Data GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner `json:"data"`
	Links *RelationshipLinks `json:"links,omitempty"`
}

type _GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction

// NewGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction instantiates a new GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction(data GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner) *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction {
	this := GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction{}
	this.Data = data
	return &this
}

// NewGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowActionWithDefaults instantiates a new GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowActionWithDefaults() *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction {
	this := GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction{}
	return &this
}

// GetData returns the Data field value
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) GetData() GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner {
	if o == nil {
		var ret GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) GetDataOk() (*GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) SetData(v GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

func (o GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction := _GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction)

	if err != nil {
		return err
	}

	*o = GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction(varGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction)

	return err
}

type NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction struct {
	value *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction
	isSet bool
}

func (v NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) Get() *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction {
	return v.value
}

func (v *NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) Set(val *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction(val *GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) *NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction {
	return &NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction{value: val, isSet: true}
}

func (v NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


