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

// checks if the GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions{}

// GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions struct for GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions
type GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions struct {
	Data []GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner `json:"data"`
	Links *RelationshipLinks `json:"links,omitempty"`
}

type _GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions

// NewGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions instantiates a new GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions(data []GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner) *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions {
	this := GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions{}
	this.Data = data
	return &this
}

// NewGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsWithDefaults instantiates a new GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsWithDefaults() *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions {
	this := GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions{}
	return &this
}

// GetData returns the Data field value
func (o *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) GetData() []GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner {
	if o == nil {
		var ret []GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) GetDataOk() ([]GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) SetData(v []GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

func (o GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) UnmarshalJSON(data []byte) (err error) {
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

	varGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions := _GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions)

	if err != nil {
		return err
	}

	*o = GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions(varGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions)

	return err
}

type NullableGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions struct {
	value *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions
	isSet bool
}

func (v NullableGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) Get() *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions {
	return v.value
}

func (v *NullableGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) Set(val *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions(val *GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) *NullableGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions {
	return &NullableGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions{value: val, isSet: true}
}

func (v NullableGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


