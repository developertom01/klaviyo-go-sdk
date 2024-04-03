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

// checks if the GetFlowMessageResponseCompoundDocumentData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFlowMessageResponseCompoundDocumentData{}

// GetFlowMessageResponseCompoundDocumentData struct for GetFlowMessageResponseCompoundDocumentData
type GetFlowMessageResponseCompoundDocumentData struct {
	Type FlowMessageEnum `json:"type"`
	Id string `json:"id"`
	Attributes FlowMessageResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
	Relationships *GetFlowMessageResponseCompoundDocumentDataAllOfRelationships `json:"relationships,omitempty"`
}

type _GetFlowMessageResponseCompoundDocumentData GetFlowMessageResponseCompoundDocumentData

// NewGetFlowMessageResponseCompoundDocumentData instantiates a new GetFlowMessageResponseCompoundDocumentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlowMessageResponseCompoundDocumentData(type_ FlowMessageEnum, id string, attributes FlowMessageResponseObjectResourceAttributes, links ObjectLinks) *GetFlowMessageResponseCompoundDocumentData {
	this := GetFlowMessageResponseCompoundDocumentData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewGetFlowMessageResponseCompoundDocumentDataWithDefaults instantiates a new GetFlowMessageResponseCompoundDocumentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlowMessageResponseCompoundDocumentDataWithDefaults() *GetFlowMessageResponseCompoundDocumentData {
	this := GetFlowMessageResponseCompoundDocumentData{}
	return &this
}

// GetType returns the Type field value
func (o *GetFlowMessageResponseCompoundDocumentData) GetType() FlowMessageEnum {
	if o == nil {
		var ret FlowMessageEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetFlowMessageResponseCompoundDocumentData) GetTypeOk() (*FlowMessageEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetFlowMessageResponseCompoundDocumentData) SetType(v FlowMessageEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetFlowMessageResponseCompoundDocumentData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetFlowMessageResponseCompoundDocumentData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetFlowMessageResponseCompoundDocumentData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *GetFlowMessageResponseCompoundDocumentData) GetAttributes() FlowMessageResponseObjectResourceAttributes {
	if o == nil {
		var ret FlowMessageResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetFlowMessageResponseCompoundDocumentData) GetAttributesOk() (*FlowMessageResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GetFlowMessageResponseCompoundDocumentData) SetAttributes(v FlowMessageResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *GetFlowMessageResponseCompoundDocumentData) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetFlowMessageResponseCompoundDocumentData) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetFlowMessageResponseCompoundDocumentData) SetLinks(v ObjectLinks) {
	o.Links = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GetFlowMessageResponseCompoundDocumentData) GetRelationships() GetFlowMessageResponseCompoundDocumentDataAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetFlowMessageResponseCompoundDocumentDataAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlowMessageResponseCompoundDocumentData) GetRelationshipsOk() (*GetFlowMessageResponseCompoundDocumentDataAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GetFlowMessageResponseCompoundDocumentData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetFlowMessageResponseCompoundDocumentDataAllOfRelationships and assigns it to the Relationships field.
func (o *GetFlowMessageResponseCompoundDocumentData) SetRelationships(v GetFlowMessageResponseCompoundDocumentDataAllOfRelationships) {
	o.Relationships = &v
}

func (o GetFlowMessageResponseCompoundDocumentData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlowMessageResponseCompoundDocumentData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

func (o *GetFlowMessageResponseCompoundDocumentData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"attributes",
		"links",
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

	varGetFlowMessageResponseCompoundDocumentData := _GetFlowMessageResponseCompoundDocumentData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFlowMessageResponseCompoundDocumentData)

	if err != nil {
		return err
	}

	*o = GetFlowMessageResponseCompoundDocumentData(varGetFlowMessageResponseCompoundDocumentData)

	return err
}

type NullableGetFlowMessageResponseCompoundDocumentData struct {
	value *GetFlowMessageResponseCompoundDocumentData
	isSet bool
}

func (v NullableGetFlowMessageResponseCompoundDocumentData) Get() *GetFlowMessageResponseCompoundDocumentData {
	return v.value
}

func (v *NullableGetFlowMessageResponseCompoundDocumentData) Set(val *GetFlowMessageResponseCompoundDocumentData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlowMessageResponseCompoundDocumentData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlowMessageResponseCompoundDocumentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlowMessageResponseCompoundDocumentData(val *GetFlowMessageResponseCompoundDocumentData) *NullableGetFlowMessageResponseCompoundDocumentData {
	return &NullableGetFlowMessageResponseCompoundDocumentData{value: val, isSet: true}
}

func (v NullableGetFlowMessageResponseCompoundDocumentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlowMessageResponseCompoundDocumentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

