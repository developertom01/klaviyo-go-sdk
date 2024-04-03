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

// checks if the SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles{}

// SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles The profile(s) to remove suppressions for.
type SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles struct {
	Data []ProfileSuppressionDeleteQueryResourceObject `json:"data"`
}

type _SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles

// NewSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles instantiates a new SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles(data []ProfileSuppressionDeleteQueryResourceObject) *SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles {
	this := SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles{}
	this.Data = data
	return &this
}

// NewSuppressionDeleteJobCreateQueryResourceObjectAttributesProfilesWithDefaults instantiates a new SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuppressionDeleteJobCreateQueryResourceObjectAttributesProfilesWithDefaults() *SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles {
	this := SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles{}
	return &this
}

// GetData returns the Data field value
func (o *SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) GetData() []ProfileSuppressionDeleteQueryResourceObject {
	if o == nil {
		var ret []ProfileSuppressionDeleteQueryResourceObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) GetDataOk() ([]ProfileSuppressionDeleteQueryResourceObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) SetData(v []ProfileSuppressionDeleteQueryResourceObject) {
	o.Data = v
}

func (o SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) UnmarshalJSON(data []byte) (err error) {
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

	varSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles := _SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles)

	if err != nil {
		return err
	}

	*o = SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles(varSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles)

	return err
}

type NullableSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles struct {
	value *SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles
	isSet bool
}

func (v NullableSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) Get() *SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles {
	return v.value
}

func (v *NullableSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) Set(val *SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) {
	v.value = val
	v.isSet = true
}

func (v NullableSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) IsSet() bool {
	return v.isSet
}

func (v *NullableSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles(val *SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) *NullableSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles {
	return &NullableSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles{value: val, isSet: true}
}

func (v NullableSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


