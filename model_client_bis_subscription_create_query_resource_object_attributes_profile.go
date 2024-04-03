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

// checks if the ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile{}

// ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile struct for ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile
type ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile struct {
	Data ProfileIdentifierDTOResourceObject `json:"data"`
}

type _ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile

// NewClientBISSubscriptionCreateQueryResourceObjectAttributesProfile instantiates a new ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientBISSubscriptionCreateQueryResourceObjectAttributesProfile(data ProfileIdentifierDTOResourceObject) *ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile {
	this := ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile{}
	this.Data = data
	return &this
}

// NewClientBISSubscriptionCreateQueryResourceObjectAttributesProfileWithDefaults instantiates a new ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientBISSubscriptionCreateQueryResourceObjectAttributesProfileWithDefaults() *ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile {
	this := ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile{}
	return &this
}

// GetData returns the Data field value
func (o *ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) GetData() ProfileIdentifierDTOResourceObject {
	if o == nil {
		var ret ProfileIdentifierDTOResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) GetDataOk() (*ProfileIdentifierDTOResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) SetData(v ProfileIdentifierDTOResourceObject) {
	o.Data = v
}

func (o ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) UnmarshalJSON(data []byte) (err error) {
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

	varClientBISSubscriptionCreateQueryResourceObjectAttributesProfile := _ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClientBISSubscriptionCreateQueryResourceObjectAttributesProfile)

	if err != nil {
		return err
	}

	*o = ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile(varClientBISSubscriptionCreateQueryResourceObjectAttributesProfile)

	return err
}

type NullableClientBISSubscriptionCreateQueryResourceObjectAttributesProfile struct {
	value *ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile
	isSet bool
}

func (v NullableClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) Get() *ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile {
	return v.value
}

func (v *NullableClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) Set(val *ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientBISSubscriptionCreateQueryResourceObjectAttributesProfile(val *ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) *NullableClientBISSubscriptionCreateQueryResourceObjectAttributesProfile {
	return &NullableClientBISSubscriptionCreateQueryResourceObjectAttributesProfile{value: val, isSet: true}
}

func (v NullableClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientBISSubscriptionCreateQueryResourceObjectAttributesProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

