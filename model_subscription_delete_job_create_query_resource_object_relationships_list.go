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

// checks if the SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList{}

// SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList struct for SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList
type SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList struct {
	Data SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsListData `json:"data"`
}

type _SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList

// NewSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList instantiates a new SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList(data SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsListData) *SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList {
	this := SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList{}
	this.Data = data
	return &this
}

// NewSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsListWithDefaults instantiates a new SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsListWithDefaults() *SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList {
	this := SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) GetData() SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsListData {
	if o == nil {
		var ret SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsListData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) GetDataOk() (*SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsListData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) SetData(v SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsListData) {
	o.Data = v
}

func (o SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList := _SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList)

	if err != nil {
		return err
	}

	*o = SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList(varSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList)

	return err
}

type NullableSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList struct {
	value *SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList
	isSet bool
}

func (v NullableSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) Get() *SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList {
	return v.value
}

func (v *NullableSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) Set(val *SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList(val *SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) *NullableSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList {
	return &NullableSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList{value: val, isSet: true}
}

func (v NullableSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


