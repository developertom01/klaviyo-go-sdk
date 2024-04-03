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

// checks if the SegmentPartialUpdateQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentPartialUpdateQuery{}

// SegmentPartialUpdateQuery struct for SegmentPartialUpdateQuery
type SegmentPartialUpdateQuery struct {
	Data SegmentPartialUpdateQueryResourceObject `json:"data"`
}

type _SegmentPartialUpdateQuery SegmentPartialUpdateQuery

// NewSegmentPartialUpdateQuery instantiates a new SegmentPartialUpdateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentPartialUpdateQuery(data SegmentPartialUpdateQueryResourceObject) *SegmentPartialUpdateQuery {
	this := SegmentPartialUpdateQuery{}
	this.Data = data
	return &this
}

// NewSegmentPartialUpdateQueryWithDefaults instantiates a new SegmentPartialUpdateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentPartialUpdateQueryWithDefaults() *SegmentPartialUpdateQuery {
	this := SegmentPartialUpdateQuery{}
	return &this
}

// GetData returns the Data field value
func (o *SegmentPartialUpdateQuery) GetData() SegmentPartialUpdateQueryResourceObject {
	if o == nil {
		var ret SegmentPartialUpdateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SegmentPartialUpdateQuery) GetDataOk() (*SegmentPartialUpdateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SegmentPartialUpdateQuery) SetData(v SegmentPartialUpdateQueryResourceObject) {
	o.Data = v
}

func (o SegmentPartialUpdateQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentPartialUpdateQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SegmentPartialUpdateQuery) UnmarshalJSON(data []byte) (err error) {
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

	varSegmentPartialUpdateQuery := _SegmentPartialUpdateQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSegmentPartialUpdateQuery)

	if err != nil {
		return err
	}

	*o = SegmentPartialUpdateQuery(varSegmentPartialUpdateQuery)

	return err
}

type NullableSegmentPartialUpdateQuery struct {
	value *SegmentPartialUpdateQuery
	isSet bool
}

func (v NullableSegmentPartialUpdateQuery) Get() *SegmentPartialUpdateQuery {
	return v.value
}

func (v *NullableSegmentPartialUpdateQuery) Set(val *SegmentPartialUpdateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentPartialUpdateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentPartialUpdateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentPartialUpdateQuery(val *SegmentPartialUpdateQuery) *NullableSegmentPartialUpdateQuery {
	return &NullableSegmentPartialUpdateQuery{value: val, isSet: true}
}

func (v NullableSegmentPartialUpdateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentPartialUpdateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

