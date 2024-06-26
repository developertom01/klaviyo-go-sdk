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

// checks if the SubscriptionDeleteJobCreateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionDeleteJobCreateQueryResourceObjectAttributes{}

// SubscriptionDeleteJobCreateQueryResourceObjectAttributes struct for SubscriptionDeleteJobCreateQueryResourceObjectAttributes
type SubscriptionDeleteJobCreateQueryResourceObjectAttributes struct {
	Profiles SubscriptionDeleteJobCreateQueryResourceObjectAttributesProfiles `json:"profiles"`
}

type _SubscriptionDeleteJobCreateQueryResourceObjectAttributes SubscriptionDeleteJobCreateQueryResourceObjectAttributes

// NewSubscriptionDeleteJobCreateQueryResourceObjectAttributes instantiates a new SubscriptionDeleteJobCreateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDeleteJobCreateQueryResourceObjectAttributes(profiles SubscriptionDeleteJobCreateQueryResourceObjectAttributesProfiles) *SubscriptionDeleteJobCreateQueryResourceObjectAttributes {
	this := SubscriptionDeleteJobCreateQueryResourceObjectAttributes{}
	this.Profiles = profiles
	return &this
}

// NewSubscriptionDeleteJobCreateQueryResourceObjectAttributesWithDefaults instantiates a new SubscriptionDeleteJobCreateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDeleteJobCreateQueryResourceObjectAttributesWithDefaults() *SubscriptionDeleteJobCreateQueryResourceObjectAttributes {
	this := SubscriptionDeleteJobCreateQueryResourceObjectAttributes{}
	return &this
}

// GetProfiles returns the Profiles field value
func (o *SubscriptionDeleteJobCreateQueryResourceObjectAttributes) GetProfiles() SubscriptionDeleteJobCreateQueryResourceObjectAttributesProfiles {
	if o == nil {
		var ret SubscriptionDeleteJobCreateQueryResourceObjectAttributesProfiles
		return ret
	}

	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDeleteJobCreateQueryResourceObjectAttributes) GetProfilesOk() (*SubscriptionDeleteJobCreateQueryResourceObjectAttributesProfiles, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// SetProfiles sets field value
func (o *SubscriptionDeleteJobCreateQueryResourceObjectAttributes) SetProfiles(v SubscriptionDeleteJobCreateQueryResourceObjectAttributesProfiles) {
	o.Profiles = v
}

func (o SubscriptionDeleteJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionDeleteJobCreateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profiles"] = o.Profiles
	return toSerialize, nil
}

func (o *SubscriptionDeleteJobCreateQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profiles",
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

	varSubscriptionDeleteJobCreateQueryResourceObjectAttributes := _SubscriptionDeleteJobCreateQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionDeleteJobCreateQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = SubscriptionDeleteJobCreateQueryResourceObjectAttributes(varSubscriptionDeleteJobCreateQueryResourceObjectAttributes)

	return err
}

type NullableSubscriptionDeleteJobCreateQueryResourceObjectAttributes struct {
	value *SubscriptionDeleteJobCreateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableSubscriptionDeleteJobCreateQueryResourceObjectAttributes) Get() *SubscriptionDeleteJobCreateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableSubscriptionDeleteJobCreateQueryResourceObjectAttributes) Set(val *SubscriptionDeleteJobCreateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDeleteJobCreateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDeleteJobCreateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDeleteJobCreateQueryResourceObjectAttributes(val *SubscriptionDeleteJobCreateQueryResourceObjectAttributes) *NullableSubscriptionDeleteJobCreateQueryResourceObjectAttributes {
	return &NullableSubscriptionDeleteJobCreateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionDeleteJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDeleteJobCreateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


