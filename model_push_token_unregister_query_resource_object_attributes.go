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

// checks if the PushTokenUnregisterQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushTokenUnregisterQueryResourceObjectAttributes{}

// PushTokenUnregisterQueryResourceObjectAttributes struct for PushTokenUnregisterQueryResourceObjectAttributes
type PushTokenUnregisterQueryResourceObjectAttributes struct {
	// A push token from APNS or FCM.
	Token string `json:"token"`
	// The platform on which the push token was created.
	Platform string `json:"platform"`
	// The vendor of the push token.
	Vendor NullableString `json:"vendor,omitempty"`
	Profile PushTokenCreateQueryResourceObjectAttributesProfile `json:"profile"`
}

type _PushTokenUnregisterQueryResourceObjectAttributes PushTokenUnregisterQueryResourceObjectAttributes

// NewPushTokenUnregisterQueryResourceObjectAttributes instantiates a new PushTokenUnregisterQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushTokenUnregisterQueryResourceObjectAttributes(token string, platform string, profile PushTokenCreateQueryResourceObjectAttributesProfile) *PushTokenUnregisterQueryResourceObjectAttributes {
	this := PushTokenUnregisterQueryResourceObjectAttributes{}
	this.Token = token
	this.Platform = platform
	this.Profile = profile
	return &this
}

// NewPushTokenUnregisterQueryResourceObjectAttributesWithDefaults instantiates a new PushTokenUnregisterQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushTokenUnregisterQueryResourceObjectAttributesWithDefaults() *PushTokenUnregisterQueryResourceObjectAttributes {
	this := PushTokenUnregisterQueryResourceObjectAttributes{}
	return &this
}

// GetToken returns the Token field value
func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *PushTokenUnregisterQueryResourceObjectAttributes) SetToken(v string) {
	o.Token = v
}

// GetPlatform returns the Platform field value
func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *PushTokenUnregisterQueryResourceObjectAttributes) SetPlatform(v string) {
	o.Platform = v
}

// GetVendor returns the Vendor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetVendor() string {
	if o == nil || IsNil(o.Vendor.Get()) {
		var ret string
		return ret
	}
	return *o.Vendor.Get()
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vendor.Get(), o.Vendor.IsSet()
}

// HasVendor returns a boolean if a field has been set.
func (o *PushTokenUnregisterQueryResourceObjectAttributes) HasVendor() bool {
	if o != nil && o.Vendor.IsSet() {
		return true
	}

	return false
}

// SetVendor gets a reference to the given NullableString and assigns it to the Vendor field.
func (o *PushTokenUnregisterQueryResourceObjectAttributes) SetVendor(v string) {
	o.Vendor.Set(&v)
}
// SetVendorNil sets the value for Vendor to be an explicit nil
func (o *PushTokenUnregisterQueryResourceObjectAttributes) SetVendorNil() {
	o.Vendor.Set(nil)
}

// UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
func (o *PushTokenUnregisterQueryResourceObjectAttributes) UnsetVendor() {
	o.Vendor.Unset()
}

// GetProfile returns the Profile field value
func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetProfile() PushTokenCreateQueryResourceObjectAttributesProfile {
	if o == nil {
		var ret PushTokenCreateQueryResourceObjectAttributesProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetProfileOk() (*PushTokenCreateQueryResourceObjectAttributesProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *PushTokenUnregisterQueryResourceObjectAttributes) SetProfile(v PushTokenCreateQueryResourceObjectAttributesProfile) {
	o.Profile = v
}

func (o PushTokenUnregisterQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushTokenUnregisterQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	toSerialize["platform"] = o.Platform
	if o.Vendor.IsSet() {
		toSerialize["vendor"] = o.Vendor.Get()
	}
	toSerialize["profile"] = o.Profile
	return toSerialize, nil
}

func (o *PushTokenUnregisterQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
		"platform",
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

	varPushTokenUnregisterQueryResourceObjectAttributes := _PushTokenUnregisterQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPushTokenUnregisterQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = PushTokenUnregisterQueryResourceObjectAttributes(varPushTokenUnregisterQueryResourceObjectAttributes)

	return err
}

type NullablePushTokenUnregisterQueryResourceObjectAttributes struct {
	value *PushTokenUnregisterQueryResourceObjectAttributes
	isSet bool
}

func (v NullablePushTokenUnregisterQueryResourceObjectAttributes) Get() *PushTokenUnregisterQueryResourceObjectAttributes {
	return v.value
}

func (v *NullablePushTokenUnregisterQueryResourceObjectAttributes) Set(val *PushTokenUnregisterQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePushTokenUnregisterQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePushTokenUnregisterQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushTokenUnregisterQueryResourceObjectAttributes(val *PushTokenUnregisterQueryResourceObjectAttributes) *NullablePushTokenUnregisterQueryResourceObjectAttributes {
	return &NullablePushTokenUnregisterQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullablePushTokenUnregisterQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushTokenUnregisterQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


