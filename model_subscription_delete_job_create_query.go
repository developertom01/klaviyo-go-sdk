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

// checks if the SubscriptionDeleteJobCreateQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionDeleteJobCreateQuery{}

// SubscriptionDeleteJobCreateQuery struct for SubscriptionDeleteJobCreateQuery
type SubscriptionDeleteJobCreateQuery struct {
	Data SubscriptionDeleteJobCreateQueryResourceObject `json:"data"`
}

type _SubscriptionDeleteJobCreateQuery SubscriptionDeleteJobCreateQuery

// NewSubscriptionDeleteJobCreateQuery instantiates a new SubscriptionDeleteJobCreateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDeleteJobCreateQuery(data SubscriptionDeleteJobCreateQueryResourceObject) *SubscriptionDeleteJobCreateQuery {
	this := SubscriptionDeleteJobCreateQuery{}
	this.Data = data
	return &this
}

// NewSubscriptionDeleteJobCreateQueryWithDefaults instantiates a new SubscriptionDeleteJobCreateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDeleteJobCreateQueryWithDefaults() *SubscriptionDeleteJobCreateQuery {
	this := SubscriptionDeleteJobCreateQuery{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionDeleteJobCreateQuery) GetData() SubscriptionDeleteJobCreateQueryResourceObject {
	if o == nil {
		var ret SubscriptionDeleteJobCreateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDeleteJobCreateQuery) GetDataOk() (*SubscriptionDeleteJobCreateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionDeleteJobCreateQuery) SetData(v SubscriptionDeleteJobCreateQueryResourceObject) {
	o.Data = v
}

func (o SubscriptionDeleteJobCreateQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionDeleteJobCreateQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SubscriptionDeleteJobCreateQuery) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionDeleteJobCreateQuery := _SubscriptionDeleteJobCreateQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionDeleteJobCreateQuery)

	if err != nil {
		return err
	}

	*o = SubscriptionDeleteJobCreateQuery(varSubscriptionDeleteJobCreateQuery)

	return err
}

type NullableSubscriptionDeleteJobCreateQuery struct {
	value *SubscriptionDeleteJobCreateQuery
	isSet bool
}

func (v NullableSubscriptionDeleteJobCreateQuery) Get() *SubscriptionDeleteJobCreateQuery {
	return v.value
}

func (v *NullableSubscriptionDeleteJobCreateQuery) Set(val *SubscriptionDeleteJobCreateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDeleteJobCreateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDeleteJobCreateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDeleteJobCreateQuery(val *SubscriptionDeleteJobCreateQuery) *NullableSubscriptionDeleteJobCreateQuery {
	return &NullableSubscriptionDeleteJobCreateQuery{value: val, isSet: true}
}

func (v NullableSubscriptionDeleteJobCreateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDeleteJobCreateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


