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

// checks if the DataPrivacyCreateDeletionJobQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataPrivacyCreateDeletionJobQueryResourceObjectAttributes{}

// DataPrivacyCreateDeletionJobQueryResourceObjectAttributes struct for DataPrivacyCreateDeletionJobQueryResourceObjectAttributes
type DataPrivacyCreateDeletionJobQueryResourceObjectAttributes struct {
	Profile DataPrivacyCreateDeletionJobQueryResourceObjectAttributesProfile `json:"profile"`
}

type _DataPrivacyCreateDeletionJobQueryResourceObjectAttributes DataPrivacyCreateDeletionJobQueryResourceObjectAttributes

// NewDataPrivacyCreateDeletionJobQueryResourceObjectAttributes instantiates a new DataPrivacyCreateDeletionJobQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPrivacyCreateDeletionJobQueryResourceObjectAttributes(profile DataPrivacyCreateDeletionJobQueryResourceObjectAttributesProfile) *DataPrivacyCreateDeletionJobQueryResourceObjectAttributes {
	this := DataPrivacyCreateDeletionJobQueryResourceObjectAttributes{}
	this.Profile = profile
	return &this
}

// NewDataPrivacyCreateDeletionJobQueryResourceObjectAttributesWithDefaults instantiates a new DataPrivacyCreateDeletionJobQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPrivacyCreateDeletionJobQueryResourceObjectAttributesWithDefaults() *DataPrivacyCreateDeletionJobQueryResourceObjectAttributes {
	this := DataPrivacyCreateDeletionJobQueryResourceObjectAttributes{}
	return &this
}

// GetProfile returns the Profile field value
func (o *DataPrivacyCreateDeletionJobQueryResourceObjectAttributes) GetProfile() DataPrivacyCreateDeletionJobQueryResourceObjectAttributesProfile {
	if o == nil {
		var ret DataPrivacyCreateDeletionJobQueryResourceObjectAttributesProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *DataPrivacyCreateDeletionJobQueryResourceObjectAttributes) GetProfileOk() (*DataPrivacyCreateDeletionJobQueryResourceObjectAttributesProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *DataPrivacyCreateDeletionJobQueryResourceObjectAttributes) SetProfile(v DataPrivacyCreateDeletionJobQueryResourceObjectAttributesProfile) {
	o.Profile = v
}

func (o DataPrivacyCreateDeletionJobQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataPrivacyCreateDeletionJobQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profile"] = o.Profile
	return toSerialize, nil
}

func (o *DataPrivacyCreateDeletionJobQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profile",
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

	varDataPrivacyCreateDeletionJobQueryResourceObjectAttributes := _DataPrivacyCreateDeletionJobQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataPrivacyCreateDeletionJobQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = DataPrivacyCreateDeletionJobQueryResourceObjectAttributes(varDataPrivacyCreateDeletionJobQueryResourceObjectAttributes)

	return err
}

type NullableDataPrivacyCreateDeletionJobQueryResourceObjectAttributes struct {
	value *DataPrivacyCreateDeletionJobQueryResourceObjectAttributes
	isSet bool
}

func (v NullableDataPrivacyCreateDeletionJobQueryResourceObjectAttributes) Get() *DataPrivacyCreateDeletionJobQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableDataPrivacyCreateDeletionJobQueryResourceObjectAttributes) Set(val *DataPrivacyCreateDeletionJobQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPrivacyCreateDeletionJobQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPrivacyCreateDeletionJobQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPrivacyCreateDeletionJobQueryResourceObjectAttributes(val *DataPrivacyCreateDeletionJobQueryResourceObjectAttributes) *NullableDataPrivacyCreateDeletionJobQueryResourceObjectAttributes {
	return &NullableDataPrivacyCreateDeletionJobQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableDataPrivacyCreateDeletionJobQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPrivacyCreateDeletionJobQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


