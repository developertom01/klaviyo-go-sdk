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

// checks if the AttributionResponseObjectResourceRelationshipsFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributionResponseObjectResourceRelationshipsFlow{}

// AttributionResponseObjectResourceRelationshipsFlow struct for AttributionResponseObjectResourceRelationshipsFlow
type AttributionResponseObjectResourceRelationshipsFlow struct {
	Data AttributionResponseObjectResourceRelationshipsFlowData `json:"data"`
}

type _AttributionResponseObjectResourceRelationshipsFlow AttributionResponseObjectResourceRelationshipsFlow

// NewAttributionResponseObjectResourceRelationshipsFlow instantiates a new AttributionResponseObjectResourceRelationshipsFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributionResponseObjectResourceRelationshipsFlow(data AttributionResponseObjectResourceRelationshipsFlowData) *AttributionResponseObjectResourceRelationshipsFlow {
	this := AttributionResponseObjectResourceRelationshipsFlow{}
	this.Data = data
	return &this
}

// NewAttributionResponseObjectResourceRelationshipsFlowWithDefaults instantiates a new AttributionResponseObjectResourceRelationshipsFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributionResponseObjectResourceRelationshipsFlowWithDefaults() *AttributionResponseObjectResourceRelationshipsFlow {
	this := AttributionResponseObjectResourceRelationshipsFlow{}
	return &this
}

// GetData returns the Data field value
func (o *AttributionResponseObjectResourceRelationshipsFlow) GetData() AttributionResponseObjectResourceRelationshipsFlowData {
	if o == nil {
		var ret AttributionResponseObjectResourceRelationshipsFlowData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AttributionResponseObjectResourceRelationshipsFlow) GetDataOk() (*AttributionResponseObjectResourceRelationshipsFlowData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AttributionResponseObjectResourceRelationshipsFlow) SetData(v AttributionResponseObjectResourceRelationshipsFlowData) {
	o.Data = v
}

func (o AttributionResponseObjectResourceRelationshipsFlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributionResponseObjectResourceRelationshipsFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AttributionResponseObjectResourceRelationshipsFlow) UnmarshalJSON(data []byte) (err error) {
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

	varAttributionResponseObjectResourceRelationshipsFlow := _AttributionResponseObjectResourceRelationshipsFlow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAttributionResponseObjectResourceRelationshipsFlow)

	if err != nil {
		return err
	}

	*o = AttributionResponseObjectResourceRelationshipsFlow(varAttributionResponseObjectResourceRelationshipsFlow)

	return err
}

type NullableAttributionResponseObjectResourceRelationshipsFlow struct {
	value *AttributionResponseObjectResourceRelationshipsFlow
	isSet bool
}

func (v NullableAttributionResponseObjectResourceRelationshipsFlow) Get() *AttributionResponseObjectResourceRelationshipsFlow {
	return v.value
}

func (v *NullableAttributionResponseObjectResourceRelationshipsFlow) Set(val *AttributionResponseObjectResourceRelationshipsFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributionResponseObjectResourceRelationshipsFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributionResponseObjectResourceRelationshipsFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributionResponseObjectResourceRelationshipsFlow(val *AttributionResponseObjectResourceRelationshipsFlow) *NullableAttributionResponseObjectResourceRelationshipsFlow {
	return &NullableAttributionResponseObjectResourceRelationshipsFlow{value: val, isSet: true}
}

func (v NullableAttributionResponseObjectResourceRelationshipsFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributionResponseObjectResourceRelationshipsFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

