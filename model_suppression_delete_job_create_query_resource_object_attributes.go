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

// checks if the SuppressionDeleteJobCreateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuppressionDeleteJobCreateQueryResourceObjectAttributes{}

// SuppressionDeleteJobCreateQueryResourceObjectAttributes struct for SuppressionDeleteJobCreateQueryResourceObjectAttributes
type SuppressionDeleteJobCreateQueryResourceObjectAttributes struct {
	Profiles SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles `json:"profiles"`
}

type _SuppressionDeleteJobCreateQueryResourceObjectAttributes SuppressionDeleteJobCreateQueryResourceObjectAttributes

// NewSuppressionDeleteJobCreateQueryResourceObjectAttributes instantiates a new SuppressionDeleteJobCreateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuppressionDeleteJobCreateQueryResourceObjectAttributes(profiles SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) *SuppressionDeleteJobCreateQueryResourceObjectAttributes {
	this := SuppressionDeleteJobCreateQueryResourceObjectAttributes{}
	this.Profiles = profiles
	return &this
}

// NewSuppressionDeleteJobCreateQueryResourceObjectAttributesWithDefaults instantiates a new SuppressionDeleteJobCreateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuppressionDeleteJobCreateQueryResourceObjectAttributesWithDefaults() *SuppressionDeleteJobCreateQueryResourceObjectAttributes {
	this := SuppressionDeleteJobCreateQueryResourceObjectAttributes{}
	return &this
}

// GetProfiles returns the Profiles field value
func (o *SuppressionDeleteJobCreateQueryResourceObjectAttributes) GetProfiles() SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles {
	if o == nil {
		var ret SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles
		return ret
	}

	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value
// and a boolean to check if the value has been set.
func (o *SuppressionDeleteJobCreateQueryResourceObjectAttributes) GetProfilesOk() (*SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// SetProfiles sets field value
func (o *SuppressionDeleteJobCreateQueryResourceObjectAttributes) SetProfiles(v SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) {
	o.Profiles = v
}

func (o SuppressionDeleteJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuppressionDeleteJobCreateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profiles"] = o.Profiles
	return toSerialize, nil
}

func (o *SuppressionDeleteJobCreateQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
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

	varSuppressionDeleteJobCreateQueryResourceObjectAttributes := _SuppressionDeleteJobCreateQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuppressionDeleteJobCreateQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = SuppressionDeleteJobCreateQueryResourceObjectAttributes(varSuppressionDeleteJobCreateQueryResourceObjectAttributes)

	return err
}

type NullableSuppressionDeleteJobCreateQueryResourceObjectAttributes struct {
	value *SuppressionDeleteJobCreateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableSuppressionDeleteJobCreateQueryResourceObjectAttributes) Get() *SuppressionDeleteJobCreateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableSuppressionDeleteJobCreateQueryResourceObjectAttributes) Set(val *SuppressionDeleteJobCreateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSuppressionDeleteJobCreateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSuppressionDeleteJobCreateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuppressionDeleteJobCreateQueryResourceObjectAttributes(val *SuppressionDeleteJobCreateQueryResourceObjectAttributes) *NullableSuppressionDeleteJobCreateQueryResourceObjectAttributes {
	return &NullableSuppressionDeleteJobCreateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableSuppressionDeleteJobCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuppressionDeleteJobCreateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


