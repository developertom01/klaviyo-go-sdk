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

// checks if the ListPartialUpdateQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPartialUpdateQuery{}

// ListPartialUpdateQuery struct for ListPartialUpdateQuery
type ListPartialUpdateQuery struct {
	Data ListPartialUpdateQueryResourceObject `json:"data"`
}

type _ListPartialUpdateQuery ListPartialUpdateQuery

// NewListPartialUpdateQuery instantiates a new ListPartialUpdateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPartialUpdateQuery(data ListPartialUpdateQueryResourceObject) *ListPartialUpdateQuery {
	this := ListPartialUpdateQuery{}
	this.Data = data
	return &this
}

// NewListPartialUpdateQueryWithDefaults instantiates a new ListPartialUpdateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPartialUpdateQueryWithDefaults() *ListPartialUpdateQuery {
	this := ListPartialUpdateQuery{}
	return &this
}

// GetData returns the Data field value
func (o *ListPartialUpdateQuery) GetData() ListPartialUpdateQueryResourceObject {
	if o == nil {
		var ret ListPartialUpdateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListPartialUpdateQuery) GetDataOk() (*ListPartialUpdateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ListPartialUpdateQuery) SetData(v ListPartialUpdateQueryResourceObject) {
	o.Data = v
}

func (o ListPartialUpdateQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPartialUpdateQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ListPartialUpdateQuery) UnmarshalJSON(data []byte) (err error) {
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

	varListPartialUpdateQuery := _ListPartialUpdateQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListPartialUpdateQuery)

	if err != nil {
		return err
	}

	*o = ListPartialUpdateQuery(varListPartialUpdateQuery)

	return err
}

type NullableListPartialUpdateQuery struct {
	value *ListPartialUpdateQuery
	isSet bool
}

func (v NullableListPartialUpdateQuery) Get() *ListPartialUpdateQuery {
	return v.value
}

func (v *NullableListPartialUpdateQuery) Set(val *ListPartialUpdateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableListPartialUpdateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableListPartialUpdateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPartialUpdateQuery(val *ListPartialUpdateQuery) *NullableListPartialUpdateQuery {
	return &NullableListPartialUpdateQuery{value: val, isSet: true}
}

func (v NullableListPartialUpdateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPartialUpdateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

