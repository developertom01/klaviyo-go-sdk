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

// checks if the GetTagFlowRelationshipsResponseCollectionDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTagFlowRelationshipsResponseCollectionDataInner{}

// GetTagFlowRelationshipsResponseCollectionDataInner struct for GetTagFlowRelationshipsResponseCollectionDataInner
type GetTagFlowRelationshipsResponseCollectionDataInner struct {
	Type FlowEnum `json:"type"`
	// The IDs of all flows that are associated with the Tag
	Id string `json:"id"`
}

type _GetTagFlowRelationshipsResponseCollectionDataInner GetTagFlowRelationshipsResponseCollectionDataInner

// NewGetTagFlowRelationshipsResponseCollectionDataInner instantiates a new GetTagFlowRelationshipsResponseCollectionDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTagFlowRelationshipsResponseCollectionDataInner(type_ FlowEnum, id string) *GetTagFlowRelationshipsResponseCollectionDataInner {
	this := GetTagFlowRelationshipsResponseCollectionDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGetTagFlowRelationshipsResponseCollectionDataInnerWithDefaults instantiates a new GetTagFlowRelationshipsResponseCollectionDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTagFlowRelationshipsResponseCollectionDataInnerWithDefaults() *GetTagFlowRelationshipsResponseCollectionDataInner {
	this := GetTagFlowRelationshipsResponseCollectionDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GetTagFlowRelationshipsResponseCollectionDataInner) GetType() FlowEnum {
	if o == nil {
		var ret FlowEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetTagFlowRelationshipsResponseCollectionDataInner) GetTypeOk() (*FlowEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetTagFlowRelationshipsResponseCollectionDataInner) SetType(v FlowEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetTagFlowRelationshipsResponseCollectionDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetTagFlowRelationshipsResponseCollectionDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetTagFlowRelationshipsResponseCollectionDataInner) SetId(v string) {
	o.Id = v
}

func (o GetTagFlowRelationshipsResponseCollectionDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTagFlowRelationshipsResponseCollectionDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *GetTagFlowRelationshipsResponseCollectionDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varGetTagFlowRelationshipsResponseCollectionDataInner := _GetTagFlowRelationshipsResponseCollectionDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTagFlowRelationshipsResponseCollectionDataInner)

	if err != nil {
		return err
	}

	*o = GetTagFlowRelationshipsResponseCollectionDataInner(varGetTagFlowRelationshipsResponseCollectionDataInner)

	return err
}

type NullableGetTagFlowRelationshipsResponseCollectionDataInner struct {
	value *GetTagFlowRelationshipsResponseCollectionDataInner
	isSet bool
}

func (v NullableGetTagFlowRelationshipsResponseCollectionDataInner) Get() *GetTagFlowRelationshipsResponseCollectionDataInner {
	return v.value
}

func (v *NullableGetTagFlowRelationshipsResponseCollectionDataInner) Set(val *GetTagFlowRelationshipsResponseCollectionDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTagFlowRelationshipsResponseCollectionDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTagFlowRelationshipsResponseCollectionDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTagFlowRelationshipsResponseCollectionDataInner(val *GetTagFlowRelationshipsResponseCollectionDataInner) *NullableGetTagFlowRelationshipsResponseCollectionDataInner {
	return &NullableGetTagFlowRelationshipsResponseCollectionDataInner{value: val, isSet: true}
}

func (v NullableGetTagFlowRelationshipsResponseCollectionDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTagFlowRelationshipsResponseCollectionDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


