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

// checks if the ServerBISSubscriptionCreateQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerBISSubscriptionCreateQuery{}

// ServerBISSubscriptionCreateQuery struct for ServerBISSubscriptionCreateQuery
type ServerBISSubscriptionCreateQuery struct {
	Data ServerBISSubscriptionCreateQueryResourceObject `json:"data"`
}

type _ServerBISSubscriptionCreateQuery ServerBISSubscriptionCreateQuery

// NewServerBISSubscriptionCreateQuery instantiates a new ServerBISSubscriptionCreateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerBISSubscriptionCreateQuery(data ServerBISSubscriptionCreateQueryResourceObject) *ServerBISSubscriptionCreateQuery {
	this := ServerBISSubscriptionCreateQuery{}
	this.Data = data
	return &this
}

// NewServerBISSubscriptionCreateQueryWithDefaults instantiates a new ServerBISSubscriptionCreateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerBISSubscriptionCreateQueryWithDefaults() *ServerBISSubscriptionCreateQuery {
	this := ServerBISSubscriptionCreateQuery{}
	return &this
}

// GetData returns the Data field value
func (o *ServerBISSubscriptionCreateQuery) GetData() ServerBISSubscriptionCreateQueryResourceObject {
	if o == nil {
		var ret ServerBISSubscriptionCreateQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ServerBISSubscriptionCreateQuery) GetDataOk() (*ServerBISSubscriptionCreateQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ServerBISSubscriptionCreateQuery) SetData(v ServerBISSubscriptionCreateQueryResourceObject) {
	o.Data = v
}

func (o ServerBISSubscriptionCreateQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerBISSubscriptionCreateQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ServerBISSubscriptionCreateQuery) UnmarshalJSON(data []byte) (err error) {
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

	varServerBISSubscriptionCreateQuery := _ServerBISSubscriptionCreateQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServerBISSubscriptionCreateQuery)

	if err != nil {
		return err
	}

	*o = ServerBISSubscriptionCreateQuery(varServerBISSubscriptionCreateQuery)

	return err
}

type NullableServerBISSubscriptionCreateQuery struct {
	value *ServerBISSubscriptionCreateQuery
	isSet bool
}

func (v NullableServerBISSubscriptionCreateQuery) Get() *ServerBISSubscriptionCreateQuery {
	return v.value
}

func (v *NullableServerBISSubscriptionCreateQuery) Set(val *ServerBISSubscriptionCreateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableServerBISSubscriptionCreateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableServerBISSubscriptionCreateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerBISSubscriptionCreateQuery(val *ServerBISSubscriptionCreateQuery) *NullableServerBISSubscriptionCreateQuery {
	return &NullableServerBISSubscriptionCreateQuery{value: val, isSet: true}
}

func (v NullableServerBISSubscriptionCreateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerBISSubscriptionCreateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


