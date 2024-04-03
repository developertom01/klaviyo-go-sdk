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

// checks if the ProfileMergeQueryResourceObjectRelationshipsProfiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileMergeQueryResourceObjectRelationshipsProfiles{}

// ProfileMergeQueryResourceObjectRelationshipsProfiles struct for ProfileMergeQueryResourceObjectRelationshipsProfiles
type ProfileMergeQueryResourceObjectRelationshipsProfiles struct {
	Data []ProfileMergeQueryResourceObjectRelationshipsProfilesDataInner `json:"data"`
}

type _ProfileMergeQueryResourceObjectRelationshipsProfiles ProfileMergeQueryResourceObjectRelationshipsProfiles

// NewProfileMergeQueryResourceObjectRelationshipsProfiles instantiates a new ProfileMergeQueryResourceObjectRelationshipsProfiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileMergeQueryResourceObjectRelationshipsProfiles(data []ProfileMergeQueryResourceObjectRelationshipsProfilesDataInner) *ProfileMergeQueryResourceObjectRelationshipsProfiles {
	this := ProfileMergeQueryResourceObjectRelationshipsProfiles{}
	this.Data = data
	return &this
}

// NewProfileMergeQueryResourceObjectRelationshipsProfilesWithDefaults instantiates a new ProfileMergeQueryResourceObjectRelationshipsProfiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileMergeQueryResourceObjectRelationshipsProfilesWithDefaults() *ProfileMergeQueryResourceObjectRelationshipsProfiles {
	this := ProfileMergeQueryResourceObjectRelationshipsProfiles{}
	return &this
}

// GetData returns the Data field value
func (o *ProfileMergeQueryResourceObjectRelationshipsProfiles) GetData() []ProfileMergeQueryResourceObjectRelationshipsProfilesDataInner {
	if o == nil {
		var ret []ProfileMergeQueryResourceObjectRelationshipsProfilesDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ProfileMergeQueryResourceObjectRelationshipsProfiles) GetDataOk() ([]ProfileMergeQueryResourceObjectRelationshipsProfilesDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ProfileMergeQueryResourceObjectRelationshipsProfiles) SetData(v []ProfileMergeQueryResourceObjectRelationshipsProfilesDataInner) {
	o.Data = v
}

func (o ProfileMergeQueryResourceObjectRelationshipsProfiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileMergeQueryResourceObjectRelationshipsProfiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ProfileMergeQueryResourceObjectRelationshipsProfiles) UnmarshalJSON(data []byte) (err error) {
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

	varProfileMergeQueryResourceObjectRelationshipsProfiles := _ProfileMergeQueryResourceObjectRelationshipsProfiles{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProfileMergeQueryResourceObjectRelationshipsProfiles)

	if err != nil {
		return err
	}

	*o = ProfileMergeQueryResourceObjectRelationshipsProfiles(varProfileMergeQueryResourceObjectRelationshipsProfiles)

	return err
}

type NullableProfileMergeQueryResourceObjectRelationshipsProfiles struct {
	value *ProfileMergeQueryResourceObjectRelationshipsProfiles
	isSet bool
}

func (v NullableProfileMergeQueryResourceObjectRelationshipsProfiles) Get() *ProfileMergeQueryResourceObjectRelationshipsProfiles {
	return v.value
}

func (v *NullableProfileMergeQueryResourceObjectRelationshipsProfiles) Set(val *ProfileMergeQueryResourceObjectRelationshipsProfiles) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileMergeQueryResourceObjectRelationshipsProfiles) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileMergeQueryResourceObjectRelationshipsProfiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileMergeQueryResourceObjectRelationshipsProfiles(val *ProfileMergeQueryResourceObjectRelationshipsProfiles) *NullableProfileMergeQueryResourceObjectRelationshipsProfiles {
	return &NullableProfileMergeQueryResourceObjectRelationshipsProfiles{value: val, isSet: true}
}

func (v NullableProfileMergeQueryResourceObjectRelationshipsProfiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileMergeQueryResourceObjectRelationshipsProfiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


