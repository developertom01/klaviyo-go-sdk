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

// checks if the AttributionResponseObjectResourceRelationshipsEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributionResponseObjectResourceRelationshipsEventData{}

// AttributionResponseObjectResourceRelationshipsEventData struct for AttributionResponseObjectResourceRelationshipsEventData
type AttributionResponseObjectResourceRelationshipsEventData struct {
	Type EventEnum `json:"type"`
	// Event
	Id string `json:"id"`
}

type _AttributionResponseObjectResourceRelationshipsEventData AttributionResponseObjectResourceRelationshipsEventData

// NewAttributionResponseObjectResourceRelationshipsEventData instantiates a new AttributionResponseObjectResourceRelationshipsEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributionResponseObjectResourceRelationshipsEventData(type_ EventEnum, id string) *AttributionResponseObjectResourceRelationshipsEventData {
	this := AttributionResponseObjectResourceRelationshipsEventData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAttributionResponseObjectResourceRelationshipsEventDataWithDefaults instantiates a new AttributionResponseObjectResourceRelationshipsEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributionResponseObjectResourceRelationshipsEventDataWithDefaults() *AttributionResponseObjectResourceRelationshipsEventData {
	this := AttributionResponseObjectResourceRelationshipsEventData{}
	return &this
}

// GetType returns the Type field value
func (o *AttributionResponseObjectResourceRelationshipsEventData) GetType() EventEnum {
	if o == nil {
		var ret EventEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AttributionResponseObjectResourceRelationshipsEventData) GetTypeOk() (*EventEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AttributionResponseObjectResourceRelationshipsEventData) SetType(v EventEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AttributionResponseObjectResourceRelationshipsEventData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttributionResponseObjectResourceRelationshipsEventData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttributionResponseObjectResourceRelationshipsEventData) SetId(v string) {
	o.Id = v
}

func (o AttributionResponseObjectResourceRelationshipsEventData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributionResponseObjectResourceRelationshipsEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AttributionResponseObjectResourceRelationshipsEventData) UnmarshalJSON(data []byte) (err error) {
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

	varAttributionResponseObjectResourceRelationshipsEventData := _AttributionResponseObjectResourceRelationshipsEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAttributionResponseObjectResourceRelationshipsEventData)

	if err != nil {
		return err
	}

	*o = AttributionResponseObjectResourceRelationshipsEventData(varAttributionResponseObjectResourceRelationshipsEventData)

	return err
}

type NullableAttributionResponseObjectResourceRelationshipsEventData struct {
	value *AttributionResponseObjectResourceRelationshipsEventData
	isSet bool
}

func (v NullableAttributionResponseObjectResourceRelationshipsEventData) Get() *AttributionResponseObjectResourceRelationshipsEventData {
	return v.value
}

func (v *NullableAttributionResponseObjectResourceRelationshipsEventData) Set(val *AttributionResponseObjectResourceRelationshipsEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributionResponseObjectResourceRelationshipsEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributionResponseObjectResourceRelationshipsEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributionResponseObjectResourceRelationshipsEventData(val *AttributionResponseObjectResourceRelationshipsEventData) *NullableAttributionResponseObjectResourceRelationshipsEventData {
	return &NullableAttributionResponseObjectResourceRelationshipsEventData{value: val, isSet: true}
}

func (v NullableAttributionResponseObjectResourceRelationshipsEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributionResponseObjectResourceRelationshipsEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

